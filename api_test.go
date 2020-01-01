package aiven

import "testing"

func Test_checkAPIResponse(t *testing.T) {
	type args struct {
		bts []byte
		r   Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"invalid-json",
			args{
				bts: []byte(`Invalid JSON`),
				r:   nil,
			},
			true,
		},
		{
			"error-response",
			args{
				bts: []byte(`{
					"message": "Authentication failed",
					"errors": [
						{
							"status": "403",
							"message": "Authentication failed"
						}
					]
				}`),
				r: nil,
			},
			true,
		},
		{
			"error-response",
			args{
				bts: []byte(`{
					"message": "",
					"errors": [
					]
				}`),
				r: nil,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkAPIResponse(tt.args.bts, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("checkAPIResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAPIResponse_GetError(t *testing.T) {
	type fields struct {
		Errors  []Error
		Message string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"empty",
			fields{
				Errors:  nil,
				Message: "",
			},
			false,
		},
		{
			"has-error",
			fields{
				Errors: []Error{
					{
						Message:  "error-message",
						MoreInfo: "some-info",
						Status:   500,
					},
				},
				Message: "error-message",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := APIResponse{
				Errors:  tt.fields.Errors,
				Message: tt.fields.Message,
			}
			if err := r.GetError(); (err != nil) != tt.wantErr {
				t.Errorf("GetError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
