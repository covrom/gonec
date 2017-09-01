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

//line ./parser/parser.y:737

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 3,
	-2, 128,
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
	28, 3,
	-2, 128,
	-1, 47,
	61, 47,
	-2, 129,
	-1, 98,
	1, 56,
	8, 56,
	14, 56,
	28, 56,
	44, 56,
	45, 56,
	53, 56,
	54, 56,
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
	28, 58,
	44, 58,
	45, 58,
	53, 58,
	54, 58,
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
	-1, 132,
	17, 0,
	18, 0,
	-2, 84,
	-1, 133,
	17, 0,
	18, 0,
	-2, 85,
	-1, 153,
	61, 48,
	-2, 42,
	-1, 155,
	1, 3,
	14, 3,
	28, 3,
	44, 3,
	45, 3,
	54, 3,
	71, 3,
	-2, 128,
	-1, 158,
	71, 3,
	-2, 128,
	-1, 159,
	71, 3,
	-2, 128,
	-1, 185,
	14, 3,
	54, 3,
	71, 3,
	-2, 128,
	-1, 233,
	61, 49,
	-2, 43,
	-1, 234,
	1, 44,
	14, 44,
	28, 44,
	44, 44,
	45, 44,
	54, 44,
	58, 44,
	61, 50,
	71, 44,
	81, 44,
	82, 44,
	-2, 51,
	-1, 242,
	1, 50,
	8, 50,
	14, 50,
	28, 50,
	44, 50,
	45, 50,
	54, 50,
	61, 50,
	71, 50,
	76, 50,
	79, 50,
	81, 50,
	82, 50,
	-2, 51,
	-1, 255,
	71, 3,
	-2, 128,
	-1, 265,
	1, 105,
	8, 105,
	14, 105,
	28, 105,
	44, 105,
	45, 105,
	53, 105,
	54, 105,
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
	-1, 267,
	1, 109,
	8, 109,
	14, 109,
	28, 109,
	44, 109,
	45, 109,
	53, 109,
	54, 109,
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
	-1, 274,
	71, 3,
	-2, 128,
	-1, 277,
	44, 3,
	45, 3,
	71, 3,
	-2, 128,
	-1, 281,
	71, 3,
	-2, 128,
	-1, 282,
	71, 3,
	-2, 128,
	-1, 287,
	1, 104,
	8, 104,
	14, 104,
	28, 104,
	44, 104,
	45, 104,
	53, 104,
	54, 104,
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
	-1, 288,
	1, 108,
	8, 108,
	14, 108,
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
	-1, 292,
	71, 3,
	-2, 128,
	-1, 296,
	71, 3,
	-2, 128,
	-1, 297,
	44, 3,
	45, 3,
	71, 3,
	-2, 128,
	-1, 303,
	71, 3,
	-2, 128,
	-1, 313,
	14, 3,
	54, 3,
	71, 3,
	-2, 128,
}

const yyPrivate = 57344

const yyLast = 3061

var yyAct = [...]int{

	84, 49, 174, 10, 44, 11, 6, 7, 260, 201,
	221, 1, 202, 2, 180, 161, 85, 46, 104, 83,
	89, 266, 91, 93, 94, 95, 96, 97, 99, 101,
	6, 7, 90, 177, 94, 102, 179, 5, 121, 107,
	109, 183, 48, 113, 116, 171, 118, 114, 10, 288,
	105, 106, 122, 257, 124, 125, 126, 127, 128, 129,
	130, 131, 132, 133, 134, 135, 136, 137, 138, 139,
	140, 141, 142, 143, 154, 264, 144, 145, 146, 147,
	207, 149, 151, 153, 148, 48, 111, 219, 152, 267,
	315, 103, 188, 287, 283, 164, 121, 256, 163, 67,
	68, 69, 70, 71, 72, 169, 162, 6, 7, 58,
	172, 215, 250, 181, 237, 182, 112, 179, 81, 314,
	312, 153, 292, 175, 155, 310, 186, 309, 154, 203,
	204, 184, 306, 154, 300, 187, 203, 204, 154, 262,
	249, 52, 246, 265, 77, 154, 79, 80, 208, 75,
	245, 192, 120, 157, 117, 121, 247, 223, 196, 82,
	189, 88, 294, 200, 8, 110, 211, 194, 193, 214,
	197, 198, 205, 217, 162, 206, 199, 203, 204, 293,
	159, 227, 224, 225, 232, 233, 220, 222, 286, 195,
	258, 216, 238, 175, 241, 236, 243, 235, 226, 218,
	213, 212, 170, 156, 248, 123, 113, 86, 115, 87,
	50, 251, 119, 4, 173, 272, 291, 47, 17, 3,
	0, 0, 0, 0, 263, 0, 0, 0, 0, 255,
	269, 0, 270, 259, 0, 261, 0, 67, 68, 69,
	70, 71, 72, 0, 0, 275, 0, 58, 0, 0,
	0, 0, 0, 0, 0, 279, 81, 0, 0, 0,
	241, 0, 0, 277, 285, 0, 0, 280, 0, 0,
	281, 282, 0, 55, 56, 57, 0, 0, 0, 52,
	0, 0, 77, 0, 79, 80, 295, 75, 0, 298,
	297, 0, 0, 301, 302, 305, 0, 303, 0, 0,
	0, 0, 0, 0, 304, 0, 0, 0, 307, 308,
	0, 22, 23, 29, 0, 311, 35, 14, 9, 15,
	45, 0, 18, 0, 0, 316, 0, 0, 0, 0,
	30, 31, 32, 16, 20, 0, 0, 0, 0, 0,
	0, 0, 0, 12, 13, 0, 0, 0, 0, 0,
	21, 0, 0, 39, 0, 40, 43, 41, 33, 0,
	0, 0, 19, 34, 42, 0, 0, 0, 0, 0,
	0, 0, 24, 28, 0, 0, 0, 37, 0, 0,
	25, 26, 27, 0, 38, 36, 0, 0, 6, 7,
	61, 62, 64, 66, 76, 78, 0, 0, 0, 0,
	0, 0, 0, 67, 68, 69, 70, 71, 72, 0,
	0, 73, 74, 58, 59, 60, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 231, 63, 65, 53, 54, 55,
	56, 57, 0, 0, 0, 52, 0, 0, 77, 230,
	79, 80, 0, 75, 61, 62, 64, 66, 76, 78,
	0, 0, 0, 0, 0, 0, 0, 67, 68, 69,
	70, 71, 72, 0, 0, 73, 74, 58, 59, 60,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 229, 63,
	65, 53, 54, 55, 56, 57, 0, 0, 0, 52,
	0, 0, 77, 228, 79, 80, 0, 75, 61, 62,
	64, 66, 76, 78, 0, 0, 0, 0, 0, 0,
	0, 67, 68, 69, 70, 71, 72, 0, 0, 73,
	74, 58, 59, 60, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 210, 0, 63, 65, 53, 54, 55, 56, 57,
	0, 0, 0, 52, 0, 0, 77, 0, 79, 80,
	209, 75, 61, 62, 64, 66, 76, 78, 0, 0,
	0, 0, 0, 0, 0, 67, 68, 69, 70, 71,
	72, 0, 0, 73, 74, 58, 59, 60, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 191, 0, 63, 65, 53,
	54, 55, 56, 57, 0, 0, 0, 52, 0, 0,
	77, 0, 79, 80, 190, 75, 61, 62, 64, 66,
	76, 78, 0, 0, 0, 0, 0, 0, 0, 67,
	68, 69, 70, 71, 72, 0, 0, 73, 74, 58,
	59, 60, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 63, 65, 53, 54, 55, 56, 57, 0, 313,
	0, 52, 0, 0, 77, 0, 79, 80, 0, 75,
	61, 62, 64, 66, 76, 78, 0, 0, 0, 0,
	0, 0, 0, 67, 68, 69, 70, 71, 72, 0,
	0, 73, 74, 58, 59, 60, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 63, 65, 53, 54, 55,
	56, 57, 0, 0, 0, 52, 0, 0, 77, 299,
	79, 80, 0, 75, 61, 62, 64, 66, 76, 78,
	0, 0, 0, 0, 0, 0, 0, 67, 68, 69,
	70, 71, 72, 0, 0, 73, 74, 58, 59, 60,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 63,
	65, 53, 54, 55, 56, 57, 0, 296, 0, 52,
	0, 0, 77, 0, 79, 80, 0, 75, 61, 62,
	64, 66, 76, 78, 0, 0, 0, 0, 0, 0,
	0, 67, 68, 69, 70, 71, 72, 0, 0, 73,
	74, 58, 59, 60, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 63, 65, 53, 54, 55, 56, 57,
	0, 0, 0, 52, 0, 0, 77, 290, 79, 80,
	0, 75, 61, 62, 64, 66, 76, 78, 0, 0,
	0, 0, 0, 0, 0, 67, 68, 69, 70, 71,
	72, 0, 0, 73, 74, 58, 59, 60, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 63, 65, 53,
	54, 55, 56, 57, 0, 0, 0, 52, 0, 0,
	77, 289, 79, 80, 0, 75, 61, 62, 64, 66,
	76, 78, 0, 0, 0, 0, 0, 0, 0, 67,
	68, 69, 70, 71, 72, 0, 0, 73, 74, 58,
	59, 60, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 63, 65, 53, 54, 55, 56, 57, 0, 0,
	0, 52, 0, 0, 77, 0, 79, 80, 278, 75,
	61, 62, 64, 66, 76, 78, 0, 0, 0, 0,
	0, 0, 0, 67, 68, 69, 70, 71, 72, 0,
	0, 73, 74, 58, 59, 60, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 276, 0, 63, 65, 53, 54, 55,
	56, 57, 0, 0, 0, 52, 0, 0, 77, 0,
	79, 80, 0, 75, 61, 62, 64, 66, 76, 78,
	0, 0, 0, 0, 0, 0, 0, 67, 68, 69,
	70, 71, 72, 0, 0, 73, 74, 58, 59, 60,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 63,
	65, 53, 54, 55, 56, 57, 0, 274, 0, 52,
	0, 0, 77, 0, 79, 80, 0, 75, 61, 62,
	64, 66, 76, 78, 0, 0, 0, 0, 0, 0,
	0, 67, 68, 69, 70, 71, 72, 0, 0, 73,
	74, 58, 59, 60, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 63, 65, 53, 54, 55, 56, 57,
	0, 0, 0, 52, 0, 0, 77, 0, 79, 80,
	273, 75, 61, 62, 64, 66, 76, 78, 0, 0,
	0, 0, 0, 0, 0, 67, 68, 69, 70, 71,
	72, 0, 0, 73, 74, 58, 59, 60, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 63, 65, 53,
	54, 55, 56, 57, 0, 0, 0, 52, 0, 0,
	77, 271, 79, 80, 0, 75, 61, 62, 64, 66,
	76, 78, 0, 0, 0, 0, 0, 0, 0, 67,
	68, 69, 70, 71, 72, 0, 0, 73, 74, 58,
	59, 60, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 63, 65, 53, 54, 55, 56, 57, 0, 0,
	0, 52, 0, 0, 77, 268, 79, 80, 0, 75,
	61, 62, 64, 66, 76, 78, 0, 0, 0, 0,
	0, 0, 0, 67, 68, 69, 70, 71, 72, 0,
	0, 73, 74, 58, 59, 60, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 254, 63, 65, 53, 54, 55,
	56, 57, 0, 0, 0, 52, 0, 0, 77, 0,
	79, 80, 0, 75, 61, 62, 64, 66, 76, 78,
	0, 0, 0, 0, 0, 0, 0, 67, 68, 69,
	70, 71, 72, 0, 0, 73, 74, 58, 59, 60,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 63,
	65, 53, 54, 55, 56, 57, 0, 0, 0, 52,
	0, 0, 77, 0, 79, 80, 253, 75, 61, 62,
	64, 66, 76, 78, 0, 0, 0, 0, 0, 0,
	0, 67, 68, 69, 70, 71, 72, 0, 0, 73,
	74, 58, 59, 60, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 244, 0, 0, 0, 0, 0,
	0, 0, 0, 63, 65, 53, 54, 55, 56, 57,
	0, 0, 0, 52, 0, 0, 77, 0, 79, 80,
	0, 75, 61, 62, 64, 66, 76, 78, 0, 0,
	0, 0, 0, 0, 0, 67, 68, 69, 70, 71,
	72, 0, 0, 73, 74, 58, 59, 60, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 63, 65, 53,
	54, 55, 56, 57, 0, 0, 0, 52, 0, 0,
	77, 0, 79, 80, 240, 75, 61, 62, 64, 66,
	76, 78, 0, 0, 0, 0, 0, 0, 0, 67,
	68, 69, 70, 71, 72, 0, 0, 73, 74, 58,
	59, 60, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 63, 65, 53, 54, 55, 56, 57, 0, 185,
	0, 52, 0, 0, 77, 0, 79, 80, 0, 75,
	61, 62, 64, 66, 76, 78, 0, 0, 0, 0,
	0, 0, 0, 67, 68, 69, 70, 71, 72, 0,
	0, 73, 74, 58, 59, 60, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 63, 65, 53, 54, 55,
	56, 57, 0, 0, 0, 52, 0, 0, 77, 176,
	79, 80, 0, 75, 61, 62, 64, 66, 76, 78,
	0, 0, 0, 0, 0, 0, 0, 67, 68, 69,
	70, 71, 72, 0, 0, 73, 74, 58, 59, 60,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 168, 63,
	65, 53, 54, 55, 56, 57, 0, 0, 0, 52,
	0, 0, 77, 0, 79, 80, 0, 75, 61, 62,
	64, 66, 76, 78, 0, 0, 0, 0, 0, 0,
	0, 67, 68, 69, 70, 71, 72, 0, 0, 73,
	74, 58, 59, 60, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 160, 0, 63, 65, 53, 54, 55, 56, 57,
	0, 0, 0, 52, 0, 0, 77, 0, 79, 80,
	0, 75, 61, 62, 64, 66, 76, 78, 0, 0,
	0, 0, 0, 0, 0, 67, 68, 69, 70, 71,
	72, 0, 0, 73, 74, 58, 59, 60, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 63, 65, 53,
	54, 55, 56, 57, 0, 158, 0, 52, 0, 0,
	77, 0, 79, 80, 0, 75, 61, 62, 64, 66,
	76, 78, 0, 0, 0, 0, 0, 0, 0, 67,
	68, 69, 70, 71, 72, 0, 0, 73, 74, 58,
	59, 60, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 0, 0, 0, 0, 51, 0, 0,
	0, 63, 65, 53, 54, 55, 56, 57, 0, 0,
	0, 52, 0, 0, 77, 0, 79, 80, 0, 75,
	22, 23, 29, 0, 0, 35, 14, 9, 15, 45,
	0, 18, 0, 0, 0, 0, 0, 0, 0, 30,
	31, 32, 16, 20, 0, 0, 0, 0, 0, 0,
	0, 0, 12, 13, 0, 0, 0, 0, 0, 21,
	0, 0, 39, 0, 40, 43, 41, 33, 0, 0,
	0, 19, 34, 42, 0, 0, 0, 0, 0, 0,
	0, 24, 28, 0, 0, 0, 37, 0, 0, 25,
	26, 27, 0, 38, 36, 61, 62, 64, 66, 76,
	78, 0, 0, 0, 0, 0, 0, 0, 67, 68,
	69, 70, 71, 72, 0, 0, 73, 74, 58, 59,
	60, 0, 0, 0, 0, 0, 0, 81, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	63, 65, 53, 54, 55, 56, 57, 0, 0, 0,
	52, 0, 0, 77, 0, 79, 80, 0, 75, 61,
	62, 64, 66, 76, 78, 0, 0, 0, 0, 0,
	0, 0, 67, 68, 69, 70, 71, 72, 0, 0,
	73, 74, 58, 59, 60, 0, 0, 0, 0, 0,
	0, 81, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 63, 65, 53, 54, 55, 56,
	57, 0, 0, 0, 52, 0, 0, 77, 0, 178,
	80, 0, 75, 61, 62, 64, 66, 76, 78, 0,
	0, 0, 0, 0, 0, 0, 67, 68, 69, 70,
	71, 72, 0, 0, 73, 74, 58, 59, 60, 0,
	0, 0, 0, 0, 0, 81, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 63, 65,
	53, 54, 55, 56, 57, 0, 0, 0, 167, 0,
	0, 77, 0, 79, 80, 0, 75, 61, 62, 64,
	66, 76, 78, 0, 0, 0, 0, 0, 0, 0,
	67, 68, 69, 70, 71, 72, 0, 0, 73, 74,
	58, 59, 60, 0, 0, 0, 0, 0, 0, 81,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 63, 65, 53, 54, 55, 56, 57, 0,
	0, 0, 166, 0, 0, 77, 0, 79, 80, 0,
	75, 61, 62, 64, 66, 0, 78, 0, 0, 0,
	0, 0, 0, 0, 67, 68, 69, 70, 71, 72,
	0, 0, 73, 74, 58, 59, 60, 0, 0, 0,
	0, 0, 0, 81, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 63, 65, 53, 54,
	55, 56, 57, 0, 0, 0, 52, 0, 0, 77,
	0, 79, 80, 0, 75, 61, 62, 64, 66, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 67, 68,
	69, 70, 71, 72, 0, 0, 73, 74, 58, 59,
	60, 0, 0, 0, 0, 0, 0, 81, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	63, 65, 53, 54, 55, 56, 57, 0, 64, 66,
	52, 0, 0, 77, 0, 79, 80, 0, 75, 67,
	68, 69, 70, 71, 72, 0, 0, 73, 74, 58,
	59, 60, 0, 0, 0, 0, 0, 0, 81, 242,
	23, 29, 0, 0, 35, 0, 0, 0, 0, 0,
	0, 63, 65, 53, 54, 55, 56, 57, 30, 31,
	32, 52, 0, 0, 77, 0, 79, 80, 0, 75,
	0, 0, 0, 0, 22, 23, 29, 0, 0, 35,
	0, 39, 0, 40, 43, 41, 33, 0, 0, 0,
	0, 34, 42, 30, 31, 32, 0, 0, 0, 0,
	24, 28, 0, 0, 0, 37, 0, 0, 25, 26,
	27, 0, 38, 36, 284, 0, 39, 0, 40, 43,
	41, 33, 0, 0, 0, 0, 34, 42, 0, 0,
	0, 0, 22, 23, 29, 24, 28, 35, 0, 0,
	37, 0, 0, 25, 26, 27, 0, 38, 36, 252,
	0, 30, 31, 32, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 22, 23, 29,
	0, 0, 35, 0, 39, 0, 40, 43, 41, 33,
	0, 0, 0, 0, 34, 42, 30, 31, 32, 0,
	0, 0, 0, 24, 28, 0, 0, 0, 37, 0,
	0, 25, 26, 27, 0, 38, 36, 239, 0, 39,
	0, 40, 43, 41, 33, 0, 0, 0, 0, 34,
	42, 0, 0, 165, 0, 22, 23, 29, 24, 28,
	35, 0, 0, 37, 0, 0, 25, 26, 27, 0,
	38, 36, 0, 0, 30, 31, 32, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 22, 23, 29, 0, 39, 35, 40,
	43, 41, 33, 0, 0, 0, 0, 34, 42, 0,
	0, 150, 30, 31, 32, 0, 24, 28, 0, 0,
	0, 37, 0, 0, 25, 26, 27, 0, 38, 36,
	0, 22, 23, 29, 0, 39, 35, 40, 43, 41,
	33, 0, 0, 0, 0, 34, 42, 0, 0, 92,
	30, 31, 32, 0, 24, 28, 0, 0, 0, 37,
	0, 0, 25, 26, 27, 0, 38, 36, 0, 242,
	23, 29, 0, 39, 35, 40, 43, 41, 33, 0,
	0, 0, 0, 34, 42, 0, 0, 0, 30, 31,
	32, 0, 24, 28, 0, 0, 0, 37, 0, 0,
	25, 26, 27, 0, 38, 36, 0, 234, 23, 29,
	0, 39, 35, 40, 43, 41, 33, 0, 0, 0,
	0, 34, 42, 0, 0, 0, 30, 31, 32, 0,
	24, 28, 0, 0, 0, 37, 0, 0, 25, 26,
	27, 0, 38, 36, 0, 0, 0, 0, 0, 39,
	0, 40, 43, 41, 33, 0, 0, 0, 0, 34,
	42, 0, 0, 0, 0, 0, 0, 0, 24, 28,
	0, 0, 0, 37, 0, 0, 25, 26, 27, 0,
	38, 36, 67, 68, 69, 70, 71, 72, 0, 0,
	73, 74, 58, 108, 23, 29, 0, 0, 35, 0,
	0, 81, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 30, 31, 32, 0, 53, 54, 55, 56,
	57, 0, 0, 0, 52, 0, 0, 77, 0, 79,
	80, 0, 75, 0, 0, 39, 0, 40, 43, 41,
	33, 0, 0, 0, 0, 34, 42, 0, 0, 0,
	0, 100, 23, 29, 24, 28, 35, 0, 0, 37,
	0, 0, 25, 26, 27, 0, 38, 36, 0, 0,
	30, 31, 32, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 98, 23, 29, 0,
	0, 35, 0, 39, 0, 40, 43, 41, 33, 0,
	0, 0, 0, 34, 42, 30, 31, 32, 0, 0,
	0, 0, 24, 28, 0, 0, 0, 37, 0, 0,
	25, 26, 27, 0, 38, 36, 0, 0, 39, 0,
	40, 43, 41, 33, 0, 0, 0, 0, 34, 42,
	0, 0, 0, 0, 0, 0, 0, 24, 28, 0,
	0, 0, 37, 0, 0, 25, 26, 27, 0, 38,
	36,
}
var yyPact = [...]int{

	-75, -1000, 1986, -75, -75, -1000, -1000, -1000, -1000, 206,
	1909, 101, -1000, -1000, 2727, 2727, 203, -1000, 157, 2727,
	-75, 2689, -54, -1000, 2727, 2727, 2727, 2982, 2947, -1000,
	-1000, -1000, -1000, -1000, 2727, 14, -75, -75, 2727, 2889,
	39, -30, 202, 2727, 93, 2727, -1000, 307, -1000, 94,
	-1000, 2727, 201, 2727, 2727, 2727, 2727, 2727, 2727, 2727,
	2727, 2727, 2727, 2727, 2727, 2727, 2727, 2727, 2727, 2727,
	2727, 2727, 2727, -1000, -1000, 2727, 2727, 2727, 2727, 2727,
	2651, 2727, 2727, 77, 2048, 2048, -75, 199, 95, 1845,
	152, 1781, -75, 2727, 2593, 69, 69, 69, -54, 2240,
	-54, 2176, 1717, 198, -32, 2727, 187, 1653, -44, 2112,
	45, -63, 2727, -1000, 2727, -36, 2048, -75, 1589, -1000,
	2727, -75, 2048, -1000, 207, 207, 69, 69, 69, 2048,
	2852, 2852, 2419, 2419, 2852, 2852, 2852, 2852, 2048, 2048,
	2048, 2048, 2048, 2048, 2048, 2304, 2048, 2368, 84, 565,
	2727, 2048, -1000, 2048, -75, -75, 173, 2727, -75, -75,
	-75, 92, 133, 72, 501, 2727, 197, 196, 2727, 35,
	183, 195, 26, -51, -1000, 97, -1000, 2727, 2727, 194,
	2727, 437, 373, 2727, 2803, -75, -1000, 191, 38, -1000,
	-1000, 2558, 1525, 2765, -1000, 2727, 1461, 79, 71, 85,
	-1000, -1000, -1000, 2727, 80, -1000, -1000, 36, -1000, -1000,
	2500, 1397, -1000, -1000, 1333, -75, 21, -23, 182, -75,
	-71, -75, 68, 2727, 67, 13, -1000, 1269, -1000, 2727,
	-1000, 2727, 1205, 2048, -54, -1000, -1000, -1000, 1141, -1000,
	-1000, 2048, -54, 1077, 2727, -1000, -1000, -1000, 1013, -75,
	-1000, 949, -1000, -1000, 2727, -75, -75, -75, 18, 2465,
	-1000, 117, -1000, 2048, 17, -1000, -27, -1000, -1000, 885,
	821, -1000, 108, -1000, -75, 757, -75, -75, -1000, 693,
	63, -75, -75, -75, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -75, -1000, 2727, 61, -75, -75, -1000, -1000,
	-1000, 56, 54, -75, 49, 629, -1000, 48, -1000, -1000,
	-1000, 19, -1000, -75, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 11, 219, 164, 218, 12, 9, 15, 216, 215,
	165, 0, 4, 5, 2, 214, 1, 13, 213, 37,
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
	11, 11, 11, 11, 11, 11, 11, 11, 17, 17,
	18, 18, 19, 19,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 2, 3, 4, 2, 3, 3,
	1, 1, 2, 2, 4, 1, 8, 9, 5, 5,
	5, 4, 1, 0, 2, 4, 8, 6, 0, 2,
	2, 2, 2, 5, 4, 3, 0, 1, 4, 0,
	1, 4, 1, 4, 4, 1, 3, 0, 1, 4,
	4, 1, 1, 2, 2, 2, 2, 4, 2, 4,
	1, 1, 1, 1, 1, 7, 3, 7, 8, 8,
	9, 5, 6, 5, 6, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 2, 3, 3,
	3, 3, 5, 4, 6, 5, 5, 4, 6, 5,
	4, 4, 6, 5, 5, 6, 5, 5, 2, 2,
	5, 4, 6, 5, 4, 6, 3, 2, 0, 1,
	1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -17, -2, -18, -19, 81, 82, -3, 11,
	-11, -13, 36, 37, 10, 12, 26, -4, 15, 55,
	27, 43, 4, 5, 65, 73, 74, 75, 66, 6,
	23, 24, 25, 51, 56, 9, 78, 70, 77, 46,
	48, 50, 57, 49, -12, 13, -17, -18, -19, -16,
	4, 58, 72, 64, 65, 66, 67, 68, 40, 41,
	42, 17, 18, 62, 19, 63, 20, 30, 31, 32,
	33, 34, 35, 38, 39, 80, 21, 75, 22, 77,
	78, 49, 58, -12, -11, -11, 4, 52, 4, -11,
	-1, -11, 60, 77, 78, -11, -11, -11, 4, -11,
	4, -11, -11, 77, 4, -17, -17, -11, 4, -11,
	-10, 47, 77, 4, 77, -10, -11, 61, -11, -3,
	58, 61, -11, 4, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -12, -11,
	60, -11, -13, -11, 61, -19, 4, 58, 70, 28,
	60, -7, -17, -12, -11, 60, 72, 72, 61, -16,
	4, 77, -12, -15, -14, 6, 76, 77, 77, 72,
	77, -11, -11, 77, -17, 70, -13, -17, 8, 76,
	79, 60, -11, -17, -1, 16, -11, -1, -1, -7,
	71, -6, -5, 44, 45, -6, -5, 8, 76, 79,
	60, -11, 4, 4, -11, 76, 8, -16, 4, 61,
	-17, 61, -17, 60, -12, -12, 4, -11, 76, 61,
	76, 61, -11, -11, 4, -1, 4, 76, -11, 79,
	79, -11, 4, -11, 53, 71, 71, 71, -11, 60,
	76, -11, 79, 79, 61, -17, 76, 76, 8, -17,
	79, -17, 71, -11, 8, 76, 8, 76, 76, -11,
	-11, 76, -9, 79, 70, -11, 60, -17, 79, -11,
	-1, -17, -17, 76, 79, -14, 71, 76, 76, 76,
	76, -8, 14, 71, 54, -1, 70, -17, -1, 76,
	71, -1, -1, -17, -1, -11, 71, -1, -1, 71,
	71, -1, 71, 70, 71, 71, -1,
}
var yyDef = [...]int{

	-2, -2, -2, 128, 129, 130, 132, 133, 4, 39,
	-2, 0, 10, 11, 47, 0, 0, 15, 0, 0,
	-2, 0, 51, 52, 0, 0, 0, 0, 0, 60,
	61, 62, 63, 64, 0, 0, 128, 128, 0, 0,
	0, 0, 0, 0, 0, 0, 2, -2, 131, 7,
	40, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 96, 97, 0, 0, 0, 0, 47,
	0, 0, 47, 12, 48, 13, 0, 0, 0, 0,
	0, 0, 28, 47, 0, 53, 54, 55, -2, 0,
	-2, 0, 0, 39, 0, 47, 36, 0, 51, 0,
	118, 119, 0, 45, 0, 0, 127, 128, 0, 5,
	47, 128, 8, 66, 76, 77, 78, 79, 80, 81,
	82, 83, -2, -2, 86, 87, 88, 89, 90, 91,
	92, 93, 94, 95, 98, 99, 100, 101, 0, 0,
	0, 126, 9, -2, 128, -2, 0, 0, -2, -2,
	28, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	40, 39, 128, 128, 37, 0, 75, 47, 47, 0,
	0, 0, 0, 0, 0, -2, 6, 0, 0, 107,
	111, 0, 0, 0, 14, 0, 0, 0, 0, 0,
	21, 31, 32, 0, 0, 29, 30, 0, 103, 110,
	0, 0, 57, 59, 0, 128, 0, 0, 40, 128,
	0, 128, 0, 0, 0, 0, 46, 0, 124, 0,
	121, 0, 0, -2, -2, 23, 41, 106, 0, 116,
	117, 49, -2, 0, 0, 18, 19, 20, 0, 128,
	102, 0, 113, 114, 0, -2, 128, 128, 0, 0,
	71, 0, 73, 35, 0, -2, 0, -2, 120, 0,
	0, 123, 0, 115, -2, 0, 128, -2, 112, 0,
	0, -2, -2, 128, 72, 38, 74, -2, -2, 125,
	122, 24, -2, 27, 0, 0, -2, -2, 34, 65,
	67, 0, 0, -2, 0, 0, 16, 0, 33, 68,
	69, 0, 26, -2, 17, 70, 25,
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
		yyDollar = yyS[yypt-4 : yypt+1]
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
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:456
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:461
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:466
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:471
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:476
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:481
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:486
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:491
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:496
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:501
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:506
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:511
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:516
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:521
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:526
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:531
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:536
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:541
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:546
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:551
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:556
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:561
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:566
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:571
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:576
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:581
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:586
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[1].tok.Lit), SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 103:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:591
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[1].tok.Lit), SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 104:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:596
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 105:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:601
		{
			yyVAL.expr = &ast.CallExpr{Name: ast.UniqueNames.Set(yyDollar[2].tok.Lit), SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 106:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:606
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 107:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:611
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 108:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:616
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 109:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:621
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 110:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:626
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 111:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:631
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 112:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:636
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 113:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:641
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Begin: yyDollar[3].expr, End: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 114:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:646
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: ast.UniqueNames.Set(yyDollar[1].tok.Lit)}, Begin: nil, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 115:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:651
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 116:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:656
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: nil}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 117:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:661
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: nil, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:666
		{
			yyVAL.expr = &ast.MakeExpr{Type: yyDollar[2].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 119:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:671
		{
			yyVAL.expr = &ast.MakeChanExpr{SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 120:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:676
		{
			yyVAL.expr = &ast.MakeChanExpr{SizeExpr: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 121:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:681
		{
			yyVAL.expr = &ast.MakeArrayExpr{LenExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 122:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:686
		{
			yyVAL.expr = &ast.MakeArrayExpr{LenExpr: yyDollar[3].expr, CapExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 123:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:691
		{
			yyVAL.expr = &ast.TypeCast{Type: yyDollar[2].typ.Name, CastExpr: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 124:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:696
		{
			yyVAL.expr = &ast.MakeExpr{TypeExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 125:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:701
		{
			yyVAL.expr = &ast.TypeCast{TypeExpr: yyDollar[3].expr, CastExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:706
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:711
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:722
		{
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:725
		{
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:730
		{
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:733
		{
		}
	}
	goto yystack /* stack new state and value */
}
