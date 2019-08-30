package u5_string

import (
	"bytes"
	"strings"
)

func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}

func IterComma(s string) string {
	var b bytes.Buffer

	for i := 0; i < len(s); i++ {
		b.WriteByte(s[i])
		if i != len(s)-1 && (i+1)%3 == 0 {
			b.WriteByte(',')
		}
	}
	return b.String()
}

func SupportFloatComma(s string) string {
	// 浮点数
	var b bytes.Buffer
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		b.WriteByte(s[0])
		s = s[1:]
	}
	if strings.Index(s, ".") != -1 {
		sl := strings.Split(s, ".")
		//ibase += IterComma(sl[0]) + "." + sl[1]
		b.WriteString(IterComma(sl[0]))
		b.WriteString(IterComma("."))
		b.WriteString(IterComma(sl[1]))
	} else {
		b.WriteString(IterComma(s))
	}
	return b.String()
}

func HasSameCharacters(s1 string, s2 string) bool {
	//temp1 := []rune(s1)
	//temp2 := []rune(s2)

	for i, v := range s1 {
		if strings.Index(s2, string(v)) == -1{
			return false
		}
		println(i,"  ", v)
	}
	for i, v := range s2 {
		if strings.Index(s1, string(v)) == -1{
			return false
		}
		println(i,"  ", v)
	}
	return true
}
