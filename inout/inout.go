package inout

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func InputTerminal() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("input something...")
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("input string: %s \n", input)
	}
}

func inputFileByLine() {
	inputFile, error := os.Open("text.txt")
	if error != nil {
		return
	}

	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readError := inputReader.ReadString('\n')
		fmt.Printf("input is : %s", inputString)
		if readError == io.EOF {
			return
		}
	}
}

func inputFileAll() {
	buf, err := ioutil.ReadFile("text.txt")
	if err != nil {
		return
	}

	fmt.Printf("%s\n", string(buf))
}

func myFrintf() {

}