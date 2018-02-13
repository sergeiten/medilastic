package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/sergeiten/medilastic"
	"github.com/sergeiten/medilastic/config"
	"github.com/sergeiten/medilastic/index"
	"github.com/sergeiten/medilastic/repository"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var configFile string
var indexName string
var mappingFile string

func init() {
	flag.StringVar(&configFile, "config", "config.json", "The file name of config file")
	flag.StringVar(&indexName, "index", "", "The index name")
	flag.StringVar(&mappingFile, "mapping", "", "The mapping file")
	flag.Parse()
}

func main() {
	ctx := context.Background()

	client, err := medilastic.NewClient(ctx)
	if err != nil {
		log.WithError(err).Fatal("failed to get elastic client")
	}

	config, err := config.New(configFile)
	if err != nil {
		log.WithError(err).Fatal("failed to get config")
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name))
	if err != nil {
		log.WithError(err).Fatal("failed to open db connection")
	}
	defer db.Close()

	repository := repository.NewRepository(indexName, db)
	if err != nil {
		log.WithError(err).Fatal("failed to get repository")
	}

	mapping, err := ioutil.ReadFile(mappingFile)
	if err != nil {
		log.WithError(err).Fatal("failed to read mapping json file")
	}

	indexer := index.NewIndexer(ctx, client, repository, string(mapping), indexName, indexName)

	err = indexer.Do()
	if err != nil {
		log.WithError(err).Fatal("failed to create index")
	}
}
