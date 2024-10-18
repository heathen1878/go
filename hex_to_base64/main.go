package hextobase64

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	var hex_value string

	fmt.Printf("Enter hex value to convert to base64: ")

	fmt.Scanln(&hex_value)

	b64Str := HexToBase64(hex_value)
	fmt.Printf("Your base64 value is: %s", b64Str)
}

func HexToBase64(hex_value string) string {
	hexBytes, _ := hex.DecodeString(hex_value)
	b64Str := base64.RawURLEncoding.EncodeToString(hexBytes)
	return b64Str
}
