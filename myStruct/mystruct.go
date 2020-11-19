package myStruct

type Rectangle struct {
	Length int `json:"length is:"`
	Width int `json:"width is:"`
}

func (r Rectangle)Area() int{
	r.Length = 1
	return r.Length * r.Width
}

func (r Rectangle)Perimeter() int {
	r.Length = 1
	return (r.Length+r.Width) * 2
}