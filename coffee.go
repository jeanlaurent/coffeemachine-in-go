package main

import (
  "fmt"
  "strconv"
)

type Order struct {
  Code string
  Price int
}

func getBeverages() (map[string]Order) {
  var beverages map[string]Order
  beverages = make(map[string]Order)
  beverages["Chocolate"] = Order{"H",50}
  beverages["Coffee"] = Order{"C",60}
  beverages["Tea"] = Order{"T",40}
  return beverages
}

func PadHasBeenPressed(beverageString string, numberOfSugar int, money int) string {
  beverage := getBeverages()[beverageString]
  difference := beverage.Price - money
  if (difference > 0) {
    return fmt.Sprintf("M: Missing %dc", difference)
  }
  touillette := ""
  sugarCode := ""
  if (numberOfSugar > 0) {
    if (numberOfSugar > 2) {
      numberOfSugar = 2
    }
    touillette = "0"
    sugarCode =  strconv.Itoa(numberOfSugar)
  }
  return fmt.Sprintf("%s:%s:%s", beverage.Code, sugarCode, touillette)
}
