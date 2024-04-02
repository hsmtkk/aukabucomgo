package boardget

import (
	"encoding/json"
	"fmt"

	"github.com/hsmtkk/aukabucomgo/base"
)

type Client interface {
	BoardGet(symbol string, marketCode MarketCode) (ResponseSchema, error)
}

func New(baseClient base.Client) Client {
	return &clientImpl{baseClient: baseClient}
}

type clientImpl struct {
	baseClient base.Client
}

/*
1	東証
3	名証
5	福証
6	札証
2	日通し
23	日中
24	夜間
*/

type MarketCode int

const (
	TOUSYOU MarketCode = iota
	ALL_DAY
	DAY
	NIGHT
)

func (clt *clientImpl) BoardGet(symbol string, marketCode MarketCode) (ResponseSchema, error) {
	marketCodeMap := map[MarketCode]string{
		TOUSYOU: "1",
		ALL_DAY: "2",
		DAY:     "23",
		NIGHT:   "24",
	}
	symbolStr := fmt.Sprintf("%s@%s", symbol, marketCodeMap[marketCode])
	path := fmt.Sprintf("/board/%s", symbolStr)
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
