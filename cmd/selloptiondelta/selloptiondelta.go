package selloptiondelta

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/hsmtkk/aukabucomgo/base"
	"github.com/hsmtkk/aukabucomgo/info"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

const PRODUCT_OP = 4
const SIDE_SELL = 1
const MARKET_CODE_ALL_DAY = 2

var SellOptionDeltaCmd = &cobra.Command{
	Use: "selloptiondelta",
	Run: sellOptionDelta,
}

type positionSchema struct {
	Symbol     string  `json:"Symbol"`
	SymbolName string  `json:"SymbolName"`
	LeavesQty  float64 `json:"LeavesQty"`
}

type boardSchema struct {
	Delta float64 `json:"Delta"`
}

func sellOptionDelta(cmd *cobra.Command, args []string) {
	positions, err := getOptionPositions()
	totalDelta := 0.0
	if err != nil {
		log.Fatal(err)
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Quarter", "Delta"})
	for _, pos := range positions {
		t.AppendRow([]interface{}{pos.symbolName, pos.leavesQty, pos.delta})
		totalDelta += float64(pos.leavesQty) * pos.delta
	}
	t.Render()
	fmt.Printf("Total delta: %f\n", totalDelta)
}

type optionPosition struct {
	symbolName string
	leavesQty  int
	delta      float64
}

func getOptionPositions() ([]optionPosition, error) {
	baseClient, err := base.New()
	if err != nil {
		log.Fatal(err)
	}
	infoClient := info.New(baseClient)
	resp, err := infoClient.PositionsGet(PRODUCT_OP, SIDE_SELL)
	if err != nil {
		log.Fatal(err)
	}
	positions := []positionSchema{}
	if err := json.Unmarshal(resp, &positions); err != nil {
		log.Fatal(err)
	}
	result := []optionPosition{}
	for _, p := range positions {
		symbol := fmt.Sprintf("%s@%d", p.Symbol, MARKET_CODE_ALL_DAY)
		resp, err := infoClient.BoardGet(symbol)
		if err != nil {
			log.Fatal(err)
		}
		board := boardSchema{}
		if err := json.Unmarshal(resp, &board); err != nil {
			log.Fatal(err)
		}
		result = append(result, optionPosition{symbolName: p.SymbolName, leavesQty: int(p.LeavesQty), delta: board.Delta})
	}
	return result, nil
}
