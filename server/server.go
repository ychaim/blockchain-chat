package server

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"mime"
	"github.com/gorilla/mux"
)
func createIndexHandler(rootDir string) http.HandlerFunc {
	return func (resp http.ResponseWriter, req * http.Request) {
		resp.Header().Add("Content-Type", "text/html")

		content, err := ioutil.ReadFile(path.Join(rootDir, "index.html"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		resp.Write(content)
	}
}

func createPathHandler(rootDir string) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		dir := mux.Vars(req)["dir"]
		file := mux.Vars(req)["path"]
		ext := filepath.Ext(file)
		resp.Header().Add("Content-Type", mime.TypeByExtension(ext))

		content, err := ioutil.ReadFile(path.Join(rootDir, dir, file))
		if err != nil || content == nil {
			fmt.Println(err.Error())
			resp.WriteHeader(404)
			resp.Write([]byte{})
			return
		}

		resp.Write(content)
	}
}

// Запускает файловый сервер на порту port, считая корневой директорией htmlDir
func Run(rootDir string, port string) {
	router := mux.NewRouter()

	//r.Methods("GET").Path("/websocket").HandlerFunc(api.WebsocketHandler(srv, log))
	router.Methods("GET").Path("/").HandlerFunc(createIndexHandler(rootDir))
	router.Methods("GET").Path("/{dir}/{path}").HandlerFunc(createPathHandler(rootDir))

	http.Handle("/", router)

	err := http.ListenAndServe("127.0.0.1:" + port, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}