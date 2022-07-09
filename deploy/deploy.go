package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", restart)
	http.HandleFunc("/done!", redirect)
	http.ListenAndServe(":5000", nil)
}

func relaunching() {
	cmd := exec.Command("sh", "./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func restart(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Deploy server: Pushing and restarting...please wait...</h1>")
	relaunching()
	http.Redirect(w, r, "127.0.0.1:5000/done!", 301)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Deploy server: Aplication successfully updated from github and works!</h1>")
}

//  env GOOS=linux GOARCH=amd64 go build  -o deploy deploy.go
