package main

import (
	"bytes"
	"fmt"
)

type Builder interface {
	makeTitle(title string)
	makeString(str string)
	makeItems(items []string)
	close()
}

// Director　はどの具象builderクラスを受けているか知らない
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) Director {
	return Director{
		builder: builder,
	}
}

func (director *Director) construct() {
	director.builder.makeTitle("Greeting")
	director.builder.makeString("朝から昼にかけて")
	director.builder.makeItems([]string{"おはようございます", "こんにちは"})
	director.builder.makeString("夜に")
	director.builder.makeItems([]string{"こんばんは", "さようなら", "おやすみなさい"})
	director.builder.close()
}

// 具象クラス
type TextBuilder struct {
	stringBuffer bytes.Buffer
}

func (tb *TextBuilder) makeTitle(title string) {
	tb.stringBuffer.WriteString("===========================\n")
	tb.stringBuffer.WriteString("『" + title + "』\n")
	tb.stringBuffer.WriteString("\n")
}

func (tb *TextBuilder) makeString(str string) {
	tb.stringBuffer.WriteString("■" + str + "\n")
	tb.stringBuffer.WriteString("\n")
}

func (tb *TextBuilder) makeItems(items []string) {
	for _, item := range items {
		tb.stringBuffer.WriteString(" ・" + item + "\n")
	}
	tb.stringBuffer.WriteString("\n")
}

func (tb *TextBuilder) close() {
	tb.stringBuffer.WriteString("===========================\n")
}

func (tb *TextBuilder) getResult() string {
	return tb.stringBuffer.String()
}

// 具象クラス
// HTML出力のものもBuilderインターフェイスを実装すればDirectorからconstract関数を用いて実行できる
type HTMLBuilder struct {
}

func main() {
	textBuilder := &TextBuilder{}
	director := Director{builder: textBuilder}
	director.construct()
	result := textBuilder.getResult()
	fmt.Println(result)
}
