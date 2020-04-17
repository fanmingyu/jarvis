package blacklist

import (
	"net/http"
	"smsgate/server/storage"
	"smsgate/utils"
)

func Add(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	mobile := r.PostFormValue("mobile")
	if len(mobile) < 10 {
		utils.ResponseJson(w, -1, "mobile is invaild.", nil)
		return
	}

	err := storage.AddBlacklist(mobile)
	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	utils.ResponseJson(w, 0, "add blacklist success.", nil)
}
