package main

import (
	"fmt"
	"net/http"
	"os"
)

var (
	addr    string = ":443"
	sitedir string = "resources"
	fs      http.Handler
)

const (
	msgfile  = "messages.txt"
	certfile = "/etc/letsencrypt/archive/gopher.ddns.net/cert1.pem"
	keyfile  = "/etc/letsencrypt/archive/gopher.ddns.net/privkey1.pem"
)

func init() {
	fs = http.FileServer(http.Dir(sitedir))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Printf("%s connected to /", r.RemoteAddr)
	}

	switch r.Method {
	// If its a post request log it.
	case "POST":
		err := r.ParseForm()
		if err != nil {
			fmt.Printf("ParseForm error: %v\n", err)
			return
		}

		name := r.FormValue("name")
		message := r.FormValue("message")

		if passesFilter(message) && passesFilter(name) && len(name) < 20 {
			data := fmt.Sprintf("[%s:%s] %s\n",
				r.RemoteAddr, name, message)
			fmt.Print(data)

			file, err := os.OpenFile(msgfile,
				os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)

			if err != nil {
				fmt.Println("Failed to open", msgfile)
				fmt.Println(err)
				break
			}

			defer file.Close()
			_, err = file.Write([]byte(data))
			if err != nil {
				fmt.Println("Failed to message write to file\n", err)
			}
		}
	}
	// Serve the code files
	fs.ServeHTTP(w, r)
}

// Yeah i get a custom handler, git gud git :L
func peepHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s entered the peep zone", r.RemoteAddr)
	fs.ServeHTTP(w, r)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s Connected via http, redirecting to https...\n", r.RemoteAddr)
	http.Redirect(w, r, "https://gopher.ddns.net", http.StatusSeeOther)
}

func passesFilter(thing string) bool {
	if len(thing) > 200 {
		return false
	}
	if thing == "" {
		return false
	}

	return true
}

func main() {
	// web redirect to https
	http.HandleFunc("/", httpHandler)
	go http.ListenAndServe(":80", nil)
	// Create server
	mux := http.NewServeMux()
	mux.HandleFunc("/peep", peepHandler)
	mux.HandleFunc("/", rootHandler)

	fmt.Printf("Listening on %s\n", addr)
	fmt.Println(http.ListenAndServeTLS(addr, certfile, keyfile, mux))
}

func handle(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
