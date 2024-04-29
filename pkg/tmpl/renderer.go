package tmpl

import (
	"bytes"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/zhb127/easydev/pkg/file"
	"github.com/zhb127/easydev/pkg/log"
)

type Renderer struct {
	kernel      *template.Template
	fileScanner *file.Scanner
}

func NewRenderer(fileScanner *file.Scanner) *Renderer {
	kernel := newRendererKernel()
	return &Renderer{
		kernel,
		fileScanner,
	}
}

// 渲染模板文本
func (r *Renderer) RenderTmplText(tmplText string, tmplVarValues map[string]interface{}) (string, error) {
	tmplInst, err := r.kernel.Parse(tmplText)
	if err != nil {
		return "", errors.WithStack(err)
	}

	resultBuff := &bytes.Buffer{}
	if err := tmplInst.Execute(resultBuff, tmplVarValues); err != nil {
		return "", errors.WithStack(err)
	}

	result := resultBuff.String()

	return result, nil
}

// 渲染模板文件
func (r *Renderer) RenderTmplFile(tmplFilePath string, tmplVarValues map[string]interface{}) (string, error) {
	tmplFileContentBytes, err := os.ReadFile(tmplFilePath)
	if err != nil {
		return "", errors.WithStack(err)
	}

	tmplFileContent := string(tmplFileContentBytes)
	return r.RenderTmplText(tmplFileContent, tmplVarValues)
}

type RenderTmplDirParams struct {
	TmplDir       string
	TmplVarValues map[string]interface{}
	TmplFileExt   string
	OutputDir     string
}

func (r *Renderer) ValidateRenderTmplDirParams(params *RenderTmplDirParams) (*RenderTmplDirParams, error) {
	if params.TmplDir == "" {
		return nil, errors.New("模板目录不能为空")
	}
	if params.TmplFileExt == "" {
		params.TmplFileExt = DefaultTmplFileExt
	}
	if params.OutputDir == "" {
		return nil, errors.New("输出目录不能为空")
	}
	if params.TmplVarValues == nil {
		params.TmplVarValues = map[string]interface{}{}
	}

	if params.TmplFileExt != DefaultTmplFileExt {
		if matched, _ := regexp.MatchString(`^\.[a-zA-Z0-9]+$`, params.TmplFileExt); !matched {
			return nil, errors.New("模板文件后缀名格式不正确，正确格式为：.xxx")
		}
	}

	return params, nil
}

// 渲染模板目录，返回渲染后的文件信息列表
func (r *Renderer) RenderTmplDir(params *RenderTmplDirParams) ([]*file.FileInfo, error) {
	log.Debugf("开始")

	log.Debugf("扫描目录：%s", params.TmplDir)
	fileInfosScanned, err := r.fileScanner.Scan(params.TmplDir)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	log.Debugf("扫描完成")

	log.Debugf("渲染扫描到的文件信息列表")
	result, err := r.renderFileInfos(fileInfosScanned, params.TmplVarValues, params.TmplFileExt, params.OutputDir)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	log.Debugf("渲染完成")

	return result, nil
}

// 渲染文件信息列表
func (r *Renderer) renderFileInfos(
	fileInfos []*file.FileInfo,
	tmplVarValues map[string]interface{},
	tmplFileExt string,
	outputDir string,
) ([]*file.FileInfo, error) {
	var result []*file.FileInfo

	for _, v := range fileInfos {
		resultFileInfo := &file.FileInfo{
			BasePath: outputDir,
			Content:  v.Content,
			IsDir:    v.IsDir,
		}

		relPathRendered, err := r.RenderTmplText(v.RelPath, tmplVarValues)
		if err != nil {
			return nil, errors.Wrapf(err, "将文件路径文本进行模板渲染错误（文件路径：%s）", v.RelPath)
		}
		resultFileInfo.RelPath = relPathRendered

		// 判断是否模板文件
		if filepath.Ext(v.RelPath) == tmplFileExt {
			contentRendered, err := r.RenderTmplText(v.Content, tmplVarValues)
			if err != nil {
				return nil, errors.Wrapf(err, "将文件内容文本进行模板渲染错误（文件路径：%s）", v.RelPath)
			}
			resultFileInfo.Content = contentRendered
			// 移除模板文件后缀名
			resultFileInfo.RelPath = strings.TrimSuffix(resultFileInfo.RelPath, tmplFileExt)
		}

		// 更新文件路径
		resultFileInfo.Path = filepath.Join(resultFileInfo.BasePath, resultFileInfo.RelPath)

		result = append(result, resultFileInfo)
	}

	return result, nil
}
