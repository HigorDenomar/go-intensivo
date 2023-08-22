package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_It_Gets_An_Error_If_ID_Is_Blank(test *testing.T) {
	order := Order{}

	assert.Error(test, order.Validate(), "id is required")
}

func Test_If_It_Gets_An_Error_If_Price_Is_Blank(test *testing.T) {
	order := Order{ID: "123"}

	assert.Error(test, order.Validate(), "invalid price")
}

func Test_If_It_Gets_An_Error_If_Tax_Is_Blank(test *testing.T) {
	order := Order{ID: "123", Price: 10.0}

	assert.Error(test, order.Validate(), "invalid tax")
}

func Test_Final_Price(test *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 1.0}

	assert.NoError(test, order.Validate())
	assert.Equal(test, "123", order.ID)
	assert.Equal(test, 10.0, order.Price)
	assert.Equal(test, 1.0, order.Tax)

	order.CalculateFinalPrice()
	assert.Equal(test, 11.0, order.FinalPrice)
}
