package gochess

import (
	"strconv"
	"strings"
)

// This is very (frustratingly) minimal, but it's intended to minimize the work the GUI has to do. Ideally, the GUI should need to
// know absolutely nothing about the rules of chess beyond the fact that there's an 8x8 grid and promotion.

type Move struct {
	Src_row        int
	Src_col        int
	Dest_row       int
	Dest_col       int
	IsPromotion    bool
	PromotionPiece Piece
}

func (move *Move) PrettyPrint(state *GameState) string {
	getPieceName := func(piece Piece) string {
		switch piece {
		case King:
			return "K"
		case Knight:
			return "N"
		case Bishop:
			return "B"
		case Queen:
			return "Q"
		case Rook:
			return "R"
		default:
			return ""
		}
	}

	var builder strings.Builder

	columnName := string(rune(104 - move.Dest_col))
	rowName := strconv.FormatInt(int64(1+move.Dest_row), 10)

	pieceName := getPieceName(state.Board.PieceAt(move.Src_row, move.Src_col))

	builder.WriteString(pieceName)
	builder.WriteString(columnName)
	builder.WriteString(rowName)

	if move.IsPromotion {
		promotionPieceName := getPieceName(move.PromotionPiece)
		builder.WriteString("=")
		builder.WriteString(promotionPieceName)
	}

	return builder.String()
}
