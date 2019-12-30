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
