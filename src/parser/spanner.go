// Code generated by goyacc - DO NOT EDIT.

package parser

import __yyfmt__ "fmt"

import (
	"strconv"
	"strings"

	"github.com/CuberL/spar/src/types"
)

type yySymType struct {
	yys       int
	empty     struct{}
	flag      bool
	i64       int64
	str       string
	strs      []string
	col       types.Column
	cols      []types.Column
	coltype   types.ColumnType
	key       types.Key
	keys      []types.Key
	keyorder  types.KeyOrder
	clstr     types.Cluster
	ondelete  types.OnDelete
	stcls     types.StoringClause
	intlr     types.Interleave
	intlrs    []types.Interleave
	alt       types.Alteration
	LastToken int
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault              = 57393
	yyEofCode              = 57344
	ACTION                 = 57361
	ADD                    = 57366
	ALTER                  = 57374
	ARRAY                  = 57353
	ASC                    = 57348
	BOOL                   = 57379
	BYTES                  = 57383
	CASCADE                = 57359
	COLUMN                 = 57367
	COMMENT                = 57369
	CREATE                 = 57373
	DATABASE               = 57376
	DATE                   = 57384
	DELETE                 = 57358
	DESC                   = 57349
	DROP                   = 57375
	FLOAT64                = 57381
	IN                     = 57351
	INDEX                  = 57378
	INT64                  = 57380
	INTERLEAVE             = 57350
	KEY                    = 57347
	MAX                    = 57362
	NO                     = 57360
	NOT                    = 57355
	NULL                   = 57356
	NULL_FILTERED          = 57364
	ON                     = 57357
	OPTIONS                = 57354
	PARENT                 = 57352
	PRIMARY                = 57346
	SET                    = 57368
	STORING                = 57365
	STRING                 = 57382
	TABLE                  = 57377
	TIMESTAMP              = 57385
	UNIQUE                 = 57363
	allow_commit_timestamp = 57372
	column_name            = 57388
	comment_content        = 57390
	database_id            = 57386
	decimal_value          = 57391
	yyErrCode              = 57345
	hex_value              = 57392
	index_name             = 57389
	null                   = 57371
	table_name             = 57387
	true                   = 57370

	yyMaxDepth = 200
	yyTabOfs   = -72
)

var (
	yyPrec = map[int]int{}

	yyXLAT = map[int]int{
		59:    0,  // ';' (51x)
		41:    1,  // ')' (43x)
		44:    2,  // ',' (38x)
		57354: 3,  // OPTIONS (14x)
		57388: 4,  // column_name (12x)
		57355: 5,  // NOT (12x)
		57374: 6,  // ALTER (11x)
		57375: 7,  // DROP (11x)
		57373: 8,  // CREATE (10x)
		57344: 9,  // $end (9x)
		62:    10, // '>' (8x)
		40:    11, // '(' (7x)
		57378: 12, // INDEX (6x)
		57350: 13, // INTERLEAVE (6x)
		57387: 14, // table_name (6x)
		57398: 15, // column_def (4x)
		57379: 16, // BOOL (3x)
		57383: 17, // BYTES (3x)
		57367: 18, // COLUMN (3x)
		57399: 19, // column_def_list (3x)
		57384: 20, // DATE (3x)
		57381: 21, // FLOAT64 (3x)
		57380: 22, // INT64 (3x)
		57412: 23, // key_part (3x)
		57364: 24, // NULL_FILTERED (3x)
		57357: 25, // ON (3x)
		57420: 26, // scalar_type (3x)
		57382: 27, // STRING (3x)
		57377: 28, // TABLE (3x)
		57385: 29, // TIMESTAMP (3x)
		57394: 30, // alter_table (2x)
		57353: 31, // ARRAY (2x)
		57395: 32, // array_type (2x)
		57401: 33, // column_type (2x)
		57403: 34, // create_database (2x)
		57404: 35, // create_index (2x)
		57405: 36, // create_table (2x)
		57391: 37, // decimal_value (2x)
		57406: 38, // drop_index (2x)
		57407: 39, // drop_table (2x)
		57392: 40, // hex_value (2x)
		57351: 41, // IN (2x)
		57389: 42, // index_name (2x)
		57408: 43, // int64_value (2x)
		57409: 44, // interleave_clause (2x)
		57413: 45, // key_part_list (2x)
		57414: 46, // length (2x)
		57362: 47, // MAX (2x)
		57415: 48, // not_null_opt (2x)
		57417: 49, // on_delete_opt (2x)
		57418: 50, // options_def (2x)
		57368: 51, // SET (2x)
		57421: 52, // statement (2x)
		60:    53, // '<' (1x)
		61:    54, // '=' (1x)
		57361: 55, // ACTION (1x)
		57366: 56, // ADD (1x)
		57372: 57, // allow_commit_timestamp (1x)
		57348: 58, // ASC (1x)
		57359: 59, // CASCADE (1x)
		57396: 60, // cluster (1x)
		57397: 61, // cluster_opt (1x)
		57400: 62, // column_name_list (1x)
		57402: 63, // comment (1x)
		57390: 64, // comment_content (1x)
		57376: 65, // DATABASE (1x)
		57386: 66, // database_id (1x)
		57358: 67, // DELETE (1x)
		57349: 68, // DESC (1x)
		57410: 69, // interleave_clause_list (1x)
		57347: 70, // KEY (1x)
		57411: 71, // key_order_opt (1x)
		57360: 72, // NO (1x)
		57371: 73, // null (1x)
		57356: 74, // NULL (1x)
		57416: 75, // null_filtered_opt (1x)
		57352: 76, // PARENT (1x)
		57346: 77, // PRIMARY (1x)
		57419: 78, // primary_key (1x)
		57422: 79, // statements (1x)
		57365: 80, // STORING (1x)
		57423: 81, // storing_clause (1x)
		57424: 82, // storing_clause_opt (1x)
		57425: 83, // table_alteration (1x)
		57426: 84, // table_column_alteration (1x)
		57370: 85, // true (1x)
		57363: 86, // UNIQUE (1x)
		57427: 87, // unique_opt (1x)
		57393: 88, // $default (0x)
		35:    89, // '#' (0x)
		47:    90, // '/' (0x)
		57369: 91, // COMMENT (0x)
		57345: 92, // error (0x)
	}

	yySymNames = []string{
		"';'",
		"')'",
		"','",
		"OPTIONS",
		"column_name",
		"NOT",
		"ALTER",
		"DROP",
		"CREATE",
		"$end",
		"'>'",
		"'('",
		"INDEX",
		"INTERLEAVE",
		"table_name",
		"column_def",
		"BOOL",
		"BYTES",
		"COLUMN",
		"column_def_list",
		"DATE",
		"FLOAT64",
		"INT64",
		"key_part",
		"NULL_FILTERED",
		"ON",
		"scalar_type",
		"STRING",
		"TABLE",
		"TIMESTAMP",
		"alter_table",
		"ARRAY",
		"array_type",
		"column_type",
		"create_database",
		"create_index",
		"create_table",
		"decimal_value",
		"drop_index",
		"drop_table",
		"hex_value",
		"IN",
		"index_name",
		"int64_value",
		"interleave_clause",
		"key_part_list",
		"length",
		"MAX",
		"not_null_opt",
		"on_delete_opt",
		"options_def",
		"SET",
		"statement",
		"'<'",
		"'='",
		"ACTION",
		"ADD",
		"allow_commit_timestamp",
		"ASC",
		"CASCADE",
		"cluster",
		"cluster_opt",
		"column_name_list",
		"comment",
		"comment_content",
		"DATABASE",
		"database_id",
		"DELETE",
		"DESC",
		"interleave_clause_list",
		"KEY",
		"key_order_opt",
		"NO",
		"null",
		"NULL",
		"null_filtered_opt",
		"PARENT",
		"PRIMARY",
		"primary_key",
		"statements",
		"STORING",
		"storing_clause",
		"storing_clause_opt",
		"table_alteration",
		"table_column_alteration",
		"true",
		"UNIQUE",
		"unique_opt",
		"$default",
		"'#'",
		"'/'",
		"COMMENT",
		"error",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {79, 1},
		2:  {79, 2},
		3:  {52, 2},
		4:  {52, 2},
		5:  {52, 2},
		6:  {52, 2},
		7:  {52, 2},
		8:  {52, 2},
		9:  {34, 3},
		10: {36, 8},
		11: {19, 0},
		12: {19, 1},
		13: {19, 4},
		14: {19, 3},
		15: {15, 4},
		16: {78, 5},
		17: {45, 1},
		18: {45, 3},
		19: {23, 2},
		20: {71, 0},
		21: {71, 1},
		22: {71, 1},
		23: {61, 0},
		24: {61, 2},
		25: {60, 5},
		26: {49, 0},
		27: {49, 3},
		28: {49, 4},
		29: {33, 1},
		30: {33, 1},
		31: {26, 1},
		32: {26, 1},
		33: {26, 1},
		34: {26, 4},
		35: {26, 4},
		36: {26, 1},
		37: {26, 1},
		38: {46, 1},
		39: {46, 1},
		40: {32, 4},
		41: {50, 0},
		42: {50, 6},
		43: {50, 6},
		44: {48, 0},
		45: {48, 2},
		46: {63, 1},
		47: {35, 12},
		48: {87, 0},
		49: {87, 1},
		50: {75, 0},
		51: {75, 1},
		52: {82, 0},
		53: {82, 1},
		54: {81, 4},
		55: {62, 1},
		56: {62, 3},
		57: {69, 0},
		58: {69, 1},
		59: {69, 3},
		60: {44, 3},
		61: {30, 4},
		62: {30, 4},
		63: {83, 3},
		64: {83, 3},
		65: {83, 2},
		66: {84, 5},
		67: {84, 5},
		68: {39, 3},
		69: {38, 3},
		70: {43, 1},
		71: {43, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [145][]uint16{
		// 0
		{6: 82, 83, 81, 30: 78, 34: 75, 77, 76, 38: 80, 79, 52: 74, 79: 73},
		{6: 82, 83, 81, 72, 30: 78, 34: 75, 77, 76, 38: 80, 79, 52: 216},
		{6: 71, 71, 71, 71},
		{215},
		{214},
		// 5
		{213},
		{212},
		{211},
		{210},
		{12: 24, 24: 24, 28: 150, 65: 149, 86: 152, 151},
		// 10
		{28: 88},
		{12: 85, 28: 84},
		{14: 87},
		{42: 86},
		{3},
		// 15
		{4},
		{14: 89},
		{6: 95, 93, 51: 94, 56: 92, 83: 90, 91},
		{11},
		{10},
		// 20
		{18: 143},
		{18: 141},
		{46, 25: 135, 49: 136},
		{18: 96},
		{4: 97},
		// 25
		{16: 100, 104, 20: 105, 102, 101, 26: 98, 103, 29: 106, 31: 107, 99, 108, 51: 109},
		{43, 43, 43, 43, 5: 43},
		{42, 42, 42, 42, 5: 42},
		{41, 41, 41, 41, 5: 41, 10: 41},
		{40, 40, 40, 40, 5: 40, 10: 40},
		// 30
		{39, 39, 39, 39, 5: 39, 10: 39},
		{11: 132},
		{11: 125},
		{36, 36, 36, 36, 5: 36, 10: 36},
		{35, 35, 35, 35, 5: 35, 10: 35},
		// 35
		{53: 122},
		{28, 5: 119, 48: 120},
		{31, 3: 110, 50: 111},
		{11: 112},
		{5},
		// 40
		{57: 113},
		{54: 114},
		{73: 116, 85: 115},
		{1: 118},
		{1: 117},
		// 45
		{29, 29, 29},
		{30, 30, 30},
		{74: 121},
		{6},
		{27, 27, 27, 27},
		// 50
		{16: 100, 104, 20: 105, 102, 101, 26: 123, 103, 29: 106},
		{10: 124},
		{32, 32, 32, 32, 5: 32},
		{37: 129, 40: 130, 43: 127, 46: 126, 128},
		{1: 131},
		// 55
		{1: 34},
		{1: 33},
		{1: 2},
		{1: 1},
		{37, 37, 37, 37, 5: 37, 10: 37},
		// 60
		{37: 129, 40: 130, 43: 127, 46: 133, 128},
		{1: 134},
		{38, 38, 38, 38, 5: 38, 10: 38},
		{67: 137},
		{7},
		// 65
		{59: 138, 72: 139},
		{45},
		{55: 140},
		{44},
		{4: 142},
		// 70
		{8},
		{4: 144, 15: 145},
		{16: 100, 104, 20: 105, 102, 101, 26: 98, 103, 29: 106, 31: 107, 99, 146},
		{9},
		{28, 28, 28, 28, 5: 119, 48: 147},
		// 75
		{31, 31, 31, 110, 50: 148},
		{57, 57, 57},
		{66: 209},
		{14: 185},
		{12: 22, 24: 154, 75: 153},
		// 80
		{12: 23, 24: 23},
		{12: 155},
		{12: 21},
		{42: 156},
		{25: 157},
		// 85
		{14: 158},
		{11: 159},
		{4: 162, 23: 160, 45: 161},
		{1: 55, 55},
		{1: 167, 166},
		// 90
		{1: 52, 52, 58: 164, 68: 165, 71: 163},
		{1: 53, 53},
		{1: 51, 51},
		{1: 50, 50},
		{4: 162, 23: 184},
		// 95
		{20, 2: 20, 13: 20, 80: 170, 169, 168},
		{15, 2: 15, 13: 179, 44: 178, 69: 177},
		{19, 2: 19, 13: 19},
		{11: 171},
		{4: 173, 62: 172},
		// 100
		{1: 174, 175},
		{1: 17, 17},
		{18, 2: 18, 13: 18},
		{4: 176},
		{1: 16, 16},
		// 105
		{25, 2: 182},
		{14, 2: 14},
		{41: 180},
		{14: 181},
		{12, 2: 12},
		// 110
		{13: 179, 44: 183},
		{13, 2: 13},
		{1: 54, 54},
		{11: 186},
		{1: 61, 4: 144, 15: 188, 19: 187},
		// 115
		{1: 194},
		{1: 60, 189},
		{1: 61, 4: 144, 15: 188, 19: 191, 63: 190, 192},
		{1: 61, 4: 144, 15: 188, 19: 193},
		{1: 58},
		// 120
		{1: 26, 4: 26},
		{1: 59},
		{77: 196, 195},
		{49, 2: 202, 61: 201},
		{70: 197},
		// 125
		{11: 198},
		{4: 162, 23: 160, 45: 199},
		{1: 200, 166},
		{56, 2: 56},
		{62},
		// 130
		{13: 204, 60: 203},
		{48},
		{41: 205},
		{76: 206},
		{14: 207},
		// 135
		{46, 25: 135, 49: 208},
		{47},
		{63},
		{6: 64, 64, 64, 64},
		{6: 65, 65, 65, 65},
		// 140
		{6: 66, 66, 66, 66},
		{6: 67, 67, 67, 67},
		{6: 68, 68, 68, 68},
		{6: 69, 69, 69, 69},
		{6: 70, 70, 70, 70},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("%q", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 92

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 9:
		{
			s := types.CreateDatabaseStatement{
				DatabaseId: yyS[yypt-0].str,
			}
			yylex.(*lexerWrapper).result.CreateDatabases = append(yylex.(*lexerWrapper).result.CreateDatabases, s)
		}
	case 10:
		{
			s := types.CreateTableStatement{
				TableName:   yyS[yypt-5].str,
				Columns:     yyS[yypt-3].cols,
				PrimaryKeys: yyS[yypt-1].keys,
				Cluster:     yyS[yypt-0].clstr,
			}
			yylex.(*lexerWrapper).result.CreateTables = append(yylex.(*lexerWrapper).result.CreateTables, s)
		}
	case 11:
		{
			yyVAL.cols = make([]types.Column, 0, 0)
		}
	case 12:
		{
			yyVAL.cols = make([]types.Column, 0, 1)
			yyVAL.cols = append(yyVAL.cols, yyS[yypt-0].col)
		}
	case 13:
		{
			yyS[yypt-3].col.Comment = yyS[yypt-1].str
			yyVAL.cols = append(yyS[yypt-0].cols, yyS[yypt-3].col)
		}
	case 14:
		{
			yyVAL.cols = append(yyS[yypt-0].cols, yyS[yypt-2].col)
		}
	case 15:
		{
			yyVAL.col = types.Column{Name: yyS[yypt-3].str, Type: yyS[yypt-2].coltype, NotNull: yyS[yypt-1].flag, Options: yyS[yypt-0].str}
		}
	case 16:
		{
			yyVAL.keys = yyS[yypt-1].keys
		}
	case 17:
		{
			yyVAL.keys = make([]types.Key, 0, 1)
			yyVAL.keys = append(yyVAL.keys, yyS[yypt-0].key)
		}
	case 18:
		{
			yyVAL.keys = append(yyS[yypt-2].keys, yyS[yypt-0].key)
		}
	case 19:
		{
			yyVAL.key = types.Key{Name: yyS[yypt-1].str, KeyOrder: yyS[yypt-0].keyorder}
		}
	case 20:
		{
			yyVAL.keyorder = types.Asc
		}
	case 21:
		{
			yyVAL.keyorder = types.Asc
		}
	case 22:
		{
			yyVAL.keyorder = types.Desc
		}
	case 23:
		{
			yyVAL.clstr = types.Cluster{}
		}
	case 24:
		{
			yyVAL.clstr = yyS[yypt-0].clstr
		}
	case 25:
		{
			yyVAL.clstr = types.Cluster{TableName: yyS[yypt-1].str, OnDelete: yyS[yypt-0].ondelete}
		}
	case 26:
		{
			// default
			yyVAL.ondelete = types.NoAction
		}
	case 27:
		{
			yyVAL.ondelete = types.Cascade
		}
	case 28:
		{
			yyVAL.ondelete = types.NoAction
		}
	case 29:
		{
			yyVAL.coltype = yyS[yypt-0].coltype
		}
	case 30:
		{
			yyVAL.coltype = yyS[yypt-0].coltype
		}
	case 31:
		{
			yyVAL.coltype = types.ColumnType{TypeTag: types.Bool}
		}
	case 32:
		{
			yyVAL.coltype = types.ColumnType{TypeTag: types.Int64}
		}
	case 33:
		{
			yyVAL.coltype = types.ColumnType{TypeTag: types.Float64}
		}
	case 34:
		{
			yyVAL.coltype = types.ColumnType{TypeTag: types.String, Length: yyS[yypt-1].i64}
		}
	case 35:
		{
			yyVAL.coltype = types.ColumnType{TypeTag: types.Bytes, Length: yyS[yypt-1].i64}
		}
	case 36:
		{
			yyVAL.coltype = types.ColumnType{TypeTag: types.Date}
		}
	case 37:
		{
			yyVAL.coltype = types.ColumnType{TypeTag: types.Timestamp}
		}
	case 38:
		{
			yyVAL.i64 = yyS[yypt-0].i64
		}
	case 39:
		{
			yyVAL.i64 = 10485760
		}
	case 40:
		{
			yyVAL.coltype = yyS[yypt-1].coltype
			yyVAL.coltype.IsArray = types.True
		}
	case 41:
		{
			yyVAL.str = ""
		}
	case 42:
		{
			yyVAL.str = yyS[yypt-3].str + "=" + yyS[yypt-1].str
		}
	case 43:
		{
			yyVAL.str = yyS[yypt-3].str + "=" + yyS[yypt-1].str
		}
	case 44:
		{
			yyVAL.flag = types.False
		}
	case 45:
		{
			yyVAL.flag = types.True
		}
	case 46:
		{
			yyVAL.str = strings.TrimSpace(strings.TrimRight(strings.TrimLeft(yyS[yypt-0].str, "/*"), "*/"))
		}
	case 47:
		{
			s := types.CreateIndexStatement{
				Unique:        yyS[yypt-10].flag,
				NullFiltered:  yyS[yypt-9].flag,
				IndexName:     yyS[yypt-7].str,
				TableName:     yyS[yypt-5].str,
				Keys:          yyS[yypt-3].keys,
				StoringClause: yyS[yypt-1].stcls,
				Interleaves:   yyS[yypt-0].intlrs,
			}
			yylex.(*lexerWrapper).result.CreateIndexes = append(yylex.(*lexerWrapper).result.CreateIndexes, s)
		}
	case 48:
		{
			yyVAL.flag = types.False
		}
	case 49:
		{
			yyVAL.flag = types.True
		}
	case 50:
		{
			yyVAL.flag = types.False
		}
	case 51:
		{
			yyVAL.flag = types.True
		}
	case 52:
		{
			yyVAL.stcls = types.StoringClause{}
		}
	case 53:
		{
			yyVAL.stcls = yyS[yypt-0].stcls
		}
	case 54:
		{
			yyVAL.stcls = types.StoringClause{ColumnNames: yyS[yypt-1].strs}
		}
	case 55:
		{
			yyVAL.strs = make([]string, 0, 1)
			yyVAL.strs = append(yyVAL.strs, yyS[yypt-0].str)
		}
	case 56:
		{
			yyVAL.strs = append(yyS[yypt-2].strs, yyS[yypt-0].str)
		}
	case 57:
		{
			yyVAL.intlrs = make([]types.Interleave, 0, 0)
		}
	case 58:
		{
			yyVAL.intlrs = make([]types.Interleave, 0, 1)
			yyVAL.intlrs = append(yyVAL.intlrs, yyS[yypt-0].intlr)
		}
	case 59:
		{
			yyVAL.intlrs = append(yyS[yypt-2].intlrs, yyS[yypt-0].intlr)
		}
	case 60:
		{
			yyVAL.intlr = types.Interleave{TableName: yyS[yypt-0].str}
		}
	case 61:
		{
			s := types.AlterTableStatement{
				TableName:  yyS[yypt-1].str,
				Alteration: yyS[yypt-0].alt,
			}
			yylex.(*lexerWrapper).result.AlterTables = append(yylex.(*lexerWrapper).result.AlterTables, s)
		}
	case 62:
		{
			s := types.AlterTableStatement{
				TableName:  yyS[yypt-1].str,
				Alteration: yyS[yypt-0].alt,
			}
			yylex.(*lexerWrapper).result.AlterTables = append(yylex.(*lexerWrapper).result.AlterTables, s)
		}
	case 63:
		{
			yyVAL.alt = &types.AddColumnTableAlteration{
				Column: yyS[yypt-0].col,
			}
		}
	case 64:
		{
			yyVAL.alt = &types.DropColumnTableAlteration{
				ColumnName: yyS[yypt-0].str,
			}
		}
	case 65:
		{
			yyVAL.alt = &types.SetTableAlteration{
				OnDelete: yyS[yypt-0].ondelete,
			}
		}
	case 66:
		{
			yyVAL.alt = &types.AlterColumnTypesAlteration{
				ColumnName: yyS[yypt-2].str,
				ColumnType: yyS[yypt-1].coltype,
				NotNull:    yyS[yypt-0].flag,
			}
		}
	case 67:
		{
			yyVAL.alt = &types.AlterColumnSetAlteration{
				ColumnName: yyS[yypt-2].str,
				Options:    yyS[yypt-0].str,
			}
		}
	case 68:
		{
			s := types.DropTableStatement{
				TableName: yyS[yypt-0].str,
			}
			yylex.(*lexerWrapper).result.DropTables = append(yylex.(*lexerWrapper).result.DropTables, s)
		}
	case 69:
		{
			s := types.DropIndexStatement{
				IndexName: yyS[yypt-0].str,
			}
			yylex.(*lexerWrapper).result.DropIndexes = append(yylex.(*lexerWrapper).result.DropIndexes, s)
		}
	case 70:
		{
			v, _ := strconv.ParseInt(yyS[yypt-0].str, 10, 64)
			yyVAL.i64 = v
		}
	case 71:
		{
			v, _ := strconv.ParseInt(yyS[yypt-0].str, 16, 64)
			yyVAL.i64 = v
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
