package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/vaxx99/swtch/conf"
	"github.com/vaxx99/swtch/es11"
	"github.com/vaxx99/swtch/fess"
	"github.com/vaxx99/swtch/si2k"
)

var cfg *conf.Config

func main() {
	conf.LoadConfig("conf.json")
	cfg = conf.GetConfig()
	fess.Cfg = cfg
	si2k.Cfg = cfg
	es11.Cfg = cfg
	os.Chdir(cfg.Path)
	go stat()
	for {
		time.Sleep(time.Second * 5)
	}
}

func stat() {
	http.HandleFunc("/", fess.Home)
	http.HandleFunc("/fess/stat", fess.Stat)
	http.HandleFunc("/fess/logs", fess.Logs)
	http.HandleFunc("/fess", fess.Home)
	http.HandleFunc("/fess/form", fess.Form)
	http.HandleFunc("/fess/call", fess.Call)
	//
	http.HandleFunc("/si2k/stat", si2k.Stat)
	http.HandleFunc("/si2k/logs", si2k.Logs)
	http.HandleFunc("/si2k", si2k.Home)
	http.HandleFunc("/si2k/form", si2k.Form)
	http.HandleFunc("/si2k/call", si2k.Call)
	//
	http.HandleFunc("/es11/stat", es11.Stat)
	http.HandleFunc("/es11/logs", es11.Logs)
	http.HandleFunc("/es11", es11.Home)
	http.HandleFunc("/es11/form", es11.Form)
	http.HandleFunc("/es11/call", es11.Call)

	http.Handle("/fess/", http.StripPrefix("/fess/", http.FileServer(http.Dir("fess"))))
	http.Handle("/si2k/", http.StripPrefix("/si2k/", http.FileServer(http.Dir("si2k"))))
	http.Handle("/es11/", http.StripPrefix("/es11/", http.FileServer(http.Dir("es11"))))
	log.Println("Start stat...")
	log.Println(fess.Cfg)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
