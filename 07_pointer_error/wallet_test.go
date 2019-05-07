package wallet

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		startingBitcoint := Bitcoin(20)
		wallet := Wallet{balance: startingBitcoint}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw with insufficient funds", func(t *testing.T) {
		startingBitcoint := Bitcoin(20)
		wallet := Wallet{balance: startingBitcoint}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBitcoint)
		assertError(t, err, ErrInsufficientFunds)

	})
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get an error but wanted")
	}

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {

	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
