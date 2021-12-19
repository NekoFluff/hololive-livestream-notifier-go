package utils

import "regexp"

/**
 * Parses content string with the given regular expression and returns the
 * group values defined in the expression.
 *
 */
func GetParams(regEx, content string) (paramsMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(content)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}
