//line ./parser/parser.y:1
package parser

import __yyfmt__ "fmt"

//line ./parser/parser.y:3
import (
	"github.com/covrom/gonec/ast"
)

//line ./parser/parser.y:29
type yySymType struct {
	yys          int
	compstmt     []ast.Stmt
	modules      []ast.Stmt
	module       ast.Stmt
	stmt_if      ast.Stmt
	stmt_default ast.Stmt
	stmt_elsif   ast.Stmt
	stmt_elsifs  []ast.Stmt
	stmt_case    ast.Stmt
	stmt_cases   []ast.Stmt
	stmts        []ast.Stmt
	stmt         ast.Stmt
	typ          ast.Type
	expr         ast.Expr
	exprs        []ast.Expr
	expr_many    []ast.Expr
	expr_pair    ast.Expr
	expr_pairs   []ast.Expr
	expr_idents  []int
	tok          ast.Token
	term         ast.Token
	terms        ast.Token
	opt_terms    ast.Token
}

const IDENT = 57346
const NUMBER = 57347
const STRING = 57348
const ARRAY = 57349
const VARARG = 57350
const FUNC = 57351
const RETURN = 57352
const VAR = 57353
const THROW = 57354
const IF = 57355
const ELSE = 57356
const FOR = 57357
const IN = 57358
const EQEQ = 57359
const NEQ = 57360
const GE = 57361
const LE = 57362
const OROR = 57363
const ANDAND = 57364
const TRUE = 57365
const FALSE = 57366
const NIL = 57367
const MODULE = 57368
const TRY = 57369
const CATCH = 57370
const FINALLY = 57371
const PLUSEQ = 57372
const MINUSEQ = 57373
const MULEQ = 57374
const DIVEQ = 57375
const ANDEQ = 57376
const OREQ = 57377
const BREAK = 57378
const CONTINUE = 57379
const PLUSPLUS = 57380
const MINUSMINUS = 57381
const POW = 57382
const SHIFTLEFT = 57383
const SHIFTRIGHT = 57384
const SWITCH = 57385
const CASE = 57386
const DEFAULT = 57387
const GO = 57388
const CHAN = 57389
const MAKE = 57390
const OPCHAN = 57391
const ARRAYLIT = 57392
const NULL = 57393
const EACH = 57394
const TO = 57395
const ELSIF = 57396
const WHILE = 57397
const TERNARY = 57398
const TYPECAST = 57399
const UNARY = 57400

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"NUMBER",
	"STRING",
	"ARRAY",
	"VARARG",
	"FUNC",
	"RETURN",
	"VAR",
	"THROW",
	"IF",
	"ELSE",
	"FOR",
	"IN",
	"EQEQ",
	"NEQ",
	"GE",
	"LE",
	"OROR",
	"ANDAND",
	"TRUE",
	"FALSE",
	"NIL",
	"MODULE",
	"TRY",
	"CATCH",
	"FINALLY",
	"PLUSEQ",
	"MINUSEQ",
	"MULEQ",
	"DIVEQ",
	"ANDEQ",
	"OREQ",
	"BREAK",
	"CONTINUE",
	"PLUSPLUS",
	"MINUSMINUS",
	"POW",
	"SHIFTLEFT",
	"SHIFTRIGHT",
	"SWITCH",
	"CASE",
	"DEFAULT",
	"GO",
	"CHAN",
	"MAKE",
	"OPCHAN",
	"ARRAYLIT",
	"NULL",
	"EACH",
	"TO",
	"ELSIF",
	"WHILE",
	"TERNARY",
	"TYPECAST",
	"'='",
	"'?'",
	"':'",
	"','",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"UNARY",
	"'{'",
	"'}'",
	"'.'",
	"'!'",
	"'^'",
	"'&'",
	"')'",
	"'('",
	"'['",
	"']'",
	"'|'",
	"';'",
	"'\\n'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line ./parser/parser.y:758

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 6,
	1, 7,
	26, 7,
	-2, 131,
	-1, 12,
	61, 50,
	-2, 5,
	-1, 17,
	61, 51,
	-2, 25,
	-1, 26,
	28, 7,
	-2, 131,
	-1, 53,
	61, 50,
	-2, 132,
	-1, 103,
	1, 59,
	8, 59,
	14, 59,
	26, 59,
	28, 59,
	44, 59,
	45, 59,
	53, 59,
	54, 59,
	58, 59,
	60, 59,
	61, 59,
	70, 59,
	71, 59,
	76, 59,
	79, 59,
	81, 59,
	82, 59,
	-2, 54,
	-1, 105,
	1, 61,
	8, 61,
	14, 61,
	26, 61,
	28, 61,
	44, 61,
	45, 61,
	53, 61,
	54, 61,
	58, 61,
	60, 61,
	61, 61,
	70, 61,
	71, 61,
	76, 61,
	79, 61,
	81, 61,
	82, 61,
	-2, 54,
	-1, 137,
	17, 0,
	18, 0,
	-2, 87,
	-1, 138,
	17, 0,
	18, 0,
	-2, 88,
	-1, 158,
	61, 51,
	-2, 45,
	-1, 162,
	71, 7,
	-2, 131,
	-1, 163,
	71, 7,
	-2, 131,
	-1, 189,
	14, 7,
	54, 7,
	71, 7,
	-2, 131,
	-1, 236,
	61, 52,
	-2, 46,
	-1, 237,
	1, 47,
	14, 47,
	26, 47,
	28, 47,
	44, 47,
	45, 47,
	54, 47,
	58, 47,
	61, 53,
	71, 47,
	81, 47,
	82, 47,
	-2, 54,
	-1, 245,
	1, 53,
	8, 53,
	14, 53,
	26, 53,
	28, 53,
	44, 53,
	45, 53,
	54, 53,
	61, 53,
	71, 53,
	76, 53,
	79, 53,
	81, 53,
	82, 53,
	-2, 54,
	-1, 258,
	71, 7,
	-2, 131,
	-1, 268,
	1, 108,
	8, 108,
	14, 108,
	26, 108,
	28, 108,
	44, 108,
	45, 108,
	53, 108,
	54, 108,
	58, 108,
	60, 108,
	61, 108,
	70, 108,
	71, 108,
	76, 108,
	79, 108,
	81, 108,
	82, 108,
	-2, 106,
	-1, 270,
	1, 112,
	8, 112,
	14, 112,
	26, 112,
	28, 112,
	44, 112,
	45, 112,
	53, 112,
	54, 112,
	58, 112,
	60, 112,
	61, 112,
	70, 112,
	71, 112,
	76, 112,
	79, 112,
	81, 112,
	82, 112,
	-2, 110,
	-1, 277,
	71, 7,
	-2, 131,
	-1, 280,
	44, 7,
	45, 7,
	71, 7,
	-2, 131,
	-1, 284,
	71, 7,
	-2, 131,
	-1, 285,
	71, 7,
	-2, 131,
	-1, 290,
	1, 107,
	8, 107,
	14, 107,
	26, 107,
	28, 107,
	44, 107,
	45, 107,
	53, 107,
	54, 107,
	58, 107,
	60, 107,
	61, 107,
	70, 107,
	71, 107,
	76, 107,
	79, 107,
	81, 107,
	82, 107,
	-2, 105,
	-1, 291,
	1, 111,
	8, 111,
	14, 111,
	26, 111,
	28, 111,
	44, 111,
	45, 111,
	53, 111,
	54, 111,
	58, 111,
	60, 111,
	61, 111,
	70, 111,
	71, 111,
	76, 111,
	79, 111,
	81, 111,
	82, 111,
	-2, 109,
	-1, 295,
	71, 7,
	-2, 131,
	-1, 299,
	71, 7,
	-2, 131,
	-1, 300,
	44, 7,
	45, 7,
	71, 7,
	-2, 131,
	-1, 306,
	71, 7,
	-2, 131,
	-1, 316,
	14, 7,
	54, 7,
	71, 7,
	-2, 131,
}

const yyPrivate = 57344

const yyLast = 3104

var yyAct = [...]int{

	90, 178, 55, 10, 204, 205, 18, 8, 9, 165,
	98, 99, 183, 17, 263, 184, 224, 187, 50, 181,
	99, 109, 222, 91, 126, 175, 94, 119, 96, 291,
	95, 100, 101, 102, 104, 106, 8, 9, 126, 260,
	89, 107, 8, 9, 269, 112, 114, 290, 118, 286,
	121, 259, 123, 218, 17, 267, 210, 192, 127, 253,
	129, 130, 131, 132, 133, 134, 135, 136, 137, 138,
	139, 140, 141, 142, 143, 144, 145, 146, 147, 148,
	240, 179, 149, 150, 151, 152, 183, 154, 156, 158,
	295, 116, 12, 318, 108, 157, 317, 159, 315, 313,
	168, 206, 207, 312, 153, 309, 52, 159, 159, 159,
	159, 173, 270, 206, 207, 122, 303, 167, 185, 252,
	186, 117, 265, 268, 211, 193, 158, 249, 250, 176,
	297, 248, 190, 226, 93, 110, 111, 161, 125, 88,
	203, 126, 115, 7, 206, 207, 289, 296, 15, 163,
	11, 3, 198, 261, 14, 219, 196, 179, 54, 239,
	6, 229, 199, 221, 216, 215, 200, 201, 53, 174,
	214, 208, 209, 217, 202, 160, 128, 118, 220, 56,
	5, 2, 92, 4, 177, 230, 275, 294, 235, 236,
	166, 120, 23, 238, 13, 1, 241, 54, 244, 246,
	227, 228, 124, 0, 0, 0, 0, 251, 0, 0,
	0, 0, 0, 0, 254, 188, 0, 0, 0, 191,
	0, 0, 0, 0, 0, 0, 0, 266, 0, 0,
	0, 0, 0, 272, 0, 273, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 278, 0,
	0, 0, 197, 0, 0, 0, 0, 166, 282, 0,
	0, 0, 283, 244, 0, 0, 288, 0, 0, 223,
	225, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 298, 0, 0, 301, 0, 0, 0, 304, 305,
	0, 0, 0, 0, 0, 0, 0, 0, 308, 307,
	0, 0, 0, 310, 311, 0, 0, 0, 0, 0,
	314, 258, 0, 0, 0, 262, 0, 264, 0, 0,
	319, 0, 0, 28, 29, 35, 0, 0, 41, 21,
	16, 22, 51, 0, 24, 0, 0, 0, 0, 0,
	0, 0, 36, 37, 38, 280, 26, 0, 0, 0,
	0, 0, 284, 285, 0, 19, 20, 0, 0, 0,
	0, 0, 27, 0, 0, 45, 0, 46, 49, 47,
	39, 0, 300, 0, 25, 40, 48, 0, 0, 306,
	0, 0, 0, 0, 30, 34, 0, 0, 0, 43,
	0, 0, 31, 32, 33, 0, 44, 42, 0, 0,
	8, 9, 67, 68, 70, 72, 82, 84, 0, 0,
	0, 0, 0, 0, 0, 73, 74, 75, 76, 77,
	78, 0, 0, 79, 80, 64, 65, 66, 0, 0,
	0, 0, 0, 0, 87, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 234, 69, 71, 59,
	60, 61, 62, 63, 0, 0, 0, 58, 0, 0,
	83, 233, 85, 86, 0, 81, 67, 68, 70, 72,
	82, 84, 0, 0, 0, 0, 0, 0, 0, 73,
	74, 75, 76, 77, 78, 0, 0, 79, 80, 64,
	65, 66, 0, 0, 0, 0, 0, 0, 87, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	232, 69, 71, 59, 60, 61, 62, 63, 0, 0,
	0, 58, 0, 0, 83, 231, 85, 86, 0, 81,
	67, 68, 70, 72, 82, 84, 0, 0, 0, 0,
	0, 0, 0, 73, 74, 75, 76, 77, 78, 0,
	0, 79, 80, 64, 65, 66, 0, 0, 0, 0,
	0, 0, 87, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 213, 0, 69, 71, 59, 60, 61,
	62, 63, 0, 0, 0, 58, 0, 0, 83, 0,
	85, 86, 212, 81, 67, 68, 70, 72, 82, 84,
	0, 0, 0, 0, 0, 0, 0, 73, 74, 75,
	76, 77, 78, 0, 0, 79, 80, 64, 65, 66,
	0, 0, 0, 0, 0, 0, 87, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 195, 0, 69,
	71, 59, 60, 61, 62, 63, 0, 0, 0, 58,
	0, 0, 83, 0, 85, 86, 194, 81, 67, 68,
	70, 72, 82, 84, 0, 0, 0, 0, 0, 0,
	0, 73, 74, 75, 76, 77, 78, 0, 0, 79,
	80, 64, 65, 66, 0, 0, 0, 0, 0, 0,
	87, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 71, 59, 60, 61, 62, 63,
	0, 316, 0, 58, 0, 0, 83, 0, 85, 86,
	0, 81, 67, 68, 70, 72, 82, 84, 0, 0,
	0, 0, 0, 0, 0, 73, 74, 75, 76, 77,
	78, 0, 0, 79, 80, 64, 65, 66, 0, 0,
	0, 0, 0, 0, 87, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 71, 59,
	60, 61, 62, 63, 0, 0, 0, 58, 0, 0,
	83, 302, 85, 86, 0, 81, 67, 68, 70, 72,
	82, 84, 0, 0, 0, 0, 0, 0, 0, 73,
	74, 75, 76, 77, 78, 0, 0, 79, 80, 64,
	65, 66, 0, 0, 0, 0, 0, 0, 87, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 71, 59, 60, 61, 62, 63, 0, 299,
	0, 58, 0, 0, 83, 0, 85, 86, 0, 81,
	67, 68, 70, 72, 82, 84, 0, 0, 0, 0,
	0, 0, 0, 73, 74, 75, 76, 77, 78, 0,
	0, 79, 80, 64, 65, 66, 0, 0, 0, 0,
	0, 0, 87, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 71, 59, 60, 61,
	62, 63, 0, 0, 0, 58, 0, 0, 83, 293,
	85, 86, 0, 81, 67, 68, 70, 72, 82, 84,
	0, 0, 0, 0, 0, 0, 0, 73, 74, 75,
	76, 77, 78, 0, 0, 79, 80, 64, 65, 66,
	0, 0, 0, 0, 0, 0, 87, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	71, 59, 60, 61, 62, 63, 0, 0, 0, 58,
	0, 0, 83, 292, 85, 86, 0, 81, 67, 68,
	70, 72, 82, 84, 0, 0, 0, 0, 0, 0,
	0, 73, 74, 75, 76, 77, 78, 0, 0, 79,
	80, 64, 65, 66, 0, 0, 0, 0, 0, 0,
	87, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 71, 59, 60, 61, 62, 63,
	0, 0, 0, 58, 0, 0, 83, 0, 85, 86,
	281, 81, 67, 68, 70, 72, 82, 84, 0, 0,
	0, 0, 0, 0, 0, 73, 74, 75, 76, 77,
	78, 0, 0, 79, 80, 64, 65, 66, 0, 0,
	0, 0, 0, 0, 87, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 279, 0, 69, 71, 59,
	60, 61, 62, 63, 0, 0, 0, 58, 0, 0,
	83, 0, 85, 86, 0, 81, 67, 68, 70, 72,
	82, 84, 0, 0, 0, 0, 0, 0, 0, 73,
	74, 75, 76, 77, 78, 0, 0, 79, 80, 64,
	65, 66, 0, 0, 0, 0, 0, 0, 87, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 71, 59, 60, 61, 62, 63, 0, 277,
	0, 58, 0, 0, 83, 0, 85, 86, 0, 81,
	67, 68, 70, 72, 82, 84, 0, 0, 0, 0,
	0, 0, 0, 73, 74, 75, 76, 77, 78, 0,
	0, 79, 80, 64, 65, 66, 0, 0, 0, 0,
	0, 0, 87, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 71, 59, 60, 61,
	62, 63, 0, 0, 0, 58, 0, 0, 83, 0,
	85, 86, 276, 81, 67, 68, 70, 72, 82, 84,
	0, 0, 0, 0, 0, 0, 0, 73, 74, 75,
	76, 77, 78, 0, 0, 79, 80, 64, 65, 66,
	0, 0, 0, 0, 0, 0, 87, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	71, 59, 60, 61, 62, 63, 0, 0, 0, 58,
	0, 0, 83, 274, 85, 86, 0, 81, 67, 68,
	70, 72, 82, 84, 0, 0, 0, 0, 0, 0,
	0, 73, 74, 75, 76, 77, 78, 0, 0, 79,
	80, 64, 65, 66, 0, 0, 0, 0, 0, 0,
	87, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 71, 59, 60, 61, 62, 63,
	0, 0, 0, 58, 0, 0, 83, 271, 85, 86,
	0, 81, 67, 68, 70, 72, 82, 84, 0, 0,
	0, 0, 0, 0, 0, 73, 74, 75, 76, 77,
	78, 0, 0, 79, 80, 64, 65, 66, 0, 0,
	0, 0, 0, 0, 87, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 257, 69, 71, 59,
	60, 61, 62, 63, 0, 0, 0, 58, 0, 0,
	83, 0, 85, 86, 0, 81, 67, 68, 70, 72,
	82, 84, 0, 0, 0, 0, 0, 0, 0, 73,
	74, 75, 76, 77, 78, 0, 0, 79, 80, 64,
	65, 66, 0, 0, 0, 0, 0, 0, 87, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 71, 59, 60, 61, 62, 63, 0, 0,
	0, 58, 0, 0, 83, 0, 85, 86, 256, 81,
	67, 68, 70, 72, 82, 84, 0, 0, 0, 0,
	0, 0, 0, 73, 74, 75, 76, 77, 78, 0,
	0, 79, 80, 64, 65, 66, 0, 0, 0, 0,
	0, 0, 87, 0, 0, 0, 247, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 71, 59, 60, 61,
	62, 63, 0, 0, 0, 58, 0, 0, 83, 0,
	85, 86, 0, 81, 67, 68, 70, 72, 82, 84,
	0, 0, 0, 0, 0, 0, 0, 73, 74, 75,
	76, 77, 78, 0, 0, 79, 80, 64, 65, 66,
	0, 0, 0, 0, 0, 0, 87, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	71, 59, 60, 61, 62, 63, 0, 0, 0, 58,
	0, 0, 83, 0, 85, 86, 243, 81, 67, 68,
	70, 72, 82, 84, 0, 0, 0, 0, 0, 0,
	0, 73, 74, 75, 76, 77, 78, 0, 0, 79,
	80, 64, 65, 66, 0, 0, 0, 0, 0, 0,
	87, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 71, 59, 60, 61, 62, 63,
	0, 189, 0, 58, 0, 0, 83, 0, 85, 86,
	0, 81, 67, 68, 70, 72, 82, 84, 0, 0,
	0, 0, 0, 0, 0, 73, 74, 75, 76, 77,
	78, 0, 0, 79, 80, 64, 65, 66, 0, 0,
	0, 0, 0, 0, 87, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 71, 59,
	60, 61, 62, 63, 0, 0, 0, 58, 0, 0,
	83, 180, 85, 86, 0, 81, 67, 68, 70, 72,
	82, 84, 0, 0, 0, 0, 0, 0, 0, 73,
	74, 75, 76, 77, 78, 0, 0, 79, 80, 64,
	65, 66, 0, 0, 0, 0, 0, 0, 87, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	172, 69, 71, 59, 60, 61, 62, 63, 0, 0,
	0, 58, 0, 0, 83, 0, 85, 86, 0, 81,
	67, 68, 70, 72, 82, 84, 0, 0, 0, 0,
	0, 0, 0, 73, 74, 75, 76, 77, 78, 0,
	0, 79, 80, 64, 65, 66, 0, 0, 0, 0,
	0, 0, 87, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 164, 0, 69, 71, 59, 60, 61,
	62, 63, 0, 0, 0, 58, 0, 0, 83, 0,
	85, 86, 0, 81, 67, 68, 70, 72, 82, 84,
	0, 0, 0, 0, 0, 0, 0, 73, 74, 75,
	76, 77, 78, 0, 0, 79, 80, 64, 65, 66,
	0, 0, 0, 0, 0, 0, 87, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	71, 59, 60, 61, 62, 63, 0, 162, 0, 58,
	0, 0, 83, 0, 85, 86, 0, 81, 67, 68,
	70, 72, 82, 84, 0, 0, 0, 0, 0, 0,
	0, 73, 74, 75, 76, 77, 78, 0, 0, 79,
	80, 64, 65, 66, 0, 0, 0, 0, 0, 0,
	87, 0, 0, 0, 0, 0, 0, 0, 0, 57,
	0, 0, 0, 69, 71, 59, 60, 61, 62, 63,
	0, 0, 0, 58, 0, 0, 83, 0, 85, 86,
	0, 81, 67, 68, 70, 72, 82, 84, 0, 0,
	0, 0, 0, 0, 0, 73, 74, 75, 76, 77,
	78, 0, 0, 79, 80, 64, 65, 66, 0, 0,
	0, 0, 0, 0, 87, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 71, 59,
	60, 61, 62, 63, 0, 0, 0, 58, 0, 0,
	83, 0, 85, 86, 0, 81, 67, 68, 70, 72,
	82, 84, 0, 0, 0, 0, 0, 0, 0, 73,
	74, 75, 76, 77, 78, 0, 0, 79, 80, 64,
	65, 66, 0, 0, 0, 0, 0, 0, 87, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 71, 59, 60, 61, 62, 63, 0, 0,
	0, 58, 0, 0, 83, 0, 182, 86, 0, 81,
	67, 68, 70, 72, 82, 84, 0, 0, 0, 0,
	0, 0, 0, 73, 74, 75, 76, 77, 78, 0,
	0, 79, 80, 64, 65, 66, 0, 0, 0, 0,
	0, 0, 87, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 71, 59, 60, 61,
	62, 63, 0, 0, 0, 171, 0, 0, 83, 0,
	85, 86, 0, 81, 67, 68, 70, 72, 82, 84,
	0, 0, 0, 0, 0, 0, 0, 73, 74, 75,
	76, 77, 78, 0, 0, 79, 80, 64, 65, 66,
	0, 0, 0, 0, 0, 0, 87, 28, 29, 35,
	0, 0, 41, 21, 16, 22, 51, 0, 24, 69,
	71, 59, 60, 61, 62, 63, 36, 37, 38, 170,
	26, 0, 83, 0, 85, 86, 0, 81, 0, 19,
	20, 0, 0, 0, 0, 0, 27, 0, 0, 45,
	0, 46, 49, 47, 39, 0, 0, 0, 25, 40,
	48, 0, 0, 0, 0, 0, 0, 0, 30, 34,
	0, 0, 0, 43, 0, 0, 31, 32, 33, 0,
	44, 42, 67, 68, 70, 72, 0, 84, 0, 0,
	0, 0, 0, 0, 0, 73, 74, 75, 76, 77,
	78, 0, 0, 79, 80, 64, 65, 66, 0, 0,
	0, 0, 0, 0, 87, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 71, 59,
	60, 61, 62, 63, 0, 0, 0, 58, 0, 0,
	83, 0, 85, 86, 0, 81, 67, 68, 70, 72,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 73,
	74, 75, 76, 77, 78, 0, 0, 79, 80, 64,
	65, 66, 0, 0, 0, 0, 0, 0, 87, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 71, 59, 60, 61, 62, 63, 0, 70,
	72, 58, 0, 0, 83, 0, 85, 86, 0, 81,
	73, 74, 75, 76, 77, 78, 0, 0, 79, 80,
	64, 65, 66, 0, 0, 0, 0, 0, 0, 87,
	245, 29, 35, 0, 0, 41, 0, 0, 0, 0,
	0, 0, 69, 71, 59, 60, 61, 62, 63, 36,
	37, 38, 58, 0, 0, 83, 0, 85, 86, 0,
	81, 0, 0, 0, 0, 28, 29, 35, 0, 0,
	41, 0, 45, 0, 46, 49, 47, 39, 0, 0,
	0, 0, 40, 48, 36, 37, 38, 0, 0, 0,
	0, 30, 34, 0, 0, 0, 43, 0, 0, 31,
	32, 33, 0, 44, 42, 287, 0, 45, 0, 46,
	49, 47, 39, 0, 0, 0, 0, 40, 48, 0,
	0, 0, 0, 28, 29, 35, 30, 34, 41, 0,
	0, 43, 0, 0, 31, 32, 33, 0, 44, 42,
	255, 0, 36, 37, 38, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 28, 29,
	35, 0, 0, 41, 0, 45, 0, 46, 49, 47,
	39, 0, 0, 0, 0, 40, 48, 36, 37, 38,
	0, 0, 0, 0, 30, 34, 0, 0, 0, 43,
	0, 0, 31, 32, 33, 0, 44, 42, 242, 0,
	45, 0, 46, 49, 47, 39, 0, 0, 0, 0,
	40, 48, 0, 0, 169, 0, 28, 29, 35, 30,
	34, 41, 0, 0, 43, 0, 0, 31, 32, 33,
	0, 44, 42, 0, 0, 36, 37, 38, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 28, 29, 35, 0, 45, 41,
	46, 49, 47, 39, 0, 0, 0, 0, 40, 48,
	0, 0, 155, 36, 37, 38, 0, 30, 34, 0,
	0, 0, 43, 0, 0, 31, 32, 33, 0, 44,
	42, 0, 28, 29, 35, 0, 45, 41, 46, 49,
	47, 39, 0, 0, 0, 0, 40, 48, 0, 0,
	97, 36, 37, 38, 0, 30, 34, 0, 0, 0,
	43, 0, 0, 31, 32, 33, 0, 44, 42, 0,
	245, 29, 35, 0, 45, 41, 46, 49, 47, 39,
	0, 0, 0, 0, 40, 48, 0, 0, 0, 36,
	37, 38, 0, 30, 34, 0, 0, 0, 43, 0,
	0, 31, 32, 33, 0, 44, 42, 0, 237, 29,
	35, 0, 45, 41, 46, 49, 47, 39, 0, 0,
	0, 0, 40, 48, 0, 0, 0, 36, 37, 38,
	0, 30, 34, 0, 0, 0, 43, 0, 0, 31,
	32, 33, 0, 44, 42, 0, 0, 0, 0, 0,
	45, 0, 46, 49, 47, 39, 0, 0, 0, 0,
	40, 48, 0, 0, 0, 0, 0, 0, 0, 30,
	34, 0, 0, 0, 43, 0, 0, 31, 32, 33,
	0, 44, 42, 73, 74, 75, 76, 77, 78, 0,
	0, 79, 80, 64, 113, 29, 35, 0, 0, 41,
	0, 0, 87, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 36, 37, 38, 0, 59, 60, 61,
	62, 63, 0, 0, 0, 58, 0, 0, 83, 0,
	85, 86, 0, 81, 0, 0, 45, 0, 46, 49,
	47, 39, 0, 0, 0, 0, 40, 48, 0, 0,
	0, 0, 105, 29, 35, 30, 34, 41, 0, 0,
	43, 0, 0, 31, 32, 33, 0, 44, 42, 0,
	0, 36, 37, 38, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 103, 29, 35,
	0, 0, 41, 0, 45, 0, 46, 49, 47, 39,
	0, 0, 0, 0, 40, 48, 36, 37, 38, 0,
	0, 0, 0, 30, 34, 0, 0, 0, 43, 0,
	0, 31, 32, 33, 0, 44, 42, 0, 0, 45,
	0, 46, 49, 47, 39, 0, 0, 0, 0, 40,
	48, 0, 0, 0, 0, 0, 0, 0, 30, 34,
	0, 0, 0, 43, 0, 0, 31, 32, 33, 0,
	44, 42, 73, 74, 75, 76, 77, 78, 0, 0,
	0, 0, 64, 73, 74, 75, 76, 77, 78, 0,
	0, 87, 0, 64, 0, 0, 0, 0, 0, 0,
	0, 0, 87, 0, 0, 0, 0, 0, 61, 62,
	63, 0, 0, 0, 58, 0, 0, 83, 0, 85,
	86, 0, 81, 0, 0, 58, 0, 0, 83, 0,
	85, 86, 0, 81,
}
var yyPact = [...]int{

	125, 125, -1000, 176, -1000, -74, -74, -1000, -1000, -1000,
	-1000, -1000, 2223, -74, -74, -1000, 175, 1921, 81, -1000,
	-1000, 2708, 2708, -1000, 130, 2708, -74, 2670, -67, -1000,
	2708, 2708, 2708, 2963, 2928, -1000, -1000, -1000, -1000, -1000,
	2708, 17, -74, -74, 2708, 2870, 44, -50, 173, 2708,
	54, 2708, -1000, 319, -1000, 80, -1000, 2708, 172, 2708,
	2708, 2708, 2708, 2708, 2708, 2708, 2708, 2708, 2708, 2708,
	2708, 2708, 2708, 2708, 2708, 2708, 2708, 2708, 2708, -1000,
	-1000, 2708, 2708, 2708, 2708, 2708, 2632, 2708, 2708, 46,
	1985, 1985, 171, 79, 1857, 121, 1793, -74, 2708, 2574,
	3023, 3023, 3023, -67, 2177, -67, 2113, 1729, 165, -52,
	2708, 151, 1665, -58, 2049, 14, -62, 2708, -1000, 2708,
	-60, 1985, -74, 1601, -1000, 2708, -74, 1985, -1000, 3012,
	3012, 3023, 3023, 3023, 1985, 2833, 2833, 2400, 2400, 2833,
	2833, 2833, 2833, 1985, 1985, 1985, 1985, 1985, 1985, 1985,
	2285, 1985, 2349, 49, 577, 2708, 1985, -1000, 1985, -74,
	136, 2708, -74, -74, -74, 69, 100, 48, 513, 2708,
	161, 160, 2708, -23, 147, 159, -39, -45, -1000, 73,
	-1000, 2708, 2708, 157, 2708, 449, 385, 2708, 2784, -74,
	-1000, 155, 4, -1000, -1000, 2539, 1537, 2746, 2708, 1473,
	60, 56, 57, -1000, -1000, -1000, 2708, 59, -1000, -1000,
	-17, -1000, -1000, 2481, 1409, -1000, -1000, 1345, -74, -25,
	-37, 145, -74, -65, -74, 51, 2708, 47, 36, -1000,
	1281, -1000, 2708, -1000, 2708, 1217, 1985, -67, -1000, -1000,
	-1000, 1153, -1000, -1000, 1985, -67, 1089, 2708, -1000, -1000,
	-1000, 1025, -74, -1000, 961, -1000, -1000, 2708, -74, -74,
	-74, -27, 2446, -1000, 75, -1000, 1985, -29, -1000, -47,
	-1000, -1000, 897, 833, -1000, 76, -1000, -74, 769, -74,
	-74, -1000, 705, 45, -74, -74, -74, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -74, -1000, 2708, 34, -74,
	-74, -1000, -1000, -1000, 32, 28, -74, 27, 641, -1000,
	25, -1000, -1000, -1000, 22, -1000, -74, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 3, 195, 181, 194, 148, 192, 5, 4, 9,
	187, 186, 142, 0, 18, 6, 1, 184, 2, 154,
	92, 143,
}
var yyR1 = [...]int{

	0, 2, 2, 2, 3, 1, 1, 4, 4, 4,
	5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 11, 11, 10, 6,
	6, 9, 9, 9, 9, 9, 8, 7, 16, 17,
	17, 17, 18, 18, 18, 15, 15, 15, 12, 12,
	14, 14, 14, 14, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 20, 20, 19, 19, 21, 21,
}
var yyR2 = [...]int{

	0, 0, 1, 2, 4, 1, 2, 0, 2, 3,
	4, 2, 3, 3, 1, 1, 2, 2, 1, 8,
	9, 5, 5, 5, 4, 1, 0, 2, 4, 8,
	6, 0, 2, 2, 2, 2, 5, 4, 3, 0,
	1, 4, 0, 1, 4, 1, 4, 4, 1, 3,
	0, 1, 4, 4, 1, 1, 2, 2, 2, 2,
	4, 2, 4, 1, 1, 1, 1, 1, 7, 3,
	7, 8, 8, 9, 5, 6, 5, 6, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	2, 3, 3, 3, 3, 5, 4, 6, 5, 5,
	4, 6, 5, 4, 4, 6, 5, 5, 6, 5,
	5, 2, 2, 5, 4, 6, 5, 4, 6, 3,
	2, 0, 1, 1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -2, -3, 26, -3, 4, -19, -21, 81, 82,
	-1, -21, -20, -4, -19, -5, 11, -13, -15, 36,
	37, 10, 12, -6, 15, 55, 27, 43, 4, 5,
	65, 73, 74, 75, 66, 6, 23, 24, 25, 51,
	56, 9, 78, 70, 77, 46, 48, 50, 57, 49,
	-14, 13, -20, -19, -21, -18, 4, 58, 72, 64,
	65, 66, 67, 68, 40, 41, 42, 17, 18, 62,
	19, 63, 20, 30, 31, 32, 33, 34, 35, 38,
	39, 80, 21, 75, 22, 77, 78, 49, 58, -14,
	-13, -13, 52, 4, -13, -1, -13, 60, 77, 78,
	-13, -13, -13, 4, -13, 4, -13, -13, 77, 4,
	-20, -20, -13, 4, -13, -12, 47, 77, 4, 77,
	-12, -13, 61, -13, -5, 58, 61, -13, 4, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -14, -13, 60, -13, -15, -13, 61,
	4, 58, 70, 28, 60, -9, -20, -14, -13, 60,
	72, 72, 61, -18, 4, 77, -14, -17, -16, 6,
	76, 77, 77, 72, 77, -13, -13, 77, -20, 70,
	-15, -20, 8, 76, 79, 60, -13, -20, 16, -13,
	-1, -1, -9, 71, -8, -7, 44, 45, -8, -7,
	8, 76, 79, 60, -13, 4, 4, -13, 76, 8,
	-18, 4, 61, -20, 61, -20, 60, -14, -14, 4,
	-13, 76, 61, 76, 61, -13, -13, 4, -1, 4,
	76, -13, 79, 79, -13, 4, -13, 53, 71, 71,
	71, -13, 60, 76, -13, 79, 79, 61, -20, 76,
	76, 8, -20, 79, -20, 71, -13, 8, 76, 8,
	76, 76, -13, -13, 76, -11, 79, 70, -13, 60,
	-20, 79, -13, -1, -20, -20, 76, 79, -16, 71,
	76, 76, 76, 76, -10, 14, 71, 54, -1, 70,
	-20, -1, 76, 71, -1, -1, -20, -1, -13, 71,
	-1, -1, 71, 71, -1, 71, 70, 71, 71, -1,
}
var yyDef = [...]int{

	1, -2, 2, 0, 3, 0, -2, 133, 135, 136,
	4, 133, -2, 131, 132, 8, 42, -2, 0, 14,
	15, 50, 0, 18, 0, 0, -2, 0, 54, 55,
	0, 0, 0, 0, 0, 63, 64, 65, 66, 67,
	0, 0, 131, 131, 0, 0, 0, 0, 0, 0,
	0, 0, 6, -2, 134, 11, 43, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 99,
	100, 0, 0, 0, 0, 50, 0, 0, 50, 16,
	51, 17, 0, 0, 0, 0, 0, 31, 50, 0,
	56, 57, 58, -2, 0, -2, 0, 0, 42, 0,
	50, 39, 0, 54, 0, 121, 122, 0, 48, 0,
	0, 130, 131, 0, 9, 50, 131, 12, 69, 79,
	80, 81, 82, 83, 84, 85, 86, -2, -2, 89,
	90, 91, 92, 93, 94, 95, 96, 97, 98, 101,
	102, 103, 104, 0, 0, 0, 129, 13, -2, 131,
	0, 0, -2, -2, 31, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 43, 42, 131, 131, 40, 0,
	78, 50, 50, 0, 0, 0, 0, 0, 0, -2,
	10, 0, 0, 110, 114, 0, 0, 0, 0, 0,
	0, 0, 0, 24, 34, 35, 0, 0, 32, 33,
	0, 106, 113, 0, 0, 60, 62, 0, 131, 0,
	0, 43, 131, 0, 131, 0, 0, 0, 0, 49,
	0, 127, 0, 124, 0, 0, -2, -2, 26, 44,
	109, 0, 119, 120, 52, -2, 0, 0, 21, 22,
	23, 0, 131, 105, 0, 116, 117, 0, -2, 131,
	131, 0, 0, 74, 0, 76, 38, 0, -2, 0,
	-2, 123, 0, 0, 126, 0, 118, -2, 0, 131,
	-2, 115, 0, 0, -2, -2, 131, 75, 41, 77,
	-2, -2, 128, 125, 27, -2, 30, 0, 0, -2,
	-2, 37, 68, 70, 0, 0, -2, 0, 0, 19,
	0, 36, 71, 72, 0, 29, -2, 20, 73, 28,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	82, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 73, 3, 3, 3, 68, 75, 3,
	77, 76, 66, 64, 61, 65, 72, 67, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 60, 81,
	63, 58, 62, 59, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 78, 3, 79, 74, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 70, 80, 71,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 69,
}
var yyTok3 = [...]int{
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
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
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
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
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
	yyn = yyPact[yystate]
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
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
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
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
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
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
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

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.y:71
		{
			yyVAL.modules = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.modules
			}
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:78
		{
			yyVAL.modules = []ast.Stmt{yyDollar[1].module}
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.modules
			}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:85
		{
			if yyDollar[2].module != nil {
				yyVAL.modules = append(yyDollar[1].modules, yyDollar[2].module)
				if l, ok := yylex.(*Lexer); ok {
					l.stmts = yyVAL.modules
				}
			}
		}
	case 4:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:96
		{
			yyVAL.module = &ast.ModuleStmt{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), Stmts: yyDollar[4].compstmt}
			yyVAL.module.SetPosition(yyDollar[1].tok.Position())
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:102
		{
			yyVAL.compstmt = nil
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:106
		{
			yyVAL.compstmt = yyDollar[1].stmts
		}
	case 7:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.y:111
		{
			yyVAL.stmts = nil
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:115
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:119
		{
			if yyDollar[3].stmt != nil {
				yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[3].stmt)
			}
		}
	case 10:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:127
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].expr_many}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:132
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: nil}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:137
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "=", Rhss: []ast.Expr{yyDollar[3].expr}}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:141
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:145
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:150
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:155
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:160
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:165
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 19:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./parser/parser.y:170
		{
			yyVAL.stmt = &ast.ForStmt{Var: ast.UniqueNames.Set(yyDollar[3].tok.Lit), Value: yyDollar[5].expr, Stmts: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./parser/parser.y:175
		{
			yyVAL.stmt = &ast.NumForStmt{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), Expr1: yyDollar[4].expr, Expr2: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:180
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:185
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[2].compstmt, Catch: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:190
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:195
		{
			yyVAL.stmt = &ast.SelectStmt{Cases: yyDollar[3].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:200
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 26:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.y:206
		{
			yyVAL.stmt_elsifs = []ast.Stmt{}
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:210
		{
			yyVAL.stmt_elsifs = append(yyDollar[1].stmt_elsifs, yyDollar[2].stmt_elsif)
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:216
		{
			yyVAL.stmt_elsif = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt}
		}
	case 29:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./parser/parser.y:222
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, ElseIf: yyDollar[5].stmt_elsifs, Else: yyDollar[7].compstmt}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 30:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:227
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, ElseIf: yyDollar[5].stmt_elsifs, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 31:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.y:233
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:237
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_case}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:241
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_default}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:245
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_case)
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:249
		{
			for _, stmt := range yyDollar[1].stmt_cases {
				if _, ok := stmt.(*ast.DefaultStmt); ok {
					yylex.Error("multiple default statement")
				}
			}
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_default)
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:260
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[5].compstmt}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:266
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyDollar[4].compstmt}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:272
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.y:277
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:281
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:285
		{
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[4].expr_pair)
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.y:290
		{
			yyVAL.expr_idents = []int{}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:294
		{
			yyVAL.expr_idents = []int{ast.UniqueNames.Set(yyDollar[1].tok.Lit)}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:298
		{
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, ast.UniqueNames.Set(yyDollar[4].tok.Lit))
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:304
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 46:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:308
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:312
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[4].tok.Lit)})
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:317
		{
			yyVAL.typ = ast.Type{Name: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:321
		{
			yyVAL.typ = ast.Type{Name: ast.UniqueNames.Set(ast.UniqueNames.Get(yyDollar[1].typ.Name) + "." + yyDollar[3].tok.Lit)}
		}
	case 50:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.y:326
		{
			yyVAL.exprs = nil
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:330
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:334
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:338
		{
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[4].tok.Lit)})
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:344
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:349
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:354
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:359
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:364
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:369
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[2].tok.Lit)}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:374
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: ast.UniqueNames.Set(yyDollar[4].tok.Lit)}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:379
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[2].tok.Lit)}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:384
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: ast.UniqueNames.Set(yyDollar[4].tok.Lit)}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:389
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:394
		{
			yyVAL.expr = &ast.ConstExpr{Value: "истина"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:399
		{
			yyVAL.expr = &ast.ConstExpr{Value: "ложь"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:404
		{
			yyVAL.expr = &ast.ConstExpr{Value: "неопределено"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:409
		{
			yyVAL.expr = &ast.ConstExpr{Value: "null"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 68:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./parser/parser.y:414
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[2].expr, Lhs: yyDollar[4].expr, Rhs: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:419
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: ast.UniqueNames.Set(yyDollar[3].tok.Lit)}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 70:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./parser/parser.y:424
		{
			yyVAL.expr = &ast.FuncExpr{Name: ast.UniqueNames.Set("<анонимная функция>"), Args: yyDollar[3].expr_idents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 71:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./parser/parser.y:429
		{
			yyVAL.expr = &ast.FuncExpr{Name: ast.UniqueNames.Set("<анонимная функция>"), Args: []int{ast.UniqueNames.Set(yyDollar[3].tok.Lit)}, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 72:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./parser/parser.y:434
		{
			yyVAL.expr = &ast.FuncExpr{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), Args: yyDollar[4].expr_idents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 73:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./parser/parser.y:439
		{
			yyVAL.expr = &ast.FuncExpr{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), Args: []int{ast.UniqueNames.Set(yyDollar[4].tok.Lit)}, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 74:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:444
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 75:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:449
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 76:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:454
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyDollar[3].expr_pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 77:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:463
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyDollar[3].expr_pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:472
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:477
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:482
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:487
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:492
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:497
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:502
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:507
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:512
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:517
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:522
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:527
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:532
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:537
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:542
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:547
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:552
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:557
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:562
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:567
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:572
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:577
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:582
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:587
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:592
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:597
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:602
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 105:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:607
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[1].tok.Lit), SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 106:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:612
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[1].tok.Lit), SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 107:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:617
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 108:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:622
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 109:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:627
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 110:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:632
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 111:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:637
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 112:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:642
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 113:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:647
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 114:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:652
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 115:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:657
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 116:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:662
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Begin: yyDollar[3].expr, End: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 117:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:667
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Begin: nil, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 118:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:672
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 119:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:677
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: nil}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 120:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:682
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: nil, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:687
		{
			yyVAL.expr = &ast.MakeExpr{Type: yyDollar[2].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:692
		{
			yyVAL.expr = &ast.MakeChanExpr{SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 123:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:697
		{
			yyVAL.expr = &ast.MakeChanExpr{SizeExpr: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 124:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:702
		{
			yyVAL.expr = &ast.MakeArrayExpr{LenExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 125:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:707
		{
			yyVAL.expr = &ast.MakeArrayExpr{LenExpr: yyDollar[3].expr, CapExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 126:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:712
		{
			yyVAL.expr = &ast.TypeCast{Type: yyDollar[2].typ.Name, CastExpr: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 127:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:717
		{
			yyVAL.expr = &ast.MakeExpr{TypeExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 128:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:722
		{
			yyVAL.expr = &ast.TypeCast{TypeExpr: yyDollar[3].expr, CastExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:727
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:732
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:743
		{
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:746
		{
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:751
		{
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:754
		{
		}
	}
	goto yystack /* stack new state and value */
}
