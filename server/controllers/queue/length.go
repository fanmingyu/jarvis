package queue

import(
	"net/http"

	"smsgate/server/workers"
	"smsgate/utils"
)

func Length(w http.ResponseWriter, r *http.Request) {
	result := make(map[string]map[string]int)
	for k, v := range workers.Workers {
		result[k] = make(map[string]int)
		result[k]["len"] = v.Queue.Len()
		result[k]["buffer"] = v.Queue.BufferLen()
	}

	utils.ResponseJson(w, 0, "success", result)
}
