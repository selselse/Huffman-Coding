package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	girditxt := "girdi.txt"
	girdi, err := os.Open(girditxt)
	if err != nil {
		fmt.Printf("Oh shit %s file seems to be in love with the %v error!!! KAWAİ", girditxt, err)
		return
	}
	defer girdi.Close()
	/*aaa*/
	reader := bufio.NewReader(girdi)
	for {
		line, err := reader.ReadString('\n')
		fmt.Print(line)
		if err == io.EOF {
			break
		}
	}
}
