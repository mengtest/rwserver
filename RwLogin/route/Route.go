package route

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	controller "../controller"
	"time"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
//输出请求执行时间
func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf("Request===>: %s\t%s\t%s\t%s", r.Method, r.RequestURI, name, time.Since(start))
	})
}

//定义路由
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{Name: "Index", Method: "GET", Pattern: "/", HandlerFunc: controller.Index},
	Route{Name: "Login", Method: "GET", Pattern: "/user/login", HandlerFunc: controller.Login},
	Route{Name: "Register", Method: "GET", Pattern: "/user/register", HandlerFunc: controller.Register},
}


