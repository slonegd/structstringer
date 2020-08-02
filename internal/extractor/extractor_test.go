package extractor

import (
	"testing"

	"github.com/slonegd/structstringer/internal/declaration"
	"github.com/slonegd/structstringer/internal/field"
	"github.com/stretchr/testify/assert"
)

func Test_extractor_ExtractFields(t *testing.T) {
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
		{
			name:     "recurcive fields",
			files:    []string{"file1_test.go", "file2_test.go"},
			typeName: "D",
			want: field.Fields{
				{Name: "i", Type: "int"},
				{Name: "b", Type: "B", Fields: field.Fields{
					{Name: "i", Type: "int"},
					{Name: "flag", Type: "bool"},
				}},
			},
		},
	}
	for _, tt := range tests {
		finder := declaration.NewFinder(tt.files)
		decl, _ := finder.Find(tt.typeName)

		extractor := NewExtractor(finder)
		got, err := extractor.ExtractFields(decl)
		if tt.wantErr != "" {
			assert.Error(t, err)
			assert.Equal(t, tt.wantErr, err.Error())
			continue
		}

		assert.NoError(t, err)
		assert.Equal(t, tt.want, got)
	}
}
