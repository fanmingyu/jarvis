package utils

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

type view struct {
	funcs      template.FuncMap
	viewPath   string
	headerPath string
	footerPath string
}

type Data struct {
	Info     interface{}
	Username string
	Cost     time.Duration
	Now      string
}

var View = &view{}

func (v *view) Execute(w http.ResponseWriter, r *http.Request, html string, info interface{}) {
	file, err := template.New(GetFileName(html)).Funcs(v.funcs).ParseFiles(v.viewPath+html, v.viewPath+v.headerPath, v.viewPath+v.footerPath)

	if err != nil {
		log.Printf("parse html failed. err:%v", err)
		return
	}

	data := processData(r, info)
	file.Execute(w, data)
}

func (v *view) Init() {
	v.viewPath = "../views/"
	v.headerPath = "common/header.html"
	v.footerPath = "common/footer.html"
}

func (v *view) SetFuncs(funcs template.FuncMap) {
	v.funcs = funcs
}

func processData(r *http.Request, info interface{}) Data {
	ctx := r.Context()
	username := ctx.Value("username").(string)

	start := ctx.Value("start").(time.Time)
	cost := time.Since(start).Truncate(time.Millisecond)
	now := time.Now().Format("2006-01-02 15:04:05")

	return Data{
		Info:     info,
		Username: username,
		Cost:     cost,
		Now:      now,
	}
}
