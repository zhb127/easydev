package file

import (
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewScanner(t *testing.T) {
	tests := []struct {
		name string
		want *Scanner
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScanner(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScanner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_Scan(t *testing.T) {
	type args struct {
		basePath string
	}
	tests := []struct {
		name    string
		fs      *Scanner
		args    args
		want    []*FileInfo
		wantErr bool
	}{

		{
			name: "",
			fs:   &Scanner{},
			args: args{
				basePath: "./test/data/dir1/hello.txt.tmpl",
			},
			want:    []*FileInfo{},
			wantErr: false,
		},
		{
			name: "",
			fs:   &Scanner{},
			args: args{
				basePath: "./test/data/dir1/",
			},
			want: []*FileInfo{
				{BasePath: "test/data/dir1", RelPath: "biz", IsDir: true},
				{BasePath: "test/data/dir1", RelPath: "biz/{{.ResourceNameLSC}}.go.tmpl", IsDir: false, Content: `package biz

type {{.ResourceNameUCC}}Biz interface {
}
`},
				{BasePath: "test/data/dir1", RelPath: "dao", IsDir: true},
				{BasePath: "test/data/dir1", RelPath: "dao/mysql", IsDir: true},
				{BasePath: "test/data/dir1", RelPath: "dao/mysql/{{.ResourceNameLSC}}.go.tmpl", IsDir: false, Content: `package mysqldao

type {{.ResourceNameUCC}}Dao interface {
}
`},
				{BasePath: "test/data/dir1", RelPath: "hello.txt.tmpl", IsDir: false, Content: "Hello, {{.Name}}"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &Scanner{}
			got, err := fs.Scan(tt.args.basePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scanner.Scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("Scanner.Scan() = %v, want %v, diff %v", got, tt.want, diff)
			}
		})
	}
}
