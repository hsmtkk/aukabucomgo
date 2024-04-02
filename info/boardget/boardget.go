package boardget

import (
	"fmt"

	"github.com/hsmtkk/aukabucomgo/base"
)

type Client interface {
	BoardGet(symbol string) ([]byte, error)
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
