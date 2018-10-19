package main

import "fmt"

type Banner struct {
	String string
}

func (banner *Banner) showWithParen() {
	fmt.Println("(", banner.String, ")")
}

func (banner *Banner) showWithAster() {
	fmt.Println("*", banner.String, "*")
}

// インターフェイス
// 実装するのはPrintBanner構造体
type Print interface {
	printWeak() string
	printStrong() string
}

type PrintBanner struct {
	banner *Banner
}

func (pb *PrintBanner) printWeak() {
	pb.banner.showWithParen()
}

func (pb *PrintBanner) printStrong() {
	pb.banner.showWithAster()
}

func NewPrintBanner(str string) *PrintBanner {
	return &PrintBanner{
		banner: &Banner{String: str},
	}
}

func main() {
	p := NewPrintBanner("サンプル")

	p.printWeak()
	p.printStrong()
}
