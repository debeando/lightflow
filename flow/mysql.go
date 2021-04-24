package flow

import (
	"fmt"

	"github.com/debeando/lightflow/common"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/mysql"
	"github.com/debeando/lightflow/flow/template"
)

func (f *Flow) mysql() {
	if ! f.isValidMySQL() {
		return
	}

	mysql := mysql.MySQL{
		Host:     f.Render(f.GetMySQLHost()),
		Port:     f.GetMySQLPort(),
		User:     f.Render(f.GetMySQLUser()),
		Password: f.Render(f.GetMySQLPassword()),
		Schema:   f.Render(f.GetMySQLSchema()),
		Query:    f.renderQuery(),
		Header:   f.GetMySQLHeader(),
		Path:     f.Render(f.GetMySQLPath()),
	}

	log.Debug(
		fmt.Sprintf(
			"%s/%s/%s Query %s",
			f.TaskName(),
			f.SubTaskName(),
			f.PipeName(),
			mysql.Query,
		),
		nil,
	)

	rows_count, row, err := mysql.Execute()
	if err != nil {
		f.Variables.Set(map[string]interface{}{"exit_code": 1})
		log.Error(err.Error(), nil)
	}

	f.Variables.Set(map[string]interface{}{
		"mysql_rows_count": rows_count,
	})

	for k,v := range row {
		f.Variables.Set(map[string]interface{}{
			k: v,
		})
	}
}

func (f *Flow) isValidMySQL() bool {
	query := f.GetMySQLQuery()

	if len(query) == 0 {
		return false
	}

	return true
}

func (f *Flow) renderQuery() string {
	var cmd = f.GetMySQLQuery()
	var vars = f.Variables.GetItems()

	// Find unknown variables:
	for _, variable := range template.Variables(cmd) {
		if f.Variables.Exist(variable) == false {
			f.Variables.Set(map[string]interface{}{variable: ""})
		}
	}

	// Find template variables to render:
	for variable, value := range f.Variables.Items {
		value_template := template.Variables(common.InterfaceToString(value))

		if len(value_template) > 0 {
			cmd, err := template.Render(common.InterfaceToString(value), vars)
			if err != nil {
				log.Warning(err.Error(), nil)
			}

			vars[variable] = cmd
		}
	}

	// Render template:
	cmd, err := template.Render(cmd, vars)
	if err != nil {
		log.Warning(err.Error(), nil)
	}

	return common.TrimNewlines(cmd)
}

func (f *Flow) GetMySQLHost() string {
	host := f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].MySQL.Host

	if len(host) == 0 {
		return common.InterfaceToString(f.Variables.Get("mysql_host"))
	}

	return host

}

func (f *Flow) GetMySQLPort() int {
	port := f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].MySQL.Port

	if port == 0 {
		port = common.StringToInt(
			common.InterfaceToString(
				f.Variables.Get("mysql_port"),
			),
		)

		if port == 0 {
			return 3306
		}
	}

	return port
}

func (f *Flow) GetMySQLUser() string {
	user := f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].MySQL.User

	if len(user) == 0 {
		return common.InterfaceToString(f.Variables.Get("mysql_user"))
	}

	return user
}

func (f *Flow) GetMySQLPassword() string {
	password := f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].MySQL.Password

	if len(password) == 0 {
		return common.InterfaceToString(f.Variables.Get("mysql_password"))
	}

	return password
}

func (f *Flow) GetMySQLSchema() string {
	schema := f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].MySQL.Schema

	if len(schema) == 0 {
		return common.InterfaceToString(f.Variables.Get("mysql_schema"))
	}

	return schema
}

func (f *Flow) GetMySQLQuery() string {
	query := f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].MySQL.Query

	if len(query) == 0 {
		return common.InterfaceToString(f.Variables.Get("mysql_query"))
	}

	return query
}

func (f *Flow) GetMySQLHeader() bool {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].MySQL.Header
}

func (f *Flow) GetMySQLPath() string {
	path := f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].MySQL.Path

	if len(path) == 0 {
		return common.InterfaceToString(f.Variables.Get("mysql_path"))
	}

	return path
}
