package utils

import "testing"

func TestCopier(t *testing.T) {
	type args struct {
		to   interface{}
		from interface{}
	}
	type PersonFrom struct {
		Name string
		Age  int
	}

	type PersonTo struct {
		Name string
		Age  int
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// Test case 1: Copying valid data
		{
			name:    "Copy Valid Data",
			args:    args{to: &PersonTo{}, from: PersonFrom{Name: "John", Age: 25}},
			wantErr: false,
		},
		// Test case 2: Copying invalid data (e.g., different types)
		{
			name:    "Copy Invalid Data",
			args:    args{to: &PersonTo{}, from: "Invalid Data"},
			wantErr: true,
		},
		// Test case 3: Copying with nil destination
		{
			name:    "Copy to Nil Destination",
			args:    args{to: nil, from: PersonFrom{Name: "Alice", Age: 30}},
			wantErr: true,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Copier(tt.args.to, tt.args.from); err != nil {
				t.Errorf("Copier() error = %v", err)
			}

		})
	}
}
