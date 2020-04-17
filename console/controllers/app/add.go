package app

import (
	"net/http"

	"smsgate/server/channels"
	"smsgate/server/workers"
	"smsgate/utils"
)

//Add 渲染添加app页面的controller
func Add(w http.ResponseWriter, r *http.Request) {
	chans := getChannels()
	worker := getWorkers()
	info := make(map[string]interface{})

	info["channels"] = chans
	info["workers"] = worker

	utils.View.Execute(w, r, "app/edit.html", info)
}

func getChannels() []string {
	chans := make([]string, 0)
	for key := range channels.Channels {
		chans = append(chans, key)
	}

	return chans
}

func getWorkers() []string {
	worker := make([]string, 0)
	for key := range workers.Workers {
		worker = append(worker, key)
	}

	return worker
}
