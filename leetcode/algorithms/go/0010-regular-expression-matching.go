package leetcode

func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	if 1 < len(p) && p[1] == '*' {
		return isMatch(s, p[2:]) || s != "" && (s[0] == p[0] || p[0] == '.') && isMatch(s[1:], p)
	}
	return s != "" && (s[0] == p[0] || p[0] == '.') && isMatch(s[1:], p[1:])
}

// import "regexp"
// func isMatch(s string, p string) bool {
// 	re := regexp.MustCompile("^" + p + "$")
// 	return re.MatchString(s)
// }

// https://golang.org/pkg/regexp/
