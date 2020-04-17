package tools

import(
	"net/http"

	"smsgate/console/modules/queue"
	"smsgate/utils"
)

func Queue(w http.ResponseWriter, r *http.Request) {
	queueData := queue.GetQueueLen()

	info := map[string]interface{}{}
	info["queue"] = queueData

	utils.View.Execute(w, r, "tools/queue.html", info)
}
