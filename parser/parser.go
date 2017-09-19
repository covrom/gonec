//line ./parser/parser.y:1
package parser

import __yyfmt__ "fmt"

//line ./parser/parser.y:3
import (
	"github.com/covrom/gonec/ast"
	"github.com/covrom/gonec/names"
)

//line ./parser/parser.y:30
type yySymType struct {
	yys          int
	compstmt     ast.Stmts
	modules      ast.Stmts
	module       ast.Stmt
	stmt_if      ast.Stmt
	stmt_default ast.Stmt
	stmt_elsif   ast.Stmt
	stmt_elsifs  ast.Stmts
	stmt_case    ast.Stmt
	stmt_cases   ast.Stmts
	stmts        ast.Stmts
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
	"')'",
	"'('",
	"'['",
	"']'",
	"'|'",
	"'&'",
	"';'",
	"'\\n'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line ./parser/parser.y:738

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 6,
	1, 7,
	25, 7,
	-2, 127,
	-1, 12,
	60, 50,
	-2, 5,
	-1, 16,
	60, 51,
	-2, 25,
	-1, 25,
	27, 7,
	-2, 127,
	-1, 50,
	60, 50,
	-2, 128,
	-1, 127,
	16, 0,
	17, 0,
	-2, 83,
	-1, 128,
	16, 0,
	17, 0,
	-2, 84,
	-1, 148,
	60, 51,
	-2, 45,
	-1, 154,
	70, 7,
	-2, 127,
	-1, 155,
	70, 7,
	-2, 127,
	-1, 179,
	13, 7,
	53, 7,
	70, 7,
	-2, 127,
	-1, 224,
	16, 0,
	60, 52,
	-2, 46,
	-1, 225,
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
	-1, 232,
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
	74, 53,
	77, 53,
	80, 53,
	81, 53,
	-2, 54,
	-1, 247,
	70, 7,
	-2, 127,
	-1, 257,
	1, 104,
	8, 104,
	13, 104,
	25, 104,
	27, 104,
	43, 104,
	44, 104,
	52, 104,
	53, 104,
	57, 104,
	59, 104,
	60, 104,
	69, 104,
	70, 104,
	74, 104,
	77, 104,
	80, 104,
	81, 104,
	-2, 102,
	-1, 259,
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
	74, 108,
	77, 108,
	80, 108,
	81, 108,
	-2, 106,
	-1, 266,
	70, 7,
	-2, 127,
	-1, 270,
	43, 7,
	44, 7,
	70, 7,
	-2, 127,
	-1, 275,
	70, 7,
	-2, 127,
	-1, 276,
	70, 7,
	-2, 127,
	-1, 281,
	1, 103,
	8, 103,
	13, 103,
	25, 103,
	27, 103,
	43, 103,
	44, 103,
	52, 103,
	53, 103,
	57, 103,
	59, 103,
	60, 103,
	69, 103,
	70, 103,
	74, 103,
	77, 103,
	80, 103,
	81, 103,
	-2, 101,
	-1, 282,
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
	74, 107,
	77, 107,
	80, 107,
	81, 107,
	-2, 105,
	-1, 286,
	70, 7,
	-2, 127,
	-1, 290,
	70, 7,
	-2, 127,
	-1, 291,
	70, 7,
	-2, 127,
	-1, 292,
	43, 7,
	44, 7,
	70, 7,
	-2, 127,
	-1, 298,
	70, 7,
	-2, 127,
	-1, 309,
	13, 7,
	53, 7,
	70, 7,
	-2, 127,
}

const yyPrivate = 57344

const yyLast = 3055

var yyAct = [...]int{

	86, 168, 163, 10, 157, 193, 194, 8, 9, 94,
	95, 173, 252, 16, 212, 177, 47, 17, 174, 210,
	171, 95, 87, 165, 111, 90, 282, 92, 173, 91,
	96, 97, 98, 312, 8, 9, 101, 85, 99, 8,
	9, 281, 104, 106, 277, 248, 110, 113, 241, 115,
	227, 16, 258, 117, 256, 119, 120, 121, 122, 123,
	124, 125, 126, 127, 128, 129, 130, 131, 132, 133,
	134, 135, 136, 137, 138, 12, 311, 139, 140, 141,
	142, 199, 144, 146, 148, 148, 180, 169, 108, 49,
	68, 69, 70, 71, 72, 73, 160, 143, 74, 75,
	59, 147, 149, 310, 150, 308, 150, 100, 205, 82,
	175, 159, 176, 306, 150, 102, 103, 109, 259, 166,
	257, 305, 249, 301, 54, 55, 56, 57, 58, 205,
	195, 196, 53, 150, 286, 295, 80, 81, 150, 76,
	78, 254, 237, 206, 195, 196, 184, 200, 114, 236,
	240, 280, 181, 187, 188, 214, 153, 238, 189, 190,
	84, 191, 203, 204, 197, 198, 7, 89, 208, 158,
	107, 192, 155, 11, 288, 218, 15, 3, 223, 224,
	186, 51, 250, 226, 228, 207, 231, 233, 215, 216,
	178, 287, 195, 196, 14, 169, 239, 152, 273, 217,
	6, 83, 209, 242, 164, 151, 118, 110, 50, 5,
	2, 167, 4, 264, 88, 255, 112, 51, 285, 22,
	13, 261, 1, 262, 0, 0, 185, 116, 0, 0,
	0, 0, 158, 0, 0, 267, 268, 0, 0, 0,
	0, 0, 211, 213, 0, 0, 272, 0, 0, 0,
	0, 274, 231, 0, 0, 279, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 0, 0, 59, 0, 0,
	289, 0, 0, 0, 293, 0, 82, 0, 0, 296,
	297, 246, 247, 0, 0, 0, 251, 0, 253, 300,
	299, 0, 0, 0, 302, 303, 304, 0, 0, 53,
	0, 0, 307, 80, 81, 0, 76, 78, 0, 0,
	0, 0, 0, 313, 0, 0, 270, 0, 0, 0,
	0, 0, 0, 0, 275, 276, 0, 0, 27, 28,
	32, 0, 0, 38, 20, 21, 48, 0, 23, 0,
	0, 0, 0, 0, 0, 292, 33, 34, 35, 0,
	25, 0, 0, 298, 0, 0, 0, 0, 0, 18,
	19, 0, 0, 0, 0, 0, 26, 0, 0, 42,
	0, 43, 46, 44, 36, 0, 0, 0, 24, 37,
	45, 0, 0, 0, 0, 0, 0, 0, 29, 0,
	0, 0, 0, 40, 0, 0, 30, 31, 0, 41,
	39, 0, 0, 0, 8, 9, 62, 63, 65, 67,
	77, 79, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 82, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	222, 64, 66, 54, 55, 56, 57, 58, 0, 0,
	0, 53, 0, 0, 221, 80, 81, 0, 76, 78,
	62, 63, 65, 67, 77, 79, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 220, 64, 66, 54, 55, 56,
	57, 58, 0, 0, 0, 53, 0, 0, 219, 80,
	81, 0, 76, 78, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 82, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 202, 0, 64,
	66, 54, 55, 56, 57, 58, 0, 0, 0, 53,
	0, 0, 0, 80, 81, 201, 76, 78, 62, 63,
	65, 67, 77, 79, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	82, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 183, 0, 64, 66, 54, 55, 56, 57, 58,
	0, 0, 0, 53, 0, 0, 0, 80, 81, 182,
	76, 78, 62, 63, 65, 67, 77, 79, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 82, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 66, 54,
	55, 56, 57, 58, 0, 309, 0, 53, 0, 0,
	0, 80, 81, 0, 76, 78, 62, 63, 65, 67,
	77, 79, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 82, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 66, 54, 55, 56, 57, 58, 0, 0,
	0, 53, 0, 0, 294, 80, 81, 0, 76, 78,
	62, 63, 65, 67, 77, 79, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 66, 54, 55, 56,
	57, 58, 0, 291, 0, 53, 0, 0, 0, 80,
	81, 0, 76, 78, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 82, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	66, 54, 55, 56, 57, 58, 0, 290, 0, 53,
	0, 0, 0, 80, 81, 0, 76, 78, 62, 63,
	65, 67, 77, 79, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	82, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 66, 54, 55, 56, 57, 58,
	0, 0, 0, 53, 0, 0, 284, 80, 81, 0,
	76, 78, 62, 63, 65, 67, 77, 79, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 82, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 66, 54,
	55, 56, 57, 58, 0, 0, 0, 53, 0, 0,
	283, 80, 81, 0, 76, 78, 62, 63, 65, 67,
	77, 79, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 82, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 66, 54, 55, 56, 57, 58, 0, 0,
	0, 53, 0, 0, 0, 80, 81, 271, 76, 78,
	62, 63, 65, 67, 77, 79, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 269, 0, 64, 66, 54, 55, 56,
	57, 58, 0, 0, 0, 53, 0, 0, 0, 80,
	81, 0, 76, 78, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 82, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	66, 54, 55, 56, 57, 58, 0, 266, 0, 53,
	0, 0, 0, 80, 81, 0, 76, 78, 62, 63,
	65, 67, 77, 79, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	82, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 66, 54, 55, 56, 57, 58,
	0, 0, 0, 53, 0, 0, 0, 80, 81, 265,
	76, 78, 62, 63, 65, 67, 77, 79, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 82, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 66, 54,
	55, 56, 57, 58, 0, 0, 0, 53, 0, 0,
	263, 80, 81, 0, 76, 78, 62, 63, 65, 67,
	77, 79, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 82, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 66, 54, 55, 56, 57, 58, 0, 0,
	0, 53, 0, 0, 260, 80, 81, 0, 76, 78,
	62, 63, 65, 67, 77, 79, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 245, 64, 66, 54, 55, 56,
	57, 58, 0, 0, 0, 53, 0, 0, 0, 80,
	81, 0, 76, 78, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 82, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	66, 54, 55, 56, 57, 58, 0, 0, 0, 53,
	0, 0, 0, 80, 81, 244, 76, 78, 62, 63,
	65, 67, 77, 79, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	82, 0, 0, 0, 235, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 66, 54, 55, 56, 57, 58,
	0, 0, 0, 53, 0, 0, 0, 80, 81, 0,
	76, 78, 62, 63, 65, 67, 77, 79, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 82, 0, 0, 0, 234, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 66, 54,
	55, 56, 57, 58, 0, 0, 0, 53, 0, 0,
	0, 80, 81, 0, 76, 78, 62, 63, 65, 67,
	77, 79, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 82, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 66, 54, 55, 56, 57, 58, 0, 0,
	0, 53, 0, 0, 0, 80, 81, 230, 76, 78,
	62, 63, 65, 67, 77, 79, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 66, 54, 55, 56,
	57, 58, 0, 179, 0, 53, 0, 0, 0, 80,
	81, 0, 76, 78, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 82, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	66, 54, 55, 56, 57, 58, 0, 0, 0, 53,
	0, 0, 170, 80, 81, 0, 76, 78, 62, 63,
	65, 67, 77, 79, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	82, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 162, 64, 66, 54, 55, 56, 57, 58,
	0, 0, 0, 53, 0, 0, 0, 80, 81, 0,
	76, 78, 62, 63, 65, 67, 77, 79, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 82, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 156, 0, 64, 66, 54,
	55, 56, 57, 58, 0, 0, 0, 53, 0, 0,
	0, 80, 81, 0, 76, 78, 62, 63, 65, 67,
	77, 79, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 82, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 66, 54, 55, 56, 57, 58, 0, 154,
	0, 53, 0, 0, 0, 80, 81, 0, 76, 78,
	62, 63, 65, 67, 77, 79, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	0, 52, 0, 0, 0, 64, 66, 54, 55, 56,
	57, 58, 0, 0, 0, 53, 0, 0, 0, 80,
	81, 0, 76, 78, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 82, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	66, 54, 55, 56, 57, 58, 0, 0, 0, 53,
	0, 0, 0, 80, 81, 0, 76, 78, 62, 63,
	65, 67, 77, 79, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	82, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 66, 54, 55, 56, 57, 58,
	0, 0, 0, 53, 0, 0, 0, 172, 81, 0,
	76, 78, 63, 65, 67, 77, 79, 0, 0, 0,
	0, 0, 0, 0, 68, 69, 70, 71, 72, 73,
	0, 0, 74, 75, 59, 60, 61, 0, 0, 0,
	0, 0, 0, 82, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 66, 54, 55,
	56, 57, 58, 0, 0, 0, 53, 0, 0, 0,
	80, 81, 0, 76, 78, 62, 63, 65, 67, 0,
	79, 0, 0, 0, 0, 0, 0, 0, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 60,
	61, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	53, 0, 0, 0, 80, 81, 0, 76, 78, 62,
	63, 65, 67, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 68, 69, 70, 71, 72, 73, 0, 0,
	74, 75, 59, 60, 61, 0, 0, 0, 0, 0,
	0, 82, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 66, 54, 55, 56, 57,
	58, 0, 65, 67, 53, 0, 0, 0, 80, 81,
	0, 76, 78, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 82, 0, 27, 28, 32, 0, 0, 38,
	20, 21, 48, 0, 23, 64, 66, 54, 55, 56,
	57, 58, 33, 34, 35, 53, 25, 0, 0, 80,
	81, 0, 76, 78, 0, 18, 19, 0, 0, 232,
	28, 32, 26, 0, 38, 42, 0, 43, 46, 44,
	36, 0, 0, 0, 24, 37, 45, 33, 34, 35,
	0, 0, 0, 0, 29, 0, 0, 0, 0, 40,
	0, 0, 30, 31, 0, 41, 39, 0, 0, 0,
	42, 0, 43, 46, 44, 36, 0, 0, 0, 0,
	37, 45, 0, 0, 0, 27, 28, 32, 0, 29,
	38, 0, 0, 0, 40, 0, 0, 30, 31, 0,
	41, 39, 278, 33, 34, 35, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	27, 28, 32, 0, 0, 38, 42, 0, 43, 46,
	44, 36, 0, 0, 0, 0, 37, 45, 33, 34,
	35, 0, 0, 0, 0, 29, 0, 0, 0, 0,
	40, 0, 0, 30, 31, 0, 41, 39, 243, 0,
	0, 42, 0, 43, 46, 44, 36, 0, 0, 0,
	0, 37, 45, 0, 0, 0, 27, 28, 32, 0,
	29, 38, 0, 0, 0, 40, 0, 0, 30, 31,
	0, 41, 39, 229, 33, 34, 35, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 42, 0, 43,
	46, 44, 36, 0, 0, 0, 0, 37, 45, 0,
	0, 161, 27, 28, 32, 0, 29, 38, 0, 0,
	0, 40, 0, 0, 30, 31, 0, 41, 39, 0,
	33, 34, 35, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 42, 0, 43, 46, 44, 36, 0,
	0, 0, 0, 37, 45, 0, 0, 145, 27, 28,
	32, 0, 29, 38, 0, 0, 0, 40, 0, 0,
	30, 31, 0, 41, 39, 0, 33, 34, 35, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 42,
	0, 43, 46, 44, 36, 0, 0, 0, 0, 37,
	45, 0, 0, 93, 27, 28, 32, 0, 29, 38,
	0, 0, 0, 40, 0, 0, 30, 31, 0, 41,
	39, 0, 33, 34, 35, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 232,
	28, 32, 0, 0, 38, 42, 0, 43, 46, 44,
	36, 0, 0, 0, 0, 37, 45, 33, 34, 35,
	0, 0, 0, 0, 29, 0, 0, 0, 0, 40,
	0, 0, 30, 31, 0, 41, 39, 0, 0, 0,
	42, 0, 43, 46, 44, 36, 0, 0, 0, 0,
	37, 45, 0, 0, 0, 225, 28, 32, 0, 29,
	38, 0, 0, 0, 40, 0, 0, 30, 31, 0,
	41, 39, 0, 33, 34, 35, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	105, 28, 32, 0, 0, 38, 42, 0, 43, 46,
	44, 36, 0, 0, 0, 0, 37, 45, 33, 34,
	35, 0, 0, 0, 0, 29, 0, 0, 0, 0,
	40, 0, 0, 30, 31, 0, 41, 39, 0, 0,
	0, 42, 0, 43, 46, 44, 36, 0, 0, 0,
	0, 37, 45, 0, 68, 69, 70, 71, 72, 73,
	29, 0, 0, 0, 59, 40, 0, 0, 30, 31,
	0, 41, 39, 82, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	56, 57, 58, 0, 0, 0, 53, 0, 0, 0,
	80, 81, 0, 76, 78,
}
var yyPact = [...]int{

	152, 152, -1000, 205, -1000, -73, -73, -1000, -1000, -1000,
	-1000, -1000, 2470, -73, -73, -1000, 2054, 144, -1000, -1000,
	2820, 2820, -1000, 163, 2820, -73, 2764, -66, -1000, 2820,
	2820, 2820, -1000, -1000, -1000, -1000, -1000, 2820, 32, -73,
	-73, 2820, 2946, 42, -51, 203, 2820, 88, 2820, -1000,
	324, -1000, 2820, 202, 2820, 2820, 2820, 2820, 2820, 2820,
	2820, 2820, 2820, 2820, 2820, 2820, 2820, 2820, 2820, 2820,
	2820, 2820, 2820, 2820, -1000, -1000, 2820, 2820, 2820, 2820,
	2820, 2708, 2820, 2820, 2820, 54, 2118, 2118, 201, 140,
	1990, 145, 1926, -73, 2820, 2652, 228, 228, 228, 1862,
	200, -52, 2820, 189, 1798, -55, 2182, -43, -57, 2820,
	-1000, 2820, -60, 2118, -73, 1734, -1000, 2118, -1000, 2975,
	2975, 228, 228, 228, 2118, 61, 61, 2424, 2424, 61,
	61, 61, 61, 2118, 2118, 2118, 2118, 2118, 2118, 2118,
	2309, 2118, 2373, 78, 582, 2820, 2118, -1000, 2118, -1000,
	-73, 165, 2820, 2820, -73, -73, -73, 101, 149, 73,
	518, 2820, 2820, 69, 177, 198, -41, -46, -1000, 96,
	-1000, 2820, 2820, 195, 2820, 454, 390, 2820, 2911, -73,
	-24, -1000, -1000, 2596, 1670, 2855, 2820, 1606, 1542, 79,
	72, 87, -1000, -1000, -1000, 2820, 91, -1000, -1000, -26,
	-1000, -1000, 2561, 1478, 1414, -73, -73, -29, 48, 174,
	-73, -65, -73, 71, 2820, 46, 44, -1000, 1350, -1000,
	2820, -1000, 2820, 1286, 2245, -66, -1000, -1000, 1222, -1000,
	-1000, 2118, -66, 1158, 2820, 2820, -1000, -1000, -1000, 1094,
	-73, -1000, 1030, -1000, -1000, 2820, 194, -73, -73, -73,
	-30, 2505, -1000, 81, -1000, 2118, -33, -1000, -48, -1000,
	-1000, 966, 902, -1000, 121, -1000, -73, 838, 774, -73,
	-73, -1000, 710, -1000, 65, -73, -73, -73, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -73, -1000, 2820, 53,
	-73, -73, -73, -1000, -1000, -1000, 51, 43, -73, 35,
	646, -1000, 33, 6, -1000, -1000, -1000, -37, -1000, -73,
	-1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 3, 222, 210, 220, 176, 219, 6, 5, 4,
	218, 213, 170, 0, 16, 17, 1, 211, 2, 194,
	75, 166,
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
	13, 13, 13, 13, 13, 13, 13, 20, 20, 19,
	19, 21, 21,
}
var yyR2 = [...]int{

	0, 0, 1, 2, 4, 1, 2, 0, 2, 3,
	3, 3, 3, 1, 1, 2, 2, 1, 8, 9,
	9, 5, 5, 5, 4, 1, 0, 2, 4, 8,
	6, 0, 2, 2, 2, 2, 5, 4, 3, 0,
	1, 4, 0, 1, 4, 1, 4, 4, 1, 3,
	0, 1, 4, 4, 1, 1, 2, 2, 2, 1,
	1, 1, 1, 1, 7, 3, 7, 8, 8, 9,
	5, 6, 5, 6, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 2, 3, 3, 3,
	3, 5, 4, 6, 5, 5, 4, 6, 5, 4,
	4, 6, 5, 5, 6, 5, 5, 2, 2, 5,
	4, 6, 5, 4, 6, 3, 2, 0, 1, 1,
	2, 1, 1,
}
var yyChk = [...]int{

	-1000, -2, -3, 25, -3, 4, -19, -21, 80, 81,
	-1, -21, -20, -4, -19, -5, -13, -15, 35, 36,
	10, 11, -6, 14, 54, 26, 42, 4, 5, 64,
	72, 73, 6, 22, 23, 24, 50, 55, 9, 76,
	69, 75, 45, 47, 49, 56, 48, -14, 12, -20,
	-19, -21, 57, 71, 63, 64, 65, 66, 67, 39,
	40, 41, 16, 17, 61, 18, 62, 19, 29, 30,
	31, 32, 33, 34, 37, 38, 78, 20, 79, 21,
	75, 76, 48, 57, 16, -14, -13, -13, 51, 4,
	-13, -1, -13, 59, 75, 76, -13, -13, -13, -13,
	75, 4, -20, -20, -13, 4, -13, -12, 46, 75,
	4, 75, -12, -13, 60, -13, -5, -13, 4, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -14, -13, 59, -13, -15, -13, -15,
	60, 4, 57, 16, 69, 27, 59, -9, -20, -14,
	-13, 59, 60, -18, 4, 75, -14, -17, -16, 6,
	74, 75, 75, 71, 75, -13, -13, 75, -20, 69,
	8, 74, 77, 59, -13, -20, 15, -13, -13, -1,
	-1, -9, 70, -8, -7, 43, 44, -8, -7, 8,
	74, 77, 59, -13, -13, 60, 74, 8, -18, 4,
	60, -20, 60, -20, 59, -14, -14, 4, -13, 74,
	60, 74, 60, -13, -13, 4, -1, 74, -13, 77,
	77, -13, 4, -13, 52, 52, 70, 70, 70, -13,
	59, 74, -13, 77, 77, 60, -20, -20, 74, 74,
	8, -20, 77, -20, 70, -13, 8, 74, 8, 74,
	74, -13, -13, 74, -11, 77, 69, -13, -13, 59,
	-20, 77, -13, 4, -1, -20, -20, 74, 77, -16,
	70, 74, 74, 74, 74, -10, 13, 70, 53, -1,
	69, 69, -20, -1, 74, 70, -1, -1, -20, -1,
	-13, 70, -1, -1, -1, 70, 70, -1, 70, 69,
	70, 70, 70, -1,
}
var yyDef = [...]int{

	1, -2, 2, 0, 3, 0, -2, 129, 131, 132,
	4, 129, -2, 127, 128, 8, -2, 0, 13, 14,
	50, 0, 17, 0, 0, -2, 0, 54, 55, 0,
	0, 0, 59, 60, 61, 62, 63, 0, 0, 127,
	127, 0, 0, 0, 0, 0, 0, 0, 0, 6,
	-2, 130, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 95, 96, 0, 0, 0, 0,
	50, 0, 0, 50, 50, 15, 51, 16, 0, 0,
	0, 0, 0, 31, 50, 0, 56, 57, 58, 0,
	42, 0, 50, 39, 0, 54, 0, 117, 118, 0,
	48, 0, 0, 126, 127, 0, 9, 10, 65, 75,
	76, 77, 78, 79, 80, 81, 82, -2, -2, 85,
	86, 87, 88, 89, 90, 91, 92, 93, 94, 97,
	98, 99, 100, 0, 0, 0, 125, 11, -2, 12,
	127, 0, 0, 0, -2, -2, 31, 0, 0, 0,
	0, 0, 0, 0, 43, 42, 127, 127, 40, 0,
	74, 50, 50, 0, 0, 0, 0, 0, 0, -2,
	0, 106, 110, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 24, 34, 35, 0, 0, 32, 33, 0,
	102, 109, 0, 0, 0, 127, 127, 0, 0, 43,
	127, 0, 127, 0, 0, 0, 0, 49, 0, 123,
	0, 120, 0, 0, -2, -2, 26, 105, 0, 115,
	116, 52, -2, 0, 0, 0, 21, 22, 23, 0,
	127, 101, 0, 112, 113, 0, 0, -2, 127, 127,
	0, 0, 70, 0, 72, 38, 0, -2, 0, -2,
	119, 0, 0, 122, 0, 114, -2, 0, 0, 127,
	-2, 111, 0, 44, 0, -2, -2, 127, 71, 41,
	73, -2, -2, 124, 121, 27, -2, 30, 0, 0,
	-2, -2, -2, 37, 64, 66, 0, 0, -2, 0,
	0, 18, 0, 0, 36, 67, 68, 0, 29, -2,
	19, 20, 69, 28,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	81, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 72, 3, 3, 3, 67, 79, 3,
	75, 74, 65, 63, 60, 64, 71, 66, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 59, 80,
	62, 57, 61, 58, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 76, 3, 77, 73, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 69, 78, 70,
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
		//line ./parser/parser.y:72
		{
			yyVAL.modules = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.modules
			}
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:79
		{
			yyVAL.modules = ast.Stmts{yyDollar[1].module}
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.modules
			}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:86
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
		//line ./parser/parser.y:97
		{
			yyVAL.module = &ast.ModuleStmt{Name: names.UniqueNames.Set(yyDollar[2].tok.Lit), Stmts: yyDollar[4].compstmt}
			yyVAL.module.SetPosition(yyDollar[1].tok.Position())
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:103
		{
			yyVAL.compstmt = nil
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:107
		{
			yyVAL.compstmt = yyDollar[1].stmts
		}
	case 7:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/parser.y:112
		{
			yyVAL.stmts = nil
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:116
		{
			yyVAL.stmts = ast.Stmts{yyDollar[2].stmt}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:120
		{
			if yyDollar[3].stmt != nil {
				yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[3].stmt)
			}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:128
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "=", Rhss: []ast.Expr{yyDollar[3].expr}}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:132
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:136
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: &ast.BinOpExpr{Lhss: yyDollar[1].expr_many, Operator: "==", Rhss: yyDollar[3].expr_many}}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:140
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:145
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:150
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:155
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:160
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 18:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./parser/parser.y:165
		{
			yyVAL.stmt = &ast.ForStmt{Var: names.UniqueNames.Set(yyDollar[3].tok.Lit), Value: yyDollar[5].expr, Stmts: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./parser/parser.y:170
		{
			yyVAL.stmt = &ast.NumForStmt{Name: names.UniqueNames.Set(yyDollar[2].tok.Lit), Expr1: yyDollar[4].expr, Expr2: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./parser/parser.y:175
		{
			yyVAL.stmt = &ast.NumForStmt{Name: names.UniqueNames.Set(yyDollar[2].tok.Lit), Expr1: yyDollar[4].expr, Expr2: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
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
			yyVAL.stmt_elsifs = ast.Stmts{}
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
			yyVAL.stmt_cases = ast.Stmts{}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:237
		{
			yyVAL.stmt_cases = ast.Stmts{yyDollar[2].stmt_case}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:241
		{
			yyVAL.stmt_cases = ast.Stmts{yyDollar[2].stmt_default}
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
			yyVAL.expr_idents = []int{names.UniqueNames.Set(yyDollar[1].tok.Lit)}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:298
		{
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, names.UniqueNames.Set(yyDollar[4].tok.Lit))
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
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit, Id: names.UniqueNames.Set(yyDollar[4].tok.Lit)})
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:317
		{
			yyVAL.typ = ast.Type{Name: names.UniqueNames.Set(yyDollar[1].tok.Lit)}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:321
		{
			yyVAL.typ = ast.Type{Name: names.UniqueNames.Set(names.UniqueNames.Get(yyDollar[1].typ.Name) + "." + yyDollar[3].tok.Lit)}
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
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit, Id: names.UniqueNames.Set(yyDollar[4].tok.Lit)})
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:344
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: names.UniqueNames.Set(yyDollar[1].tok.Lit)}
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
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:369
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:374
		{
			yyVAL.expr = &ast.ConstExpr{Value: "истина"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:379
		{
			yyVAL.expr = &ast.ConstExpr{Value: "ложь"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:384
		{
			yyVAL.expr = &ast.ConstExpr{Value: "неопределено"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:389
		{
			yyVAL.expr = &ast.ConstExpr{Value: "null"}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 64:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./parser/parser.y:394
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[2].expr, Lhs: yyDollar[4].expr, Rhs: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:399
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: names.UniqueNames.Set(yyDollar[3].tok.Lit)}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 66:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./parser/parser.y:404
		{
			yyVAL.expr = &ast.FuncExpr{Name: names.UniqueNames.Set("<анонимная функция>"), Args: yyDollar[3].expr_idents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 67:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./parser/parser.y:409
		{
			yyVAL.expr = &ast.FuncExpr{Name: names.UniqueNames.Set("<анонимная функция>"), Args: []int{names.UniqueNames.Set(yyDollar[3].tok.Lit)}, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 68:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./parser/parser.y:414
		{
			yyVAL.expr = &ast.FuncExpr{Name: names.UniqueNames.Set(yyDollar[2].tok.Lit), Args: yyDollar[4].expr_idents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./parser/parser.y:419
		{
			yyVAL.expr = &ast.FuncExpr{Name: names.UniqueNames.Set(yyDollar[2].tok.Lit), Args: []int{names.UniqueNames.Set(yyDollar[4].tok.Lit)}, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 70:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:424
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 71:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:429
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 72:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:434
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
		//line ./parser/parser.y:443
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
		//line ./parser/parser.y:452
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:457
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "+", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:462
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "-", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:467
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "*", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:472
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "/", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:477
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "%", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:482
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "**", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:487
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "<<", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:492
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: ">>", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:497
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "==", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:502
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "!=", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:507
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: ">", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:512
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: ">=", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:517
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "<", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:522
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "<=", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:527
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:532
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:537
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:542
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:547
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:552
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:557
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:562
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:567
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "|", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:572
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "||", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:577
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "&", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:582
		{
			yyVAL.expr = &ast.BinOpExpr{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "&&", Rhss: []ast.Expr{yyDollar[3].expr}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:587
		{
			yyVAL.expr = &ast.CallExpr{Name: names.UniqueNames.Set(yyDollar[1].tok.Lit), SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 102:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:592
		{
			yyVAL.expr = &ast.CallExpr{Name: names.UniqueNames.Set(yyDollar[1].tok.Lit), SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 103:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:597
		{
			yyVAL.expr = &ast.CallExpr{Name: names.UniqueNames.Set(yyDollar[2].tok.Lit), SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 104:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:602
		{
			yyVAL.expr = &ast.CallExpr{Name: names.UniqueNames.Set(yyDollar[2].tok.Lit), SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 105:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:607
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 106:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:612
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 107:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:617
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 108:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:622
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 109:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:627
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: names.UniqueNames.Set(yyDollar[1].tok.Lit)}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 110:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:632
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 111:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:637
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: names.UniqueNames.Set(yyDollar[1].tok.Lit)}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 112:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:642
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: names.UniqueNames.Set(yyDollar[1].tok.Lit)}, Begin: yyDollar[3].expr, End: &ast.NoneExpr{}}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 113:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:647
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit, Id: names.UniqueNames.Set(yyDollar[1].tok.Lit)}, Begin: &ast.NoneExpr{}, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 114:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:652
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 115:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:657
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: &ast.NoneExpr{}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 116:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:662
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: &ast.NoneExpr{}, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:667
		{
			yyVAL.expr = &ast.MakeExpr{Type: yyDollar[2].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:672
		{
			yyVAL.expr = &ast.MakeChanExpr{SizeExpr: &ast.NoneExpr{}}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 119:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:677
		{
			yyVAL.expr = &ast.MakeChanExpr{SizeExpr: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 120:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:682
		{
			yyVAL.expr = &ast.MakeArrayExpr{LenExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 121:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:687
		{
			yyVAL.expr = &ast.MakeArrayExpr{LenExpr: yyDollar[3].expr, CapExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 122:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/parser.y:692
		{
			yyVAL.expr = &ast.TypeCast{Type: yyDollar[2].typ.Name, CastExpr: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 123:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/parser.y:697
		{
			yyVAL.expr = &ast.MakeExpr{TypeExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 124:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser/parser.y:702
		{
			yyVAL.expr = &ast.TypeCast{TypeExpr: yyDollar[3].expr, CastExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/parser.y:707
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:712
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:723
		{
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/parser.y:726
		{
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:731
		{
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/parser.y:734
		{
		}
	}
	goto yystack /* stack new state and value */
}
