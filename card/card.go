package card

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var unknownCreditCard = errors.New("unknown credit card")

type Card struct {
	Number string
}

func NewCard(number string) *Card {
	// Remove spaces
	return &Card{
		Number: strings.ReplaceAll(number, " ", ""),
	}
}

// GetScheme Get card scheme
func (c *Card) GetScheme() string {

	for _, scheme := range schemeList {
		re := regexp.MustCompile(scheme.regExRule)
		if re.MatchString(c.Number) {
			return scheme.name
		}
	}
	return unknownCreditCard.Error()
}

// Validate the card number
func (c *Card) Validate() bool {
	var sum, mod int
	var alternate bool
	var err error

	numberLen := len(c.Number)
	// Base on the table given in the code challenge
	// valid range is 12 to 19
	if numberLen < 12 || numberLen > 19 {
		return false
	}

	// Parse all numbers of the card into a for loop
	for i := numberLen - 1; i > -1; i-- {
		// Takes the mod, converting the current number in integer
		if mod, err = strconv.Atoi(string(c.Number[i])); err != nil {
			return false
		}
		if alternate {
			mod *= 2
			if mod > 9 {
				mod = (mod % 10) + 1
			}
		}
		alternate = !alternate
		sum += mod
	}

	return sum%10 == 0
}

/** Notice: this is another solution to detect scheme, but it's not flexible as the main solution
func (c *Card) GetScheme() string {

	switch {
	case c.isMasterCard():
		return MasterCard
	case c.isVisa():
		return Visa
	case c.isMaestro():
		return Maestro
	case c.isJCB():
		return JCB
	case c.isAmericanExpress():
		return AmericanExpress
	default:
		return unknownCreditCard.Error()
	}
}

func (c *Card) isMasterCard() bool {
	re := regexp.MustCompile(`^5[1-5]\d{14}|(222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)\d{12}$`)
	return re.MatchString(c.Number)
}

func (c *Card) isVisa() bool {
	re := regexp.MustCompile(`^4\d{12}(?:\d{3,6})?$`)
	return re.MatchString(c.Number)
}
func (c *Card) isMaestro() bool {
	re := regexp.MustCompile(`^50|5[6-8]\d{10,17}|6\d{11,18}$`)
	return re.MatchString(c.Number)
}
func (c *Card) isJCB() bool {
	re := regexp.MustCompile(`^(352[89]|35[3-8][0-9])\d{11,15}$`)
	return re.MatchString(c.Number)
}
func (c *Card) isAmericanExpress() bool {
	re := regexp.MustCompile(`^3[4-7]\d{13}$`)
	return re.MatchString(c.Number)
}

*/
