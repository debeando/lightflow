package plugins

import (
	"github.com/debeando/lightflow/common"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/template"
	"github.com/debeando/lightflow/plugins/mysql"
	"github.com/debeando/lightflow/variables"
)

type PluginMySQL struct {
	Config    mysql.MySQL
	Variables variables.List
}

func (p *PluginMySQL) Retrieve(fn func(rowCount int, columns []string, row []string) bool) {
	if !p.isValid() {
		return
	}

	p.Variables = *variables.Load()

	p.Variables.Set(map[string]interface{}{
		"mysql_query": p.QueryRendered(),
	})

	mysql := mysql.MySQL{
		Host:     p.Render(p.Host()),
		Port:     p.Port(),
		User:     p.Render(p.User()),
		Password: p.Render(p.Password()),
		Schema:   p.Render(p.Schema()),
		Query:    p.QueryRendered(),
	}

	err := mysql.Execute(func(rowCount int, columns []string, row []string) bool {
		p.Variables.Set(map[string]interface{}{
			"mysql_rows_count": rowCount,
		})

		if rowCount == 1 {
			for k, v := range row {
				p.Variables.Set(map[string]interface{}{
					columns[k]: string(v),
				})
			}
		}
		return fn(rowCount, columns, row)
	})
	if err != nil {
		p.Variables.Set(map[string]interface{}{"exit_code": 1})
		log.Error(err.Error(), nil)
	}
}

func (p *PluginMySQL) isValid() bool {
	query := p.Query()

	if len(query) == 0 {
		return false
	}

	return true
}

func (p *PluginMySQL) Query() string {
	query := p.Config.Query

	if len(query) == 0 {
		return common.InterfaceToString(
			p.Variables.Get("mysql_query"),
		)
	}

	return query
}

func (p *PluginMySQL) QueryRendered() string {
	var query = p.Query()
	var vars = p.Variables.GetItems()

	// Find unknown variables:
	for _, variable := range template.Variables(query) {
		if p.Variables.Exist(variable) == false {
			p.Variables.Set(map[string]interface{}{variable: ""})
		}
	}

	// Find template variables to render:
	for variable, value := range p.Variables.Items {
		value_template := template.Variables(common.InterfaceToString(value))

		if len(value_template) > 0 {
			query, err := template.Render(common.InterfaceToString(value), vars)
			if err != nil {
				log.Warning(err.Error(), nil)
			}

			vars[variable] = query
		}
	}

	// Render template:
	query, err := template.Render(query, vars)
	if err != nil {
		log.Warning(err.Error(), nil)
	}

	return common.TrimNewlines(query)
}

func (p *PluginMySQL) Host() string {
	host := p.Config.Host

	if len(host) == 0 {
		return common.InterfaceToString(
			p.Variables.Get("mysql_host"),
		)
	}

	return host
}

func (p *PluginMySQL) Port() int {
	port := p.Config.Port

	if port == 0 {
		port = common.StringToInt(
			common.InterfaceToString(
				p.Variables.Get("mysql_port"),
			),
		)

		if port == 0 {
			return 3306
		}
	}

	return port
}

func (p *PluginMySQL) User() string {
	user := p.Config.User

	if len(user) == 0 {
		return common.InterfaceToString(
			p.Variables.Get("mysql_user"),
		)
	}

	return user
}

func (p *PluginMySQL) Password() string {
	password := p.Config.Password

	if len(password) == 0 {
		return common.InterfaceToString(
			p.Variables.Get("mysql_password"),
		)
	}

	return password
}

func (p *PluginMySQL) Schema() string {
	schema := p.Config.Schema

	if len(schema) == 0 {
		return common.InterfaceToString(
			p.Variables.Get("mysql_schema"),
		)
	}

	return schema
}

func (c *PluginMySQL) Render(s string) string {
	r, err := template.Render(s, c.Variables.GetItems())
	if err != nil {
		log.Warning(err.Error(), nil)
	}
	return r
}
