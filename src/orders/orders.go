package orders

import (
	"errors"
	"strings"
	"time"

	"github.com/mitsukomegumi/Crypto-Go/src/accounts"
	"github.com/mitsukomegumi/Crypto-Go/src/common"
	"github.com/mitsukomegumi/Crypto-Go/src/pairs"
)

// Order - definition of order, fields attributed to a single order
type Order struct {
	Filled bool `json:"filled"`

	IssuanceTime time.Time `json:"issuancetime"`
	FillTime     time.Time `json:"filletime"`

	Amount float64 `json:"amount"`

	OrderType string     `json:"ordertype"`
	OrderFee  float64    `json:"orderfee"`
	OrderPair pairs.Pair `json:"tradingpair"`

	Issuer *accounts.Account `json:"issuer"`

	OrderID string `json:"orderid"`
}

// NewOrder - creates, retursn new instance of order struct
func NewOrder(account *accounts.Account, ordertype string, tradingpair pairs.Pair, amount float64) (Order, error) {
	ordertype = strings.ToUpper(ordertype)
	if amount < account.Balance {
		rOrder := Order{Filled: false, IssuanceTime: time.Now().UTC(), Amount: (1.0 - common.FeeRate) * amount, OrderType: ordertype, OrderPair: tradingpair, Issuer: account, OrderID: "", OrderFee: common.FeeRate * amount}

		hash, err := common.Hash(rOrder)

		if err != nil {
			return rOrder, err
		}

		rOrder.OrderID = hash

		return rOrder, nil
	}
	return Order{}, errors.New("insufficient balance")
}
