package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"gopkg.in/ini.v1"
)

func hello(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	inputs, ok := r.URL.Query()["input"]
    
    if !ok || len(inputs[0]) < 1 {
        return
    }
	input := inputs[0]
	elapsed := time.Since(start).Seconds()
	if (input == "1000") {
		log.Println("Warning: Unicorn Refund Request: c4a5010e-6734-41e8-974b", elapsed)
	} else {
		fmt.Fprintf(w, "/lFaYJBBiveYZ/G4bCimw/Fa7ipI0tDl9f1uuXsKyw==")		
		log.Println("Got result in ", elapsed)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: Create a nice looking homepage.")
}

func main() {

	cfg, err := ini.Load("main.ini")
    if err != nil {
        fmt.Printf("Fail to read file: %v", err)
        os.Exit(1)
	}
	logPath := cfg.Section("").Key("log").String()
	
	f, err := os.OpenFile(logPath, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	http.HandleFunc("/calc", hello)
	http.HandleFunc("/", home)
 
    fmt.Printf("Done!\n")
    if err := http.ListenAndServe(":80", nil); err != nil {
        log.Fatal(err)
    }
}
