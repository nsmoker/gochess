package gochess

import (
	"strconv"
	"strings"
	"unicode"
)

// Note that the FEN functions here are do not absolve the GUI of responsibility, it still needs to display FEN strings.

func ToFEN(state *GameState) string {
	var fen strings.Builder
	counter := 0
	blacks := [6]string{"k", "q", "r", "b", "n", "p"}
	whites := [6]string{"K", "Q", "R", "B", "N", "P"}
	for row := 7; row >= 0; row -= 1 {
		for col := 7; col >= 0; col -= 1 {
			if state.Board.PieceAt(row, col) != Empty {
				if counter > 0 {
					fen.WriteString(strconv.Itoa(counter))
					counter = 0
				}
				if state.Board.SideAt(row, col) == Black {
					fen.WriteString(blacks[state.Board.PieceAt(row, col)])
				} else {
					fen.WriteString(whites[state.Board.PieceAt(row, col)])
				}
			} else {
				counter += 1
			}

			if col == 0 {
				if counter > 0 {
					fen.WriteString(strconv.Itoa(counter))
					counter = 0
				}
				fen.WriteString("/")
			}
		}
	}

	if state.IsWhiteTurn {
		fen.WriteString(" w")
	} else {
		fen.WriteString(" b")
	}

	fen.WriteString(" ")

	if state.WhiteCanCastleKingside {
		fen.WriteString("K")
	}

	if state.WhiteCanCastleQueenside {
		fen.WriteString("Q")
	}

	if state.BlackCanCastleKingside {
		fen.WriteString("k")
	}

	if state.BlackCanCastleQueenside {
		fen.WriteString("q")
	}

	if !(state.BlackCanCastleKingside || state.WhiteCanCastleKingside || state.BlackCanCastleQueenside || state.WhiteCanCastleQueenside) {
		fen.WriteString("-")
	}

	fen.WriteString(" -")

	return fen.String()
}

func ParseFEN(fen string) (GameState, error) {
	var state GameState
	row := 7
	col := 7
	i := 0
	spaceReached := false

	// First field
	for !spaceReached {
		switch fen[i] {
		case 'r':
			state.Board.PlacePiece(row, col, Rook, Black)
		case 'n':
			state.Board.PlacePiece(row, col, Knight, Black)
		case 'b':
			state.Board.PlacePiece(row, col, Bishop, Black)
		case 'k':
			state.Board.PlacePiece(row, col, King, Black)
		case 'q':
			state.Board.PlacePiece(row, col, Queen, Black)
		case 'p':
			state.Board.PlacePiece(row, col, Pawn, Black)
		case 'R':
			state.Board.PlacePiece(row, col, Rook, White)
		case 'N':
			state.Board.PlacePiece(row, col, Knight, White)
		case 'B':
			state.Board.PlacePiece(row, col, Bishop, White)
		case 'K':
			state.Board.PlacePiece(row, col, King, White)
		case 'Q':
			state.Board.PlacePiece(row, col, Queen, White)
		case 'P':
			state.Board.PlacePiece(row, col, Pawn, White)
		case '/':
			row--
			col = 7
			i++
			continue
		case ' ':
			spaceReached = true
		default:
			if unicode.IsNumber(rune(fen[i])) {
				val, _ := strconv.Atoi(string(fen[i]))
				for val > 0 {
					state.Board.RemovePiece(row, col)
					col--
					val--
				}
				if col < 0 {
					col = 7
				}
				i++
				continue
			}
		}
		col--
		i++
	}

	// Second field
	switch fen[i] {
	case 'w':
		state.IsWhiteTurn = true
	case 'b':
		state.IsWhiteTurn = false
	}

	i += 2
	// Third field
	spaceReached = false
	for !spaceReached {
		switch fen[i] {
		case 'K':
			state.WhiteCanCastleKingside = true
		case 'Q':
			state.WhiteCanCastleQueenside = true
		case 'k':
			state.BlackCanCastleKingside = true
		case 'q':
			state.BlackCanCastleQueenside = true
		case ' ':
			spaceReached = true
		}
		i++
	}

	// Fourth field
	spaceReached = false
	if fen[i] == '-' {
		i += 2
	} else {
		col := 104 - fen[i]
		row := fen[i+1]
		state.PreviousMove.DstRow = int(row)
		state.PreviousMove.DstCol = int(col)
	}

	return state, nil
}
