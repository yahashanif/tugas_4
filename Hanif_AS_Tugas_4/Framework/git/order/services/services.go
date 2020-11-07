package services

import (
	"context"
	"fmt"
	"time"

	cm "Hanif_AS_Tugas_4/Framework/git/order/common"
)

//SubscriberServices is service definition
type PaymentServices interface {
	OrderHandler(context.Context, cm.Message) cm.Message
	CustomerHandler(context.Context, cm.Customers) cm.Customers
	ProductHandler(context.Context, cm.Products) cm.Products
}

type PaymentService struct{}

type ServiceMiddleware func(PaymentServices) PaymentServices

func utc() string {
	return time.Now().Format("2006-01-02 15:04:05 +0700")
}

func panicRecovery() {
	if r := recover(); r != nil {
		fmt.Printf("Recovering from panic: %v \n", r)
	}
}
