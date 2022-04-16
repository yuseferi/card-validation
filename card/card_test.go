package card

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type cardTestStruct struct {
	cardNumber     string
	expectedScheme string
	isValid        bool
}

func TestCard_Validate(t *testing.T) {
	var testCases = map[string]cardTestStruct{
		"test validation : validated card": {
			cardNumber:     "5237 2516 2477 8133",
			expectedScheme: MasterCard,
			isValid:        true,
		},
		"test validation : inValid card": {
			cardNumber:     "5237251624770000",
			expectedScheme: MasterCard,
			isValid:        false,
		},
		"getSchema - MasterCard": {
			cardNumber:     "5105105105105100",
			expectedScheme: MasterCard,
			isValid:        true,
		},
		"getSchema - Visa": {
			cardNumber:     "4012888888881881",
			expectedScheme: Visa,
			isValid:        true,
		},
		"getSchema - Maestro": {
			cardNumber:     "6759649826438453",
			expectedScheme: Maestro,
			isValid:        true,
		},
		"getSchema - JCB": {
			cardNumber:     "3530111333300000",
			expectedScheme: JCB,
			isValid:        true,
		},
		"getSchema - American Express": {
			cardNumber:     "378282246310005",
			expectedScheme: AmericanExpress,
			isValid:        true,
		},
		"getSchema - not defined scheme": {
			cardNumber:     "998282246310005",
			expectedScheme: unknownCreditCard.Error(),
			isValid:        true,
		},
	}
	for testCase, data := range testCases {
		t.Run(testCase, func(t *testing.T) {
			card := NewCard(data.cardNumber)
			assert.Equal(t, data.isValid, card.Validate())
			assert.Equal(t, data.expectedScheme, card.GetScheme())

		})
	}

}
