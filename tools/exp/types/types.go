package types

// UserConfigSchemaDeprecationInfo is a struct that contains the deprecation info for a user config schema entry.
type UserConfigSchemaDeprecationInfo struct {
	IsDeprecated      bool   `yaml:"is_deprecated,omitempty"`
	DeprecationNotice string `yaml:"deprecation_notice,omitempty"`
}

// UserConfigSchemaEnumValue is a struct that contains the enum value for a user config schema entry.
type UserConfigSchemaEnumValue struct {
	UserConfigSchemaDeprecationInfo `yaml:",inline"`

	Value interface{} `yaml:"value"`
}

// UserConfigSchema represents an output schema for the user config.
type UserConfigSchema struct {
	UserConfigSchemaDeprecationInfo `yaml:",inline"`

	Title       string                      `yaml:"title,omitempty"`
	Description string                      `yaml:"description,omitempty"`
	Type        interface{}                 `yaml:"type,omitempty"`
	Default     interface{}                 `yaml:"default,omitempty"`
	Required    []string                    `yaml:"required,omitempty"`
	Properties  map[string]UserConfigSchema `yaml:"properties,omitempty"`
	Items       *UserConfigSchema           `yaml:"items,omitempty"`
	OneOf       []UserConfigSchema          `yaml:"one_of,omitempty"`
	Enum        []UserConfigSchemaEnumValue `yaml:"enum,omitempty"`
	Minimum     float64                     `yaml:"minimum,omitempty"`
	Maximum     float64                     `yaml:"maximum,omitempty"`
	MinLength   int                         `yaml:"min_length,omitempty"`
	MaxLength   int                         `yaml:"max_length,omitempty"`
	MaxItems    int                         `yaml:"max_items,omitempty"`
	CreateOnly  bool                        `yaml:"create_only,omitempty"`
	Pattern     string                      `yaml:"pattern,omitempty"`
	Example     interface{}                 `yaml:"example,omitempty"`
	UserError   string                      `yaml:"user_error,omitempty"`
}

// GenerationResult represents the result of a generation.
type GenerationResult map[int]map[string]UserConfigSchema

// ReadResult represents the result of a read.
type ReadResult map[int]map[string]UserConfigSchema

// DiffResult represents the result of a diff.
type DiffResult map[int]map[string]UserConfigSchema

const (
	// KeyServiceTypes is the key for the service types.
	KeyServiceTypes int = iota

	// KeyIntegrationTypes is the key for the integration types.
	KeyIntegrationTypes

	// KeyIntegrationEndpointTypes is the key for the integration endpoint types.
	KeyIntegrationEndpointTypes
)
