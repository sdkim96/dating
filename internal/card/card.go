package card

type Card struct {
	ID       int      `json:"id,omitempty"`
	Human    Human    `json:"human"`
	Position string   `json:"position,omitempty"`
	Company  *Company `json:"company,omitempty"`
}

type Human struct {
	Name  string `json:"name" jsonschema:"required,description=The name of the person"`
	Email string `json:"email" jsonschema:"required,description=The email of the person"`
}

type Company struct {
	Name    string       `json:"name" jsonschema:"required,description=The name of the company"`
	Website string       `json:"website,omitempty" jsonschema:"description=The website of the company"`
	Address string       `json:"address,omitempty" jsonschema:"description=The address of the company"`
	Meta    *CompanyMeta `json:"meta,omitempty" jsonschema:"description=Additional metadata about the company"`
}

type CompanyMeta struct {
	BusinessDetail string `json:"business_detail,omitempty"`
	WorkerCount    *int   `json:"worker_count,omitempty"`
	Revenue        *int64 `json:"revenue,omitempty"`
}
