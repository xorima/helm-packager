package files

import (
	"github.com/stretchr/testify/assert"
	"github.com/youshy/logger"
	"github.xom/xorima/helm-variant-packager/internal/models"
	"testing"
)

func TestFileHandler_Discover(t *testing.T) {
	t.Run("returns a list of chart variants", func(t *testing.T) {
		h := NewFileHandler(logger.NewLogger(logger.FATAL, false))
		want := make(models.ChartVariants)
		want[models.ChartName("chart-1")] = []models.Variant{
			{
				FileName: "values.dev.yaml",
				Name:     "dev",
			},
			{
				FileName: "values.prod.yaml",
				Name:     "prod",
			}}
		got, err := h.Discover("testdata")
		assert.NoError(t, err)
		assert.Equal(t, want, got[0])
	})
}
