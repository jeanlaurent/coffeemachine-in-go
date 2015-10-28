package main

import (
  "fmt"
  "strconv"
)

type Order struct {
  Code string
  Price int
  IsExtraHot bool
}

func getBeverages() (map[string]Order) {
  var beverages map[string]Order
  beverages = make(map[string]Order)
  beverages["Chocolate"] = Order{"H",50, true}
  beverages["Coffee"] = Order{"C",60, true}
  beverages["Tea"] = Order{"T",40, true}
  beverages["Orange"] = Order{"O",60, false}
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

func handleSugarAndTouillette(beverage Order, numberOfSugar int) (string, string) {
  if (numberOfSugar > 0) {
    if (numberOfSugar > 2) {
      numberOfSugar = 2
    }
    return strconv.Itoa(numberOfSugar), "0"
  }
  return "",""
}

func handleExtraHot(beverage Order, extraHotRequest bool) string {
  if (extraHotRequest && beverage.IsExtraHot) {
    return "h"
  }
  return ""
}
