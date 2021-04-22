package token

type Token int

const (
	Illegal Token = iota

	operator_beg
	OR  // or
	XOR // xor
	AND // and
	SHL // <<
	SHR // >>

	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // mod
	EXP // pow
	operator_end

	LPAREN // (
	RPAREN // )

	ASSIGN // =

	NUM
	VAR

	AS
	TO
)

var tokenToString = map[Token]string{
	Illegal: "ILLEGAL",

	OR: "OR",
	XOR: "XOR",
	AND: "AND",
	SHL: "<<",
	SHR: ">>",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
	REM: "mod",
	EXP: "pow",

	LPAREN: "(",
	RPAREN: ")",

	ASSIGN: "=",

	NUM: "NUM",
	VAR: "VAR",

	AS: "as",
	TO: "to",
}

func (t Token) String() string {
	return tokenToString[t]
}

const (
	LowestPrec = 0
)

func (t Token) Precedence() int {
	switch t {
	case TO, AS:
		return 1
	case ADD, SUB, OR, XOR:
		return 2
	case MUL, QUO, REM, SHL, SHR, AND, EXP:
		return 3
	}
	return LowestPrec
}

var StringToOperator = map[string]Token{
	"or":  OR,
	"xor": XOR,
	"and": AND,
	"<<":  SHL,
	">>":  SHR,

	"+":   ADD,
	"-":   SUB,
	"*":   MUL,
	"/":   QUO,
	"mod": REM,

	"pow": EXP,
}

func init() {
	for x := operator_beg + 1; x < operator_end; x++ {
		StringToOperator[x.String()] = x
	}
}
