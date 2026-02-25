package card

type CardCreate struct {
	Human    Human    `json:"human" jsonschema:"required,description=The identity of the person"`
	Position string   `json:"position,omitempty" jsonschema:"description=The job position of the person"`
	Company  *Company `json:"company,omitempty" jsonschema:"description=The company associated with the person"`
}
