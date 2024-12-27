package main

import (
	"bytes"
	"io"
	"log"
	"os/exec"
)

func main() {
	// go run code-user/main.go
	cmd := exec.Command("go", "run", "code-user/main.go")
	var out, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &out
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalln(err)
	}
	io.WriteString(stdinPipe, "23 11\n")

	// 根据测试的输入案例，进行运行，拿到输出结果和标准的输出的结果进行比对
	if err := cmd.Run(); err != nil {
		log.Fatalln(err, stderr.String())
	}
	println("Err:", string(stderr.String()))
	log.Println(out.String())

	// 判断结果
	println(out.String() == "34\n")
}
