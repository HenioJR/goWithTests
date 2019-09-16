package pointers

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want string) {
		t.Helper()

		if got == nil {
			t.Error("wanted an error bug didn't get one")
		}

		if got.Error() != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		t.Helper()

		if got != nil {
			t.Error("got an error but didn't want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(5))
		assertBalance(t, wallet, Bitcoin(15))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, ErrInsufficientFunds.Error())
	})
}
