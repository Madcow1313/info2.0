package model

import (
	"database/sql"
	"encoding/json"

	_ "github.com/lib/pq"
)

type IModel interface {
	ConnectToDB(configFile []byte) error
	SendResult() ([]byte, error)
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
	connectStr := "user=" + cf.Username + " dbname=postgresql password=" + cf.Password
	db, err := sql.Open("postgres", connectStr)
	if err != nil {
		return err
	}
	m.DB = db
	m.StatusConnected = true
	return nil
}

func (m *Model) Create(values...string) {

}
