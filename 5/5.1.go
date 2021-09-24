package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",echo)
	err := http.ListenAndServe("127.0.0.1:9486",nil)
	if err != nil{
		log.Fatal(err)
	}
}

func echo(wr http.ResponseWriter,r *http.Request){

}