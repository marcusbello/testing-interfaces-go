package main

import (
	"context"
)

type FinanceData struct {
	ResponseData interface{}
}

type UserData struct {
	ResponseData interface{}
}

type FinanceActions interface {
	WithdrawMoney(ctx context.Context, user string, amount float64) (*FinanceData, error)
}


type UserActions struct {
	Financials FinanceActions
	// ...
}

func (ua *UserActions) UserWithdrawMoney(user string, amount float64) (*UserData, error)  {
	var err error

	//some checks
	err = KYCStatus(user)
	if err != nil {
		return nil, err
	}
	err = CompareBalance(user, amount)
	if err != nil {
		return nil, err
	} 

	// withdraw money for user
	data, err := ua.Financials.WithdrawMoney(context.Background(), user, amount)
	if err != nil {
		return nil, err
	}
	response := &UserData{
		ResponseData: data,
	}

	return response, nil
}

func KYCStatus(user string) error {
	return nil
}

func CompareBalance(user string, amount float64) error {
	return nil
}


