package main

import (
  "fmt"
  "testing"
)

func TestOrderAChocolate(t *testing.T) {
  equals(PadHasBeenPressed("Chocolate",0),"H::", t)
}

func TestOrderACoffee(t *testing.T) {
  equals(PadHasBeenPressed("Coffee",0),"C::", t)
}

func TestOrderATea(t *testing.T) {
  equals(PadHasBeenPressed("Tea", 0),"T::", t)
}

func TestOrderOneSugar(t *testing.T) {
  equals(PadHasBeenPressed("Chocolate",1),"H:1:0", t)
}

func TestOrderTwoSugar(t *testing.T) {
  equals(PadHasBeenPressed("Chocolate",2),"H:2:0", t)
}

func TestOrderThreeSugar(t *testing.T) {
  equals(PadHasBeenPressed("Chocolate",3),"H:2:0", t)
}

func equals(value string, expected string, t *testing.T) {
  if value != expected {
    t.Error(fmt.Sprintf("expected [%s] is not equal to [%s]", expected, value))
  }

}
