package utils

import(
	"testing"
	"fmt"
	"crypto/md5"
)

//TestMobileString单元测试
func TestMobileString(t *testing.T) {
	mobile := []MobileString{"15900000633", "123", "12376565763275", "", "15000006338,12000006440"}
	s := fmt.Sprintf("%v", mobile[0])
	mdS := md5.Sum([]byte(mobile[0]))
	expect := fmt.Sprintf("1******0633(%x)", mdS)
	if s != expect {
		t.Errorf("mobile[0] result is:%v", s)
	}

	s = fmt.Sprintf("%v", mobile[1])
	mdS = md5.Sum([]byte(mobile[1]))
	expect = fmt.Sprintf("1**(%x)", mdS)
	if s != expect {
		t.Errorf("mobile[1] result is:%v", s)
	}

	s = fmt.Sprintf("%v", mobile[2])
	mdS = md5.Sum([]byte(mobile[2]))
	expect = fmt.Sprintf("1******5763275(%x)", mdS)
	if s != expect {
		t.Errorf("mobile[2] result is:%v", s)
	}

	s = fmt.Sprintf("%v", mobile[3])
	mdS = md5.Sum([]byte(mobile[3]))
	expect = fmt.Sprintf("(%x)", mdS)
	if s != expect {
		t.Errorf("mobile[3] result is:%v", s)
	}

	s = fmt.Sprintf("%v", mobile[4])
	mdS1 := md5.Sum([]byte("15000006338"))
	mdS2 := md5.Sum([]byte("12000006440"))
	expect = fmt.Sprintf("1******6338(%x),1******6440(%x)", mdS1, mdS2)
	if s != expect {
		t.Errorf("mobile[4] result is:%v", s)
	}

	t.Logf("test MobileString success")
}

//TestMobileFormat 手机号脱敏单元测试
func TestMobileFormat(t *testing.T) {
	mobile := "15900000000"
	result := MobileFormat(mobile)
	if result != "159****0000" {
		t.Errorf("mobile format error")
	}

	mobile = "15900001234"
	result = MobileFormat(mobile)
	if result != "159****1234" {
		t.Errorf("mobile format error")
	}
}

//BenchmarkMobileString 手机号自动脱敏的benchmark测试
func BenchmarkMobileString(b *testing.B) {
	mobile := MobileString("15900007889,12600009889")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%v\n", mobile)
	}
}
