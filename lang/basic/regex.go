package basic

import (
	"fmt"
	"os"
	"regexp"
)

func RunPathRegexTest() {
	var path = "caw_ear.ear\\jsf.war\\WEB-INF\\lib\\core_jsf.jar/xxx"

	pgex, _ := regexp.Compile("[\\\\/]")

	if pgex != nil {
		path = pgex.ReplaceAllString(path, string(os.PathSeparator))
		//path = pgex.ReplaceAllString(path, "\\")
	}
	//test
	fmt.Println(path)
}
