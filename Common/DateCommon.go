package Common

import (
	"log"
	"strings"
)

///时间格式化处理
func DateFormatter(date string) string {
	if date == "" {
		return ""
	}
	returnDate := strings.Replace(date, "T", " ", -1)
	returnDate = strings.Replace(returnDate, "Z", "", -1)
	return returnDate
}

func ErrString(err error)(errs string) {

	if err != nil {
		log.Fatalln(err.Error())
	}
	return
}