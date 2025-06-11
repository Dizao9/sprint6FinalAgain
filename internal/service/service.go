package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func MorzeDetect(text string) string {
	if len(text) == 0 {
		return " "
	}
	flag := false
	for _, v := range text {
		if v != '-' && v != ' ' && v != '.' {
			flag = true
		}
	}
	if flag == true {
		return morse.ToMorse(text)
	}
	return morse.ToText(text)
}
