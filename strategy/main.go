package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

const (
	GUU = iota
	CHO
	PAA
)

var hands []*hand

func init() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatal("invalid args")
	}
	i, _ := strconv.ParseInt(flag.Arg(0), 10, 64)
	rand.Seed(i)

	hands = []*hand{
		&hand{GUU},
		&hand{CHO},
		&hand{PAA},
	}
}

type hand struct {
	handValue int
}

func getHand(handValue int) *hand {
	return hands[handValue]
}

func (hand *hand) isStrongerThan(h *hand) bool {
	return hand.fight(h) == 1
}

func (hand *hand) isWeakerThan(h *hand) bool {
	return hand.fight(h) == -1
}

func (hand *hand) fight(h *hand) int {
	if hand == h {
		return 0
	} else if (hand.handValue+1)%3 == h.handValue {
		return 1
	} else {
		return -1
	}
}

// Strategy インターフェイス (抽象クラス)
type Strategy interface {
	NextHand() *hand
	study(win bool)
}

type WinningStrategy struct {
	won      bool
	prevHand *hand
}

func NewWinningStrategy() *WinningStrategy {
	return &WinningStrategy{
		won: false,
	}
}

func (ws *WinningStrategy) NextHand() *hand {
	if !ws.won {
		ws.prevHand = getHand(rand.Intn(3))
	}
	return ws.prevHand
}

func (ws *WinningStrategy) study(win bool) {
	ws.won = win
}

// Playerクラス ストラテジーデザインパターンおけるContextの役割
type Player struct {
	name      string
	strategy  Strategy
	wincount  int
	losecount int
	gamecount int
}

func NewPlayer(name string, strategy Strategy) *Player {
	return &Player{
		name:     name,
		strategy: strategy,
	}
}

func (p *Player) nextHand() *hand {
	return p.strategy.NextHand()
}

func (p *Player) win() {
	p.strategy.study(true)
	p.wincount++
	p.gamecount++
}

func (p *Player) lose() {
	p.strategy.study(false)
	p.losecount++
	p.gamecount++
}

func (p *Player) even() {
	p.gamecount++
}

func (p *Player) toString() string {
	return "[" + p.name + ":" + strconv.Itoa(p.gamecount) + " games, " + strconv.Itoa(p.wincount) + " win, " + strconv.Itoa(p.losecount) + " lose " + "]"
}

func main() {
	player1 := NewPlayer("渡辺", NewWinningStrategy())
	player2 := NewPlayer("竣介", NewWinningStrategy())

	for i := 0; i < 10000; i++ {
		nextHand1 := player1.nextHand()
		nextHand2 := player2.nextHand()
		if nextHand1.isStrongerThan(nextHand2) {
			fmt.Println("Winner: " + player1.name)
			player1.win()
			player2.lose()
		} else if nextHand2.isStrongerThan(nextHand1) {
			fmt.Println("Winner: " + player2.name)
			player2.win()
			player1.lose()
		} else {
			fmt.Println("Even...")
			player1.even()
			player2.even()
		}
	}

	fmt.Println("Total Result:")
	fmt.Println(player1.toString())
	fmt.Println(player2.toString())
}
