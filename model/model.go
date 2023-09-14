package model

import (
	"database/sql"
	"encoding/json"
	"reflect"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type IModel interface {
	ConnectToDB(rawConfigFile json.RawMessage) error
	SendResult() ([]byte, error)
	Read(tableName string) ([][]string, error)
	ExecuteQuery(query string) ([][]string, error)
}

type Config struct {
	Username string `JSON:"username"`
	Password string `JSON:"password"`
}

type Model struct {
	DB              *sql.DB
	StatusConnected bool
}

func (m *Model) ConnectToDB(rawConfigFile json.RawMessage) error {
	var cf Config
	err := json.Unmarshal(rawConfigFile, &cf)
	if err != nil {
		return err
	}
	connectStr := "user=" + cf.Username + " dbname=postgres password=" + cf.Password + " sslmode=disable"
	db, err := sql.Open("postgres", connectStr)
	if err != nil {
		return err
	}
	m.DB = db
	m.StatusConnected = true
	return nil
}

func (m *Model) ExecuteQuery(query string) ([][]string, error) {
	db := m.DB

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	cols, _ := rows.Columns()
	dest := []interface{}{ // Standard MySQL columns
	}
	for i := 0; i < len(cols); i++ {
		dest = append(dest, new(string))
	}
	result := make([][]string, 0)
	result = append(result, cols)
	for rows.Next() {
		err := rows.Scan(dest...)
		parsed := make([]string, 0)
		if err == nil {
			for i, v := range dest {
				typeOfRow, _ := rows.ColumnTypes()
				switch typeOfRow[i].ScanType().String() {
				case "time.Time":
					t, _ := time.Parse(time.RFC3339, reflect.ValueOf(v).Elem().String())
					format := time.Time.Format(t, "2006-01-02 15:04:05")
					format = strings.TrimSuffix(format, "00:00:00")
					format = strings.TrimPrefix(format, "0000-01-01")
					parsed = append(parsed, format)
				default:
					parsed = append(parsed, reflect.ValueOf(v).Elem().String())
				}
			}
		}
		result = append(result, parsed)
	}
	return result, nil
}

func (m *Model) SendResult() ([]byte, error) { return nil, nil }

func (m *Model) Create(values ...string) ([][]string, error) {
	result, err := m.ExecuteQuery("SELECT * FROM ")
	return result, err
}

func (m *Model) Read(tableName string) ([][]string, error) {
	result, err := m.ExecuteQuery("SELECT * FROM " + tableName)
	return result, err
}
