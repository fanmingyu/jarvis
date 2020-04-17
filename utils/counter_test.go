package utils

import (
	"testing"
)

func TestSmsCounter(t *testing.T) {
	content := "1234567890123456789012345678901234567890123456789012345678901234567890"
	n := SmsCounter(content)
	t.Logf("n = %v", n)
	if n != 1 {
		t.Errorf("sms count error")
	}

	content = "一1234567890123456789012345678901234567890123456789012345678901234567890"
	n = SmsCounter(content)
	t.Logf("n = %v", n)
	if n != 2 {
		t.Errorf("sms count error")
	}

	content = "123是67890123456789012345678901234567890123456789012345678901234567890"
	n = SmsCounter(content)
	t.Logf("n = %v", n)
	if n != 1 {
		t.Errorf("sms count error")
	}

	content = "123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012340"
	n = SmsCounter(content)
	t.Logf("n = %v", n)
	if n != 3 {
		t.Errorf("sms count error")
	}

	content = "23456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012340"
	n = SmsCounter(content)
	t.Logf("n = %v", n)
	if n != 2 {
		t.Errorf("sms count error")
	}
}
