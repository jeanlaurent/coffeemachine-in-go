package main

type External interface {
  NotifyMissingDrink(drink string)
  IsEmpty(drink string) bool
}

type Machine struct {}

func (m Machine) IsEmpty(drink string) bool { return false }

func (m Machine) NotifyMissingDrink(drink string) { /* noop */ }
