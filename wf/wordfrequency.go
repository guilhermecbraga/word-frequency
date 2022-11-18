package wf

import (
	"log"
	"regexp"
	"strings"
)

type WordFrequencer map[string]int

func CleanText(s string) string {
	r, err := regexp.Compile(`[A-Za-záàâãéèêíóôõúçÁÀÂÃÉÈÍÓÔÕÚÇ ]`)
	if err != nil {
		log.Fatal(err.Error())
	}
	return strings.Join(r.FindAllString(strings.ToLower(s), -1), "")
}
