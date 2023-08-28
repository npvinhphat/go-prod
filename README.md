# go-prod

[![GitHub License](https://img.shields.io/github/license/npvinhphat/go-prod)](https://github.com/npvinhphat/go-prod/blob/main/LICENSE)

**go-prod** is a command-line tool written in Go that facilitates the management
of service production readiness. It helps you generate production readiness
guidelines and checklists based on various technology stacks.

## Installation

To install go-prod, you can use the `go install` command:

```sh
go install github.com/npvinhphat/go-prod
```

## Usage

### Generating Production Readiness Guideline

Use the following command to generate a production readiness guideline:

```sh
go-prod generate guideline --stacks=default,kubernetes,pagerduty,slack,wavefront > docs/examples/guideline.md
```

Replace the `--stacks` flag value with the desired technology stacks.

### Generating Production Readiness Checklist

Use the following command to generate a production readiness checklist:

```sh
go-prod generate checklist --stacks=default,kubernetes,pagerduty,slack,wavefront --level=a > docs/examples/checklist.md
```

Replace the `--stacks` flag value with the desired technology stacks and
`--level` with the desired readiness level.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for
improvements, feel free to open an issue or a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
