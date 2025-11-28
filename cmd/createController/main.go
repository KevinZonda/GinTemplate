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

const TEMPLATE_NAME = "template"

func main() {
	tgtName := os.Args[1]
	run("cp", "-r", "controller/"+TEMPLATE_NAME, "controller/"+tgtName)
	_, err := run("bash", "-c", replaceAllCmd("controller/"+tgtName, TEMPLATE_NAME, tgtName))
	if err != nil {
		fmt.Println("replaceAllCmd error:", err)
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Controller Created:", tgtName)
}

func getSed() string {
	if runtime.GOOS == "darwin" {
		return "gsed"
	}
	return "sed"
}

func replaceAllCmd(path string, strFrom string, stringTo string) string {
	return fmt.Sprintf("grep -rl '%s' '%s' | xargs %s -i 's/%s/%s/g'", strFrom, path, getSed(), strFrom, stringTo)
}
