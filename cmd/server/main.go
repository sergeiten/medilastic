package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sergeiten/medilastic"
	"github.com/sergeiten/medilastic/config"
	"github.com/sergeiten/medilastic/search"

	log "github.com/sirupsen/logrus"
)

var host string
var port string
var configFile string

func init() {
	flag.StringVar(&host, "host", "0.0.0.0", "Host")
	flag.StringVar(&port, "port", "8888", "Port")
	flag.StringVar(&configFile, "config", "config.json", "The file name of config file")
	flag.Parse()
}

func main() {
	f, err := os.OpenFile("/logs/medilastic.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed to opening log file: %v", err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.WithError(err).Error("failed to close file")
		}
	}()

	log.SetOutput(f)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/search", Search)

	fmt.Printf("Starting up on %s:%s\n", host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, router))
}

func Search(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")
	if query == "" {
		serveError(w, "Query is empty")
		return
	}
	indexNames := r.FormValue("index_names")
	if indexNames == "" {
		serveError(w, "Indexes are empty")
		return
	}

	var err error

	from := 0
	size := 10

	if r.FormValue("from") != "" {
		from, err = strconv.Atoi(r.FormValue("from"))
		if err != nil {
			log.WithError(err).Fatal("failed to parse from value")
		}
	}

	if r.FormValue("size") != "" {
		size, err = strconv.Atoi(r.FormValue("size"))
		if err != nil {
			log.WithError(err).Fatal("failed to parse size value")
		}
	}

	ctx := context.Background()

	config, err := config.New(configFile)
	if err != nil {
		log.WithError(err).Fatal("failed to get config")
	}

	url := fmt.Sprintf("http://%s:%s", config.Elasticsearch.Host, config.Elasticsearch.Port)

	client, err := medilastic.NewClient(ctx, url)
	if err != nil {
		log.WithError(err).Fatal("failed to get elastic client")
	}

	indexes := strings.Split(indexNames, ",")

	result := make(map[string][]map[string]string)

	for _, index := range indexes {
		search := search.NewSearch(index, ctx, client)
		result[index], err = search.Search(query, from, size)
		if err != nil {
			log.WithError(err).Error("failed to get search result")
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.WithError(err).Fatal("failed to encode result")
	}
}

func serveError(w http.ResponseWriter, message string) {
	http.Error(w, message, 400)
}
