package api

import (
	"github.com/ericlamnguyen/simple-bank/util"
	"github.com/go-playground/validator/v10"
)

// validCurrency will validate that the requested currency is among supported currencies
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	// check if provided field is a valid string
	// using type assertion to check if a value can be converted to a string type
	// The reflection part is in fieldLevel.Field().Interface(), which dynamically retrieves a field from a struct as an interface{},
	// and later attempts to assert it as a string. This approach is commonly used when dealing with unknown or dynamic types.
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// provided field is a string, now check if currency is supported
		return util.IsSupportedCurrency(currency)
	}

	// provided field is not a valid string
	return false
}
