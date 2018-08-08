package main

import (    "net/http"  //http services
            "fmt"       //data output
            "time"      //for the sleep
            "log"       //error handling
       )

func main() {
    http.HandleFunc("/hash", hash)  //add a Handler for /hash
    log.Fatal(http.ListenAndServe(":8080", nil))    //listen on port 8080
}

func hash(w http.ResponseWriter, r *http.Request) {
    password := r.PostFormValue("password") //get the POST data
    time.Sleep(time.Millisecond*5000)       //wait 5 seconds
    fmt.Fprintf(w, "%s\n", encodePassword(password))    //output the hash
}

