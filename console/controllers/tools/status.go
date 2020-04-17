package tools

import(
	"net/http"
	"strings"

	"smsgate/console/modules/status"
	"smsgate/utils"
)

func Status(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	statusCode := strings.TrimSpace(vars.Get("status"))

	data := make(map[string]status.ReportStatus)
	if statusCode == "" {
		data = status.Status
	} else {
		data[statusCode] = status.Status[statusCode]
	}

	info := make(map[string]interface{})

	info["data"] = data
	info["status"] = statusCode

	utils.View.Execute(w, r, "tools/status.html", info)
}
