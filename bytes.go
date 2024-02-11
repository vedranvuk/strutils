package strutils

func IsUpper(c byte) bool { return c >= 'A' && c <= 'Z' }

func IsLower(c byte) bool { return c >= 'a' && c <= 'z' }

func IsLetter(c byte) bool { return IsLower(c) || IsUpper(c) }

func IsDigit(c byte) bool { return c >= '0' && c <= '9' }

func IsAlphanumeric(c byte) bool { return IsLower(c) || IsUpper(c) || IsDigit(c) }

func ToLower(c byte) byte {
	if IsUpper(c) {
		return c + ('a' - 'A')
	}
	return c
}

func ToUpper(c byte) byte {
	if IsLower(c) {
		return c - ('a' - 'A')
	}
	return c
}
