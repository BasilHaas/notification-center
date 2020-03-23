package main

import (
	"github.com/gorilla/mux"
	"net/http"
	make_message "notification-center/controllers"
)

//const (
//	ModeMarkdown = "Markdown"
//	ModeHTML     = "HTML"
//)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/telegram", make_message.GetInfo).Methods("POST")
	_ = http.ListenAndServe(":8000", r)
}
