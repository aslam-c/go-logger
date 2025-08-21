package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    query:=r.URL.Query()
    a,_ :=strconv.Atoi(query.Get("a"))
    b,_ :=strconv.Atoi(query.Get("b"))
    
    fmt.Fprintf(w, "ഈ ലിങ്ക് തുറന്നവന് ഊള %s!%d", time.August,a+b)
    
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
