package slotopoldeluxe

import (
	"context"
	"fmt"
	"time"

	"github.com/slotopol/server/game/slot"
	"github.com/slotopol/server/game/slot/megajack/slotopol"
)

func CalcStat(ctx context.Context, mrtp float64) float64 {
	fmt.Printf("*bonus games calculations*\n")
	slotopol.ExpEldorado()
	slotopol.ExpMonopoly()
	fmt.Printf("*reels calculations*\n")
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
	var qmje1 = float64(s.BonusCount[mje1]) / reshuf / sln
	var rtpmje1 = slotopol.Emje * 1 * qmje1 * 100
	var qmje3 = float64(s.BonusCount[mje3]) / reshuf / sln
	var rtpmje3 = slotopol.Emje * 3 * qmje3 * 100
	var qmje6 = float64(s.BonusCount[mje6]) / reshuf / sln
	var rtpmje6 = slotopol.Emje * 6 * qmje6 * 100
	var qmjm = float64(s.BonusCount[mjm]) / reshuf / sln
	var rtpmjm = slotopol.Emjm * qmjm * 100
	var rtp = rtpsym + rtpmje1 + rtpmje3 + rtpmje6 + rtpmjm
	fmt.Printf("completed %.5g%%, selected %d lines, time spent %v\n", reshuf/float64(s.Planned())*100, g.Sel, dur)
	fmt.Printf("reels lengths [%d, %d, %d, %d, %d], total reshuffles %d\n",
		len(reels.Reel(1)), len(reels.Reel(2)), len(reels.Reel(3)), len(reels.Reel(4)), len(reels.Reel(5)), reels.Reshuffles())
	fmt.Printf("symbols: %.5g(lined) + %.5g(scatter) = %.6f%%\n", lrtp, srtp, rtpsym)
	fmt.Printf("spin1 bonuses: frequency 1/%.5g, rtp = %.6f%%\n", reshuf/float64(s.BonusCount[mje1]), rtpmje1)
	fmt.Printf("spin3 bonuses: frequency 1/%.5g, rtp = %.6f%%\n", reshuf/float64(s.BonusCount[mje3]), rtpmje3)
	fmt.Printf("spin6 bonuses: frequency 1/%.5g, rtp = %.6f%%\n", reshuf/float64(s.BonusCount[mje6]), rtpmje6)
	fmt.Printf("monopoly bonuses: frequency 1/%.5g, rtp = %.6f%%\n", reshuf/float64(s.BonusCount[mjm]), rtpmjm)
	if s.JackCount[mjj] > 0 {
		fmt.Printf("jackpots: count %d, frequency 1/%.12g\n", s.JackCount[mjj], reshuf/float64(s.JackCount[mjj]))
	}
	fmt.Printf("RTP = %.5g(sym) + %.5g(mje) + %.5g(mjm) = %.6f%%\n", rtpsym, rtpmje1+rtpmje3+rtpmje6, rtpmjm, rtp)
	return rtp
}
