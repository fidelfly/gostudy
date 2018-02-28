package basic

import (
	"fmt"
	"strconv"
	"strings"
)

func CompareVersion(ver1 string, ver2 string) int {
	ver1Value := ResolveVersion(ver1)
	ver2Value := ResolveVersion(ver2)

	for i := 0; i < 3; i++ {
		if ver1Value[i] == ver2Value[i] {
			continue
		} else if ver1Value[i] > ver2Value[i] {
			return 1
		} else {
			return -1
		}
	}

	return 0
}

func ResolveVersion(version string) [3]int64 {
	result := [3]int64{0, 0, 0}

	if len(version) > 0 {
		index := strings.Index(version, "_")
		if index > 0 {
			rs := []rune(version)
			rs = rs[:index]

			version = string(rs)
		}

		vers := strings.Split(version, ".")
		for pos, ver := range vers {
			if pos < 3 {
				result[pos], _ = strconv.ParseInt(ver, 10, 64)
			} else {
				break
			}
		}
	}

	return result
}

func RunVersionTest() {
	ResolveVersion("1.2.393847_test")

	fmt.Println("Compare 1.3.234556_xxxx and 1.5.345_yyy : ", CompareVersion("1.5.234556_xxxx", "1.5.345_yd"))
}
