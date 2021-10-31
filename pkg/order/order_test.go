package order_test

import (
	"errors"
	"testing"

	"github.com/gustialfian/go-simple-payment/pkg/order"
)

func TestOrderIsValid(t *testing.T) {
	tests := []struct {
		input    order.Order
		expected bool
	}{
		{
			input: order.Order{
				Entries: []order.OrderEntries{
					{Name: "Trip fare", Amount: -18, UserID: "0"},
					{Name: "Service fee", Amount: -2, UserID: "0"},
					{Name: "Trip fare", Amount: 18, UserID: "1"},
					{Name: "Service fee", Amount: 2, UserID: "3"},
				},
			},
			expected: true,
		},
		{
			input: order.Order{
				Entries: []order.OrderEntries{
					{Name: "Trip fare", Amount: -18, UserID: "0"},
					{Name: "Service fee", Amount: -2, UserID: "0"},
					{Name: "Trip fare", Amount: 18, UserID: "1"},
					{Name: "Service fee", Amount: 3, UserID: "3"},
				},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		if test.input.IsValid() != test.expected {
			t.Error("TestOrderIsValid: {} inputed, {} expected", test.input, test.expected)
		}
	}
}

type OrderProccessRepositoryMock struct {
	User []struct {
		UserID string
		Amount int
	}
}

func (repo OrderProccessRepositoryMock) GetAmountEntries(oe order.OrderEntries) (int, error) {
	for _, u := range repo.User {
		if u.UserID == oe.UserID {
			return u.Amount, nil
		}
	}
	return 0, errors.New("user not found")
}

func (repo OrderProccessRepositoryMock) Deduct(amount int, oe order.OrderEntries) error {
	for _, u := range repo.User {
		if u.UserID == oe.UserID {
			u.Amount = amount
			return nil
		}
	}
	return errors.New("user not found")
}

func TestOrderProccess(t *testing.T) {
	tests := []struct {
		input    order.Order
		expected error
	}{
		{
			input: order.Order{
				Entries: []order.OrderEntries{
					{Name: "Trip fare", Amount: -18, UserID: "0"},
					{Name: "Service fee", Amount: -2, UserID: "0"},
					{Name: "Trip fare", Amount: 18, UserID: "1"},
					{Name: "Service fee", Amount: 2, UserID: "2"},
				},
			},
			expected: nil,
		},
		{
			input: order.Order{
				Entries: []order.OrderEntries{
					{Name: "Trip fare", Amount: -18, UserID: "0"},
					{Name: "Service fee", Amount: -2, UserID: "0"},
					{Name: "Trip fare", Amount: 18, UserID: "1"},
					{Name: "Service fee", Amount: 3, UserID: "3"},
				},
			},
			expected: errors.New("user not found"),
		},
	}

	repoMock := OrderProccessRepositoryMock{
		User: []struct {
			UserID string
			Amount int
		}{
			{
				UserID: "0",
				Amount: 100,
			},
			{
				UserID: "1",
				Amount: 100,
			},
			{
				UserID: "2",
				Amount: 100,
			},
		},
	}

	for _, test := range tests {
		if res := test.input.Proccess(repoMock); res != nil && res.Error() != test.expected.Error() {
			t.Error("TestOrderProccess: {} inputed, {} expected", res, test.expected)
		}
	}
}
