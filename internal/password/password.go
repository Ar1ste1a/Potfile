package password

import (
	"encoding/hex"
	"regexp"
	"strings"
)

func IsHexPassword(hexString string) bool {
	return strings.Contains(hexString, "$HEX[")
}

func HexToPassword(hexString string) string {
	hexList := separateHex(hexString)
	if len(hexList) == 2 {
		decodedValue := decodeHex(hexList[0])
		decodedValue += decodeHex(hexList[1])
		return decodedValue
	} else {
		decodedValue := decodeHex(hexList[0])
		return decodedValue
	}
}

func separateHex(hexString string) []string {
	var out []string

	if wholeStringHex(hexString) {
		return []string{hexString}
	}

	out = make([]string, 2)
	hexMatch := `\$HEX\[[0-9a-fA-F]+\]`
	re := regexp.MustCompile(hexMatch)
	match := re.FindAllString(hexString, -1)
	if len(match) == 2 {
		out[0] = match[0]
		out[1] = match[1]
		return out
	} else {
		if strings.HasPrefix(hexString, "$HEX[") {
			out[0] = match[0]
			startIndex := strings.Index(hexString, match[0]) + len(match[0])
			out[1] = hexString[startIndex:]
		} else if strings.HasSuffix(hexString, "]") {
			out[1] = match[0]
			endIndex := strings.Index(hexString, match[0])
			out[0] = hexString[:endIndex]
		}
	}
	return out
}

func wholeStringHex(hexString string) bool {
	// Example of a whole string hex: $HEX[4f63746f626572313938393a29]
	// Example of a partial string hex: |^{ZEM+$HEX[333e43415c3a4b] or $HEX[3a32433a514853]!,<M^O9

	if strings.HasPrefix(hexString, "$HEX[") && strings.HasSuffix(hexString, "]") {
		return true
	}

	return false
}

func decodeHex(hexValue string) string {
	if strings.Contains(hexValue, "$HEX[") {
		hexValue = strings.TrimSuffix(strings.TrimPrefix(hexValue, "$HEX["), "]")
		decodedHex, err := hex.DecodeString(hexValue)
		if err != nil {
			return ""
		}

		return string(decodedHex)
	}

	return hexValue
}

func SplitPassword(line string) (string, string) {
	var (
		hash string
		pw   string
	)

	startPassword := strings.LastIndex(line, ":")
	if startPassword == -1 {
		return line, ""
	}

	hash = line[:startPassword]
	pw = line[startPassword+1:]

	return hash, pw
}
