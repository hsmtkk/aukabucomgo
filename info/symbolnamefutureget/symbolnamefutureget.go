package symbolnamefutureget

import (
	"time"

	"github.com/hsmtkk/aukabucomgo/base"
)

type Client interface {
	SymbolNameFutureGet(futureCode FutureCode, year, month int) ([]byte, error)
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

func (clt *clientImpl) SymbolNameFutureGet(futureCode FutureCode, year, month int) ([]byte, error) {
	futureCodeMap := map[FutureCode]string{
		NK225micro: "NK225micro",
	}
	futureCodeStr := futureCodeMap[futureCode]
	yearMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, nil)
	derivMonth := yearMonth.Format("200601")
	return clt.baseClient.Get("/symbolname/future", map[string]string{"FutureCode": futureCodeStr, "DerivMonth": derivMonth})
}
