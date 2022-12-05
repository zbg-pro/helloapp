package models

import (
	"strconv"
	"time"
)

var (
	OrderList = make(map[string]*Order)
)

type Order struct {
	Id       string
	Amount   string
	Price    string
	marketId int64
}

func AddOrder(o Order) string {
	o.Id = "order_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	OrderList[o.Id] = &o
	return o.Id
}

func getAllOrders() map[string]*Order {
	return OrderList
}
