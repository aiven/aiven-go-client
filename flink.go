package aiven

type (
	// shared fields by some responses
	flinkPosition struct {
		CharacterNumber    int `json:"character_number"`
		EndCharacterNumber int `json:"end_character_number"`
		EndLineNumber      int `json:"end_line_number"`
		LineNumber         int `json:"line_number"`
	}
)
