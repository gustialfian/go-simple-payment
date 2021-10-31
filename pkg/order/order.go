package order

import "errors"

type Order struct {
	Entries []OrderEntries `json:"entries"`
}

type OrderEntries struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	UserID string `json:"userID"`
}

func (o *Order) Proccess(orderRepository OrderProccessRepository) error {
	for _, entry := range o.Entries {
		amount, err := orderRepository.GetAmountEntries(entry)
		if err != nil {
			return err
		}
		newAmount := amount + entry.Amount
		if newAmount <= 0 {
			return errors.New("new amount minus")
		}
		orderRepository.Deduct(newAmount, entry)
	}
	return nil
}

func (o *Order) IsValid() bool {
	total := 0

	for i := 0; i < len(o.Entries); i++ {
		total += o.Entries[i].Amount
	}

	return total == 0
}
