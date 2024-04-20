package powerstars

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/slotopol/server/game"
)

func BruteForce5x(ctx context.Context, s game.Stater, g game.SlotGame, reels game.Reels, wc2, wc3, wc4 float64) {
	var screen = g.NewScreen()
	defer screen.Free()
	var ws game.WinScan
	var r1 = reels.Reel(1)
	var r2 = reels.Reel(2)
	var r3 = reels.Reel(3)
	var r4 = reels.Reel(4)
	var r5 = reels.Reel(5)
	for i1 := range r1 {
		screen.SetCol(1, r1, i1)
		for i2 := range r2 {
			screen.SetCol(2, r2, i2)
			for i3 := range r3 {
				screen.SetCol(3, r3, i3)
				for i4 := range r4 {
					screen.SetCol(4, r4, i4)
					for i5 := range r5 {
						screen.SetCol(5, r5, i5)
						var sym2, sym3, sym4 = screen.At(2, 1), screen.At(3, 1), screen.At(4, 1)
						if rand.Float64() < wc2 {
							screen.Set(2, 1, wild)
						}
						if rand.Float64() < wc3 {
							screen.Set(3, 1, wild)
						}
						if rand.Float64() < wc4 {
							screen.Set(4, 1, wild)
						}
						g.Scanner(screen, &ws)
						screen.Set(2, 1, sym2)
						screen.Set(3, 1, sym3)
						screen.Set(4, 1, sym4)
						s.Update(&ws)
						ws.Reset()
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
	}
}

func CalcStatStars(ctx context.Context, wc2, wc3, wc4 float64) float64 {
	var reels = &Reels
	var g = NewGame("nil")
	g.SBL = game.MakeSblNum(1)
	var sbl = float64(g.SBL.Num())
	var s game.Stat

	var wcsym = func(wc float64) byte {
		if wc == 1. {
			return '*'
		}
		if wc == 0. {
			return '-'
		}
		return 'x'
	}
	fmt.Printf("calculations of star combinations [%c%c%c]\n", wcsym(wc2), wcsym(wc3), wcsym(wc4))

	var total = float64(reels.Reshuffles())
	var dur = func() time.Duration {
		var t0 = time.Now()
		var ctx2, cancel2 = context.WithCancel(ctx)
		defer cancel2()
		go s.Progress(ctx2, time.NewTicker(2*time.Second), sbl, total)
		BruteForce5x(ctx2, &s, g, reels, wc2, wc3, wc4)
		return time.Since(t0)
	}()

	var reshuf = float64(s.Reshuffles)
	var lrtp = s.LinePay / reshuf / sbl * 100
	_ = jid
	fmt.Printf("completed %.5g%%, selected %d lines, time spent %v\n", reshuf/total*100, g.SBL.Num(), dur)
	fmt.Printf("reels lengths [%d, %d, %d, %d, %d], total reshuffles %d\n",
		len(reels.Reel(1)), len(reels.Reel(2)), len(reels.Reel(3)), len(reels.Reel(4)), len(reels.Reel(5)), reels.Reshuffles())
	fmt.Printf("RTP[%c%c%c] = %.6f%%\n", wcsym(wc2), wcsym(wc3), wcsym(wc4), lrtp)
	return lrtp
}

func CalcStat(ctx context.Context, rn string) (rtp float64) {
	var wc float64
	if rn != "" {
		var ok bool
		if wc, ok = ChanceMap[rn]; !ok {
			return 0
		}
	} else {
		wc, rn = ChanceMap["95"], "95"
	}

	var b = 1 / wc
	fmt.Printf("wild chance %.5g, b = %.5g\n", wc, b)
	var rtp000 = CalcStatStars(ctx, 0, 0, 0)
	var rtp100 = CalcStatStars(ctx, 1, 0, 0)
	var rtp010 = CalcStatStars(ctx, 0, 1, 0)
	var rtp001 = CalcStatStars(ctx, 0, 0, 1)
	var rtp110 = CalcStatStars(ctx, 1, 1, 0)
	var rtp011 = CalcStatStars(ctx, 0, 1, 1)
	var rtp101 = CalcStatStars(ctx, 1, 0, 1)
	var rtp111 = CalcStatStars(ctx, 1, 1, 1)
	var q = AnyStarProb(b)
	var rtpfs = ((rtp100+rtp010+rtp001)*(b-1)*(b-1) + (rtp110+rtp011+rtp101)*(b-1) + rtp111) / (b*b + (b-1)*b + (b-1)*(b-1))
	rtp = rtp000 + q*rtpfs
	fmt.Printf("free spins: q = %.5g, 1/q = %.5g, rtpfs = %.6f%%\n", q, 1/q, rtpfs)
	fmt.Printf("RTP = %.5g(sym) + q*%.5g(fg) = %.6f%%\n", rtp000, rtpfs, rtp)
	return
}
