package payment

import (
	"strings"
)

// CryptoCurrency represents a supported cryptocurrency
type CryptoCurrency struct {
	Symbol      string  // BTC, ETH, LTC, etc.
	Name        string  // Bitcoin, Ethereum, Litecoin, etc.
	Decimals    int     // Number of decimal places
	MinAmount   float64 // Minimum amount for transactions
	IsActive    bool    // Whether this crypto is currently supported
}

// Supported cryptocurrencies (based on NowPayments)
var (
	Bitcoin = CryptoCurrency{
		Symbol:    "BTC",
		Name:      "Bitcoin", 
		Decimals:  8,
		MinAmount: 0.0001,
		IsActive:  true,
	}
	
	Ethereum = CryptoCurrency{
		Symbol:    "ETH",
		Name:      "Ethereum",
		Decimals:  18,
		MinAmount: 0.001,
		IsActive:  true,
	}
	
	Litecoin = CryptoCurrency{
		Symbol:    "LTC", 
		Name:      "Litecoin",
		Decimals:  8,
		MinAmount: 0.001,
		IsActive:  true,
	}
	
	BitcoinCash = CryptoCurrency{
		Symbol:    "BCH",
		Name:      "Bitcoin Cash",
		Decimals:  8,
		MinAmount: 0.001,
		IsActive:  true,
	}
	
	Ripple = CryptoCurrency{
		Symbol:    "XRP",
		Name:      "Ripple",
		Decimals:  6,
		MinAmount: 1.0,
		IsActive:  true,
	}
	
	Dogecoin = CryptoCurrency{
		Symbol:    "DOGE",
		Name:      "Dogecoin",
		Decimals:  8,
		MinAmount: 1.0,
		IsActive:  true,
	}
)

// GetSupportedCryptoCurrencies returns all supported cryptocurrencies
func GetSupportedCryptoCurrencies() []CryptoCurrency {
	return []CryptoCurrency{
		Bitcoin,
		Ethereum,
		Litecoin,
		BitcoinCash,
		Ripple,
		Dogecoin,
	}
}

// GetCryptoCurrencyBySymbol returns a cryptocurrency by its symbol
func GetCryptoCurrencyBySymbol(symbol string) (CryptoCurrency, error) {
	symbol = strings.ToUpper(strings.TrimSpace(symbol))
	
	cryptos := GetSupportedCryptoCurrencies()
	for _, crypto := range cryptos {
		if crypto.Symbol == symbol && crypto.IsActive {
			return crypto, nil
		}
	}
	
	return CryptoCurrency{}, ErrUnsupportedCrypto
}

// IsSupported checks if a cryptocurrency symbol is supported
func IsSupported(symbol string) bool {
	_, err := GetCryptoCurrencyBySymbol(symbol)
	return err == nil
}

// ValidateAmount checks if the amount meets minimum requirements
func (c CryptoCurrency) ValidateAmount(amount float64) error {
	if amount <= 0 {
		return ErrInvalidCryptoAmount
	}
	
	if amount < c.MinAmount {
		return ErrInvalidCryptoAmount
	}
	
	return nil
}

// FormatAmount formats the amount according to the cryptocurrency's decimals
func (c CryptoCurrency) FormatAmount(amount float64) float64 {
	// This would typically use proper decimal formatting
	// For simplicity, we'll just return the amount
	return amount
}

// IsActive checks if the cryptocurrency is currently active/supported
func (c CryptoCurrency) IsActiveCurrency() bool {
	return c.IsActive
}

// GetSymbol returns the cryptocurrency symbol
func (c CryptoCurrency) GetSymbol() string {
	return c.Symbol
}

// GetName returns the cryptocurrency name
func (c CryptoCurrency) GetName() string {
	return c.Name
}

// GetMinAmount returns the minimum transaction amount
func (c CryptoCurrency) GetMinAmount() float64 {
	return c.MinAmount
}