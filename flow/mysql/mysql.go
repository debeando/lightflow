package mysql

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	Host     string
	Port     int
	User     string
	Password string
	Schema   string
	Query    string
	Header   bool
	Path     string
}

func (m *MySQL) Execute() (int, map[string]string, error) {
	var err error
	var RowsCount int
	var row []string
	oneRow := map[string]string{}

	ch := make(chan []string)
	defer close(ch)

	if len(m.Query) == 0 {
		return 0, nil, nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&multiStatements=true",
		m.User,
		m.Password,
		m.Host,
		m.Port,
		m.Schema,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return 0, nil, err
	}

	err = db.Ping()
	if err != nil {
		return 0, nil, err
	}

	rows, err := db.Query(
		"SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED; " +
		"SET SQL_BUFFER_RESULT=1; " +
		m.Query)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return 0, nil, err
	}

	//Values: all values of a row. Put each field of each row into values. Values length = = number of columns
	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	if err = m.ValidPath(); err != nil {
		return 0, nil, err
	}

	go m.SaveCSVRow(columns, ch)

	for rows.Next() {
		row = nil
		RowsCount += 1

		//Add the contents of each line to scanArgs, and also to values:
		err = rows.Scan(scanArgs...)
		if err != nil {
			return 0, nil, err
		}

		for _, v := range values {
			row = append(row, string(v))
		}

		// Send row to save into file:
		ch <- row
	}

	if len(values) > 0 && len(row) == 1 && RowsCount == 1 {
		for k, v := range columns {
			oneRow[v] = string(values[k])
		}
	}

	if err = rows.Err(); err != nil {
		return 0, nil, err
	}

	return RowsCount, oneRow, nil
}

func (m *MySQL) SaveCSVRow(columns []string, ch <-chan []string) error {
	f, err := os.Create(m.Path)
	defer f.Close()
	if err != nil {
		return err
	}

	w := csv.NewWriter(f)

	if m.Header {
		w.Write(columns)
	}

	for row := range ch {
		w.Write(row)
	}

	w.Flush()
	return nil
}

func (m *MySQL) ValidPath() error {
	dir, err := os.Stat(m.Path)

	if err != nil {
		return errors.New(fmt.Sprintf("Failed to open file: %s", m.Path))
	}

	if dir.IsDir() {
		return errors.New(fmt.Sprintf("The path is a directory and is not valid: %s", m.Path))
	}

	if filepath.Ext(m.Path) != ".csv" {
		return errors.New(fmt.Sprintf("File extension ins't equal to .csv: %s", m.Path))
	}

	return err
}
