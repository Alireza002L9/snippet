package main

import (
	"flag"
	"log"
	"net/http"
	// "os"
)

func main() {

	addr := flag.String("addr", ":4000", "Http Network Address")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Importantly, we use the flag.Parse() function to parse the command-line flag.
	// This reads in the command-line flag value and assigns it to the addr
	// variable. You need to call this *before* you use the addr variable
	// otherwise it will always contain the default value of ":4000". If any errors are
	// encountered during parsing the application will be terminated.
	flag.Parse()
	// we also have flag.Int() , flag.Bool() and flag.Float64()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)


	srv:= &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}

	// Ports 0-1023 are restricted and (typically) can only be used by services 
	// which have root privileges.
	// use -help to see all flag line commands
	infoLog.Printf("Starting server on http://localhost%s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
	// avoid using the Panic() and Fatal() variations outside ofyour main() function
}

// ***Logging to a file***
// f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
// if err != nil {
// log.Fatal(err)
// }
// defer f.Close()
// infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)