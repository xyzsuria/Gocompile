// code/code.go

package code

// 字节的集合
type Instructions []byte

// 操作码
type Opcode byte

const (
	OpConstant Opcode = iota
)
