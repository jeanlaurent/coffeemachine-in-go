package main

import (
  "fmt"
  "strconv"
)

type Beverage struct {
  Code string
  Price int
  IsExtraHot bool
}

var beverages = initBeverages()

func initBeverages() (map[string]Beverage) {
  var beveragesMap map[string]Beverage
  beveragesMap = make(map[string]Beverage)
  beveragesMap["Chocolate"] = Beverage{"H",50, true}
  beveragesMap["Coffee"] = Beverage{"C",60, true}
  beveragesMap["Tea"] = Beverage{"T",40, true}
  beveragesMap["Orange"] = Beverage{"O",60, false}
  return beveragesMap
}

func PadHasBeenPressed(beverageString string, numberOfSugar int, money int, extraHotRequest bool) string {
  beverage := beverages[beverageString]

  difference := beverage.Price - money
  if (difference > 0) {
    return fmt.Sprintf("M: Missing %dc", difference)
  }

  extraHot := handleExtraHot(beverage, extraHotRequest)
  sugarCode, touillette := handleSugarAndTouillette(beverage, numberOfSugar)

  return fmt.Sprintf("%s%s:%s:%s", beverage.Code, extraHot, sugarCode, touillette)
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
