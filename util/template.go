package util

import (
	"bytes"
	"text/template"
	"unicode"
)

func Render(tpl string, data interface{}) ([]byte, error) {
	t, err := template.New("").Funcs(fm).Parse(tpl)
	if err != nil {
		return []byte{}, err
	}

	buf := bytes.NewBufferString("")
	err = t.Execute(buf, data)
	if err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

var fm template.FuncMap = map[string]interface{}{
	"ToTitle": ToTitle,
}

// order_id  =>  OrderId
func ToTitle(s string) string {
	results := []rune{}
	flag := true
	for _, r := range s {
		if flag {
			results = append(results, unicode.ToUpper(r))
			flag = false
			continue
		}
		if r == rune('_') {
			flag = true
			continue
		}
		results = append(results, r)
	}
	return string(results)
}
