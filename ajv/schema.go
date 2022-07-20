package ajv

import "encoding/json"

type Schema struct {
	Typ        string            `json:"type"`
	Properties map[string]Schema `json:"properties,omitempty"`
	Items      *Schema           `json:"items,omitempty"`
	Nullable   bool              `json:"nullable,omitempty"`
	Required   []string          `json:"required,omitempty"`
}

func (s Schema) String() (string, error) {

	raw, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return "", err
	}

	return string(raw), nil
}
