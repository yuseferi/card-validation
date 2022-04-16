package card

// you can add a new schemes here
var schemeList = []scheme{
	{name: MasterCard, regExRule: `^5[1-5]\d{14}|(222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)\d{12}$`},
	{name: Visa, regExRule: `^4\d{12}(?:\d{3,6})?$`},
	{name: Maestro, regExRule: `^50|5[6-8]\d{10,17}|6\d{11,18}$`},
	{name: JCB, regExRule: `^(352[89]|35[3-8][0-9])\d{11,15}$`},
	{name: AmericanExpress, regExRule: `^3[4-7]\d{13}$`},
}

const (
	MasterCard      = "Master"
	Visa            = "Visa"
	Maestro         = "Maestro"
	JCB             = "JCB"
	AmericanExpress = "American Express"
)
