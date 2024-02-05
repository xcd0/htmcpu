package main

import (
	"fmt"
	"strings"
)

// TD4で使用する4bit変数型
// math/bitを使用する。

// TD4のレジスタ
type Register struct {
	carry bool
	a     uint
	b     uint
	pc    uint
	i     uint
	o     uint
}

// 命令群保持ROM用
type Memory [16]uint

/*
adda im | 0000
addb im | 0101
mova im | 0011
movb im | 0111
movab   | 0001
movba   | 0100
jmp im  | 1111
jnc im  | 1110
ina     | 0010
out im  | 1001
outb    | 10001
*/

func main() {

	var reg Register
	var mem Memory

	var program string = `
# コメント

mova 1
movb 1
adda 1
addb 2
`

	Asemble(&program, &mem)
	Reset(&reg)
	Run(&reg, &mem)
}
func Run(reg *Register, mem *Memory) {

	for ; reg.pc < len(mem); reg.pc++ {
		ins := Fetch(mem, reg.pc)

	}

}

func Asemble(program *string, mem *Memory) {

	// 文字列を改行で分割
	lines := strings.Split(text, "\n")

	// 分割した各行に対してループ
	for _, line := range lines {
		fmt.Println(line)
	}
}

func Fetch(mem *Memory, pc uint) uint {
	return (*mem)[pc]
}

func DecodeInstruction(ins uint) {
}

func SetInstruction(mem *Memory) {
	mem[0] = 0b0000_0000 // 直書き?
}

func Reset(reg *Register) {
	reg.carry = false // bool
	reg.a = 0         // uint
	reg.b = 0         // uint
	reg.pc = 0        // uint
	reg.i = 0         // uint
	reg.o = 0         // uint
}
