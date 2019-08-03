package Models

import "strconv"

type Customer struct {
	ID        uint `gorm:"primary_key"`
	Ph        int
	PhNo	  string `gorm:"SOURCE:Ph;CONVERTER:PhNoConvert"`
	NewPhNo   int `gorm:"SOURCE:PhNo;CONVERTER:PhNoConvertToInt"`
}

func (c Customer) PhNoConvert(ph int) (phNo string) {
	return strconv.Itoa(ph)
}

func (c Customer) PhNoConvertToInt(ph string) (phNo int) {
	out, _ := strconv.Atoi(ph)
	return out / 10
}
