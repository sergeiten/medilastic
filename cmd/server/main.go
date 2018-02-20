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
var cfg config.Config

func init() {
	flag.StringVar(&host, "host", "0.0.0.0", "Host")
	flag.StringVar(&port, "port", "8888", "Port")
	flag.StringVar(&configFile, "config", "config.json", "The file name of config file")
	flag.Parse()

	var err error
	cfg, err = config.New(configFile)
	if err != nil {
		log.WithError(err).Fatal("failed to load config file")
	}
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
	router.HandleFunc("/search", searchRequest)

	fmt.Printf("Starting up on %s:%s\n", host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, router))
}

func searchRequest(w http.ResponseWriter, r *http.Request) {
	err := validateSearchRequest(r)
	if err != nil {
		serveError(w, err.Error())
		return
	}

	from, err := validateNumber(r, "from", 0)
	if err != nil {
		log.WithError(err).Printf("failed to validate number")
	}
	size, err := validateNumber(r, "size", 10)
	if err != nil {
		log.WithError(err).Printf("failed to validate number")
	}

	ctx := context.Background()

	url := fmt.Sprintf("http://%s:%s", cfg.Elasticsearch.Host, cfg.Elasticsearch.Port)

	client, err := medilastic.NewClient(ctx, url)
	if err != nil {
		log.WithError(err).Fatal("failed to get elastic client")
	}

	indexes := strings.Split(r.FormValue("index_names"), ",")

	result := make(map[string][]map[string]string)

	for _, index := range indexes {
		search := search.NewSearch(ctx, index, client)
		result[index], err = search.Search(r.FormValue("query"), from, size)
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

func validateSearchRequest(r *http.Request) error {
	if r.FormValue("query") == "" {
		return fmt.Errorf("Query is empty")
	}
	if r.FormValue("index_names") == "" {
		return fmt.Errorf("Indexes are empty")
	}
	return nil
}

func validateNumber(r *http.Request, key string, defaultValue int) (int, error) {
	if r.FormValue(key) == "" {
		return defaultValue, nil
	}

	value, err := strconv.Atoi(r.FormValue(key))
	if err != nil {
		return defaultValue, fmt.Errorf("failed to parse %s to string", key)
	}
	return value, nil
}

func serveError(w http.ResponseWriter, message string) {
	http.Error(w, message, 400)
}
