package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

func (app *application) numberfunction() int {
	return 1
}

func testnew() int {
	return 2
}

func newcode() string {
	return "rohan"
}

func msama() int {
	return 7
}
