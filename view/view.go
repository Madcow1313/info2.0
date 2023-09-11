package view

import (
	"fmt"
	"html/template"
	"info2_0/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IView interface {
	Init(model.IModel)
	HandleQuery(query string)
	SetData(page string, data string)
	GET(request []interface{})
	POST(request []interface{})
}
type Fields struct {
	Fields map[string]any
}

type q struct {
	q string `query:"value"`
}

type View struct {
	Router    *gin.Engine
	Data      Fields
	model     model.IModel
	querydata q
}

func (d *Fields) FillMain() {
	d.Fields["header"] = template.HTML(`<div style="height: 123px; background-attachment: fixed; background-color: #1D2633;; margin: -10px; margin-top: -18px;"><div><button id="homepage_btn" class="btn"
	style="border: none !important; position: fixed; width: 120px; height: 121px; background:transparent;"><img id="logo" src="logo.svg" width="120px" height="121px" alt="no image"></img></button></p></div>`)
	d.Fields["main_buttons"] = template.HTML(`<div class="parent"><button id="about_btn" class="buttons" href="about.html";">About me</button>
	<button id="data_btn" class="buttons" href="data.html">Data</button>
	<button id="operations_btn" class="buttons" href="operations.html">Operations</button>
	<script src="./static/js/main.js"></script></div>`)
}

func (v *View) prettyHTML(res [][]string) string {
	var htmlString string
	divStart := "<div class=\"parent\">"
	divEnd := "</div>"
	for _, str := range res {
		htmlString += divStart
		for _, str2 := range str {
			htmlString += `<p style="padding: 0px; text-align: center; width: 300px; border: 1px solid; border-color: #44EB99; border-radius: 8px;>`
			htmlString += `<label">`
			htmlString += str2
			htmlString += " "
			htmlString += "</p>"
		}
		htmlString += divEnd
	}
	return htmlString
}

func (v *View) FillBaseData(c *gin.Context) {
	param := c.Param("btn")
	if param == ":create_btn" {
		v.Data.Fields["data"] = template.HTML(`
			<label style="margin-top: -1px; margin-left: 0px; padding-top: 40px;">insert values</label>
			<form>
			<input id="insert_values" class="input_fields" type="text" placeholder="insert here">
			</form>
			<input type="submit" class="submit_buttons">
		`)
	} else if param == ":read_btn" {
		res, _ := v.model.Read(v.querydata.q)
		htmlString := v.prettyHTML(res)
		v.Data.Fields["data"] = template.HTML(`
				<label style="margin-top: -1px; margin-left: 0px; padding-top: 40px;"></label>
		` + htmlString)
	} else if param == ":update_btn" {
		v.Data.Fields["data"] = template.HTML(`
			<label style="margin-top: -1px; margin-left: 0px; padding-top: 40px;">update</label>
			<form>
			<input id="insert_values" class="input_fields" type="text" placeholder="id, name, e.t.c.">
			</form>
			<br>
			<label style="margin-top: -1px; margin-left: 0px; padding-top: 40px;">values</label>
			<form>
			<input id="insert_values" class="input_fields" type="text" placeholder="1, Komi, e.t.c.">
			</form>
			<br>
			<label style="margin-top: -1px; margin-left: 0px; padding-top: 40px;">where</label>
			<form>
			<input id="insert_values" class="input_fields" type="text" placeholder="condition">
			</form>
			<input type="submit" class="submit_buttons">
		`)
	} else if param == ":delete_btn" {
		v.Data.Fields["data"] = template.HTML(`
			<label style="margin-top: -1px; margin-left: 0px; padding-top: 40px;">delete</label>
			<form>
			<input id="insert_values" class="input_fields" type="text" placeholder="wildcard or empty">
			</form>
			<br>
			<label style="margin-top: -1px; margin-left: 0px; padding-top: 40px;">where</label>
			<form>
			<input id="insert_values" class="input_fields" type="text" placeholder="condition">
			</form>
			<input type="submit" class="submit_buttons">
		`)
	}
}

func loadFiles(engine *gin.Engine) {
	engine.LoadHTMLFiles("index.html", "about.html", "operations.html", "data.html")
	engine.StaticFile("/logo.svg", "logo.svg")
	engine.StaticFile("/static/js/main.js", "./static/js/main.js")
	engine.StaticFile("style.css", "./static/css/style.css")
	engine.StaticFile("/static/js/crud.js", "./static/js/crud.js")
}

func (v *View) Init(m model.IModel) {
	var data Fields
	data.Fields = make(map[string]any)
	data.FillMain()

	router := gin.Default()
	v.Data = data
	v.Router = router
	v.model = m
	loadFiles(v.Router)
	v.GET(nil)
	v.Router.Run()
}

func (v *View) GET(request []interface{}) {
	v.Router.GET("/", func(c *gin.Context) {
		v.Data.Fields["data"] = ""
		c.HTML(http.StatusOK, "index.html", v.Data.Fields)
	})
	v.Router.GET("/about.html", func(ctx *gin.Context) {
		v.Data.Fields["data"] = ""
		ctx.HTML(http.StatusOK, "about.html", v.Data.Fields)
	})
	v.Router.GET("/data.html", func(ctx *gin.Context) {
		v.FillBaseData(ctx)
		ctx.HTML(http.StatusOK, "data.html", v.Data.Fields)
	})
	v.Router.GET("/data.html/:btn", func(ctx *gin.Context) {
		var query q
		if ctx.ShouldBind(&query) == nil {
			v.querydata = query
			fmt.Println("query is **", v.querydata.q)
		} else {
			fmt.Println("none")
		}
		v.FillBaseData(ctx)
		ctx.HTML(http.StatusOK, "data.html", v.Data.Fields)
	})
	v.Router.GET("/operations.html", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "operations.html", v.Data.Fields)
	})
}

func (v *View) POST(request []interface{})       {}
func (v *View) HandleQuery(string)               {}
func (v *View) SetData(page string, data string) {}
