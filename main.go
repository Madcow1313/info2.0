package main

import (
	"fmt"
	"info2_0/model"
	"info2_0/view"
	"os"
)

func initView(iv view.IView, m model.IModel) {
	iv.Init(m)
}

func initModel(im model.IModel, pathToConfig string) error {
	b, err := os.ReadFile(pathToConfig)
	if err == nil {
		im.ConnectToDB(b)
	} else {
		fmt.Printf("Couldn't open config.json: %v", err)
		return err
	}
	return nil
}

func main() {

	m := new(model.Model)

	if err := initModel(m, "config.json"); err == nil {
		m.StatusConnected = true
		fmt.Println("DB connected")
		defer m.DB.Close()
	}
	v := new(view.View)
	initView(v, m)
}
