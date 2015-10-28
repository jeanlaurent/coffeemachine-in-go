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

func getBeverages() (map[string]Beverage) {
  var beverages map[string]Beverage
  beverages = make(map[string]Beverage)
  beverages["Chocolate"] = Beverage{"H",50, true}
  beverages["Coffee"] = Beverage{"C",60, true}
  beverages["Tea"] = Beverage{"T",40, true}
  beverages["Orange"] = Beverage{"O",60, false}
  return beverages
}

func PadHasBeenPressed(beverageString string, numberOfSugar int, money int, extraHotRequest bool) string {
  beverage := getBeverages()[beverageString]

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
