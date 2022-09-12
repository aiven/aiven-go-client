package convert

import (
	"errors"

	"github.com/aiven/aiven-go-client"
	"github.com/aiven/aiven-go-client/tools/exp/types"
	"github.com/mitchellh/copystructure"
)

// errUnexpected is the error that is returned when an unexpected error occurs.
var errUnexpected = errors.New("unexpected conversion error")

// UserConfigSchema converts aiven.UserConfigSchema to UserConfigSchema.
func UserConfigSchema(v aiven.UserConfigSchema) (*types.UserConfigSchema, error) {
	var cnp map[string]types.UserConfigSchema = nil

	if len(v.Properties) != 0 {
		cnp = make(map[string]types.UserConfigSchema, len(v.Properties))

		p, err := copystructure.Copy(v.Properties)
		if err != nil {
			return nil, err
		}

		ap, ok := p.(map[string]aiven.UserConfigSchema)
		if !ok {
			return nil, errUnexpected
		}

		for k, v := range ap {
			var cv *types.UserConfigSchema

			cv, err = UserConfigSchema(v)
			if err != nil {
				return nil, err
			}

			cnp[k] = *cv
		}
	}

	var cni *types.UserConfigSchema = nil

	if v.Items != nil {
		var err error

		cni, err = UserConfigSchema(*v.Items)
		if err != nil {
			return nil, err
		}
	}

	var cno []types.UserConfigSchema = nil

	if len(v.OneOf) != 0 {
		cno = make([]types.UserConfigSchema, len(v.OneOf))

		o, err := copystructure.Copy(v.OneOf)
		if err != nil {
			return nil, err
		}

		ao, ok := o.([]aiven.UserConfigSchema)
		if !ok {
			return nil, errUnexpected
		}

		for i, v := range ao {
			var cv *types.UserConfigSchema

			cv, err = UserConfigSchema(v)
			if err != nil {
				return nil, err
			}

			cno[i] = *cv
		}
	}

	var e []types.UserConfigSchemaEnumValue

	for _, v := range v.Enum {
		e = append(e, types.UserConfigSchemaEnumValue{Value: v})
	}

	return &types.UserConfigSchema{
		Title:       v.Title,
		Description: v.Description,
		Type:        v.Type,
		Default:     v.Default,
		Properties:  cnp,
		Items:       cni,
		OneOf:       cno,
		Enum:        e,
		Minimum:     v.Minimum,
		Maximum:     v.Maximum,
		MinLength:   v.MinLength,
		MaxLength:   v.MaxLength,
		MaxItems:    v.MaxItems,
		CreateOnly:  v.CreateOnly,
		Pattern:     v.Pattern,
		Example:     v.Example,
		UserError:   v.UserError,
	}, nil
}
