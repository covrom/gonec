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
const THROW = 57353
const IF = 57354
const ELSE = 57355
const FOR = 57356
const IN = 57357
const EQEQ = 57358
const NEQ = 57359
const GE = 57360
const LE = 57361
const OROR = 57362
const ANDAND = 57363
const TRUE = 57364
const FALSE = 57365
const NIL = 57366
const MODULE = 57367
const TRY = 57368
const CATCH = 57369
const FINALLY = 57370
const PLUSEQ = 57371
const MINUSEQ = 57372
const MULEQ = 57373
const DIVEQ = 57374
const ANDEQ = 57375
const OREQ = 57376
const BREAK = 57377
const CONTINUE = 57378
const PLUSPLUS = 57379
const MINUSMINUS = 57380
const POW = 57381
const SHIFTLEFT = 57382
const SHIFTRIGHT = 57383
const SWITCH = 57384
const CASE = 57385
const DEFAULT = 57386
const GO = 57387
const CHAN = 57388
const MAKE = 57389
const OPCHAN = 57390
const ARRAYLIT = 57391
const NULL = 57392
const EACH = 57393
const TO = 57394
const ELSIF = 57395
const WHILE = 57396
const TERNARY = 57397
const TYPECAST = 57398
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

//line ./parser/parser.y:757

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 6,
	1, 7,
	25, 7,
	-2, 131,
	-1, 12,
	60, 50,
	-2, 5,
	-1, 16,
	60, 51,
	-2, 25,
	-1, 25,
	27, 7,
	-2, 131,
	-1, 52,
	60, 50,
	-2, 132,
	-1, 101,
	1, 59,
	8, 59,
	13, 59,
	25, 59,
	27, 59,
	43, 59,
	44, 59,
	52, 59,
	53, 59,
	57, 59,
	59, 59,
	60, 59,
	69, 59,
	70, 59,
	75, 59,
	78, 59,
	80, 59,
	81, 59,
	-2, 54,
	-1, 103,
	1, 61,
	8, 61,
	13, 61,
	25, 61,
	27, 61,
	43, 61,
	44, 61,
	52, 61,
	53, 61,
	57, 61,
	59, 61,
	60, 61,
	69, 61,
	70, 61,
	75, 61,
	78, 61,
	80, 61,
	81, 61,
	-2, 54,
	-1, 133,
	16, 0,
	17, 0,
	-2, 87,
	-1, 134,
	16, 0,
	17, 0,
	-2, 88,
	-1, 154,
	60, 51,
	-2, 45,
	-1, 160,
	70, 7,
	-2, 131,
	-1, 161,
	70, 7,
	-2, 131,
	-1, 187,
	13, 7,
	53, 7,
	70, 7,
	-2, 131,
	-1, 234,
	16, 0,
	60, 52,
	-2, 46,
	-1, 235,
	1, 47,
	13, 47,
	16, 47,
	25, 47,
	27, 47,
	43, 47,
	44, 47,
	53, 47,
	57, 47,
	60, 53,
	70, 47,
	80, 47,
	81, 47,
	-2, 54,
	-1, 242,
	1, 53,
	8, 53,
	13, 53,
	25, 53,
	27, 53,
	43, 53,
	44, 53,
	53, 53,
	60, 53,
	70, 53,
	75, 53,
	78, 53,
	80, 53,
	81, 53,
	-2, 54,
	-1, 257,
	70, 7,
	-2, 131,
	-1, 267,
	1, 108,
	8, 108,
	13, 108,
	25, 108,
	27, 108,
	43, 108,
	44, 108,
	52, 108,
	53, 108,
	57, 108,
	59, 108,
	60, 108,
	69, 108,
	70, 108,
	75, 108,
	78, 108,
	80, 108,
	81, 108,
	-2, 106,
	-1, 269,
	1, 112,
	8, 112,
	13, 112,
	25, 112,
	27, 112,
	43, 112,
	44, 112,
	52, 112,
	53, 112,
	57, 112,
	59, 112,
	60, 112,
	69, 112,
	70, 112,
	75, 112,
	78, 112,
	80, 112,
	81, 112,
	-2, 110,
	-1, 276,
	70, 7,
	-2, 131,
	-1, 280,
	43, 7,
	44, 7,
	70, 7,
	-2, 131,
	-1, 285,
	70, 7,
	-2, 131,
	-1, 286,
	70, 7,
	-2, 131,
	-1, 291,
	1, 107,
	8, 107,
	13, 107,
	25, 107,
	27, 107,
	43, 107,
	44, 107,
	52, 107,
	53, 107,
	57, 107,
	59, 107,
	60, 107,
	69, 107,
	70, 107,
	75, 107,
	78, 107,
	80, 107,
	81, 107,
	-2, 105,
	-1, 292,
	1, 111,
	8, 111,
	13, 111,
	25, 111,
	27, 111,
	43, 111,
	44, 111,
	52, 111,
	53, 111,
	57, 111,
	59, 111,
	60, 111,
	69, 111,
	70, 111,
	75, 111,
	78, 111,
	80, 111,
	81, 111,
	-2, 109,
	-1, 296,
	70, 7,
	-2, 131,
	-1, 300,
	70, 7,
	-2, 131,
	-1, 301,
	70, 7,
	-2, 131,
	-1, 302,
	43, 7,
	44, 7,
	70, 7,
	-2, 131,
	-1, 308,
	70, 7,
	-2, 131,
	-1, 319,
	13, 7,
	53, 7,
	70, 7,
	-2, 131,
}

const yyPrivate = 57344

const yyLast = 3330

var yyAct = [...]int{

	88, 176, 171, 10, 201, 202, 262, 17, 182, 163,
	8, 9, 181, 16, 173, 222, 49, 185, 220, 96,
	97, 117, 89, 179, 97, 92, 215, 94, 215, 93,
	98, 99, 100, 102, 104, 8, 9, 87, 8, 9,
	105, 259, 107, 216, 110, 112, 292, 291, 116, 119,
	287, 121, 258, 16, 251, 123, 268, 125, 126, 127,
	128, 129, 130, 131, 132, 133, 134, 135, 136, 137,
	138, 139, 140, 141, 142, 143, 144, 266, 237, 145,
	146, 147, 148, 181, 150, 152, 154, 154, 207, 188,
	114, 322, 12, 153, 155, 203, 204, 177, 166, 149,
	70, 71, 72, 73, 74, 75, 51, 321, 156, 296,
	61, 320, 318, 165, 106, 316, 183, 315, 184, 84,
	115, 311, 248, 269, 305, 174, 203, 204, 264, 156,
	247, 246, 156, 120, 108, 109, 58, 59, 60, 159,
	156, 156, 55, 250, 267, 80, 224, 82, 83, 298,
	78, 113, 192, 200, 15, 208, 189, 86, 91, 195,
	196, 290, 203, 204, 197, 198, 297, 7, 211, 205,
	206, 214, 199, 161, 11, 3, 218, 194, 260, 14,
	158, 217, 53, 228, 177, 6, 233, 234, 164, 283,
	227, 236, 238, 52, 241, 243, 225, 226, 85, 118,
	219, 213, 212, 172, 249, 90, 157, 122, 124, 116,
	5, 252, 2, 186, 4, 175, 274, 295, 22, 13,
	53, 1, 0, 0, 0, 265, 0, 0, 0, 0,
	0, 271, 0, 272, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 277, 278, 0, 0, 193,
	0, 0, 0, 0, 0, 164, 282, 0, 0, 0,
	0, 284, 241, 0, 0, 289, 0, 221, 223, 0,
	0, 0, 0, 0, 0, 0, 242, 28, 34, 0,
	299, 40, 0, 0, 303, 0, 0, 0, 0, 306,
	307, 0, 0, 0, 35, 36, 37, 0, 0, 310,
	309, 0, 0, 0, 312, 313, 314, 0, 256, 257,
	0, 0, 317, 261, 0, 263, 0, 44, 0, 45,
	48, 46, 38, 323, 0, 0, 0, 39, 47, 0,
	0, 0, 0, 0, 0, 0, 29, 33, 0, 0,
	0, 42, 0, 280, 30, 31, 32, 0, 43, 41,
	288, 285, 286, 0, 0, 27, 28, 34, 0, 0,
	40, 20, 21, 50, 0, 23, 0, 0, 0, 0,
	0, 0, 302, 35, 36, 37, 0, 25, 0, 0,
	308, 0, 0, 0, 0, 0, 18, 19, 0, 0,
	0, 0, 0, 26, 0, 0, 44, 0, 45, 48,
	46, 38, 0, 0, 0, 24, 39, 47, 0, 0,
	0, 0, 0, 0, 0, 29, 33, 0, 0, 0,
	42, 0, 0, 30, 31, 32, 0, 43, 41, 0,
	0, 8, 9, 64, 65, 67, 69, 79, 81, 0,
	0, 0, 0, 0, 0, 0, 70, 71, 72, 73,
	74, 75, 0, 0, 76, 77, 61, 62, 63, 0,
	0, 0, 0, 0, 0, 84, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 232, 66, 68,
	56, 57, 58, 59, 60, 0, 0, 0, 55, 0,
	0, 80, 231, 82, 83, 0, 78, 64, 65, 67,
	69, 79, 81, 0, 0, 0, 0, 0, 0, 0,
	70, 71, 72, 73, 74, 75, 0, 0, 76, 77,
	61, 62, 63, 0, 0, 0, 0, 0, 0, 84,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 230, 66, 68, 56, 57, 58, 59, 60, 0,
	0, 0, 55, 0, 0, 80, 229, 82, 83, 0,
	78, 64, 65, 67, 69, 79, 81, 0, 0, 0,
	0, 0, 0, 0, 70, 71, 72, 73, 74, 75,
	0, 0, 76, 77, 61, 62, 63, 0, 0, 0,
	0, 0, 0, 84, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 210, 0, 66, 68, 56, 57,
	58, 59, 60, 0, 0, 0, 55, 0, 0, 80,
	0, 82, 83, 209, 78, 64, 65, 67, 69, 79,
	81, 0, 0, 0, 0, 0, 0, 0, 70, 71,
	72, 73, 74, 75, 0, 0, 76, 77, 61, 62,
	63, 0, 0, 0, 0, 0, 0, 84, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 191, 0,
	66, 68, 56, 57, 58, 59, 60, 0, 0, 0,
	55, 0, 0, 80, 0, 82, 83, 190, 78, 64,
	65, 67, 69, 79, 81, 0, 0, 0, 0, 0,
	0, 0, 70, 71, 72, 73, 74, 75, 0, 0,
	76, 77, 61, 62, 63, 0, 0, 0, 0, 0,
	0, 84, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 68, 56, 57, 58, 59,
	60, 0, 319, 0, 55, 0, 0, 80, 0, 82,
	83, 0, 78, 64, 65, 67, 69, 79, 81, 0,
	0, 0, 0, 0, 0, 0, 70, 71, 72, 73,
	74, 75, 0, 0, 76, 77, 61, 62, 63, 0,
	0, 0, 0, 0, 0, 84, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 68,
	56, 57, 58, 59, 60, 0, 0, 0, 55, 0,
	0, 80, 304, 82, 83, 0, 78, 64, 65, 67,
	69, 79, 81, 0, 0, 0, 0, 0, 0, 0,
	70, 71, 72, 73, 74, 75, 0, 0, 76, 77,
	61, 62, 63, 0, 0, 0, 0, 0, 0, 84,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 66, 68, 56, 57, 58, 59, 60, 0,
	301, 0, 55, 0, 0, 80, 0, 82, 83, 0,
	78, 64, 65, 67, 69, 79, 81, 0, 0, 0,
	0, 0, 0, 0, 70, 71, 72, 73, 74, 75,
	0, 0, 76, 77, 61, 62, 63, 0, 0, 0,
	0, 0, 0, 84, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 68, 56, 57,
	58, 59, 60, 0, 300, 0, 55, 0, 0, 80,
	0, 82, 83, 0, 78, 64, 65, 67, 69, 79,
	81, 0, 0, 0, 0, 0, 0, 0, 70, 71,
	72, 73, 74, 75, 0, 0, 76, 77, 61, 62,
	63, 0, 0, 0, 0, 0, 0, 84, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	66, 68, 56, 57, 58, 59, 60, 0, 0, 0,
	55, 0, 0, 80, 294, 82, 83, 0, 78, 64,
	65, 67, 69, 79, 81, 0, 0, 0, 0, 0,
	0, 0, 70, 71, 72, 73, 74, 75, 0, 0,
	76, 77, 61, 62, 63, 0, 0, 0, 0, 0,
	0, 84, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 68, 56, 57, 58, 59,
	60, 0, 0, 0, 55, 0, 0, 80, 293, 82,
	83, 0, 78, 64, 65, 67, 69, 79, 81, 0,
	0, 0, 0, 0, 0, 0, 70, 71, 72, 73,
	74, 75, 0, 0, 76, 77, 61, 62, 63, 0,
	0, 0, 0, 0, 0, 84, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 68,
	56, 57, 58, 59, 60, 0, 0, 0, 55, 0,
	0, 80, 0, 82, 83, 281, 78, 64, 65, 67,
	69, 79, 81, 0, 0, 0, 0, 0, 0, 0,
	70, 71, 72, 73, 74, 75, 0, 0, 76, 77,
	61, 62, 63, 0, 0, 0, 0, 0, 0, 84,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	279, 0, 66, 68, 56, 57, 58, 59, 60, 0,
	0, 0, 55, 0, 0, 80, 0, 82, 83, 0,
	78, 64, 65, 67, 69, 79, 81, 0, 0, 0,
	0, 0, 0, 0, 70, 71, 72, 73, 74, 75,
	0, 0, 76, 77, 61, 62, 63, 0, 0, 0,
	0, 0, 0, 84, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 68, 56, 57,
	58, 59, 60, 0, 276, 0, 55, 0, 0, 80,
	0, 82, 83, 0, 78, 64, 65, 67, 69, 79,
	81, 0, 0, 0, 0, 0, 0, 0, 70, 71,
	72, 73, 74, 75, 0, 0, 76, 77, 61, 62,
	63, 0, 0, 0, 0, 0, 0, 84, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	66, 68, 56, 57, 58, 59, 60, 0, 0, 0,
	55, 0, 0, 80, 0, 82, 83, 275, 78, 64,
	65, 67, 69, 79, 81, 0, 0, 0, 0, 0,
	0, 0, 70, 71, 72, 73, 74, 75, 0, 0,
	76, 77, 61, 62, 63, 0, 0, 0, 0, 0,
	0, 84, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 68, 56, 57, 58, 59,
	60, 0, 0, 0, 55, 0, 0, 80, 273, 82,
	83, 0, 78, 64, 65, 67, 69, 79, 81, 0,
	0, 0, 0, 0, 0, 0, 70, 71, 72, 73,
	74, 75, 0, 0, 76, 77, 61, 62, 63, 0,
	0, 0, 0, 0, 0, 84, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 68,
	56, 57, 58, 59, 60, 0, 0, 0, 55, 0,
	0, 80, 270, 82, 83, 0, 78, 64, 65, 67,
	69, 79, 81, 0, 0, 0, 0, 0, 0, 0,
	70, 71, 72, 73, 74, 75, 0, 0, 76, 77,
	61, 62, 63, 0, 0, 0, 0, 0, 0, 84,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 255, 66, 68, 56, 57, 58, 59, 60, 0,
	0, 0, 55, 0, 0, 80, 0, 82, 83, 0,
	78, 64, 65, 67, 69, 79, 81, 0, 0, 0,
	0, 0, 0, 0, 70, 71, 72, 73, 74, 75,
	0, 0, 76, 77, 61, 62, 63, 0, 0, 0,
	0, 0, 0, 84, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 68, 56, 57,
	58, 59, 60, 0, 0, 0, 55, 0, 0, 80,
	0, 82, 83, 254, 78, 64, 65, 67, 69, 79,
	81, 0, 0, 0, 0, 0, 0, 0, 70, 71,
	72, 73, 74, 75, 0, 0, 76, 77, 61, 62,
	63, 0, 0, 0, 0, 0, 0, 84, 0, 0,
	0, 245, 0, 0, 0, 0, 0, 0, 0, 0,
	66, 68, 56, 57, 58, 59, 60, 0, 0, 0,
	55, 0, 0, 80, 0, 82, 83, 0, 78, 64,
	65, 67, 69, 79, 81, 0, 0, 0, 0, 0,
	0, 0, 70, 71, 72, 73, 74, 75, 0, 0,
	76, 77, 61, 62, 63, 0, 0, 0, 0, 0,
	0, 84, 0, 0, 0, 244, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 68, 56, 57, 58, 59,
	60, 0, 0, 0, 55, 0, 0, 80, 0, 82,
	83, 0, 78, 64, 65, 67, 69, 79, 81, 0,
	0, 0, 0, 0, 0, 0, 70, 71, 72, 73,
	74, 75, 0, 0, 76, 77, 61, 62, 63, 0,
	0, 0, 0, 0, 0, 84, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 68,
	56, 57, 58, 59, 60, 0, 0, 0, 55, 0,
	0, 80, 0, 82, 83, 240, 78, 64, 65, 67,
	69, 79, 81, 0, 0, 0, 0, 0, 0, 0,
	70, 71, 72, 73, 74, 75, 0, 0, 76, 77,
	61, 62, 63, 0, 0, 0, 0, 0, 0, 84,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 66, 68, 56, 57, 58, 59, 60, 0,
	187, 0, 55, 0, 0, 80, 0, 82, 83, 0,
	78, 64, 65, 67, 69, 79, 81, 0, 0, 0,
	0, 0, 0, 0, 70, 71, 72, 73, 74, 75,
	0, 0, 76, 77, 61, 62, 63, 0, 0, 0,
	0, 0, 0, 84, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 68, 56, 57,
	58, 59, 60, 0, 0, 0, 55, 0, 0, 80,
	178, 82, 83, 0, 78, 64, 65, 67, 69, 79,
	81, 0, 0, 0, 0, 0, 0, 0, 70, 71,
	72, 73, 74, 75, 0, 0, 76, 77, 61, 62,
	63, 0, 0, 0, 0, 0, 0, 84, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 170,
	66, 68, 56, 57, 58, 59, 60, 0, 0, 0,
	55, 0, 0, 80, 0, 82, 83, 0, 78, 64,
	65, 67, 69, 79, 81, 0, 0, 0, 0, 0,
	0, 0, 70, 71, 72, 73, 74, 75, 0, 0,
	76, 77, 61, 62, 63, 0, 0, 0, 0, 0,
	0, 84, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 162, 0, 66, 68, 56, 57, 58, 59,
	60, 0, 0, 0, 55, 0, 0, 80, 0, 82,
	83, 0, 78, 64, 65, 67, 69, 79, 81, 0,
	0, 0, 0, 0, 0, 0, 70, 71, 72, 73,
	74, 75, 0, 0, 76, 77, 61, 62, 63, 0,
	0, 0, 0, 0, 0, 84, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 68,
	56, 57, 58, 59, 60, 0, 160, 0, 55, 0,
	0, 80, 0, 82, 83, 0, 78, 64, 65, 67,
	69, 79, 81, 0, 0, 0, 0, 0, 0, 0,
	70, 71, 72, 73, 74, 75, 0, 0, 76, 77,
	61, 62, 63, 0, 0, 0, 0, 0, 0, 84,
	0, 0, 0, 0, 0, 0, 0, 0, 54, 0,
	0, 0, 66, 68, 56, 57, 58, 59, 60, 0,
	0, 0, 55, 0, 0, 80, 0, 82, 83, 0,
	78, 64, 65, 67, 69, 79, 81, 0, 0, 0,
	0, 0, 0, 0, 70, 71, 72, 73, 74, 75,
	0, 0, 76, 77, 61, 62, 63, 0, 0, 0,
	0, 0, 0, 84, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 68, 56, 57,
	58, 59, 60, 0, 0, 0, 55, 0, 0, 80,
	0, 82, 83, 0, 78, 64, 65, 67, 69, 79,
	81, 0, 0, 0, 0, 0, 0, 0, 70, 71,
	72, 73, 74, 75, 0, 0, 76, 77, 61, 62,
	63, 0, 0, 0, 0, 0, 0, 84, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	66, 68, 56, 57, 58, 59, 60, 0, 0, 0,
	55, 0, 0, 80, 0, 180, 83, 0, 78, 64,
	65, 67, 69, 79, 81, 0, 0, 0, 0, 0,
	0, 0, 70, 71, 72, 73, 74, 75, 0, 0,
	76, 77, 61, 62, 63, 0, 0, 0, 0, 0,
	0, 84, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 68, 56, 57, 58, 59,
	60, 0, 0, 0, 169, 0, 0, 80, 0, 82,
	83, 0, 78, 64, 65, 67, 69, 79, 81, 0,
	0, 0, 0, 0, 0, 0, 70, 71, 72, 73,
	74, 75, 0, 0, 76, 77, 61, 62, 63, 0,
	0, 0, 0, 0, 0, 84, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 68,
	56, 57, 58, 59, 60, 0, 0, 0, 168, 0,
	0, 80, 0, 82, 83, 0, 78, 65, 67, 69,
	79, 81, 0, 0, 0, 0, 0, 0, 0, 70,
	71, 72, 73, 74, 75, 0, 0, 76, 77, 61,
	62, 63, 0, 0, 0, 0, 0, 0, 84, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 66, 68, 56, 57, 58, 59, 60, 0, 0,
	0, 55, 0, 0, 80, 0, 82, 83, 0, 78,
	64, 65, 67, 69, 0, 81, 0, 0, 0, 0,
	0, 0, 0, 70, 71, 72, 73, 74, 75, 0,
	0, 76, 77, 61, 62, 63, 0, 0, 0, 0,
	0, 0, 84, 0, 27, 28, 34, 0, 0, 40,
	20, 21, 50, 0, 23, 66, 68, 56, 57, 58,
	59, 60, 35, 36, 37, 55, 25, 0, 80, 0,
	82, 83, 0, 78, 0, 18, 19, 0, 0, 0,
	0, 0, 26, 0, 0, 44, 0, 45, 48, 46,
	38, 0, 0, 0, 24, 39, 47, 0, 0, 0,
	0, 0, 0, 0, 29, 33, 0, 0, 0, 42,
	0, 0, 30, 31, 32, 0, 43, 41, 64, 65,
	67, 69, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 70, 71, 72, 73, 74, 75, 0, 0, 76,
	77, 61, 62, 63, 0, 0, 0, 0, 0, 0,
	84, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 66, 68, 56, 57, 58, 59, 60,
	0, 67, 69, 55, 0, 0, 80, 0, 82, 83,
	0, 78, 70, 71, 72, 73, 74, 75, 0, 0,
	76, 77, 61, 62, 63, 0, 0, 0, 0, 0,
	0, 84, 0, 27, 28, 34, 0, 0, 40, 0,
	0, 0, 0, 0, 66, 68, 56, 57, 58, 59,
	60, 35, 36, 37, 55, 0, 0, 80, 0, 82,
	83, 0, 78, 0, 0, 0, 0, 0, 27, 28,
	34, 0, 0, 40, 44, 0, 45, 48, 46, 38,
	0, 0, 0, 0, 39, 47, 35, 36, 37, 0,
	0, 0, 0, 29, 33, 0, 0, 0, 42, 0,
	0, 30, 31, 32, 0, 43, 41, 253, 0, 44,
	0, 45, 48, 46, 38, 0, 0, 0, 0, 39,
	47, 0, 0, 0, 0, 27, 28, 34, 29, 33,
	40, 0, 0, 42, 0, 0, 30, 31, 32, 0,
	43, 41, 239, 35, 36, 37, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 44, 0, 45, 48,
	46, 38, 0, 0, 0, 0, 39, 47, 0, 0,
	167, 0, 27, 28, 34, 29, 33, 40, 0, 0,
	42, 0, 0, 30, 31, 32, 0, 43, 41, 0,
	35, 36, 37, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 44, 0, 45, 48, 46, 38, 0,
	0, 0, 0, 39, 47, 0, 0, 151, 0, 27,
	28, 34, 29, 33, 40, 0, 0, 42, 0, 0,
	30, 31, 32, 0, 43, 41, 0, 35, 36, 37,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	44, 0, 45, 48, 46, 38, 0, 0, 0, 0,
	39, 47, 0, 0, 95, 0, 27, 28, 34, 29,
	33, 40, 0, 0, 42, 0, 0, 30, 31, 32,
	0, 43, 41, 0, 35, 36, 37, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 242, 28, 34, 0, 0, 40, 44, 0, 45,
	48, 46, 38, 0, 0, 0, 0, 39, 47, 35,
	36, 37, 0, 0, 0, 0, 29, 33, 0, 0,
	0, 42, 0, 0, 30, 31, 32, 0, 43, 41,
	0, 0, 44, 0, 45, 48, 46, 38, 0, 0,
	0, 0, 39, 47, 0, 0, 0, 0, 235, 28,
	34, 29, 33, 40, 0, 0, 42, 0, 0, 30,
	31, 32, 0, 43, 41, 0, 35, 36, 37, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 44,
	0, 45, 48, 46, 38, 0, 0, 0, 0, 39,
	47, 0, 0, 0, 0, 0, 0, 0, 29, 33,
	0, 0, 0, 42, 0, 0, 30, 31, 32, 0,
	43, 41, 70, 71, 72, 73, 74, 75, 0, 0,
	76, 77, 61, 111, 28, 34, 0, 0, 40, 0,
	0, 84, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 35, 36, 37, 0, 0, 56, 57, 58, 59,
	60, 0, 0, 0, 55, 0, 0, 80, 0, 82,
	83, 0, 78, 0, 44, 0, 45, 48, 46, 38,
	0, 0, 0, 0, 39, 47, 0, 0, 0, 0,
	103, 28, 34, 29, 33, 40, 0, 0, 42, 0,
	0, 30, 31, 32, 0, 43, 41, 0, 35, 36,
	37, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 101, 28, 34, 0, 0,
	40, 44, 0, 45, 48, 46, 38, 0, 0, 0,
	0, 39, 47, 35, 36, 37, 0, 0, 0, 0,
	29, 33, 0, 0, 0, 42, 0, 0, 30, 31,
	32, 0, 43, 41, 0, 0, 44, 0, 45, 48,
	46, 38, 0, 0, 0, 0, 39, 47, 0, 0,
	0, 0, 0, 0, 0, 29, 33, 0, 0, 0,
	42, 0, 0, 30, 31, 32, 0, 43, 41, 70,
	71, 72, 73, 74, 75, 0, 0, 0, 0, 61,
	0, 0, 0, 0, 0, 0, 0, 0, 84, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 55, 0, 0, 80, 0, 82, 83, 0, 78,
}
var yyPact = [...]int{

	150, 150, -1000, 206, -1000, -70, -70, -1000, -1000, -1000,
	-1000, -1000, 2510, -70, -70, -1000, 2081, 141, -1000, -1000,
	2932, 2932, -1000, 154, 2932, -70, 2875, -57, -1000, 2932,
	2932, 2932, 3201, 3166, -1000, -1000, -1000, -1000, -1000, 2932,
	38, -70, -70, 2932, 3109, 44, -55, 205, 2932, 73,
	2932, -1000, 351, -1000, 2932, 204, 2932, 2932, 2932, 2932,
	2932, 2932, 2932, 2932, 2932, 2932, 2932, 2932, 2932, 2932,
	2932, 2932, 2932, 2932, 2932, 2932, -1000, -1000, 2932, 2932,
	2932, 2932, 2932, 2818, 2932, 2932, 2932, 72, 2145, 2145,
	202, 123, 2017, 146, 1953, -70, 2932, 2761, 3250, 3250,
	3250, -57, 2337, -57, 2273, 1889, 199, -62, 2932, 178,
	1825, -53, 2209, 12, -68, 2932, -1000, 2932, -59, 2145,
	-70, 1761, -1000, 2145, -1000, 71, 71, 3250, 3250, 3250,
	2145, 3073, 3073, 2623, 2623, 3073, 3073, 3073, 3073, 2145,
	2145, 2145, 2145, 2145, 2145, 2145, 2464, 2145, 2572, 81,
	609, 2932, 2145, -1000, 2145, -1000, -70, 162, 2932, 2932,
	-70, -70, -70, 83, 119, 80, 545, 2932, 198, 197,
	2932, -32, 173, 196, -42, -45, -1000, 87, -1000, 2932,
	2932, 186, 2932, 481, 417, 2932, 3024, -70, 3, -1000,
	-1000, 2704, 1697, 2967, 2932, 1633, 1569, 61, 60, 52,
	-1000, -1000, -1000, 2932, 84, -1000, -1000, -21, -1000, -1000,
	2669, 1505, -1000, -1000, 1441, -70, -70, -23, -34, 170,
	-70, -72, -70, 58, 2932, 69, 48, -1000, 1377, -1000,
	2932, -1000, 2932, 1313, 2400, -57, -1000, -1000, 1249, -1000,
	-1000, 2145, -57, 1185, 2932, 2932, -1000, -1000, -1000, 1121,
	-70, -1000, 1057, -1000, -1000, 2932, 185, -70, -70, -70,
	-25, 272, -1000, 91, -1000, 2145, -28, -1000, -29, -1000,
	-1000, 993, 929, -1000, 96, -1000, -70, 865, 801, -70,
	-70, -1000, 737, -1000, 54, -70, -70, -70, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -70, -1000, 2932, 51,
	-70, -70, -70, -1000, -1000, -1000, 47, 45, -70, 42,
	673, -1000, 41, 37, -1000, -1000, -1000, 21, -1000, -70,
	-1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 3, 221, 212, 219, 154, 218, 5, 4, 9,
	217, 216, 151, 0, 16, 7, 1, 215, 2, 179,
	92, 167,
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
	3, 3, 3, 1, 1, 2, 2, 1, 8, 9,
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

	-1000, -2, -3, 25, -3, 4, -19, -21, 80, 81,
	-1, -21, -20, -4, -19, -5, -13, -15, 35, 36,
	10, 11, -6, 14, 54, 26, 42, 4, 5, 64,
	72, 73, 74, 65, 6, 22, 23, 24, 50, 55,
	9, 77, 69, 76, 45, 47, 49, 56, 48, -14,
	12, -20, -19, -21, 57, 71, 63, 64, 65, 66,
	67, 39, 40, 41, 16, 17, 61, 18, 62, 19,
	29, 30, 31, 32, 33, 34, 37, 38, 79, 20,
	74, 21, 76, 77, 48, 57, 16, -14, -13, -13,
	51, 4, -13, -1, -13, 59, 76, 77, -13, -13,
	-13, 4, -13, 4, -13, -13, 76, 4, -20, -20,
	-13, 4, -13, -12, 46, 76, 4, 76, -12, -13,
	60, -13, -5, -13, 4, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -14,
	-13, 59, -13, -15, -13, -15, 60, 4, 57, 16,
	69, 27, 59, -9, -20, -14, -13, 59, 71, 71,
	60, -18, 4, 76, -14, -17, -16, 6, 75, 76,
	76, 71, 76, -13, -13, 76, -20, 69, 8, 75,
	78, 59, -13, -20, 15, -13, -13, -1, -1, -9,
	70, -8, -7, 43, 44, -8, -7, 8, 75, 78,
	59, -13, 4, 4, -13, 60, 75, 8, -18, 4,
	60, -20, 60, -20, 59, -14, -14, 4, -13, 75,
	60, 75, 60, -13, -13, 4, -1, 75, -13, 78,
	78, -13, 4, -13, 52, 52, 70, 70, 70, -13,
	59, 75, -13, 78, 78, 60, -20, -20, 75, 75,
	8, -20, 78, -20, 70, -13, 8, 75, 8, 75,
	75, -13, -13, 75, -11, 78, 69, -13, -13, 59,
	-20, 78, -13, 4, -1, -20, -20, 75, 78, -16,
	70, 75, 75, 75, 75, -10, 13, 70, 53, -1,
	69, 69, -20, -1, 75, 70, -1, -1, -20, -1,
	-13, 70, -1, -1, -1, 70, 70, -1, 70, 69,
	70, 70, 70, -1,
}
var yyDef = [...]int{

	1, -2, 2, 0, 3, 0, -2, 133, 135, 136,
	4, 133, -2, 131, 132, 8, -2, 0, 13, 14,
	50, 0, 17, 0, 0, -2, 0, 54, 55, 0,
	0, 0, 0, 0, 63, 64, 65, 66, 67, 0,
	0, 131, 131, 0, 0, 0, 0, 0, 0, 0,
	0, 6, -2, 134, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 99, 100, 0, 0,
	0, 0, 50, 0, 0, 50, 50, 15, 51, 16,
	0, 0, 0, 0, 0, 31, 50, 0, 56, 57,
	58, -2, 0, -2, 0, 0, 42, 0, 50, 39,
	0, 54, 0, 121, 122, 0, 48, 0, 0, 130,
	131, 0, 9, 10, 69, 79, 80, 81, 82, 83,
	84, 85, 86, -2, -2, 89, 90, 91, 92, 93,
	94, 95, 96, 97, 98, 101, 102, 103, 104, 0,
	0, 0, 129, 11, -2, 12, 131, 0, 0, 0,
	-2, -2, 31, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 43, 42, 131, 131, 40, 0, 78, 50,
	50, 0, 0, 0, 0, 0, 0, -2, 0, 110,
	114, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	24, 34, 35, 0, 0, 32, 33, 0, 106, 113,
	0, 0, 60, 62, 0, 131, 131, 0, 0, 43,
	131, 0, 131, 0, 0, 0, 0, 49, 0, 127,
	0, 124, 0, 0, -2, -2, 26, 109, 0, 119,
	120, 52, -2, 0, 0, 0, 21, 22, 23, 0,
	131, 105, 0, 116, 117, 0, 0, -2, 131, 131,
	0, 0, 74, 0, 76, 38, 0, -2, 0, -2,
	123, 0, 0, 126, 0, 118, -2, 0, 0, 131,
	-2, 115, 0, 44, 0, -2, -2, 131, 75, 41,
	77, -2, -2, 128, 125, 27, -2, 30, 0, 0,
	-2, -2, -2, 37, 68, 70, 0, 0, -2, 0,
	0, 18, 0, 0, 36, 71, 72, 0, 29, -2,
	19, 20, 73, 28,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	81, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 72, 3, 3, 3, 67, 74, 3,
	76, 75, 65, 63, 60, 64, 71, 66, 3, 3,
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
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:127
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "=", Rhss: []ast.Expr{yyDollar[3].expr}}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:131
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:135
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: &ast.BinOpExpr{Lhss: yyDollar[1].expr_many, Operator: "==", Rhss: yyDollar[3].expr_many}}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:139
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:144
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:149
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:154
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:159
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 18:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./parser/parser.y:164
		{
			yyVAL.stmt = &ast.ForStmt{Var: ast.UniqueNames.Set(yyDollar[3].tok.Lit), Value: yyDollar[5].expr, Stmts: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./parser/parser.y:169
		{
			yyVAL.stmt = &ast.NumForStmt{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), Expr1: yyDollar[4].expr, Expr2: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./parser/parser.y:174
		{
			yyVAL.stmt = &ast.NumForStmt{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), Expr1: yyDollar[4].expr, Expr2: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:179
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:184
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[2].compstmt, Catch: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:189
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:194
		{
			yyVAL.stmt = &ast.SelectStmt{Cases: yyDollar[3].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:199
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 26:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.y:205
		{
			yyVAL.stmt_elsifs = []ast.Stmt{}
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:209
		{
			yyVAL.stmt_elsifs = append(yyDollar[1].stmt_elsifs, yyDollar[2].stmt_elsif)
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:215
		{
			yyVAL.stmt_elsif = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt}
		}
	case 29:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./parser/parser.y:221
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, ElseIf: yyDollar[5].stmt_elsifs, Else: yyDollar[7].compstmt}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 30:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:226
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, ElseIf: yyDollar[5].stmt_elsifs, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 31:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.y:232
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:236
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_case}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:240
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_default}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:244
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_case)
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:248
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
		//line ./parser/parser.y:259
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[5].compstmt}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:265
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyDollar[4].compstmt}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:271
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.y:276
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:280
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:284
		{
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[4].expr_pair)
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.y:289
		{
			yyVAL.expr_idents = []int{}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:293
		{
			yyVAL.expr_idents = []int{ast.UniqueNames.Set(yyDollar[1].tok.Lit)}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:297
		{
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, ast.UniqueNames.Set(yyDollar[4].tok.Lit))
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:303
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 46:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:307
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:311
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[4].tok.Lit)})
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:316
		{
			yyVAL.typ = ast.Type{Name: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:320
		{
			yyVAL.typ = ast.Type{Name: ast.UniqueNames.Set(ast.UniqueNames.Get(yyDollar[1].typ.Name) + "." + yyDollar[3].tok.Lit)}
		}
	case 50:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.y:325
		{
			yyVAL.exprs = nil
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:329
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:333
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:337
		{
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[4].tok.Lit)})
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:343
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:348
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:353
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:358
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:363
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:368
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[2].tok.Lit)}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:373
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: ast.UniqueNames.Set(yyDollar[4].tok.Lit)}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:378
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[2].tok.Lit)}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:383
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: ast.UniqueNames.Set(yyDollar[4].tok.Lit)}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:388
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:393
		{
			yyVAL.expr = &ast.ConstExpr{Value: "истина"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:398
		{
			yyVAL.expr = &ast.ConstExpr{Value: "ложь"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:403
		{
			yyVAL.expr = &ast.ConstExpr{Value: "неопределено"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:408
		{
			yyVAL.expr = &ast.ConstExpr{Value: "null"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 68:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./parser/parser.y:413
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[2].expr, Lhs: yyDollar[4].expr, Rhs: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:418
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: ast.UniqueNames.Set(yyDollar[3].tok.Lit)}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 70:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./parser/parser.y:423
		{
			yyVAL.expr = &ast.FuncExpr{Name: ast.UniqueNames.Set("<анонимная функция>"), Args: yyDollar[3].expr_idents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 71:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./parser/parser.y:428
		{
			yyVAL.expr = &ast.FuncExpr{Name: ast.UniqueNames.Set("<анонимная функция>"), Args: []int{ast.UniqueNames.Set(yyDollar[3].tok.Lit)}, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 72:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./parser/parser.y:433
		{
			yyVAL.expr = &ast.FuncExpr{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), Args: yyDollar[4].expr_idents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 73:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./parser/parser.y:438
		{
			yyVAL.expr = &ast.FuncExpr{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), Args: []int{ast.UniqueNames.Set(yyDollar[4].tok.Lit)}, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 74:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:443
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 75:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:448
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 76:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:453
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
		//line ./parser/parser.y:462
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
		//line ./parser/parser.y:471
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:476
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "+", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:481
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "-", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:486
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "*", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:491
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "/", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:496
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "%", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:501
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "**", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:506
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "<<", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:511
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: ">>", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:516
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "==", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:521
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "!=", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:526
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: ">", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:531
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: ">=", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:536
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "<", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:541
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "<=", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:546
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:551
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:556
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:561
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:566
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:571
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:576
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:581
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:586
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "|", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:591
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "||", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:596
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "&", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:601
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "&&", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 105:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:606
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[1].tok.Lit), SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 106:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:611
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[1].tok.Lit), SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 107:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:616
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 108:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:621
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 109:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:626
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 110:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:631
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 111:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:636
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 112:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:641
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 113:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:646
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 114:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:651
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 115:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:656
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 116:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:661
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Begin: yyDollar[3].expr, End: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 117:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:666
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Begin: nil, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 118:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:671
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 119:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:676
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: nil}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 120:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:681
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: nil, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:686
		{
			yyVAL.expr = &ast.MakeExpr{Type: yyDollar[2].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:691
		{
			yyVAL.expr = &ast.MakeChanExpr{SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 123:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:696
		{
			yyVAL.expr = &ast.MakeChanExpr{SizeExpr: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 124:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:701
		{
			yyVAL.expr = &ast.MakeArrayExpr{LenExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 125:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:706
		{
			yyVAL.expr = &ast.MakeArrayExpr{LenExpr: yyDollar[3].expr, CapExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 126:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:711
		{
			yyVAL.expr = &ast.TypeCast{Type: yyDollar[2].typ.Name, CastExpr: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 127:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:716
		{
			yyVAL.expr = &ast.MakeExpr{TypeExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 128:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:721
		{
			yyVAL.expr = &ast.TypeCast{TypeExpr: yyDollar[3].expr, CastExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:726
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:731
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:742
		{
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:745
		{
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:750
		{
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:753
		{
		}
	}
	goto yystack /* stack new state and value */
}
