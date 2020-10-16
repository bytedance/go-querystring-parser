%{
package querystring
%}

%union {
	s 	string
	n 	int
	q 	Condition
}

%token tSTRING tPHRASE tNUMBER
%token tAND tOR tNOT tTO tPLUS tMINUS tCOLON
%token tLEFTBRACKET tRIGHTBRACKET tLEFTRANGE tRIGHTRANGE
%token tGREATER tLESS tEQUAL

%type <s>                tSTRING
%type <s>                tPHRASE
%type <s>                tNUMBER
%type <s>                posOrNegNumber
%type <q>                searchBase searchParts searchPart searchLogicPart searchLogicSimplePart
%type <n>                searchPrefix

%%

input:
searchParts {
	yylex.(*lexerWrapper).query = $1
};

searchParts:
searchPart searchParts {
	$$ = NewAndCondition($1, $2)
}
|
searchPart {
	$$ = $1
};

searchPart:
searchPrefix searchBase {
	switch($1) {
	case queryMustNot:
		$$ = NewNotCondition($2)
	default:
		$$ = $2
	}
}
|
searchLogicPart {
	$$ = $1
};

searchLogicPart:
searchLogicSimplePart {
	$$ = $1
}
|
searchLogicSimplePart tAND searchLogicPart {
	$$ = NewAndCondition($1, $3)
}
|
searchLogicSimplePart tOR searchLogicPart {
	$$ = NewOrCondition($1, $3)
};

searchLogicSimplePart:
searchBase {
	$$ = $1
}
|
tLEFTBRACKET searchLogicPart tRIGHTBRACKET {
	$$ = $2
}
|
tNOT searchLogicSimplePart {
	$$ = NewNotCondition($2)
};

searchPrefix:
tPLUS {
	$$ = queryMust
}
|
tMINUS {
	$$ = queryMustNot
};

searchBase:
tSTRING {
	$$ = newStringCondition($1)
}
|
tNUMBER {
	$$ = NewMatchCondition($1)
}
|
tPHRASE {
	phrase := $1
	q := NewMatchCondition(phrase)
	$$ = q
}
|
tSTRING tCOLON tSTRING {
	q := newStringCondition($3)
	q.SetField($1)
	$$ = q
}
|
tSTRING tCOLON posOrNegNumber {
	val := $3
	q := NewNumberRangeCondition(&val, &val, true, true)
	q.SetField($1)
	$$ = q
}
|
tSTRING tCOLON tPHRASE {
	q := NewMatchCondition($3)
	q.SetField($1)
	$$ = q
}
|
tSTRING tCOLON tGREATER posOrNegNumber {
	val := $4
	q := NewNumberRangeCondition(&val, nil, false, false)
	q.SetField($1)
	$$ = q
}
|
tSTRING tCOLON tGREATER tEQUAL posOrNegNumber {
	val := $5
	q := NewNumberRangeCondition(&val, nil, true, false)
	q.SetField($1)
	$$ = q
}
|
tSTRING tCOLON tLESS posOrNegNumber {
	val := $4
	q := NewNumberRangeCondition(nil, &val, false, false)
	q.SetField($1)
	$$ = q
}
|
tSTRING tCOLON tLESS tEQUAL posOrNegNumber {
	val := $5
	q := NewNumberRangeCondition(nil, &val, false, true)
	q.SetField($1)
	$$ = q
}
|
tSTRING tCOLON tGREATER tPHRASE {
	phrase := $4
	q := NewTimeRangeCondition(&phrase, nil, false, false)
	q.SetField($1)
	$$ = q
}
|
tSTRING tCOLON tGREATER tEQUAL tPHRASE {
	phrase := $5
	q := NewTimeRangeCondition(&phrase, nil, true, false)
	q.SetField($1)
	$$ = q
}
|
tSTRING tCOLON tLESS tPHRASE {
	phrase := $4
	q := NewTimeRangeCondition(nil, &phrase, false, false)
	q.SetField($1)
	$$ = q
}
|
tSTRING tCOLON tLESS tEQUAL tPHRASE {
	phrase := $5
	q := NewTimeRangeCondition(nil, &phrase, false, true)
	q.SetField($1)
	$$ = q
}
|
tSTRING tCOLON tLEFTRANGE posOrNegNumber tTO posOrNegNumber tRIGHTRANGE {
	min := $4
	max := $6
	q := NewNumberRangeCondition(&min, &max, true, true)
	q.SetField($1)
	$$ = q
}
|
tSTRING tCOLON tLEFTRANGE tPHRASE tTO tPHRASE tRIGHTRANGE {
	min := $4
	max := $6
	q := NewTimeRangeCondition(&min, &max, true, true)
	q.SetField($1)
	$$ = q
};

posOrNegNumber:
tNUMBER {
	$$ = $1
}
|
tMINUS tNUMBER {
	$$ = "-" + $2
};
