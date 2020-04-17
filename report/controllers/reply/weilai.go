package reply

import (
	"log"
	"net/http"

	"smsgate/report/modules/reply"
	"smsgate/utils"
)

func Weilai(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	log.Printf("weilai reply. reply data:%v", r.PostForm)
	msgid := r.PostFormValue("msg_id")

	spCode := r.PostFormValue("sp_code")

	mobile := r.PostFormValue("src_mobile")
	if mobile == "" {
		log.Printf("weilai reply. mobile is invaild. data:%v", r.PostForm)
		utils.ResponseJson(w, -1, "mobile is invaild", nil)
		return
	}

	content := r.PostFormValue("msg_content")

	recvTimeStr := r.PostFormValue("recv_time")
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
		Account:  "weilai",
	}

	err := smsReply.Insert()
	if err != nil {
		log.Printf("weilai reply. insert into mysql failed. err:%v, smsreply:%v", err, smsReply)
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	utils.ResponseJson(w, 0, "success", nil)
}
