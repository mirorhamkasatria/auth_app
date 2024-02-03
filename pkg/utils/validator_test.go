package utils

import (
	"testing"

	"github.com/auth_app/app/transports/requests"
	"github.com/go-playground/validator/v10"
)

func TestValidateStruct(t *testing.T) {
	type args struct {
		v    *validator.Validate
		data interface{}
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// Test case 1: Valid user data
		{
			name: "Valid User",
			args: args{
				v:    validator.New(),
				data: &requests.Login{Password: "john_doe", Email: "john@example.com"},
			},
			wantErr: false,
		},
		// Test case 2: Invalid user data (missing required field)
		{
			name: "Invalid User (Missing Field)",
			args: args{
				v:    validator.New(),
				data: &requests.Login{Email: "john@example.com"},
			},
			wantErr: true,
		},
		// Test case 3: Invalid user data (email format)
		{
			name: "Invalid User (Invalid Email)",
			args: args{
				v:    validator.New(),
				data: &requests.Login{Password: "john_doe", Email: "invalid-email"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateStruct(tt.args.v, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ValidateStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
