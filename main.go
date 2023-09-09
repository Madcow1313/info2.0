package main

import (
	"fmt"
	"info2_0/model"
	"info2_0/view"
	"os"
)

func initView(iv view.IView) {
	iv.Init()
}

func main() {

	var m model.Model

	b, err := os.ReadFile("config.json")
	if err == nil {
		fmt.Printf("Couldn't open config.json: %v", err)
		m.StatusConnected = false
		m.ConnectToDB(b)
	}
	defer m.DB.Close()
	v := new(view.View)
	initView(v)
}
