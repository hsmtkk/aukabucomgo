package boardget

import (
	"encoding/json"
	"fmt"

	"github.com/hsmtkk/aukabucomgo/base"
)

type Client interface {
	BoardGet(symbol string) (ResponseSchema, error)
}

func New(baseClient base.Client) Client {
	return &clientImpl{baseClient: baseClient}
}

type clientImpl struct {
	baseClient base.Client
}

func (clt *clientImpl) BoardGet(symbol string) (ResponseSchema, error) {
	path := fmt.Sprintf("/board/%s", symbol)
	respBytes, err := clt.baseClient.Get(path, nil)
	if err != nil {
		return ResponseSchema{}, err
	}
	decoded := ResponseSchema{}
	if err := json.Unmarshal(respBytes, &decoded); err != nil {
		return ResponseSchema{}, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return decoded, nil
}

type ResponseSchema struct {
	Delta float64 `json:"Delta"`
}
