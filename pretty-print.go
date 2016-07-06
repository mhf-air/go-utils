package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {

}

//====================================================================================================
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ln(a ...interface{}) {
	for _, v := range a {
		b, err := json.MarshalIndent(v, "", "  ")
		check(err)

		//ignore the case when \n is in " "
		lines := strings.Split(string(b), "\n")
		lst := []string{}
		for _, l := range lines {
			pureline := strings.TrimSpace(l)
			frontBlank := strings.Repeat(" ", len(l)-len(pureline))
			if strings.HasPrefix(pureline, "]") || strings.HasPrefix(pureline, "}") {
				continue
			}
			if strings.HasSuffix(pureline, ",") {
				pureline = pureline[:len(pureline)-1]
			}
			lst = append(lst, frontBlank+pureline+"\n")
		}
		str := strings.Join(lst, "")
		fmt.Printf("%s", str)
	}
}
