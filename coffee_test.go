package main

import (
  "fmt"
  "testing"
)

func TestOrderAChocolate(t *testing.T) {
  equals(PadHasBeenPressed("Chocolate",0, 100),"H::", t)
}

func TestOrderACoffee(t *testing.T) {
  equals(PadHasBeenPressed("Coffee",0, 100),"C::", t)
}

func TestOrderATea(t *testing.T) {
  equals(PadHasBeenPressed("Tea", 0, 100),"T::", t)
}

func TestOrderOneSugar(t *testing.T) {
  equals(PadHasBeenPressed("Chocolate",1, 100),"H:1:0", t)
}

func TestOrderTwoSugar(t *testing.T) {
  equals(PadHasBeenPressed("Chocolate",2, 100),"H:2:0", t)
}

func TestOrderThreeSugar(t *testing.T) {
  equals(PadHasBeenPressed("Chocolate",3, 100),"H:2:0", t)
}

func TestDontAllowOrderIfMoneyIfNotEnough(t *testing.T) {
  equals(PadHasBeenPressed("Chocolate",0, 50),"H::", t)
  equals(PadHasBeenPressed("Chocolate",0, 40),"M: Missing 10c", t)
}

func TestDontAllowOrderIfMoneyIfNotEnoughForTea(t *testing.T) {
  equals(PadHasBeenPressed("Tea",0, 40),"T::", t)
  equals(PadHasBeenPressed("Tea",0, 30),"M: Missing 10c", t)
}

func TestDontAllowOrderIfMoneyIfNotEnoughForCoffee(t *testing.T) {
  equals(PadHasBeenPressed("Coffee",0, 60),"C::", t)
  equals(PadHasBeenPressed("Coffee",0, 50),"M: Missing 10c", t)
}

func equals(value string, expected string, t *testing.T) {
  if value != expected {
    t.Error(fmt.Sprintf("expected [%s] is not equal to [%s]", expected, value))
  }
}
