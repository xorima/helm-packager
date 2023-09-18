package config

import (
	"errors"
)

var (
	MissingChartPathError  = "missing chart path"
	MissingOutputPathError = "missing output path"
	DefaultConfig          = NewEmptyConfig() // This is the main config object we will use.
)

func NewMissingChartPathError() error {
	return errors.New(MissingChartPathError)
}

func NewMissingOutputPathError() error {
	return errors.New(MissingOutputPathError)
}

func NewEmptyConfig() *Config {
	return &Config{}
}

type Config struct {
	ChartPath  string
	OutputPath string
}

func (c *Config) GetChartPath() string {
	return c.ChartPath
}

func (c *Config) GetOutputPath() string {
	return c.OutputPath
}

func (c *Config) Validate() error {
	if c.ChartPath == "" {
		return NewMissingChartPathError()
	}
	if c.OutputPath == "" {
		return NewMissingOutputPathError()
	}
	return nil
}
