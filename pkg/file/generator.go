package file

import (
	"os"

	"github.com/pkg/errors"
	"github.com/zhb127/easydev/pkg/log"
)

type Generator struct {
	fileScanner *Scanner
}

func NewGenerator(fileScanner *Scanner) *Generator {
	return &Generator{
		fileScanner,
	}
}

// 按文件信息列表生成文件，返回生成后的文件列表
func (g *Generator) GenFilesByFileInfos(fileInfos []*FileInfo, dryRun bool) ([]*FileInfo, error) {
	log.Debugf("开始")

	result := []*FileInfo{}
	for _, v := range fileInfos {
		if v.IsDir {
			log.Debugf("创建目录：%s", v.Path)
			if Exists(v.Path) {
				log.Debugf("目录已存在，跳过")
				continue
			}
			if !dryRun {
				if err := os.MkdirAll(v.Path, os.ModePerm); err != nil {
					log.Debugf(err.Error())
					return nil, errors.WithStack(err)
				}
				log.Debugf("目录创建成功")
			} else {
				log.Debugf("目录创建成功（试运行模式，不执行真正操作）")
			}
			result = append(result, v)
			continue
		}

		log.Debugf("创建文件：" + v.Path)
		if Exists(v.Path) {
			log.Debugf("文件已存在，跳过")
			continue
		}
		if !dryRun {
			if err := os.WriteFile(v.Path, []byte(v.Content), os.ModePerm); err != nil {
				log.Debugf(err.Error())
				return nil, errors.WithStack(err)
			}
			log.Debugf("文件创建成功")
		} else {
			log.Debugf("文件创建成功（试运行模式，不执行真正操作）")
		}
		result = append(result, v)
	}

	return result, nil
}
