package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"

	_ "expvar"
	_ "net/http/pprof"

	"github.com/Giantmen/bitsoon/config"
	"github.com/Giantmen/bitsoon/log"
	"github.com/Giantmen/bitsoon/service"
)

var (
	cfgpath = flag.String("config", "config.toml", "config file path")
	debug   = flag.Bool("d", false, "log to stderr")
)

type Test struct {
}

func (t *Test) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func initLog(cfg *config.Config) {
	log.SetLevelByString(cfg.LogLevel)
	if !cfg.Debug {
		log.SetHighlighting(false)
		err := log.SetOutputByName(cfg.LogPath)
		if err != nil {
			log.Fatal(err)
		}
		log.SetRotateByDay()
	}
}

func setloglevel(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	level := r.FormValue("level")
	if len(level) == 0 {
		return
	}
	log.SetLevelByString(level)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()
	cfg, err := config.Parse(*cfgpath)
	if err != nil {
		log.Fatal(err)
	}

	// init log
	initLog(cfg)

	gm, err := service.NewGoodsManager(cfg)
	if err != nil {
		log.Fatal(err)
	}
	um ,err := service.NewUserManager(cfg)
	if err != nil {
		log.Fatal(err)
	}

	mux := mux.NewRouter()

	mux.HandleFunc("/goods/queryall", gm.QueryAllHandler)
	mux.HandleFunc("/goods/query/{goodsID}", gm.QueryOneHandler)
	mux.HandleFunc("/goods/insert", gm.InsertHandler)
	mux.HandleFunc("/goods/update", gm.UpdateHandler)
	mux.HandleFunc("/goods/delete", gm.DeleteHandler)

	mux.HandleFunc("/user/queryall", um.QueryAllHandler)
	mux.HandleFunc("/user/query/{userID}", um.QueryOneHandler)
	mux.HandleFunc("/user/insert", um.InsertHandler)
	mux.HandleFunc("/user/update", um.UpdateHandler)
	mux.HandleFunc("/user/delete", um.DeleteHandler)

	mux.HandleFunc("/loglevel", setloglevel)

	http.Handle("/", mux)

	log.Infof("start http listen on %s", cfg.Listen)
	log.Fatal(http.ListenAndServe(cfg.Listen, nil))

}
