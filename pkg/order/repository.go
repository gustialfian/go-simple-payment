package order

import (
	"database/sql"
	"log"
)

type OrderProccessRepository interface {
	GetAmountEntries(OrderEntries) (int, error)
	Deduct(int, OrderEntries) error
}

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (or *OrderRepository) GetAmountEntries(e OrderEntries) (int, error) {
	var amount int
	err := or.db.QueryRow("select amount from users where id = $1;", e.UserID).Scan(&amount)
	if err != nil {
		log.Fatal(err.Error())
		return 0, err
	}
	return amount, nil
}

func (or *OrderRepository) Deduct(amount int, e OrderEntries) error {
	var after int
	err := or.db.QueryRow("update users set amount = $1 where id=$2 returning amount;", amount, e.UserID).Scan(&after)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	log.Printf("userId: %v, amount: %v, before: %v, after: %v", e.UserID, e.Amount, amount, after)
	return nil
}
