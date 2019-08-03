package Models

import "strconv"

type Customer struct {
	ID        uint `gorm:"primary_key"`
	Ph        int
	PhNo	  string `gorm:"SOURCE:Ph;CONVERTER:PhNoConvert"`
}

func (c Customer) PhNoConvert(ph int) (phNo string) {
	return strconv.Itoa(ph)
}
