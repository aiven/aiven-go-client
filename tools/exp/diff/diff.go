package diff

import (
	"context"
	"fmt"

	"github.com/aiven/aiven-go-client/tools/exp/types"
	"github.com/aiven/aiven-go-client/tools/exp/util"
	"golang.org/x/exp/maps"
	"golang.org/x/sync/errgroup"
)

// logger is a pointer to the logger.
var logger *util.Logger

// genResult is the result of the generation process.
var genResult types.GenerationResult

// readResult is the result of the read process.
var readResult types.ReadResult

// result is the result of the diff process.
var result types.DiffResult

// diff is a function that diffs two maps.
func diff(gen map[string]types.UserConfigSchema, read map[string]types.UserConfigSchema) (map[string]types.UserConfigSchema, error) {
	if len(read) == 0 {
		return gen, nil
	}

	resultSchema := map[string]types.UserConfigSchema{}

	for k, v := range read {
		nv := v

		d, err := diff(gen[k].Properties, nv.Properties)
		if err != nil {
			return nil, err
		}

		nv.Properties = d

		if nv.Items != nil && gen[k].Items != nil {
			nv.Items.Title = gen[k].Items.Title

			nv.Items.Description = gen[k].Items.Description

			nv.Items.Type = gen[k].Items.Type

			d, err = diff(gen[k].Items.Properties, nv.Items.Properties)
			if err != nil {
				return nil, err
			}

			nv.Items.Properties = d

			if len(nv.Items.OneOf) != 0 {
				for kn, vn := range nv.Items.OneOf {
					if len(gen[k].Items.OneOf) > kn {
						d, err = diff(gen[k].Items.OneOf[kn].Properties, vn.Properties)
						if err != nil {
							return nil, err
						}

						nv.Items.OneOf[kn].Properties = d
					}

					genExists := false

					for _, vg := range gen[k].Items.OneOf {
						if vn.Title == vg.Title {
							genExists = true

							break
						}
					}

					if !genExists {
						nv.Items.OneOf[kn].IsDeprecated = true

						if nv.Items.OneOf[kn].DeprecationNotice == "" {
							nv.Items.OneOf[kn].DeprecationNotice = "This item is deprecated."
						}
					}
				}

				if len(gen[k].Items.OneOf) != 0 {
					for _, vn := range gen[k].Items.OneOf {
						readExists := false

						for k, vr := range nv.Items.OneOf {
							if vn.Title == vr.Title {
								nv.Items.OneOf[k].Description = vn.Description

								nv.Items.OneOf[k].Type = vn.Type

								nv.Items.OneOf[k].MaxLength = vn.MaxLength

								nv.Items.OneOf[k].Pattern = vn.Pattern

								nv.Items.OneOf[k].Example = vn.Example

								readExists = true

								break
							}
						}

						if !readExists {
							nv.Items.OneOf = append(nv.Items.OneOf, vn)
						}
					}
				}
			}
		}

		for kn, vn := range nv.Enum {
			genExists := false

			vnv := fmt.Sprintf("%v", vn.Value)

			for _, vg := range gen[k].Enum {
				vgv := fmt.Sprintf("%v", vg.Value)

				if vnv == vgv {
					genExists = true

					break
				}
			}

			if !genExists {
				nv.Enum[kn].IsDeprecated = true

				if nv.Enum[kn].DeprecationNotice == "" {
					nv.Enum[kn].DeprecationNotice = "This value is deprecated."
				}
			}

			for _, vn := range gen[k].Enum {
				readExists := false

				vnv := fmt.Sprintf("%v", vn.Value)

				for _, vr := range nv.Enum {
					vrv := fmt.Sprintf("%v", vr.Value)

					if vnv == vrv {
						readExists = true

						break
					}
				}

				if !readExists {
					nv.Enum = append(nv.Enum, vn)
				}
			}
		}

		if _, ok := gen[k]; !ok {
			nv.IsDeprecated = true

			if nv.DeprecationNotice == "" {
				nv.DeprecationNotice = "This property is deprecated."
			}
		} else {
			nv.Title = gen[k].Title

			nv.Description = gen[k].Description

			nv.Type = gen[k].Type

			nv.Default = gen[k].Default

			if len(nv.Properties) == 0 {
				nv.Properties = gen[k].Properties
			}

			if nv.Items == nil {
				nv.Items = gen[k].Items
			}

			if len(nv.OneOf) == 0 {
				nv.OneOf = gen[k].OneOf
			}

			if len(nv.Enum) == 0 {
				nv.Enum = gen[k].Enum
			}

			nv.Minimum = gen[k].Minimum

			nv.Maximum = gen[k].Maximum

			nv.MinLength = gen[k].MinLength

			nv.MaxLength = gen[k].MaxLength

			nv.MaxItems = gen[k].MaxItems

			nv.CreateOnly = gen[k].CreateOnly

			nv.Pattern = gen[k].Pattern

			nv.Example = gen[k].Example

			nv.UserError = gen[k].UserError
		}

		resultSchema[k] = nv
	}

	kg := maps.Keys(gen)

	for _, k := range kg {
		if _, ok := read[k]; !ok {
			resultSchema[k] = gen[k]
		}
	}

	return resultSchema, nil
}

// diffServiceTypes diffs the service types.
func diffServiceTypes() error {
	defer util.MeasureExecutionTime(logger)()

	schema, err := diff(genResult[types.KeyServiceTypes], readResult[types.KeyServiceTypes])
	if err != nil {
		return err
	}

	result[types.KeyServiceTypes] = schema

	return nil
}

// diffIntegrationTypes diffs the integration types.
func diffIntegrationTypes() error {
	defer util.MeasureExecutionTime(logger)()

	schema, err := diff(genResult[types.KeyIntegrationTypes], readResult[types.KeyIntegrationTypes])
	if err != nil {
		return err
	}

	result[types.KeyIntegrationTypes] = schema

	return nil
}

func diffIntegrationEndpointTypes() error {
	defer util.MeasureExecutionTime(logger)()

	schema, err := diff(genResult[types.KeyIntegrationEndpointTypes], readResult[types.KeyIntegrationEndpointTypes])
	if err != nil {
		return err
	}

	result[types.KeyIntegrationEndpointTypes] = schema

	return nil
}

// setup sets up the diff.
func setup(l *util.Logger, gr types.GenerationResult, rr types.ReadResult) {
	logger = l
	genResult = gr
	readResult = rr

	result = types.DiffResult{}
}

// Run runs the diff.
func Run(ctx context.Context, logger *util.Logger, genResult types.GenerationResult, readResult types.ReadResult) (types.DiffResult, error) {
	setup(logger, genResult, readResult)

	errs, _ := errgroup.WithContext(ctx)

	errs.Go(diffServiceTypes)
	errs.Go(diffIntegrationTypes)
	errs.Go(diffIntegrationEndpointTypes)

	err := errs.Wait()
	if err != nil {
		return nil, err
	}

	return result, nil
}
