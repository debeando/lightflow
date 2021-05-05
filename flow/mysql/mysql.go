package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/debeando/lightflow/flow/mysql/csv"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	Host         string
	Port         int
	User         string
	Password     string
	Schema       string
	Query        string
	Header       bool
	Path         string
	Connection  *sql.DB
	Columns    []string
	Row          map[string]string
}

func (m *MySQL) Run() (int, map[string]string, error) {
	var err error
	var export bool
	var rowsCount int

	oneRow := map[string]string{}

	file := csv.CSV {
		Path: m.Path,
	}

	if err := file.IsValidPath(); err != nil {
		if err.Error() != "File name is empty." {
			return rowsCount, nil, err
		}
	} else {
		if err := file.Create(); err != nil {
			return rowsCount, nil, err
		} else {
			export = true
		}		
	}

	if err := m.connect(); err != nil {
		return rowsCount, nil, err
	}

	if ! export {
		rowsCount, oneRow, err = m.row()
	} else {
		chCSV := make(chan []string)
		defer close(chCSV)

		go func() {
			err = file.Write(chCSV)
			if err != nil {
				panic(err)
			}
		}()

		rowsCount, err = m.dump(chCSV)
	}

	if err != nil {
		return rowsCount, nil, err
	}

	return rowsCount, oneRow, nil
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

func (m *MySQL) execute(fn func(int,[]string,[]string) bool) (err error) {
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

func (m *MySQL) row() (int, map[string]string, error) {
	oneRow := map[string]string{}
	rowsCount := 0

	err := m.execute(func(rowCount int, columns []string, row []string) bool {
		for k, v := range columns {
			oneRow[v] = string(row[k])
		}
		rowsCount = rowCount
		return true
	})

	return rowsCount, oneRow, err
}

func (m *MySQL) dump(chOut chan<- []string) (int, error) {
	var rowsCount int

	err := m.execute(func(rowCount int, columns []string, row []string) bool {
		rowsCount = rowCount

		if rowsCount == 1 {
			chOut <- columns
		}

		chOut <- row

		return false
	})

	return rowsCount, err
}
