package utils

import (
	"crypto/md5"
	"encoding/hex"
)

var MapHash = make(map[string]string)

const NUMS = 10

var arrayPasswords [NUMS]string = [NUMS]string{"a", "b", "c", "d", "e",
	"F", "G", "H", "I", "J"}
var arrayValues [NUMS]string = [NUMS]string{"hello world", "first app", "imac", "library", "golang", "python3.8",
	"goodbye", "unix", "linux", "mac os"}

func AppendMap() {
	for i := 0; i < NUMS; i++ {
		MapHash[GetMd5(arrayPasswords[i])] = arrayValues[i]
	}
}

func GetMd5(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
