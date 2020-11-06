package mystring

import (
	"testing"
)

func TestHelloWorld_String(t *testing.T) {
	//h := HelloWorld{
	//	User: "jiahao",
	//	Pass: "haoge",
	//}
	//fmt.Println(h)
}

func TestEncodingString_String(t *testing.T) {
	//encodingString("⌘")
}

func TestAppendStringI_String(t *testing.T) {
	appendStringI("mmmm")
}

func TestNewString(t *testing.T) {
	str := NewString2(8)
	t.Log(str)
}

func BenchmarkNewString1(b *testing.B) {
	for i:=0;i<b.N;i++{
		NewString1(i)
	}
}

func BenchmarkNewString2(b *testing.B) {
	for i:=0;i<b.N;i++{
		NewString2(i)
	}
}

func BenchmarkNewString3(b *testing.B) {
	for i:=0;i<b.N;i++{
		NewString3(i)
	}
}