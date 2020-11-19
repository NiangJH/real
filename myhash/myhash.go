package myhash

import (
	"crypto/md5"
	"fmt"
)

func myMd5() {
	md5er := md5.New()
	json := md5er.Sum([]byte{41, 23, 11})
	fmt.Println(json)
}