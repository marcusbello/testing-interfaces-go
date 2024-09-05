package main

import (
	"context"
	"errors"
	"testing"
)

type MockFinanceActions struct {
	WithdrawMoneyFunc func(ctx context.Context, user string, amount float64) (*FinanceData, error)
}

func (mfa *MockFinanceActions) WithdrawMoney(ctx context.Context, user string, amount float64) (*FinanceData, error) {
	return mfa.WithdrawMoneyFunc(ctx, user, amount)
}

func TestWithdrawMoney(t *testing.T) {
	userAlice := "alice"
	userDonny := "donny"

	expectedDataforDonny := &FinanceData{
		ResponseData: map[string]interface{}{
			"id": 7,
			"user": userDonny,
			"message": "success",
			"balance": 5435.22,
		},
	}

	mock := MockFinanceActions{}
	mock.WithdrawMoneyFunc = func(ctx context.Context, user string, amount float64) (*FinanceData, error) {

		if user == userAlice {
			return nil, errors.New("alice suspended")
		}

		if user == userDonny {
			data := expectedDataforDonny
			return data, nil
		}

		return nil, nil
	}

	ua := &UserActions{
		Financials: &mock,
	}

	aliceData, err := ua.UserWithdrawMoney(userAlice, 500.2)
	if err == nil {
		t.Errorf("failed: %v", aliceData)
	}

	donnyData, err := ua.UserWithdrawMoney(userDonny, 9670.2)
	if err != nil {
		t.Log(donnyData)
		t.Errorf("failed: %v", donnyData)
	}

}