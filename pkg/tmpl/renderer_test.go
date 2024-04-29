package tmpl

import (
	"reflect"
	"testing"
)

func TestNewRenderer(t *testing.T) {
	tests := []struct {
		name string
		want *Renderer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRenderer(nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRenderer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRenderer_RenderTmplText(t *testing.T) {
	type args struct {
		tmplText      string
		tmplVarValues map[string]interface{}
	}
	tests := []struct {
		name    string
		r       *Renderer
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "",
			r: func() *Renderer {
				return NewRenderer(nil)
			}(),
			args: args{
				tmplText: "Hello, {{.Name}}",
				tmplVarValues: map[string]interface{}{
					"Name": "World",
				},
			},
			want:    "Hello, World",
			wantErr: false,
		},
		{
			name: "",
			r: func() *Renderer {
				return NewRenderer(nil)
			}(),
			args: args{
				tmplText: `
{{"ApiToken" | ToSnakeCase}}
{{"api token" | ToUpperCamelCase}}
{{.ResourceName | ToLowerCamelCase}}
{{.ResourceName | ToKebabCase}}
{{.ResourceName | ToPlural}}
{{.ResourceName | ToSingular}}
{{.ResourceName | ToLowerCamelCase | ToPlural}}
`,
				tmplVarValues: map[string]interface{}{
					"ResourceName": "api_token",
				},
			},
			want: `
api_token
ApiToken
apiToken
api-token
api_tokens
api_token
apiTokens
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.r
			got, err := r.RenderTmplText(tt.args.tmplText, tt.args.tmplVarValues)
			if (err != nil) != tt.wantErr {
				t.Errorf("Renderer.RenderTmplText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Renderer.RenderTmplText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRenderer_RenderTmplFile(t *testing.T) {
	type args struct {
		tmplFilePath  string
		tmplVarValues map[string]interface{}
	}
	tests := []struct {
		name    string
		r       *Renderer
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "",
			r: func() *Renderer {
				return NewRenderer(nil)
			}(),
			args: args{
				tmplFilePath: "./test/data/demo1/hello.txt.tmpl",
				tmplVarValues: map[string]interface{}{
					"Name": "World",
				},
			},
			want:    "Hello, World",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.r
			got, err := r.RenderTmplFile(tt.args.tmplFilePath, tt.args.tmplVarValues)
			if (err != nil) != tt.wantErr {
				t.Errorf("Renderer.RenderFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Renderer.RenderFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
