package main

import (
	"bytes"
	"fmt"
	"go/build"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
)

const (
	// http://127.0.0.1:8083
	webComponentsV0Token = "AkuXKVDyq7S8wMrSAh8heevmtL2E0S7p7qOOETSiIdNc0EA6JNOu88v9gO5TlPZ6ik33U7iquKiiVOjLpJDRcAoAAABSeyJvcmlnaW4iOiJodHRwOi8vMTI3LjAuMC4xOjgwODMiLCJmZWF0dXJlIjoiV2ViQ29tcG9uZW50c1YwIiwiZXhwaXJ5IjoxNjEyMjIzOTk5fQ=="
)

func main() {
	pkg, err := build.Default.Import("cmd", "", build.FindOnly)
	if err != nil {
		fmt.Println("err:", err)
	}

	tracePath := pkg.Dir + filepath.FromSlash("/trace/trace.go")
	traceContent, err := ioutil.ReadFile(tracePath)
	if err != nil {
		log.Fatal(err)
	}

	patchHead := fmt.Sprintf(`<head><meta http-equiv="origin-trial" content="%s">`, webComponentsV0Token)
	patchedContent := bytes.Replace(traceContent, []byte("<head>"), []byte(patchHead), 1)
	err = ioutil.WriteFile(tracePath, patchedContent, 0644)
	if err != nil {
		fmt.Println("err:", err)
	}

	_, err = exec.Command("go", "install", "cmd/trace").CombinedOutput()
	if err != nil {
		fmt.Println("err:", err)
	}
}
