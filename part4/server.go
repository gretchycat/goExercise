package main

import (
            "encoding/json" //stats output
            "context"       //for the Shutdown call
            "net/http"      //http services
            "fmt"           //formatted string output
            "time"          //for the sleep and timing queries
       )

type serverInfo struct {
    activeConnections int    //define memory for connection count
    totalTime int            //define memory for total time count
    countQueries int         //define memory for total queries
}

var srvInfo *serverInfo = new(serverInfo)
var srv=&http.Server{Addr: ":8080"}     //initialize server parameters to use port 80

func main() {
    srvInfo.activeConnections = 0           //initialize connection count to 0
    srvInfo.countQueries = 0                //initialize number of queries to 0
    srvInfo.totalTime = 0                   //initialize the time counter to 0
    http.HandleFunc("/hash", hash)          //add a Handler for /hash
    http.HandleFunc("/shutdown", shutdown)  //add a Handler for /shutdown
    http.HandleFunc("/stats", stats)        //add a Handler for /stats
    srv.ListenAndServe()                    //listen on port 8080
}

func hash(w http.ResponseWriter, r *http.Request) {
    srvInfo.activeConnections++
    start := time.Now()                                 //get initial time to calculate query time
    password := r.PostFormValue("password")             //get the POST data
    time.Sleep(time.Millisecond*5000)                   //wait 5 seconds
    fmt.Fprintf(w, "%s\n", encodePassword(password))    //output the hash
    srvInfo.activeConnections--
    srvInfo.countQueries++                                     //only count completed queries for stats
    srvInfo.totalTime+=int(time.Now().Sub(start)/1000000)      //count time for stats
}

func shutdown(w http.ResponseWriter, r *http.Request) { //TODO: I'd think the Shutdown method would be better behaved, but in practice it quits immediately
    for (srvInfo.activeConnections>0) {}    //TODO: wait for all active connections. A sleep in here would make it less CPU intensive
    srv.Shutdown( context.Background())     //TODO: this call doesnt seem to want to wait for active connections to naturally close. Use a different context?
}

func stats(w http.ResponseWriter, r *http.Request) {
    avg:=0                                  //default the average time to zero
    if(srvInfo.countQueries>0) {
        avg=srvInfo.totalTime / srvInfo.countQueries    //cant divide by zero
    }
    encoder := json.NewEncoder(w)           //create the decoder and output to the web stream
    data := map[string]int{"total":srvInfo.countQueries, "average": avg}    //set the data in memory
    encoder.Encode(data)                    //encode and output the encoded structure to the stream
}
