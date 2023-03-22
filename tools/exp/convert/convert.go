package convert

import (
	"errors"

	"github.com/mitchellh/copystructure"
	"golang.org/x/exp/slices"

	"github.com/aiven/aiven-go-client"
	"github.com/aiven/aiven-go-client/tools/exp/types"
)

// errUnexpected is the error that is returned when an unexpected error occurs.
var errUnexpected = errors.New("unexpected conversion error")

// maxSafeNumber the last number in double precision floating point that can make valid comparison
const maxSafeNumber = float64(1<<53 - 1)

// UserConfigSchema converts aiven.UserConfigSchema to UserConfigSchema.
func UserConfigSchema(v aiven.UserConfigSchema) (*types.UserConfigSchema, error) {
	var r []string

	for _, v := range v.Required {
		r = append(r, v)
	}

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

	// YAML uses scientific notation for floats, they won't change that
	// https://github.com/go-yaml/yaml/issues/669
	var max *float64
	if v.Maximum != nil {
		// If this is an integer it has to be lte maxSafeNumber
		// Otherwise, uses it as is
		if !slices.Contains(normalizeTypes(v.Type), "integer") || *v.Maximum <= maxSafeNumber {
			max = v.Maximum
		}
	}

	return &types.UserConfigSchema{
		Title:       v.Title,
		Description: v.Description,
		Type:        v.Type,
		Default:     v.Default,
		Required:    r,
		Properties:  cnp,
		Items:       cni,
		OneOf:       cno,
		Enum:        e,
		Minimum:     v.Minimum,
		Maximum:     max,
		MinLength:   v.MinLength,
		MaxLength:   v.MaxLength,
		MaxItems:    v.MaxItems,
		CreateOnly:  v.CreateOnly,
		Pattern:     v.Pattern,
		Example:     v.Example,
		UserError:   v.UserError,
	}, nil
}

// normalizeTypes json type field can be a string or list of strings.
// Turns into a slice of strings
func normalizeTypes(t any) []string {
	s, ok := t.(string)
	if ok {
		return []string{s}
	}

	typeList := make([]string, 0)
	a, ok := t.([]any)
	if !ok {
		return typeList
	}

	for _, v := range a {
		if vv, ok := v.(string); ok {
			typeList = append(typeList, vv)
		}
	}

	return typeList
}
