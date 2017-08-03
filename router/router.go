package router

import (
	"fmt"
	"net/http"
	"build_golang_web_station/controller"
)

func init()  {
	fmt.Println("loading router")
}


func ImportRouter(){
	http.HandleFunc("/index",controller.Index)
}



