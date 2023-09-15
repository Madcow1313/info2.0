package controller

import (
	"fmt"
	"html/template"
	"info2_0/model"
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Model       model.IModel
	Current     string
	Querydata   string
	TableNames  []string
	TableFields map[string]string
	Data        map[string]any
}

func (c *Controller) ExtractTableFields() {
	var fields string

	c.TableFields = make(map[string]string)
	for _, table := range c.TableNames {
		rows, err := c.Model.ExecuteQuery("SELECT * FROM " + table)
		if err == nil {
			for _, r := range rows[0] {
				fields += r + ","
			}
		}
		c.TableFields[table] = strings.TrimSuffix(fields, ",")
		fields = ""
	}
}

func (c *Controller) ExtractTableNames() {
	res, err := c.Model.ExecuteQuery(`select table_name from information_schema.tables where table_schema='public'`)
	if err == nil {
		tableNames := res[1:]
		var options string
		for _, table := range tableNames {
			options += `<option value="` + table[0] + `" label="` + table[0] + `" class="dropdown-content"></option>`
			c.TableNames = append(c.TableNames, table[0])
		}
		c.Data["dropdown"] = template.HTML(options)
	} else {
		c.Data["dropdown"] = template.HTML(`no tables in bd`)
	}
}

func (c *Controller) PrettyHTML(raw [][]string) string {
	var htmlString string
	divStart := "<div class=\"parent\">"
	divEnd := "</div>"
	for _, str := range raw {
		htmlString += divStart
		for _, str2 := range str {
			htmlString += `<p style="padding: 0px; text-align: center; width: 300px; border: 1px solid;
			border-color: #44EB99; border-radius: 8px;>`
			htmlString += `<label">`
			htmlString += str2
			htmlString += " "
			htmlString += "</p>"
		}
		htmlString += divEnd
	}
	return htmlString
}

func (c *Controller) FillBaseData(ctx *gin.Context) {
	param := ctx.Param("btn")
	if param == "" {
		return
	}
	if param != ":changeDrop" {
		c.Current = param
		c.Querydata = ctx.Query("value")
	}
	fmt.Println(c.Current, c.Querydata)
	if param == ":create_btn" || (param == ":changeDrop" && c.Current == ":create_btn") {
		c.Data["data"] = template.HTML(`
			<div class="parent">
			<label style="margin-top: -1px; margin-left: 0px; padding-top: 40px;">insert values</label>
			<form>
			<input id="insert_values" class="input_fields" type="text" placeholder="` + c.TableFields[c.Querydata] + `"
			title="` + c.TableFields[c.Querydata] + `" >
			</form>
			<input type="submit" class="submit_buttons">
			</div>
		`)
	} else if param == ":read_btn" {
		res, _ := c.Model.Read(c.Querydata)
		htmlString := c.PrettyHTML(res)
		c.Data["data"] = template.HTML(`
				<div class="parent">
				<label style="margin-top: -1px; margin-left: 0px; font-size: 30px;">` + c.Querydata + ` table data</label></div>` + htmlString)
	} else if param == ":update_btn" || (param == ":changeDrop" && c.Current == ":update_btn") {
		c.Data["data"] = template.HTML(`
			<div class="parent">
			<label style="margin-top: -1px; margin-left: 0px; padding-top: 40px;">update</label>
			<form>
			<input id="insert_values" class="input_fields" type="text" placeholder="` + c.TableFields[c.Querydata] + `"
			title="` + c.TableFields[c.Querydata] + `">
			</form>
			<br>
			<label style="margin-top: -1px; margin-left: 0px; padding-top: 40px;">values</label>
			<form>
			<input id="insert_values" class="input_fields" type="text" placeholder="` + c.TableFields[c.Querydata] + `"
			title="` + c.TableFields[c.Querydata] + `">
			</form>
			<br>
			<label style="margin-top: -1px; margin-left: 0px; padding-top: 40px;">where</label>
			<form>
			<input id="insert_values" class="input_fields" type="text" placeholder="condition">
			</form>
			<input type="submit" class="submit_buttons">
			</div>
		`)
	} else if param == ":delete_btn" || (param == ":changeDrop" && c.Current == ":delete_btn") {
		c.Data["data"] = template.HTML(`
			<div class="parent">
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
			</div>
		`)
	}
}
