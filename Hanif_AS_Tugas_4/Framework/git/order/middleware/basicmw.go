package middleware

import (
	"time"

	"context"

	cm "Hanif_AS_Tugas_4/Framework/git/order/common"
	"Hanif_AS_Tugas_4/Framework/git/order/services"

	log "github.com/Sirupsen/logrus"
)

func BasicMiddleware() services.ServiceMiddleware {
	return func(next services.PaymentServices) services.PaymentServices {
		return BasicMiddlewareStruct{next}
	}
}

type BasicMiddlewareStruct struct {
	services.PaymentServices
}

func (mw BasicMiddlewareStruct) OrderHandler(ctx context.Context, request cm.Message) cm.Message {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("OrderHandler ends")
	}(time.Now())

	log.WithField("request", request).Info("OrderHandler begins")

	return mw.PaymentServices.OrderHandler(ctx, request)

}
