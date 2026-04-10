package foundation

import (
	"encoding/base64"
	"log"
)

func SendMockEmail(to, subject, body string) {
	log.Printf("Sending mock email to %s:\n subject=%q\n body:\n%q", to, subject, body)
}
func ToBase64(s string) string {
	enconded := base64.StdEncoding.EncodeToString([]byte(s))
	return enconded
}
func FromBase64(s string) string {
	decodedBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	decoded := string(decodedBytes)
	return decoded
}
