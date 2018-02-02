package basic

import (
	"time"
	"fmt"
)

func RunTimeFormatTest() {
	now := time.Now()

	fmt.Println("now.String():" + now.String())

	fmt.Println("Format(2006-01-02) : " + now.Format("2006-01-02"))

	fmt.Println(now.Unix())
}
