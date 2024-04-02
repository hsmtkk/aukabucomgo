package symbolnameoptionget

import (
	"strconv"

	"github.com/hsmtkk/aukabucomgo/base"
)

type Client interface {
	SymbolNameOptionGet(optionCode OptionCode, derivMonth int, putOrCall PutOrCall, strikePrice int) ([]byte, error)
}

func New(baseClient base.Client) Client {
	return &clientImpl{baseClient: baseClient}
}

type clientImpl struct {
	baseClient base.Client
}

type OptionCode int

const (
	NK225op OptionCode = iota
	NK225miniop
)

type PutOrCall int

const (
	PUT PutOrCall = iota
	CALL
)

func (clt *clientImpl) SymbolNameOptionGet(optionCode OptionCode, derivMonth int, putOrCall PutOrCall, strikePrice int) ([]byte, error) {
	optionCodeMap := map[OptionCode]string{
		NK225op:     "NK225op",
		NK225miniop: "NK225miniop",
	}
	optionCodeStr := optionCodeMap[optionCode]
	putOrCallMap := map[PutOrCall]string{
		PUT:  "P",
		CALL: "C",
	}
	putOrCallStr := putOrCallMap[putOrCall]
	return clt.baseClient.Get("/symbolname/future", map[string]string{"OptionCode": optionCodeStr, "DerivMonth": strconv.Itoa(derivMonth), "PutOrCall": putOrCallStr, "StrikePrice": strconv.Itoa(strikePrice)})
}
