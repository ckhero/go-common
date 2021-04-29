package databases

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"net"
	"net/http"
	"time"
	config3 "youmi-micro-cluster/src/common/config"
	"youmi-micro-cluster/src/common/logger"
)

var ElasticClient *elastic.Client

func RegisterElasticSearchClient(cfg *config3.Elastic) {
	ctx := context.TODO()
	defer func() {
		if err := recover(); err != nil {
			logger.GetLogger(ctx).Errorf(`registerElasticSearchClient error: %v`, err)
		}
	}()

	if ElasticClient != nil {
		return
	}

	var (
		url      = cfg.Host
		username = cfg.Username
		password = cfg.Password
		index    = cfg.Index
	)

	var sniff = false //<4>
	//cfgES := &config2.Config{
	//	URL:      url,
	//	Username: username,
	//	Password: password,
	//}
	//
	//cfgES.Sniff = &sniff
	//var esClient, err = elastic.NewClientFromConfig(cfgES)

	// 自定义 设置
	httpClient := &http.Client{}
	httpClient.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:    10 * time.Second,
		MaxIdleConns:           cfg.MaxIdleConns,
		MaxIdleConnsPerHost:    cfg.MaxIdleConnsPerHost,
		MaxConnsPerHost:        cfg.MaxIdleConnsPerHost,
		IdleConnTimeout:        90 * time.Second,
		ExpectContinueTimeout:  10 * time.Second,
	}
	esClient, err := elastic.NewClient(
			elastic.SetHttpClient(httpClient),
			elastic.SetSniff(sniff),
			elastic.SetURL(url),
			elastic.SetBasicAuth(username, password),
		)
	fmt.Println(cfg.MaxIdleConnsPerHost)
	fmt.Println(cfg.MaxIdleConnsPerHost)
	if err != nil || esClient == nil {
		logger.GetLogger(ctx).Errorf(`elastic.NewClientFromConfig error: %v`, err)
		panic(err)
	}

	ElasticClient = esClient

	// 创建索引
	exists, err := ElasticClient.IndexExists(index).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		createIndex, err := ElasticClient.CreateIndex(index).Do(ctx)
		if err != nil {
			logger.GetLogger(ctx).Errorf(`ElasticClient.CreateIndex error: %v`, err)
		}

		logger.GetLogger(ctx).Infof(`createIndex: %v`, createIndex.Acknowledged)
	}
}

func GetESClient() *elastic.Client {
	return ElasticClient
}

