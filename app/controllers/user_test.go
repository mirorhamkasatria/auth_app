package controllers

import (
	"testing"

	"github.com/auth_app/app/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func Test_controllerUser_LoginUser(t *testing.T) {
	mockService := new(services.MockServiceUser)
	mockValidator := validator.New()
	type fields struct {
		sv        services.ServiceUser
		validator *validator.Validate
	}
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "valid_login",
			fields: fields{
				sv:        mockService,
				validator: mockValidator,
			},
			args:    args{ctx: createTestContext("/login", "POST")},
			wantErr: false,
		},
		{
			name: "invalid_login",
			fields: fields{
				sv:        mockService,
				validator: mockValidator,
			},
			args:    args{ctx: createTestContext("/login", "POST")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controllerUser{
				sv:        tt.fields.sv,
				validator: tt.fields.validator,
			}
			err := c.LoginUser(tt.args.ctx)

			if tt.wantErr {
				// Check the status code
				if tt.args.ctx.Response().StatusCode() == fiber.StatusOK {
					t.Errorf("Expected status code %d, but got %d", fiber.StatusOK, tt.args.ctx.Response().StatusCode())
				}

			} else {
				if err != nil {
					t.Errorf("Expected no error, but got %v", err)
				}
			}
		})
	}
}

func Test_controllerUser_RegisterUser(t *testing.T) {
	mockService := new(services.MockServiceUser)
	mockValidator := validator.New()

	// Helper function to create a test context
	// createTestContext := createTestContext(path, method)

	type fields struct {
		sv        services.ServiceUser
		validator *validator.Validate
	}
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "register",
			fields: fields{
				sv:        mockService,
				validator: mockValidator,
			},
			args:    args{ctx: createTestContext("/register", "POST")},
			wantErr: true,
		},
		{
			name: "successful_registration",
			fields: fields{
				sv:        mockService,
				validator: mockValidator,
			},
			args:    args{ctx: createTestContext("/register", "POST")},
			wantErr: false,
		},
		// Add a scenario for failed registration
		{
			name: "failed_registration",
			fields: fields{
				sv:        mockService,
				validator: mockValidator,
			},
			args:    args{ctx: createTestContext("/register", "POST")},
			wantErr: true,
		},
		// Add other test cases...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controllerUser{
				sv:        tt.fields.sv,
				validator: tt.fields.validator,
			}
			err := c.RegisterUser(tt.args.ctx)

			if tt.wantErr {
				// Check the status code
				if tt.args.ctx.Response().StatusCode() == fiber.StatusOK {
					t.Errorf("Expected status code %d, but got %d", fiber.StatusOK, tt.args.ctx.Response().StatusCode())
				}

			} else {
				if err != nil {
					t.Errorf("Expected no error, but got %v", err)
				}
			}
		})
	}
}

func createTestContext(path string, method string) *fiber.Ctx {

	app := fiber.New()

	// Use fasthttp.RequestCtx
	ctx := app.AcquireCtx(new(fasthttp.RequestCtx))

	// Populate the RequestCtx with the desired path and method
	ctx.Request().SetRequestURI(path)
	ctx.Request().Header.SetMethod(method)

	return ctx

}
