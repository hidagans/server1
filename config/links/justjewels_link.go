//go:build !prod || full || novomatic

package links

import (
	"context"

	"github.com/slotopol/server/game/justjewels"
	"github.com/spf13/pflag"
)

func init() {
	var gi = GameInfo{
		Aliases: []GameAlias{
			{"justjewels", "Just Jewels"},
			{"justjewelsdeluxe", "Just Jewels Deluxe"},
		},
		Provider: "Novomatic",
		ScrnX:    5,
		ScrnY:    3,
		RtpList:  MakeRtpList(justjewels.ReelsMap),
	}
	GameList = append(GameList, gi)

	for _, ga := range gi.Aliases {
		ScanIters = append(ScanIters, func(flags *pflag.FlagSet, ctx context.Context) {
			if is, _ := flags.GetBool(ga.ID); is {
				var rn, _ = flags.GetString("reels")
				justjewels.CalcStat(ctx, rn)
			}
		})
		GameFactory[ga.ID] = func() any {
			return justjewels.NewGame()
		}
	}
}
