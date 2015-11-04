package main

import (
  "fmt"
  "strings"
  "strconv"
  "sort"
)

var beverages = new(Beverages)

func PadHasBeenPressed(beverageString string, numberOfSugar int, money int, extraHotRequest bool) string {
  beverage := beverages.Get(beverageString)

  difference := beverage.Price - money
  if difference > 0 {
    return fmt.Sprintf("M: Missing %dc", difference)
  }
  if beverages.machine.IsEmpty(beverageString) {
    beverages.machine.NotifyMissingDrink(beverageString)
    return fmt.Sprintf("M: we have a shortage of %s.", beverageString)
  }

  beverages.Record(beverage)

  extraHot := handleExtraHot(*beverage, extraHotRequest)
  sugarCode, touillette := handleSugarAndTouillette(*beverage, numberOfSugar)

  return fmt.Sprintf("%s%s:%s:%s", beverage.Code, extraHot, sugarCode, touillette)
}

func printReport() string {
  myStrings := make([]string,0)
  for _, beverage := range beverages.List {
    myStrings = append(myStrings, fmt.Sprintf("%s: %d", beverage.Name, beverage.Count))
  }
  sort.Strings(myStrings)
  myStrings = append(myStrings, fmt.Sprintf(" ,Money: %d", beverages.Sum))
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

func main() {
	fmt.Println("Nothing to see here, move along, run test instead")
}
