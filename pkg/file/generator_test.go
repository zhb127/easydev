package file

import (
	"reflect"
	"testing"
)

func TestNewGenerator(t *testing.T) {
	type args struct {
		fileScanner *Scanner
	}
	tests := []struct {
		name string
		args args
		want *Generator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGenerator(tt.args.fileScanner); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_GenFilesByFileInfos(t *testing.T) {
	type fields struct {
		fileScanner *Scanner
	}
	type args struct {
		fileInfos []*FileInfo
		dryRun    bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*FileInfo
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				fileScanner: nil,
			},
			args: args{
				fileInfos: []*FileInfo{
					{
						Path:     "",
						BasePath: "",
						RelPath:  "",
						Content:  "",
						IsDir:    false,
					},
				},
				dryRun: true,
			},
			want: []*FileInfo{
				{
					Path:     "",
					BasePath: "",
					RelPath:  "",
					Content:  "",
					IsDir:    false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{
				fileScanner: tt.fields.fileScanner,
			}
			got, err := g.GenFilesByFileInfos(tt.args.fileInfos, tt.args.dryRun)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generator.GenFilesByFileInfos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generator.GenFilesByFileInfos() = %v, want %v", got, tt.want)
			}
		})
	}
}
