package infrastructure

import (
	log "book-lending-api/pkg/logger"
	elastic "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	elogrus "gopkg.in/sohlich/elogrus.v7"
)

func SetupElasticLogger() {
	logger := log.NewLogger()

	client, err := elastic.NewClient(elastic.SetURL("http://elasticsearch:9200"))
	if err != nil {
		logger.Warnf("Elasticsearch unavailable, skipping log hook: %v", err)
		return
	}

	hook, err := elogrus.NewElasticHook(client, "localhost", logrus.InfoLevel, "book-lending-logs")
	if err != nil {
		logger.Warnf("Failed to add ES log hook: %v", err)
		return
	}

	logger.AddHook(hook)
	logger.SetFormatter(&logrus.JSONFormatter{})
}
