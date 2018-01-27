package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		cmd := exec.Command("vmstat")
		out, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		cmd.Start()
		io.Copy(w, out)
	})

	http.ListenAndServe(":3333", nil)

}
