package declaration

import (
	"testing"

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
		finder := NewFinder(tt.files)
		_, err := finder.Find(tt.typeName)
		if tt.wantErr != "" {
			assert.Error(t, err)
			assert.Equal(t, tt.wantErr, err.Error())
			continue
		}

		assert.NoError(t, err)
	}
}
