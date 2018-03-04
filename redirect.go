/*
 * Copyright (C) 2018 Giuliano Pasqualotto (github.com/giulianopa)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
 package main

 import (
     "fmt"
     "log"
     "net/http"
     "github.com/gorilla/mux"
 )


//-------
// Types
//-------

// Remote host descriptor.
type RemoteHost struct {
    ip string   // RemoteAddr, from the original HttpRequest.
}


//------------
// Parameters
//------------

const SERVICE_PORT string = ":8080"
var Hosts []RemoteHost


//-----------
// Functions
//-----------

// Redirect and log
 func Redirect(w http.ResponseWriter, r *http.Request) {
     vars := mux.Vars(r)
     url := vars["url"]
     logMsg := r.RemoteAddr + "," + url
     Hosts = append(Hosts, RemoteHost{ip: logMsg + ";"})
     fmt.Printf(logMsg + "\n")
     http.Redirect(w, r,  "http://" + url, 301)
 }

// Log collected information
 func Log(w http.ResponseWriter, r *http.Request) {
     fmt.Fprintf(w, "%v\n", Hosts)
 }

// Initialize
func init() {
    // NOOP
}


//------
// Main
//------

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/log/", Log)
    router.HandleFunc("/r/{url}", Redirect)
    log.Fatal(http.ListenAndServe(SERVICE_PORT, router))
}
