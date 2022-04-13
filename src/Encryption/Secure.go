package Encryption

import (
	"log"
)

func DataEncrypt(content []byte)([]byte,[]byte){
	priK,pubK := GenRsaKey()
	log.Println("your prikey",string(priK))
	log.Println("your pubkey",string(pubK))
	ciphertext := RsaEncrypt(content, pubK)
	return ciphertext,priK
}
//func DataDncrypt(priK string,ciphertext string)string{
//
//}
