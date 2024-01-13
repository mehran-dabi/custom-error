package customerror

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testData struct {
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"min=18"`
	Email string `json:"email" validate:"required,email"`
}

func TestErrorValidator(t *testing.T) {
	validate := validator.New()
	testCases := []struct {
		desc     string
		input    testData
		expected map[string]string
	}{
		{
			desc:     "missing required fields",
			input:    testData{},
			expected: map[string]string{"name": "is required", "age": "must be greater than 18", "email": "is required"},
		},
		{
			desc:     "invalid age",
			input:    testData{Name: "John Doe", Age: 17, Email: "john@example.com"},
			expected: map[string]string{"age": "must be greater than 18"},
		},
		{
			desc:     "valid input",
			input:    testData{Name: "John Doe", Age: 30, Email: "john@example.com"},
			expected: map[string]string{},
		},
		{
			desc:     "invalid email",
			input:    testData{Name: "John Doe", Age: 30, Email: "invalid_email"},
			expected: map[string]string{"email": "validation failed on 'email' condition"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := validate.Struct(tc.input)
			if err != nil {
				got := ErrorValidator(err, tc.input)
				assert.Equal(t, tc.expected, got)
			} else {
				assert.Empty(t, tc.expected)
			}
		})
	}

}
