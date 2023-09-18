package files

import (
	"github.xom/xorima/helm-variant-packager/internal/models"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"regexp"
)

var (
	// doing here as it will be faster
	valuesPattern = regexp.MustCompile(`^values\.(.*).yaml$`)
)

type DirReader interface {
	ReadDir(dirname string) ([]os.FileInfo, error)
}

type FileHandler struct {
	log       *zap.SugaredLogger
	DirReader DirReader
}

func NewFileHandler(log *zap.SugaredLogger) *FileHandler {
	return &FileHandler{
		log: log,
	}
}

func (f *FileHandler) Discover(path string) ([]models.ChartVariants, error) {
	charts, err := f.discoverCharts(path)
	if err != nil {
		return nil, err
	}
	cVariants := make([]models.ChartVariants, 0)
	for _, chart := range charts {
		f.log.Infof("Discovered chart %s", chart)
		cv := make(models.ChartVariants)
		cv[models.ChartName(chart)] = make([]models.Variant, 0)
		v, err := f.discoverValuesFiles(filepath.Join(path, chart))
		if err != nil {
			return nil, err
		}
		cv[models.ChartName(chart)] = append(cv[models.ChartName(chart)], v...)
		cVariants = append(cVariants, cv)
	}
	return cVariants, nil
}

func (f *FileHandler) DeleteDir(path string) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return os.RemoveAll(path)
	})
	if err != nil {
		return err
	}
	return nil
}

func (f *FileHandler) DeleteFile(path string) error {
	return os.RemoveAll(path)
}

func (f *FileHandler) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (f *FileHandler) WriteFile(path string, content []byte) error {
	return os.WriteFile(path, content, 0644)
}

func (f *FileHandler) discoverCharts(path string) ([]string, error) {
	charts, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var resp []string
	for _, chart := range charts {
		if chart.IsDir() {
			resp = append(resp, chart.Name())
		}
	}
	return resp, nil
}

func (f *FileHandler) discoverValuesFiles(path string) ([]models.Variant, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	pattern := valuesPattern
	resp := make([]models.Variant, 0)
	for _, file := range files {
		if pattern.MatchString(file.Name()) {
			variantName := pattern.FindStringSubmatch(file.Name())[1]
			f.log.Infof("Discovered variant %s", variantName)
			v := models.Variant{FileName: file.Name(),
				Name: variantName,
			}
			resp = append(resp, v)
		}
	}
	return resp, nil
}

func (f *FileHandler) createDir(path string) error {
	if f.dirExist(path) {
		f.log.Warnf("Directory %s already exists, deleting", path)
		f.DeleteDir(path)
	}
	return os.MkdirAll(path, 0755)
}

func (f *FileHandler) dirExist(path string) bool {
	srcInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return srcInfo.IsDir()
}

func (f *FileHandler) dirNotExist(path string) bool {
	return !f.dirExist(path)
}
