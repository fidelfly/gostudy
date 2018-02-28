package basic

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCmdTest() {
	cmd := exec.Command("/bin/bash", "test.bash")
	cmd.Dir = "/Users/fidelfly/workshop"
	//var out bytes.Buffer
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(out.String())
}
