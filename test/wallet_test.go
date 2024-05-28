package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err != nil {
			t.Errorf("Error: %v", err.Error())
		}
	}
	wallet := Wallet{}
	t.Run("wallet deposit", func(t *testing.T) {
		wallet.Deposit(Bitcoin(1000))
		assertBalance(t, wallet, 1000)
	})
	t.Run("wallet withdraw", func(t *testing.T) {
		err := wallet.Withdraw(500)
		assertError(t, err)
		assertBalance(t, wallet, 500)
	})

	t.Run("Wallet", func(t *testing.T) { fmt.Print("as") })
}
