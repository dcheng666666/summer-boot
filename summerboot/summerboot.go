package summerboot

import (
	"net/http"
	"fmt"
	"os"
	"errors"
	"io/ioutil"
	"strings"
)


type Handler func (http.ResponseWriter, *http.Request)
type GroupHandler map[string]Handler

func matchPath(path string, pathStr string) bool {
	pathList := strings.Split(path, "/")
	pathStrList := strings.Split(pathStr, "/")
	pathParamMap := make(map[string]string)

	if len(pathList) != len(pathStrList) {
		return false
	}

	for idx := 0; idx< len(pathList); idx++ {
		if strings.HasPrefix(pathList[idx], ":") {
			pathParamMap[pathList[idx]] = pathStrList[idx]
		} else {
			if pathList[idx] != pathStrList[idx] {
				return false
			}
		}
	}

	return true
}

// SummerBoot definition
type SummerBoot struct {
	groupHandlerList map[string]GroupHandler
	staticResourcePath string
}

var summerBoot *SummerBoot

func GetInstance() *SummerBoot  {
	if (summerBoot == nil) {
		summerBoot = &SummerBoot{}
		summerBoot.init()
	}
	return summerBoot
}

func (summerBoot *SummerBoot)init()  {
	summerBoot.groupHandlerList = make(map[string]GroupHandler)
}

func (summerBoot *SummerBoot)findHandler(uri string, method string) (Handler, error) {
	for path, groupHandler := range GetInstance().groupHandlerList {
		if match := matchPath(path, uri); match && groupHandler != nil {
			if handler := groupHandler[method]; handler == nil {
				return nil, errors.New(fmt.Sprintf("method %s with uri %s not found", method, uri))
			} else {
				return handler, nil
			}
		}
	}

	return nil, errors.New(fmt.Sprintf("url %s not found", uri))
}

func defaultHandler(responseWriter http.ResponseWriter, request *http.Request) {
	uri := request.URL.Path

	if strings.HasSuffix(uri, ".html") {
		GetInstance().staticResourceHandler(uri, responseWriter)
	} else {
		GetInstance().restApiHandler(request, responseWriter)
	}
}

func (summerBoot *SummerBoot)staticResourceHandler(uri string, responseWriter http.ResponseWriter) {
	htmlPath := summerBoot.staticResourcePath + uri
	if _, err := os.Stat(htmlPath); os.IsNotExist(err) {
		responseWriter.WriteHeader(http.StatusNotFound)

		fmt.Fprintf(responseWriter, "%s", err)
	} else {
		if data, err := ioutil.ReadFile(htmlPath); err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(responseWriter, "%s", err)
		} else {
			fmt.Fprintf(responseWriter, "%s", data)
		}
	}
}

func (summerBoot *SummerBoot)restApiHandler(request *http.Request, responseWriter http.ResponseWriter) {
	uri := request.URL.Path
	method := strings.ToLower(request.Method)

	groupHandler, err := summerBoot.findHandler(uri, method)
	if (err != nil) {
		responseWriter.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(responseWriter, "%s", err)
	} else {
		groupHandler(responseWriter, request)
	}
}

func (summerBoot *SummerBoot)AddRoute(uri string, method string, handler func (http.ResponseWriter, *http.Request))  {
	groupHandler := summerBoot.groupHandlerList[uri]
	if groupHandler == nil {
		groupHandler = make(GroupHandler)
		summerBoot.groupHandlerList[uri] = groupHandler		
	}
	groupHandler[strings.ToLower(method)] = handler 
}

func (summerBoot *SummerBoot)SetStaticResource(path string)  {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(err)
	}

	summerBoot.staticResourcePath = path
}
	
func (summerBoot *SummerBoot)Start() {
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", nil)
}

