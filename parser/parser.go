//line .\parser\parser.y:1
package parser

import __yyfmt__ "fmt"

//line .\parser\parser.y:3
import (
	"github.com/covrom/gonec/ast"
)

//line .\parser\parser.y:29
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

//line .\parser\parser.y:772

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 6,
	1, 7,
	26, 7,
	-2, 134,
	-1, 12,
	61, 53,
	-2, 5,
	-1, 17,
	61, 54,
	-2, 28,
	-1, 26,
	28, 7,
	-2, 134,
	-1, 53,
	61, 53,
	-2, 135,
	-1, 104,
	1, 62,
	8, 62,
	11, 62,
	14, 62,
	26, 62,
	28, 62,
	44, 62,
	45, 62,
	53, 62,
	54, 62,
	58, 62,
	60, 62,
	61, 62,
	70, 62,
	71, 62,
	76, 62,
	79, 62,
	81, 62,
	82, 62,
	-2, 57,
	-1, 106,
	1, 64,
	8, 64,
	11, 64,
	14, 64,
	26, 64,
	28, 64,
	44, 64,
	45, 64,
	53, 64,
	54, 64,
	58, 64,
	60, 64,
	61, 64,
	70, 64,
	71, 64,
	76, 64,
	79, 64,
	81, 64,
	82, 64,
	-2, 57,
	-1, 138,
	17, 0,
	18, 0,
	-2, 90,
	-1, 139,
	17, 0,
	18, 0,
	-2, 91,
	-1, 159,
	61, 54,
	-2, 48,
	-1, 165,
	71, 7,
	-2, 134,
	-1, 166,
	71, 7,
	-2, 134,
	-1, 192,
	14, 7,
	54, 7,
	71, 7,
	-2, 134,
	-1, 240,
	17, 0,
	61, 55,
	-2, 49,
	-1, 241,
	1, 50,
	11, 50,
	14, 50,
	17, 50,
	26, 50,
	28, 50,
	44, 50,
	45, 50,
	54, 50,
	58, 50,
	61, 56,
	71, 50,
	81, 50,
	82, 50,
	-2, 57,
	-1, 250,
	1, 56,
	8, 56,
	14, 56,
	26, 56,
	28, 56,
	44, 56,
	45, 56,
	54, 56,
	61, 56,
	71, 56,
	76, 56,
	79, 56,
	81, 56,
	82, 56,
	-2, 57,
	-1, 264,
	71, 7,
	-2, 134,
	-1, 274,
	1, 111,
	8, 111,
	11, 111,
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
	-1, 276,
	1, 115,
	8, 115,
	11, 115,
	14, 115,
	26, 115,
	28, 115,
	44, 115,
	45, 115,
	53, 115,
	54, 115,
	58, 115,
	60, 115,
	61, 115,
	70, 115,
	71, 115,
	76, 115,
	79, 115,
	81, 115,
	82, 115,
	-2, 113,
	-1, 284,
	71, 7,
	-2, 134,
	-1, 288,
	44, 7,
	45, 7,
	71, 7,
	-2, 134,
	-1, 292,
	71, 7,
	-2, 134,
	-1, 293,
	71, 7,
	-2, 134,
	-1, 298,
	1, 110,
	8, 110,
	11, 110,
	14, 110,
	26, 110,
	28, 110,
	44, 110,
	45, 110,
	53, 110,
	54, 110,
	58, 110,
	60, 110,
	61, 110,
	70, 110,
	71, 110,
	76, 110,
	79, 110,
	81, 110,
	82, 110,
	-2, 108,
	-1, 299,
	1, 114,
	8, 114,
	11, 114,
	14, 114,
	26, 114,
	28, 114,
	44, 114,
	45, 114,
	53, 114,
	54, 114,
	58, 114,
	60, 114,
	61, 114,
	70, 114,
	71, 114,
	76, 114,
	79, 114,
	81, 114,
	82, 114,
	-2, 112,
	-1, 303,
	71, 7,
	-2, 134,
	-1, 308,
	71, 7,
	-2, 134,
	-1, 309,
	71, 7,
	-2, 134,
	-1, 310,
	44, 7,
	45, 7,
	71, 7,
	-2, 134,
	-1, 316,
	71, 7,
	-2, 134,
	-1, 328,
	14, 7,
	54, 7,
	71, 7,
	-2, 134,
}

const yyPrivate = 57344

const yyLast = 3241

var yyAct = [...]int{

	91, 208, 18, 181, 55, 209, 168, 187, 50, 8,
	9, 99, 100, 17, 269, 10, 228, 186, 226, 184,
	100, 178, 190, 92, 120, 110, 95, 275, 97, 299,
	90, 101, 102, 103, 105, 107, 8, 9, 8, 9,
	127, 108, 96, 119, 273, 113, 115, 298, 294, 214,
	122, 265, 124, 259, 17, 266, 195, 245, 128, 186,
	130, 131, 132, 133, 134, 135, 136, 137, 138, 139,
	140, 141, 142, 143, 144, 145, 146, 147, 148, 149,
	161, 127, 150, 151, 152, 153, 117, 155, 157, 159,
	159, 158, 160, 303, 154, 276, 222, 161, 109, 210,
	211, 171, 161, 182, 12, 210, 211, 319, 170, 161,
	161, 332, 274, 331, 176, 330, 118, 215, 52, 188,
	179, 189, 327, 325, 196, 324, 256, 159, 320, 193,
	313, 271, 207, 305, 255, 254, 126, 123, 164, 127,
	258, 230, 15, 89, 210, 211, 7, 111, 112, 94,
	304, 127, 166, 11, 3, 116, 201, 199, 282, 267,
	223, 54, 182, 14, 202, 203, 56, 244, 297, 6,
	233, 212, 225, 218, 206, 213, 221, 53, 220, 163,
	219, 204, 205, 224, 88, 177, 162, 129, 234, 119,
	5, 239, 240, 231, 232, 243, 125, 93, 180, 246,
	54, 249, 251, 169, 121, 2, 281, 4, 242, 302,
	23, 257, 13, 1, 0, 0, 0, 0, 260, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 191, 0,
	0, 272, 194, 0, 0, 0, 0, 278, 0, 279,
	0, 73, 74, 75, 76, 77, 78, 0, 0, 79,
	80, 64, 0, 285, 286, 0, 0, 0, 0, 0,
	87, 0, 0, 0, 290, 0, 200, 0, 0, 249,
	0, 0, 169, 0, 296, 59, 60, 61, 62, 63,
	291, 0, 0, 58, 227, 229, 83, 306, 85, 86,
	0, 81, 0, 0, 0, 0, 0, 0, 0, 0,
	307, 0, 0, 0, 311, 0, 318, 0, 314, 315,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 317,
	159, 0, 329, 0, 321, 322, 323, 264, 0, 0,
	0, 268, 326, 270, 28, 29, 35, 0, 0, 41,
	21, 16, 22, 51, 333, 24, 0, 0, 0, 0,
	0, 0, 0, 36, 37, 38, 0, 26, 0, 0,
	0, 0, 0, 288, 0, 0, 19, 20, 0, 0,
	292, 293, 0, 27, 0, 0, 45, 0, 46, 49,
	47, 39, 0, 0, 0, 25, 40, 48, 0, 0,
	0, 0, 310, 0, 0, 30, 34, 0, 0, 316,
	43, 0, 0, 31, 32, 33, 0, 44, 42, 0,
	0, 8, 9, 67, 68, 70, 72, 82, 84, 0,
	0, 0, 0, 0, 0, 0, 73, 74, 75, 76,
	77, 78, 0, 0, 79, 80, 64, 65, 66, 0,
	0, 0, 0, 0, 0, 87, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 238, 69, 71,
	59, 60, 61, 62, 63, 0, 0, 0, 58, 0,
	0, 83, 237, 85, 86, 0, 81, 67, 68, 70,
	72, 82, 84, 0, 0, 0, 0, 0, 0, 0,
	73, 74, 75, 76, 77, 78, 0, 0, 79, 80,
	64, 65, 66, 0, 0, 0, 0, 0, 0, 87,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 236, 69, 71, 59, 60, 61, 62, 63, 0,
	0, 0, 58, 0, 0, 83, 235, 85, 86, 0,
	81, 67, 68, 70, 72, 82, 84, 0, 0, 0,
	0, 0, 0, 0, 73, 74, 75, 76, 77, 78,
	0, 0, 79, 80, 64, 65, 66, 0, 0, 0,
	0, 0, 0, 87, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 217, 0, 69, 71, 59, 60,
	61, 62, 63, 0, 0, 0, 58, 0, 0, 83,
	0, 85, 86, 216, 81, 67, 68, 70, 72, 82,
	84, 0, 0, 0, 0, 0, 0, 0, 73, 74,
	75, 76, 77, 78, 0, 0, 79, 80, 64, 65,
	66, 0, 0, 0, 0, 0, 0, 87, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 198, 0,
	69, 71, 59, 60, 61, 62, 63, 0, 0, 0,
	58, 0, 0, 83, 0, 85, 86, 197, 81, 67,
	68, 70, 72, 82, 84, 0, 0, 0, 0, 0,
	0, 0, 73, 74, 75, 76, 77, 78, 0, 0,
	79, 80, 64, 65, 66, 0, 0, 0, 0, 0,
	0, 87, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 71, 59, 60, 61, 62,
	63, 0, 328, 0, 58, 0, 0, 83, 0, 85,
	86, 0, 81, 67, 68, 70, 72, 82, 84, 0,
	0, 0, 0, 0, 0, 0, 73, 74, 75, 76,
	77, 78, 0, 0, 79, 80, 64, 65, 66, 0,
	0, 0, 0, 0, 0, 87, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 71,
	59, 60, 61, 62, 63, 0, 0, 0, 58, 0,
	0, 83, 312, 85, 86, 0, 81, 67, 68, 70,
	72, 82, 84, 0, 0, 0, 0, 0, 0, 0,
	73, 74, 75, 76, 77, 78, 0, 0, 79, 80,
	64, 65, 66, 0, 0, 0, 0, 0, 0, 87,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 69, 71, 59, 60, 61, 62, 63, 0,
	309, 0, 58, 0, 0, 83, 0, 85, 86, 0,
	81, 67, 68, 70, 72, 82, 84, 0, 0, 0,
	0, 0, 0, 0, 73, 74, 75, 76, 77, 78,
	0, 0, 79, 80, 64, 65, 66, 0, 0, 0,
	0, 0, 0, 87, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 69, 71, 59, 60,
	61, 62, 63, 0, 308, 0, 58, 0, 0, 83,
	0, 85, 86, 0, 81, 67, 68, 70, 72, 82,
	84, 0, 0, 0, 0, 0, 0, 0, 73, 74,
	75, 76, 77, 78, 0, 0, 79, 80, 64, 65,
	66, 0, 0, 0, 0, 0, 0, 87, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	69, 71, 59, 60, 61, 62, 63, 0, 0, 0,
	58, 0, 0, 83, 301, 85, 86, 0, 81, 67,
	68, 70, 72, 82, 84, 0, 0, 0, 0, 0,
	0, 0, 73, 74, 75, 76, 77, 78, 0, 0,
	79, 80, 64, 65, 66, 0, 0, 0, 0, 0,
	0, 87, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 71, 59, 60, 61, 62,
	63, 0, 0, 0, 58, 0, 0, 83, 300, 85,
	86, 0, 81, 67, 68, 70, 72, 82, 84, 0,
	0, 0, 0, 0, 0, 0, 73, 74, 75, 76,
	77, 78, 0, 0, 79, 80, 64, 65, 66, 0,
	0, 0, 0, 0, 0, 87, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 71,
	59, 60, 61, 62, 63, 0, 0, 0, 58, 0,
	0, 83, 0, 85, 86, 289, 81, 67, 68, 70,
	72, 82, 84, 0, 0, 0, 0, 0, 0, 0,
	73, 74, 75, 76, 77, 78, 0, 0, 79, 80,
	64, 65, 66, 0, 0, 0, 0, 0, 0, 87,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	287, 0, 69, 71, 59, 60, 61, 62, 63, 0,
	0, 0, 58, 0, 0, 83, 0, 85, 86, 0,
	81, 67, 68, 70, 72, 82, 84, 0, 0, 0,
	0, 0, 0, 0, 73, 74, 75, 76, 77, 78,
	0, 0, 79, 80, 64, 65, 66, 0, 0, 0,
	0, 0, 0, 87, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 69, 71, 59, 60,
	61, 62, 63, 0, 284, 0, 58, 0, 0, 83,
	0, 85, 86, 0, 81, 67, 68, 70, 72, 82,
	84, 0, 0, 0, 0, 0, 0, 0, 73, 74,
	75, 76, 77, 78, 0, 0, 79, 80, 64, 65,
	66, 0, 0, 0, 0, 0, 0, 87, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	69, 71, 59, 60, 61, 62, 63, 0, 0, 0,
	58, 0, 0, 83, 0, 85, 86, 283, 81, 67,
	68, 70, 72, 82, 84, 0, 0, 0, 0, 0,
	0, 0, 73, 74, 75, 76, 77, 78, 0, 0,
	79, 80, 64, 65, 66, 0, 0, 0, 0, 0,
	0, 87, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 71, 59, 60, 61, 62,
	63, 0, 0, 0, 58, 0, 0, 83, 280, 85,
	86, 0, 81, 67, 68, 70, 72, 82, 84, 0,
	0, 0, 0, 0, 0, 0, 73, 74, 75, 76,
	77, 78, 0, 0, 79, 80, 64, 65, 66, 0,
	0, 0, 0, 0, 0, 87, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 71,
	59, 60, 61, 62, 63, 0, 0, 0, 58, 0,
	0, 83, 277, 85, 86, 0, 81, 67, 68, 70,
	72, 82, 84, 0, 0, 0, 0, 0, 0, 0,
	73, 74, 75, 76, 77, 78, 0, 0, 79, 80,
	64, 65, 66, 0, 0, 0, 0, 0, 0, 87,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 263, 69, 71, 59, 60, 61, 62, 63, 0,
	0, 0, 58, 0, 0, 83, 0, 85, 86, 0,
	81, 67, 68, 70, 72, 82, 84, 0, 0, 0,
	0, 0, 0, 0, 73, 74, 75, 76, 77, 78,
	0, 0, 79, 80, 64, 65, 66, 0, 0, 0,
	0, 0, 0, 87, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 69, 71, 59, 60,
	61, 62, 63, 0, 0, 0, 58, 0, 0, 83,
	0, 85, 86, 262, 81, 67, 68, 70, 72, 82,
	84, 0, 0, 0, 0, 0, 0, 0, 73, 74,
	75, 76, 77, 78, 0, 0, 79, 80, 64, 65,
	66, 0, 0, 0, 0, 0, 0, 87, 0, 0,
	0, 253, 0, 0, 0, 0, 0, 0, 0, 0,
	69, 71, 59, 60, 61, 62, 63, 0, 0, 0,
	58, 0, 0, 83, 0, 85, 86, 0, 81, 67,
	68, 70, 72, 82, 84, 0, 0, 0, 0, 0,
	0, 0, 73, 74, 75, 76, 77, 78, 0, 0,
	79, 80, 64, 65, 66, 0, 0, 0, 0, 0,
	0, 87, 0, 0, 0, 252, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 71, 59, 60, 61, 62,
	63, 0, 0, 0, 58, 0, 0, 83, 0, 85,
	86, 0, 81, 67, 68, 70, 72, 82, 84, 0,
	0, 0, 0, 0, 0, 0, 73, 74, 75, 76,
	77, 78, 0, 0, 79, 80, 64, 65, 66, 0,
	0, 0, 0, 0, 0, 87, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 71,
	59, 60, 61, 62, 63, 0, 0, 0, 58, 0,
	0, 83, 0, 85, 86, 248, 81, 67, 68, 70,
	72, 82, 84, 0, 0, 0, 0, 0, 0, 0,
	73, 74, 75, 76, 77, 78, 0, 0, 79, 80,
	64, 65, 66, 0, 0, 0, 0, 0, 0, 87,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 69, 71, 59, 60, 61, 62, 63, 0,
	192, 0, 58, 0, 0, 83, 0, 85, 86, 0,
	81, 67, 68, 70, 72, 82, 84, 0, 0, 0,
	0, 0, 0, 0, 73, 74, 75, 76, 77, 78,
	0, 0, 79, 80, 64, 65, 66, 0, 0, 0,
	0, 0, 0, 87, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 69, 71, 59, 60,
	61, 62, 63, 0, 0, 0, 58, 0, 0, 83,
	183, 85, 86, 0, 81, 67, 68, 70, 72, 82,
	84, 0, 0, 0, 0, 0, 0, 0, 73, 74,
	75, 76, 77, 78, 0, 0, 79, 80, 64, 65,
	66, 0, 0, 0, 0, 0, 0, 87, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 175,
	69, 71, 59, 60, 61, 62, 63, 0, 0, 0,
	58, 0, 0, 83, 0, 85, 86, 0, 81, 67,
	68, 70, 72, 82, 84, 0, 0, 0, 0, 0,
	0, 0, 73, 74, 75, 76, 77, 78, 0, 0,
	79, 80, 64, 65, 66, 0, 0, 0, 0, 0,
	0, 87, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 167, 0, 69, 71, 59, 60, 61, 62,
	63, 0, 0, 0, 58, 0, 0, 83, 0, 85,
	86, 0, 81, 67, 68, 70, 72, 82, 84, 0,
	0, 0, 0, 0, 0, 0, 73, 74, 75, 76,
	77, 78, 0, 0, 79, 80, 64, 65, 66, 0,
	0, 0, 0, 0, 0, 87, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 71,
	59, 60, 61, 62, 63, 0, 165, 0, 58, 0,
	0, 83, 0, 85, 86, 0, 81, 67, 68, 70,
	72, 82, 84, 0, 0, 0, 0, 0, 0, 0,
	73, 74, 75, 76, 77, 78, 0, 0, 79, 80,
	64, 65, 66, 0, 0, 0, 0, 0, 0, 87,
	0, 0, 0, 0, 0, 0, 0, 0, 57, 0,
	0, 0, 69, 71, 59, 60, 61, 62, 63, 0,
	0, 0, 58, 0, 0, 83, 0, 85, 86, 0,
	81, 67, 68, 70, 72, 82, 84, 0, 0, 0,
	0, 0, 0, 0, 73, 74, 75, 76, 77, 78,
	0, 0, 79, 80, 64, 65, 66, 0, 0, 0,
	0, 0, 0, 87, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 69, 71, 59, 60,
	61, 62, 63, 0, 0, 0, 58, 0, 0, 83,
	0, 85, 86, 0, 81, 67, 68, 70, 72, 82,
	84, 0, 0, 0, 0, 0, 0, 0, 73, 74,
	75, 76, 77, 78, 0, 0, 79, 80, 64, 65,
	66, 0, 0, 0, 0, 0, 0, 87, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	69, 71, 59, 60, 61, 62, 63, 0, 0, 0,
	58, 0, 0, 83, 0, 185, 86, 0, 81, 67,
	68, 70, 72, 82, 84, 0, 0, 0, 0, 0,
	0, 0, 73, 74, 75, 76, 77, 78, 0, 0,
	79, 80, 64, 65, 66, 0, 0, 0, 0, 0,
	0, 87, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 71, 59, 60, 61, 62,
	63, 0, 0, 0, 174, 0, 0, 83, 0, 85,
	86, 0, 81, 67, 68, 70, 72, 82, 84, 0,
	0, 0, 0, 0, 0, 0, 73, 74, 75, 76,
	77, 78, 0, 0, 79, 80, 64, 65, 66, 0,
	0, 0, 0, 0, 0, 87, 28, 29, 35, 0,
	0, 41, 21, 16, 22, 51, 0, 24, 69, 71,
	59, 60, 61, 62, 63, 36, 37, 38, 173, 26,
	0, 83, 0, 85, 86, 0, 81, 0, 19, 20,
	0, 0, 0, 0, 0, 27, 0, 0, 45, 0,
	46, 49, 47, 39, 0, 0, 0, 25, 40, 48,
	0, 0, 0, 0, 0, 0, 0, 30, 34, 0,
	0, 0, 43, 0, 0, 31, 32, 33, 0, 44,
	42, 68, 70, 72, 82, 84, 0, 0, 0, 0,
	0, 0, 0, 73, 74, 75, 76, 77, 78, 0,
	0, 79, 80, 64, 65, 66, 0, 0, 0, 0,
	0, 0, 87, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 71, 59, 60, 61,
	62, 63, 0, 0, 0, 58, 0, 0, 83, 0,
	85, 86, 0, 81, 67, 68, 70, 72, 0, 84,
	0, 0, 0, 0, 0, 0, 0, 73, 74, 75,
	76, 77, 78, 0, 0, 79, 80, 64, 65, 66,
	0, 0, 0, 0, 0, 0, 87, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	71, 59, 60, 61, 62, 63, 0, 0, 0, 58,
	0, 0, 83, 0, 85, 86, 0, 81, 67, 68,
	70, 72, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 73, 74, 75, 76, 77, 78, 0, 0, 79,
	80, 64, 65, 66, 0, 0, 0, 0, 0, 0,
	87, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 71, 59, 60, 61, 62, 63,
	0, 70, 72, 58, 0, 0, 83, 0, 85, 86,
	0, 81, 73, 74, 75, 76, 77, 78, 0, 0,
	79, 80, 64, 65, 66, 0, 0, 0, 0, 0,
	0, 87, 250, 29, 35, 0, 0, 41, 0, 0,
	0, 0, 0, 0, 69, 71, 59, 60, 61, 62,
	63, 36, 37, 38, 58, 0, 0, 83, 0, 85,
	86, 0, 81, 0, 0, 0, 0, 28, 29, 35,
	0, 0, 41, 0, 45, 0, 46, 49, 47, 39,
	0, 0, 0, 0, 40, 48, 36, 37, 38, 0,
	0, 0, 0, 30, 34, 0, 0, 0, 43, 0,
	0, 31, 32, 33, 0, 44, 42, 295, 0, 45,
	0, 46, 49, 47, 39, 0, 0, 0, 0, 40,
	48, 0, 0, 0, 0, 28, 29, 35, 30, 34,
	41, 0, 0, 43, 0, 0, 31, 32, 33, 0,
	44, 42, 261, 0, 36, 37, 38, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	28, 29, 35, 0, 0, 41, 0, 45, 0, 46,
	49, 47, 39, 0, 0, 0, 0, 40, 48, 36,
	37, 38, 0, 0, 0, 0, 30, 34, 0, 0,
	0, 43, 0, 0, 31, 32, 33, 0, 44, 42,
	247, 0, 45, 0, 46, 49, 47, 39, 0, 0,
	0, 0, 40, 48, 0, 0, 172, 0, 28, 29,
	35, 30, 34, 41, 0, 0, 43, 0, 0, 31,
	32, 33, 0, 44, 42, 0, 0, 36, 37, 38,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 28, 29, 35, 0,
	45, 41, 46, 49, 47, 39, 0, 0, 0, 0,
	40, 48, 0, 0, 156, 36, 37, 38, 0, 30,
	34, 0, 0, 0, 43, 0, 0, 31, 32, 33,
	0, 44, 42, 0, 28, 29, 35, 0, 45, 41,
	46, 49, 47, 39, 0, 0, 0, 0, 40, 48,
	0, 0, 98, 36, 37, 38, 0, 30, 34, 0,
	0, 0, 43, 0, 0, 31, 32, 33, 0, 44,
	42, 0, 250, 29, 35, 0, 45, 41, 46, 49,
	47, 39, 0, 0, 0, 0, 40, 48, 0, 0,
	0, 36, 37, 38, 0, 30, 34, 0, 0, 0,
	43, 0, 0, 31, 32, 33, 0, 44, 42, 0,
	241, 29, 35, 0, 45, 41, 46, 49, 47, 39,
	0, 0, 0, 0, 40, 48, 0, 0, 0, 36,
	37, 38, 0, 30, 34, 0, 0, 0, 43, 0,
	0, 31, 32, 33, 0, 44, 42, 0, 114, 29,
	35, 0, 45, 41, 46, 49, 47, 39, 0, 0,
	0, 0, 40, 48, 0, 0, 0, 36, 37, 38,
	0, 30, 34, 0, 0, 0, 43, 0, 0, 31,
	32, 33, 0, 44, 42, 0, 106, 29, 35, 0,
	45, 41, 46, 49, 47, 39, 0, 0, 0, 0,
	40, 48, 0, 0, 0, 36, 37, 38, 0, 30,
	34, 0, 0, 0, 43, 0, 0, 31, 32, 33,
	0, 44, 42, 0, 104, 29, 35, 0, 45, 41,
	46, 49, 47, 39, 0, 0, 0, 0, 40, 48,
	0, 0, 0, 36, 37, 38, 0, 30, 34, 0,
	0, 0, 43, 0, 0, 31, 32, 33, 0, 44,
	42, 0, 0, 0, 0, 0, 45, 0, 46, 49,
	47, 39, 0, 0, 0, 0, 40, 48, 0, 0,
	0, 0, 0, 0, 0, 30, 34, 0, 0, 0,
	43, 0, 0, 31, 32, 33, 0, 44, 42, 73,
	74, 75, 76, 77, 78, 0, 0, 0, 0, 64,
	73, 74, 75, 76, 77, 78, 0, 0, 87, 0,
	64, 0, 0, 0, 0, 0, 0, 0, 0, 87,
	0, 0, 0, 0, 0, 61, 62, 63, 0, 0,
	0, 58, 0, 0, 83, 0, 85, 86, 0, 81,
	0, 0, 58, 0, 0, 83, 0, 85, 86, 0,
	81,
}
var yyPact = [...]int{

	128, 128, -1000, 186, -1000, -72, -72, -1000, -1000, -1000,
	-1000, -1000, 2362, -72, -72, -1000, 162, 2060, 126, -1000,
	-1000, 2910, 2910, -1000, 145, 2910, -72, 2872, -66, -1000,
	2910, 2910, 2910, 3100, 3062, -1000, -1000, -1000, -1000, -1000,
	2910, 21, -72, -72, 2910, 3024, 39, -53, 185, 2910,
	76, 2910, -1000, 330, -1000, 78, -1000, 2910, 183, 2910,
	2910, 2910, 2910, 2910, 2910, 2910, 2910, 2910, 2910, 2910,
	2910, 2910, 2910, 2910, 2910, 2910, 2910, 2910, 2910, -1000,
	-1000, 2910, 2910, 2910, 2910, 2910, 2834, 2910, 2910, 2910,
	49, 2124, 2124, 182, 121, 1996, 124, 1932, -72, 2910,
	2776, 3160, 3160, 3160, -66, 2316, -66, 2252, 1868, 181,
	-56, 2910, 156, 1804, -58, 2188, -13, -70, 2910, -1000,
	2910, -55, 2124, -72, 1740, -1000, 2910, -72, 2124, -1000,
	3149, 3149, 3160, 3160, 3160, 2124, 211, 211, 2602, 2602,
	211, 211, 211, 211, 2124, 2124, 2124, 2124, 2124, 2124,
	2124, 2487, 2124, 2551, 48, 588, 2910, 2124, -1000, 2124,
	-1000, -72, 140, 2910, 2910, -72, -72, -72, 61, 100,
	41, 524, 2910, 176, 174, 2910, 20, 152, 168, -43,
	-45, -1000, 81, -1000, 2910, 2910, 166, 2910, 460, 396,
	2910, 2986, -72, -1000, 163, -19, -1000, -1000, 2741, 1676,
	2948, 2910, 1612, 1548, 64, 63, 55, -1000, -1000, -1000,
	2910, 80, -1000, -1000, -23, -1000, -1000, 2683, 1484, -1000,
	-1000, 1420, -72, -25, -21, 151, -72, -65, -72, 60,
	2910, 36, 19, -1000, 1356, -1000, 2910, -1000, 2910, 1292,
	2423, -66, -1000, 147, -1000, -1000, 1228, -1000, -1000, 2124,
	-66, 1164, 2910, 2910, -1000, -1000, -1000, 1100, -72, -1000,
	1036, -1000, -1000, 2910, -72, -72, -72, -28, 2648, -1000,
	97, -1000, 2124, -29, -1000, -47, -1000, -1000, 972, 908,
	-1000, 79, 162, -1000, -72, 844, 780, -72, -72, -1000,
	716, 59, -72, -72, -72, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -72, -1000, 2910, 90, 57, -72, -72,
	-72, -1000, -1000, -1000, 54, 52, -72, 51, 652, 2910,
	-1000, 44, 42, -1000, -1000, -1000, 40, -1000, -72, -1000,
	-1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 15, 213, 205, 212, 142, 210, 5, 1, 6,
	209, 206, 155, 0, 8, 2, 3, 198, 4, 163,
	104, 195, 146,
}
var yyR1 = [...]int{

	0, 2, 2, 2, 3, 1, 1, 4, 4, 4,
	21, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5, 5, 11,
	11, 10, 6, 6, 9, 9, 9, 9, 9, 8,
	7, 16, 17, 17, 17, 18, 18, 18, 15, 15,
	15, 12, 12, 14, 14, 14, 14, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 20, 20, 19, 19, 22, 22,
}
var yyR2 = [...]int{

	0, 0, 1, 2, 4, 1, 2, 0, 2, 3,
	0, 9, 2, 3, 3, 3, 1, 1, 2, 2,
	1, 8, 9, 9, 5, 5, 5, 4, 1, 0,
	2, 4, 8, 6, 0, 2, 2, 2, 2, 5,
	4, 3, 0, 1, 4, 0, 1, 4, 1, 4,
	4, 1, 3, 0, 1, 4, 4, 1, 1, 2,
	2, 2, 2, 4, 2, 4, 1, 1, 1, 1,
	1, 7, 3, 7, 8, 8, 9, 5, 6, 5,
	6, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 2, 3, 3, 3, 3, 5, 4,
	6, 5, 5, 4, 6, 5, 4, 4, 6, 5,
	5, 6, 5, 5, 2, 2, 5, 4, 6, 5,
	4, 6, 3, 2, 0, 1, 1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -2, -3, 26, -3, 4, -19, -22, 81, 82,
	-1, -22, -20, -4, -19, -5, 11, -13, -15, 36,
	37, 10, 12, -6, 15, 55, 27, 43, 4, 5,
	65, 73, 74, 75, 66, 6, 23, 24, 25, 51,
	56, 9, 78, 70, 77, 46, 48, 50, 57, 49,
	-14, 13, -20, -19, -22, -18, 4, 58, 72, 64,
	65, 66, 67, 68, 40, 41, 42, 17, 18, 62,
	19, 63, 20, 30, 31, 32, 33, 34, 35, 38,
	39, 80, 21, 75, 22, 77, 78, 49, 58, 17,
	-14, -13, -13, 52, 4, -13, -1, -13, 60, 77,
	78, -13, -13, -13, 4, -13, 4, -13, -13, 77,
	4, -20, -20, -13, 4, -13, -12, 47, 77, 4,
	77, -12, -13, 61, -13, -5, 58, 61, -13, 4,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -14, -13, 60, -13, -15, -13,
	-15, 61, 4, 58, 17, 70, 28, 60, -9, -20,
	-14, -13, 60, 72, 72, 61, -18, 4, 77, -14,
	-17, -16, 6, 76, 77, 77, 72, 77, -13, -13,
	77, -20, 70, -15, -20, 8, 76, 79, 60, -13,
	-20, 16, -13, -13, -1, -1, -9, 71, -8, -7,
	44, 45, -8, -7, 8, 76, 79, 60, -13, 4,
	4, -13, 76, 8, -18, 4, 61, -20, 61, -20,
	60, -14, -14, 4, -13, 76, 61, 76, 61, -13,
	-13, 4, -1, -21, 4, 76, -13, 79, 79, -13,
	4, -13, 53, 53, 71, 71, 71, -13, 60, 76,
	-13, 79, 79, 61, -20, 76, 76, 8, -20, 79,
	-20, 71, -13, 8, 76, 8, 76, 76, -13, -13,
	76, -11, 11, 79, 70, -13, -13, 60, -20, 79,
	-13, -1, -20, -20, 76, 79, -16, 71, 76, 76,
	76, 76, -10, 14, 71, 54, -18, -1, 70, 70,
	-20, -1, 76, 71, -1, -1, -20, -1, -13, 17,
	71, -1, -1, -1, 71, 71, -1, 71, 70, -15,
	71, 71, 71, -1,
}
var yyDef = [...]int{

	1, -2, 2, 0, 3, 0, -2, 136, 138, 139,
	4, 136, -2, 134, 135, 8, 45, -2, 0, 16,
	17, 53, 0, 20, 0, 0, -2, 0, 57, 58,
	0, 0, 0, 0, 0, 66, 67, 68, 69, 70,
	0, 0, 134, 134, 0, 0, 0, 0, 0, 0,
	0, 0, 6, -2, 137, 12, 46, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 102,
	103, 0, 0, 0, 0, 53, 0, 0, 53, 53,
	18, 54, 19, 0, 0, 0, 0, 0, 34, 53,
	0, 59, 60, 61, -2, 0, -2, 0, 0, 45,
	0, 53, 42, 0, 57, 0, 124, 125, 0, 51,
	0, 0, 133, 134, 0, 9, 53, 134, 13, 72,
	82, 83, 84, 85, 86, 87, 88, 89, -2, -2,
	92, 93, 94, 95, 96, 97, 98, 99, 100, 101,
	104, 105, 106, 107, 0, 0, 0, 132, 14, -2,
	15, 134, 0, 0, 0, -2, -2, 34, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 46, 45, 134,
	134, 43, 0, 81, 53, 53, 0, 0, 0, 0,
	0, 0, -2, 10, 0, 0, 113, 117, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 27, 37, 38,
	0, 0, 35, 36, 0, 109, 116, 0, 0, 63,
	65, 0, 134, 0, 0, 46, 134, 0, 134, 0,
	0, 0, 0, 52, 0, 130, 0, 127, 0, 0,
	-2, -2, 29, 0, 47, 112, 0, 122, 123, 55,
	-2, 0, 0, 0, 24, 25, 26, 0, 134, 108,
	0, 119, 120, 0, -2, 134, 134, 0, 0, 77,
	0, 79, 41, 0, -2, 0, -2, 126, 0, 0,
	129, 0, 45, 121, -2, 0, 0, 134, -2, 118,
	0, 0, -2, -2, 134, 78, 44, 80, -2, -2,
	131, 128, 30, -2, 33, 0, 0, 0, -2, -2,
	-2, 40, 71, 73, 0, 0, -2, 0, 0, 53,
	21, 0, 0, 39, 74, 75, 0, 32, -2, 11,
	22, 23, 76, 31,
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
		//line .\parser\parser.y:71
		{
			yyVAL.modules = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.modules
			}
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:78
		{
			yyVAL.modules = []ast.Stmt{yyDollar[1].module}
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.modules
			}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:85
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
		//line .\parser\parser.y:96
		{
			yyVAL.module = &ast.ModuleStmt{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), Stmts: yyDollar[4].compstmt}
			yyVAL.module.SetPosition(yyDollar[1].tok.Position())
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:102
		{
			yyVAL.compstmt = nil
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:106
		{
			yyVAL.compstmt = yyDollar[1].stmts
		}
	case 7:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line .\parser\parser.y:111
		{
			yyVAL.stmts = nil
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:115
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:119
		{
			if yyDollar[3].stmt != nil {
				yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[3].stmt)
			}
		}
	case 10:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:127
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].expr_many}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line .\parser\parser.y:132
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].expr_many}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:137
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: nil}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:142
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "=", Rhss: []ast.Expr{yyDollar[3].expr}}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:146
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:150
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: &ast.BinOpExpr{Lhss: yyDollar[1].expr_many, Operator: "==", Rhss: yyDollar[3].expr_many}}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:154
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:159
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:164
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:169
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:174
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 21:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line .\parser\parser.y:179
		{
			yyVAL.stmt = &ast.ForStmt{Var: ast.UniqueNames.Set(yyDollar[3].tok.Lit), Value: yyDollar[5].expr, Stmts: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line .\parser\parser.y:184
		{
			yyVAL.stmt = &ast.NumForStmt{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), Expr1: yyDollar[4].expr, Expr2: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 23:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line .\parser\parser.y:189
		{
			yyVAL.stmt = &ast.NumForStmt{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), Expr1: yyDollar[4].expr, Expr2: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 24:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:194
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 25:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:199
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[2].compstmt, Catch: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:204
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:209
		{
			yyVAL.stmt = &ast.SelectStmt{Cases: yyDollar[3].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:214
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 29:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line .\parser\parser.y:220
		{
			yyVAL.stmt_elsifs = []ast.Stmt{}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:224
		{
			yyVAL.stmt_elsifs = append(yyDollar[1].stmt_elsifs, yyDollar[2].stmt_elsif)
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:230
		{
			yyVAL.stmt_elsif = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt}
		}
	case 32:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line .\parser\parser.y:236
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, ElseIf: yyDollar[5].stmt_elsifs, Else: yyDollar[7].compstmt}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 33:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line .\parser\parser.y:241
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, ElseIf: yyDollar[5].stmt_elsifs, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 34:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line .\parser\parser.y:247
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:251
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_case}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:255
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_default}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:259
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_case)
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:263
		{
			for _, stmt := range yyDollar[1].stmt_cases {
				if _, ok := stmt.(*ast.DefaultStmt); ok {
					yylex.Error("multiple default statement")
				}
			}
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_default)
		}
	case 39:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:274
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[5].compstmt}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:280
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyDollar[4].compstmt}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:286
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line .\parser\parser.y:291
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:295
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:299
		{
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[4].expr_pair)
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line .\parser\parser.y:304
		{
			yyVAL.expr_idents = []int{}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:308
		{
			yyVAL.expr_idents = []int{ast.UniqueNames.Set(yyDollar[1].tok.Lit)}
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:312
		{
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, ast.UniqueNames.Set(yyDollar[4].tok.Lit))
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:318
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 49:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:322
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:326
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[4].tok.Lit)})
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:331
		{
			yyVAL.typ = ast.Type{Name: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:335
		{
			yyVAL.typ = ast.Type{Name: ast.UniqueNames.Set(ast.UniqueNames.Get(yyDollar[1].typ.Name) + "." + yyDollar[3].tok.Lit)}
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line .\parser\parser.y:340
		{
			yyVAL.exprs = nil
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:344
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:348
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:352
		{
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[4].tok.Lit)})
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:358
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:363
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:368
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:373
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:378
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:383
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[2].tok.Lit)}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 63:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:388
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: ast.UniqueNames.Set(yyDollar[4].tok.Lit)}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:393
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[2].tok.Lit)}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 65:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:398
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: ast.UniqueNames.Set(yyDollar[4].tok.Lit)}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:403
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:408
		{
			yyVAL.expr = &ast.ConstExpr{Value: "истина"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:413
		{
			yyVAL.expr = &ast.ConstExpr{Value: "ложь"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:418
		{
			yyVAL.expr = &ast.ConstExpr{Value: "неопределено"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:423
		{
			yyVAL.expr = &ast.ConstExpr{Value: "null"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 71:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line .\parser\parser.y:428
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[2].expr, Lhs: yyDollar[4].expr, Rhs: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:433
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: ast.UniqueNames.Set(yyDollar[3].tok.Lit)}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 73:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line .\parser\parser.y:438
		{
			yyVAL.expr = &ast.FuncExpr{Name: ast.UniqueNames.Set("<анонимная функция>"), Args: yyDollar[3].expr_idents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 74:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line .\parser\parser.y:443
		{
			yyVAL.expr = &ast.FuncExpr{Name: ast.UniqueNames.Set("<анонимная функция>"), Args: []int{ast.UniqueNames.Set(yyDollar[3].tok.Lit)}, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 75:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line .\parser\parser.y:448
		{
			yyVAL.expr = &ast.FuncExpr{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), Args: yyDollar[4].expr_idents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 76:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line .\parser\parser.y:453
		{
			yyVAL.expr = &ast.FuncExpr{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), Args: []int{ast.UniqueNames.Set(yyDollar[4].tok.Lit)}, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 77:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:458
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 78:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line .\parser\parser.y:463
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 79:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:468
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
	case 80:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line .\parser\parser.y:477
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
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:486
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:491
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "+", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:496
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "-", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:501
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "*", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:506
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "/", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:511
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "%", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:516
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "**", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:521
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "<<", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:526
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: ">>", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:531
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "==", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:536
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "!=", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:541
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: ">", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:546
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: ">=", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:551
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "<", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:556
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "<=", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:561
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:566
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:571
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:576
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:581
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:586
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:591
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:596
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:601
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "|", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:606
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "||", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:611
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "&", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:616
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "&&", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 108:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:621
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[1].tok.Lit), SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 109:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:626
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[1].tok.Lit), SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 110:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line .\parser\parser.y:631
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 111:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:636
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 112:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:641
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 113:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:646
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 114:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line .\parser\parser.y:651
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 115:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:656
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 116:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:661
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 117:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:666
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 118:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line .\parser\parser.y:671
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 119:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:676
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Begin: yyDollar[3].expr, End: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 120:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:681
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Begin: nil, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 121:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line .\parser\parser.y:686
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 122:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:691
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: nil}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 123:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:696
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: nil, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:701
		{
			yyVAL.expr = &ast.MakeExpr{Type: yyDollar[2].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:706
		{
			yyVAL.expr = &ast.MakeChanExpr{SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 126:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:711
		{
			yyVAL.expr = &ast.MakeChanExpr{SizeExpr: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 127:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:716
		{
			yyVAL.expr = &ast.MakeArrayExpr{LenExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 128:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line .\parser\parser.y:721
		{
			yyVAL.expr = &ast.MakeArrayExpr{LenExpr: yyDollar[3].expr, CapExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 129:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:726
		{
			yyVAL.expr = &ast.TypeCast{Type: yyDollar[2].typ.Name, CastExpr: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 130:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:731
		{
			yyVAL.expr = &ast.MakeExpr{TypeExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 131:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line .\parser\parser.y:736
		{
			yyVAL.expr = &ast.TypeCast{TypeExpr: yyDollar[3].expr, CastExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:741
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:746
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:757
		{
		}
	case 137:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:760
		{
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:765
		{
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:768
		{
		}
	}
	goto yystack /* stack new state and value */
}
