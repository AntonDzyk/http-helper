package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func codes(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")
	intCode, err := strconv.Atoi(code)
	statusText := http.StatusText(intCode)

	if err != nil || (intCode < 100 || intCode > 999) || statusText == "" {
		notExists(w, code)
		return
	}

	w.WriteHeader(intCode)
	_, err = w.Write([]byte(`<p>code: ` + code + `<br>message: ` + statusText + `</p>`))
	if err != nil {
		log.Printf("error: %s\n", err.Error())
	}
}

func timeout(w http.ResponseWriter, r *http.Request) {
	timeoutValue := r.URL.Query().Get("timeout")
	tm, err := strconv.Atoi(timeoutValue)
	if err != nil {
		sendMessage(w, "timeout is not valid: "+timeoutValue)
	}

	time.Sleep(time.Duration(tm) * time.Second)
	sendMessage(w, "OK")
}

func notExists(w http.ResponseWriter, code string) {
	sendMessage(w, "<b>This http code does not exist - '"+code+"'</b>")
}

func sendMessage(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Printf("error: %s\n", err.Error())
	}
}

func helper(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(
		[]byte(`<p>This is test service<br>Use entrypoint /http?code=NUMBER and get response with code NUMBER</p>`))
	if err != nil {
		log.Printf("error: %s\n", err.Error())
	}
}

func main() {
	http.HandleFunc("/", helper)
	http.HandleFunc("/http", codes)
	http.HandleFunc("/wait", timeout)
	status := http.ListenAndServe(":8080", nil)
	fmt.Printf("Finish Code: %s\n", status)
}
