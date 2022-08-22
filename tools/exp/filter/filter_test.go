package filter

import (
	"reflect"
	"testing"

	"github.com/aiven/aiven-go-client"
)

// TestServiceTypes is a test for ServiceTypes.
func TestServiceTypes(t *testing.T) {
	type args struct {
		f map[string]aiven.ServiceType
	}

	tests := []struct {
		name    string
		args    args
		want    map[string]aiven.ServiceType
		wantErr error
	}{
		{
			name: "basic",
			args: args{
				f: map[string]aiven.ServiceType{
					"foo":   {},
					"kafka": {},
				},
			},
			want: map[string]aiven.ServiceType{
				"kafka": {},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ServiceTypes(tt.args.f)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("ServiceTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServiceTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestIntegrationTypes is a test for IntegrationTypes.
func TestIntegrationTypes(t *testing.T) {
	type args struct {
		f []aiven.IntegrationType
	}

	tests := []struct {
		name    string
		args    args
		want    []aiven.IntegrationType
		wantErr error
	}{
		{
			name: "basic",
			args: args{
				f: []aiven.IntegrationType{
					{
						DestServiceTypes:   []string{"pg"},
						SourceServiceTypes: []string{"kafka"},
					},
					{
						DestServiceTypes:   []string{"flink", "foo"},
						SourceServiceTypes: []string{"kafka"},
					},
				},
			},
			want: []aiven.IntegrationType{
				{
					DestServiceTypes:   []string{"pg"},
					SourceServiceTypes: []string{"kafka"},
				},
				{
					DestServiceTypes:   []string{"flink"},
					SourceServiceTypes: []string{"kafka"},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntegrationTypes(tt.args.f)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("IntegrationTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntegrationTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestIntegrationEndpointTypes is a test for IntegrationEndpointTypes.
func TestIntegrationEndpointTypes(t *testing.T) {
	type args struct {
		f []aiven.IntegrationEndpointType
	}

	tests := []struct {
		name    string
		args    args
		want    []aiven.IntegrationEndpointType
		wantErr error
	}{
		{
			name: "basic",
			args: args{
				f: []aiven.IntegrationEndpointType{
					{
						ServiceTypes: []string{
							"kafka",
							"flink",
						},
					},
					{
						ServiceTypes: []string{
							"pg",
							"foo",
						},
					},
				},
			},
			want: []aiven.IntegrationEndpointType{
				{
					ServiceTypes: []string{
						"kafka",
						"flink",
					},
				},
				{
					ServiceTypes: []string{
						"pg",
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntegrationEndpointTypes(tt.args.f)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("IntegrationEndpointTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntegrationEndpointTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}
