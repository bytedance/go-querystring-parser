package querystring

import (
	"strings"
	"time"
)

// Condition .
type Condition interface {
}

// AndCondition .
type AndCondition struct {
	Left  Condition
	Right Condition
}

// NewAndCondition .
func NewAndCondition(left, right Condition) *AndCondition {
	return &AndCondition{Left: left, Right: right}
}

// OrCondition .
type OrCondition struct {
	Left  Condition
	Right Condition
}

// NewOrCondition .
func NewOrCondition(left, right Condition) *OrCondition {
	return &OrCondition{Left: left, Right: right}
}

// NotCondition .
type NotCondition struct {
	Condition Condition
}

// NewNotCondition .
func NewNotCondition(q Condition) *NotCondition {
	return &NotCondition{Condition: q}
}

// FieldableCondition .
type FieldableCondition interface {
	SetField(field string)
}

// MatchCondition .
type MatchCondition struct {
	Field string
	Value string
}

// NewMatchCondition .
func NewMatchCondition(s string) *MatchCondition {
	return &MatchCondition{
		Value: s,
	}
}

// SetField .
func (q *MatchCondition) SetField(field string) {
	q.Field = field
}

// RegexpCondition .
type RegexpCondition struct {
	Field string
	Value string
}

// NewRegexpCondition .
func NewRegexpCondition(s string) *RegexpCondition {
	return &RegexpCondition{
		Value: s,
	}
}

// SetField .
func (q *RegexpCondition) SetField(field string) {
	q.Field = field
}

// WildcardCondition .
type WildcardCondition struct {
	Field string
	Value string
}

// NewWildcardCondition .
func NewWildcardCondition(s string) *WildcardCondition {
	return &WildcardCondition{
		Value: s,
	}
}

// SetField .
func (q *WildcardCondition) SetField(field string) {
	q.Field = field
}

// NumberRangeCondition .
type NumberRangeCondition struct {
	Field        string
	Start        *string
	End          *string
	IncludeStart bool
	IncludeEnd   bool
}

// NewNumberRangeCondition .
func NewNumberRangeCondition(start, end *string, includeStart, includeEnd bool) *NumberRangeCondition {
	return &NumberRangeCondition{
		Start:        start,
		End:          end,
		IncludeStart: includeStart,
		IncludeEnd:   includeEnd,
	}
}

// SetField .
func (q *NumberRangeCondition) SetField(field string) {
	q.Field = field
}

// TimeRangeCondition .
type TimeRangeCondition struct {
	Field        string
	Start        *string
	End          *string
	IncludeStart bool
	IncludeEnd   bool
}

// NewTimeRangeCondition .
func NewTimeRangeCondition(start, end *string, includeStart, includeEnd bool) *TimeRangeCondition {
	return &TimeRangeCondition{
		Start:        start,
		End:          end,
		IncludeStart: includeStart,
		IncludeEnd:   includeEnd,
	}
}

// SetField .
func (q *TimeRangeCondition) SetField(field string) {
	q.Field = field
}

func queryTimeFromString(t string) (time.Time, error) {
	return time.Parse(time.RFC3339, t)
}

func newStringCondition(str string) FieldableCondition {
	if strings.HasPrefix(str, "/") && strings.HasSuffix(str, "/") {
		return NewRegexpCondition(str[1 : len(str)-1])
	}

	if strings.ContainsAny(str, "*?") {
		return NewWildcardCondition(str)
	}

	return NewMatchCondition(str)
}
