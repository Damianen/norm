package model

type AisleStocker struct {
	Id         int
	StockerId  int
	AisleId    int
	ShiftId    int
	Norm       float32
	NormalTime float32
	TimeWorked float32
}
