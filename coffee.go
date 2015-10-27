package main

import (
  "fmt"
  "strconv"
)

func PadHasBeenPressed(boisson string, numberOfSugar int) string {
  beverageCode := ""
  if "Chocolate" == boisson {
    beverageCode = "H"
  }
  if "Coffee" == boisson {
    beverageCode = "C"
  }
  if "Tea" == boisson {
    beverageCode = "T"
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
  return fmt.Sprintf("%s:%s:%s", beverageCode, sugarCode, touillette)
}
