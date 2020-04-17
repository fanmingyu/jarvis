package storage

import(
	"strings"
)

func AddBlacklist(mobile string) error {
	data := SmsBlacklist {
		Mobile: mobile,
	}

	_, err := Mysql.Engine.Insert(&data)
	return err
}

func RemoveBlacklist(mobile string) error {
	var data SmsBlacklist

	_, err := Mysql.Engine.Where("(mobile) = ?", mobile).Delete(&data)
	return err
}

func InBlacklist(mobile string) bool {
	if CacheData.GetBlacklist(mobile) == 1 {
		return true
	}
	return false
}

func DivideBlacklist(mobile string) (white []string, black []string) {
	mobiles := strings.Split(mobile, ",")
	white = mobiles[:0]

	for _, value := range mobiles {
		if InBlacklist(value){
			black = append(black, value)
		} else {
			white = append(white, value)
		}
	}
	return
}
