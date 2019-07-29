package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	log.Fatal(http.ListenAndServe(":4721", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "could not read body: "+err.Error(), http.StatusInternalServerError)
			return
		}
		resource := strings.TrimSpace(string(buf))

		cmd := exec.Command("open", resource)
		out, err := cmd.CombinedOutput()
		if err != nil {
			errMsg := "executing command: " + err.Error() + "\n" + string(out)
			http.Error(w, errMsg, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})))
}
