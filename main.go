package main

import (
	"fmt"
	"os"
)

func main() {
	girditxt := "girdi.txt"
	girdi, err := os.Open(girditxt)
	if err != nil {
		fmt.Printf("Oh shit")
		return
	}
	defer girdi.Close()
	/*aaa*/
}
