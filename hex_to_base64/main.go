package hextobase64

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(hex_value string) string {
	hexBytes, _ := hex.DecodeString(hex_value)
	b64Str := base64.RawURLEncoding.EncodeToString(hexBytes)
	return b64Str
}
