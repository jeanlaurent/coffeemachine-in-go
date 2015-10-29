package main

type Beverages struct {
  Sum int
  List map[string]*Beverage
  machine External
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
  b.machine = Machine{}
}
