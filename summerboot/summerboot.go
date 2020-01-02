package summerboot

import (
	"net/http"
)

// SummerBoot definition
type SummerBoot struct {

}

func (summerBoot *SummerBoot)AddRoute(uri string, handler func (http.ResponseWriter, *http.Request))  {
	http.HandleFunc(uri, handler)
}

func (summerBoot *SummerBoot)Start() {
	http.ListenAndServe(":8080", nil)
}

// func AddRoute(summerBoot *SummerBoot, uri string, handler func (http.ResponseWriter, *http.Request))  {
// 	http.HandleFunc(uri, handler)
// }

// func Start(summerBoot *SummerBoot) {
// 	http.ListenAndServe(":8080", nil)
// }
