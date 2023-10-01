package utils

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestCreateCsv(t *testing.T) {
	// t.Parallel()

	tests := []struct {
		name    string
		fPath   string
		fName   string
		want    *csv.Writer
		want1   *os.File
		wantErr bool
	}{
		{
			name:    "TestCreateCsv OK",
			fPath:   "test",
			fName:   "test",
			want:    nil,
			want1:   nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, got1, err := CreateCsv(tt, tt.fPath, tt.fName)

			if err != nil {
				t.Errorf("CreateCsv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil {
				t.Errorf("CreateCsv() got = %v, want %v", got, tt.want)
			}
			if got1 == nil {
				t.Errorf("CreateCsv() got1 = %v, want %v", got1, tt.want1)
			}

		})
	}

}
