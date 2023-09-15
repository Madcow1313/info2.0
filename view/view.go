package view

import (
	"html/template"
	"info2_0/controller"
	"info2_0/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IView interface {
	Init(model.IModel, controller.Controller)
	HandleQuery(query string)
	SetData(page string, data string)
	GET(request []interface{})
	POST(request []interface{})
}
type Fields struct {
	Fields map[string]any
}

type q struct {
	Q string `query:"value"`
}

type View struct {
	Router      *gin.Engine
	Data        Fields
	model       model.IModel
	Querydata   string
	TableNames  []string
	TableFields map[string]string
	Current     string
	controller  controller.Controller
}

func (d *Fields) FillMain() {
	d.Fields["header"] = template.HTML(`<div style="height: 123px; background-attachment: fixed; background-color: #1D2633;; margin: -10px; margin-top: -18px;"><div><button id="homepage_btn" class="btn"
	style="border: none !important; position: fixed; width: 120px; height: 121px; background:transparent;"><img id="logo" src="logo.svg" width="120px" height="121px" alt="no image"></img></button></p></div>`)
	d.Fields["main_buttons"] = template.HTML(`<div class="parent"><button id="about_btn" class="buttons" href="about.html";">About me</button>
	<button id="data_btn" class="buttons" href="data.html">Data</button>
	<button id="operations_btn" class="buttons" href="operations.html">Operations</button>
	<script src="./static/js/main.js"></script></div>`)
}

func loadFiles(engine *gin.Engine) {
	engine.LoadHTMLFiles("index.html", "about.html", "operations.html", "data.html")
	engine.StaticFile("/logo.svg", "logo.svg")
	engine.StaticFile("/static/js/main.js", "./static/js/main.js")
	engine.StaticFile("style.css", "./static/css/style.css")
	engine.StaticFile("/static/js/crud.js", "./static/js/crud.js")
}

func (v *View) Init(m model.IModel, c controller.Controller) {
	var data Fields
	data.Fields = make(map[string]any)
	data.FillMain()

	router := gin.Default()
	v.Data = data
	v.Router = router
	v.model = m
	v.controller = c
	v.controller.Model = v.model
	v.controller.Data = v.Data.Fields
	loadFiles(v.Router)
	// v.extractTableNames()
	// v.extractTableFields()
	v.controller.ExtractTableNames()
	v.controller.ExtractTableFields()
	v.TableFields = v.controller.TableFields
	v.TableNames = v.controller.TableNames
	v.GET(nil)
	v.Router.Run()
}

func (v *View) SetController() {
	// if v.Current != "" {
	// 	v.controller.Current = v.Current
	// }
	v.controller.Querydata = v.Querydata
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
		v.controller.FillBaseData(ctx)
		v.SetController()
		ctx.HTML(http.StatusOK, "data.html", v.controller.Data)
	})
	v.Router.GET("/data.html/:btn", func(ctx *gin.Context) {
		str, _ := ctx.GetQuery("value")
		v.Querydata = str
		v.SetController()
		v.controller.FillBaseData(ctx)
		ctx.HTML(http.StatusOK, "data.html", v.controller.Data)
	})
	v.Router.GET("/operations.html", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "operations.html", v.Data.Fields)
	})
}

func (v *View) POST(request []interface{})       {}
func (v *View) HandleQuery(string)               {}
func (v *View) SetData(page string, data string) {}
