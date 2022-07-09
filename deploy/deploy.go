package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", restart)
	http.ListenAndServe(":5000", nil)
}

func relaunching() {
	cmd := exec.Command("sh", "./deploy.sh")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func restart(w http.ResponseWriter, _ *http.Request) {
	relaunching()
	io.WriteString(w, "<h1>Deploy server: Aplication successfully updated from gitlab and works!</h1>")
}

//  env GOOS=linux GOARCH=amd64 go build  -o deploy deploy.go
