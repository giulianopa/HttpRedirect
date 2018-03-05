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
    "time"
    "net/http"
    "container/ring"
    "github.com/gorilla/mux"
)


//------------
// Parameters
//------------

const SERVICE_PORT string = ":8080"
const REDIRECT_PROTO string = "http://"
const RING_BUF_SZ int = 100
var Hosts *ring.Ring


//-------
// Types
//-------

// Remote host descriptor.
type RemoteHost struct {
    ts time.Time    // Timestamp.
    ip string       // RemoteAddr, from the original HttpRequest.
    destUrl string  // The URL the remote address was redirected to.
}

// Conversion to string
/*
    Called when printing a RemoteHost instance.
 */
func (h RemoteHost) String() string {
    return fmt.Sprintf("%v;%v;%v;", h.ts.Format(time.RFC3339), h.ip, h.destUrl)
}


//-----------
// Functions
//-----------

// Redirect and log
func Redirect(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    url := vars["url"]
    host := RemoteHost{ip: r.RemoteAddr, destUrl: url, ts: time.Now().UTC()}
    Hosts.Value = host
    fmt.Println(host)
    Hosts = Hosts.Next()
    http.Redirect(w, r,  REDIRECT_PROTO + url, 301)
}

// Log collected information
func Log(w http.ResponseWriter, r *http.Request) {
    Hosts.Do(func(p interface{}) {
        if p.(RemoteHost).ip != "" {
            fmt.Fprintln(w, p.(RemoteHost))
        }
    })
}

// Initialize
func init() {
    Hosts = ring.New(RING_BUF_SZ)
    for i := 0; i < Hosts.Len(); i++ {
        Hosts.Value = RemoteHost{ip: ""}
        Hosts = Hosts.Next()
    }
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
