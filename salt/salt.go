package gosecurity

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func AddSalt(salt, obj string) string {
	r := md5.Sum([]byte(salt + obj))
	return hex.EncodeToString(r[:])
}

func VerifySalt(salt, obj, dish string) bool {
	if len(salt) == 0 || len(obj) == 0 || len(dish) == 0 {
		return false
	}
	return strings.EqualFold(strings.ToLower(dish), strings.ToLower(AddSalt(salt, obj)))
}
