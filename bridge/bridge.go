package main

import "fmt"

// 機能のクラス階層 ---------------

type Display struct {
	impl DisplayImpl
}

func (d Display) open() {
	d.impl.rawOpen()
}

func (d Display) print() {
	d.impl.rawPrint()
}

func (d Display) close() {
	d.impl.rawClose()
}

type CountDisplay struct {
	*Display
}

func (cd CountDisplay) multiDisplay(times int) {
	cd.impl.rawOpen()
	for i := 0; i < times; i++ {
		cd.impl.rawPrint()
	}
	cd.impl.rawClose()
}

// 実装のクラス階層 ---------------

type DisplayImpl interface {
	rawOpen()
	rawPrint()
	rawClose()
}

type StringDisplay struct {
	Str   string
	Width int
}

func NewStringDisplay(s string) *StringDisplay {
	return &StringDisplay{
		Str:   s,
		Width: len(s),
	}
}

func (sd *StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < sd.Width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func (sd *StringDisplay) rawOpen() {
	sd.printLine()
}

func (sd *StringDisplay) rawPrint() {
	fmt.Println("|" + sd.Str + "|")
}

func (sd *StringDisplay) rawClose() {
	sd.printLine()
}

func main() {
	strDisplayImpl := NewStringDisplay("Hello Shunsuke")

	display := Display{
		impl: strDisplayImpl,
	}

	display.open()
	display.print()
	display.close()

	countDisplay := CountDisplay{
		Display: &Display{
			impl: strDisplayImpl,
		},
	}

	countDisplay.multiDisplay(5)
}
