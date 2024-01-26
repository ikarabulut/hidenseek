package util

import (
	"crypto/md5"
	"encoding/hex"
)


func CreateMd5Hex(secret string) (secretHex string) {
	data := []byte(secret)
  md5Raw := md5.Sum(data)
	secretHex = hex.EncodeToString(md5Raw[:])
	return secretHex
}