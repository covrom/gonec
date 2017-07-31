//line .\parser\parser.y:1
package parser

import __yyfmt__ "fmt"

//line .\parser\parser.y:3
import (
	"github.com/covrom/gonec/ast"
)

//line .\parser\parser.y:27
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
const TERNARY = 57399
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
	"TERNARY",
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

//line .\parser\parser.y:702

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 3,
	-2, 121,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	61, 46,
	-2, 1,
	-1, 10,
	61, 47,
	-2, 21,
	-1, 20,
	29, 3,
	-2, 121,
	-1, 46,
	61, 46,
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
	58, 55,
	60, 55,
	61, 55,
	70, 55,
	71, 55,
	76, 55,
	79, 55,
	81, 55,
	82, 55,
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
	58, 57,
	60, 57,
	61, 57,
	70, 57,
	71, 57,
	76, 57,
	79, 57,
	81, 57,
	82, 57,
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
	61, 47,
	-2, 41,
	-1, 148,
	71, 3,
	-2, 121,
	-1, 151,
	71, 3,
	-2, 121,
	-1, 152,
	71, 3,
	-2, 121,
	-1, 175,
	14, 3,
	55, 3,
	71, 3,
	-2, 121,
	-1, 213,
	61, 48,
	-2, 42,
	-1, 214,
	1, 43,
	14, 43,
	29, 43,
	45, 43,
	46, 43,
	55, 43,
	58, 43,
	61, 49,
	71, 43,
	81, 43,
	82, 43,
	-2, 50,
	-1, 220,
	1, 49,
	8, 49,
	14, 49,
	29, 49,
	45, 49,
	46, 49,
	55, 49,
	61, 49,
	71, 49,
	76, 49,
	79, 49,
	81, 49,
	82, 49,
	-2, 50,
	-1, 236,
	71, 3,
	-2, 121,
	-1, 247,
	1, 105,
	8, 105,
	14, 105,
	29, 105,
	45, 105,
	46, 105,
	54, 105,
	55, 105,
	58, 105,
	60, 105,
	61, 105,
	70, 105,
	71, 105,
	76, 105,
	79, 105,
	81, 105,
	82, 105,
	-2, 103,
	-1, 249,
	1, 109,
	8, 109,
	14, 109,
	29, 109,
	45, 109,
	46, 109,
	54, 109,
	55, 109,
	58, 109,
	60, 109,
	61, 109,
	70, 109,
	71, 109,
	76, 109,
	79, 109,
	81, 109,
	82, 109,
	-2, 107,
	-1, 255,
	71, 3,
	-2, 121,
	-1, 262,
	71, 3,
	-2, 121,
	-1, 263,
	71, 3,
	-2, 121,
	-1, 268,
	1, 104,
	8, 104,
	14, 104,
	29, 104,
	45, 104,
	46, 104,
	54, 104,
	55, 104,
	58, 104,
	60, 104,
	61, 104,
	70, 104,
	71, 104,
	76, 104,
	79, 104,
	81, 104,
	82, 104,
	-2, 102,
	-1, 269,
	1, 108,
	8, 108,
	14, 108,
	29, 108,
	45, 108,
	46, 108,
	54, 108,
	55, 108,
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
	-1, 273,
	71, 3,
	-2, 121,
	-1, 277,
	71, 3,
	-2, 121,
	-1, 279,
	45, 3,
	46, 3,
	71, 3,
	-2, 121,
	-1, 284,
	71, 3,
	-2, 121,
	-1, 292,
	45, 3,
	46, 3,
	71, 3,
	-2, 121,
	-1, 299,
	14, 3,
	55, 3,
	71, 3,
	-2, 121,
}

const yyPrivate = 57344

const yyLast = 2536

var yyAct = [...]int{

	83, 164, 227, 10, 48, 241, 228, 11, 43, 115,
	203, 1, 102, 201, 161, 167, 84, 6, 7, 109,
	88, 106, 90, 82, 238, 93, 94, 95, 97, 99,
	6, 7, 89, 6, 7, 100, 91, 92, 206, 105,
	206, 108, 210, 110, 207, 112, 269, 10, 169, 92,
	268, 116, 248, 118, 119, 120, 121, 122, 123, 124,
	125, 126, 127, 128, 129, 130, 131, 132, 133, 134,
	135, 136, 137, 251, 264, 138, 139, 140, 141, 2,
	143, 144, 146, 45, 206, 101, 246, 142, 250, 145,
	237, 233, 115, 155, 66, 67, 68, 69, 70, 71,
	154, 217, 72, 73, 57, 147, 159, 197, 190, 165,
	252, 302, 162, 80, 300, 146, 103, 104, 229, 230,
	249, 206, 176, 298, 273, 171, 295, 52, 53, 54,
	55, 56, 294, 290, 148, 51, 281, 178, 76, 147,
	78, 79, 243, 74, 226, 225, 224, 221, 147, 114,
	111, 185, 115, 258, 247, 205, 150, 81, 87, 196,
	183, 147, 8, 186, 187, 275, 199, 66, 67, 68,
	69, 70, 71, 168, 267, 213, 191, 57, 208, 209,
	152, 274, 218, 219, 184, 222, 80, 215, 211, 212,
	147, 174, 231, 5, 234, 177, 232, 239, 47, 198,
	229, 230, 54, 55, 56, 179, 244, 86, 51, 113,
	165, 76, 245, 78, 79, 216, 74, 172, 168, 200,
	173, 195, 194, 160, 256, 149, 117, 182, 85, 49,
	257, 4, 163, 189, 253, 46, 260, 272, 188, 17,
	47, 219, 202, 204, 266, 3, 0, 0, 261, 0,
	0, 0, 270, 271, 66, 67, 68, 69, 70, 71,
	0, 0, 0, 0, 57, 0, 0, 276, 0, 0,
	0, 0, 0, 80, 282, 283, 289, 236, 0, 0,
	0, 240, 0, 242, 0, 288, 0, 0, 297, 291,
	0, 293, 0, 0, 0, 51, 296, 0, 76, 0,
	78, 79, 0, 74, 301, 0, 0, 0, 0, 0,
	0, 304, 0, 0, 0, 0, 0, 262, 263, 0,
	0, 0, 0, 0, 0, 0, 22, 23, 29, 0,
	0, 35, 14, 9, 15, 44, 0, 18, 279, 0,
	0, 0, 0, 0, 284, 39, 30, 31, 32, 16,
	20, 0, 0, 0, 0, 0, 0, 0, 292, 12,
	13, 0, 0, 0, 0, 0, 21, 0, 0, 40,
	0, 41, 42, 0, 33, 0, 0, 0, 19, 34,
	0, 0, 0, 0, 0, 0, 0, 24, 28, 0,
	0, 0, 37, 0, 0, 25, 26, 27, 0, 38,
	36, 0, 0, 6, 7, 60, 61, 63, 65, 75,
	77, 0, 0, 0, 0, 0, 0, 0, 0, 66,
	67, 68, 69, 70, 71, 0, 0, 72, 73, 57,
	58, 59, 0, 0, 0, 0, 0, 0, 80, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 287,
	62, 64, 52, 53, 54, 55, 56, 0, 0, 0,
	51, 0, 0, 76, 286, 78, 79, 0, 74, 60,
	61, 63, 65, 75, 77, 0, 0, 0, 0, 0,
	0, 0, 0, 66, 67, 68, 69, 70, 71, 0,
	0, 72, 73, 57, 58, 59, 0, 0, 0, 0,
	0, 0, 80, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 193, 0, 62, 64, 52, 53, 54, 55,
	56, 0, 0, 0, 51, 0, 0, 76, 0, 78,
	79, 192, 74, 60, 61, 63, 65, 75, 77, 0,
	0, 0, 0, 0, 0, 0, 0, 66, 67, 68,
	69, 70, 71, 0, 0, 72, 73, 57, 58, 59,
	0, 0, 0, 0, 0, 0, 80, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 181, 0, 62, 64,
	52, 53, 54, 55, 56, 0, 0, 0, 51, 0,
	0, 76, 0, 78, 79, 180, 74, 60, 61, 63,
	65, 75, 77, 0, 0, 0, 0, 0, 0, 0,
	0, 66, 67, 68, 69, 70, 71, 0, 0, 72,
	73, 57, 58, 59, 0, 0, 0, 0, 0, 0,
	80, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 62, 64, 52, 53, 54, 55, 56, 0,
	0, 0, 51, 0, 0, 76, 303, 78, 79, 0,
	74, 60, 61, 63, 65, 75, 77, 0, 0, 0,
	0, 0, 0, 0, 0, 66, 67, 68, 69, 70,
	71, 0, 0, 72, 73, 57, 58, 59, 0, 0,
	0, 0, 0, 0, 80, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 62, 64, 52, 53,
	54, 55, 56, 0, 299, 0, 51, 0, 0, 76,
	0, 78, 79, 0, 74, 60, 61, 63, 65, 75,
	77, 0, 0, 0, 0, 0, 0, 0, 0, 66,
	67, 68, 69, 70, 71, 0, 0, 72, 73, 57,
	58, 59, 0, 0, 0, 0, 0, 0, 80, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	62, 64, 52, 53, 54, 55, 56, 0, 0, 0,
	51, 0, 0, 76, 285, 78, 79, 0, 74, 60,
	61, 63, 65, 75, 77, 0, 0, 0, 0, 0,
	0, 0, 0, 66, 67, 68, 69, 70, 71, 0,
	0, 72, 73, 57, 58, 59, 0, 0, 0, 0,
	0, 0, 80, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 62, 64, 52, 53, 54, 55,
	56, 0, 0, 0, 51, 0, 0, 76, 280, 78,
	79, 0, 74, 60, 61, 63, 65, 75, 77, 0,
	0, 0, 0, 0, 0, 0, 0, 66, 67, 68,
	69, 70, 71, 0, 0, 72, 73, 57, 58, 59,
	0, 0, 0, 0, 0, 0, 80, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 278, 0, 62, 64,
	52, 53, 54, 55, 56, 0, 0, 0, 51, 0,
	0, 76, 0, 78, 79, 0, 74, 60, 61, 63,
	65, 75, 77, 0, 0, 0, 0, 0, 0, 0,
	0, 66, 67, 68, 69, 70, 71, 0, 0, 72,
	73, 57, 58, 59, 0, 0, 0, 0, 0, 0,
	80, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 62, 64, 52, 53, 54, 55, 56, 0,
	277, 0, 51, 0, 0, 76, 0, 78, 79, 0,
	74, 60, 61, 63, 65, 75, 77, 0, 0, 0,
	0, 0, 0, 0, 0, 66, 67, 68, 69, 70,
	71, 0, 0, 72, 73, 57, 58, 59, 0, 0,
	0, 0, 0, 0, 80, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 62, 64, 52, 53,
	54, 55, 56, 0, 0, 0, 51, 0, 0, 76,
	0, 78, 79, 259, 74, 60, 61, 63, 65, 75,
	77, 0, 0, 0, 0, 0, 0, 0, 0, 66,
	67, 68, 69, 70, 71, 0, 0, 72, 73, 57,
	58, 59, 0, 0, 0, 0, 0, 0, 80, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	62, 64, 52, 53, 54, 55, 56, 0, 255, 0,
	51, 0, 0, 76, 0, 78, 79, 0, 74, 60,
	61, 63, 65, 75, 77, 0, 0, 0, 0, 0,
	0, 0, 0, 66, 67, 68, 69, 70, 71, 0,
	0, 72, 73, 57, 58, 59, 0, 0, 0, 0,
	0, 0, 80, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 62, 64, 52, 53, 54, 55,
	56, 0, 0, 0, 51, 0, 0, 76, 0, 78,
	79, 254, 74, 60, 61, 63, 65, 75, 77, 0,
	0, 0, 0, 0, 0, 0, 0, 66, 67, 68,
	69, 70, 71, 0, 0, 72, 73, 57, 58, 59,
	0, 0, 0, 0, 0, 0, 80, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 235, 62, 64,
	52, 53, 54, 55, 56, 0, 0, 0, 51, 0,
	0, 76, 0, 78, 79, 0, 74, 60, 61, 63,
	65, 75, 77, 0, 0, 0, 0, 0, 0, 0,
	0, 66, 67, 68, 69, 70, 71, 0, 0, 72,
	73, 57, 58, 59, 0, 0, 0, 0, 0, 0,
	80, 0, 0, 0, 223, 0, 0, 0, 0, 0,
	0, 0, 62, 64, 52, 53, 54, 55, 56, 0,
	0, 0, 51, 0, 0, 76, 0, 78, 79, 0,
	74, 60, 61, 63, 65, 75, 77, 0, 0, 0,
	0, 0, 0, 0, 0, 66, 67, 68, 69, 70,
	71, 0, 0, 72, 73, 57, 58, 59, 0, 0,
	0, 0, 0, 0, 80, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 62, 64, 52, 53,
	54, 55, 56, 0, 175, 0, 51, 0, 0, 76,
	0, 78, 79, 0, 74, 60, 61, 63, 65, 75,
	77, 0, 0, 0, 0, 0, 0, 0, 0, 66,
	67, 68, 69, 70, 71, 0, 0, 72, 73, 57,
	58, 59, 0, 0, 0, 0, 0, 0, 80, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	62, 64, 52, 53, 54, 55, 56, 0, 0, 0,
	51, 0, 0, 76, 166, 78, 79, 0, 74, 60,
	61, 63, 65, 75, 77, 0, 0, 0, 0, 0,
	0, 0, 0, 66, 67, 68, 69, 70, 71, 0,
	0, 72, 73, 57, 58, 59, 0, 0, 0, 0,
	0, 0, 80, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 158, 62, 64, 52, 53, 54, 55,
	56, 0, 0, 0, 51, 0, 0, 76, 0, 78,
	79, 0, 74, 60, 61, 63, 65, 75, 77, 0,
	0, 0, 0, 0, 0, 0, 0, 66, 67, 68,
	69, 70, 71, 0, 0, 72, 73, 57, 58, 59,
	0, 0, 0, 0, 0, 0, 80, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 153, 0, 62, 64,
	52, 53, 54, 55, 56, 0, 0, 0, 51, 0,
	0, 76, 0, 78, 79, 0, 74, 60, 61, 63,
	65, 75, 77, 0, 0, 0, 0, 0, 0, 0,
	0, 66, 67, 68, 69, 70, 71, 0, 0, 72,
	73, 57, 58, 59, 0, 0, 0, 0, 0, 0,
	80, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 62, 64, 52, 53, 54, 55, 56, 0,
	151, 0, 51, 0, 0, 76, 0, 78, 79, 0,
	74, 60, 61, 63, 65, 75, 77, 0, 0, 0,
	0, 0, 0, 0, 0, 66, 67, 68, 69, 70,
	71, 0, 0, 72, 73, 57, 58, 59, 0, 0,
	0, 0, 0, 0, 80, 0, 0, 0, 0, 0,
	0, 0, 50, 0, 0, 0, 62, 64, 52, 53,
	54, 55, 56, 0, 0, 0, 51, 0, 0, 76,
	0, 78, 79, 0, 74, 60, 61, 63, 65, 75,
	77, 0, 0, 0, 0, 0, 0, 0, 0, 66,
	67, 68, 69, 70, 71, 0, 0, 72, 73, 57,
	58, 59, 0, 0, 0, 0, 0, 0, 80, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	62, 64, 52, 53, 54, 55, 56, 0, 0, 0,
	51, 0, 0, 76, 0, 78, 79, 0, 74, 60,
	61, 63, 65, 75, 77, 0, 0, 0, 0, 0,
	0, 0, 0, 66, 67, 68, 69, 70, 71, 0,
	0, 72, 73, 57, 58, 59, 0, 0, 0, 0,
	0, 0, 80, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 62, 64, 52, 53, 54, 55,
	56, 0, 0, 0, 51, 0, 0, 76, 0, 170,
	79, 0, 74, 60, 61, 63, 65, 75, 77, 0,
	0, 0, 0, 0, 0, 0, 0, 66, 67, 68,
	69, 70, 71, 0, 0, 72, 73, 57, 58, 59,
	0, 0, 0, 0, 0, 0, 80, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 62, 64,
	52, 53, 54, 55, 56, 0, 0, 0, 157, 0,
	0, 76, 0, 78, 79, 0, 74, 60, 61, 63,
	65, 75, 77, 0, 0, 0, 0, 0, 0, 0,
	0, 66, 67, 68, 69, 70, 71, 0, 0, 72,
	73, 57, 58, 59, 0, 0, 0, 0, 0, 0,
	80, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 62, 64, 52, 53, 54, 55, 56, 0,
	0, 0, 156, 0, 0, 76, 0, 78, 79, 0,
	74, 22, 23, 29, 0, 0, 35, 14, 9, 15,
	44, 0, 18, 0, 0, 0, 0, 0, 0, 0,
	39, 30, 31, 32, 16, 20, 0, 0, 0, 0,
	0, 0, 0, 0, 12, 13, 0, 0, 0, 0,
	0, 21, 0, 0, 40, 0, 41, 42, 0, 33,
	0, 0, 0, 19, 34, 0, 0, 0, 0, 0,
	0, 0, 24, 28, 0, 0, 0, 37, 0, 0,
	25, 26, 27, 0, 38, 36, 60, 61, 63, 65,
	0, 77, 0, 0, 0, 0, 0, 0, 0, 0,
	66, 67, 68, 69, 70, 71, 0, 0, 72, 73,
	57, 58, 59, 0, 0, 0, 0, 0, 0, 80,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 62, 64, 52, 53, 54, 55, 56, 0, 0,
	0, 51, 0, 0, 76, 0, 78, 79, 0, 74,
	60, 61, 63, 65, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 67, 68, 69, 70, 71,
	0, 0, 72, 73, 57, 58, 59, 0, 0, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 62, 64, 52, 53, 54,
	55, 56, 63, 65, 0, 51, 0, 0, 76, 0,
	78, 79, 0, 74, 66, 67, 68, 69, 70, 71,
	0, 0, 72, 73, 57, 58, 59, 0, 0, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 62, 64, 52, 53, 54,
	55, 56, 220, 23, 29, 51, 0, 35, 76, 0,
	78, 79, 0, 74, 0, 0, 0, 0, 0, 0,
	0, 39, 30, 31, 32, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 22, 23, 29,
	0, 0, 35, 0, 0, 40, 0, 41, 42, 0,
	33, 0, 0, 0, 0, 34, 39, 30, 31, 32,
	0, 0, 0, 24, 28, 0, 0, 0, 37, 0,
	0, 25, 26, 27, 0, 38, 36, 265, 0, 0,
	40, 0, 41, 42, 0, 33, 0, 0, 0, 0,
	34, 0, 0, 0, 0, 220, 23, 29, 24, 28,
	35, 0, 0, 37, 0, 0, 25, 26, 27, 0,
	38, 36, 0, 0, 39, 30, 31, 32, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	214, 23, 29, 0, 0, 35, 0, 0, 40, 0,
	41, 42, 0, 33, 0, 0, 0, 0, 34, 39,
	30, 31, 32, 0, 0, 0, 24, 28, 0, 0,
	0, 37, 0, 0, 25, 26, 27, 0, 38, 36,
	0, 0, 0, 40, 0, 41, 42, 0, 33, 0,
	0, 0, 0, 34, 0, 0, 0, 0, 107, 23,
	29, 24, 28, 35, 0, 0, 37, 0, 0, 25,
	26, 27, 0, 38, 36, 0, 0, 39, 30, 31,
	32, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 98, 23, 29, 0, 0, 35, 0,
	0, 40, 0, 41, 42, 0, 33, 0, 0, 0,
	0, 34, 39, 30, 31, 32, 0, 0, 0, 24,
	28, 0, 0, 0, 37, 0, 0, 25, 26, 27,
	0, 38, 36, 0, 0, 0, 40, 0, 41, 42,
	0, 33, 0, 0, 0, 0, 34, 0, 0, 0,
	0, 96, 23, 29, 24, 28, 35, 0, 0, 37,
	0, 0, 25, 26, 27, 0, 38, 36, 0, 0,
	39, 30, 31, 32, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 40, 0, 41, 42, 0, 33,
	0, 0, 0, 0, 34, 0, 0, 0, 0, 0,
	0, 0, 24, 28, 0, 0, 0, 37, 0, 0,
	25, 26, 27, 0, 38, 36,
}
var yyPact = [...]int{

	-64, -1000, 1937, -64, -64, -1000, -1000, -1000, -1000, 225,
	1604, 99, -1000, -1000, 2213, 2213, 224, -1000, 154, 2213,
	-64, 2213, -41, -1000, 2213, 2213, 2213, 2457, 2399, -1000,
	-1000, -1000, -1000, -1000, 2213, 8, -64, -64, 2213, -56,
	2364, -58, 2213, 89, 2213, -1000, 322, -1000, 91, -1000,
	2213, 222, 2213, 2213, 2213, 2213, 2213, 2213, 2213, 2213,
	2213, 2213, 2213, 2213, 2213, 2213, 2213, 2213, 2213, 2213,
	2213, 2213, -1000, -1000, 2213, 2213, 2213, 2213, 2213, 2213,
	2213, 2213, 87, 1668, 1668, 64, 221, 98, 1540, 151,
	1476, 2213, 2213, 223, 223, 223, -41, 1860, -41, 1796,
	1412, 219, -63, 2213, 204, 1348, 214, -29, 1732, 169,
	1668, -64, 1284, -1000, 2213, -64, 1668, -1000, 136, 136,
	223, 223, 223, 1668, 63, 63, 2113, 2113, 63, 63,
	63, 63, 1668, 1668, 1668, 1668, 1668, 1668, 1668, 1999,
	1668, 2063, 129, 516, 1668, -1000, 1668, -64, -64, 168,
	2213, -64, -64, -64, 100, 452, 218, 217, 2213, 31,
	191, 215, -48, -51, -1000, 95, -1000, -32, -1000, 2213,
	2213, -34, 214, 214, 2306, -64, -1000, 211, 25, -1000,
	-1000, 2213, 2271, 76, 2213, 1220, 75, 74, 73, 155,
	15, -1000, -1000, 2213, -1000, -1000, 1156, -64, 14, -52,
	189, -64, -74, -64, 71, 2213, 208, -1000, 78, 44,
	-1000, 12, 49, 1668, -41, -1000, -1000, -1000, 1092, 1668,
	-41, -1000, 1028, 2213, -1000, -1000, -1000, -1000, -1000, 2213,
	93, -1000, -1000, -1000, 964, 2213, -64, -64, -64, -2,
	2178, -1000, 103, -1000, 1668, -1000, -26, -1000, -30, -1000,
	-1000, 2213, 2213, 110, -1000, -64, 900, 836, -64, -1000,
	772, 65, -64, -64, -64, -1000, -1000, -1000, -1000, -1000,
	708, 388, -1000, -64, -1000, 2213, 62, -64, -64, -64,
	-1000, -1000, 61, 55, -64, -1000, -1000, 2213, 52, 644,
	-1000, 43, -64, -1000, -1000, -1000, 40, 580, -1000, -64,
	-1000, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 11, 245, 162, 239, 6, 2, 238, 237, 234,
	15, 0, 8, 7, 1, 232, 4, 79, 231, 193,
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
	1, 1, 1, 1, 7, 3, 7, 8, 8, 9,
	5, 6, 5, 6, 3, 4, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 2, 3, 3,
	3, 3, 5, 4, 6, 5, 5, 4, 6, 5,
	4, 4, 6, 6, 4, 5, 7, 7, 9, 3,
	2, 0, 1, 1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -17, -2, -18, -19, 81, 82, -3, 11,
	-11, -13, 37, 38, 10, 12, 27, -4, 15, 56,
	28, 44, 4, 5, 65, 73, 74, 75, 66, 6,
	24, 25, 26, 52, 57, 9, 78, 70, 77, 23,
	47, 49, 50, -12, 13, -17, -18, -19, -16, 4,
	58, 72, 64, 65, 66, 67, 68, 41, 42, 43,
	17, 18, 62, 19, 63, 20, 31, 32, 33, 34,
	35, 36, 39, 40, 80, 21, 75, 22, 77, 78,
	50, 58, -12, -11, -11, 4, 53, 4, -11, -1,
	-11, 77, 78, -11, -11, -11, 4, -11, 4, -11,
	-11, 77, 4, -17, -17, -11, 77, 4, -11, 77,
	-11, 61, -11, -3, 58, 61, -11, 4, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -12, -11, -11, -13, -11, 61, 70, 4,
	58, 70, 29, 60, -12, -11, 72, 72, 61, -16,
	4, 77, -12, -15, -14, 6, 76, -10, 4, 77,
	77, -10, 48, 51, -17, 70, -13, -17, 8, 76,
	79, 60, -17, -1, 16, -11, -1, -1, -7, -17,
	8, 76, 79, 60, 4, 4, -11, 76, 8, -16,
	4, 61, -17, 61, -17, 60, 72, 76, -12, -12,
	76, -10, -10, -11, 4, -1, 4, 76, -11, -11,
	4, 71, -11, 54, 71, 71, 71, -6, -5, 45,
	46, -6, -5, 76, -11, 61, -17, 76, 76, 8,
	-17, 79, -17, 71, -11, 4, 8, 76, 8, 76,
	76, 61, 61, -9, 79, 70, -11, -11, 60, 79,
	-11, -1, -17, -17, 76, 79, -14, 71, 76, 76,
	-11, -11, -8, 14, 71, 55, -1, 70, 60, -17,
	76, 71, -1, -1, -17, 76, 76, 61, -1, -11,
	71, -1, -17, -1, 71, 71, -1, -11, 71, 70,
	71, -1, 71, 76, -1,
}
var yyDef = [...]int{

	-2, -2, -2, 121, 122, 123, 125, 126, 4, 38,
	-2, 0, 10, 11, 46, 0, 0, 15, 0, 0,
	-2, 0, 50, 51, 0, 0, 0, 0, 0, 59,
	60, 61, 62, 63, 0, 0, 121, 121, 0, 0,
	0, 0, 0, 0, 0, 2, -2, 124, 7, 39,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 96, 97, 0, 0, 0, 0, 46, 0,
	0, 46, 12, 47, 13, 0, 0, 0, 0, 0,
	0, 46, 0, 52, 53, 54, -2, 0, -2, 0,
	0, 38, 0, 46, 35, 0, 0, 50, 0, 0,
	120, 121, 0, 5, 46, 121, 8, 65, 76, 77,
	78, 79, 80, 81, 82, 83, -2, -2, 86, 87,
	88, 89, 90, 91, 92, 93, 94, 95, 98, 99,
	100, 101, 0, 0, 119, 9, -2, 121, -2, 0,
	0, -2, -2, 27, 0, 0, 0, 0, 0, 0,
	39, 38, 121, 121, 36, 0, 74, 0, 44, 46,
	46, 0, 0, 0, 0, -2, 6, 0, 0, 107,
	111, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 103, 110, 0, 56, 58, 0, 121, 0, 0,
	39, 121, 0, 121, 0, 0, 0, 75, 0, 0,
	114, 0, 0, -2, -2, 22, 40, 106, 0, 48,
	-2, 14, 0, 0, 18, 19, 20, 30, 31, 0,
	0, 28, 29, 102, 0, 0, -2, 121, 121, 0,
	0, 70, 0, 72, 34, 45, 0, -2, 0, -2,
	115, 0, 0, 0, 113, -2, 0, 0, 121, 112,
	0, 0, -2, -2, 121, 71, 37, 73, -2, -2,
	0, 0, 23, -2, 26, 0, 0, -2, 121, -2,
	64, 66, 0, 0, -2, 116, 117, 0, 0, 0,
	16, 0, -2, 33, 67, 68, 0, 0, 25, -2,
	17, 32, 69, 118, 24,
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
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:67
		{
			yyVAL.compstmt = nil
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:71
		{
			yyVAL.compstmt = yyDollar[1].stmts
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line .\parser\parser.y:76
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:83
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:90
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
		//line .\parser\parser.y:101
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].expr_many}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:106
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: nil}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:111
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "=", Rhss: []ast.Expr{yyDollar[3].expr}}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:115
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:119
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:124
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:129
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:134
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 14:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:139
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyDollar[2].tok.Lit, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:144
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 16:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line .\parser\parser.y:149
		{
			yyVAL.stmt = &ast.ForStmt{Var: yyDollar[3].tok.Lit, Value: yyDollar[5].expr, Stmts: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line .\parser\parser.y:154
		{
			yyVAL.stmt = &ast.NumForStmt{Name: yyDollar[2].tok.Lit, Expr1: yyDollar[4].expr, Expr2: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:159
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:164
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[2].compstmt, Catch: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:169
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:174
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line .\parser\parser.y:180
		{
			yyVAL.stmt_elsifs = []ast.Stmt{}
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:184
		{
			yyVAL.stmt_elsifs = append(yyDollar[1].stmt_elsifs, yyDollar[2].stmt_elsif)
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:190
		{
			yyVAL.stmt_elsif = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt}
		}
	case 25:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line .\parser\parser.y:196
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, ElseIf: yyDollar[5].stmt_elsifs, Else: yyDollar[7].compstmt}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 26:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line .\parser\parser.y:201
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, ElseIf: yyDollar[5].stmt_elsifs, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 27:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line .\parser\parser.y:207
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:211
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_case}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:215
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_default}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:219
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_case)
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:223
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
		//line .\parser\parser.y:234
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[5].compstmt}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:240
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyDollar[4].compstmt}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:246
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 35:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line .\parser\parser.y:251
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:255
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:259
		{
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[4].expr_pair)
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line .\parser\parser.y:264
		{
			yyVAL.expr_idents = []string{}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:268
		{
			yyVAL.expr_idents = []string{yyDollar[1].tok.Lit}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:272
		{
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, yyDollar[4].tok.Lit)
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:278
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:282
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 43:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:286
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:291
		{
			yyVAL.typ = ast.Type{Name: yyDollar[1].tok.Lit}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:295
		{
			yyVAL.typ = ast.Type{Name: yyDollar[1].typ.Name + "." + yyDollar[3].tok.Lit}
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line .\parser\parser.y:300
		{
			yyVAL.exprs = nil
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:304
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:308
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 49:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:312
		{
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:318
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:323
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:328
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:333
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:338
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:343
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:348
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:353
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:358
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:363
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:368
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:373
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:378
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:383
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 64:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line .\parser\parser.y:388
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[2].expr, Lhs: yyDollar[4].expr, Rhs: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:393
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: yyDollar[3].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 66:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line .\parser\parser.y:398
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyDollar[3].expr_idents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 67:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line .\parser\parser.y:403
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyDollar[3].tok.Lit}, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 68:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line .\parser\parser.y:408
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Args: yyDollar[4].expr_idents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line .\parser\parser.y:413
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Args: []string{yyDollar[4].tok.Lit}, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 70:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:418
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 71:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line .\parser\parser.y:423
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 72:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:428
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
		//line .\parser\parser.y:437
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
		//line .\parser\parser.y:446
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:451
		{
			yyVAL.expr = &ast.NewExpr{Type: yyDollar[3].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:456
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:461
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:466
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:471
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:476
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:481
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:486
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:491
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:496
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:501
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:506
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:511
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:516
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:521
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:526
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:531
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:536
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:541
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:546
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:551
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:556
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:561
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:566
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:571
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:576
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:581
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:586
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 103:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:591
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 104:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line .\parser\parser.y:596
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 105:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:601
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 106:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:606
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 107:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:611
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 108:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line .\parser\parser.y:616
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 109:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:621
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 110:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:626
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 111:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:631
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 112:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line .\parser\parser.y:636
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 113:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line .\parser\parser.y:641
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 114:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser\parser.y:646
		{
			yyVAL.expr = &ast.MakeExpr{Type: yyDollar[3].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 115:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser\parser.y:651
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].typ.Name, SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 116:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line .\parser\parser.y:656
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].typ.Name, SizeExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 117:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line .\parser\parser.y:661
		{
			yyVAL.expr = &ast.MakeArrayExpr{Type: yyDollar[4].typ.Name, LenExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 118:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line .\parser\parser.y:666
		{
			yyVAL.expr = &ast.MakeArrayExpr{Type: yyDollar[4].typ.Name, LenExpr: yyDollar[6].expr, CapExpr: yyDollar[8].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser\parser.y:671
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:676
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:687
		{
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser\parser.y:690
		{
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:695
		{
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser\parser.y:698
		{
		}
	}
	goto yystack /* stack new state and value */
}
