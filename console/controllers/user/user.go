package user

import(
	"net/http"

	"smsgate/utils"
)

//Username 获取http header中的Authgate-Username
func Username(w http.ResponseWriter, r *http.Request) {
	authgateUsername := r.Header["Authgate-Username"]
	username := ""
	if len(authgateUsername) > 0 {
		username = authgateUsername[0]
	}

	utils.ResponseJson(w, 0, "", username)
}
