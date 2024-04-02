package positionsget

import (
	"encoding/json"
	"fmt"

	"github.com/hsmtkk/aukabucomgo/base"
)

type Client interface {
	PositionsGet(product Product, side Side) (ResponseSchema, error)
}

func New(baseClient base.Client) Client {
	return &clientImpl{baseClient: baseClient}
}

type clientImpl struct {
	baseClient base.Client
}

type Product int

const (
	ALL    Product = iota // 0
	FUTURE                // 3
	OPTION                // 4
)

type Side int

const (
	BUY  Side = iota // 1
	SELL             // 2
)

type ResponseSchema []Position

type Position struct {
	Symbol     string  `json:"Symbol"`
	SymbolName string  `json:"SymbolName"`
	LeavesQty  float64 `json:"LeavesQty"`
}

func (clt *clientImpl) PositionsGet(product Product, side Side) (ResponseSchema, error) {
	sideStr := "undef"
	if side == BUY {
		sideStr = "1"
	} else if side == SELL {
		sideStr = "2"
	}
	productStr := "undef"
	switch product {
	case ALL:
		productStr = "0"
	case FUTURE:
		productStr = "3"
	case OPTION:
		productStr = "4"
	}
	respBytes, err := clt.baseClient.Get("/positions", map[string]string{"product": productStr, "side": sideStr})
	if err != nil {
		return ResponseSchema{}, err
	}
	decoded := ResponseSchema{}
	if err := json.Unmarshal(respBytes, &decoded); err != nil {
		return ResponseSchema{}, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return decoded, nil
}
