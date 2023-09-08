package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

type Data struct {
	fields map[string]any
}

/*TODO: change from hard-coding to parsing from file*/
func (d *Data) fillMain() {
	d.fields["header"] = template.HTML(`<div style="height: 123px; background-attachment: fixed; background-color: #1D2633;; margin: -10px; margin-top: -18px;"><div><button id="homepage_btn" class="btn"
	style="border: none !important; position: fixed; width: 120px; height: 121px; background:transparent;"><img id="logo" src="logo.svg" width="120px" height="121px" alt="no image"></img></button></p></div>`)
	d.fields["main_buttons"] = template.HTML(`<div class="parent"><button id="about_btn" class="buttons" href="about.html";">About me</button>
	<button id="data_btn" class="buttons" href="data.html">Data</button>
	<button id="operations_btn" class="buttons" href="operations.html">Operations</button>
	<script src="./static/js/main.js"></script></div>`)
}


func (d *Data) fillCRUD() {
	
} 

func loadFiles(engine *gin.Engine) {
	engine.LoadHTMLFiles("index.html", "about.html", "operations.html", "data.html")
	engine.StaticFile("/logo.svg", "logo.svg")
	engine.StaticFile("/static/js/main.js", "./static/js/main.js")
	engine.StaticFile("style.css", "./static/css/style.css")
	engine.StaticFile("/static/js/crud.js", "./static/js/crud.js")
}
