package querystring

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, expected, cond)
}

func TestParserMixedCondition(t *testing.T) {
	cond, err := Parse("a: 1 OR (b: 2 and c: 4)")
	if err != nil {
		t.Errorf("parse return error, %s", err)
		return
	}

	assert.Equal(t, &OrCondition{
		Left: &NumberRangeCondition{
			Field:        "a",
			Start:        pointer("1"),
			End:          pointer("1"),
			IncludeStart: true,
			IncludeEnd:   true,
		},
		Right: &AndCondition{
			Left: &NumberRangeCondition{
				Field:        "b",
				Start:        pointer("2"),
				End:          pointer("2"),
				IncludeStart: true,
				IncludeEnd:   true,
			},
			Right: &NumberRangeCondition{
				Field:        "c",
				Start:        pointer("4"),
				End:          pointer("4"),
				IncludeStart: true,
				IncludeEnd:   true,
			},
		},
	}, cond)
}

func pointer(s string) *string {
	return &s
}
