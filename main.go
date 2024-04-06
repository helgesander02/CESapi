package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SourceItem struct {
	Source  string        `json:"source"`
	Targets []*TargetItem `json:"targets"`
}

type TargetItem struct {
	Target string  `json:"target"`
	FX     float64 `json:"fx"`
}

func NewTargetItems(twd float64, jpy float64, usd float64) []*TargetItem {
	return []*TargetItem{
		{Target: "TWD", FX: twd},
		{Target: "JPY", FX: jpy},
		{Target: "USD", FX: usd},
	}
}

// TargetItem inject SourceItem
// Reference https://tehub.com/a/c0W0jZ5qR8
func NewSourceItem(sourcename string, targetItem []*TargetItem) *SourceItem {
	return &SourceItem{Source: sourcename, Targets: targetItem}
}

var (
	// Currency exchange rates
	TWDch = NewTargetItems(1, 3.669, 0.03281)
	JPYch = NewTargetItems(0.26956, 1, 0.00885)
	USDch = NewTargetItems(30.444, 111.801, 1)

	//
	TWD = NewSourceItem("TWD", TWDch)
	JPY = NewSourceItem("JPY", JPYch)
	USD = NewSourceItem("USD", USDch)

	//
	Currencies = []*SourceItem{TWD, JPY, USD}
)

func getCurrencies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Currencies)
}

func main() {
	r := gin.Default()
	r.GET("/", getCurrencies)
	r.Run()
}
