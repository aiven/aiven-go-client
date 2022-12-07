package aiven

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

// TestAccessControlMarshalJSON test AccessControl JSON marshalling process
func TestAccessControlMarshalJSON(t *testing.T) {
	strPtr := func(s string) *string {
		return &s
	}
	type notEmpty struct {
		M3Group      string   `json:"m3_group"`
		RedisACLKeys []string `json:"redis_acl_keys"`
	}

	tests := []struct {
		name     string
		input    *AccessControl
		wantType interface{}
		want     interface{}
	}{
		{
			name:     "empty",
			input:    &AccessControl{},
			wantType: &map[string]interface{}{},
			want:     &map[string]interface{}{},
		},
		{
			name: "notempty",
			input: &AccessControl{
				M3Group:      strPtr("foo"),
				RedisACLKeys: []string{"bar"},
			},
			wantType: &notEmpty{},
			want: &notEmpty{
				M3Group:      "foo",
				RedisACLKeys: []string{"bar"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.Marshal(tt.input)
			if err != nil {
				t.Errorf("MarshalJSON() error = %v", err)
				return
			}

			dec := json.NewDecoder(bytes.NewReader(b))
			dec.DisallowUnknownFields()
			if err := dec.Decode(tt.wantType); err != nil {
				t.Errorf("Decode() error = %v", err)
			}
			if !reflect.DeepEqual(tt.wantType, tt.want) {
				t.Errorf("DeepEqual() got = %+v, want = %+v", tt.wantType, tt.want)
			}
		})
	}
}
