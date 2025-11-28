package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func run(prog string, args ...string) (string, error) {
	bs, err := exec.Command(prog, args...).CombinedOutput()
	rst := ""
	if err == nil {
		rst = string(bs)
	}
	return rst, err
}

const PKG_NAME = "github.com/KevinZonda/GinTemplate"

func main() {
	tgtName := os.Args[1]
	fmt.Println("[INFO] Renaming:", PKG_NAME, "->", tgtName)
	_, err := run("bash", "-c", replaceAllCmd(".", PKG_NAME, tgtName))
	if err != nil {
		fmt.Println("[ERROR] replaceAllCmd error:", err)
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("[SUCCESS] Renamed:", PKG_NAME, "->", tgtName)
}

func getSed() string {
	bin := "sed"
	if runtime.GOOS == "darwin" {
		bin = "gsed"
	}
	fmt.Println("[INFO] sed application:", bin)
	return bin
}

func replaceAllCmd(path string, strFrom string, stringTo string) string {
	return fmt.Sprintf("grep -rl '%s' '%s' | xargs %s -i 's@%s@%s@g'", strFrom, path, getSed(), strFrom, stringTo)
}
