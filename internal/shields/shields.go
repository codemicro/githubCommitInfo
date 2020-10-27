package shields

type ShieldsResponse struct {
	SchemaVersion int `json:"schemaVersion"`
	Label string `json:"label"`
	Message string `json:"message"`
	Colour string `json:"color"`
}

func NewShield(label, message, colour string) *ShieldsResponse {
	return &ShieldsResponse{
		SchemaVersion: 1,
		Label: label,
		Message: message,
		Colour: colour,
	}
}