package main

import "fmt"

type Product interface {
	createClone() Product
	use(s string)
}

type Manager struct {
	showcase map[string]Product
}

func (m *Manager) register(name string, proto Product) {
	m.showcase[name] = proto
}

func (m *Manager) create(protoname string) Product {
	p := m.showcase[protoname]
	return p.createClone()
}

// MessageBox構造体はProduct インターフェイスを実装する
type MessageBox struct {
	decochar string
}

func (mb *MessageBox) use(s string) {
	length := len(s)
	for i := 0; i < length+4; i++ {
		fmt.Print(mb.decochar)
	}
	fmt.Println("")
	fmt.Println(mb.decochar + " " + s + " " + mb.decochar)
	for i := 0; i < length+4; i++ {
		fmt.Print(mb.decochar)
	}
	fmt.Println("")
}

func (mb *MessageBox) createClone() Product {
	return &MessageBox{
		decochar: mb.decochar,
	}
}

// UnderlinePen構造体はProductインターフェイスを実装する
type UnderlinePen struct {
	ulchar string
}

func (up *UnderlinePen) use(s string) {
	length := len(s)
	fmt.Println("\"" + s + "\"")
	for i := 0; i < length; i++ {
		fmt.Print(up.ulchar)
	}
	fmt.Println("")
}

func (up *UnderlinePen) createClone() Product {
	return &UnderlinePen{
		ulchar: up.ulchar,
	}
}

func main() {
	manager := Manager{showcase: make(map[string]Product)}
	upen := UnderlinePen{ulchar: "~"}
	mbox := MessageBox{decochar: "*"}
	sbox := MessageBox{decochar: "/"}

	manager.register("strong message", &upen)
	manager.register("warning box", &mbox)
	manager.register("slash box", &sbox)

	// 生成
	p1 := manager.create("strong message")
	p1.use("Hello, World")

	p2 := manager.create("warning box")
	p2.use("Hello, World")

	p3 := manager.create("slash box")
	p3.use("Hello, World")
}
