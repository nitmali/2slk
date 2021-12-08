package utils

import (
	"log"
	"os"
)

var HOST string
var HTTPS string
var PROXY bool
var PORT string

var URL string

func Init_env_vars() {
	UHOST, err := os.LookupEnv("HOST")
	if !err {
		log.Fatal("HOST enviroment variable not found, please set it!")
	}
	HOST = UHOST

	UHTTPS, _ := os.LookupEnv("HTTPS")
	if UHTTPS != "true" {
		HTTPS = "http"
	} else {
		HTTPS = "https"
	}

	UPROXY, _ := os.LookupEnv("PROXY")
	if UPROXY != "true" {
		PROXY = false
	} else {
		PROXY = true
	}

	UPORT, err := os.LookupEnv("PORT")
	if !err {
		PORT = "3000"
	} else {
		PORT = UPORT
	}

	create_string()

}

func create_string() {
	if !PROXY {
		URL = HTTPS + "://" + HOST + ":" + PORT + "/"
	} else {
		URL = HTTPS + "://" + HOST + "/"
	}
}