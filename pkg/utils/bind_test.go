package utils

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func TestBodyParserAndValidate(t *testing.T) {
	type args struct {
		v    *validator.Validate
		ctx  *fiber.Ctx
		body interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := BodyParserAndValidate(tt.args.v, tt.args.ctx, tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("BodyParserAndValidate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
