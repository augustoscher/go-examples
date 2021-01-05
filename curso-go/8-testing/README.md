# Running tests

## Inside directory

- Enter on test folder:

```bash
cd go-examples/curso-go/8-testing/mathematics
```

- Run:

```bash
go test
```

## All tests from project

- Enter on project folder:

```bash
cd go-examples/curso-go/8-testing/
```

- Run:

```bash
go test ./...
```

## With verbose mode

- Enter on project folder:

```bash
cd go-examples/curso-go/8-testing/dataset
```

- Run:

```bash
go test -v
```

Result:

```bash
=== RUN   TestIndex
    TestIndex: strings_test.go:29: Test: {Testing is greate Testing 0}
    TestIndex: strings_test.go:29: Test: {  0}
    TestIndex: strings_test.go:29: Test: {Hey hey -1}
    TestIndex: strings_test.go:29: Test: {Hi, how are you? how 4}
    TestIndex: strings_test.go:29: Test: {augusto g 2}
--- PASS: TestIndex (0.00s)
PASS
```

## Coverage

To get code coverage, the test needs to be on GOPATH folder

- Enter on folder:

```bash
cd /home/AUGUSTO.SCHER/go/src/github.com/augustoscher/math
```

- Run:

```bash
go test -cover
```

```bash
go test -coverprofile=result.out

go tool cover -func=result.out

go tool cover -html=result.out
```
