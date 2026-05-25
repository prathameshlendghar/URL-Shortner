package utils

func MakeShortBase62(counter int64) string {
	base62 := "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ+/"
	var shortStr string = ""
	for counter > 0 {
		mod := counter % 62
		shortStr = shortStr + string(base62[mod])
		counter /= 62
	}
	return shortStr
}
