package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {

	splitDomain := strings.Split(email, "@")
	splitTLD := strings.Split(splitDomain[1], ".")

	var TLD string
	if len(splitTLD) > 2 {
		TLD = strings.Join(splitTLD[1:], ".")
	} else {
		TLD = splitTLD[1]
	}

	return "Domain: " + splitTLD[0] + " dan TLD: " + TLD // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
