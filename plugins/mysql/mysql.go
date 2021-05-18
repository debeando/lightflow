package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Schema     string `yaml:"schema"`
	Query      string `yaml:"query"`
	Connection *sql.DB
}

func (m *MySQL) Execute(fn func(int, []string, []string) bool) error {
	var err error

	if err = m.connect(); err != nil {
		return err
	}

	err = m.execute(func(rowCount int, columns []string, row []string) bool {
		return fn(rowCount, columns, row)
	})

	if err != nil {
		return err
	}

	return nil
}

func (m *MySQL) connect() error {
	if m.Connection != nil {
		return nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&multiStatements=true",
		m.User,
		m.Password,
		m.Host,
		m.Port,
		m.Schema,
	)

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = conn.Ping()
	if err != nil {
		return err
	}

	m.Connection = conn

	return nil
}

func (m *MySQL) execute(fn func(int, []string, []string) bool) (err error) {
	var rowCount int
	var columns []string
	var row []string

	if len(m.Query) == 0 {
		return errors.New("Query is empty.")
	}

	rows, err := m.Connection.Query(
		"SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED; " +
			"SET SQL_BUFFER_RESULT=1; " +
			m.Query)
	if err != nil {
		return err
	}
	defer rows.Close()

	columns, err = rows.Columns()
	if err != nil {
		return err
	}

	// Values: all values of a row. Put each field of each row into values.
	// Values length == number of columns
	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		row = nil
		rowCount += 1

		//Add the contents of each line to scanArgs, and also to values:
		err = rows.Scan(scanArgs...)
		if err != nil {
			return err
		}

		for _, v := range values {
			row = append(row, string(v))
		}

		if fn(rowCount, columns, row) {
			break
		}
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}
