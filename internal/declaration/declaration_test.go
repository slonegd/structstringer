package declaration

import (
	"testing"

	"github.com/slonegd/structstringer/internal/field"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	tests := []struct {
		name     string
		files    []string
		typeName string
		wantErr  string
	}{
		{
			name:     "no file",
			files:    []string{"no_file.go"},
			typeName: "A",
			wantErr:  `parse file: open no_file.go: no such file or directory`,
		},
		{
			name:     "no type",
			files:    []string{"file2_test.go"},
			typeName: "A",
			wantErr:  `cant find type "A"`,
		},
		{
			name:     "happy path",
			files:    []string{"file1_test.go"},
			typeName: "A",
		},
	}
	for _, tt := range tests {
		_, err := Find(tt.files, tt.typeName)
		if tt.wantErr != "" {
			assert.Error(t, err)
			assert.Equal(t, tt.wantErr, err.Error())
			continue
		}

		assert.NoError(t, err)
	}
}

func TestExtractFields(t *testing.T) {
	tests := []struct {
		name     string
		files    []string
		typeName string
		want     field.Fields
		wantErr  string
	}{
		{
			name:     "happy path",
			files:    []string{"file1_test.go", "file2_test.go"},
			typeName: "A",
			want: field.Fields{
				{Name: "i", Type: "int"},
				{Name: "flag", Type: "bool"},
			},
		},
		{
			name:     "not a struct",
			files:    []string{"file1_test.go", "file2_test.go"},
			typeName: "C",
			wantErr:  `type "C" not a struct`,
		},
	}
	for _, tt := range tests {
		gotDecl, _ := Find(tt.files, tt.typeName)
		got, err := ExtractFields(gotDecl)
		if tt.wantErr != "" {
			assert.Error(t, err)
			assert.Equal(t, tt.wantErr, err.Error())
			continue
		}

		assert.NoError(t, err)
		assert.Equal(t, tt.want, got)
	}
}
