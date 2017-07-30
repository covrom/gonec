//line parser.y:1
package parser

import __yyfmt__ "fmt"

//line parser.y:3
import (
	"github.com/covrom/gonec/ast"
)

//line parser.y:27
type yySymType struct {
	yys          int
	compstmt     []ast.Stmt
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
	expr_idents  []string
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
const NEW = 57365
const TRUE = 57366
const FALSE = 57367
const NIL = 57368
const MODULE = 57369
const TRY = 57370
const CATCH = 57371
const FINALLY = 57372
const PLUSEQ = 57373
const MINUSEQ = 57374
const MULEQ = 57375
const DIVEQ = 57376
const ANDEQ = 57377
const OREQ = 57378
const BREAK = 57379
const CONTINUE = 57380
const PLUSPLUS = 57381
const MINUSMINUS = 57382
const POW = 57383
const SHIFTLEFT = 57384
const SHIFTRIGHT = 57385
const SWITCH = 57386
const CASE = 57387
const DEFAULT = 57388
const GO = 57389
const CHAN = 57390
const MAKE = 57391
const OPCHAN = 57392
const ARRAYLIT = 57393
const NULL = 57394
const EACH = 57395
const TO = 57396
const ELSIF = 57397
const WHILE = 57398
const UNARY = 57399

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
	"NEW",
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
	"'('",
	"')'",
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

//line parser.y:702

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 3,
	-2, 121,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	60, 46,
	-2, 1,
	-1, 10,
	60, 47,
	-2, 21,
	-1, 20,
	29, 3,
	-2, 121,
	-1, 45,
	60, 46,
	-2, 122,
	-1, 96,
	1, 55,
	8, 55,
	14, 55,
	29, 55,
	45, 55,
	46, 55,
	54, 55,
	55, 55,
	57, 55,
	59, 55,
	60, 55,
	69, 55,
	70, 55,
	76, 55,
	78, 55,
	80, 55,
	81, 55,
	-2, 50,
	-1, 98,
	1, 57,
	8, 57,
	14, 57,
	29, 57,
	45, 57,
	46, 57,
	54, 57,
	55, 57,
	57, 57,
	59, 57,
	60, 57,
	69, 57,
	70, 57,
	76, 57,
	78, 57,
	80, 57,
	81, 57,
	-2, 50,
	-1, 126,
	17, 0,
	18, 0,
	-2, 84,
	-1, 127,
	17, 0,
	18, 0,
	-2, 85,
	-1, 146,
	60, 47,
	-2, 41,
	-1, 148,
	70, 3,
	-2, 121,
	-1, 151,
	70, 3,
	-2, 121,
	-1, 152,
	70, 3,
	-2, 121,
	-1, 174,
	14, 3,
	55, 3,
	70, 3,
	-2, 121,
	-1, 212,
	60, 48,
	-2, 42,
	-1, 213,
	1, 43,
	14, 43,
	29, 43,
	45, 43,
	46, 43,
	55, 43,
	57, 43,
	60, 49,
	70, 43,
	80, 43,
	81, 43,
	-2, 50,
	-1, 220,
	1, 49,
	8, 49,
	14, 49,
	29, 49,
	45, 49,
	46, 49,
	55, 49,
	60, 49,
	70, 49,
	76, 49,
	78, 49,
	80, 49,
	81, 49,
	-2, 50,
	-1, 235,
	70, 3,
	-2, 121,
	-1, 246,
	1, 105,
	8, 105,
	14, 105,
	29, 105,
	45, 105,
	46, 105,
	54, 105,
	55, 105,
	57, 105,
	59, 105,
	60, 105,
	69, 105,
	70, 105,
	76, 105,
	78, 105,
	80, 105,
	81, 105,
	-2, 103,
	-1, 248,
	1, 109,
	8, 109,
	14, 109,
	29, 109,
	45, 109,
	46, 109,
	54, 109,
	55, 109,
	57, 109,
	59, 109,
	60, 109,
	69, 109,
	70, 109,
	76, 109,
	78, 109,
	80, 109,
	81, 109,
	-2, 107,
	-1, 254,
	70, 3,
	-2, 121,
	-1, 260,
	70, 3,
	-2, 121,
	-1, 261,
	70, 3,
	-2, 121,
	-1, 266,
	1, 104,
	8, 104,
	14, 104,
	29, 104,
	45, 104,
	46, 104,
	54, 104,
	55, 104,
	57, 104,
	59, 104,
	60, 104,
	69, 104,
	70, 104,
	76, 104,
	78, 104,
	80, 104,
	81, 104,
	-2, 102,
	-1, 267,
	1, 108,
	8, 108,
	14, 108,
	29, 108,
	45, 108,
	46, 108,
	54, 108,
	55, 108,
	57, 108,
	59, 108,
	60, 108,
	69, 108,
	70, 108,
	76, 108,
	78, 108,
	80, 108,
	81, 108,
	-2, 106,
	-1, 271,
	70, 3,
	-2, 121,
	-1, 275,
	70, 3,
	-2, 121,
	-1, 277,
	45, 3,
	46, 3,
	70, 3,
	-2, 121,
	-1, 281,
	70, 3,
	-2, 121,
	-1, 289,
	45, 3,
	46, 3,
	70, 3,
	-2, 121,
	-1, 296,
	14, 3,
	55, 3,
	70, 3,
	-2, 121,
}

const yyPrivate = 57344

const yyLast = 2351

var yyAct = [...]int{

	83, 163, 227, 10, 47, 228, 42, 166, 11, 240,
	202, 1, 267, 200, 6, 7, 84, 91, 160, 92,
	88, 82, 90, 114, 266, 93, 94, 95, 97, 99,
	6, 7, 89, 6, 7, 168, 250, 92, 104, 237,
	107, 262, 109, 247, 111, 245, 10, 205, 236, 190,
	115, 116, 249, 118, 119, 120, 121, 122, 123, 124,
	125, 126, 127, 128, 129, 130, 131, 132, 133, 134,
	135, 136, 137, 178, 205, 138, 139, 140, 141, 209,
	143, 144, 146, 205, 114, 142, 233, 2, 206, 217,
	145, 44, 108, 155, 105, 147, 101, 147, 154, 271,
	196, 147, 229, 230, 251, 158, 164, 147, 299, 161,
	297, 248, 295, 246, 146, 205, 170, 191, 292, 291,
	287, 278, 175, 102, 103, 147, 242, 226, 225, 224,
	221, 148, 113, 110, 152, 114, 257, 204, 150, 81,
	273, 179, 87, 167, 229, 230, 184, 5, 238, 197,
	8, 185, 46, 164, 244, 272, 215, 167, 199, 195,
	183, 194, 159, 186, 187, 198, 149, 100, 117, 85,
	265, 48, 162, 4, 212, 207, 208, 45, 216, 210,
	211, 252, 218, 219, 270, 222, 214, 171, 188, 17,
	172, 86, 231, 46, 234, 232, 112, 3, 173, 0,
	0, 0, 176, 0, 0, 243, 0, 0, 0, 0,
	0, 0, 0, 66, 67, 68, 69, 70, 71, 0,
	0, 0, 0, 57, 255, 0, 0, 0, 0, 0,
	256, 0, 80, 0, 0, 182, 0, 0, 0, 0,
	219, 189, 0, 264, 0, 0, 0, 259, 0, 201,
	203, 268, 269, 51, 0, 0, 76, 78, 0, 79,
	0, 74, 0, 0, 0, 0, 274, 0, 0, 0,
	0, 0, 279, 280, 286, 0, 0, 0, 0, 0,
	0, 0, 0, 285, 235, 294, 0, 288, 239, 290,
	241, 0, 0, 293, 0, 0, 0, 0, 0, 0,
	0, 298, 0, 0, 0, 0, 0, 0, 301, 22,
	23, 29, 0, 0, 34, 14, 9, 15, 43, 0,
	18, 0, 0, 0, 260, 261, 0, 0, 38, 30,
	31, 32, 16, 20, 0, 0, 0, 0, 0, 0,
	0, 0, 12, 13, 0, 277, 0, 0, 0, 21,
	281, 0, 39, 0, 40, 41, 0, 33, 0, 0,
	0, 19, 0, 0, 289, 0, 0, 0, 0, 24,
	28, 0, 0, 0, 36, 0, 0, 25, 26, 27,
	37, 0, 35, 0, 0, 6, 7, 60, 61, 63,
	65, 75, 77, 0, 0, 0, 0, 0, 0, 0,
	0, 66, 67, 68, 69, 70, 71, 0, 0, 72,
	73, 57, 58, 59, 0, 0, 0, 0, 0, 0,
	80, 0, 0, 0, 0, 0, 0, 0, 50, 0,
	284, 62, 64, 52, 53, 54, 55, 56, 0, 0,
	0, 51, 0, 0, 76, 78, 283, 79, 0, 74,
	60, 61, 63, 65, 75, 77, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 67, 68, 69, 70, 71,
	0, 0, 72, 73, 57, 58, 59, 0, 0, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 0, 0,
	0, 50, 193, 0, 62, 64, 52, 53, 54, 55,
	56, 0, 0, 0, 51, 0, 0, 76, 78, 0,
	79, 192, 74, 60, 61, 63, 65, 75, 77, 0,
	0, 0, 0, 0, 0, 0, 0, 66, 67, 68,
	69, 70, 71, 0, 0, 72, 73, 57, 58, 59,
	0, 0, 0, 0, 0, 0, 80, 0, 0, 0,
	0, 0, 0, 0, 50, 181, 0, 62, 64, 52,
	53, 54, 55, 56, 0, 0, 0, 51, 0, 0,
	76, 78, 0, 79, 180, 74, 60, 61, 63, 65,
	75, 77, 0, 0, 0, 0, 0, 0, 0, 0,
	66, 67, 68, 69, 70, 71, 0, 0, 72, 73,
	57, 58, 59, 0, 0, 0, 0, 0, 0, 80,
	0, 0, 0, 0, 0, 0, 0, 50, 0, 0,
	62, 64, 52, 53, 54, 55, 56, 0, 0, 0,
	51, 0, 0, 76, 78, 300, 79, 0, 74, 60,
	61, 63, 65, 75, 77, 0, 0, 0, 0, 0,
	0, 0, 0, 66, 67, 68, 69, 70, 71, 0,
	0, 72, 73, 57, 58, 59, 0, 0, 0, 0,
	0, 0, 80, 0, 0, 0, 0, 0, 0, 0,
	50, 0, 0, 62, 64, 52, 53, 54, 55, 56,
	0, 296, 0, 51, 0, 0, 76, 78, 0, 79,
	0, 74, 60, 61, 63, 65, 75, 77, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 67, 68, 69,
	70, 71, 0, 0, 72, 73, 57, 58, 59, 0,
	0, 0, 0, 0, 0, 80, 0, 0, 0, 0,
	0, 0, 0, 50, 0, 0, 62, 64, 52, 53,
	54, 55, 56, 0, 0, 0, 51, 0, 0, 76,
	78, 282, 79, 0, 74, 60, 61, 63, 65, 75,
	77, 0, 0, 0, 0, 0, 0, 0, 0, 66,
	67, 68, 69, 70, 71, 0, 0, 72, 73, 57,
	58, 59, 0, 0, 0, 0, 0, 0, 80, 0,
	0, 0, 0, 0, 0, 0, 50, 276, 0, 62,
	64, 52, 53, 54, 55, 56, 0, 0, 0, 51,
	0, 0, 76, 78, 0, 79, 0, 74, 60, 61,
	63, 65, 75, 77, 0, 0, 0, 0, 0, 0,
	0, 0, 66, 67, 68, 69, 70, 71, 0, 0,
	72, 73, 57, 58, 59, 0, 0, 0, 0, 0,
	0, 80, 0, 0, 0, 0, 0, 0, 0, 50,
	0, 0, 62, 64, 52, 53, 54, 55, 56, 0,
	275, 0, 51, 0, 0, 76, 78, 0, 79, 0,
	74, 60, 61, 63, 65, 75, 77, 0, 0, 0,
	0, 0, 0, 0, 0, 66, 67, 68, 69, 70,
	71, 0, 0, 72, 73, 57, 58, 59, 0, 0,
	0, 0, 0, 0, 80, 0, 0, 0, 0, 0,
	0, 0, 50, 0, 0, 62, 64, 52, 53, 54,
	55, 56, 0, 0, 0, 51, 0, 0, 76, 78,
	0, 79, 258, 74, 60, 61, 63, 65, 75, 77,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 67,
	68, 69, 70, 71, 0, 0, 72, 73, 57, 58,
	59, 0, 0, 0, 0, 0, 0, 80, 0, 0,
	0, 0, 0, 0, 0, 50, 0, 0, 62, 64,
	52, 53, 54, 55, 56, 0, 254, 0, 51, 0,
	0, 76, 78, 0, 79, 0, 74, 60, 61, 63,
	65, 75, 77, 0, 0, 0, 0, 0, 0, 0,
	0, 66, 67, 68, 69, 70, 71, 0, 0, 72,
	73, 57, 58, 59, 0, 0, 0, 0, 0, 0,
	80, 0, 0, 0, 0, 0, 0, 0, 50, 0,
	0, 62, 64, 52, 53, 54, 55, 56, 0, 0,
	0, 51, 0, 0, 76, 78, 0, 79, 253, 74,
	60, 61, 63, 65, 75, 77, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 67, 68, 69, 70, 71,
	0, 0, 72, 73, 57, 58, 59, 0, 0, 0,
	0, 0, 0, 80, 0, 0, 0, 223, 0, 0,
	0, 50, 0, 0, 62, 64, 52, 53, 54, 55,
	56, 0, 0, 0, 51, 0, 0, 76, 78, 0,
	79, 0, 74, 60, 61, 63, 65, 75, 77, 0,
	0, 0, 0, 0, 0, 0, 0, 66, 67, 68,
	69, 70, 71, 0, 0, 72, 73, 57, 58, 59,
	0, 0, 0, 0, 0, 0, 80, 0, 0, 0,
	0, 0, 0, 0, 50, 177, 0, 62, 64, 52,
	53, 54, 55, 56, 0, 0, 0, 51, 0, 0,
	76, 78, 0, 79, 0, 74, 60, 61, 63, 65,
	75, 77, 0, 0, 0, 0, 0, 0, 0, 0,
	66, 67, 68, 69, 70, 71, 0, 0, 72, 73,
	57, 58, 59, 0, 0, 0, 0, 0, 0, 80,
	0, 0, 0, 0, 0, 0, 0, 50, 0, 0,
	62, 64, 52, 53, 54, 55, 56, 0, 174, 0,
	51, 0, 0, 76, 78, 0, 79, 0, 74, 60,
	61, 63, 65, 75, 77, 0, 0, 0, 0, 0,
	0, 0, 0, 66, 67, 68, 69, 70, 71, 0,
	0, 72, 73, 57, 58, 59, 0, 0, 0, 0,
	0, 0, 80, 0, 0, 0, 0, 0, 0, 0,
	50, 0, 0, 62, 64, 52, 53, 54, 55, 56,
	0, 0, 0, 51, 0, 0, 76, 78, 165, 79,
	0, 74, 60, 61, 63, 65, 75, 77, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 67, 68, 69,
	70, 71, 0, 0, 72, 73, 57, 58, 59, 0,
	0, 0, 0, 0, 0, 80, 0, 0, 0, 0,
	0, 0, 0, 50, 153, 0, 62, 64, 52, 53,
	54, 55, 56, 0, 0, 0, 51, 0, 0, 76,
	78, 0, 79, 0, 74, 60, 61, 63, 65, 75,
	77, 0, 0, 0, 0, 0, 0, 0, 0, 66,
	67, 68, 69, 70, 71, 0, 0, 72, 73, 57,
	58, 59, 0, 0, 0, 0, 0, 0, 80, 0,
	0, 0, 0, 0, 0, 0, 50, 0, 0, 62,
	64, 52, 53, 54, 55, 56, 0, 151, 0, 51,
	0, 0, 76, 78, 0, 79, 0, 74, 60, 61,
	63, 65, 75, 77, 0, 0, 0, 0, 0, 0,
	0, 0, 66, 67, 68, 69, 70, 71, 0, 0,
	72, 73, 57, 58, 59, 0, 0, 0, 0, 0,
	0, 80, 0, 0, 0, 0, 0, 0, 49, 50,
	0, 0, 62, 64, 52, 53, 54, 55, 56, 0,
	0, 0, 51, 0, 0, 76, 78, 0, 79, 0,
	74, 60, 61, 63, 65, 75, 77, 0, 0, 0,
	0, 0, 0, 0, 0, 66, 67, 68, 69, 70,
	71, 0, 0, 72, 73, 57, 58, 59, 0, 0,
	0, 0, 0, 0, 80, 0, 0, 0, 0, 0,
	0, 0, 50, 0, 0, 62, 64, 52, 53, 54,
	55, 56, 0, 0, 0, 51, 0, 0, 76, 78,
	0, 79, 0, 74, 60, 61, 63, 65, 75, 77,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 67,
	68, 69, 70, 71, 0, 0, 72, 73, 57, 58,
	59, 0, 0, 0, 0, 0, 0, 80, 0, 0,
	0, 0, 0, 0, 0, 50, 0, 0, 62, 64,
	52, 53, 54, 55, 56, 0, 0, 0, 51, 0,
	0, 76, 169, 0, 79, 0, 74, 60, 61, 63,
	65, 75, 77, 0, 0, 0, 0, 0, 0, 0,
	0, 66, 67, 68, 69, 70, 71, 0, 0, 72,
	73, 57, 58, 59, 0, 0, 0, 0, 0, 0,
	80, 0, 0, 0, 0, 0, 0, 0, 50, 0,
	0, 62, 64, 52, 53, 54, 55, 56, 0, 0,
	0, 157, 0, 0, 76, 78, 0, 79, 0, 74,
	60, 61, 63, 65, 75, 77, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 67, 68, 69, 70, 71,
	0, 0, 72, 73, 57, 58, 59, 0, 0, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 0, 0,
	0, 50, 0, 0, 62, 64, 52, 53, 54, 55,
	56, 0, 0, 0, 156, 0, 0, 76, 78, 0,
	79, 0, 74, 60, 61, 63, 65, 0, 77, 0,
	0, 0, 0, 0, 0, 0, 0, 66, 67, 68,
	69, 70, 71, 0, 0, 72, 73, 57, 58, 59,
	0, 0, 0, 0, 0, 0, 80, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 62, 64, 52,
	53, 54, 55, 56, 0, 0, 0, 51, 0, 0,
	76, 78, 0, 79, 0, 74, 22, 23, 29, 0,
	0, 34, 14, 9, 15, 43, 0, 18, 0, 0,
	0, 0, 0, 0, 0, 38, 30, 31, 32, 16,
	20, 0, 0, 0, 0, 0, 0, 0, 0, 12,
	13, 0, 0, 0, 0, 0, 21, 0, 0, 39,
	0, 40, 41, 0, 33, 0, 0, 0, 19, 0,
	0, 0, 0, 0, 0, 0, 24, 28, 0, 0,
	0, 36, 0, 0, 25, 26, 27, 37, 0, 35,
	60, 61, 63, 65, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 67, 68, 69, 70, 71,
	0, 0, 72, 73, 57, 58, 59, 0, 0, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 62, 64, 52, 53, 54, 55,
	56, 63, 65, 0, 51, 0, 0, 76, 78, 0,
	79, 0, 74, 66, 67, 68, 69, 70, 71, 0,
	0, 72, 73, 57, 58, 59, 0, 0, 0, 0,
	0, 0, 80, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 62, 64, 52, 53, 54, 55, 56,
	220, 23, 29, 51, 0, 34, 76, 78, 0, 79,
	0, 74, 0, 0, 0, 0, 0, 0, 0, 38,
	30, 31, 32, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 39, 0, 40, 41, 0, 33, 0,
	0, 0, 0, 66, 67, 68, 69, 70, 71, 0,
	24, 28, 0, 57, 0, 36, 0, 0, 25, 26,
	27, 37, 80, 35, 263, 66, 67, 68, 69, 70,
	71, 0, 0, 72, 73, 57, 0, 54, 55, 56,
	0, 0, 0, 51, 80, 0, 76, 78, 0, 79,
	0, 74, 0, 0, 0, 0, 0, 52, 53, 54,
	55, 56, 22, 23, 29, 51, 0, 34, 76, 78,
	0, 79, 0, 74, 0, 0, 0, 0, 0, 0,
	0, 38, 30, 31, 32, 0, 0, 0, 0, 0,
	0, 0, 0, 220, 23, 29, 0, 0, 34, 0,
	0, 0, 0, 0, 0, 39, 0, 40, 41, 0,
	33, 0, 38, 30, 31, 32, 0, 0, 0, 0,
	0, 0, 24, 28, 213, 23, 29, 36, 0, 34,
	25, 26, 27, 37, 0, 35, 39, 0, 40, 41,
	0, 33, 0, 38, 30, 31, 32, 0, 0, 0,
	0, 0, 0, 24, 28, 106, 23, 29, 36, 0,
	34, 25, 26, 27, 37, 0, 35, 39, 0, 40,
	41, 0, 33, 0, 38, 30, 31, 32, 0, 0,
	0, 0, 0, 0, 24, 28, 98, 23, 29, 36,
	0, 34, 25, 26, 27, 37, 0, 35, 39, 0,
	40, 41, 0, 33, 0, 38, 30, 31, 32, 0,
	0, 0, 0, 0, 0, 24, 28, 96, 23, 29,
	36, 0, 34, 25, 26, 27, 37, 0, 35, 39,
	0, 40, 41, 0, 33, 0, 38, 30, 31, 32,
	0, 0, 0, 0, 0, 0, 24, 28, 0, 0,
	0, 36, 0, 0, 25, 26, 27, 37, 0, 35,
	39, 0, 40, 41, 0, 33, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 24, 28, 0,
	0, 0, 36, 0, 0, 25, 26, 27, 37, 0,
	35,
}
var yyPact = [...]int{

	-66, -1000, 1832, -66, -66, -1000, -1000, -1000, -1000, 167,
	1441, 82, -1000, -1000, 2118, 2118, 165, -1000, 138, 2118,
	-66, 2118, -58, -1000, 2118, 2118, 2118, 2273, 2242, -1000,
	-1000, -1000, -1000, -1000, 92, -66, -66, 2118, 19, 2211,
	17, 2118, 73, 2118, -1000, 305, -1000, 75, -1000, 2118,
	2118, 164, 2118, 2118, 2118, 2118, 2118, 2118, 2118, 2118,
	2118, 2118, 2118, 2118, 2118, 2118, 2118, 2118, 2118, 2118,
	2118, 2118, -1000, -1000, 2118, 2118, 2118, 2118, 2118, 2118,
	2118, 2118, 47, 1504, 1504, 62, 162, 81, 1378, 105,
	1315, 2118, 2118, 182, 182, 182, -58, 1693, -58, 1630,
	158, -57, 2118, 147, 1252, 153, -40, 1567, 139, 1504,
	-66, 1189, -1000, 2118, -66, 1504, 1126, -1000, 2032, 2032,
	182, 182, 182, 1504, 2054, 2054, 1942, 1942, 2054, 2054,
	2054, 2054, 1504, 1504, 1504, 1504, 1504, 1504, 1504, 1756,
	1504, 1893, 65, 496, 1504, -1000, 1504, -66, -66, 130,
	2118, -66, -66, -66, 41, 433, 157, 155, 24, 141,
	154, -47, -50, -1000, 78, -1000, 12, -1000, 2118, 2118,
	3, 153, 153, 2180, -66, -1000, 152, 2118, 13, -1000,
	-1000, 2118, 2149, 60, 2118, 1063, 59, 58, 57, 99,
	10, -1000, -1000, 2118, -1000, -1000, -66, -28, -37, 140,
	-66, -69, -66, 56, 2118, 150, -1000, 37, 35, -1000,
	-24, 44, 1504, -58, -1000, -1000, 1504, -1000, 1000, 1504,
	-58, -1000, 937, 2118, -1000, -1000, -1000, -1000, -1000, 2118,
	77, -1000, -1000, -1000, 874, -66, -66, -66, -35, 2006,
	-1000, 100, -1000, 1504, -1000, -52, -1000, -64, -1000, -1000,
	2118, 2118, 85, -1000, -66, 811, 748, -66, -1000, 51,
	-66, -66, -66, -1000, -1000, -1000, -1000, -1000, 685, 370,
	-1000, -66, -1000, 2118, 50, -66, -66, -66, -1000, 49,
	48, -66, -1000, -1000, 2118, 42, 622, -1000, 40, -66,
	-1000, -1000, -1000, 38, 559, -1000, -66, -1000, -1000, -1000,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 11, 197, 150, 189, 5, 2, 188, 184, 181,
	7, 0, 6, 8, 1, 172, 4, 87, 173, 147,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 9, 9, 8, 4, 4, 7, 7, 7,
	7, 7, 6, 5, 14, 15, 15, 15, 16, 16,
	16, 13, 13, 13, 10, 10, 12, 12, 12, 12,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 17, 17, 18, 18, 19, 19,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 2, 3, 4, 2, 3, 3,
	1, 1, 2, 2, 5, 1, 8, 9, 5, 5,
	5, 1, 0, 2, 4, 8, 6, 0, 2, 2,
	2, 2, 5, 4, 3, 0, 1, 4, 0, 1,
	4, 1, 4, 4, 1, 3, 0, 1, 4, 4,
	1, 1, 2, 2, 2, 2, 4, 2, 4, 1,
	1, 1, 1, 1, 5, 3, 7, 8, 8, 9,
	5, 6, 5, 6, 3, 4, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 2, 3, 3,
	3, 3, 5, 4, 6, 5, 5, 4, 6, 5,
	4, 4, 6, 6, 4, 5, 7, 7, 9, 3,
	2, 0, 1, 1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -17, -2, -18, -19, 80, 81, -3, 11,
	-11, -13, 37, 38, 10, 12, 27, -4, 15, 56,
	28, 44, 4, 5, 64, 72, 73, 74, 65, 6,
	24, 25, 26, 52, 9, 77, 69, 75, 23, 47,
	49, 50, -12, 13, -17, -18, -19, -16, 4, 57,
	58, 71, 63, 64, 65, 66, 67, 41, 42, 43,
	17, 18, 61, 19, 62, 20, 31, 32, 33, 34,
	35, 36, 39, 40, 79, 21, 74, 22, 75, 77,
	50, 57, -12, -11, -11, 4, 53, 4, -11, -1,
	-11, 75, 77, -11, -11, -11, 4, -11, 4, -11,
	75, 4, -17, -17, -11, 75, 4, -11, 75, -11,
	60, -11, -3, 57, 60, -11, -11, 4, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -12, -11, -11, -13, -11, 60, 69, 4,
	57, 69, 29, 59, -12, -11, 71, 71, -16, 4,
	75, -12, -15, -14, 6, 76, -10, 4, 75, 75,
	-10, 48, 51, -17, 69, -13, -17, 59, 8, 76,
	78, 59, -17, -1, 16, -11, -1, -1, -7, -17,
	8, 76, 78, 59, 4, 4, 76, 8, -16, 4,
	60, -17, 60, -17, 59, 71, 76, -12, -12, 76,
	-10, -10, -11, 4, -1, 4, -11, 76, -11, -11,
	4, 70, -11, 54, 70, 70, 70, -6, -5, 45,
	46, -6, -5, 76, -11, -17, 76, 76, 8, -17,
	78, -17, 70, -11, 4, 8, 76, 8, 76, 76,
	60, 60, -9, 78, 69, -11, -11, 59, 78, -1,
	-17, -17, 76, 78, -14, 70, 76, 76, -11, -11,
	-8, 14, 70, 55, -1, 69, 59, -17, 70, -1,
	-1, -17, 76, 76, 60, -1, -11, 70, -1, -17,
	-1, 70, 70, -1, -11, 70, 69, 70, -1, 70,
	76, -1,
}
var yyDef = [...]int{

	-2, -2, -2, 121, 122, 123, 125, 126, 4, 38,
	-2, 0, 10, 11, 46, 0, 0, 15, 0, 0,
	-2, 0, 50, 51, 0, 0, 0, 0, 0, 59,
	60, 61, 62, 63, 0, 121, 121, 0, 0, 0,
	0, 0, 0, 0, 2, -2, 124, 7, 39, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 96, 97, 0, 0, 0, 0, 46, 0,
	0, 46, 12, 47, 13, 0, 0, 0, 0, 0,
	0, 46, 0, 52, 53, 54, -2, 0, -2, 0,
	38, 0, 46, 35, 0, 0, 50, 0, 0, 120,
	121, 0, 5, 46, 121, 8, 0, 65, 76, 77,
	78, 79, 80, 81, 82, 83, -2, -2, 86, 87,
	88, 89, 90, 91, 92, 93, 94, 95, 98, 99,
	100, 101, 0, 0, 119, 9, -2, 121, -2, 0,
	0, -2, -2, 27, 0, 0, 0, 0, 0, 39,
	38, 121, 121, 36, 0, 74, 0, 44, 46, 46,
	0, 0, 0, 0, -2, 6, 0, 0, 0, 107,
	111, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 103, 110, 0, 56, 58, 121, 0, 0, 39,
	121, 0, 121, 0, 0, 0, 75, 0, 0, 114,
	0, 0, -2, -2, 22, 40, 64, 106, 0, 48,
	-2, 14, 0, 0, 18, 19, 20, 30, 31, 0,
	0, 28, 29, 102, 0, -2, 121, 121, 0, 0,
	70, 0, 72, 34, 45, 0, -2, 0, -2, 115,
	0, 0, 0, 113, -2, 0, 0, 121, 112, 0,
	-2, -2, 121, 71, 37, 73, -2, -2, 0, 0,
	23, -2, 26, 0, 0, -2, 121, -2, 66, 0,
	0, -2, 116, 117, 0, 0, 0, 16, 0, -2,
	33, 67, 68, 0, 0, 25, -2, 17, 32, 69,
	118, 24,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	81, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 72, 3, 3, 3, 67, 74, 3,
	75, 76, 65, 63, 60, 64, 71, 66, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 59, 80,
	62, 57, 61, 58, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 77, 3, 78, 73, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 69, 79, 70,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 68,
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
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:67
		{
			yyVAL.compstmt = nil
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:71
		{
			yyVAL.compstmt = yyDollar[1].stmts
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:76
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:83
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:90
		{
			if yyDollar[3].stmt != nil {
				yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[3].stmt)
				if l, ok := yylex.(*Lexer); ok {
					l.stmts = yyVAL.stmts
				}
			}
		}
	case 6:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:101
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].expr_many}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:106
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: nil}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:111
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "=", Rhss: []ast.Expr{yyDollar[3].expr}}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:115
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:119
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:124
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:129
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:134
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 14:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:139
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyDollar[2].tok.Lit, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:144
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 16:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:149
		{
			yyVAL.stmt = &ast.ForStmt{Var: yyDollar[3].tok.Lit, Value: yyDollar[5].expr, Stmts: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.y:154
		{
			yyVAL.stmt = &ast.NumForStmt{Name: yyDollar[2].tok.Lit, Expr1: yyDollar[4].expr, Expr2: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:159
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:164
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[2].compstmt, Catch: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:169
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:174
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:180
		{
			yyVAL.stmt_elsifs = []ast.Stmt{}
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:184
		{
			yyVAL.stmt_elsifs = append(yyDollar[1].stmt_elsifs, yyDollar[2].stmt_elsif)
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:190
		{
			yyVAL.stmt_elsif = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt}
		}
	case 25:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:196
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, ElseIf: yyDollar[5].stmt_elsifs, Else: yyDollar[7].compstmt}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 26:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:201
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, ElseIf: yyDollar[5].stmt_elsifs, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 27:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:207
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:211
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_case}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:215
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_default}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:219
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_case)
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:223
		{
			for _, stmt := range yyDollar[1].stmt_cases {
				if _, ok := stmt.(*ast.DefaultStmt); ok {
					yylex.Error("multiple default statement")
				}
			}
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_default)
		}
	case 32:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:234
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[5].compstmt}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:240
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyDollar[4].compstmt}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:246
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 35:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:251
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:255
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:259
		{
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[4].expr_pair)
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:264
		{
			yyVAL.expr_idents = []string{}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:268
		{
			yyVAL.expr_idents = []string{yyDollar[1].tok.Lit}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:272
		{
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, yyDollar[4].tok.Lit)
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:278
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:282
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 43:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:286
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:291
		{
			yyVAL.typ = ast.Type{Name: yyDollar[1].tok.Lit}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:295
		{
			yyVAL.typ = ast.Type{Name: yyDollar[1].typ.Name + "." + yyDollar[3].tok.Lit}
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:300
		{
			yyVAL.exprs = nil
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:304
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:308
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 49:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:312
		{
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:318
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:323
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:328
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:333
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:338
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:343
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:348
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:353
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:358
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:363
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:368
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:373
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:378
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:383
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 64:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:388
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[1].expr, Lhs: yyDollar[3].expr, Rhs: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:393
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: yyDollar[3].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 66:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:398
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyDollar[3].expr_idents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 67:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:403
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyDollar[3].tok.Lit}, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 68:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:408
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Args: yyDollar[4].expr_idents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.y:413
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Args: []string{yyDollar[4].tok.Lit}, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 70:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:418
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 71:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:423
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 72:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:428
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
	case 73:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:437
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
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:446
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:451
		{
			yyVAL.expr = &ast.NewExpr{Type: yyDollar[3].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:456
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:461
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:466
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:471
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:476
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:481
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:486
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:491
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:496
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:501
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:506
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:511
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:516
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:521
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:526
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:531
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:536
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:541
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:546
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:551
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:556
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:561
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:566
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:571
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:576
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:581
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:586
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 103:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:591
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 104:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:596
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 105:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:601
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 106:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:606
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 107:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:611
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 108:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:616
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 109:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:621
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 110:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:626
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 111:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:631
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 112:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:636
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 113:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:641
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 114:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:646
		{
			yyVAL.expr = &ast.MakeExpr{Type: yyDollar[3].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 115:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:651
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].typ.Name, SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 116:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:656
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].typ.Name, SizeExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 117:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:661
		{
			yyVAL.expr = &ast.MakeArrayExpr{Type: yyDollar[4].typ.Name, LenExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 118:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.y:666
		{
			yyVAL.expr = &ast.MakeArrayExpr{Type: yyDollar[4].typ.Name, LenExpr: yyDollar[6].expr, CapExpr: yyDollar[8].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:671
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:676
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:687
		{
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:690
		{
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:695
		{
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:698
		{
		}
	}
	goto yystack /* stack new state and value */
}
