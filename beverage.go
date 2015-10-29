package main

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
