# GoExchangeCoins

A small command-line currency exchange tool written in Go. It reads a value and
two currencies from standard input and prints the converted amount, looping until
you type `close`.

Supported currencies and their rates (relative to USD) live in `cotations` in
[`main.go`](main.go):

| Code | Rate |
|------|------|
| USD  | 1.00 |
| BRL  | 5.10 |
| EUR  | 0.87 |

Conversion is `value * rate(target) / rate(origin)`.

## Requirements

- [Go](https://go.dev/dl/) 1.26 or newer (see [`go.mod`](go.mod)).

## Project structure

```Diagram
GoExchangeCoins/
├── go.mod            # Module definition and Go version
├── main.go           # Entry point, the I/O loop (run), and the convert logic
├── convert_test.go   # Table-driven unit tests for convert()
├── main_test.go      # Tests for the run() loop using in-memory I/O
└── README.md
```

A few notes on the design:

- **`convert(value, origin, target)`** holds the pure conversion logic with no
  I/O, which makes it trivial to unit-test.
- **`run(in io.Reader, out io.Writer)`** contains the interactive loop. It takes
  its input and output as parameters instead of using `os.Stdin`/`os.Stdout`
  directly, so tests can feed it a string and capture its output.
- **`main()`** is a thin wrapper that wires the real stdin/stdout into `run`.

## Build

Compile a binary into the current directory:

```sh
go build -o goexchangecoins .
```

Then run it:

```sh
./goexchangecoins
```

## Run

To run directly without producing a binary:

```sh
go run .
```

### Usage

The program prompts you for three things in order, then prints the result:

```bash
Type the value you want to exchange ->
2
Type the origin currency ->
USD
Type the target currency ->
BRL
Result: BRL 10.20
```

Type `close` at any prompt to exit the program.

## Test

Run the full test suite:

```sh
go test ./...
```

For verbose output showing each individual case:

```sh
go test -v ./...
```

Run a single test by name:

```sh
go test -run TestConvert -v ./...
```

With coverage:

```sh
go test -cover ./...
```
