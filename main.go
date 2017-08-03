package main

import (
	"fmt"
	"build_golang_web_station/router"
	"net/http"
)


func main(){
	fmt.Println("start serve:","gogogo")
	router.ImportRouter()
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		fmt.Println("error serve")
	}
}

