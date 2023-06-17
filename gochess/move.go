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

func getColumnName(column int) string {
	return string(rune(104 - column))
}

func getRowName(row int) string {
	return strconv.FormatInt(int64(1 + row), 10)
}

func getPieceName(piece Piece) string {
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

func (move *Move) PrettyPrint(state *GameState) string {
	var builder strings.Builder
	piece := state.Board.PieceAt(move.Src_row, move.Src_col)
	//side := state.Board.SideAt(move.Src_row, move.Src_col)

	columnName := getColumnName(move.Dest_col)
	rowName := getRowName(move.Dest_row)

	pieceName := getPieceName(piece)

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
