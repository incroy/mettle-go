package workspace

import "os/exec"

func Generate(args ...string) {
	name := args[0]
	cmd := exec.Command("go mod init ", name)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
