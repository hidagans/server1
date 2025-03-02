package ultrasevens

import (
	"context"
	"fmt"
	"time"

	"github.com/slotopol/server/game/slot"
)

func CalcStat(ctx context.Context, mrtp float64) float64 {
	var reels, _ = slot.FindClosest(ReelsMap, mrtp)
	var g = NewGame()
	var sln float64 = 1
	g.Sel = int(sln)
	var s slot.Stat

	var dur = slot.ScanReels5x(ctx, &s, g, reels,
		time.Tick(5*time.Second), time.Tick(2*time.Second))

	var reshuf = float64(s.Reshuffles)
	var lrtp, srtp = s.LinePay / reshuf / sln * 100, s.ScatPay / reshuf / sln * 100
	if srtp > 0 {
		panic("scatters have no pays")
	}
	var rtpsym = lrtp + srtp
	fmt.Printf("completed %.5g%%, selected %d lines, time spent %v\n", reshuf/float64(s.Planned())*100, g.Sel, dur)
	fmt.Printf("reels lengths [%d, %d, %d, %d, %d], total reshuffles %d\n",
		len(reels.Reel(1)), len(reels.Reel(2)), len(reels.Reel(3)), len(reels.Reel(4)), len(reels.Reel(5)), reels.Reshuffles())
	fmt.Printf("jackpots1: count %d, frequency 1/%.12g\n", s.JackCount[ssj1], reshuf/float64(s.JackCount[ssj1]))
	fmt.Printf("jackpots2: count %d, frequency 1/%.12g\n", s.JackCount[ssj2], reshuf/float64(s.JackCount[ssj2]))
	fmt.Printf("jackpots3: count %d, frequency 1/%.12g\n", s.JackCount[ssj3], reshuf/float64(s.JackCount[ssj3]))
	fmt.Printf("RTP = %.6f%%\n", rtpsym)
	return rtpsym
}
