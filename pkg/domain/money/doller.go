package money

type Doller struct {
	Amount int `json:"doller"`
}

func NewDoller(amount int) *Doller {
	return &Doller{amount}
}

func (doller *Doller) times(multiplier int) *Doller {
	return NewDoller(doller.Amount * multiplier)
}
