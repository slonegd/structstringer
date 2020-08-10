package extractor

import (
	"testing"

	"github.com/slonegd/structstringer/internal/declaration"
	"github.com/slonegd/structstringer/internal/field"
	"github.com/stretchr/testify/assert"
)

func Test_extractor_ExtractFields(t *testing.T) {
	tests := []struct {
		name        string
		files       []string
		typeName    string
		packageName string
		want        field.Fields
		wantErr     string
	}{
		{
			name:        "happy path",
			files:       []string{"file1_test.go", "file2_test.go"},
			typeName:    "A",
			packageName: "extractor",
			want: field.Fields{
				{Name: "i", PathToValue: "i", Type: "int", Package: "extractor"},
				{Name: "flag", PathToValue: "flag", Type: "bool", Package: "extractor"},
			},
		},
		{
			name:        "not a struct",
			files:       []string{"file1_test.go", "file2_test.go"},
			typeName:    "C",
			packageName: "extractor",
			wantErr:     `type "C" not a struct`,
		},
		{
			name:        "recurcive fields",
			files:       []string{"file1_test.go", "file2_test.go"},
			typeName:    "D",
			packageName: "extractor",
			want: field.Fields{
				{Name: "i", PathToValue: "i", Type: "int", Package: "extractor"},
				{Name: "b", PathToValue: "b", Type: "B", Package: "extractor", Fields: field.Fields{
					{Name: "i", PathToValue: "b.i", Type: "int", Package: "extractor"},
					{Name: "flag", PathToValue: "b.flag", Type: "bool", Package: "extractor"},
				}},
			},
		},
		{
			name:        "other package",
			files:       []string{"file1_test.go", "file2_test.go"},
			typeName:    "E",
			packageName: "extractor",
			want: field.Fields{
				{Name: "i", PathToValue: "i", Type: "int", Package: "extractor"},
				{Name: "b", PathToValue: "b", Type: "B", Package: "simple", IsOtherPackage: true, Fields: field.Fields{
					{Name: "I", PathToValue: "b.I", Type: "int", Package: "simple"},
				}},
			},
		},
	}
	for _, tt := range tests {
		finder := declaration.NewFinder(tt.files)
		decl, _ := finder.Find(tt.typeName, "")

		extractor := NewExtractor(finder, tt.packageName)
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
