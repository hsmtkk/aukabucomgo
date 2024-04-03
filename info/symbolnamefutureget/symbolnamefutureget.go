package symbolnamefutureget

import (
	"encoding/json"
	"time"

	"github.com/hsmtkk/aukabucomgo/base"
)

type Client interface {
	SymbolNameFutureGet(futureCode FutureCode, year, month int) (ResponseSchema, error)
}

func New(baseClient base.Client) Client {
	return &clientImpl{baseClient: baseClient}
}

type clientImpl struct {
	baseClient base.Client
}

type FutureCode int

const (
	NK225micro FutureCode = iota
)

type ResponseSchema struct {
	Symbol     string `json:"Symbol"`
	SymbolName string `json:"SymbolName"`
}

func (clt *clientImpl) SymbolNameFutureGet(futureCode FutureCode, year, month int) (ResponseSchema, error) {
	futureCodeMap := map[FutureCode]string{
		NK225micro: "NK225micro",
	}
	futureCodeStr := futureCodeMap[futureCode]
	yearMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	derivMonth := yearMonth.Format("200601")
	respBytes, err := clt.baseClient.Get("/symbolname/future", map[string]string{"FutureCode": futureCodeStr, "DerivMonth": derivMonth})
	if err != nil {
		return ResponseSchema{}, err
	}
	decoded := ResponseSchema{}
	if err := json.Unmarshal(respBytes, &decoded); err != nil {
		return ResponseSchema{}, err
	}
	return decoded, nil
}
