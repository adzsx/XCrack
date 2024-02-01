package utils

import (
	"regexp"
)

var (
	reg map[string]string
)

func init() {
	reg = make(map[string]string)

	reg["md5"] = "^[a-fA-F0-9]{32}$"
	reg["sha1"] = "^[a-fA-F0-9]{40}$"
	reg["sha256"] = "^[a-fA-F0-9]{64}$"
	reg["sha512"] = "[a-fA-F0-9]{128}$"
}

func GetHashType(hash string) string {
	for key, element := range reg {
		r, _ := regexp.Compile(element)
		if r.MatchString(hash) {
			return key
		}
	}
	return ""
}
