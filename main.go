package main

import (
	"errors"
	"net/http"
	"strconv"

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

// currency exchange rate table
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
	// Create currency exchange rate table
	TWDch = NewTargetItems(1, 3.669, 0.03281)
	JPYch = NewTargetItems(0.26956, 1, 0.00885)
	USDch = NewTargetItems(30.444, 111.801, 1)

	// Create dependency injection
	TWD = NewSourceItem("TWD", TWDch)
	JPY = NewSourceItem("JPY", JPYch)
	USD = NewSourceItem("USD", USDch)

	// Packaging
	Currencies = []*SourceItem{TWD, JPY, USD}
)

// 1. Display currency exchange rate table on home page
func getCurrencies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Currencies)
}

// 2. Present currency conversion results
func convertCurrency(c *gin.Context) {
	source := c.Query("source")
	target := c.Query("target")
	amountStr := c.Query("amount")

	// Validate input parameters
	if source == "" || target == "" || amountStr == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing required parameters: source, target, or amount"})
		return
	}

	//
	var amount float64
	var err error
	if amount, err = strconv.ParseFloat(amountStr, 64); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid amount format"})
		return
	}

	// Find source item
	sourceItem, err := getItemBySource(source)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Source currency not found"})
		return
	}

	// Find conversion rate
	conversionRate, found := sourceItem.getConversionRateByTarget(target)
	if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Target currency not found for source"})
		return
	}

	// Perform conversion calculation
	convertedAmount := amount * conversionRate
	ca := convertedAmount

	// Respond with conversion result
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "success",
		"amount":  ca,
	})
}

func (SI SourceItem) getConversionRateByTarget(target string) (float64, bool) {
	var conversionRate float64
	found := false
	for _, targetItem := range SI.Targets {
		if targetItem.Target == target {
			conversionRate = targetItem.FX
			found = true
			break
		}
	}

	return conversionRate, found
}

func getItemBySource(source string) (*SourceItem, error) {
	for _, item := range Currencies {
		if item.Source == source {
			return item, nil
		}
	}
	return nil, errors.New("source currency not found")
}

// main
func main() {
	r := gin.Default()
	r.GET("/", getCurrencies)
	r.GET("/convert", convertCurrency)
	r.Run()
}
