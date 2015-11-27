package main
import (
	"net/http"
	"log"
)


const port = "8080"
const assets = "assets"
const tpl = "tpl"
func main(){

	http.Handle("/assets/", http.FileServer(http.Dir("")))
	http.HandleFunc("/", serveEn)
	http.HandleFunc("/da", serveDa)
	http.HandleFunc("/da/", redirectDK)
	http.HandleFunc("/robots.txt", serveFile)
	http.HandleFunc("/sitemap.xml", serveFile)
	log.Println("Listening on port: ", port)
	http.ListenAndServe(":"+port, nil)


	//http.Handle("/da", http.StripPrefix("/d/", http.FileServer(http.Dir("assets"))))
	//log.Fatal(http.ListenAndServe(":"+port, http.FileServer(http.Dir("assets"))))
}
func serveEn(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/"{
		http.Redirect(rw, req, `http://eutopia.fail`, 302)
		return
	}
	http.ServeFile(rw,req, "tpl/index.html")
}

func serveDa(rw http.ResponseWriter, req *http.Request) {
	http.ServeFile(rw,req, "tpl/da/index.html")
}

func redirectDK(rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, `http://eutopia.fail/da`, 302)
}

func serveFile(rw http.ResponseWriter, req *http.Request) {
	http.ServeFile(rw, req, "etc" + req.URL.Path)
}




