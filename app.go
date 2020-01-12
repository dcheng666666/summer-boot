package main

import (
	"fmt"
	"net/http"
	summerboot "./summerboot"
)

func getOpportunity(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "getOpportunity")
}

func getContactByOpportunity(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "getContactByOpportunity")
}

func main()  {
	summerBoot := summerboot.GetInstance()
	summerBoot.AddRoute("/opportunity", "GET", getOpportunity)
	summerBoot.AddRoute("/opportunity/:id/contact/:id", "GET", getContactByOpportunity)
	summerBoot.SetStaticResource("/Users/dcheng/Github/psa-be-gopher-web-framework/summer-boot/static/")
	summerBoot.Start()
}
