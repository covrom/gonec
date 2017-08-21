//line ./parser/parser.y:1
package parser

import __yyfmt__ "fmt"

//line ./parser/parser.y:3
import (
	"github.com/covrom/gonec/ast"
)

//line ./parser/parser.y:27
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
const TYPECAST = 57400
const UNARY = 57401

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

//line ./parser/parser.y:742

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 3,
	-2, 129,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	62, 47,
	-2, 1,
	-1, 10,
	62, 48,
	-2, 22,
	-1, 20,
	29, 3,
	-2, 129,
	-1, 48,
	62, 47,
	-2, 130,
	-1, 99,
	1, 56,
	8, 56,
	14, 56,
	29, 56,
	45, 56,
	46, 56,
	54, 56,
	55, 56,
	59, 56,
	61, 56,
	62, 56,
	71, 56,
	72, 56,
	77, 56,
	80, 56,
	82, 56,
	83, 56,
	-2, 51,
	-1, 101,
	1, 58,
	8, 58,
	14, 58,
	29, 58,
	45, 58,
	46, 58,
	54, 58,
	55, 58,
	59, 58,
	61, 58,
	62, 58,
	71, 58,
	72, 58,
	77, 58,
	80, 58,
	82, 58,
	83, 58,
	-2, 51,
	-1, 134,
	17, 0,
	18, 0,
	-2, 85,
	-1, 135,
	17, 0,
	18, 0,
	-2, 86,
	-1, 155,
	62, 48,
	-2, 42,
	-1, 157,
	72, 3,
	-2, 129,
	-1, 160,
	72, 3,
	-2, 129,
	-1, 161,
	72, 3,
	-2, 129,
	-1, 188,
	14, 3,
	55, 3,
	72, 3,
	-2, 129,
	-1, 237,
	62, 49,
	-2, 43,
	-1, 238,
	1, 44,
	14, 44,
	29, 44,
	45, 44,
	46, 44,
	55, 44,
	59, 44,
	62, 50,
	72, 44,
	82, 44,
	83, 44,
	-2, 51,
	-1, 246,
	1, 50,
	8, 50,
	14, 50,
	29, 50,
	45, 50,
	46, 50,
	55, 50,
	62, 50,
	72, 50,
	77, 50,
	80, 50,
	82, 50,
	83, 50,
	-2, 51,
	-1, 260,
	72, 3,
	-2, 129,
	-1, 270,
	1, 106,
	8, 106,
	14, 106,
	29, 106,
	45, 106,
	46, 106,
	54, 106,
	55, 106,
	59, 106,
	61, 106,
	62, 106,
	71, 106,
	72, 106,
	77, 106,
	80, 106,
	82, 106,
	83, 106,
	-2, 104,
	-1, 272,
	1, 110,
	8, 110,
	14, 110,
	29, 110,
	45, 110,
	46, 110,
	54, 110,
	55, 110,
	59, 110,
	61, 110,
	62, 110,
	71, 110,
	72, 110,
	77, 110,
	80, 110,
	82, 110,
	83, 110,
	-2, 108,
	-1, 279,
	72, 3,
	-2, 129,
	-1, 282,
	45, 3,
	46, 3,
	72, 3,
	-2, 129,
	-1, 286,
	72, 3,
	-2, 129,
	-1, 287,
	72, 3,
	-2, 129,
	-1, 292,
	1, 105,
	8, 105,
	14, 105,
	29, 105,
	45, 105,
	46, 105,
	54, 105,
	55, 105,
	59, 105,
	61, 105,
	62, 105,
	71, 105,
	72, 105,
	77, 105,
	80, 105,
	82, 105,
	83, 105,
	-2, 103,
	-1, 293,
	1, 109,
	8, 109,
	14, 109,
	29, 109,
	45, 109,
	46, 109,
	54, 109,
	55, 109,
	59, 109,
	61, 109,
	62, 109,
	71, 109,
	72, 109,
	77, 109,
	80, 109,
	82, 109,
	83, 109,
	-2, 107,
	-1, 297,
	72, 3,
	-2, 129,
	-1, 301,
	72, 3,
	-2, 129,
	-1, 302,
	45, 3,
	46, 3,
	72, 3,
	-2, 129,
	-1, 308,
	72, 3,
	-2, 129,
	-1, 318,
	14, 3,
	55, 3,
	72, 3,
	-2, 129,
}

const yyPrivate = 57344

const yyLast = 3137

var yyAct = [...]int{

	85, 176, 50, 10, 45, 204, 11, 265, 112, 205,
	224, 1, 163, 222, 6, 7, 86, 94, 95, 84,
	90, 105, 92, 180, 95, 96, 97, 98, 100, 102,
	6, 7, 91, 6, 7, 103, 183, 271, 173, 108,
	269, 111, 115, 210, 182, 118, 182, 120, 227, 10,
	191, 186, 117, 124, 116, 126, 127, 128, 129, 130,
	131, 132, 133, 134, 135, 136, 137, 138, 139, 140,
	141, 142, 143, 144, 145, 109, 293, 146, 147, 148,
	149, 292, 151, 153, 155, 150, 113, 288, 123, 261,
	154, 156, 255, 123, 156, 104, 166, 156, 2, 165,
	241, 297, 47, 262, 156, 320, 272, 171, 218, 270,
	182, 174, 211, 206, 207, 184, 114, 185, 179, 192,
	206, 207, 319, 155, 177, 317, 315, 314, 311, 189,
	305, 267, 251, 250, 247, 106, 107, 157, 122, 156,
	252, 123, 299, 119, 254, 226, 161, 203, 159, 83,
	206, 207, 198, 195, 263, 219, 177, 240, 230, 298,
	199, 89, 221, 8, 216, 215, 115, 172, 214, 197,
	208, 217, 200, 201, 209, 202, 220, 5, 158, 125,
	87, 51, 49, 175, 231, 228, 229, 236, 237, 4,
	291, 277, 164, 48, 296, 242, 17, 245, 3, 248,
	239, 68, 69, 70, 71, 72, 73, 253, 0, 0,
	88, 59, 121, 0, 256, 0, 0, 0, 187, 0,
	82, 0, 190, 0, 0, 0, 49, 268, 0, 0,
	0, 0, 0, 0, 274, 0, 275, 56, 57, 58,
	0, 0, 0, 53, 0, 0, 78, 0, 80, 81,
	280, 76, 0, 0, 0, 196, 0, 0, 0, 0,
	284, 164, 0, 0, 0, 245, 0, 0, 290, 0,
	0, 0, 285, 223, 225, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 0, 0, 59, 0, 0,
	0, 300, 0, 0, 303, 0, 82, 0, 306, 307,
	310, 0, 0, 0, 0, 0, 0, 0, 0, 309,
	0, 0, 0, 312, 313, 0, 0, 260, 0, 53,
	316, 264, 78, 266, 80, 81, 0, 76, 0, 0,
	321, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 22, 23, 29, 0, 0, 35, 14,
	9, 15, 46, 282, 18, 0, 0, 0, 0, 0,
	286, 287, 39, 30, 31, 32, 16, 20, 0, 0,
	0, 0, 0, 0, 0, 0, 12, 13, 0, 0,
	302, 0, 0, 21, 0, 0, 40, 308, 41, 44,
	42, 33, 0, 0, 0, 19, 34, 43, 0, 0,
	0, 0, 0, 0, 0, 24, 28, 0, 0, 0,
	37, 0, 0, 25, 26, 27, 0, 38, 36, 0,
	0, 6, 7, 62, 63, 65, 67, 77, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 82, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 235, 64,
	66, 54, 55, 56, 57, 58, 0, 0, 0, 53,
	0, 0, 78, 234, 80, 81, 0, 76, 62, 63,
	65, 67, 77, 79, 0, 0, 0, 0, 0, 0,
	0, 0, 68, 69, 70, 71, 72, 73, 0, 0,
	74, 75, 59, 60, 61, 0, 0, 0, 0, 0,
	0, 82, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 233, 64, 66, 54, 55, 56, 57,
	58, 0, 0, 0, 53, 0, 0, 78, 232, 80,
	81, 0, 76, 62, 63, 65, 67, 77, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 82, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 213, 0, 64,
	66, 54, 55, 56, 57, 58, 0, 0, 0, 53,
	0, 0, 78, 0, 80, 81, 212, 76, 62, 63,
	65, 67, 77, 79, 0, 0, 0, 0, 0, 0,
	0, 0, 68, 69, 70, 71, 72, 73, 0, 0,
	74, 75, 59, 60, 61, 0, 0, 0, 0, 0,
	0, 82, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 194, 0, 64, 66, 54, 55, 56, 57,
	58, 0, 0, 0, 53, 0, 0, 78, 0, 80,
	81, 193, 76, 22, 23, 29, 0, 0, 35, 14,
	9, 15, 46, 0, 18, 0, 0, 0, 0, 0,
	0, 0, 39, 30, 31, 32, 16, 20, 0, 0,
	0, 0, 0, 0, 0, 0, 12, 13, 0, 0,
	0, 0, 0, 21, 0, 0, 40, 0, 41, 44,
	42, 33, 0, 0, 0, 19, 34, 43, 0, 0,
	0, 0, 0, 0, 0, 24, 28, 0, 0, 0,
	37, 0, 0, 25, 26, 27, 0, 38, 36, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 66, 54, 55, 56,
	57, 58, 0, 318, 0, 53, 0, 0, 78, 0,
	80, 81, 0, 76, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 0, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 60,
	61, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	53, 0, 0, 78, 304, 80, 81, 0, 76, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 66, 54, 55, 56,
	57, 58, 0, 301, 0, 53, 0, 0, 78, 0,
	80, 81, 0, 76, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 0, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 60,
	61, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	53, 0, 0, 78, 295, 80, 81, 0, 76, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 66, 54, 55, 56,
	57, 58, 0, 0, 0, 53, 0, 0, 78, 294,
	80, 81, 0, 76, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 0, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 60,
	61, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	53, 0, 0, 78, 0, 80, 81, 283, 76, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 281, 0, 64, 66, 54, 55, 56,
	57, 58, 0, 0, 0, 53, 0, 0, 78, 0,
	80, 81, 0, 76, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 0, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 60,
	61, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 279, 0,
	53, 0, 0, 78, 0, 80, 81, 0, 76, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 66, 54, 55, 56,
	57, 58, 0, 0, 0, 53, 0, 0, 78, 0,
	80, 81, 278, 76, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 0, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 60,
	61, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	53, 0, 0, 78, 276, 80, 81, 0, 76, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 66, 54, 55, 56,
	57, 58, 0, 0, 0, 53, 0, 0, 78, 273,
	80, 81, 0, 76, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 0, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 60,
	61, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 259,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	53, 0, 0, 78, 0, 80, 81, 0, 76, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 66, 54, 55, 56,
	57, 58, 0, 0, 0, 53, 0, 0, 78, 0,
	80, 81, 258, 76, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 0, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 60,
	61, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 249, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	53, 0, 0, 78, 0, 80, 81, 0, 76, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 66, 54, 55, 56,
	57, 58, 0, 0, 0, 53, 0, 0, 78, 0,
	80, 81, 244, 76, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 0, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 60,
	61, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 188, 0,
	53, 0, 0, 78, 0, 80, 81, 0, 76, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 66, 54, 55, 56,
	57, 58, 0, 0, 0, 53, 0, 0, 78, 178,
	80, 81, 0, 76, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 0, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 60,
	61, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 170,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	53, 0, 0, 78, 0, 80, 81, 0, 76, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 162, 0, 64, 66, 54, 55, 56,
	57, 58, 0, 0, 0, 53, 0, 0, 78, 0,
	80, 81, 0, 76, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 0, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 60,
	61, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 160, 0,
	53, 0, 0, 78, 0, 80, 81, 0, 76, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 52, 0, 0, 0, 64, 66, 54, 55, 56,
	57, 58, 0, 0, 0, 53, 0, 0, 78, 0,
	80, 81, 0, 76, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 0, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 60,
	61, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	53, 0, 0, 78, 0, 80, 81, 0, 76, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 66, 54, 55, 56,
	57, 58, 0, 0, 0, 53, 0, 0, 78, 0,
	181, 81, 0, 76, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 0, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 60,
	61, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	169, 0, 0, 78, 0, 80, 81, 0, 76, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 66, 54, 55, 56,
	57, 58, 0, 0, 0, 168, 0, 0, 78, 0,
	80, 81, 0, 76, 62, 63, 65, 67, 0, 79,
	0, 0, 0, 0, 0, 0, 0, 0, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 60,
	61, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	53, 0, 0, 78, 0, 80, 81, 0, 76, 62,
	63, 65, 67, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 66, 54, 55, 56,
	57, 58, 65, 67, 0, 53, 0, 0, 78, 0,
	80, 81, 0, 76, 68, 69, 70, 71, 72, 73,
	0, 0, 74, 75, 59, 60, 61, 0, 0, 0,
	0, 0, 0, 82, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 66, 54, 55,
	56, 57, 58, 246, 23, 29, 53, 0, 35, 78,
	0, 80, 81, 0, 76, 0, 0, 0, 0, 0,
	0, 0, 39, 30, 31, 32, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 22,
	23, 29, 0, 0, 35, 0, 40, 0, 41, 44,
	42, 33, 0, 0, 0, 0, 34, 43, 39, 30,
	31, 32, 0, 0, 0, 24, 28, 0, 0, 0,
	37, 0, 0, 25, 26, 27, 0, 38, 36, 289,
	0, 0, 40, 0, 41, 44, 42, 33, 0, 0,
	0, 0, 34, 43, 0, 0, 0, 0, 22, 23,
	29, 24, 28, 35, 0, 0, 37, 0, 0, 25,
	26, 27, 0, 38, 36, 257, 0, 39, 30, 31,
	32, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 22, 23, 29, 0, 0, 35,
	0, 40, 0, 41, 44, 42, 33, 0, 0, 0,
	0, 34, 43, 39, 30, 31, 32, 0, 0, 0,
	24, 28, 0, 0, 0, 37, 0, 0, 25, 26,
	27, 0, 38, 36, 243, 0, 0, 40, 0, 41,
	44, 42, 33, 0, 0, 0, 0, 34, 43, 0,
	0, 167, 0, 22, 23, 29, 24, 28, 35, 0,
	0, 37, 0, 0, 25, 26, 27, 0, 38, 36,
	0, 0, 39, 30, 31, 32, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 22, 23, 29, 0, 40, 35, 41, 44,
	42, 33, 0, 0, 0, 0, 34, 43, 0, 0,
	152, 39, 30, 31, 32, 24, 28, 0, 0, 0,
	37, 0, 0, 25, 26, 27, 0, 38, 36, 0,
	0, 22, 23, 29, 0, 40, 35, 41, 44, 42,
	33, 0, 0, 0, 0, 34, 43, 0, 0, 93,
	39, 30, 31, 32, 24, 28, 0, 0, 0, 37,
	0, 0, 25, 26, 27, 0, 38, 36, 246, 23,
	29, 0, 0, 35, 40, 0, 41, 44, 42, 33,
	0, 0, 0, 0, 34, 43, 0, 39, 30, 31,
	32, 0, 0, 24, 28, 0, 0, 0, 37, 0,
	0, 25, 26, 27, 0, 38, 36, 238, 23, 29,
	0, 40, 35, 41, 44, 42, 33, 0, 0, 0,
	0, 34, 43, 0, 0, 0, 39, 30, 31, 32,
	24, 28, 0, 0, 0, 37, 0, 0, 25, 26,
	27, 0, 38, 36, 110, 23, 29, 0, 0, 35,
	40, 0, 41, 44, 42, 33, 0, 0, 0, 0,
	34, 43, 0, 39, 30, 31, 32, 0, 0, 24,
	28, 0, 0, 0, 37, 0, 0, 25, 26, 27,
	0, 38, 36, 101, 23, 29, 0, 40, 35, 41,
	44, 42, 33, 0, 0, 0, 0, 34, 43, 0,
	0, 0, 39, 30, 31, 32, 24, 28, 0, 0,
	0, 37, 0, 0, 25, 26, 27, 0, 38, 36,
	99, 23, 29, 0, 0, 35, 40, 0, 41, 44,
	42, 33, 0, 0, 0, 0, 34, 43, 0, 39,
	30, 31, 32, 0, 0, 24, 28, 0, 0, 0,
	37, 0, 0, 25, 26, 27, 0, 38, 36, 0,
	0, 0, 0, 40, 0, 41, 44, 42, 33, 0,
	0, 0, 0, 34, 43, 0, 0, 0, 0, 0,
	0, 0, 24, 28, 0, 0, 0, 37, 0, 0,
	25, 26, 27, 0, 38, 36, 68, 69, 70, 71,
	72, 73, 0, 0, 74, 75, 59, 0, 0, 0,
	0, 0, 0, 0, 0, 82, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	54, 55, 56, 57, 58, 0, 0, 0, 53, 0,
	0, 78, 0, 80, 81, 0, 76,
}
var yyPact = [...]int{

	-68, -1000, 679, -68, -68, -1000, -1000, -1000, -1000, 177,
	2042, 90, -1000, -1000, 2817, 2817, 176, -1000, 157, 2817,
	-68, 2778, -61, -1000, 2817, 2817, 2817, 3006, 2969, -1000,
	-1000, -1000, -1000, -1000, 2817, 17, -68, -68, 2817, -3,
	2930, 38, -24, 162, 2817, 81, 2817, -1000, 339, -1000,
	79, -1000, 2817, 175, 2817, 2817, 2817, 2817, 2817, 2817,
	2817, 2817, 2817, 2817, 2817, 2817, 2817, 2817, 2817, 2817,
	2817, 2817, 2817, 2817, -1000, -1000, 2817, 2817, 2817, 2817,
	2817, 2739, 2817, 2817, 77, 2107, 2107, 66, 174, 89,
	1977, 117, 1912, -68, 2817, 2680, 246, 246, 246, -61,
	2302, -61, 2237, 1847, 163, -40, 2817, 150, 1782, 162,
	-55, 2172, 37, -42, 2817, -1000, 2817, -27, 2107, -68,
	1717, -1000, 2817, -68, 2107, -1000, 170, 170, 246, 246,
	246, 2107, 3055, 3055, 2483, 2483, 3055, 3055, 3055, 3055,
	2107, 2107, 2107, 2107, 2107, 2107, 2107, 2367, 2107, 2432,
	42, 601, 2817, 2107, -1000, 2107, -68, -68, 136, 2817,
	-68, -68, -68, 75, 105, 35, 536, 2817, 161, 160,
	2817, 31, 147, 158, -49, -52, -1000, 84, -1000, -29,
	2817, 2817, 154, 2817, 471, 406, 2817, 2893, -68, -1000,
	153, 23, -1000, -1000, 2644, 1652, 2854, 62, 2817, 1587,
	61, 60, 68, -1000, -1000, -1000, 2817, 83, -1000, -1000,
	15, -1000, -1000, 2585, 1522, -1000, -1000, 1457, -68, 12,
	26, 146, -68, -73, -68, 59, 2817, -1000, 32, 29,
	-1000, 1392, -1000, 2817, -1000, 2817, 1327, 2107, -61, -1000,
	-1000, -1000, 1262, -1000, -1000, 2107, -61, -1000, 1197, 2817,
	-1000, -1000, -1000, 1132, -68, -1000, 1067, -1000, -1000, 2817,
	-68, -68, -68, 10, 2549, -1000, 118, -1000, 2107, 4,
	-1000, -1, -1000, -1000, 1002, 937, -1000, 87, -1000, -68,
	872, -68, -68, -1000, 807, 58, -68, -68, -68, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -68, -1000, 2817,
	56, -68, -68, -1000, -1000, -1000, 55, 54, -68, 53,
	742, -1000, 50, -1000, -1000, -1000, 33, -1000, -68, -1000,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 11, 198, 163, 196, 9, 5, 12, 194, 191,
	8, 0, 4, 6, 1, 183, 2, 98, 189, 177,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 9, 9, 8, 4, 4, 7, 7,
	7, 7, 7, 6, 5, 14, 15, 15, 15, 16,
	16, 16, 13, 13, 13, 10, 10, 12, 12, 12,
	12, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 17,
	17, 18, 18, 19, 19,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 2, 3, 4, 2, 3, 3,
	1, 1, 2, 2, 5, 1, 8, 9, 5, 5,
	5, 4, 1, 0, 2, 4, 8, 6, 0, 2,
	2, 2, 2, 5, 4, 3, 0, 1, 4, 0,
	1, 4, 1, 4, 4, 1, 3, 0, 1, 4,
	4, 1, 1, 2, 2, 2, 2, 4, 2, 4,
	1, 1, 1, 1, 1, 7, 3, 7, 8, 8,
	9, 5, 6, 5, 6, 3, 4, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 3,
	3, 3, 3, 5, 4, 6, 5, 5, 4, 6,
	5, 4, 4, 6, 5, 5, 6, 5, 5, 2,
	2, 5, 4, 6, 5, 4, 6, 3, 2, 0,
	1, 1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -17, -2, -18, -19, 82, 83, -3, 11,
	-11, -13, 37, 38, 10, 12, 27, -4, 15, 56,
	28, 44, 4, 5, 66, 74, 75, 76, 67, 6,
	24, 25, 26, 52, 57, 9, 79, 71, 78, 23,
	47, 49, 51, 58, 50, -12, 13, -17, -18, -19,
	-16, 4, 59, 73, 65, 66, 67, 68, 69, 41,
	42, 43, 17, 18, 63, 19, 64, 20, 31, 32,
	33, 34, 35, 36, 39, 40, 81, 21, 76, 22,
	78, 79, 50, 59, -12, -11, -11, 4, 53, 4,
	-11, -1, -11, 61, 78, 79, -11, -11, -11, 4,
	-11, 4, -11, -11, 78, 4, -17, -17, -11, 78,
	4, -11, -10, 48, 78, 4, 78, -10, -11, 62,
	-11, -3, 59, 62, -11, 4, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-12, -11, 61, -11, -13, -11, 62, 71, 4, 59,
	71, 29, 61, -7, -17, -12, -11, 61, 73, 73,
	62, -16, 4, 78, -12, -15, -14, 6, 77, -10,
	78, 78, 73, 78, -11, -11, 78, -17, 71, -13,
	-17, 8, 77, 80, 61, -11, -17, -1, 16, -11,
	-1, -1, -7, 72, -6, -5, 45, 46, -6, -5,
	8, 77, 80, 61, -11, 4, 4, -11, 77, 8,
	-16, 4, 62, -17, 62, -17, 61, 77, -12, -12,
	4, -11, 77, 62, 77, 62, -11, -11, 4, -1,
	4, 77, -11, 80, 80, -11, 4, 72, -11, 54,
	72, 72, 72, -11, 61, 77, -11, 80, 80, 62,
	-17, 77, 77, 8, -17, 80, -17, 72, -11, 8,
	77, 8, 77, 77, -11, -11, 77, -9, 80, 71,
	-11, 61, -17, 80, -11, -1, -17, -17, 77, 80,
	-14, 72, 77, 77, 77, 77, -8, 14, 72, 55,
	-1, 71, -17, -1, 77, 72, -1, -1, -17, -1,
	-11, 72, -1, -1, 72, 72, -1, 72, 71, 72,
	72, -1,
}
var yyDef = [...]int{

	-2, -2, -2, 129, 130, 131, 133, 134, 4, 39,
	-2, 0, 10, 11, 47, 0, 0, 15, 0, 0,
	-2, 0, 51, 52, 0, 0, 0, 0, 0, 60,
	61, 62, 63, 64, 0, 0, 129, 129, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 2, -2, 132,
	7, 40, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 97, 98, 0, 0, 0, 0,
	47, 0, 0, 47, 12, 48, 13, 0, 0, 0,
	0, 0, 0, 28, 47, 0, 53, 54, 55, -2,
	0, -2, 0, 0, 39, 0, 47, 36, 0, 0,
	51, 0, 119, 120, 0, 45, 0, 0, 128, 129,
	0, 5, 47, 129, 8, 66, 77, 78, 79, 80,
	81, 82, 83, 84, -2, -2, 87, 88, 89, 90,
	91, 92, 93, 94, 95, 96, 99, 100, 101, 102,
	0, 0, 0, 127, 9, -2, 129, -2, 0, 0,
	-2, -2, 28, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 40, 39, 129, 129, 37, 0, 75, 0,
	47, 47, 0, 0, 0, 0, 0, 0, -2, 6,
	0, 0, 108, 112, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 21, 31, 32, 0, 0, 29, 30,
	0, 104, 111, 0, 0, 57, 59, 0, 129, 0,
	0, 40, 129, 0, 129, 0, 0, 76, 0, 0,
	46, 0, 125, 0, 122, 0, 0, -2, -2, 23,
	41, 107, 0, 117, 118, 49, -2, 14, 0, 0,
	18, 19, 20, 0, 129, 103, 0, 114, 115, 0,
	-2, 129, 129, 0, 0, 71, 0, 73, 35, 0,
	-2, 0, -2, 121, 0, 0, 124, 0, 116, -2,
	0, 129, -2, 113, 0, 0, -2, -2, 129, 72,
	38, 74, -2, -2, 126, 123, 24, -2, 27, 0,
	0, -2, -2, 34, 65, 67, 0, 0, -2, 0,
	0, 16, 0, 33, 68, 69, 0, 26, -2, 17,
	70, 25,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	83, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 74, 3, 3, 3, 69, 76, 3,
	78, 77, 67, 65, 62, 66, 73, 68, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 61, 82,
	64, 59, 63, 60, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 79, 3, 80, 75, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 71, 81, 72,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 70,
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
		//line ./parser/parser.y:67
		{
			yyVAL.compstmt = nil
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:71
		{
			yyVAL.compstmt = yyDollar[1].stmts
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.y:76
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:83
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:90
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
		//line ./parser/parser.y:101
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].expr_many}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:106
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: nil}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:111
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "=", Rhss: []ast.Expr{yyDollar[3].expr}}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:115
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:119
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:124
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:129
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:134
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 14:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:139
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:144
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 16:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./parser/parser.y:149
		{
			yyVAL.stmt = &ast.ForStmt{Var: ast.UniqueNames.Set(yyDollar[3].tok.Lit), Value: yyDollar[5].expr, Stmts: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./parser/parser.y:154
		{
			yyVAL.stmt = &ast.NumForStmt{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), Expr1: yyDollar[4].expr, Expr2: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:159
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:164
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[2].compstmt, Catch: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:169
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:174
		{
			yyVAL.stmt = &ast.SelectStmt{Cases: yyDollar[3].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:179
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.y:185
		{
			yyVAL.stmt_elsifs = []ast.Stmt{}
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:189
		{
			yyVAL.stmt_elsifs = append(yyDollar[1].stmt_elsifs, yyDollar[2].stmt_elsif)
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:195
		{
			yyVAL.stmt_elsif = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt}
		}
	case 26:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./parser/parser.y:201
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, ElseIf: yyDollar[5].stmt_elsifs, Else: yyDollar[7].compstmt}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 27:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:206
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, ElseIf: yyDollar[5].stmt_elsifs, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.y:212
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:216
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_case}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:220
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_default}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:224
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_case)
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:228
		{
			for _, stmt := range yyDollar[1].stmt_cases {
				if _, ok := stmt.(*ast.DefaultStmt); ok {
					yylex.Error("multiple default statement")
				}
			}
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_default)
		}
	case 33:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:239
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[5].compstmt}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:245
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyDollar[4].compstmt}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:251
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.y:256
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:260
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:264
		{
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[4].expr_pair)
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.y:269
		{
			yyVAL.expr_idents = []int{}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:273
		{
			yyVAL.expr_idents = []int{ast.UniqueNames.Set(yyDollar[1].tok.Lit)}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:277
		{
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, ast.UniqueNames.Set(yyDollar[4].tok.Lit))
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:283
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 43:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:287
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:291
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[4].tok.Lit)})
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:296
		{
			yyVAL.typ = ast.Type{Name: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:300
		{
			yyVAL.typ = ast.Type{Name: ast.UniqueNames.Set(ast.UniqueNames.Get(yyDollar[1].typ.Name) + "." + yyDollar[3].tok.Lit)}
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.y:305
		{
			yyVAL.exprs = nil
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:309
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 49:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:313
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:317
		{
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[4].tok.Lit)})
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:323
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:328
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:333
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:338
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:343
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:348
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[2].tok.Lit)}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:353
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: ast.UniqueNames.Set(yyDollar[4].tok.Lit)}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:358
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[2].tok.Lit)}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 59:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:363
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: ast.UniqueNames.Set(yyDollar[4].tok.Lit)}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:368
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:373
		{
			yyVAL.expr = &ast.ConstExpr{Value: "истина"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:378
		{
			yyVAL.expr = &ast.ConstExpr{Value: "ложь"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:383
		{
			yyVAL.expr = &ast.ConstExpr{Value: "неопределено"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:388
		{
			yyVAL.expr = &ast.ConstExpr{Value: "null"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 65:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./parser/parser.y:393
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[2].expr, Lhs: yyDollar[4].expr, Rhs: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:398
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: ast.UniqueNames.Set(yyDollar[3].tok.Lit)}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 67:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./parser/parser.y:403
		{
			yyVAL.expr = &ast.FuncExpr{Name: ast.UniqueNames.Set("<анонимная функция>"), Args: yyDollar[3].expr_idents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 68:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./parser/parser.y:408
		{
			yyVAL.expr = &ast.FuncExpr{Name: ast.UniqueNames.Set("<анонимная функция>"), Args: []int{ast.UniqueNames.Set(yyDollar[3].tok.Lit)}, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./parser/parser.y:413
		{
			yyVAL.expr = &ast.FuncExpr{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), Args: yyDollar[4].expr_idents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 70:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./parser/parser.y:418
		{
			yyVAL.expr = &ast.FuncExpr{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), Args: []int{ast.UniqueNames.Set(yyDollar[4].tok.Lit)}, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 71:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:423
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 72:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:428
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 73:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:433
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
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:442
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
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:451
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:456
		{
			yyVAL.expr = &ast.NewExpr{Type: yyDollar[3].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:461
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:466
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:471
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:476
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:481
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:486
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:491
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:496
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:501
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:506
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:511
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:516
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:521
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:526
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:531
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:536
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:541
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:546
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:551
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:556
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:561
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:566
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:571
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:576
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:581
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:586
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 103:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:591
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[1].tok.Lit), SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 104:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:596
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[1].tok.Lit), SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 105:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:601
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 106:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:606
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 107:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:611
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 108:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:616
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 109:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:621
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 110:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:626
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 111:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:631
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:636
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 113:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:641
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 114:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:646
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Begin: yyDollar[3].expr, End: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 115:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:651
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Begin: nil, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 116:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:656
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 117:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:661
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: nil}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 118:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:666
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: nil, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 119:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:671
		{
			yyVAL.expr = &ast.MakeExpr{Type: yyDollar[2].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:676
		{
			yyVAL.expr = &ast.MakeChanExpr{SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 121:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:681
		{
			yyVAL.expr = &ast.MakeChanExpr{SizeExpr: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 122:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:686
		{
			yyVAL.expr = &ast.MakeArrayExpr{LenExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 123:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:691
		{
			yyVAL.expr = &ast.MakeArrayExpr{LenExpr: yyDollar[3].expr, CapExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 124:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:696
		{
			yyVAL.expr = &ast.TypeCast{Type: yyDollar[2].typ.Name, CastExpr: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 125:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:701
		{
			yyVAL.expr = &ast.MakeExpr{TypeExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 126:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:706
		{
			yyVAL.expr = &ast.TypeCast{TypeExpr: yyDollar[3].expr, CastExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:711
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:716
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:727
		{
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:730
		{
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:735
		{
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:738
		{
		}
	}
	goto yystack /* stack new state and value */
}
