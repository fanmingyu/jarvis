package controllers

import (
	"fmt"
	"time"

	"smsgate/server/channels"
	"smsgate/server/workers"
)

//ViewFuncs 将模板要使用的funcs加载到view中
var ViewFuncs = map[string]interface{}{
	"getChannelName": getChannelName,
	"getWorkerName":  getWorkerName,
	"calcPercent":    calcPercent,
	"default":    defaultValue,
	"formatTime":    formatTime,
}

func getChannelName(channel string) string {
	return channels.Channels[channel].Name
}

func getWorkerName(worker string) string {
	return workers.Workers[worker].Name
}

func calcPercent(num int, sum int) string {
	if sum == 0 {
		return "0.00"
	}
	rate := float64(num)/float64(sum) * 100.00
	result := fmt.Sprintf("%.2f", rate)

	return result
}

func defaultValue(defaultValue interface{}, value interface{}) interface{} {
	if v, ok := value.(int); ok && v == 0 {
		return defaultValue
	}

	if v, ok := value.(string); ok && v == "" {
		return defaultValue
	}

	return value
}

func formatTime(timestamp int) string {
	if timestamp == 0 {
		return ""
	}

	return time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
}
