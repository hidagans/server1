package valkyrie

import (
	"context"
	"fmt"
	"time"

	"github.com/slotopol/server/game/slot"
)

func BruteForceBon(ctx context.Context, s slot.Stater, g slot.SlotGame, reels *slot.Reels5x) {
	var r1 = reels.Reel(1)
	var r5 = reels.Reel(3)
	var screen = g.Screen()
	var wins slot.Wins
	for i1 := range r1 {
		screen.SetCol(1, r1, i1)
		for _, big := range BonusReel {
			var x slot.Pos
			for x = 2; x <= 4; x++ {
				screen.Set(x, 1, big)
				screen.Set(x, 2, big)
				screen.Set(x, 3, big)
			}
			for i5 := range r5 {
				screen.SetCol(3, r5, i5)
				g.Scanner(&wins)
				s.Update(wins)
				wins.Reset()
				if s.Count()&100 == 0 {
					select {
					case <-ctx.Done():
						return
					default:
					}
				}
			}
		}
	}
}

func CalcStatBon(ctx context.Context, mrtp float64) float64 {
	var reels, _ = slot.FindClosest(ReelsMap, mrtp)
	var g = NewGame()
	var sln float64 = 1
	g.Sel = int(sln)
	g.FSR = 15 // set free spins mode
	var s slot.Stat

	var ctx2, cancel2 = context.WithCancel(ctx)
	defer cancel2()
	s.SetPlan(uint64(len(reels.Reel(1))) * uint64(len(BonusReel)) * uint64(len(reels.Reel(5))))
	BruteForceBon(ctx2, &s, g, reels)

	var reshuf = float64(s.Reshuffles)
	var lrtp, srtp = s.LinePay / reshuf / sln * 100, s.ScatPay / reshuf / sln * 100
	if srtp > 0 {
		panic("scatters are presents on bonus games")
	}
	var rtp = lrtp + srtp
	fmt.Printf("completed %.5g%%, selected %d lines\n", reshuf/float64(s.Planned())*100, g.Sel)
	fmt.Printf("reels lengths [%d, [%d], %d], total reshuffles %d\n",
		len(reels.Reel(1)), len(BonusReel), len(reels.Reel(5)), reels.Reshuffles())
	fmt.Printf("RTP = %.6f%%\n", rtp)
	return rtp
}

func CalcStatReg(ctx context.Context, mrtp float64) float64 {
	fmt.Printf("*bonus reels calculations*\n")
	var rtpfs = CalcStatBon(ctx, mrtp)
	if ctx.Err() != nil {
		return 0
	}
	fmt.Printf("*regular reels calculations*\n")
	var reels, _ = slot.FindClosest(ReelsMap, mrtp)
	var g = NewGame()
	var sln float64 = 1
	g.Sel = int(sln)
	var s slot.Stat

	var dur = slot.ScanReels5x(ctx, &s, g, reels,
		time.Tick(2*time.Second), time.Tick(2*time.Second))

	var reshuf = float64(s.Reshuffles)
	var lrtp, srtp = s.LinePay / reshuf / sln * 100, s.ScatPay / reshuf / sln * 100
	var rtpsym = lrtp + srtp
	var q = float64(s.FreeCount) / reshuf
	var rtp = rtpsym + q*rtpfs
	fmt.Printf("completed %.5g%%, selected %d lines, time spent %v\n", reshuf/float64(s.Planned())*100, g.Sel, dur)
	fmt.Printf("reels lengths [%d, %d, %d, %d, %d], total reshuffles %d\n",
		len(reels.Reel(1)), len(reels.Reel(2)), len(reels.Reel(3)), len(reels.Reel(4)), len(reels.Reel(5)), reels.Reshuffles())
	fmt.Printf("symbols: %.5g(lined) + %.5g(scatter) = %.6f%%\n", lrtp, srtp, rtpsym)
	fmt.Printf("free spins %d, q = %.6f\n", s.FreeCount, q)
	fmt.Printf("free games frequency: 1/%.5g\n", reshuf/float64(s.FreeHits))
	fmt.Printf("RTP = %.5g(sym) + %.5g*%.5g(fg) = %.6f%%\n", rtpsym, q, rtpfs, rtp)
	return rtp
}
