//go:build !prod || full || novomatic

package links

import (
	"context"

	slot "github.com/slotopol/server/game/slot/beetlemania"
	"github.com/spf13/pflag"
)

func init() {
	var gi = GameInfo{
		Aliases: []GameAlias{
			{ID: "beetlemania", Name: "Beetle Mania"},
			{ID: "beetlemaniadeluxe", Name: "Beetle Mania Deluxe"},
			{ID: "hottarget", Name: "Hot Target"},
		},
		Provider: "Novomatic",
		SX:       5,
		SY:       3,
		LN:       10,
		FG:       FGhas,
		BN:       0,
		RTP:      MakeRtpList(slot.ReelsMap),
	}
	GameList = append(GameList, gi)

	for _, ga := range gi.Aliases {
		ScanIters = append(ScanIters, func(flags *pflag.FlagSet, ctx context.Context) {
			if is, _ := flags.GetBool(ga.ID); is {
				var rn, _ = flags.GetString("reels")
				if rn == "bon" || rn == "bonu" {
					slot.CalcStatBon(ctx, rn)
				} else {
					slot.CalcStatReg(ctx, rn)
				}
			}
		})
		GameFactory[ga.ID] = func() any {
			return slot.NewGame()
		}
	}
}
