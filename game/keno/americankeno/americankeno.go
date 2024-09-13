package americankeno

import (
	keno "github.com/slotopol/server/game/keno"
)

// RTP[ 2] = 86.075949%
// RTP[ 3] = 90.214216%
// RTP[ 4] = 89.036596%
// RTP[ 5] = 88.206730%
// RTP[ 6] = 88.123522%
// RTP[ 7] = 90.974272%
// RTP[ 8] = 89.819236%
// RTP[ 9] = 90.589102%
// RTP[10] = 90.212487%
// RTP[game] = 89.250235%
var Paytable = keno.KenoPaytable{
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},             //  0 sel
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},             //  1 sel
	{1, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0},             //  2 sel
	{1, 0, 2, 15, 0, 0, 0, 0, 0, 0, 0},            //  3 sel
	{1, 0, 1, 5, 50, 0, 0, 0, 0, 0, 0},            //  4 sel
	{1, 0, 0, 3, 12, 400, 0, 0, 0, 0, 0},          //  5 sel
	{1, 0, 0, 2, 6, 50, 1000, 0, 0, 0, 0},         //  6 sel
	{1, 0, 0, 1, 2, 25, 300, 3000, 0, 0, 0},       //  7 sel
	{1, 0, 0, 1, 2, 10, 40, 800, 6000, 0, 0},      //  8 sel
	{1, 0, 0, 0, 1, 5, 45, 400, 2000, 8000, 0},    //  9 sel
	{1, 0, 0, 0, 1, 5, 18, 100, 500, 2500, 10000}, // 10 sel
}

type Game struct {
	keno.Keno80 `yaml:",inline"`
}

func NewGame() *Game {
	return &Game{
		Keno80: keno.Keno80{
			Bet: 1,
		},
	}
}

func (g *Game) Scanner(scrn *keno.Screen, wins *keno.Wins) {
	for i := range 80 {
		if scrn[i] == keno.KSselhit {
			wins.Num++
		}
	}
	wins.Pay = Paytable[len(g.Sel)][wins.Num] * g.Bet
}