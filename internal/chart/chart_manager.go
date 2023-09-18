package chart

import (
	"github.xom/xorima/helm-variant-packager/internal/config"
	"github.xom/xorima/helm-variant-packager/internal/files"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"path/filepath"
)

type helmChart struct {
	Name         string                `yaml:"name"`
	ApiVersion   string                `yaml:"apiVersion"`
	Description  string                `yaml:"description"`
	Type         string                `yaml:"type"`
	Version      string                `yaml:"version"`
	AppVersion   string                `yaml:"appVersion"`
	Dependencies []helmChartDependency `yaml:"dependencies"`
}

// helmChartDependency is a subset of the helmChart struct
// to handle the dependencies section
// note that Version is not required and in our current
// usecase will not be used
type helmChartDependency struct {
	Name       string `yaml:"name"`
	Repository string `yaml:"repository"`
}

type Handler struct {
	log       *zap.SugaredLogger
	cfg       *config.Config
	files     *files.FileHandler
	chartName string
}

func NewChartHandler(log *zap.SugaredLogger, cfg *config.Config, handler *files.FileHandler, chartName string) *Handler {
	return &Handler{chartName: chartName, files: handler, log: log, cfg: cfg}
}

func (h *Handler) Handle() error {
	h.log.Infof("Discovering charts in %s", h.cfg.GetChartPath())
	charts, err := h.files.Discover(h.cfg.GetChartPath())
	if err != nil {
		h.log.Fatalf("Failed to discover charts: %s", err)
	}
	err = h.processCharts(charts)
	if err != nil {
		h.log.Fatalf("Failed to process charts: %s", err)
	}
	return nil
}

func (h *Handler) getChartPath() (string, error) {
	absPath, err := filepath.Abs(filepath.Join(h.cfg.GetChartPath(), h.chartName))
	if err != nil {
		return "", err
	}
	return absPath, nil
}

func (h *Handler) indentYaml(input []byte, indentTitle string) ([]byte, error) {
	var content interface{}
	err := yaml.Unmarshal(input, &content)
	if err != nil {
		return nil, err
	}

	c := map[string]any{
		indentTitle: content,
	}
	return yaml.Marshal(c)
}

func (h *Handler) updateChartYamlName(input []byte, name string) ([]byte, error) {
	var chart helmChart
	err := yaml.Unmarshal(input, &chart)
	if err != nil {
		return nil, err
	}

	chart.Name = name
	return yaml.Marshal(chart)
}
