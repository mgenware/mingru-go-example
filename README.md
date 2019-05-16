# mingru-go-example

A Go example demonstrating mingru.

## Usage

### Directory structure

- `da` data access Go code generated by mingru
- `mingru` mingru models and actions source
- The rest of the files are ordinary Go project files

### To run the example

- You need the [employees sample database](https://github.com/datacharmer/test_db)
- `dep ensure` installs dependencies
- `go run main.go` runs the example
- The example tries to connect local MySQL instance with a password of `123456`

### To set up mingru and regenerate code

- `cd mingru` navigates to the `mingru` directory
- `yarn` installs dependencies
- `yarn build` generates code to `../da` directory
