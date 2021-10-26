package models

type BlockTotals struct {
	Id           int     `json:"-"`
	BlockNumber  int     `json:"-"`
	Transactions int     `json:"transactions"`
	Amount       float64 `json:"amount"`
}
