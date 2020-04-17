package controllers

import (
	"net/http"

	"smsgate/server/controllers/blacklist"
	"smsgate/server/controllers/queue"
	"smsgate/server/controllers/sms"
)

var Routers = map[string]http.HandlerFunc{
	// 发送短信(兼容旧版接口)
	"/send":             sms.Send,
	// 批量发送(兼容旧版接口)
	"/batch":            sms.Batch,
	// 发送短信
	"/sms/send":         sms.Send,
	//获取上行短信
	"/sms/reply":        sms.Reply,
	// 添加黑名单
	"/blacklist/add":    blacklist.Add,
	// 移除黑名单
	"/blacklist/remove": blacklist.Remove,
	// 获取队列长度
	"/queue/length":     queue.Length,
}
