package main

import (
	"crypto/md5"
	"encoding/hex"
)

var map_hashs map[string]string

const NUMS = 10

var array_passwords [NUMS]string = [NUMS]string{"a", "b", "c", "d", "e",
	"F", "G", "H", "I", "J"}
var array_values [NUMS]string = [NUMS]string{"hello world", "first app", "imac", "library", "golang", "python3.8",
	"goodbye", "unix", "linux", "mac os"}

func GetMd5(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func append_map() {
	for i := 0; i < NUMS; i++ {
		map_hashs[GetMd5(array_passwords[i])] = array_values[i]
	}
}
