package main

import (
	"net/http"
	"time"

	"github.com/bitbrewers/lappy"
	"github.com/bitbrewers/tranx2"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type envConf struct {
	DSN        string `split_words:"true" required:"true"`
	ListenAddr string `split_words:"true" required:"true"`
	SerialPort string `split_words:"true" required:"true"`
	LogLevel   string `split_words:"true" required:"false" default:"info"`
}

func main() {
	log := logrus.New()
	conf := &envConf{}
	if err := envconfig.Process("", conf); err != nil {
		log.Fatal(err)
	}

	level, err := logrus.ParseLevel(conf.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(level)
	log.Info("Starting lappy")
	log.Debugf("config: %+v", conf)

	s, err := lappy.NewSqliteStorage(conf.DSN)
	if err != nil {
		log.Fatal(err)
	}

	p := lappy.NewSsePublisher()
	h := &lappy.Handler{
		Log:       log,
		Storage:   s,
		Publisher: p,
	}

	http.HandleFunc("/api/events", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Accel-Buffering", "no;")
		p.Server.HTTPHandler(w, r)
	})

	http.HandleFunc("/api/start", func(w http.ResponseWriter, r *http.Request) {
		log.Info("start race")
		if _, err := s.StartRace(time.Now()); err != nil {
			log.Errorf("could not start race: %s", err)
		}
	})
	http.HandleFunc("/api/stop", func(w http.ResponseWriter, r *http.Request) {
		log.Info("stop race")
		if err := s.StopRace(time.Now()); err != nil {
			log.Errorf("could not stop race: %s", err)
		}
	})

	go func() {
		http.ListenAndServe(conf.ListenAddr, nil)
	}()

	c := tranx2.NewClient(conf.SerialPort, h)
	if err := c.Listen(); err != nil {
		log.Fatal(err)
	}

	c.Serve()
}
