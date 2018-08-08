package main

import (
            "context"   //for the Shutdown call.
            "net/http"  //http services
            "fmt"       //data output
            "time"      //for the sleep
       )

var activeConnections *int  = new(int)  //define memory for connection count
var srv=&http.Server{Addr: ":8080"}     //initialize server parameters to use port 80

func main() {
    *activeConnections = 0                  //initialize connection count to 0
    http.HandleFunc("/hash", hash)          //add a Handler for /hash
    http.HandleFunc("/shutdown", shutdown)  //add a Handler for /shutdown
    srv.ListenAndServe()                    //listen on port 8080
}

func hash(w http.ResponseWriter, r *http.Request) {
    *activeConnections++
    password := r.PostFormValue("password")             //get the POST data
    time.Sleep(time.Millisecond*5000)                   //wait 5 seconds
    fmt.Fprintf(w, "%s\n", encodePassword(password))    //output the hash
    *activeConnections--
}

func shutdown(w http.ResponseWriter, r *http.Request) {
    for (*activeConnections>0) {}       //wait for all active connections. A sleep in here would make it less CPU intensive
    srv.Shutdown( context.Background()) //this call doesnt seem to want to wait for active connections to naturally close. Different context?
}
