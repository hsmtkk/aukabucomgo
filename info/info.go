package info

import (
	"fmt"
	"strconv"

	"github.com/hsmtkk/aukabucomgo/base"
)

type Client interface {
	BoardGet(symbol string) ([]byte, error)
	PositionsGet(product int, side int) ([]byte, error)
}

func New(baseClient base.Client) Client {
	return &clientImpl{baseClient: baseClient}
}

type clientImpl struct {
	baseClient base.Client
}

func (clt *clientImpl) BoardGet(symbol string) ([]byte, error) {
	path := fmt.Sprintf("/board/%s", symbol)
	return clt.baseClient.Get(path, nil)
}

func (clt *clientImpl) PositionsGet(product int, side int) ([]byte, error) {
	return clt.baseClient.Get("/positions", map[string]string{"product": strconv.Itoa(product), "side": strconv.Itoa(side)})
}
