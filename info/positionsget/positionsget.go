package positionsget

import (
	"strconv"

	"github.com/hsmtkk/aukabucomgo/base"
)

type Client interface {
	PositionsGet(product int, side Side) ([]byte, error)
}

func New(baseClient base.Client) Client {
	return &clientImpl{baseClient: baseClient}
}

type clientImpl struct {
	baseClient base.Client
}

type Side int

const (
	BUY Side = iota
	SELL
)

func (clt *clientImpl) PositionsGet(product int, side Side) ([]byte, error) {
	sideStr := "undef"
	if side == BUY {
		sideStr = "1"
	} else if side == SELL {
		sideStr = "2"
	}
	return clt.baseClient.Get("/positions", map[string]string{"product": strconv.Itoa(product), "side": sideStr})
}
