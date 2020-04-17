package controllers

import (
	"net/http"

	"smsgate/report/controllers/reply"
	"smsgate/report/controllers/report"
	"smsgate/report/controllers/index"
)

//Controllers 路由配置
var Controllers = map[string]http.HandlerFunc{
	//默认页
	"/":             index.Index,
	//上行
	"/reply/weilai": reply.Weilai,
	"/reply/welink": reply.Welink,

	//回执
	"/report/weilai": report.Weilai,
	"/report/welink": report.Welink,
}
