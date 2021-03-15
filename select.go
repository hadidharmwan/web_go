package main

import (
	"net/http"
	"log"
	"html/template"
)

type RadioButton struct {
	Name string
	Value string
	IsDisabled bool
	IsChecked bool
	Text string
}

type PageVariables struct {
	PageTitle string
	PageRadioButtons []RadioButton
	Answer string
}

func main() {
	http.HandleFunc("/", DisplayRadioButtons)
	http.HandleFunc("/selected", UserSelected)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func DisplayRadioButtons(w http.ResponseWriter, r *http.Request){

}

func UserSelected(w http.ResponseWriter, r *http.Request){
	
}