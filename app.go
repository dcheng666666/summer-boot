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
	summerBoot := summerboot.GetInstance()
	summerBoot.AddRoute("/aa", "GET", httpHandler)
	summerBoot.SetStaticResource("/Users/dcheng/Github/psa-be-gopher-web-framework/summer-boot/static/")
	summerBoot.Start()
}
