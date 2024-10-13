//go:build !prod || full || playngo

package links

import (
	"context"

	slot "github.com/slotopol/server/game/slot/copsnrobbers"
	"github.com/spf13/pflag"
)

func init() {
	var gi = GameInfo{
		Aliases: []GameAlias{
			{ID: "copsnrobbers", Name: "Cops'n'Robbers"},
		},
		Provider: "Play'n GO",
		SX:       5,
		SY:       3,
		GP:       GPsel | GPfghas | GPfgreel | GPfgmult | GPscat | GPwild,
		SN:       len(slot.LinePay),
		LN:       len(slot.BetLines),
		BN:       0,
		RTP:      MakeRtpList(slot.ReelsMap),
	}
	GameList = append(GameList, gi)

	for _, ga := range gi.Aliases {
		ScanIters = append(ScanIters, func(flags *pflag.FlagSet, ctx context.Context) {
			if is, _ := flags.GetBool(ga.ID); is {
				var mrtp, _ = flags.GetFloat64("reels")
				slot.CalcStatReg(ctx, mrtp)
			}
		})
		GameFactory[ga.ID] = func() any {
			return slot.NewGame()
		}
	}
}