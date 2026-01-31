package main

import (
	"mysqldb/crud/data/data"
	"mysqldb/crud/data/mysqldb/crud"
	"net/http"
	"strconv"
	"text/template"
)

const baseUrlTemplateFolder = "templates/"

func index(w http.ResponseWriter, req *http.Request) {
	var ViewAll = data.AllColuns{
		ViewAll: crud.SelectAll(),
	}
	tmpl := template.Must(template.ParseFiles(baseUrlTemplateFolder + "index.html"))
	tmpl.Execute(w, ViewAll)
}

func addProd(w http.ResponseWriter, req *http.Request) {
	var data data.Coluns
	if req.Method == "POST" {

		req.ParseForm()
		data.NomeProd = req.Form.Get("nome")
		preco, _ := strconv.Atoi(req.Form.Get("preco"))
		val, _ := strconv.Atoi(req.Form.Get("val"))
		code, _ := strconv.Atoi(req.Form.Get("code"))
		data.Preco = preco
		data.ValAno = val
		data.CodeBar = code
	}
	crud.Insert(data)
	tmpl := template.Must(template.ParseFiles(baseUrlTemplateFolder + "addProd.html"))
	tmpl.Execute(w, data)
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/add", addProd)
	http.ListenAndServe(":8090", mux)
}
