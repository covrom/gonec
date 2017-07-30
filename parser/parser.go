//line parser.y:1
package parser

import __yyfmt__ "fmt"

//line parser.y:3
import (
	"github.com/covrom/gonec/ast"
)

//line parser.y:28
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
	expr_lets    ast.Expr
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
	"';'",
	"'.'",
	"'!'",
	"'^'",
	"'&'",
	"'('",
	"')'",
	"'['",
	"']'",
	"'|'",
	"'\\n'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:714

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 3,
	-2, 123,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	60, 48,
	-2, 1,
	-1, 10,
	60, 49,
	-2, 22,
	-1, 20,
	29, 3,
	-2, 123,
	-1, 45,
	60, 48,
	-2, 124,
	-1, 89,
	60, 49,
	-2, 43,
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
	71, 57,
	77, 57,
	79, 57,
	81, 57,
	-2, 52,
	-1, 100,
	1, 59,
	8, 59,
	14, 59,
	29, 59,
	45, 59,
	46, 59,
	54, 59,
	55, 59,
	57, 59,
	59, 59,
	60, 59,
	69, 59,
	70, 59,
	71, 59,
	77, 59,
	79, 59,
	81, 59,
	-2, 52,
	-1, 128,
	17, 0,
	18, 0,
	-2, 86,
	-1, 129,
	17, 0,
	18, 0,
	-2, 87,
	-1, 148,
	60, 49,
	-2, 43,
	-1, 150,
	70, 3,
	-2, 123,
	-1, 154,
	70, 3,
	-2, 123,
	-1, 156,
	70, 3,
	-2, 123,
	-1, 178,
	14, 3,
	55, 3,
	70, 3,
	-2, 123,
	-1, 218,
	60, 50,
	-2, 44,
	-1, 219,
	1, 45,
	14, 45,
	29, 45,
	45, 45,
	46, 45,
	55, 45,
	57, 45,
	60, 51,
	70, 45,
	71, 45,
	81, 45,
	-2, 52,
	-1, 226,
	1, 51,
	8, 51,
	14, 51,
	29, 51,
	45, 51,
	46, 51,
	55, 51,
	60, 51,
	70, 51,
	71, 51,
	77, 51,
	79, 51,
	81, 51,
	-2, 52,
	-1, 242,
	70, 3,
	-2, 123,
	-1, 253,
	1, 107,
	8, 107,
	14, 107,
	29, 107,
	45, 107,
	46, 107,
	54, 107,
	55, 107,
	57, 107,
	59, 107,
	60, 107,
	69, 107,
	70, 107,
	71, 107,
	77, 107,
	79, 107,
	81, 107,
	-2, 105,
	-1, 255,
	1, 111,
	8, 111,
	14, 111,
	29, 111,
	45, 111,
	46, 111,
	54, 111,
	55, 111,
	57, 111,
	59, 111,
	60, 111,
	69, 111,
	70, 111,
	71, 111,
	77, 111,
	79, 111,
	81, 111,
	-2, 109,
	-1, 261,
	70, 3,
	-2, 123,
	-1, 268,
	70, 3,
	-2, 123,
	-1, 269,
	70, 3,
	-2, 123,
	-1, 274,
	1, 106,
	8, 106,
	14, 106,
	29, 106,
	45, 106,
	46, 106,
	54, 106,
	55, 106,
	57, 106,
	59, 106,
	60, 106,
	69, 106,
	70, 106,
	71, 106,
	77, 106,
	79, 106,
	81, 106,
	-2, 104,
	-1, 275,
	1, 110,
	8, 110,
	14, 110,
	29, 110,
	45, 110,
	46, 110,
	54, 110,
	55, 110,
	57, 110,
	59, 110,
	60, 110,
	69, 110,
	70, 110,
	71, 110,
	77, 110,
	79, 110,
	81, 110,
	-2, 108,
	-1, 279,
	70, 3,
	-2, 123,
	-1, 283,
	70, 3,
	-2, 123,
	-1, 284,
	70, 3,
	-2, 123,
	-1, 286,
	45, 3,
	46, 3,
	70, 3,
	-2, 123,
	-1, 290,
	70, 3,
	-2, 123,
	-1, 299,
	45, 3,
	46, 3,
	70, 3,
	-2, 123,
	-1, 306,
	14, 3,
	55, 3,
	70, 3,
	-2, 123,
}

const yyPrivate = 57344

const yyLast = 2414

var yyAct = [...]int{

	83, 167, 47, 10, 170, 234, 235, 6, 247, 275,
	208, 1, 274, 93, 11, 94, 84, 7, 42, 116,
	89, 6, 92, 270, 206, 95, 96, 97, 99, 101,
	243, 7, 91, 82, 90, 6, 244, 254, 106, 240,
	109, 252, 111, 223, 113, 7, 10, 172, 2, 94,
	117, 118, 44, 120, 121, 122, 123, 124, 125, 126,
	127, 128, 129, 130, 131, 132, 133, 134, 135, 136,
	137, 138, 139, 196, 257, 140, 141, 142, 143, 103,
	145, 146, 148, 211, 104, 105, 211, 211, 215, 149,
	153, 256, 212, 149, 116, 159, 147, 144, 164, 66,
	67, 68, 69, 70, 71, 162, 255, 72, 73, 57,
	253, 202, 158, 110, 107, 174, 148, 258, 80, 182,
	310, 168, 308, 165, 307, 149, 305, 150, 279, 211,
	179, 52, 53, 54, 55, 56, 302, 301, 236, 237,
	51, 296, 197, 76, 78, 265, 79, 287, 74, 249,
	232, 102, 231, 189, 190, 227, 148, 115, 149, 112,
	116, 177, 187, 233, 210, 180, 191, 204, 193, 281,
	192, 149, 87, 155, 226, 23, 29, 152, 218, 34,
	216, 217, 222, 81, 280, 273, 224, 225, 183, 228,
	220, 213, 214, 38, 30, 31, 32, 188, 186, 245,
	241, 238, 239, 236, 237, 156, 195, 203, 168, 171,
	251, 250, 221, 171, 207, 209, 8, 39, 205, 40,
	41, 86, 33, 66, 67, 68, 69, 70, 71, 201,
	262, 263, 200, 57, 24, 28, 163, 264, 151, 36,
	119, 85, 80, 25, 26, 27, 37, 225, 35, 271,
	272, 242, 48, 175, 267, 246, 176, 248, 276, 277,
	5, 166, 114, 88, 51, 46, 259, 76, 78, 278,
	79, 4, 74, 282, 194, 45, 17, 3, 0, 0,
	288, 289, 295, 0, 0, 0, 0, 0, 0, 0,
	0, 294, 268, 269, 304, 297, 298, 0, 300, 0,
	0, 0, 303, 0, 0, 0, 46, 0, 0, 0,
	0, 309, 0, 0, 286, 0, 0, 0, 312, 290,
	22, 23, 29, 0, 0, 34, 14, 9, 15, 43,
	0, 18, 0, 0, 299, 0, 0, 0, 0, 38,
	30, 31, 32, 16, 20, 0, 0, 0, 0, 0,
	0, 0, 0, 12, 13, 0, 0, 0, 0, 0,
	21, 0, 0, 39, 0, 40, 41, 0, 33, 0,
	0, 0, 19, 0, 0, 0, 0, 0, 0, 0,
	24, 28, 0, 0, 0, 36, 0, 6, 0, 25,
	26, 27, 37, 0, 35, 0, 0, 7, 60, 61,
	63, 65, 75, 77, 0, 0, 0, 0, 0, 0,
	0, 0, 66, 67, 68, 69, 70, 71, 0, 0,
	72, 73, 57, 58, 59, 0, 0, 0, 0, 0,
	0, 80, 0, 0, 0, 0, 0, 0, 0, 50,
	0, 293, 62, 64, 52, 53, 54, 55, 56, 0,
	0, 0, 0, 51, 0, 0, 76, 78, 292, 79,
	0, 74, 60, 61, 63, 65, 75, 77, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 67, 68, 69,
	70, 71, 0, 0, 72, 73, 57, 58, 59, 0,
	0, 0, 0, 0, 0, 80, 0, 0, 0, 0,
	0, 0, 0, 50, 199, 0, 62, 64, 52, 53,
	54, 55, 56, 0, 0, 0, 0, 51, 0, 0,
	76, 78, 0, 79, 198, 74, 60, 61, 63, 65,
	75, 77, 0, 0, 0, 0, 0, 0, 0, 0,
	66, 67, 68, 69, 70, 71, 0, 0, 72, 73,
	57, 58, 59, 0, 0, 0, 0, 0, 0, 80,
	0, 0, 0, 0, 0, 0, 0, 50, 185, 0,
	62, 64, 52, 53, 54, 55, 56, 0, 0, 0,
	0, 51, 0, 0, 76, 78, 0, 79, 184, 74,
	60, 61, 63, 65, 75, 77, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 67, 68, 69, 70, 71,
	0, 0, 72, 73, 57, 58, 59, 0, 0, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 0, 0,
	0, 50, 0, 0, 62, 64, 52, 53, 54, 55,
	56, 0, 0, 0, 0, 51, 0, 0, 76, 78,
	311, 79, 0, 74, 60, 61, 63, 65, 75, 77,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 67,
	68, 69, 70, 71, 0, 0, 72, 73, 57, 58,
	59, 0, 0, 0, 0, 0, 0, 80, 0, 0,
	0, 0, 0, 0, 0, 50, 0, 0, 62, 64,
	52, 53, 54, 55, 56, 0, 306, 0, 0, 51,
	0, 0, 76, 78, 0, 79, 0, 74, 60, 61,
	63, 65, 75, 77, 0, 0, 0, 0, 0, 0,
	0, 0, 66, 67, 68, 69, 70, 71, 0, 0,
	72, 73, 57, 58, 59, 0, 0, 0, 0, 0,
	0, 80, 0, 0, 0, 0, 0, 0, 0, 50,
	0, 0, 62, 64, 52, 53, 54, 55, 56, 0,
	0, 0, 0, 51, 0, 0, 76, 78, 291, 79,
	0, 74, 60, 61, 63, 65, 75, 77, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 67, 68, 69,
	70, 71, 0, 0, 72, 73, 57, 58, 59, 0,
	0, 0, 0, 0, 0, 80, 0, 0, 0, 0,
	0, 0, 0, 50, 285, 0, 62, 64, 52, 53,
	54, 55, 56, 0, 0, 0, 0, 51, 0, 0,
	76, 78, 0, 79, 0, 74, 60, 61, 63, 65,
	75, 77, 0, 0, 0, 0, 0, 0, 0, 0,
	66, 67, 68, 69, 70, 71, 0, 0, 72, 73,
	57, 58, 59, 0, 0, 0, 0, 0, 0, 80,
	0, 0, 0, 0, 0, 0, 0, 50, 0, 0,
	62, 64, 52, 53, 54, 55, 56, 0, 284, 0,
	0, 51, 0, 0, 76, 78, 0, 79, 0, 74,
	60, 61, 63, 65, 75, 77, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 67, 68, 69, 70, 71,
	0, 0, 72, 73, 57, 58, 59, 0, 0, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 0, 0,
	0, 50, 0, 0, 62, 64, 52, 53, 54, 55,
	56, 0, 283, 0, 0, 51, 0, 0, 76, 78,
	0, 79, 0, 74, 60, 61, 63, 65, 75, 77,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 67,
	68, 69, 70, 71, 0, 0, 72, 73, 57, 58,
	59, 0, 0, 0, 0, 0, 0, 80, 0, 0,
	0, 0, 0, 0, 0, 50, 0, 0, 62, 64,
	52, 53, 54, 55, 56, 0, 0, 0, 0, 51,
	0, 0, 76, 78, 0, 79, 266, 74, 60, 61,
	63, 65, 75, 77, 0, 0, 0, 0, 0, 0,
	0, 0, 66, 67, 68, 69, 70, 71, 0, 0,
	72, 73, 57, 58, 59, 0, 0, 0, 0, 0,
	0, 80, 0, 0, 0, 0, 0, 0, 0, 50,
	0, 0, 62, 64, 52, 53, 54, 55, 56, 0,
	261, 0, 0, 51, 0, 0, 76, 78, 0, 79,
	0, 74, 60, 61, 63, 65, 75, 77, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 67, 68, 69,
	70, 71, 0, 0, 72, 73, 57, 58, 59, 0,
	0, 0, 0, 0, 0, 80, 0, 0, 0, 0,
	0, 0, 0, 50, 0, 0, 62, 64, 52, 53,
	54, 55, 56, 0, 0, 0, 0, 51, 0, 0,
	76, 78, 0, 79, 260, 74, 60, 61, 63, 65,
	75, 77, 0, 0, 0, 0, 0, 0, 0, 0,
	66, 67, 68, 69, 70, 71, 0, 0, 72, 73,
	57, 58, 59, 0, 0, 0, 0, 0, 0, 80,
	0, 0, 0, 0, 0, 0, 0, 50, 0, 0,
	62, 64, 52, 53, 54, 55, 56, 0, 0, 0,
	230, 51, 0, 0, 76, 78, 0, 79, 0, 74,
	60, 61, 63, 65, 75, 77, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 67, 68, 69, 70, 71,
	0, 0, 72, 73, 57, 58, 59, 0, 0, 0,
	0, 0, 0, 80, 0, 0, 0, 229, 0, 0,
	0, 50, 0, 0, 62, 64, 52, 53, 54, 55,
	56, 0, 0, 0, 0, 51, 0, 0, 76, 78,
	0, 79, 0, 74, 60, 61, 63, 65, 75, 77,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 67,
	68, 69, 70, 71, 0, 0, 72, 73, 57, 58,
	59, 0, 0, 0, 0, 0, 0, 80, 0, 0,
	0, 0, 0, 0, 0, 50, 181, 0, 62, 64,
	52, 53, 54, 55, 56, 0, 0, 0, 0, 51,
	0, 0, 76, 78, 0, 79, 0, 74, 60, 61,
	63, 65, 75, 77, 0, 0, 0, 0, 0, 0,
	0, 0, 66, 67, 68, 69, 70, 71, 0, 0,
	72, 73, 57, 58, 59, 0, 0, 0, 0, 0,
	0, 80, 0, 0, 0, 0, 0, 0, 0, 50,
	0, 0, 62, 64, 52, 53, 54, 55, 56, 0,
	178, 0, 0, 51, 0, 0, 76, 78, 0, 79,
	0, 74, 60, 61, 63, 65, 75, 77, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 67, 68, 69,
	70, 71, 0, 0, 72, 73, 57, 58, 59, 0,
	0, 0, 0, 0, 0, 80, 0, 0, 0, 0,
	0, 0, 0, 50, 0, 0, 62, 64, 52, 53,
	54, 55, 56, 0, 0, 0, 0, 51, 0, 0,
	76, 78, 169, 79, 0, 74, 60, 61, 63, 65,
	75, 77, 0, 0, 0, 0, 0, 0, 0, 0,
	66, 67, 68, 69, 70, 71, 0, 0, 72, 73,
	57, 58, 59, 0, 0, 0, 0, 0, 0, 80,
	0, 0, 0, 0, 0, 0, 0, 50, 157, 0,
	62, 64, 52, 53, 54, 55, 56, 0, 0, 0,
	0, 51, 0, 0, 76, 78, 0, 79, 0, 74,
	60, 61, 63, 65, 75, 77, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 67, 68, 69, 70, 71,
	0, 0, 72, 73, 57, 58, 59, 0, 0, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 0, 0,
	0, 50, 0, 0, 62, 64, 52, 53, 54, 55,
	56, 0, 154, 0, 0, 51, 0, 0, 76, 78,
	0, 79, 0, 74, 60, 61, 63, 65, 75, 77,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 67,
	68, 69, 70, 71, 0, 0, 72, 73, 57, 58,
	59, 0, 0, 0, 0, 0, 0, 80, 0, 0,
	0, 0, 0, 0, 49, 50, 0, 0, 62, 64,
	52, 53, 54, 55, 56, 0, 0, 0, 0, 51,
	0, 0, 76, 78, 0, 79, 0, 74, 60, 61,
	63, 65, 75, 77, 0, 0, 0, 0, 0, 0,
	0, 0, 66, 67, 68, 69, 70, 71, 0, 0,
	72, 73, 57, 58, 59, 0, 0, 0, 0, 0,
	0, 80, 0, 0, 0, 0, 0, 0, 0, 50,
	0, 0, 62, 64, 52, 53, 54, 55, 56, 0,
	0, 0, 0, 51, 0, 0, 76, 78, 0, 79,
	0, 74, 60, 61, 63, 65, 75, 77, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 67, 68, 69,
	70, 71, 0, 0, 72, 73, 57, 58, 59, 0,
	0, 0, 0, 0, 0, 80, 0, 0, 0, 0,
	0, 0, 0, 50, 0, 0, 62, 64, 52, 53,
	54, 55, 56, 0, 0, 0, 0, 51, 0, 0,
	76, 173, 0, 79, 0, 74, 60, 61, 63, 65,
	75, 77, 0, 0, 0, 0, 0, 0, 0, 0,
	66, 67, 68, 69, 70, 71, 0, 0, 72, 73,
	57, 58, 59, 0, 0, 0, 0, 0, 0, 80,
	0, 0, 0, 0, 0, 0, 0, 50, 0, 0,
	62, 64, 52, 53, 54, 55, 56, 0, 0, 0,
	0, 161, 0, 0, 76, 78, 0, 79, 0, 74,
	60, 61, 63, 65, 75, 77, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 67, 68, 69, 70, 71,
	0, 0, 72, 73, 57, 58, 59, 0, 0, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 0, 0,
	0, 50, 0, 0, 62, 64, 52, 53, 54, 55,
	56, 60, 61, 63, 65, 160, 77, 0, 76, 78,
	0, 79, 0, 74, 0, 66, 67, 68, 69, 70,
	71, 0, 0, 72, 73, 57, 58, 59, 0, 0,
	0, 0, 0, 0, 80, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 62, 64, 52, 53, 54,
	55, 56, 0, 0, 0, 0, 51, 0, 0, 76,
	78, 0, 79, 0, 74, 22, 23, 29, 0, 0,
	34, 14, 9, 15, 43, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 38, 30, 31, 32, 16, 20,
	0, 0, 0, 0, 0, 0, 0, 0, 12, 13,
	0, 0, 0, 0, 0, 21, 0, 0, 39, 0,
	40, 41, 0, 33, 0, 0, 0, 19, 0, 0,
	0, 0, 0, 0, 0, 24, 28, 0, 0, 0,
	36, 0, 0, 0, 25, 26, 27, 37, 0, 35,
	60, 61, 63, 65, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 67, 68, 69, 70, 71,
	0, 0, 72, 73, 57, 58, 59, 0, 0, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 62, 64, 52, 53, 54, 55,
	56, 0, 63, 65, 0, 51, 0, 0, 76, 78,
	0, 79, 0, 74, 66, 67, 68, 69, 70, 71,
	0, 0, 72, 73, 57, 58, 59, 0, 0, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 62, 64, 52, 53, 54, 55,
	56, 22, 23, 29, 0, 51, 34, 0, 76, 78,
	0, 79, 0, 74, 0, 0, 0, 0, 0, 0,
	38, 30, 31, 32, 0, 0, 0, 0, 0, 0,
	0, 226, 23, 29, 0, 0, 34, 0, 0, 0,
	0, 0, 0, 0, 39, 0, 40, 41, 0, 33,
	38, 30, 31, 32, 0, 0, 0, 0, 0, 0,
	0, 24, 28, 219, 23, 29, 36, 0, 34, 0,
	25, 26, 27, 37, 39, 35, 40, 41, 0, 33,
	0, 0, 38, 30, 31, 32, 0, 0, 0, 0,
	0, 24, 28, 108, 23, 29, 36, 0, 34, 0,
	25, 26, 27, 37, 0, 35, 39, 0, 40, 41,
	0, 33, 38, 30, 31, 32, 0, 0, 0, 0,
	0, 0, 0, 24, 28, 100, 23, 29, 36, 0,
	34, 0, 25, 26, 27, 37, 39, 35, 40, 41,
	0, 33, 0, 0, 38, 30, 31, 32, 0, 0,
	0, 0, 0, 24, 28, 98, 23, 29, 36, 0,
	34, 0, 25, 26, 27, 37, 0, 35, 39, 0,
	40, 41, 0, 33, 38, 30, 31, 32, 0, 0,
	0, 0, 0, 0, 0, 24, 28, 0, 0, 0,
	36, 0, 0, 0, 25, 26, 27, 37, 39, 35,
	40, 41, 0, 33, 66, 67, 68, 69, 70, 71,
	0, 0, 0, 0, 57, 24, 28, 0, 0, 0,
	36, 0, 0, 80, 25, 26, 27, 37, 0, 35,
	0, 0, 0, 0, 0, 0, 0, 0, 54, 55,
	56, 0, 0, 0, 0, 51, 0, 0, 76, 78,
	0, 79, 0, 74,
}
var yyPact = [...]int{

	-64, -1000, 1981, -64, -64, -1000, -1000, -1000, -1000, 248,
	1597, 126, -1000, -1000, 2157, 2157, 237, -1000, 168, 2157,
	-64, 2157, -63, -1000, 2157, 2157, 2157, 2311, 2281, -1000,
	-1000, -1000, -1000, -1000, 75, -64, -64, 2157, 38, 2249,
	37, 2157, 99, 2157, -1000, 316, -1000, 100, -1000, 2157,
	2157, 236, 2157, 2157, 2157, 2157, 2157, 2157, 2157, 2157,
	2157, 2157, 2157, 2157, 2157, 2157, 2157, 2157, 2157, 2157,
	2157, 2157, -1000, -1000, 2157, 2157, 2157, 2157, 2157, 2157,
	2157, 2157, 98, 1661, 1661, 58, 234, 120, 19, 1533,
	116, 176, 1469, 2157, 2157, 192, 192, 192, -63, 1853,
	-63, 1789, 232, 22, 2157, 202, 1405, 209, -29, 1725,
	205, 1661, -64, 1341, -1000, 2157, -64, 1661, 1277, -1000,
	2333, 2333, 192, 192, 192, 1661, 68, 68, 2093, 2093,
	68, 68, 68, 68, 1661, 1661, 1661, 1661, 1661, 1661,
	1661, 1904, 1661, 2043, 111, 509, 1661, -1000, 1661, -64,
	-64, 181, 2157, 2157, -64, 2157, -64, -64, 65, 445,
	228, 225, 34, 199, 214, -36, -50, -1000, 105, -1000,
	15, -1000, 2157, 2157, 11, 209, 209, 2219, -64, -1000,
	208, 2157, -34, -1000, -1000, 2157, 2187, 85, 2157, 1213,
	1149, 82, -1000, 80, 93, 158, -38, -1000, -1000, 2157,
	-1000, -1000, -64, -47, -41, 191, -64, -71, -64, 79,
	2157, 206, -1000, 33, 29, -1000, 14, 57, 1661, -63,
	-1000, -1000, 1661, -1000, 1085, 1661, -63, -1000, 1021, 2157,
	2157, -1000, -1000, -1000, -1000, -1000, 2157, 86, -1000, -1000,
	-1000, 957, -64, -64, -64, -54, 170, -1000, 115, -1000,
	1661, -1000, -65, -1000, -68, -1000, -1000, 2157, 2157, 114,
	-1000, -64, 893, 829, 765, -64, -1000, 77, -64, -64,
	-64, -1000, -1000, -1000, -1000, -1000, 701, 381, -1000, -64,
	-1000, 2157, 71, -64, -64, -64, -64, -1000, 67, 66,
	-64, -1000, -1000, 2157, 56, 637, -1000, 54, 52, -64,
	-1000, -1000, -1000, 50, 573, -1000, -64, -1000, -1000, -1000,
	-1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 11, 277, 216, 276, 6, 5, 274, 269, 266,
	4, 0, 18, 14, 263, 1, 261, 2, 48, 271,
	260,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 9, 9, 8, 4, 4, 7, 7,
	7, 7, 7, 6, 5, 15, 16, 16, 16, 17,
	17, 17, 14, 13, 13, 13, 10, 10, 12, 12,
	12, 12, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 18, 18, 19, 19, 20, 20,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 2, 3, 4, 2, 3, 3,
	1, 1, 2, 2, 5, 1, 8, 9, 9, 5,
	5, 5, 1, 0, 2, 4, 8, 6, 0, 2,
	2, 2, 2, 5, 4, 3, 0, 1, 4, 0,
	1, 4, 3, 1, 4, 4, 1, 3, 0, 1,
	4, 4, 1, 1, 2, 2, 2, 2, 4, 2,
	4, 1, 1, 1, 1, 1, 5, 3, 7, 8,
	8, 9, 5, 6, 5, 6, 3, 4, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	3, 3, 3, 3, 5, 4, 6, 5, 5, 4,
	6, 5, 4, 4, 6, 6, 4, 5, 7, 7,
	9, 3, 2, 0, 1, 1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -18, -2, -19, -20, 71, 81, -3, 11,
	-11, -13, 37, 38, 10, 12, 27, -4, 15, 56,
	28, 44, 4, 5, 64, 73, 74, 75, 65, 6,
	24, 25, 26, 52, 9, 78, 69, 76, 23, 47,
	49, 50, -12, 13, -18, -19, -20, -17, 4, 57,
	58, 72, 63, 64, 65, 66, 67, 41, 42, 43,
	17, 18, 61, 19, 62, 20, 31, 32, 33, 34,
	35, 36, 39, 40, 80, 21, 75, 22, 76, 78,
	50, 57, -12, -11, -11, 4, 53, 4, -14, -11,
	-13, -1, -11, 76, 78, -11, -11, -11, 4, -11,
	4, -11, 76, 4, -18, -18, -11, 76, 4, -11,
	76, -11, 60, -11, -3, 57, 60, -11, -11, 4,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -12, -11, -11, -13, -11, 60,
	69, 4, 57, 71, 69, 57, 29, 59, -12, -11,
	72, 72, -17, 4, 76, -12, -16, -15, 6, 77,
	-10, 4, 76, 76, -10, 48, 51, -18, 69, -13,
	-18, 59, 8, 77, 79, 59, -18, -1, 16, -11,
	-11, -1, -13, -1, -7, -18, 8, 77, 79, 59,
	4, 4, 77, 8, -17, 4, 60, -18, 60, -18,
	59, 72, 77, -12, -12, 77, -10, -10, -11, 4,
	-1, 4, -11, 77, -11, -11, 4, 70, -11, 54,
	71, 70, 70, 70, -6, -5, 45, 46, -6, -5,
	77, -11, -18, 77, 77, 8, -18, 79, -18, 70,
	-11, 4, 8, 77, 8, 77, 77, 60, 60, -9,
	79, 69, -11, -11, -11, 59, 79, -1, -18, -18,
	77, 79, -15, 70, 77, 77, -11, -11, -8, 14,
	70, 55, -1, 69, 69, 59, -18, 70, -1, -1,
	-18, 77, 77, 60, -1, -11, 70, -1, -1, -18,
	-1, 70, 70, -1, -11, 70, 69, 70, 70, -1,
	70, 77, -1,
}
var yyDef = [...]int{

	-2, -2, -2, 123, 124, 125, 127, 128, 4, 39,
	-2, 0, 10, 11, 48, 0, 0, 15, 0, 48,
	-2, 0, 52, 53, 0, 0, 0, 0, 0, 61,
	62, 63, 64, 65, 0, 123, 123, 0, 0, 0,
	0, 0, 0, 0, 2, -2, 126, 7, 40, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 98, 99, 0, 0, 0, 0, 48, 0,
	0, 48, 12, 49, 13, 0, 0, 0, 0, -2,
	0, 0, 0, 48, 0, 54, 55, 56, -2, 0,
	-2, 0, 39, 0, 48, 36, 0, 0, 52, 0,
	0, 122, 123, 0, 5, 48, 123, 8, 0, 67,
	78, 79, 80, 81, 82, 83, 84, 85, -2, -2,
	88, 89, 90, 91, 92, 93, 94, 95, 96, 97,
	100, 101, 102, 103, 0, 0, 121, 9, -2, 123,
	-2, 0, 0, 0, -2, 48, -2, 28, 0, 0,
	0, 0, 0, 40, 39, 123, 123, 37, 0, 76,
	0, 46, 48, 48, 0, 0, 0, 0, -2, 6,
	0, 0, 0, 109, 113, 0, 0, 0, 0, 0,
	0, 0, 42, 0, 0, 0, 0, 105, 112, 0,
	58, 60, 123, 0, 0, 40, 123, 0, 123, 0,
	0, 0, 77, 0, 0, 116, 0, 0, -2, -2,
	23, 41, 66, 108, 0, 50, -2, 14, 0, 0,
	0, 19, 20, 21, 31, 32, 0, 0, 29, 30,
	104, 0, -2, 123, 123, 0, 0, 72, 0, 74,
	35, 47, 0, -2, 0, -2, 117, 0, 0, 0,
	115, -2, 0, 0, 0, 123, 114, 0, -2, -2,
	123, 73, 38, 75, -2, -2, 0, 0, 24, -2,
	27, 0, 0, -2, -2, 123, -2, 68, 0, 0,
	-2, 118, 119, 0, 0, 0, 16, 0, 0, -2,
	34, 69, 70, 0, 0, 26, -2, 17, 18, 33,
	71, 120, 25,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	81, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 73, 3, 3, 3, 67, 75, 3,
	76, 77, 65, 63, 60, 64, 72, 66, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 59, 71,
	62, 57, 61, 58, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 78, 3, 79, 74, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 69, 80, 70,
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
		//line parser.y:69
		{
			yyVAL.compstmt = nil
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:73
		{
			yyVAL.compstmt = yyDollar[1].stmts
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:78
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:85
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:92
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
		//line parser.y:103
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].expr_many}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:108
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: nil}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:113
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "=", Rhss: []ast.Expr{yyDollar[3].expr}}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:117
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:121
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:126
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:131
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:136
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 14:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:141
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyDollar[2].tok.Lit, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:146
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 16:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:151
		{
			yyVAL.stmt = &ast.ForStmt{Var: yyDollar[3].tok.Lit, Value: yyDollar[5].expr, Stmts: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.y:156
		{
			yyVAL.stmt = &ast.NumForStmt{Name: yyDollar[2].tok.Lit, Expr1: yyDollar[4].expr, Expr2: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.y:161
		{
			yyVAL.stmt = &ast.CForStmt{Expr1: yyDollar[2].expr_lets, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:166
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:171
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[2].compstmt, Catch: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:176
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:181
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:187
		{
			yyVAL.stmt_elsifs = []ast.Stmt{}
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:191
		{
			yyVAL.stmt_elsifs = append(yyDollar[1].stmt_elsifs, yyDollar[2].stmt_elsif)
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:197
		{
			yyVAL.stmt_elsif = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt}
		}
	case 26:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:203
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, ElseIf: yyDollar[5].stmt_elsifs, Else: yyDollar[7].compstmt}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 27:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:208
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, ElseIf: yyDollar[5].stmt_elsifs, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:214
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:218
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_case}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:222
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_default}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:226
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_case)
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:230
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
		//line parser.y:241
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[5].compstmt}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:247
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyDollar[4].compstmt}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:253
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:258
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:262
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:266
		{
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[4].expr_pair)
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:271
		{
			yyVAL.expr_idents = []string{}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:275
		{
			yyVAL.expr_idents = []string{yyDollar[1].tok.Lit}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:279
		{
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, yyDollar[4].tok.Lit)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:284
		{
			yyVAL.expr_lets = &ast.LetsExpr{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:290
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:294
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:298
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:303
		{
			yyVAL.typ = ast.Type{Name: yyDollar[1].tok.Lit}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:307
		{
			yyVAL.typ = ast.Type{Name: yyDollar[1].typ.Name + "." + yyDollar[3].tok.Lit}
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:312
		{
			yyVAL.exprs = nil
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:316
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:320
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:324
		{
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:330
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:335
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:340
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:345
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:350
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:355
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:360
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:365
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:370
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:375
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:380
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:385
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:390
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:395
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 66:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:400
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[1].expr, Lhs: yyDollar[3].expr, Rhs: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:405
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: yyDollar[3].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 68:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:410
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyDollar[3].expr_idents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:415
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyDollar[3].tok.Lit}, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 70:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:420
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Args: yyDollar[4].expr_idents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 71:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.y:425
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Args: []string{yyDollar[4].tok.Lit}, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 72:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:430
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 73:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:435
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 74:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:440
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
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:449
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
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:458
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:463
		{
			yyVAL.expr = &ast.NewExpr{Type: yyDollar[3].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:468
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:473
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:478
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:483
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:488
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:493
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:498
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:503
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:508
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:513
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:518
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:523
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:528
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:533
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:538
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:543
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:548
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:553
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:558
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:563
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:568
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:573
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:578
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:583
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:588
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:593
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 104:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:598
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 105:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:603
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 106:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:608
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 107:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:613
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 108:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:618
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 109:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:623
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 110:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:628
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 111:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:633
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:638
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 113:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:643
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 114:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:648
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 115:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:653
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 116:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:658
		{
			yyVAL.expr = &ast.MakeExpr{Type: yyDollar[3].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 117:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:663
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].typ.Name, SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 118:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:668
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].typ.Name, SizeExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 119:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:673
		{
			yyVAL.expr = &ast.MakeArrayExpr{Type: yyDollar[4].typ.Name, LenExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 120:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.y:678
		{
			yyVAL.expr = &ast.MakeArrayExpr{Type: yyDollar[4].typ.Name, LenExpr: yyDollar[6].expr, CapExpr: yyDollar[8].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:683
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:688
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:699
		{
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:702
		{
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:707
		{
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:710
		{
		}
	}
	goto yystack /* stack new state and value */
}
