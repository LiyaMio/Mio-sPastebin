package Encryption

import (
	"crypto/md5"
	"encoding/hex"
)

func AesKey(v string) []byte{
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return []byte(hex.EncodeToString(m.Sum(nil)))
}