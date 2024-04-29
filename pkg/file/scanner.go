package file

import (
	"os"
	"path/filepath"
)

type Scanner struct{}

func NewScanner() *Scanner {
	return &Scanner{}
}

// 扫描指定目录，返回文件信息列表
func (s *Scanner) Scan(basePath string) ([]*FileInfo, error) {
	basePath = filepath.Clean(basePath)

	basePathInfo, err := os.Stat(basePath)
	if err != nil {
		return nil, err
	}
	if !basePathInfo.IsDir() {
		return nil, NewErrScanBasePathNotDir()
	}

	var result []*FileInfo
	if err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(basePath, path)
		if err != nil {
			return err
		}

		if relPath == "." {
			return nil
		}

		fileInfo := &FileInfo{
			BasePath: basePath,
			RelPath:  relPath,
			IsDir:    info.IsDir(),
		}

		if !fileInfo.IsDir {
			contentBytes, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			fileInfo.Content = string(contentBytes)
		}

		result = append(result, fileInfo)

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
