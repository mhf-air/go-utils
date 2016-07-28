package util

import (
	"encoding/json"
	"fmt"
	"strings"
)

func debug(n int, a ...interface{}) {
	lst := []string{"<Line: " + fmt.Sprintf("%d", n) + ">\n"}
	for _, s := range a {
		b, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			fmt.Println("error in debug():", err)
		}
		//ignore the case when \n is in " "
		lines := strings.Split(string(b), "\n")
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
		fmt.Printf("%s\n", str)
		lst = []string{}
	}
}
