package main

import (
	"flag"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

var (
	port       int
	staticPath string
)

func main() {
	startServer()
}

func startServer() {
  flag.IntVar(&port, "port", 3000, "The port to run the server on")
	flag.StringVar(&staticPath, "staticPath", "public", "The port to run the server on")

	flag.Parse()

	myMux := mux.NewRouter()

	myMux.HandleFunc("/", IndexRoute)
	myMux.HandleFunc("/{componentUrlPath:[a-zA-Z0-9\\-\\_\\/]*}", IndexRoute)
	myMux.HandleFunc("/static/{cache_id}/{filename:[a-zA-Z0-9\\.\\-\\_\\/]*}", FileServer)

	middleWare := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir(".")),
	)
	middleWare.UseHandler(myMux)
	middleWare.Run(":" + strconv.Itoa(port))
}

func IndexRoute(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", mime.TypeByExtension("html"))
	http.ServeFile(writer, request, "./" + staticPath + "/index.html")
}

func FileServer(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	file := vars["filename"]

	writer.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(file)))
	http.ServeFile(writer, request, "./" + staticPath + "/" + file)
}
