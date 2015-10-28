package main

import (
  "fmt"
  "strings"
  "strconv"
  "sort"
)

type Beverage struct {
  Code string
  Price int
  IsExtraHot bool
  Count int
  Name string
}

func (b *Beverage) Record() {
  b.Count++
}

var beverages = initBeverages()
var sum = 0

func initBeverages() (map[string]*Beverage) {
  beveragesMap := make(map[string]*Beverage)
  beveragesMap["Chocolate"] = &Beverage{"H",50, true, 0, "Chocolate"}
  beveragesMap["Coffee"] = &Beverage{"C",60, true, 0, "Coffee"}
  beveragesMap["Tea"] = &Beverage{"T",40, true, 0, "Tea"}
  beveragesMap["Orange"] = &Beverage{"O",60, false, 0, "Orange"}
  return beveragesMap
}

func Reset() {
  for _, beverage := range beverages {
    beverage.Count = 0
  }
  sum = 0
}

func PadHasBeenPressed(beverageString string, numberOfSugar int, money int, extraHotRequest bool) string {
  beverage := beverages[beverageString]

  difference := beverage.Price - money
  if (difference > 0) {
    return fmt.Sprintf("M: Missing %dc", difference)
  }

  beverage.Record()
  sum += beverage.Price

  extraHot := handleExtraHot(*beverage, extraHotRequest)
  sugarCode, touillette := handleSugarAndTouillette(*beverage, numberOfSugar)

  return fmt.Sprintf("%s%s:%s:%s", beverage.Code, extraHot, sugarCode, touillette)
}

func printReport() string {
  myStrings := make([]string,0)
  for _, beverage := range beverages {
    myStrings = append(myStrings, fmt.Sprintf("%s: %d", beverage.Name, beverage.Count))
  }
  sort.Strings(myStrings)
  myStrings = append(myStrings, fmt.Sprintf(" ,Money: %d", sum))
  return strings.Join(myStrings, ", ")
}

func handleSugarAndTouillette(beverage Beverage, numberOfSugar int) (string, string) {
  if (numberOfSugar > 0) {
    if (numberOfSugar > 2) {
      numberOfSugar = 2
    }
    return strconv.Itoa(numberOfSugar), "0"
  }
  return "",""
}

func handleExtraHot(beverage Beverage, extraHotRequest bool) string {
  if (extraHotRequest && beverage.IsExtraHot) {
    return "h"
  }
  return ""
}
