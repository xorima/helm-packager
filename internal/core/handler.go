package core

import (
	"fmt"
	"github.xom/xorima/helm-variant-packager/internal/config"
	"github.xom/xorima/helm-variant-packager/internal/files"
	"github.xom/xorima/helm-variant-packager/internal/models"
	"go.uber.org/zap"
	"path/filepath"
)

const valuesYaml = "values.yaml"

type Handler struct {
	log   *zap.SugaredLogger
	cfg   *config.Config
	files *files.FileHandler
}

func NewHandler(log *zap.SugaredLogger, cfg *config.Config) *Handler {
	return &Handler{
		log:   log,
		cfg:   cfg,
		files: files.NewFileHandler(log),
	}
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

func (h *Handler) processCharts(charts []models.ChartVariants) error {
	for _, chart := range charts {
		// helm package root chart

		for c, v := range chart {
			h.log.Infof("Found chart %s", c)
			for _, variant := range v {
				// need to work out the absolute path (path, _ := filepath.Abs(dir)) for the chart so we can use it in dependancies
				// then helm dependency update
				// then helm package

				// note we cannot use a packaged chart as a dependancy
				// it has to be a chart directory

				h.log.Infof("Found variant %s", variant.Name)
				dstPath := h.variantPath(string(c), variant.Name)
				// we only need to copy 2 files, the Chart.yaml and the values.yaml
				h.files.createDir(dstPath)
				err := h.setupChartYaml(string(c), variant.Name)
				if err != nil {
					return err
				}
				err = h.setupChartValues(string(c), variant)
				if err != nil {
					return err
				}
				// delete variant.values
				// helm package root chart
				// add in dependancies
				// helm package each variant with dependancies
			}
		}
	}
	return nil

}

func (h *Handler) variantName(chart, variant string) string {
	return fmt.Sprintf("%s-%s", chart, variant)
}

func (h *Handler) variantPath(chart, variant string) string {
	return filepath.Join(h.cfg.GetOutputPath(), h.variantName(chart, variant))
}

func (h *Handler) setupChartYaml(chart, variantName string) error {
	filepath.Join(h.cfg.GetChartPath(), chart)
	chartAbsPath, err := filepath.Abs(filepath.Join(h.cfg.GetChartPath(), chart))
	if err != nil {
		return err
	}

	content, err := h.files.ReadFile(filepath.Join(chartAbsPath, "Chart.yaml"))
	if err != nil {
		return err
	}
	updatedContent, err := UpdateChartName(content, h.variantName(chart, variantName))
	if err != nil {
		return err
	}
	return h.files.WriteFile(filepath.Join(h.variantPath(chart, variantName), "Chart.yaml"), updatedContent)
}

func (h *Handler) setupChartValues(chart string, variant variant) error {
	content, err := h.files.ReadFile(filepath.Join(h.cfg.GetChartPath(), chart, variant.fileName))
	if err != nil {
		return err
	}
	updatedContent, err := IndentYaml(content, chart)
	if err != nil {
		return err
	}
	return h.files.WriteFile(filepath.Join(h.variantPath(chart, variant.name), "values.yaml"), updatedContent)

}
