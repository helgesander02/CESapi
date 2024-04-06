package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SourceItem struct {
	Source  string       `json:"source"`
	Targets []TargetItem `json:"targets"`
}

type TargetItem struct {
	Target string  `json:"target"`
	FX     float64 `json:"fx"`
}

var (
	//
	TWDch = []TargetItem{
		{Target: "TWD", FX: 1},
		{Target: "JPY", FX: 3.669},
		{Target: "USD", FX: 0.03281},
	}
	JPYch = []TargetItem{
		{Target: "TWD", FX: 0.26956},
		{Target: "JPY", FX: 1},
		{Target: "USD", FX: 0.00885},
	}
	USDch = []TargetItem{
		{Target: "TWD", FX: 30.444},
		{Target: "JPY", FX: 111.801},
		{Target: "USD", FX: 1},
	}
	//
	TWD = SourceItem{Source: "TWD", Targets: TWDch}
	JPY = SourceItem{Source: "JPY", Targets: JPYch}
	USD = SourceItem{Source: "USD", Targets: USDch}
	//
	Currencies = []SourceItem{TWD, JPY, USD}
)

func getCurrencies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Currencies)
}

func getSource(c *gin.Context) {
	source := c.Query("source")
	sourceitem, err := getItemBySource(source)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found"})
	}

	c.IndentedJSON(http.StatusOK, sourceitem)
}

func getItemBySource(source string) (*SourceItem, error) {
	for _, item := range Currencies {
		if item.Source == source {
			return &item, nil
		}
	}

	return nil, errors.New("todo not found")
}

func main() {
	r := gin.Default()

	r.GET("/", getCurrencies)
	r.GET("/source", getSource)

	r.Run()
}
