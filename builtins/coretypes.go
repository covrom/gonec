package core

type VMOperation int

const (
	_    VMOperation = iota
	ADD              // +
	SUB              // -
	MUL              // *
	QUO              // /
	REM              // %
	EQL              // ==
	NEQ              // !=
	GTR              // >
	GEQ              // >=
	LSS              // <
	LEQ              // <=
	OR               // |
	LOR              // ||
	AND              // &
	LAND             // &&
	POW              //**
	SHL              // <<
	SHR              // >>
)

var OperMap = map[string]VMOperation{
	"+":  ADD,  // +
	"-":  SUB,  // -
	"*":  MUL,  // *
	"/":  QUO,  // /
	"%":  REM,  // %
	"==": EQL,  // ==
	"!=": NEQ,  // !=
	">":  GTR,  // >
	">=": GEQ,  // >=
	"<":  LSS,  // <
	"<=": LEQ,  // <=
	"|":  OR,   // |
	"||": LOR,  // ||
	"&":  AND,  // &
	"&&": LAND, // &&
	"**": POW,  //**
	"<<": SHL,  // <<
	">>": SHR,  // >>
}

var OperMapR = map[VMOperation]string{
	ADD:  "+",  // +
	SUB:  "-",  // -
	MUL:  "*",  // *
	QUO:  "/",  // /
	REM:  "%",  // %
	EQL:  "==", // ==
	NEQ:  "!=", // !=
	GTR:  ">",  // >
	GEQ:  ">=", // >=
	LSS:  "<",  // <
	LEQ:  "<=", // <=
	OR:   "|",  // |
	LOR:  "||", // ||
	AND:  "&",  // &
	LAND: "&&", // &&
	POW:  "**", //**
	SHL:  "<<", // <<
	SHR:  ">>", // >>
}

// тип NULL

type VMNullType struct{}

func (x VMNullType) vmval()                 {}
func (x VMNullType) null()                  {}
func (x VMNullType) String() string         { return "NULL" }
func (x VMNullType) Interface() interface{} { return x }

var VMNullVar = VMNullType{}

// старые типы

type VMChannel chan interface{}
