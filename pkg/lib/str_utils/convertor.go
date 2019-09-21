package str_utils

import (
	"strings"
	"unicode"
)

func ConvertToCamelFormat(lowerStr string) string {
	tmpStrs := []string{}
	splitedLowerStr := strings.Split(lowerStr, ".")
	for _, str := range splitedLowerStr {
		tmpStrs = append(tmpStrs, strings.ToUpper(str[:1])+strings.ToLower(str[1:]))
	}
	return strings.Join(tmpStrs, "")
}

func ConvertToLowerFormat(camelStr string) string {
	runes := []rune{}
	for i, r := range camelStr {
		if i == 0 {
			runes = append(runes, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			runes = append(runes, '.', unicode.ToLower(r))
		} else {
			runes = append(runes, r)
		}
	}
	return string(runes)
}

func SplitActionDataName(name string) (string, string) {
	actionRunes := []rune{}
	dataRunes := []rune{}
	isAction := false
	for i, r := range name {
		if i == 0 {
			isAction = true
		} else if unicode.IsUpper(r) {
			isAction = false
		}
		if isAction {
			actionRunes = append(actionRunes, r)
		} else {
			dataRunes = append(dataRunes, r)
		}
	}

	return string(actionRunes), string(dataRunes)
}
