package main

import (
  "fmt"
  "strings"
  "strconv"
  "sort"
)

type Beverages struct {
  Sum int
  List map[string]*Beverage
}

func (b *Beverages) Record(beverage *Beverage) {
  beverage.Record()
  b.Sum += beverage.Price
}

func (b *Beverages) Get(name string) *Beverage{
    return b.List[name]
}

func (b *Beverages) Init() {
  b.Sum = 0
  b.List = make(map[string]*Beverage)
  b.List["Chocolate"] = &Beverage{"H",50, true, 0, "Chocolate"}
  b.List["Coffee"] = &Beverage{"C",60, true, 0, "Coffee"}
  b.List["Tea"] = &Beverage{"T",40, true, 0, "Tea"}
  b.List["Orange"] = &Beverage{"O",60, false, 0, "Orange"}
}

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

var beverages = new(Beverages)

func PadHasBeenPressed(beverageString string, numberOfSugar int, money int, extraHotRequest bool) string {
  beverage := beverages.Get(beverageString)

  difference := beverage.Price - money
  if (difference > 0) {
    return fmt.Sprintf("M: Missing %dc", difference)
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
