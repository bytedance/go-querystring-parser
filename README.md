# go-querystring-parser
A golang Elasticsearch Querystring parser

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

The AST parsed from above querystring:

```
// print via github.com/davecgh/go-spew.Dump
(*querystring.AndCondition)(0xc00000c200)({
    Left: (*querystring.MatchCondition)(0xc00000c1e0)({
        Field: (string) (len=7) "message",
        Value: (string) (len=10) "test value"
    }),
    Right: (*querystring.TimeRangeCondition)(0xc000076750)({
        Field: (string) (len=8) "datetime",
        Start: (*string)(0xc0000545b0)((len=19) "2020-01-01T00:00:00"),
        End: (*string)(0xc0000545c0)((len=19) "2020-12-31T00:00:00"),
        IncludeStart: (bool) true,
        IncludeEnd: (bool) true
    })
})
```

## For Developers

After edit querystring.y, gen code via run:

```shell
go generate
```
