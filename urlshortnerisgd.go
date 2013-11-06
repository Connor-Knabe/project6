package main

import (
	  "fmt"
	  "code.google.com/p/odie"
    "flag"
    "urlgo"
)

func helloHandler(w *odie.ResponseWriter, req *odie.Request, vars odie.Context) {

	url := vars.Get("url")

	if url == "" {
		url = req.Context.Get("url")
    
    uri, err := goisgd.Shorten(url)
    if err != nil {
      fmt.Fprintln(os.Stderr,err)
	}

	if uri == "" {
    fmt.Fprintf(w, "Enter the url to be shortened:")
    w.Prompt("url:")
	} else {
		w.AddContext("url", uri)
    fmt.Fprintf(w, "Your shortened url is: %s", uri)
	}
}




func init () {

	odie.Handle("url", helloHandler) 
	odie.Handle("my name is $name", helloHandler) 
	odie.Handle("what is my name", whatHandler)
}

                
