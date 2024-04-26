package models

type MenuItem interface {
	GetName() string
	GetFlavor() string
	GetPrice() uint64
	GetStock() uint16
	PrintDetail()
	AddStock(amount int)
	ReduceStock(amount int)
	Cook()
}
