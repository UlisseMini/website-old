package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
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
	certfile = "/letsencrypt/cert1.pem"
	keyfile  = "/letsencrypt/privkey1.pem"
	logfile  = "/logs/latest.log"
)

func init() {
	fs = http.FileServer(http.Dir(sitedir))
}

func main() {
	// initalize loggers
	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open %s: %v\n", logfile, err)
	}
	defer f.Close()
	log.SetOutput(f)

	// web redirect to https
	http.HandleFunc("/", httpHandler)
	go func() {
		log.Errorln(http.ListenAndServe(":80", nil))
	}()

	// Create serve mux
	mux := http.NewServeMux()
	mux.HandleFunc("/peep", peepHandler)
	mux.HandleFunc("/", rootHandler)

	// Start listening
	log.Infof("Listening on %s\n", addr)
	err = http.ListenAndServeTLS(addr, certfile, keyfile, mux)
	log.Fatal(err)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		log.Debugf("%s connected to /", r.RemoteAddr)
	}

	switch r.Method {
	// If its a post request log it.
	case "POST":
		err := r.ParseForm()
		if err != nil {
			log.Errorf("ParseForm error: %v\n", err)
			return
		}

		name := r.FormValue("name")
		message := r.FormValue("message")

		if passesFilter(message) && passesFilter(name) && len(name) < 20 {
			data := fmt.Sprintf("[%s:%s] %s\n",
				r.RemoteAddr, name, message)
			log.Info(data)

			file, err := os.OpenFile(msgfile,
				os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)

			if err != nil {
				log.Errorf("Failed to open %s: %v\n", msgfile, err)
				break
			}

			defer file.Close()
			_, err = file.Write([]byte(data))
			if err != nil {
				log.Errorf("Failed to message write to file: %v\n", err)
			}
		}
	}
	// Serve the code files
	fs.ServeHTTP(w, r)
}

// Yeah i get a custom mustNotr, git gud git :L
func peepHandler(w http.ResponseWriter, r *http.Request) {
	log.Tracef("%s entered the peep zone", r.RemoteAddr)
	fs.ServeHTTP(w, r)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	log.Tracef("%s Connected via http, redirecting to https...\n", r.RemoteAddr)
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
