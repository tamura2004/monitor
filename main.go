package main

import (
	"github.com/pressly/chi"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		cmd := exec.Command("netstat", "-na")
		out, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		cmd.Start()
		io.Copy(w, out)
	})

	http.ListenAndServe(":3333", r)
}
