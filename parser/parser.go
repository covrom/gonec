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

//line ./parser/parser.y:707

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 3,
	-2, 122,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	61, 47,
	-2, 1,
	-1, 10,
	61, 48,
	-2, 22,
	-1, 20,
	29, 3,
	-2, 122,
	-1, 47,
	61, 47,
	-2, 123,
	-1, 98,
	1, 56,
	8, 56,
	14, 56,
	29, 56,
	45, 56,
	46, 56,
	54, 56,
	55, 56,
	58, 56,
	60, 56,
	61, 56,
	70, 56,
	71, 56,
	76, 56,
	79, 56,
	81, 56,
	82, 56,
	-2, 51,
	-1, 100,
	1, 58,
	8, 58,
	14, 58,
	29, 58,
	45, 58,
	46, 58,
	54, 58,
	55, 58,
	58, 58,
	60, 58,
	61, 58,
	70, 58,
	71, 58,
	76, 58,
	79, 58,
	81, 58,
	82, 58,
	-2, 51,
	-1, 131,
	17, 0,
	18, 0,
	-2, 85,
	-1, 132,
	17, 0,
	18, 0,
	-2, 86,
	-1, 151,
	61, 48,
	-2, 42,
	-1, 153,
	71, 3,
	-2, 122,
	-1, 156,
	71, 3,
	-2, 122,
	-1, 157,
	71, 3,
	-2, 122,
	-1, 181,
	14, 3,
	55, 3,
	71, 3,
	-2, 122,
	-1, 225,
	61, 49,
	-2, 43,
	-1, 226,
	1, 44,
	14, 44,
	29, 44,
	45, 44,
	46, 44,
	55, 44,
	58, 44,
	61, 50,
	71, 44,
	81, 44,
	82, 44,
	-2, 51,
	-1, 232,
	1, 50,
	8, 50,
	14, 50,
	29, 50,
	45, 50,
	46, 50,
	55, 50,
	61, 50,
	71, 50,
	76, 50,
	79, 50,
	81, 50,
	82, 50,
	-2, 51,
	-1, 244,
	71, 3,
	-2, 122,
	-1, 254,
	1, 106,
	8, 106,
	14, 106,
	29, 106,
	45, 106,
	46, 106,
	54, 106,
	55, 106,
	58, 106,
	60, 106,
	61, 106,
	70, 106,
	71, 106,
	76, 106,
	79, 106,
	81, 106,
	82, 106,
	-2, 104,
	-1, 256,
	1, 110,
	8, 110,
	14, 110,
	29, 110,
	45, 110,
	46, 110,
	54, 110,
	55, 110,
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
	-1, 261,
	71, 3,
	-2, 122,
	-1, 264,
	45, 3,
	46, 3,
	71, 3,
	-2, 122,
	-1, 268,
	71, 3,
	-2, 122,
	-1, 269,
	71, 3,
	-2, 122,
	-1, 274,
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
	-1, 275,
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
	-1, 278,
	71, 3,
	-2, 122,
	-1, 282,
	71, 3,
	-2, 122,
	-1, 283,
	45, 3,
	46, 3,
	71, 3,
	-2, 122,
	-1, 289,
	71, 3,
	-2, 122,
	-1, 299,
	14, 3,
	55, 3,
	71, 3,
	-2, 122,
}

const yyPrivate = 57344

const yyLast = 2689

var yyAct = [...]int{

	84, 171, 49, 10, 196, 111, 44, 11, 197, 249,
	1, 159, 215, 6, 7, 104, 85, 213, 93, 94,
	89, 83, 91, 175, 94, 95, 96, 97, 99, 101,
	178, 90, 6, 7, 120, 102, 255, 6, 7, 107,
	253, 110, 168, 202, 115, 184, 117, 177, 10, 246,
	275, 218, 121, 114, 123, 124, 125, 126, 127, 128,
	129, 130, 131, 132, 133, 134, 135, 136, 137, 138,
	139, 140, 141, 142, 108, 274, 143, 144, 145, 146,
	270, 148, 149, 151, 120, 245, 147, 2, 103, 152,
	150, 46, 241, 152, 229, 162, 152, 278, 152, 209,
	161, 177, 198, 199, 256, 301, 166, 172, 254, 198,
	199, 203, 169, 185, 174, 179, 300, 298, 296, 295,
	151, 292, 286, 251, 105, 106, 237, 182, 238, 236,
	233, 153, 152, 119, 116, 195, 120, 240, 280, 217,
	155, 88, 82, 113, 157, 5, 8, 198, 199, 190,
	48, 247, 210, 172, 279, 228, 191, 221, 212, 207,
	206, 113, 167, 154, 189, 200, 208, 192, 193, 201,
	194, 211, 273, 122, 86, 50, 170, 259, 277, 222,
	160, 225, 219, 220, 17, 3, 0, 112, 230, 231,
	87, 234, 227, 48, 118, 0, 4, 0, 0, 239,
	47, 0, 0, 0, 180, 0, 242, 0, 183, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 252, 67,
	68, 69, 70, 71, 72, 258, 0, 0, 0, 58,
	0, 0, 0, 0, 0, 0, 262, 0, 81, 0,
	188, 0, 0, 0, 266, 0, 160, 0, 0, 231,
	0, 0, 272, 0, 0, 267, 0, 214, 216, 0,
	52, 0, 0, 77, 0, 79, 80, 0, 75, 0,
	0, 0, 281, 0, 0, 284, 0, 0, 0, 287,
	288, 291, 0, 0, 0, 0, 0, 0, 0, 290,
	0, 0, 0, 293, 294, 0, 0, 244, 0, 0,
	297, 248, 0, 250, 0, 0, 0, 0, 0, 0,
	302, 0, 0, 0, 0, 0, 22, 23, 29, 0,
	0, 35, 14, 9, 15, 45, 0, 18, 264, 0,
	0, 0, 0, 268, 269, 39, 30, 31, 32, 16,
	20, 0, 0, 0, 0, 0, 0, 0, 0, 12,
	13, 283, 0, 0, 0, 0, 21, 0, 289, 40,
	0, 41, 43, 42, 33, 0, 0, 0, 19, 34,
	0, 0, 0, 0, 0, 0, 0, 24, 28, 0,
	0, 0, 37, 0, 0, 25, 26, 27, 0, 38,
	36, 0, 0, 6, 7, 61, 62, 64, 66, 76,
	78, 0, 0, 0, 0, 0, 0, 0, 0, 67,
	68, 69, 70, 71, 72, 0, 0, 73, 74, 58,
	59, 60, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 224,
	63, 65, 53, 54, 55, 56, 57, 0, 0, 0,
	52, 0, 0, 77, 223, 79, 80, 0, 75, 61,
	62, 64, 66, 76, 78, 0, 0, 0, 0, 0,
	0, 0, 0, 67, 68, 69, 70, 71, 72, 0,
	0, 73, 74, 58, 59, 60, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 205, 0, 63, 65, 53, 54, 55, 56,
	57, 0, 0, 0, 52, 0, 0, 77, 0, 79,
	80, 204, 75, 61, 62, 64, 66, 76, 78, 0,
	0, 0, 0, 0, 0, 0, 0, 67, 68, 69,
	70, 71, 72, 0, 0, 73, 74, 58, 59, 60,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 187, 0, 63, 65,
	53, 54, 55, 56, 57, 0, 0, 0, 52, 0,
	0, 77, 0, 79, 80, 186, 75, 61, 62, 64,
	66, 76, 78, 0, 0, 0, 0, 0, 0, 0,
	0, 67, 68, 69, 70, 71, 72, 0, 0, 73,
	74, 58, 59, 60, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 63, 65, 53, 54, 55, 56, 57, 0,
	299, 0, 52, 0, 0, 77, 0, 79, 80, 0,
	75, 61, 62, 64, 66, 76, 78, 0, 0, 0,
	0, 0, 0, 0, 0, 67, 68, 69, 70, 71,
	72, 0, 0, 73, 74, 58, 59, 60, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 63, 65, 53, 54,
	55, 56, 57, 0, 0, 0, 52, 0, 0, 77,
	285, 79, 80, 0, 75, 61, 62, 64, 66, 76,
	78, 0, 0, 0, 0, 0, 0, 0, 0, 67,
	68, 69, 70, 71, 72, 0, 0, 73, 74, 58,
	59, 60, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	63, 65, 53, 54, 55, 56, 57, 0, 282, 0,
	52, 0, 0, 77, 0, 79, 80, 0, 75, 61,
	62, 64, 66, 76, 78, 0, 0, 0, 0, 0,
	0, 0, 0, 67, 68, 69, 70, 71, 72, 0,
	0, 73, 74, 58, 59, 60, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 63, 65, 53, 54, 55, 56,
	57, 0, 0, 0, 52, 0, 0, 77, 276, 79,
	80, 0, 75, 61, 62, 64, 66, 76, 78, 0,
	0, 0, 0, 0, 0, 0, 0, 67, 68, 69,
	70, 71, 72, 0, 0, 73, 74, 58, 59, 60,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 63, 65,
	53, 54, 55, 56, 57, 0, 0, 0, 52, 0,
	0, 77, 0, 79, 80, 265, 75, 61, 62, 64,
	66, 76, 78, 0, 0, 0, 0, 0, 0, 0,
	0, 67, 68, 69, 70, 71, 72, 0, 0, 73,
	74, 58, 59, 60, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	263, 0, 63, 65, 53, 54, 55, 56, 57, 0,
	0, 0, 52, 0, 0, 77, 0, 79, 80, 0,
	75, 61, 62, 64, 66, 76, 78, 0, 0, 0,
	0, 0, 0, 0, 0, 67, 68, 69, 70, 71,
	72, 0, 0, 73, 74, 58, 59, 60, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 63, 65, 53, 54,
	55, 56, 57, 0, 261, 0, 52, 0, 0, 77,
	0, 79, 80, 0, 75, 61, 62, 64, 66, 76,
	78, 0, 0, 0, 0, 0, 0, 0, 0, 67,
	68, 69, 70, 71, 72, 0, 0, 73, 74, 58,
	59, 60, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	63, 65, 53, 54, 55, 56, 57, 0, 0, 0,
	52, 0, 0, 77, 0, 79, 80, 260, 75, 61,
	62, 64, 66, 76, 78, 0, 0, 0, 0, 0,
	0, 0, 0, 67, 68, 69, 70, 71, 72, 0,
	0, 73, 74, 58, 59, 60, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 63, 65, 53, 54, 55, 56,
	57, 0, 0, 0, 52, 0, 0, 77, 257, 79,
	80, 0, 75, 61, 62, 64, 66, 76, 78, 0,
	0, 0, 0, 0, 0, 0, 0, 67, 68, 69,
	70, 71, 72, 0, 0, 73, 74, 58, 59, 60,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 243, 63, 65,
	53, 54, 55, 56, 57, 0, 0, 0, 52, 0,
	0, 77, 0, 79, 80, 0, 75, 61, 62, 64,
	66, 76, 78, 0, 0, 0, 0, 0, 0, 0,
	0, 67, 68, 69, 70, 71, 72, 0, 0, 73,
	74, 58, 59, 60, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 235, 0, 0, 0, 0, 0,
	0, 0, 63, 65, 53, 54, 55, 56, 57, 0,
	0, 0, 52, 0, 0, 77, 0, 79, 80, 0,
	75, 61, 62, 64, 66, 76, 78, 0, 0, 0,
	0, 0, 0, 0, 0, 67, 68, 69, 70, 71,
	72, 0, 0, 73, 74, 58, 59, 60, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 63, 65, 53, 54,
	55, 56, 57, 0, 181, 0, 52, 0, 0, 77,
	0, 79, 80, 0, 75, 61, 62, 64, 66, 76,
	78, 0, 0, 0, 0, 0, 0, 0, 0, 67,
	68, 69, 70, 71, 72, 0, 0, 73, 74, 58,
	59, 60, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	63, 65, 53, 54, 55, 56, 57, 0, 0, 0,
	52, 0, 0, 77, 173, 79, 80, 0, 75, 61,
	62, 64, 66, 76, 78, 0, 0, 0, 0, 0,
	0, 0, 0, 67, 68, 69, 70, 71, 72, 0,
	0, 73, 74, 58, 59, 60, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 165, 63, 65, 53, 54, 55, 56,
	57, 0, 0, 0, 52, 0, 0, 77, 0, 79,
	80, 0, 75, 61, 62, 64, 66, 76, 78, 0,
	0, 0, 0, 0, 0, 0, 0, 67, 68, 69,
	70, 71, 72, 0, 0, 73, 74, 58, 59, 60,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 158, 0, 63, 65,
	53, 54, 55, 56, 57, 0, 0, 0, 52, 0,
	0, 77, 0, 79, 80, 0, 75, 61, 62, 64,
	66, 76, 78, 0, 0, 0, 0, 0, 0, 0,
	0, 67, 68, 69, 70, 71, 72, 0, 0, 73,
	74, 58, 59, 60, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 63, 65, 53, 54, 55, 56, 57, 0,
	156, 0, 52, 0, 0, 77, 0, 79, 80, 0,
	75, 61, 62, 64, 66, 76, 78, 0, 0, 0,
	0, 0, 0, 0, 0, 67, 68, 69, 70, 71,
	72, 0, 0, 73, 74, 58, 59, 60, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 0,
	0, 0, 51, 0, 0, 0, 63, 65, 53, 54,
	55, 56, 57, 0, 0, 0, 52, 0, 0, 77,
	0, 79, 80, 0, 75, 22, 23, 29, 0, 0,
	35, 14, 9, 15, 45, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 39, 30, 31, 32, 16, 20,
	0, 0, 0, 0, 0, 0, 0, 0, 12, 13,
	0, 0, 0, 0, 0, 21, 0, 0, 40, 0,
	41, 43, 42, 33, 0, 0, 0, 19, 34, 0,
	0, 0, 0, 0, 0, 0, 24, 28, 0, 0,
	0, 37, 0, 0, 25, 26, 27, 0, 38, 36,
	61, 62, 64, 66, 76, 78, 0, 0, 0, 0,
	0, 0, 0, 0, 67, 68, 69, 70, 71, 72,
	0, 0, 73, 74, 58, 59, 60, 0, 0, 0,
	0, 0, 0, 81, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 63, 65, 53, 54, 55,
	56, 57, 0, 0, 0, 52, 0, 0, 77, 0,
	79, 80, 0, 75, 61, 62, 64, 66, 76, 78,
	0, 0, 0, 0, 0, 0, 0, 0, 67, 68,
	69, 70, 71, 72, 0, 0, 73, 74, 58, 59,
	60, 0, 0, 0, 0, 0, 0, 81, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 63,
	65, 53, 54, 55, 56, 57, 0, 0, 0, 52,
	0, 0, 77, 0, 176, 80, 0, 75, 61, 62,
	64, 66, 76, 78, 0, 0, 0, 0, 0, 0,
	0, 0, 67, 68, 69, 70, 71, 72, 0, 0,
	73, 74, 58, 59, 60, 0, 0, 0, 0, 0,
	0, 81, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 63, 65, 53, 54, 55, 56, 57,
	0, 0, 0, 164, 0, 0, 77, 0, 79, 80,
	0, 75, 61, 62, 64, 66, 76, 78, 0, 0,
	0, 0, 0, 0, 0, 0, 67, 68, 69, 70,
	71, 72, 0, 0, 73, 74, 58, 59, 60, 0,
	0, 0, 0, 0, 0, 81, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 63, 65, 53,
	54, 55, 56, 57, 0, 0, 0, 163, 0, 0,
	77, 0, 79, 80, 0, 75, 61, 62, 64, 66,
	0, 78, 0, 0, 0, 0, 0, 0, 0, 0,
	67, 68, 69, 70, 71, 72, 0, 0, 73, 74,
	58, 59, 60, 0, 0, 0, 0, 0, 0, 81,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 63, 65, 53, 54, 55, 56, 57, 0, 0,
	0, 52, 0, 0, 77, 0, 79, 80, 0, 75,
	61, 62, 64, 66, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 67, 68, 69, 70, 71, 72,
	0, 0, 73, 74, 58, 59, 60, 0, 0, 0,
	0, 0, 0, 81, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 63, 65, 53, 54, 55,
	56, 57, 64, 66, 0, 52, 0, 0, 77, 0,
	79, 80, 0, 75, 67, 68, 69, 70, 71, 72,
	0, 0, 73, 74, 58, 59, 60, 0, 0, 0,
	0, 0, 0, 81, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 63, 65, 53, 54, 55,
	56, 57, 232, 23, 29, 52, 0, 35, 77, 0,
	79, 80, 0, 75, 0, 0, 0, 0, 0, 0,
	0, 39, 30, 31, 32, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 22, 23, 29,
	0, 0, 35, 0, 0, 40, 0, 41, 43, 42,
	33, 0, 0, 0, 0, 34, 39, 30, 31, 32,
	0, 0, 0, 24, 28, 0, 0, 0, 37, 0,
	0, 25, 26, 27, 0, 38, 36, 271, 0, 0,
	40, 0, 41, 43, 42, 33, 0, 0, 0, 0,
	34, 0, 0, 92, 0, 22, 23, 29, 24, 28,
	35, 0, 0, 37, 0, 0, 25, 26, 27, 0,
	38, 36, 0, 0, 39, 30, 31, 32, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	232, 23, 29, 0, 0, 35, 0, 0, 40, 0,
	41, 43, 42, 33, 0, 0, 0, 0, 34, 39,
	30, 31, 32, 0, 0, 0, 24, 28, 0, 0,
	0, 37, 0, 0, 25, 26, 27, 0, 38, 36,
	0, 0, 0, 40, 0, 41, 43, 42, 33, 0,
	0, 0, 0, 34, 0, 0, 0, 0, 226, 23,
	29, 24, 28, 35, 0, 0, 37, 0, 0, 25,
	26, 27, 0, 38, 36, 0, 0, 39, 30, 31,
	32, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 40, 0, 41, 43, 42, 33, 0, 0, 0,
	0, 34, 0, 0, 0, 0, 0, 0, 0, 24,
	28, 0, 0, 0, 37, 0, 0, 25, 26, 27,
	0, 38, 36, 67, 68, 69, 70, 71, 72, 0,
	0, 73, 74, 58, 0, 0, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 53, 54, 55, 56,
	57, 109, 23, 29, 52, 0, 35, 77, 0, 79,
	80, 0, 75, 0, 0, 0, 0, 0, 0, 0,
	39, 30, 31, 32, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 100, 23, 29, 0,
	0, 35, 0, 0, 40, 0, 41, 43, 42, 33,
	0, 0, 0, 0, 34, 39, 30, 31, 32, 0,
	0, 0, 24, 28, 0, 0, 0, 37, 0, 0,
	25, 26, 27, 0, 38, 36, 0, 0, 0, 40,
	0, 41, 43, 42, 33, 0, 0, 0, 0, 34,
	0, 0, 0, 0, 98, 23, 29, 24, 28, 35,
	0, 0, 37, 0, 0, 25, 26, 27, 0, 38,
	36, 0, 0, 39, 30, 31, 32, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 40, 0, 41,
	43, 42, 33, 0, 0, 0, 0, 34, 0, 0,
	0, 0, 0, 0, 0, 24, 28, 0, 0, 0,
	37, 0, 0, 25, 26, 27, 0, 38, 36, 67,
	68, 69, 70, 71, 72, 0, 0, 0, 0, 58,
	0, 0, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 55, 56, 57, 0, 0, 0,
	52, 0, 0, 77, 0, 79, 80, 0, 75,
}
var yyPact = [...]int{

	-68, -1000, 1671, -68, -68, -1000, -1000, -1000, -1000, 171,
	1594, 84, -1000, -1000, 2261, 2261, 170, -1000, 137, 2261,
	-68, 2203, -59, -1000, 2261, 2261, 2261, 2560, 2502, -1000,
	-1000, -1000, -1000, -1000, 2261, 11, -68, -68, 2261, -3,
	2467, 139, -24, 2261, 73, 2261, -1000, 312, -1000, 75,
	-1000, 2261, 169, 2261, 2261, 2261, 2261, 2261, 2261, 2261,
	2261, 2261, 2261, 2261, 2261, 2261, 2261, 2261, 2261, 2261,
	2261, 2261, 2261, -1000, -1000, 2261, 2261, 2261, 2261, 2261,
	2261, 2261, 2261, 71, 1733, 1733, 61, 159, 82, 1530,
	115, 1466, -68, 2261, 2261, 188, 188, 188, -59, 1925,
	-59, 1861, 1402, 158, -35, 2261, 147, 1338, 157, -54,
	1797, 29, -47, -1000, 2261, 1733, -68, 1274, -1000, 2261,
	-68, 1733, -1000, 2608, 2608, 188, 188, 188, 1733, 2402,
	2402, 2103, 2103, 2402, 2402, 2402, 2402, 1733, 1733, 1733,
	1733, 1733, 1733, 1733, 1989, 1733, 2053, 37, 506, 1733,
	-1000, 1733, -68, -68, 133, 2261, -68, -68, -68, 64,
	102, 35, 442, 156, 155, 2261, 23, 144, 154, -44,
	-49, -1000, 79, -1000, -25, 2261, 2261, 153, 2261, 378,
	2354, -68, -1000, 151, 18, -1000, -1000, 2261, 2296, 59,
	2261, 1210, 58, 55, 57, -1000, -1000, -1000, 2261, 77,
	-1000, -1000, 16, -1000, -1000, 2261, -1000, -1000, 1146, -68,
	9, -27, 143, -68, -70, -68, 52, 2261, -1000, 32,
	28, -1000, 1082, -1000, 2261, 1733, -59, -1000, -1000, -1000,
	1018, 1733, -59, -1000, 954, 2261, -1000, -1000, -1000, 890,
	-68, -1000, 826, 2261, -68, -68, -68, 4, 2168, -1000,
	101, -1000, 1733, -1, -1000, -26, -1000, -1000, 762, 83,
	-1000, -68, 698, -68, -68, -1000, 634, 51, -68, -68,
	-68, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -68, -1000,
	2261, 50, -68, -68, -1000, -1000, -1000, 48, 47, -68,
	46, 570, -1000, 45, -1000, -1000, -1000, 34, -1000, -68,
	-1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 10, 185, 146, 184, 8, 4, 11, 178, 177,
	5, 0, 6, 7, 1, 176, 2, 87, 196, 145,
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
	11, 11, 17, 17, 18, 18, 19, 19,
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
	5, 4, 4, 6, 6, 2, 2, 5, 4, 6,
	3, 2, 0, 1, 1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -17, -2, -18, -19, 81, 82, -3, 11,
	-11, -13, 37, 38, 10, 12, 27, -4, 15, 56,
	28, 44, 4, 5, 65, 73, 74, 75, 66, 6,
	24, 25, 26, 52, 57, 9, 78, 70, 77, 23,
	47, 49, 51, 50, -12, 13, -17, -18, -19, -16,
	4, 58, 72, 64, 65, 66, 67, 68, 41, 42,
	43, 17, 18, 62, 19, 63, 20, 31, 32, 33,
	34, 35, 36, 39, 40, 80, 21, 75, 22, 77,
	78, 50, 58, -12, -11, -11, 4, 53, 4, -11,
	-1, -11, 60, 77, 78, -11, -11, -11, 4, -11,
	4, -11, -11, 77, 4, -17, -17, -11, 77, 4,
	-11, -10, 48, 4, 77, -11, 61, -11, -3, 58,
	61, -11, 4, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -12, -11, -11,
	-13, -11, 61, 70, 4, 58, 70, 29, 60, -7,
	-17, -12, -11, 72, 72, 61, -16, 4, 77, -12,
	-15, -14, 6, 76, -10, 77, 77, 72, 77, -11,
	-17, 70, -13, -17, 8, 76, 79, 60, -17, -1,
	16, -11, -1, -1, -7, 71, -6, -5, 45, 46,
	-6, -5, 8, 76, 79, 60, 4, 4, -11, 76,
	8, -16, 4, 61, -17, 61, -17, 60, 76, -12,
	-12, 4, -11, 76, 61, -11, 4, -1, 4, 76,
	-11, -11, 4, 71, -11, 54, 71, 71, 71, -11,
	60, 76, -11, 61, -17, 76, 76, 8, -17, 79,
	-17, 71, -11, 8, 76, 8, 76, 76, -11, -9,
	79, 70, -11, 60, -17, 79, -11, -1, -17, -17,
	76, 79, -14, 71, 76, 76, 76, -8, 14, 71,
	55, -1, 70, -17, -1, 76, 71, -1, -1, -17,
	-1, -11, 71, -1, -1, 71, 71, -1, 71, 70,
	71, 71, -1,
}
var yyDef = [...]int{

	-2, -2, -2, 122, 123, 124, 126, 127, 4, 39,
	-2, 0, 10, 11, 47, 0, 0, 15, 0, 0,
	-2, 0, 51, 52, 0, 0, 0, 0, 0, 60,
	61, 62, 63, 64, 0, 0, 122, 122, 0, 0,
	0, 0, 0, 0, 0, 0, 2, -2, 125, 7,
	40, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 97, 98, 0, 0, 0, 0, 47,
	0, 0, 47, 12, 48, 13, 0, 0, 0, 0,
	0, 0, 28, 47, 0, 53, 54, 55, -2, 0,
	-2, 0, 0, 39, 0, 47, 36, 0, 0, 51,
	0, 115, 116, 45, 0, 121, 122, 0, 5, 47,
	122, 8, 66, 77, 78, 79, 80, 81, 82, 83,
	84, -2, -2, 87, 88, 89, 90, 91, 92, 93,
	94, 95, 96, 99, 100, 101, 102, 0, 0, 120,
	9, -2, 122, -2, 0, 0, -2, -2, 28, 0,
	0, 0, 0, 0, 0, 0, 0, 40, 39, 122,
	122, 37, 0, 75, 0, 47, 47, 0, 0, 0,
	0, -2, 6, 0, 0, 108, 112, 0, 0, 0,
	0, 0, 0, 0, 0, 21, 31, 32, 0, 0,
	29, 30, 0, 104, 111, 0, 57, 59, 0, 122,
	0, 0, 40, 122, 0, 122, 0, 0, 76, 0,
	0, 46, 0, 118, 0, -2, -2, 23, 41, 107,
	0, 49, -2, 14, 0, 0, 18, 19, 20, 0,
	122, 103, 0, 0, -2, 122, 122, 0, 0, 71,
	0, 73, 35, 0, -2, 0, -2, 117, 0, 0,
	114, -2, 0, 122, -2, 113, 0, 0, -2, -2,
	122, 72, 38, 74, -2, -2, 119, 24, -2, 27,
	0, 0, -2, -2, 34, 65, 67, 0, 0, -2,
	0, 0, 16, 0, 33, 68, 69, 0, 26, -2,
	17, 70, 25,
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
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:378
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:383
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:388
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
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
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:646
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:651
		{
			yyVAL.expr = &ast.MakeExpr{Type: yyDollar[2].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:656
		{
			yyVAL.expr = &ast.MakeChanExpr{SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 117:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:661
		{
			yyVAL.expr = &ast.MakeChanExpr{SizeExpr: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 118:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:666
		{
			yyVAL.expr = &ast.MakeArrayExpr{LenExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 119:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:671
		{
			yyVAL.expr = &ast.MakeArrayExpr{LenExpr: yyDollar[3].expr, CapExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:676
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:681
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:692
		{
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:695
		{
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:700
		{
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:703
		{
		}
	}
	goto yystack /* stack new state and value */
}
