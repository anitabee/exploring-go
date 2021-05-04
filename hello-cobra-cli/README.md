# Hello Cobra CLI

It's really simple CLI application example using [Cobra library](https://github.com/spf13/cobra)

## Cobra CLI example

Prerequisites: </br>
Go 1.16+ for building the binary </br>
Go 1.16+ for running the test suite </br>
Cobra v1.1+ for initializing commands </br>

## Tutorial

To start cobra project from scratch

```bash
cobra init --config .cobra.yaml --pkg-name github.com/anitabee/exploring-go/hello-cobra-cli
```

To add new command

```bash
cobra add hello --config .cobra.yaml
```

Build app

```bash
go build -o cli -ldflags="-w -s" 
```

Run cli tool

```bash
// Unix like systems
./cli hello

// Windows systems
cli.exe hello
```
