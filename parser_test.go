package querystring

import (
	"testing"

	"github.com/go-test/deep"
)

func TestParser(t *testing.T) {
	cond, err := Parse("message: test\\ value AND datetime: [\"2020-01-01T00:00:00\" TO \"2020-12-31T00:00:00\"]")
	if err != nil {
		t.Errorf("parse return error, %s", err)
		return
	}

	expected := Condition(&AndCondition{
		Left: &MatchCondition{
			Field: "message",
			Value: "test value",
		},
		Right: &TimeRangeCondition{
			Field:        "datetime",
			Start:        pointer("2020-01-01T00:00:00"),
			End:          pointer("2020-12-31T00:00:00"),
			IncludeStart: true,
			IncludeEnd:   true,
		},
	})

	if diff := deep.Equal(cond, expected); diff != nil {
		t.Errorf("returned condition unexpected: diff= %s", diff)
		return
	}
}

func pointer(s string) *string {
	return &s
}
