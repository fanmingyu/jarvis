package test

import(
	"testing"
)

func TestSend(t *testing.T) {
	msgid, err := Send("15567289292", "testfail")
	if msgid != "" || err == nil{
		t.Errorf("test wrong send failed.")
	}

	msgid, err = Send("155234566", "尊敬的xxx你的xxx已过期")
	if msgid == "" || err != nil {
		t.Errorf("test right send failed")
	}
	t.Logf("%s", msgid)
}
