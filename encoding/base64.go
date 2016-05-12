package gosecurity

import "encoding/base64"

func EncodeBase64String(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func DecodeBase64String(data string) string {
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return ""
	}
	return string(b)
}
