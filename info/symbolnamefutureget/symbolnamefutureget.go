package symbolnamefutureget

import (
	"strconv"

	"github.com/hsmtkk/aukabucomgo/base"
)

type Client interface {
	SymbolNameFutureGet(futureCode FutureCode, derivMonth int) ([]byte, error)
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

func (clt *clientImpl) SymbolNameFutureGet(futureCode FutureCode, derivMonth int) ([]byte, error) {
	futureCodeMap := map[FutureCode]string{
		NK225micro: "NK225micro",
	}
	futureCodeStr := futureCodeMap[futureCode]
	return clt.baseClient.Get("/symbolname/future", map[string]string{"FutureCode": futureCodeStr, "DerivMonth": strconv.Itoa(derivMonth)})
}
