// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package utils

// ToLower is the equivalent of strings.ToLower
func ToLower(b string) string {
	var res = make([]byte, len(b))
	copy(res, b)
	for i := 0; i < len(res); i++ {
		res[i] = toLowerTable[res[i]]
	}

	return GetString(res)
}

// ToUpper is the equivalent of strings.ToUpper
func ToUpper(b string) string {
	var res = make([]byte, len(b))
	copy(res, b)
	for i := 0; i < len(res); i++ {
		res[i] = toUpperTable[res[i]]
	}

	return GetString(res)
}

// TrimLeft is the equivalent of strings.TrimLeft
func TrimLeft(s string, cutset byte) string {
	lenStr, start := len(s), 0
	for start < lenStr && s[start] == cutset {
		start++
	}
	return s[start:]
}

// Trim is the equivalent of strings.Trim
func Trim(s string, cutset byte) string {
	i, j := 0, len(s)-1
	for ; i < j; i++ {
		if s[i] != cutset {
			break
		}
	}
	for ; i < j; j-- {
		if s[j] != cutset {
			break
		}
	}

	return s[i : j+1]
}

// TrimRight is the equivalent of strings.TrimRight
func TrimRight(s string, cutset byte) string {
	lenStr := len(s)
	for lenStr > 0 && s[lenStr-1] == cutset {
		lenStr--
	}
	return s[:lenStr]
}

// Default returns the provided fallback value if string is empty
func Default(value string, defaultValue string) string {
	if len(value) <= 0 {
		return defaultValue
	}
	return value
}
