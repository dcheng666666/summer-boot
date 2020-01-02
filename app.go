package main

import (
	"fmt"
	"net/http"
	"html"
	summerboot "./summerboot"
)

func httpHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello, %s", html.EscapeString(r.URL.EscapedPath()))
}

func main()  {
	summerBoot := &summerboot.SummerBoot{}
	summerBoot.AddRoute("/", httpHandler)
	summerBoot.Start()
}
