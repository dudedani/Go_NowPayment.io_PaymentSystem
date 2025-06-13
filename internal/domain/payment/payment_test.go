package payment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test helper functions

// createTestPaymentMethod creates a valid payment method for testing
func createTestPaymentMethod() PaymentMethod {
	method, _ := NewPaymentMethod("BTC", "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", 30)
	return method
}

// Tests for PaymentStatus

func TestPaymentStatus(t *testing.T) {
	t.Run("valid statuses", func(t *testing.T) {
		validStatuses := []PaymentStatus{
			StatusPending, StatusConfirming, StatusConfirmed, 
			StatusFailed, StatusExpired, StatusRefunded, StatusCancelled,
		}
		
		for _, status := range validStatuses {
			assert.True(t, status.IsValid())
		}
	})
	
	t.Run("invalid status", func(t *testing.T) {
		invalidStatus := PaymentStatus("INVALID")
		assert.False(t, invalidStatus.IsValid())
	})
	
	t.Run("completed status", func(t *testing.T) {
		assert.True(t, StatusConfirmed.IsCompleted())
		assert.False(t, StatusPending.IsCompleted())
		assert.False(t, StatusFailed.IsCompleted())
	})
	
	t.Run("final status", func(t *testing.T) {
		finalStatuses := []PaymentStatus{
			StatusConfirmed, StatusFailed, StatusExpired, StatusRefunded, StatusCancelled,
		}
		
		for _, status := range finalStatuses {
			assert.True(t, status.IsFinal())
		}
		
		assert.False(t, StatusPending.IsFinal())
		assert.False(t, StatusConfirming.IsFinal())
	})
	
	t.Run("can be refunded", func(t *testing.T) {
		assert.True(t, StatusConfirmed.CanBeRefunded())
		assert.False(t, StatusPending.CanBeRefunded())
		assert.False(t, StatusFailed.CanBeRefunded())
	})
	
	t.Run("can be cancelled", func(t *testing.T) {
		assert.True(t, StatusPending.CanBeCancelled())
		assert.True(t, StatusConfirming.CanBeCancelled())
		assert.False(t, StatusConfirmed.CanBeCancelled())
		assert.False(t, StatusFailed.CanBeCancelled())
	})
}

// Tests for CryptoCurrency

func TestCryptoCurrency(t *testing.T) {
	t.Run("get supported cryptocurrencies", func(t *testing.T) {
		cryptos := GetSupportedCryptoCurrencies()
		
		assert.GreaterOrEqual(t, len(cryptos), 6) // At least 6 cryptos
		
		// Check specific cryptocurrencies
		symbols := make(map[string]bool)
		for _, crypto := range cryptos {
			symbols[crypto.Symbol] = true
		}
		
		assert.True(t, symbols["BTC"])
		assert.True(t, symbols["ETH"])
		assert.True(t, symbols["LTC"])
	})
	
	t.Run("get cryptocurrency by symbol", func(t *testing.T) {
		crypto, err := GetCryptoCurrencyBySymbol("BTC")
		
		assert.NoError(t, err)
		assert.Equal(t, "BTC", crypto.Symbol)
		assert.Equal(t, "Bitcoin", crypto.Name)
		assert.True(t, crypto.IsActive)
	})
	
	t.Run("get cryptocurrency by symbol case insensitive", func(t *testing.T) {
		crypto, err := GetCryptoCurrencyBySymbol("btc")
		
		assert.NoError(t, err)
		assert.Equal(t, "BTC", crypto.Symbol)
	})
	
	t.Run("unsupported cryptocurrency", func(t *testing.T) {
		_, err := GetCryptoCurrencyBySymbol("INVALID")
		
		assert.Error(t, err)
		assert.Equal(t, ErrUnsupportedCrypto, err)
	})
	
	t.Run("is supported", func(t *testing.T) {
		assert.True(t, IsSupported("BTC"))
		assert.True(t, IsSupported("ETH"))
		assert.False(t, IsSupported("INVALID"))
	})
	
	t.Run("validate amount", func(t *testing.T) {
		crypto := Bitcoin
		
		// Valid amount
		err := crypto.ValidateAmount(0.001)
		assert.NoError(t, err)
		
		// Amount too small
		err = crypto.ValidateAmount(0.00001)
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidCryptoAmount, err)
		
		// Negative amount
		err = crypto.ValidateAmount(-0.1)
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidCryptoAmount, err)
	})
	
	t.Run("cryptocurrency properties", func(t *testing.T) {
		crypto := Bitcoin
		
		assert.Equal(t, "BTC", crypto.GetSymbol())
		assert.Equal(t, "Bitcoin", crypto.GetName())
		assert.Equal(t, 0.0001, crypto.GetMinAmount())
		assert.True(t, crypto.IsActiveCurrency())
	})
}

// Tests for PaymentMethod

func TestPaymentMethod(t *testing.T) {
	t.Run("create valid payment method", func(t *testing.T) {
		method, err := NewPaymentMethod("BTC", "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", 30)
		
		assert.NoError(t, err)
		assert.Equal(t, TypeCryptocurrency, method.Type)
		assert.Equal(t, "BTC", method.CryptoCurrency.Symbol)
		assert.Equal(t, "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", method.WalletAddress)
		assert.False(t, method.IsExpired())
	})
	
	t.Run("cannot create payment method with empty wallet address", func(t *testing.T) {
		_, err := NewPaymentMethod("BTC", "", 30)
		
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidWalletAddress, err)
	})
	
	t.Run("cannot create payment method with unsupported crypto", func(t *testing.T) {
		_, err := NewPaymentMethod("INVALID", "address123", 30)
		
		assert.Error(t, err)
		assert.Equal(t, ErrUnsupportedCrypto, err)
	})
	
	t.Run("payment method expiration", func(t *testing.T) {
		// Create method that expires in 0 minutes (immediately)
		method, _ := NewPaymentMethod("BTC", "address123", 0)
		
		// Should be expired immediately
		time.Sleep(1 * time.Millisecond)
		assert.True(t, method.IsExpired())
		assert.Equal(t, time.Duration(0), method.TimeUntilExpiry())
	})
	
	t.Run("payment method time until expiry", func(t *testing.T) {
		method, _ := NewPaymentMethod("BTC", "address123", 30)
		
		duration := method.TimeUntilExpiry()
		assert.Greater(t, duration, 29*time.Minute)
		assert.LessOrEqual(t, duration, 30*time.Minute)
	})
	
	t.Run("validate amount", func(t *testing.T) {
		method := createTestPaymentMethod()
		
		// Valid amount
		err := method.ValidateAmount(0.001)
		assert.NoError(t, err)
		
		// Invalid amount
		err = method.ValidateAmount(0.00001)
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidCryptoAmount, err)
	})
	
	t.Run("payment method properties", func(t *testing.T) {
		method := createTestPaymentMethod()
		
		assert.Equal(t, "BTC", method.GetCryptoSymbol())
		assert.Equal(t, "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", method.GetWalletAddress())
		assert.True(t, method.IsActive())
	})
}

// Tests for Payment Entity

func TestNewPayment(t *testing.T) {
	t.Run("create valid payment", func(t *testing.T) {
		payment, err := NewPayment("order-123", 100.0, "USD", "BTC", "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", 30)
		
		assert.NoError(t, err)
		assert.NotEmpty(t, payment.ID)
		assert.Equal(t, "order-123", payment.OrderID)
		assert.Equal(t, 100.0, payment.Amount)
		assert.Equal(t, "USD", payment.Currency)
		assert.Equal(t, "BTC", payment.CryptoCurrency.Symbol)
		assert.Equal(t, StatusPending, payment.Status)
		assert.False(t, payment.IsExpired())
		assert.True(t, payment.IsPending())
		assert.False(t, payment.IsCompleted())
	})
	
	t.Run("cannot create payment with empty order ID", func(t *testing.T) {
		_, err := NewPayment("", 100.0, "USD", "BTC", "address123", 30)
		
		assert.Error(t, err)
		assert.Equal(t, ErrEmptyOrderID, err)
	})
	
	t.Run("cannot create payment with invalid amount", func(t *testing.T) {
		_, err := NewPayment("order-123", -50.0, "USD", "BTC", "address123", 30)
		
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidAmount, err)
	})
	
	t.Run("cannot create payment with empty currency", func(t *testing.T) {
		_, err := NewPayment("order-123", 100.0, "", "BTC", "address123", 30)
		
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidCurrency, err)
	})
	
	t.Run("cannot create payment with unsupported crypto", func(t *testing.T) {
		_, err := NewPayment("order-123", 100.0, "USD", "INVALID", "address123", 30)
		
		assert.Error(t, err)
		assert.Equal(t, ErrUnsupportedCrypto, err)
	})
}

func TestPaymentCryptoAmount(t *testing.T) {
	t.Run("update crypto amount", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		
		err := payment.UpdateCryptoAmount(0.001)
		
		assert.NoError(t, err)
		assert.Equal(t, 0.001, payment.CryptoAmount)
	})
	
	t.Run("cannot update crypto amount on final payment", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		payment.Status = StatusConfirmed
		
		err := payment.UpdateCryptoAmount(0.001)
		
		assert.Error(t, err)
		assert.Equal(t, ErrCannotUpdateFinalPayment, err)
	})
	
	t.Run("cannot update with invalid crypto amount", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		
		err := payment.UpdateCryptoAmount(0.00001) // Below minimum for BTC
		
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidCryptoAmount, err)
	})
}

func TestPaymentStatusTransitions(t *testing.T) {
	t.Run("mark as confirming", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		
		err := payment.MarkAsConfirming("abc123")
		
		assert.NoError(t, err)
		assert.Equal(t, StatusConfirming, payment.Status)
		assert.Equal(t, "abc123", payment.TransactionHash)
		assert.Equal(t, 0, payment.Confirmations)
		assert.True(t, payment.IsConfirming())
	})
	
	t.Run("cannot mark as confirming from wrong status", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		payment.Status = StatusConfirmed
		
		err := payment.MarkAsConfirming("abc123")
		
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidStatusTransition, err)
	})
	
	t.Run("cannot mark as confirming with empty transaction hash", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		
		err := payment.MarkAsConfirming("")
		
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidTransactionHash, err)
	})
	
	t.Run("update confirmations", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		payment.MarkAsConfirming("abc123")
		
		err := payment.UpdateConfirmations(1)
		
		assert.NoError(t, err)
		assert.Equal(t, 1, payment.Confirmations)
		assert.Equal(t, StatusConfirming, payment.Status)
	})
	
	t.Run("auto confirm with enough confirmations", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		payment.MarkAsConfirming("abc123")
		
		err := payment.UpdateConfirmations(2) // BTC requires 2 confirmations
		
		assert.NoError(t, err)
		assert.Equal(t, 2, payment.Confirmations)
		assert.Equal(t, StatusConfirmed, payment.Status)
		assert.True(t, payment.IsCompleted())
	})
	
	t.Run("manual confirm payment", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		
		err := payment.MarkAsConfirmed()
		
		assert.NoError(t, err)
		assert.Equal(t, StatusConfirmed, payment.Status)
		assert.True(t, payment.IsCompleted())
	})
	
	t.Run("cannot confirm already confirmed payment", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		payment.Status = StatusConfirmed
		
		err := payment.MarkAsConfirmed()
		
		assert.Error(t, err)
		assert.Equal(t, ErrPaymentAlreadyConfirmed, err)
	})
	
	t.Run("mark as failed", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		
		err := payment.MarkAsFailed()
		
		assert.NoError(t, err)
		assert.Equal(t, StatusFailed, payment.Status)
	})
	
	t.Run("mark as expired", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		
		err := payment.MarkAsExpired()
		
		assert.NoError(t, err)
		assert.Equal(t, StatusExpired, payment.Status)
	})
}

func TestPaymentCancellation(t *testing.T) {
	t.Run("cancel pending payment", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		
		err := payment.Cancel()
		
		assert.NoError(t, err)
		assert.Equal(t, StatusCancelled, payment.Status)
	})
	
	t.Run("cancel confirming payment", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		payment.Status = StatusConfirming
		
		err := payment.Cancel()
		
		assert.NoError(t, err)
		assert.Equal(t, StatusCancelled, payment.Status)
	})
	
	t.Run("cannot cancel confirmed payment", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		payment.Status = StatusConfirmed
		
		err := payment.Cancel()
		
		assert.Error(t, err)
		assert.Equal(t, ErrCannotCancelPayment, err)
	})
}

func TestPaymentRefunds(t *testing.T) {
	t.Run("full refund confirmed payment", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		payment.UpdateCryptoAmount(0.001)
		payment.Status = StatusConfirmed
		
		err := payment.Refund()
		
		assert.NoError(t, err)
		assert.Equal(t, 0.001, payment.RefundedAmount)
		assert.Equal(t, StatusRefunded, payment.Status)
		assert.True(t, payment.IsFullyRefunded())
		assert.NotNil(t, payment.RefundedAt)
	})
	
	t.Run("partial refund confirmed payment", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		payment.UpdateCryptoAmount(0.001)
		payment.Status = StatusConfirmed
		
		err := payment.PartialRefund(0.0005)
		
		assert.NoError(t, err)
		assert.Equal(t, 0.0005, payment.RefundedAmount)
		assert.Equal(t, StatusConfirmed, payment.Status) // Still confirmed, not fully refunded
		assert.False(t, payment.IsFullyRefunded())
		assert.Equal(t, 0.0005, payment.GetRemainingRefundableAmount())
	})
	
	t.Run("cannot refund non-confirmed payment", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		
		err := payment.Refund()
		
		assert.Error(t, err)
		assert.Equal(t, ErrCannotRefundPayment, err)
	})
	
	t.Run("cannot refund more than payment amount", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		payment.UpdateCryptoAmount(0.001)
		payment.Status = StatusConfirmed
		
		err := payment.PartialRefund(0.002)
		
		assert.Error(t, err)
		assert.Equal(t, ErrRefundAmountExceedsPayment, err)
	})
	
	t.Run("set refund transaction hash", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		payment.UpdateCryptoAmount(0.001)
		payment.Status = StatusConfirmed
		payment.Refund()
		
		err := payment.SetRefundTransactionHash("refund-hash-123")
		
		assert.NoError(t, err)
		assert.Equal(t, "refund-hash-123", payment.RefundTransactionHash)
	})
}

func TestPaymentValidation(t *testing.T) {
	t.Run("validate exact amount", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		payment.UpdateCryptoAmount(0.001)
		
		err := payment.ValidateAmount(0.001)
		
		assert.NoError(t, err)
	})
	
	t.Run("validate amount within tolerance", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		payment.UpdateCryptoAmount(0.001)
		
		err := payment.ValidateAmount(0.0009995) // Slightly less but within tolerance
		
		assert.NoError(t, err)
	})
	
	t.Run("insufficient amount", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		payment.UpdateCryptoAmount(0.001)
		
		err := payment.ValidateAmount(0.0005)
		
		assert.Error(t, err)
		assert.Equal(t, ErrInsufficientAmount, err)
	})
	
	t.Run("excessive amount", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		payment.UpdateCryptoAmount(0.001)
		
		err := payment.ValidateAmount(0.002)
		
		assert.Error(t, err)
		assert.Equal(t, ErrExcessiveAmount, err)
	})
}

func TestPaymentExternalService(t *testing.T) {
	t.Run("set now payments ID", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		
		err := payment.SetNowPaymentsID("np-123456")
		
		assert.NoError(t, err)
		assert.Equal(t, "np-123456", payment.NowPaymentsID)
	})
	
	t.Run("cannot set empty now payments ID", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		
		err := payment.SetNowPaymentsID("")
		
		assert.Error(t, err)
		assert.Equal(t, ErrEmptyPaymentID, err)
	})
	
	t.Run("set callback URL", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		
		err := payment.SetCallbackURL("https://example.com/webhook")
		
		assert.NoError(t, err)
		assert.Equal(t, "https://example.com/webhook", payment.CallbackURL)
	})
}

func TestPaymentExpiration(t *testing.T) {
	t.Run("payment expiration", func(t *testing.T) {
		// Create payment that expires immediately
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 0)
		
		// Wait a moment for expiration
		time.Sleep(1 * time.Millisecond)
		
		assert.True(t, payment.IsExpired())
		assert.Equal(t, time.Duration(0), payment.GetTimeUntilExpiry())
	})
	
	t.Run("cannot mark expired payment as confirming", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 0)
		time.Sleep(1 * time.Millisecond)
		
		err := payment.MarkAsConfirming("abc123")
		
		assert.Error(t, err)
		assert.Equal(t, ErrPaymentExpired, err)
	})
}

func TestPaymentQueryMethods(t *testing.T) {
	t.Run("query methods on pending payment", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		
		assert.True(t, payment.IsPending())
		assert.False(t, payment.IsConfirming())
		assert.False(t, payment.IsCompleted())
		assert.True(t, payment.CanBeCancelled())
		assert.False(t, payment.CanBeRefunded())
		assert.Equal(t, "BTC", payment.GetCryptoSymbol())
		assert.Equal(t, "address123", payment.GetWalletAddress())
		assert.Equal(t, 0.0, payment.GetRemainingRefundableAmount())
	})
	
	t.Run("query methods on confirmed payment", func(t *testing.T) {
		payment, _ := NewPayment("order-123", 100.0, "USD", "BTC", "address123", 30)
		payment.UpdateCryptoAmount(0.001)
		payment.Status = StatusConfirmed
		
		assert.False(t, payment.IsPending())
		assert.False(t, payment.IsConfirming())
		assert.True(t, payment.IsCompleted())
		assert.False(t, payment.CanBeCancelled())
		assert.True(t, payment.CanBeRefunded())
		assert.Equal(t, 0.001, payment.GetRemainingRefundableAmount())
	})
}

func TestRequiredConfirmations(t *testing.T) {
	testCases := []struct {
		crypto   string
		expected int
	}{
		{"BTC", 2},
		{"ETH", 12},
		{"LTC", 6},
		{"BCH", 6},
		{"XRP", 1},
		{"DOGE", 6},
	}
	
	for _, tc := range testCases {
		t.Run(tc.crypto+" required confirmations", func(t *testing.T) {
			payment, _ := NewPayment("order-123", 100.0, "USD", tc.crypto, "address123", 30)
			
			assert.Equal(t, tc.expected, payment.RequiredConfirmations)
		})
	}
}