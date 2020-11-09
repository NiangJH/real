package array

import (
	"bytes"
	"fmt"
)

func initArray() {
	//
}

func readBuffer() {
	myBuffer := bytes.Buffer{}
	myBuffer.WriteString("h")
	myBuffer.WriteString("e")
	myBuffer.WriteString("l")
	myBuffer.WriteString("l")
	myBuffer.WriteString("o")

	fmt.Printf("myBuffer %s \n", myBuffer.String())
}
