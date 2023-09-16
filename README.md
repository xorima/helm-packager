# helm-variant-packager

A Helm Plugin for Chart Packaging to handle variant/values overrides for you, to create a separate chart for each values
file that exists so when releasing you do not need to be concerned about remembering the right values file to use.

Note: This repository follows semantic versioning.

## Usage

### Install

```bash
    helm plugin install https://github.com/xorima/helm-variant-packager .git --version v1.0.0
```

### Package

```bash
    helm variant-packager package <chart-path> <output-path>
```

### Push

```bash
    helm variant-packager push <chart-path> <output-path> <repository-url>
```

### help

```bash
    helm variant-packager help
```

## Development

### Build

```bash
    make build
```

### Test

```bash
    make test
```

### Install

```bash
    make install
```

### Uninstall

```bash
    make uninstall
```
