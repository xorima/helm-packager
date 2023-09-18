package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMissingChartPathError(t *testing.T) {
	err := NewMissingChartPathError()
	assert.Equal(t, MissingChartPathError, err.Error(), "they should be equal")
}

func TestNewMissingOutputPathError(t *testing.T) {
	err := NewMissingOutputPathError()
	assert.Equal(t, MissingOutputPathError, err.Error(), "they should be equal")
}

func TestGetChartPath(t *testing.T) {
	chartPath := "testChartPath"
	config := Config{ChartPath: chartPath}
	assert.Equal(t, chartPath, config.GetChartPath(), "they should be equal")
}

func TestGetOutputPath(t *testing.T) {
	outputPath := "testOutputPath"
	config := Config{OutputPath: outputPath}
	assert.Equal(t, outputPath, config.GetOutputPath(), "they should be equal")
}

func TestValidate(t *testing.T) {
	// Test missing chart path
	config := Config{}
	err := config.Validate()
	assert.EqualError(t, err, MissingChartPathError)

	// Test missing output path
	chartPath := "testChartPath"
	config = Config{ChartPath: chartPath}
	err = config.Validate()
	assert.EqualError(t, err, MissingOutputPathError)

	// Test valid config
	outputPath := "testOutputPath"
	config = Config{ChartPath: chartPath, OutputPath: outputPath}
	err = config.Validate()
	assert.Nil(t, err)
}
