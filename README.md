
# ProofStr

ProofStr is a deterministic string generator that uses a secret key to produce unique, reproducible strings. It relies on HMAC-SHA256 to ensure that the generated strings can only be recreated by someone who knows the secret key.

## Example Usage

```bash
go run main.go --secret=mySecretKey --num=10 --length=10
```

This command will generate 10 strings, each 10 characters long, using the secret key `mySecretKey`.

## Command-Line Flags

- `--secret` (required): The secret key used to generate the strings.
- `--num` (optional): The number of strings to generate. Defaults to 5.
- `--length` (optional): The length of the generated strings. Defaults to 7.
- `--help`: Displays the help menu.

## Help

To display the help menu, run:

```bash
go run main.go --help
```

You will see:

```
Usage: ./ProofStr [options]

Options:
  --secret string      Secret key to use for generating strings (required)
  --num int            Number of strings to generate (default 5)
  --length int         Length of the generated strings (default 7)
  --help               Show this help menu
```
