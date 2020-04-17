package tpl

import(
	"net/http"

	"smsgate/utils"
)

//Add 添加一个模板
func Add(w http.ResponseWriter, r *http.Request) {
	utils.View.Execute(w, r, "tpl/edit.html", nil)
}
