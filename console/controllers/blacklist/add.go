package blacklist

import(
	"net/http"

	"smsgate/utils"
)

//Add 添加一条黑名单
func Add(w http.ResponseWriter, r *http.Request) {
	utils.View.Execute(w, r, "blacklist/edit.html", nil)
}
