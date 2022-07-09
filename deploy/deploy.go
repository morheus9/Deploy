package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
)

func relaunching() {
	cmd := exec.Command("sh", "./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func restart(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "<h1>Deploy server: Pushing and restarting...please wait...</h1>")
	relaunching()
	io.WriteString(w, "<h1>Deploy server: Aplication successfully updated from github and works!</h1>")
}

func main() {
	http.HandleFunc("/", restart)
	http.ListenAndServe(":5000", nil)
}

//  env GOOS=linux GOARCH=amd64 go build  -o deploy deploy.go
