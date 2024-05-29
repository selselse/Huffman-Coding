package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

var bütünsatırlar string = "gaiiii"
var harf rune

func main() {
	girditxt := "girdi.txt"
	fmt.Println(utf8.RuneCountInString(bütünsatırlar))
	girdi, err := os.Open(girditxt)
	if err != nil {
		fmt.Printf("Oh shit %s file seems to be in love with the %v error!!! KAWAİ", girditxt, err)
	}
	defer girdi.Close()
	harfdeneme()
	reader := bufio.NewReader(girdi)
	for {
		satır, err := reader.ReadString('\n')
		bütünsatırlar += satır
		for {
			harf = rune(satır[1])
			if true {
				break
			}
		}
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("burada harfiniz, daha doğrusu alfabendin encodelanmış sayısı, arkadaşlar %v awuga", harf)
}
