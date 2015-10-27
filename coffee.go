package main

import (
  "fmt"
  "strconv"
)

func getBeverages() (map[string]string) {
  var beverages map[string]string
  beverages = make(map[string]string)
  beverages["Chocolate"] = "H"
  beverages["Coffee"] = "C"
  beverages["Tea"] = "T"
  return beverages
}

func PadHasBeenPressed(beverage string, numberOfSugar int) string {
  beverages := getBeverages()
  beverageCode := beverages[beverage]
  touillette := ""
  sugarCode := ""
  if (numberOfSugar > 0) {
    if (numberOfSugar > 2) {
      numberOfSugar = 2
    }
    touillette = "0"
    sugarCode =  strconv.Itoa(numberOfSugar)
  }
  return fmt.Sprintf("%s:%s:%s", beverageCode, sugarCode, touillette)
}
