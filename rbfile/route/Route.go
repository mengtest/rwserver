package route

import (
	"github.com/gorilla/mux"
	TQ "../../rbwork/base"
	"net/http"
	Ctrl "../controller"
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
		r.ParseForm()
		inner.ServeHTTP(w, r)
		TQ.LogInfo("completed:"+r.RequestURI," time:", time.Since(start))
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
	Route{Name: "Index", Method: "GET", Pattern: "/", HandlerFunc: Ctrl.Index},
	Route{Name: "Upload", Method: "GET", Pattern: "/upload", HandlerFunc: Ctrl.Upload},
	Route{Name: "Download", Method: "GET", Pattern: "/download/{fileName}", HandlerFunc: Ctrl.Download},//下载指定文件

}


