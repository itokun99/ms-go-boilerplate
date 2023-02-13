package util

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

func LowerCamelCase(v string) string {
	s := strings.Split(v, "")
	s[0] = strings.ToLower(s[0])

	res := strings.Join(s, "")

	return res
}

func ToLowerCamelCase(input string) string {
	return strcase.ToLowerCamel(input)
}

func CustomLowerCamelCase(input string) string {
	var str string
	var tempArr []string
	var res string
	var tempsStr []string
	if strings.ContainsAny(input, "-") {
		tempsStr = strings.Split(input, "-")
	} else if strings.ContainsAny(input, "_") {
		tempsStr = strings.Split(input, "_")
		fmt.Println(tempsStr)
	} else {
		fmt.Println(ToLowerCamelCase("PascalCase"))
	}

	if tempArr != nil {
		str = strings.ToLower(tempsStr[0])
		tempArr = append(tempArr, str)
		for i := 1; i < len(tempsStr); i++ {
			tempS2 := strings.Split(tempsStr[i], "")
			joinTempS2 := strings.Join(tempS2[1:], "")
			lowerS2 := strings.ToLower(joinTempS2)
			tempArr = append(tempArr, tempS2[0]+lowerS2)
		}
	}
	res = strings.Join(tempArr, "")

	return res
}
