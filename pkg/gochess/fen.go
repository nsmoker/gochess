package gochess

import "unicode"
import "strconv"

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
		row := fen[i + 1]
		state.PreviousMove.dest_row = int(row)
		state.PreviousMove.dest_col = int(col)
	}

	return state, nil
}
