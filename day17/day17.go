package day17

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	InputPath = "day17/input.txt"
)

func Runner() error {
	cpu, program, err := load(InputPath)
	if err != nil {
		return err
	}
	fmt.Printf("\n:: DAY 17 ::\n")
	//RunWithInit(cpu, program, 5)
	fmt.Printf("Part 1: %v\n", Part1(cpu, program))
	cpu.Clear()
	fmt.Printf("Part 2: %v\n", Part2(cpu, program))
	return nil
}

func Part1(cpu *CPU, program string) string {
	program = strings.Replace(program, ",", "", -1)
	return cpu.Run(program)
}

func Part2(cpu *CPU, program string) int {
	cpu.A = 0
	program = strings.Replace(program, ",", "", -1)
	return FindSelf(cpu, program[:len(program)-2], program, 0)
}

func RunWithInit(cpu *CPU, program string, initA int) string {
	cpu.Clear()
	cpu.A = initA
	cpu.Run(program)
	out := cpu.Output.String()
	return out
}

// I really need to study this a bit more to understand it. I marginally do, enough to feel comfortable putting it onto my own github. But I don't have it internalized.
func FindSelf(cpu *CPU, program string, targetProgram string, initA int) int {
	if len(targetProgram) == 0 {
		return initA
	}
	for i := 0; i <= 7; i++ {
		cpu.Clear()
		// Every new A, we move through possible inputs 0 - 7 to see if it generates our next needed value
		// A little wonky, since part 1 wipes out A.
		// We try the next three bits as the determinant of our output.
		tryA := (initA << 3) + i
		cpu.A = tryA
		cpu.Run(program)
		targetOut := string(targetProgram[len(targetProgram)-1])
		if cpu.Output.String() == targetOut {
			if v := FindSelf(cpu, program, targetProgram[:len(targetProgram)-1], tryA); v != -1 {
				return v
			}
		}
	}
	return -1 // Not a candidate value
}

// Sure, I don't need to write this as methods on a type, but it feels nice. We'll see.
type CPU struct {
	A, B, C  int // registers
	IP       int // instruction pointer
	IPJumped bool
	Output   strings.Builder
}

func NewCPU(a, b, c int) *CPU {
	return &CPU{
		A: a,
		B: b,
		C: c,
	}
}

func (c *CPU) Clear() {
	c.A = 0
	c.B = 0
	c.C = 0
	c.Output = strings.Builder{}
	c.IP = 0
	c.IPJumped = false
}

func (c *CPU) Run(program string) string {
	for {
		if c.IP >= len(program) {
			break
		}
		op := int(program[c.IP] - '0')
		operand := int(program[c.IP+1] - '0')
		c.DoOp(op, operand)
		if !c.IPJumped {
			c.IP += 2
		}
		c.IPJumped = false
	}
	return c.Output.String()
}

func (c *CPU) DoOp(op int, operand int) {
	switch op {
	case 0:
		c.Adv(operand)
	case 1:
		c.Bxl(operand)
	case 2:
		c.Bst(operand)
	case 3:
		c.Jnz(operand)
	case 4:
		c.Bxc(operand)
	case 5:
		c.Out(operand)
	case 6:
		c.Bdv(operand)
	case 7:
		c.Cdv(operand)
	}
}

func (c *CPU) Adv(operand int) {
	c.A = c.A / pow(2, c.ComboOp(operand))
}

func (c *CPU) Bxl(operand int) {
	c.B = c.B ^ operand
}

func (c *CPU) Bst(operand int) {
	c.B = c.ComboOp(operand) % 8
}

func (c *CPU) Jnz(operand int) {
	if c.A != 0 {
		c.IP = operand
		c.IPJumped = true
	}
}

func (c *CPU) Bxc(operand int) {
	c.B = c.B ^ c.C
}

func (c *CPU) Out(operand int) {
	if c.Output.Len() > 0 {
		c.Output.WriteRune(',')
	}
	c.Output.WriteString(strconv.Itoa(c.ComboOp(operand) % 8))
}

func (c *CPU) Bdv(operand int) {
	c.B = c.A / pow(2, c.ComboOp(operand))
}

func (c *CPU) Cdv(operand int) {
	c.C = c.A / pow(2, c.ComboOp(operand))
}

func (c *CPU) ComboOp(operand int) int {
	var res int
	switch operand {
	case 0, 1, 2, 3:
		res = operand
	case 4:
		res = c.A
	case 5:
		res = c.B
	case 6:
		res = c.C
	}
	return res
}

func pow(a, b int) int {
	return int(math.Floor(math.Pow(float64(a), float64(b))))
}

var (
	regRegister = regexp.MustCompile("Register \\w: (\\d+)")
	regProgram  = regexp.MustCompile("Program: (.+)")
)

func load(inputPath string) (*CPU, string, error) {
	f, err := os.Open(inputPath)
	if err != nil {
		return nil, "", err
	}
	buf := bufio.NewScanner(f)
	buf.Scan()
	a := regRegister.FindStringSubmatch(buf.Text())[1]
	aI, _ := strconv.Atoi(a)
	buf.Scan()
	b := regRegister.FindStringSubmatch(buf.Text())[1]
	bI, _ := strconv.Atoi(b)
	buf.Scan()
	c := regRegister.FindStringSubmatch(buf.Text())[1]
	cI, _ := strconv.Atoi(c)
	buf.Scan()
	buf.Scan()
	program := regProgram.FindStringSubmatch(buf.Text())[1]
	return NewCPU(aI, bI, cI), program, nil
}
