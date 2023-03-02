// Package token implements the token-related parts of the staking API.
package token

const (
	// ModuleName is a unique module name for the staking/token module.
	ModuleName = "staking/token"

	// Maximum length of the token symbol.
	TokenSymbolMaxLength = 8
	// Regular expression defining valid token symbol characters.
	TokenSymbolRegexp = "^[A-Z]+$" // nolint: gosec // Not that kind of token :).
	// Maximum value of token's value base-10 exponent.
	TokenValueExponentMaxValue = 20
)

// ErrInvalidTokenValueExponent is the error returned when an invalid token's
// value base-10 exponent is specified.
// removed var statement
