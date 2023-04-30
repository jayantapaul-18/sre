# Development steps & examples

```bash
go mod init github.com/jayantapaul-18/sre
go mod tidy
go install github.com/spf13/cobra-cli@latest
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
cobra init --pkg-name github.com/jayantapaul-18/sre
go mod tidy
go run .
go build .
```

# Local build & Run

```bash
Run below command to fixed Makefile if your build failing with error as "Makefile:7: \*\*\* missing separator. Stop."
perl -pi -e 's/^ \*/\t/' Makefile
To View & validate:
cat -e -t -v Makefile
```

```bash
./build.sh
OR,
make build
make build-with-config
make run
cz bump
```

## Local

brew install pre-commit

# Local Command to validate diffrrent GO packages

```bash
go env
ls -la $GOPATH/bin
ls $GOPATH/bin/cobra
which go
/usr/local/go/bin/go
echo `go env GOPATH`/bin
ls $GOPATH/bin/cobra
go version -m ~/go/bin/goimports
go build -o sre -ldflags="-X 'main.Version=v0.0.1'"
```
