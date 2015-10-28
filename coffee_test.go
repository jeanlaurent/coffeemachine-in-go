package main

import (
  "fmt"
  "testing"
)

func TestOrderAChocolate(t *testing.T) {
  equals(PadHasBeenPressed("Chocolate",0, 100, false),"H::", t)
}

func TestOrderACoffee(t *testing.T) {
  equals(PadHasBeenPressed("Coffee",0, 100, false),"C::", t)
}

func TestOrderATea(t *testing.T) {
  equals(PadHasBeenPressed("Tea", 0, 100, false),"T::", t)
}

func TestOrderOneSugar(t *testing.T) {
  equals(PadHasBeenPressed("Chocolate",1, 100, false),"H:1:0", t)
}

func TestOrderTwoSugar(t *testing.T) {
  equals(PadHasBeenPressed("Chocolate",2, 100, false),"H:2:0", t)
}

func TestOrderThreeSugar(t *testing.T) {
  equals(PadHasBeenPressed("Chocolate",3, 100, false),"H:2:0", t)
}

func TestDontAllowOrderIfMoneyIfNotEnough(t *testing.T) {
  equals(PadHasBeenPressed("Chocolate",0, 50, false),"H::", t)
  equals(PadHasBeenPressed("Chocolate",0, 40, false),"M: Missing 10c", t)
}

func TestDontAllowOrderIfMoneyIfNotEnoughForTea(t *testing.T) {
  equals(PadHasBeenPressed("Tea",0, 40, false),"T::", t)
  equals(PadHasBeenPressed("Tea",0, 30, false),"M: Missing 10c", t)
}

func TestDontAllowOrderIfMoneyIfNotEnoughForCoffee(t *testing.T) {
  equals(PadHasBeenPressed("Coffee",0, 60, false),"C::", t)
  equals(PadHasBeenPressed("Coffee",0, 50, false),"M: Missing 10c", t)
}

func TestDontAllowOrderIfMoneyIfNotEnoughForOrange(t *testing.T) {
  equals(PadHasBeenPressed("Orange",0, 60, false),"O::", t)
  equals(PadHasBeenPressed("Orange",0, 50, false),"M: Missing 10c", t)
}

func TestOrderAnExtraHotChocolate(t *testing.T) {
  equals(PadHasBeenPressed("Chocolate",0, 100, true),"Hh::", t)
}

func TestOrderAnExtraHotCoffee(t *testing.T) {
  equals(PadHasBeenPressed("Coffee",0, 100, true),"Ch::", t)
}

func TestOrderAnExtraHotTea(t *testing.T) {
  equals(PadHasBeenPressed("Tea", 0, 100, true),"Th::", t)
}

func TestOrderAnExtraHotOrange(t *testing.T) {
  equals(PadHasBeenPressed("Orange", 0, 100, true),"O::", t)
}

func equals(value string, expected string, t *testing.T) {
  if value != expected {
    t.Error(fmt.Sprintf("expected [%s] is not equal to [%s]", expected, value))
  }
}
