package reply

import(
	"net/http"
	"log"

	"smsgate/report/modules/reply"
	"smsgate/utils"
)

func Welink(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	log.Printf("welink reply. reply data:%v", vars)
	user := vars.Get("AccountId")

	msgid := vars.Get("MsgId")

	spCode := vars.Get("Up_YourNum")

	mobile := vars.Get("Up_UserTel")
	if mobile == "" {
		log.Printf("welink reply. mobile is invaild. data:%v", vars)
		utils.ResponseJson(w, -1, "mobile is invaild", nil)
		return
	}

	content := vars.Get("Up_UserMsg")

	recvTimeStr := vars.Get("MoTime")
	recvTimeUnix := int(utils.GetUnixTime(recvTimeStr, "2006-01-02 15:04:05"))
	if recvTimeUnix < 0 {
		recvTimeUnix = 0
	}

	smsReply := &reply.SmsReply{
		Msgid:    msgid,
		SpCode:   spCode,
		Mobile:   mobile,
		Content:  content,
		RecvTime: recvTimeUnix,
		Account:  user,
	}

	err := smsReply.Insert()
	if err != nil {
		log.Printf("welink reply. insert into mysql failed. err:%v, smsreply:%v", err, smsReply)
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	utils.ResponseJson(w, 0, "success", nil)
}
