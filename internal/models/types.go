package models

type (
	ChartName string
	Variant   struct {
		FileName string
		Name     string
	}
	ChartVariants map[ChartName][]Variant

	HelmChart struct {
		Name         string                `yaml:"name"`
		ApiVersion   string                `yaml:"apiVersion"`
		Description  string                `yaml:"description"`
		Type         string                `yaml:"type"`
		Version      string                `yaml:"version"`
		AppVersion   string                `yaml:"appVersion"`
		Dependencies []HelmChartDependency `yaml:"dependencies"`
	}
	HelmChartDependency struct {
		Name       string `yaml:"name"`
		Version    string `yaml:"version"`
		Repository string `yaml:"repository"`
	}
)
