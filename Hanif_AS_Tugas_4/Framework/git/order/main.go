package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	cm "Hanif_AS_Tugas_4/Framework/git/order/common"
	"Hanif_AS_Tugas_4/Framework/git/order/middleware"
	"Hanif_AS_Tugas_4/Framework/git/order/services"
	"Hanif_AS_Tugas_4/Framework/git/order/transport"

	log "github.com/Sirupsen/logrus"
	httptransport "github.com/go-kit/kit/transport/http"
)

func initHandlers() {
	var svc services.PaymentServices

	svc = services.PaymentService{}
	svc = middleware.BasicMiddleware()(svc)

	root := cm.Config.RootURL
	fmt.Println(cm.Config.RootURL)
	http.Handle(fmt.Sprintf("%s/orders", root), httptransport.NewServer(
		transport.OrderEndpoint(svc), transport.DecodeRequest, transport.EncodeResponse,
	))

	http.Handle(fmt.Sprintf("%s/customers", root), httptransport.NewServer(
		transport.CustomerEndpoint(svc), transport.DecodeCustomerRequest, transport.EncodeResponse,
	))

	http.Handle(fmt.Sprintf("%s/product", root), httptransport.NewServer(
		transport.ProductEndpoint(svc), transport.DecodeProductRequest, transport.EncodeResponse,
	))

}

var logger *log.Entry

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
	})

	//log.SetReportCaller(true)
}

func main() {
	configFile := flag.String("conf", "conf-dev.yml", "main configuration file")
	flag.Parse()
	initLogger()
	log.WithField("file", *configFile).Info("Loading configuration file")
	cm.LoadConfigFromFile(configFile)
	initHandlers()

	var err error
	if cm.Config.RootURL != "" || cm.Config.ListenPort != "" {
		err = http.ListenAndServe(cm.Config.ListenPort, nil)
	}

	if err != nil {
		log.WithField("error", err).Error("Unable to start the server")
		os.Exit(1)
	}
}
