package exp

import (
	"reflect"
	"testing"

	"github.com/aiven/aiven-go-client"
)

// TestExp_filterServiceTypes is a test for filterServiceTypes.
func TestExp_filterServiceTypes(t *testing.T) {
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
			got, err := filterServiceTypes(tt.args.f)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("filterServiceTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterServiceTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}
