package gochess

import "testing"
import "github.com/google/go-cmp/cmp"

func TestParseFEN(t *testing.T) {
	fen := "r1b1k2r/p4ppp/1pR1p3/6B1/1q1P4/n2QPN2/P3KPPP/7R b kq - 0 17"

	parseResult, _ := ParseFEN(fen)

	want := MailboxBoard{Pieces: [64]Piece{
		Rook,  Empty,  Empty,  Empty, Empty, Empty,  Empty,  Empty,
        Pawn,  Pawn,   Pawn,   King,  Empty, Empty,  Empty,  Pawn,
        Empty, Empty,  Knight, Pawn,  Queen, Empty,  Empty,  Knight,
        Empty, Empty,  Empty,  Empty, Pawn,  Empty,  Queen,  Empty,
        Empty, Bishop, Empty,  Empty, Empty, Empty,  Empty,  Empty,
        Empty, Empty,  Empty,  Pawn,  Empty, Rook,   Pawn,   Empty,
        Pawn,  Pawn,   Pawn,   Empty, Empty, Empty,  Empty,  Pawn,
        Rook,  Empty,  Empty,  King,  Empty, Bishop, Empty,  Rook},
        Sides: [64]Side{
			White, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
            White, White, White, White, Empty, Empty, Empty, White,
            Empty, Empty, White, White, White, Empty, Empty, Black,
            Empty, Empty, Empty, Empty, White, Empty, Black, Empty,
            Empty, White, Empty, Empty, Empty, Empty, Empty, Empty,
            Empty, Empty, Empty, Black, Empty, White, Black, Empty,
            Black, Black, Black, Empty, Empty, Empty, Empty, Black,
            Black, Empty, Empty, Black, Empty, Black, Empty, Black}}
	
	if !cmp.Equal(parseResult.Board, want) {
		t.Log(cmp.Diff(parseResult.Board, want))
		t.Fatalf("Got %s from %s, want %s", parseResult.Board.PrettyPrint(), fen, want.PrettyPrint())
	}
}