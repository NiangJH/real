package myStruct

import "testing"

func TestRectangle_Area(t *testing.T) {
	myRec := new(Rectangle)
	myRec.Length = 5
	myRec.Width = 2

	area := myRec.Area()
	perimeter := myRec.Perimeter()

	t.Logf("area is %d, perimeter is:%d \n", area, perimeter)
	t.Logf("width is %d, length is:%d \n", myRec.Width, myRec.Length)
}

func TestRectangle_Perimeter(t *testing.T) {
	myRec := &Rectangle{Length: 5, Width: 2}

	area := myRec.Area()
	perimeter := myRec.Perimeter()

	t.Logf("area is %d, perimeter is:%d \n", area, perimeter)
	t.Logf("width is %d, length is:%d \n", myRec.Width, myRec.Length)
}
