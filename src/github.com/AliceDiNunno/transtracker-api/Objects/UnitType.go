package Objects

/*Ok so a little note about how the Units values are set
   1 - First the symbol. If the power of the unit is below 0, the value will be negative (-) instead it will be positive.
       For exemple, "micro" (-6), "nano" (-9), "pico" (-12) etc... are negative while "" (0) and kilo (3) are positive
   2 - the first number arbitrarily represent a unit. Here, 1 is "mol", 2 is "grams", etc...
   3 - the last two numbers represent the actual power of the unit so 09 will either be nano or giga depending on the sign

   So I will take two examples:

       nmol is -109 because:

       1 - it's negative because it is "nano"
       2 - the mol is the number 1
       3 - nano is of power -9 so it is 09

       gram is 200

       1 - it's positive because it is of power 0
       2 - the gram is the number 2
       3 - since it is of power 0, its representation is 00

PS: Special items like "pill" or "pression" are numbered from 1 to 99. They aren't a unit per se but we can use them as well.
   For exemple if you take 2 pills, we store it as value 2 with unit "pill". Tho it could be bad practice
PS2: Also since we aren't using all the units, we aren't setting all of them. But it should be done for evolutivity later.
*/

const (
	//Special
	UNIT_PILL = 1
	UNIT_PRESSION = 2


	//Substance units
	UNIT_NMOL = -109

	//Mass units
	UNIT_PG = -212
	UNIT_UG = -206
	UNIT_MG = -203

	//Volume units
	UNIT_ML = -303
	UNIT_L = 300

	//Time units
	//Here we have a problem because time units are not linear, we don't really progress from one to one
	//For now I will take 0 for seconds, 1 for minutes and so on
	UNIT_H = 402
)


type Diffusion struct {
	Id string `json:"id"`
	Value float64 `json:"value"`
	Unit int64 `json:"unit"`
}

type Dosage struct {
	Id string `json:"id"`
	Value float64 `json:"value"`
	Unit int64 `json:"unit"`

	Diffusion *Diffusion `json:"diffusion"`
}

type Product struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Laboratory string `json:"laboratory"`
	Manual string `json:"manual"`
	Dosage Dosage `json:"dosage"`
	Administration int64 `json:"administration"`
	Type int64 `json:"type"`
	NumberOfIntakesInABox int `json:"numberOfIntakesInABox"`
}

