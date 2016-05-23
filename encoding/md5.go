package gosecurity

import (
	"crypto/md5"
	"encoding/hex"
)

const EmptyMD5Hash string = "00000000000000000000000000000000"

func GetMD5Hash(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func ValidMD5Hash(s string) bool {
	return len(s) == 32
}
