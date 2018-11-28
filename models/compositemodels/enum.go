//go:generate go-enum -f=enum.go
package compositemodels
//x ENUM(
// ADDITION,
// SUBTRACTION,
// MULTIPLICATION,
// DIVISION,
// RANDOM
// )
type Op int

/*const(
	ADDITION Op = 1 + iota
	SUBTRACTION
	MULTIPLICATION
	DIVISION
	RANDOM
)*/