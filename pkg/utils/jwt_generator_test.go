package utils

import (
	"testing"

	"github.com/auth_app/app/models"
)

func TestGenerateNewAccessToken(t *testing.T) {
	type args struct {
		data *models.User
	}
	tests := []struct {
		name    string
		args    args
		want    *string
		wantErr bool
	}{
		// Test case 1: Valid user data
		{
			name: "Valid User",
			args: args{
				data: &models.User{
					ID: 1,
				},
			},
			wantErr: false,
		},
		// Test case 2: Invalid user data (e.g., nil data)
		{
			name: "Invalid User (Nil Data)",
			args: args{
				data: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GenerateNewAccessToken(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateNewAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
