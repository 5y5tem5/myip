package main

import (
        "fmt"
        "net/http"
        "errors"
        "os"
        "strings"
        "encoding/json"
        "crypto/tls"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
        type IP struct {
        IP string `json:"ip"`
        }
        ip := IP{strings.Split(r.RemoteAddr, ":")[0]}
        fmt.Printf("request from:")
        fmt.Fprintf(os.Stdout,strings.Split(r.RemoteAddr, ":")[0]+"\n")
        json.NewEncoder(w).Encode(ip)
}

func main() {
        http.HandleFunc("/", getRoot)
        //err := http.ListenAndServe(":80", nil)
        server := &http.Server{
        Addr: ":443",
        TLSConfig: &tls.Config{
                MinVersion: tls.VersionTLS12,
                MaxVersion: tls.VersionTLS13,
                },
        }
        err := server.ListenAndServeTLS( "fullchain.pem", "privkey.pem" )
        if errors.Is(err, http.ErrServerClosed) {
                fmt.Printf("server closed\n")
        } else if err != nil {
                fmt.Printf("error starting server: %s\n", err)
                os.Exit(1)
        }
}
