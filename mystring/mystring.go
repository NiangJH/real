package mystring

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type HelloWorld struct {
	User string `json:"user is:"`
	Pass string `json:"pass is:"`
}

func (h HelloWorld) String() string {
	return "momoda:" + h.User + "," + h.Pass
}

func (h HelloWorld) checkWord() bool {
	return true
}

func encodingString(str string) {
	fmt.Printf("%+q\n", str)
	fmt.Printf("%q\n", 0x2318)

	fmt.Printf("Hex:")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
	fmt.Printf("\n")
}

func appendStringI(str string) {
	fmt.Printf("unicode:")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x", str[i])
	}
	fmt.Printf("\n")

	fmt.Printf("utf8:")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	fmt.Printf("\n")
}

func NewString1(strlen int) string {
	rand.Seed(time.Now().Unix())
	var word = "abcdefghjiklmnopqrstuvwxyz1234567890"
	rand.Seed(time.Now().Unix())

	var tmpStr = make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		index := rand.Intn(len(word))
		tmpStr[i] = word[index]
	}

	return string(tmpStr)
}

func NewString2(strlen int) string {
	var word = "abcdefghjiklmnopqrstuvwxyz1234567890"
	rand.Seed(time.Now().Unix())

	var tmpStr string
	for i := 0; i < strlen; i++ {
		index := rand.Intn(len(word))
		tmpStr += string(word[index])
	}

	return tmpStr
}

func NewString3(strlen int) string {
	rand.Seed(time.Now().Unix())
	var word = "abcdefghjiklmnopqrstuvwxyz1234567890"
	rand.Seed(time.Now().Unix())

	var tmpStr = bytes.Buffer{}
	for i := 0; i < strlen; i++ {
		index := rand.Intn(len(word))
		tmpStr.WriteString(string(word[index]))
	}

	return tmpStr.String()
}
