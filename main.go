package main

import (
	"fmt"
	"regexp"
)

func main() {

	r, _ := regexp.Compile(`[^'",!().\d]`)
	fmt.Println(r.FindAllString(`User can d'água enter text (or cut and paste) into the input box. This input	box must allow the entry of large blocks of text (maximum of 2048 characters).`, -1))

}
