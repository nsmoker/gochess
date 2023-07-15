package gochess

import (
	"fmt"
	"testing"
)

func TestPgnParsing(t *testing.T) {
    pgn := `[Event "Live Chess"]
	[Site "Chess.com"]
	[Date "2023.05.09"]
	[Round "-"]
	[White "4gateftw"]
	[Black "kandu12345"]
	[Result "1-0"]
	[WhiteElo "1407"]
	[BlackElo "1348"]
	[TimeControl "900+10"]
	[Termination "4gateftw won by resignation"]
	[UTCDate "2023.05.10"]
	[UTCTime "05:06:53"]
	[Variant "Standard"]
	[ECO "A50"]
	[Opening "Slav Indian"]
	[Annotator "https://lichess.org/@/fourgateftw"]
	
	1. d4 { [%clk 0:15:10] } 1... Nf6 { [%clk 0:15:03] } 2. c4 { [%clk 0:15:16] } 2... c6 { [%clk 0:15:09] } 3. Nc3 { [%clk 0:15:23] } 3... e6 { [%clk 0:15:10] } 4. e4 { [%clk 0:15:27] } 4... Bb4 { [%clk 0:15:12] } 5. e5 { [%clk 0:14:49] } 5... Ne4 { [%clk 0:15:19] } 6. Qc2 { [%clk 0:14:56] } 6... Nxc3 { [%clk 0:15:13] } 7. bxc3 { [%clk 0:15:02] } 7... Ba5 { [%clk 0:15:13] } 8. Nf3 { [%clk 0:14:58] } 8... h6 { [%clk 0:15:18] } 9. Bd3 { [%clk 0:14:58] } 9... g6 { [%clk 0:15:16] } 10. Bxg6 { [%clk 0:14:39] } 10... Rg8 { [%clk 0:14:45] } 11. Bh5 { [%clk 0:14:12] } 11... Rh8 { [%clk 0:14:17] } 12. Ba3 { [%clk 0:08:53] } 12... b5 { [%clk 0:13:25] } 13. cxb5 { [%clk 0:08:07] } 13... cxb5 { [%clk 0:13:29] } 14. O-O { [%clk 0:05:48] } 14... Bb7 { [%clk 0:13:35] } 15. Qd2 { [%clk 0:05:32] } 15... Nc6 { [%clk 0:13:14] } 16. Qf4 { [%clk 0:05:28] } 16... Rf8 { [%clk 0:12:53] } 17. Bxf8 { 1-0 White wins. } { [%clk 0:05:32] } 1-0
	`

	asd := ParsePgn(pgn)
	for len(asd.Next) > 0 {
		fmt.Println(asd.State.Board.PrettyPrint())
		asd = *asd.Next[0]
	}
}