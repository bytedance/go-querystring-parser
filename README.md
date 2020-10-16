# go-querystring-parser
A golang querystring parser

## Usage

```go
import (
    querystring "github.com/bytedance/go-querystring-parser"
)
```

Then use:

```go
query := "message: test\\ value AND datetime: [\"2020-01-01T00:00:00\" TO \"2020-12-31T00:00:00\"]"
ast, err := querystring.Parse(query)
if err != nil {
    // error handling
}

// do something with ast
```

## For Developers

After edit querystring.y, gen code via run:

```shell
go generate
```
