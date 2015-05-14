package keyword

import (
	"strings"
)

type Checker struct {
	sensitive bool
	keyword   []string
	exclude   []string
}

func New(keyword, exclude string, sensitive bool) *Checker {
	//lower case all if no sensitive
	if !sensitive {
		keyword = strings.ToLower(keyword)
		exclude = strings.ToLower(exclude)
	}

	//split
	keywordArr := strings.Split(keyword, ",")
	excludeArr := strings.Split(exclude, ",")

	//trim
	keywordArr = trimArr(keywordArr)
	excludeArr = trimArr(excludeArr)

	return &Checker{
		keyword:   keywordArr,
		exclude:   excludeArr,
		sensitive: sensitive,
	}
}

func (k *Checker) CheckKeyword(str string) bool {
	if !k.sensitive {
		str = strings.ToLower(str)
	}

	// if keyword is empty -> always true
	if len(k.keyword) == 0 {
		return true
	}

	return test(k.keyword, str)
}

func (k *Checker) CheckExclude(str string) bool {
	if !k.sensitive {
		str = strings.ToLower(str)
	}

	return test(k.exclude, str)
}

func (k *Checker) Check(str string) bool {
	return k.CheckKeyword(str) && !k.CheckExclude(str)
}

func test(strArr []string, str string) bool {
	for i := 0; i < len(strArr); i++ {
		if strings.Contains(str, strArr[i]) {
			return true
		}
	}

	return false
}

func trimArr(arr []string) (trimmedArr []string) {
	for i := 0; i < len(arr); i++ {
		str := strings.Trim(arr[i], " ")

		//remove empty string
		if str == "" {
			continue
		}

		trimmedArr = append(trimmedArr, str)
	}
	return
}
