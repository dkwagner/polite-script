package util

func IsLetter(ch byte) bool {

	if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
		return true
	}

	return false
}

func IsDigit(ch byte) bool {

	if ch >= '0' && ch <= '9' {
		return true
	}

	return false
}

func IsWhitespace(ch byte) bool {

	if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' || ch == '\v' || ch == '\f' {
		return true
	}

	return false
}

func ContainsString(testString string, list []string) bool {

	for _, str := range list {

		if str == testString {
			return true
		}
	}

	return false
}

func ContainsByte(testByte byte, list []byte) bool {

	for _, b := range list {

		if b == testByte {
			return true
		}
	}

	return false
}
