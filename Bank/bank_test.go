package bank

import "testing"

func TestBank(t *testing.T) {
	t.Run("Test balance check", func(t *testing.T) {
		bank := Bank{balance: 30}
		// The expected amount.
		want := 30
		// The actual amount.
		got := bank.Balance()
		assertBank(t, got, want)
	})
	t.Run("Test withdraw", func(t *testing.T) {
		bank := Bank{balance: 30}
		bank.Withdraw(10)
		// The expected amount.
		want := 20
		// The actual amount.
		got := bank.Balance()
		assertBank(t, got, want)
	})
	t.Run("Test deposit", func(t *testing.T) {
		bank := Bank{balance: 30}
		bank.Deposit(10)
		// The expected amount.
		want := 40
		// The actual amount.
		got := bank.Balance()
		assertBank(t, got, want)
	})
	t.Run("Test withdraw error", func(t *testing.T) {
		bank := Bank{balance: 30}
		got := bank.Withdraw(40)
		// The expected amount.
		want := ErrInsufficientFunds
		assertError(t, got, want)
	})

}

func assertBank(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
