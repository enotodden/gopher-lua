// Code generated by goyacc -o parser.go parser.go.y. DO NOT EDIT.

//line parser.go.y:2
package parse

import __yyfmt__ "fmt"

//line parser.go.y:2

import (
	"github.com/enotodden/gopher-lua/ast"
)

//line parser.go.y:34
type yySymType struct {
	yys   int
	token ast.Token

	stmts []ast.Stmt
	stmt  ast.Stmt

	funcname *ast.FuncName
	funcexpr *ast.FunctionExpr

	exprlist []ast.Expr
	expr     ast.Expr

	fieldlist []*ast.Field
	field     *ast.Field
	fieldsep  string

	namelist []string
	parlist  *ast.ParList
}

const TAnd = 57346
const TBreak = 57347
const TDo = 57348
const TElse = 57349
const TElseIf = 57350
const TEnd = 57351
const TFalse = 57352
const TFor = 57353
const TFunction = 57354
const TIf = 57355
const TIn = 57356
const TLocal = 57357
const TNil = 57358
const TNot = 57359
const TOr = 57360
const TReturn = 57361
const TRepeat = 57362
const TThen = 57363
const TTrue = 57364
const TUntil = 57365
const TWhile = 57366
const TGoto = 57367
const TEqeq = 57368
const TNeq = 57369
const TLte = 57370
const TGte = 57371
const T2Comma = 57372
const T3Comma = 57373
const T2Colon = 57374
const TIdent = 57375
const TNumber = 57376
const TString = 57377
const UNARY = 57378

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TAnd",
	"TBreak",
	"TDo",
	"TElse",
	"TElseIf",
	"TEnd",
	"TFalse",
	"TFor",
	"TFunction",
	"TIf",
	"TIn",
	"TLocal",
	"TNil",
	"TNot",
	"TOr",
	"TReturn",
	"TRepeat",
	"TThen",
	"TTrue",
	"TUntil",
	"TWhile",
	"TGoto",
	"TEqeq",
	"TNeq",
	"TLte",
	"TGte",
	"T2Comma",
	"T3Comma",
	"T2Colon",
	"TIdent",
	"TNumber",
	"TString",
	"'{'",
	"'('",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'|'",
	"'%'",
	"UNARY",
	"'^'",
	"';'",
	"'='",
	"','",
	"':'",
	"'.'",
	"'['",
	"']'",
	"'#'",
	"')'",
	"'}'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:529

func TokenName(c int) string {
	if c >= TAnd && c-TAnd < len(yyToknames) {
		if yyToknames[c-TAnd] != "" {
			return yyToknames[c-TAnd]
		}
	}
	return string([]byte{byte(c)})
}

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 19,
	49, 33,
	50, 33,
	-2, 71,
	-1, 98,
	49, 34,
	50, 34,
	-2, 71,
}

const yyPrivate = 57344

const yyLast = 633

var yyAct = [...]uint8{
	26, 93, 52, 25, 47, 89, 160, 58, 144, 119,
	143, 169, 54, 69, 56, 55, 35, 141, 41, 42,
	49, 110, 34, 67, 63, 69, 50, 139, 162, 64,
	173, 149, 51, 145, 48, 46, 45, 86, 87, 88,
	113, 114, 109, 96, 116, 111, 100, 97, 24, 78,
	50, 85, 138, 104, 23, 69, 51, 111, 22, 79,
	80, 81, 82, 83, 84, 112, 85, 43, 44, 90,
	120, 121, 122, 123, 124, 125, 126, 127, 128, 129,
	130, 131, 132, 133, 134, 135, 136, 81, 82, 83,
	84, 117, 85, 41, 42, 49, 146, 33, 140, 172,
	9, 155, 62, 157, 156, 155, 115, 148, 151, 150,
	153, 152, 40, 71, 154, 19, 50, 102, 101, 50,
	159, 158, 51, 64, 66, 51, 65, 70, 61, 57,
	107, 194, 21, 191, 186, 76, 77, 75, 74, 78,
	185, 161, 99, 96, 163, 179, 164, 72, 73, 79,
	80, 81, 82, 83, 84, 68, 85, 98, 175, 176,
	174, 171, 166, 170, 105, 118, 53, 1, 142, 177,
	92, 137, 178, 32, 180, 20, 8, 182, 181, 60,
	59, 28, 3, 39, 167, 189, 188, 27, 37, 4,
	190, 2, 0, 29, 0, 193, 71, 0, 0, 0,
	0, 0, 31, 0, 94, 30, 41, 42, 22, 0,
	70, 0, 36, 0, 0, 0, 0, 0, 76, 77,
	75, 74, 78, 0, 95, 0, 38, 71, 91, 0,
	72, 73, 79, 80, 81, 82, 83, 84, 0, 85,
	0, 70, 0, 0, 0, 0, 165, 0, 0, 76,
	77, 75, 74, 78, 0, 0, 0, 0, 0, 0,
	0, 72, 73, 79, 80, 81, 82, 83, 84, 28,
	85, 39, 0, 0, 0, 27, 37, 147, 0, 0,
	0, 29, 0, 0, 0, 0, 0, 0, 0, 0,
	31, 0, 23, 30, 41, 42, 22, 28, 0, 39,
	36, 0, 0, 27, 37, 0, 0, 0, 0, 29,
	0, 0, 0, 0, 38, 103, 0, 0, 31, 0,
	94, 30, 41, 42, 22, 28, 0, 39, 36, 0,
	0, 27, 37, 0, 0, 0, 0, 29, 0, 71,
	95, 183, 38, 0, 0, 0, 31, 0, 23, 30,
	41, 42, 22, 70, 0, 0, 36, 0, 0, 0,
	0, 76, 77, 75, 74, 78, 0, 71, 0, 0,
	38, 0, 0, 72, 73, 79, 80, 81, 82, 83,
	84, 70, 85, 0, 0, 184, 0, 0, 0, 76,
	77, 75, 74, 78, 0, 71, 0, 192, 0, 0,
	0, 72, 73, 79, 80, 81, 82, 83, 84, 70,
	85, 0, 0, 168, 0, 0, 0, 76, 77, 75,
	74, 78, 0, 71, 0, 0, 0, 0, 0, 72,
	73, 79, 80, 81, 82, 83, 84, 70, 85, 0,
	187, 0, 0, 0, 0, 76, 77, 75, 74, 78,
	0, 71, 0, 0, 0, 0, 0, 72, 73, 79,
	80, 81, 82, 83, 84, 70, 85, 0, 108, 0,
	0, 0, 0, 76, 77, 75, 74, 78, 0, 71,
	0, 106, 0, 0, 0, 72, 73, 79, 80, 81,
	82, 83, 84, 70, 85, 0, 0, 0, 0, 0,
	0, 76, 77, 75, 74, 78, 0, 71, 0, 0,
	0, 0, 0, 72, 73, 79, 80, 81, 82, 83,
	84, 70, 85, 0, 0, 0, 0, 0, 0, 76,
	77, 75, 74, 78, 71, 0, 0, 0, 0, 0,
	0, 72, 73, 79, 80, 81, 82, 83, 84, 0,
	85, 0, 0, 0, 0, 0, 76, 77, 75, 74,
	78, 0, 0, 0, 0, 0, 0, 0, 72, 73,
	79, 80, 81, 82, 83, 84, 0, 85, 7, 10,
	0, 0, 0, 0, 14, 15, 13, 0, 16, 0,
	0, 0, 6, 12, 0, 0, 0, 11, 18, 0,
	0, 0, 0, 0, 0, 17, 23, 0, 0, 0,
	22, 76, 77, 75, 74, 78, 0, 0, 0, 0,
	0, 5, 0, 72, 73, 79, 80, 81, 82, 83,
	84, 0, 85,
}

var yyPact = [...]int16{
	-1000, -1000, 573, 0, -1000, -1000, 315, -1000, 18, -17,
	-1000, 315, -1000, 315, 96, 95, 90, 93, 91, -1000,
	-1000, -1000, 315, -1000, -1000, -37, 503, -1000, -1000, -1000,
	-1000, -1000, -1000, -17, -1000, -1000, 315, 315, 315, 32,
	-1000, -1000, 171, 315, 21, 315, 85, -1000, 84, 259,
	-1000, -1000, 155, -1000, 475, 107, 447, -7, 7, 32,
	-11, -1000, 73, -5, -1000, 59, -1000, 109, -47, 315,
	315, 315, 315, 315, 315, 315, 315, 315, 315, 315,
	315, 315, 315, 315, 315, 315, 4, 4, 4, -1000,
	-4, -1000, -40, -1000, -16, 315, 503, -37, -1000, -17,
	223, -1000, 58, -1000, -25, -1000, -1000, 315, -1000, 315,
	315, 72, -1000, 71, 70, 32, 315, -1000, -1000, -1000,
	503, 530, 585, 19, 19, 19, 19, 19, 19, 19,
	45, 45, 4, 4, 4, 4, 4, -50, -1000, -1000,
	-22, -1000, 287, -1000, -1000, 315, 192, -1000, -1000, -1000,
	153, 503, -1000, 363, 5, -1000, -1000, -1000, -1000, -37,
	-1000, 152, 68, -1000, 503, -19, -1000, 151, 315, -1000,
	136, -1000, -1000, 315, -1000, -1000, 315, 335, 131, -1000,
	503, 125, 419, -1000, 315, -1000, -1000, -1000, 124, 391,
	-1000, -1000, -1000, 122, -1000,
}

var yyPgo = [...]uint8{
	0, 166, 191, 2, 189, 184, 182, 180, 179, 176,
	112, 7, 3, 0, 22, 97, 132, 175, 4, 173,
	5, 171, 16, 170, 1, 168,
}

var yyR1 = [...]int8{
	0, 1, 1, 1, 2, 2, 2, 3, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 5, 5, 6, 6, 6, 7,
	7, 8, 8, 9, 9, 10, 10, 10, 11, 11,
	12, 12, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	14, 15, 15, 15, 15, 17, 16, 16, 18, 18,
	18, 18, 19, 20, 20, 21, 21, 21, 22, 22,
	23, 23, 23, 24, 24, 24, 25, 25,
}

var yyR2 = [...]int8{
	0, 1, 2, 3, 0, 2, 2, 1, 3, 1,
	3, 5, 4, 6, 8, 9, 11, 7, 3, 4,
	4, 2, 3, 2, 0, 5, 1, 2, 1, 1,
	3, 1, 3, 1, 3, 1, 4, 3, 1, 3,
	1, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 2,
	1, 1, 1, 1, 3, 3, 2, 4, 2, 3,
	1, 1, 2, 5, 4, 1, 1, 3, 2, 3,
	1, 3, 2, 3, 5, 1, 1, 1,
}

var yyChk = [...]int16{
	-1000, -1, -2, -6, -4, 48, 19, 5, -9, -15,
	6, 24, 20, 13, 11, 12, 15, 32, 25, -10,
	-17, -16, 37, 33, 48, -12, -13, 16, 10, 22,
	34, 31, -19, -15, -14, -22, 41, 17, 55, 12,
	-10, 35, 36, 49, 50, 53, 52, -18, 51, 37,
	-22, -14, -3, -1, -13, -3, -13, 33, -11, -7,
	-8, 33, 12, -11, 33, 33, 33, -13, -16, 50,
	18, 4, 38, 39, 29, 28, 26, 27, 30, 40,
	41, 42, 43, 44, 45, 47, -13, -13, -13, -20,
	37, 57, -23, -24, 33, 53, -13, -12, -10, -15,
	-13, 33, 33, 56, -12, 9, 6, 23, 21, 49,
	14, 50, -20, 51, 52, 33, 49, 32, 56, 56,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, -21, 56, 31,
	-11, 57, -25, 50, 48, 49, -13, 54, -18, 56,
	-3, -13, -3, -13, -12, 33, 33, 33, -20, -12,
	56, -3, 50, -24, -13, 54, 9, -5, 50, 6,
	-3, 9, 31, 49, 9, 7, 8, -13, -3, 9,
	-13, -3, -13, 6, 50, 9, 9, 21, -3, -13,
	-3, 9, 6, -3, 9,
}

var yyDef = [...]int8{
	4, -2, 1, 2, 5, 6, 26, 28, 0, 9,
	4, 0, 4, 0, 0, 0, 0, 0, 0, -2,
	72, 73, 0, 35, 3, 27, 40, 42, 43, 44,
	45, 46, 47, 48, 49, 50, 0, 0, 0, 0,
	71, 70, 0, 0, 0, 0, 0, 76, 0, 0,
	80, 81, 0, 7, 0, 0, 0, 38, 0, 0,
	29, 31, 0, 21, 38, 0, 23, 0, 73, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 67, 68, 69, 82,
	0, 88, 0, 90, 35, 0, 95, 8, -2, 0,
	0, 37, 0, 78, 0, 10, 4, 0, 4, 0,
	0, 0, 18, 0, 0, 0, 0, 22, 74, 75,
	41, 51, 52, 53, 54, 55, 56, 57, 58, 59,
	60, 61, 62, 63, 64, 65, 66, 0, 4, 85,
	86, 89, 92, 96, 97, 0, 0, 36, 77, 79,
	0, 12, 24, 0, 0, 39, 30, 32, 19, 20,
	4, 0, 0, 91, 93, 0, 11, 0, 0, 4,
	0, 84, 87, 0, 13, 4, 0, 0, 0, 83,
	94, 0, 0, 4, 0, 17, 14, 4, 0, 0,
	25, 15, 4, 0, 16,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 55, 3, 45, 3, 3,
	37, 56, 42, 40, 50, 41, 52, 43, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 51, 48,
	39, 49, 38, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 53, 3, 54, 47, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 36, 44, 57,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 46,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
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
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:73
		{
			yyVAL.stmts = yyDollar[1].stmts
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.stmts
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:79
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmt)
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.stmts
			}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:85
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmt)
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:93
		{
			yyVAL.stmts = []ast.Stmt{}
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:96
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmt)
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:99
		{
			yyVAL.stmts = yyDollar[1].stmts
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:104
		{
			yyVAL.stmts = yyDollar[1].stmts
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:109
		{
			yyVAL.stmt = &ast.AssignStmt{Lhs: yyDollar[1].exprlist, Rhs: yyDollar[3].exprlist}
			yyVAL.stmt.SetLine(yyDollar[1].exprlist[0].Line())
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:114
		{
			if _, ok := yyDollar[1].expr.(*ast.FuncCallExpr); !ok {
				yylex.(*Lexer).Error("parse error")
			} else {
				yyVAL.stmt = &ast.FuncCallStmt{Expr: yyDollar[1].expr}
				yyVAL.stmt.SetLine(yyDollar[1].expr.Line())
			}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:122
		{
			yyVAL.stmt = &ast.DoBlockStmt{Stmts: yyDollar[2].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[3].token.Pos.Line)
		}
	case 11:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:127
		{
			yyVAL.stmt = &ast.WhileStmt{Condition: yyDollar[2].expr, Stmts: yyDollar[4].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[5].token.Pos.Line)
		}
	case 12:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:132
		{
			yyVAL.stmt = &ast.RepeatStmt{Condition: yyDollar[4].expr, Stmts: yyDollar[2].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[4].expr.Line())
		}
	case 13:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:137
		{
			yyVAL.stmt = &ast.IfStmt{Condition: yyDollar[2].expr, Then: yyDollar[4].stmts}
			cur := yyVAL.stmt
			for _, elseif := range yyDollar[5].stmts {
				cur.(*ast.IfStmt).Else = []ast.Stmt{elseif}
				cur = elseif
			}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[6].token.Pos.Line)
		}
	case 14:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.go.y:147
		{
			yyVAL.stmt = &ast.IfStmt{Condition: yyDollar[2].expr, Then: yyDollar[4].stmts}
			cur := yyVAL.stmt
			for _, elseif := range yyDollar[5].stmts {
				cur.(*ast.IfStmt).Else = []ast.Stmt{elseif}
				cur = elseif
			}
			cur.(*ast.IfStmt).Else = yyDollar[7].stmts
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[8].token.Pos.Line)
		}
	case 15:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.go.y:158
		{
			yyVAL.stmt = &ast.NumberForStmt{Name: yyDollar[2].token.Str, Init: yyDollar[4].expr, Limit: yyDollar[6].expr, Stmts: yyDollar[8].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[9].token.Pos.Line)
		}
	case 16:
		yyDollar = yyS[yypt-11 : yypt+1]
//line parser.go.y:163
		{
			yyVAL.stmt = &ast.NumberForStmt{Name: yyDollar[2].token.Str, Init: yyDollar[4].expr, Limit: yyDollar[6].expr, Step: yyDollar[8].expr, Stmts: yyDollar[10].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[11].token.Pos.Line)
		}
	case 17:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.go.y:168
		{
			yyVAL.stmt = &ast.GenericForStmt{Names: yyDollar[2].namelist, Exprs: yyDollar[4].exprlist, Stmts: yyDollar[6].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[7].token.Pos.Line)
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:173
		{
			yyVAL.stmt = &ast.FuncDefStmt{Name: yyDollar[2].funcname, Func: yyDollar[3].funcexpr}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[3].funcexpr.LastLine())
		}
	case 19:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:178
		{
			yyVAL.stmt = &ast.LocalAssignStmt{Names: []string{yyDollar[3].token.Str}, Exprs: []ast.Expr{yyDollar[4].funcexpr}}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[4].funcexpr.LastLine())
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:183
		{
			yyVAL.stmt = &ast.LocalAssignStmt{Names: yyDollar[2].namelist, Exprs: yyDollar[4].exprlist}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:187
		{
			yyVAL.stmt = &ast.LocalAssignStmt{Names: yyDollar[2].namelist, Exprs: []ast.Expr{}}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:191
		{
			yyVAL.stmt = &ast.LabelStmt{Name: yyDollar[2].token.Str}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:195
		{
			yyVAL.stmt = &ast.GotoStmt{Label: yyDollar[2].token.Str}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:201
		{
			yyVAL.stmts = []ast.Stmt{}
		}
	case 25:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:204
		{
			yyVAL.stmts = append(yyDollar[1].stmts, &ast.IfStmt{Condition: yyDollar[3].expr, Then: yyDollar[5].stmts})
			yyVAL.stmts[len(yyVAL.stmts)-1].SetLine(yyDollar[2].token.Pos.Line)
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:210
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: nil}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:214
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprlist}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:218
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:224
		{
			yyVAL.funcname = yyDollar[1].funcname
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:227
		{
			yyVAL.funcname = &ast.FuncName{Func: nil, Receiver: yyDollar[1].funcname.Func, Method: yyDollar[3].token.Str}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:232
		{
			yyVAL.funcname = &ast.FuncName{Func: &ast.IdentExpr{Value: yyDollar[1].token.Str}}
			yyVAL.funcname.Func.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:236
		{
			key := &ast.StringExpr{Value: yyDollar[3].token.Str}
			key.SetLine(yyDollar[3].token.Pos.Line)
			fn := &ast.AttrGetExpr{Object: yyDollar[1].funcname.Func, Key: key}
			fn.SetLine(yyDollar[3].token.Pos.Line)
			yyVAL.funcname = &ast.FuncName{Func: fn}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:245
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:248
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:253
		{
			yyVAL.expr = &ast.IdentExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 36:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:257
		{
			yyVAL.expr = &ast.AttrGetExpr{Object: yyDollar[1].expr, Key: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:261
		{
			key := &ast.StringExpr{Value: yyDollar[3].token.Str}
			key.SetLine(yyDollar[3].token.Pos.Line)
			yyVAL.expr = &ast.AttrGetExpr{Object: yyDollar[1].expr, Key: key}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:269
		{
			yyVAL.namelist = []string{yyDollar[1].token.Str}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:272
		{
			yyVAL.namelist = append(yyDollar[1].namelist, yyDollar[3].token.Str)
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:277
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:280
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:285
		{
			yyVAL.expr = &ast.NilExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:289
		{
			yyVAL.expr = &ast.FalseExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:293
		{
			yyVAL.expr = &ast.TrueExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:297
		{
			yyVAL.expr = &ast.NumberExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:301
		{
			yyVAL.expr = &ast.Comma3Expr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:305
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:308
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:311
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:314
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:317
		{
			yyVAL.expr = &ast.LogicalOpExpr{Lhs: yyDollar[1].expr, Operator: "or", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:321
		{
			yyVAL.expr = &ast.LogicalOpExpr{Lhs: yyDollar[1].expr, Operator: "and", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:325
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:329
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:333
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:337
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:341
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:345
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "~=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:349
		{
			yyVAL.expr = &ast.StringConcatOpExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:353
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:357
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:361
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:365
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:369
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:373
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:377
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "^", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:381
		{
			yyVAL.expr = &ast.UnaryMinusOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:385
		{
			yyVAL.expr = &ast.UnaryNotOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:389
		{
			yyVAL.expr = &ast.UnaryLenOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:395
		{
			yyVAL.expr = &ast.StringExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:401
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:404
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:407
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:410
		{
			if ex, ok := yyDollar[2].expr.(*ast.Comma3Expr); ok {
				ex.AdjustRet = true
			}
			yyVAL.expr = yyDollar[2].expr
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:419
		{
			yyDollar[2].expr.(*ast.FuncCallExpr).AdjustRet = true
			yyVAL.expr = yyDollar[2].expr
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:425
		{
			yyVAL.expr = &ast.FuncCallExpr{Func: yyDollar[1].expr, Args: yyDollar[2].exprlist}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:429
		{
			yyVAL.expr = &ast.FuncCallExpr{Method: yyDollar[3].token.Str, Receiver: yyDollar[1].expr, Args: yyDollar[4].exprlist}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:435
		{
			if yylex.(*Lexer).PNewLine {
				yylex.(*Lexer).TokenError(yyDollar[1].token, "ambiguous syntax (function call x new statement)")
			}
			yyVAL.exprlist = []ast.Expr{}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:441
		{
			if yylex.(*Lexer).PNewLine {
				yylex.(*Lexer).TokenError(yyDollar[1].token, "ambiguous syntax (function call x new statement)")
			}
			yyVAL.exprlist = yyDollar[2].exprlist
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:447
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:450
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:455
		{
			yyVAL.expr = &ast.FunctionExpr{ParList: yyDollar[2].funcexpr.ParList, Stmts: yyDollar[2].funcexpr.Stmts}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.expr.SetLastLine(yyDollar[2].funcexpr.LastLine())
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:462
		{
			yyVAL.funcexpr = &ast.FunctionExpr{ParList: yyDollar[2].parlist, Stmts: yyDollar[4].stmts}
			yyVAL.funcexpr.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.funcexpr.SetLastLine(yyDollar[5].token.Pos.Line)
		}
	case 84:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:467
		{
			yyVAL.funcexpr = &ast.FunctionExpr{ParList: &ast.ParList{HasVargs: false, Names: []string{}}, Stmts: yyDollar[3].stmts}
			yyVAL.funcexpr.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.funcexpr.SetLastLine(yyDollar[4].token.Pos.Line)
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:474
		{
			yyVAL.parlist = &ast.ParList{HasVargs: true, Names: []string{}}
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:477
		{
			yyVAL.parlist = &ast.ParList{HasVargs: false, Names: []string{}}
			yyVAL.parlist.Names = append(yyVAL.parlist.Names, yyDollar[1].namelist...)
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:481
		{
			yyVAL.parlist = &ast.ParList{HasVargs: true, Names: []string{}}
			yyVAL.parlist.Names = append(yyVAL.parlist.Names, yyDollar[1].namelist...)
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:488
		{
			yyVAL.expr = &ast.TableExpr{Fields: []*ast.Field{}}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:492
		{
			yyVAL.expr = &ast.TableExpr{Fields: yyDollar[2].fieldlist}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:499
		{
			yyVAL.fieldlist = []*ast.Field{yyDollar[1].field}
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:502
		{
			yyVAL.fieldlist = append(yyDollar[1].fieldlist, yyDollar[3].field)
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:505
		{
			yyVAL.fieldlist = yyDollar[1].fieldlist
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:510
		{
			yyVAL.field = &ast.Field{Key: &ast.StringExpr{Value: yyDollar[1].token.Str}, Value: yyDollar[3].expr}
			yyVAL.field.Key.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 94:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:514
		{
			yyVAL.field = &ast.Field{Key: yyDollar[2].expr, Value: yyDollar[5].expr}
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:517
		{
			yyVAL.field = &ast.Field{Value: yyDollar[1].expr}
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:522
		{
			yyVAL.fieldsep = ","
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:525
		{
			yyVAL.fieldsep = ";"
		}
	}
	goto yystack /* stack new state and value */
}
