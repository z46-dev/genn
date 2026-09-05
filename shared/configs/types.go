package configs

type (
	Type         uint8
	DefinitionID uint16

	Definition struct {
		ID      DefinitionID  `json:"id"`
		Type    Type          `json:"type,omitempty"`
		Name    string        `json:"name,omitempty"`
		Parents []*Definition `json:"parents,omitempty"`
	}

	Gun struct{}

	Turret struct {
		Definitions []*Definition `json:"definitions,omitempty"`
	}
)
