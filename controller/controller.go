package controller

import (
	"net/http"
	"fmt"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fmt.Fprintf(w, r.URL.Path+"\r\n")

	for k,v := range r.Form {
		fmt.Fprintf(w, "key:"+k+"\r\n")
		fmt.Fprintf(w, "val:"+strings.Join(v,"")+"\r\n")
	}
}
