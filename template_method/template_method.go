package main

import "fmt"

type printer interface {
	open() string
	print() string
	close() string
}

// 抽象クラス
// Displayメソッドのみを実装。実際の処理はどうなっているのかこの構造体は知らない
type AbstractDisplay struct {
}

func (ad *AbstractDisplay) Display(printer printer) string {
	result := printer.open()
	for i := 0; i < 5; i++ {
		result += printer.print()
	}
	result += printer.close()

	return result
}

// 具象クラス
type CharDisplay struct {
	*AbstractDisplay
	Char string
}

func (cd CharDisplay) open() string {
	return ">>"
}

func (cd CharDisplay) close() string {
	return "<<"
}

func (cd CharDisplay) print() string {
	return cd.Char
}

// 具象クラス
type StrDisplay struct {
	*AbstractDisplay
	Str string
}

func (sd StrDisplay) printLine() string {
	line := "+-"
	for _, _ = range sd.Str {
		line += "-"
	}
	line += "-+\n"
	return line
}

func (sd StrDisplay) open() string {
	return sd.printLine()
}

func (sd StrDisplay) print() string {
	return "| " + sd.Str + " |\n"
}

func (sd StrDisplay) close() string {
	return sd.printLine()
}

func main() {
	abstractDisplay := AbstractDisplay{}

	cd := CharDisplay{
		AbstractDisplay: &abstractDisplay,
		Char:            "具象クラス1",
	}

	sd := StrDisplay{
		AbstractDisplay: &abstractDisplay,
		Str:             "具象クラス2",
	}

	fmt.Println(abstractDisplay.Display(cd))
	fmt.Println(abstractDisplay.Display(sd))
}
