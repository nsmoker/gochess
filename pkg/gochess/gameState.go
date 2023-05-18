package gochess

import (
	"math"
)

type GameState struct {
	Board MailboxBoard

	BlackCanCastleKingside  bool
	BlackCanCastleQueenside bool
	WhiteCanCastleKingside  bool
	WhiteCanCastleQueenside bool

	IsWhiteTurn   bool
	PreviousMove  *Move
	PreviousState *GameState
}

func MakeStartingState() GameState {
	var ret GameState
	ret.Board = MakeStartingBoard()
	ret.BlackCanCastleQueenside = true
	ret.WhiteCanCastleKingside = true
	ret.BlackCanCastleKingside = true
	ret.WhiteCanCastleQueenside = true
	ret.IsWhiteTurn = true
	return ret
}

func (state *GameState) IsMoveLegal(move Move) bool {
	movingPiece := state.Board.PieceAt(move.src_row, move.src_col)
	movingSide := state.Board.SideAt(move.src_row, move.src_col)

	if movingSide == Black && state.IsWhiteTurn || movingSide == White && !state.IsWhiteTurn {
		return false
	}

	var canMoveTo bool

	switch movingPiece {
	case Empty:
		return false
	case Pawn:
		canMove := state.pawnCanMove(movingSide, move.src_row, move.src_col, move.dest_row, move.dest_col)
		canMoveTo = state.pawnCanSee(movingSide, move.src_row, move.src_col, move.dest_row, move.dest_col) || canMove
	case Bishop:
		canMoveTo = state.canSeeDiagonal(move.src_row, move.src_col, move.dest_row, move.dest_col)
	case Knight:
		canMoveTo = state.knightCanSee(move.src_row, move.src_col, move.dest_row, move.dest_col)
	case King:
		canMoveTo = state.kingCanSee(movingSide, move.src_row, move.src_col, move.dest_row, move.dest_col)
	case Queen:
		canSeeDiagonal := state.canSeeDiagonal(move.src_row, move.src_col, move.dest_row, move.dest_col)
		canMoveTo = state.canSeeRook(move.src_row, move.src_col, move.dest_row, move.dest_col) || canSeeDiagonal
	case Rook:
		canMoveTo = state.canSeeRook(move.src_row, move.src_col, move.dest_row, move.dest_col)
	}

	legalWithoutCheck := canMoveTo && !state.squareTaken(movingSide, move.dest_row, move.dest_col)

	var testState GameState
	testState.Board = state.Board
	testState.PreviousMove = state.PreviousMove
	testState.PreviousState = state.PreviousState

	testState.Board.PlacePiece(move.dest_row, move.dest_col, movingPiece, movingSide)
	testState.Board.RemovePiece(move.src_row, move.src_col)

	return legalWithoutCheck && !testState.IsInCheck(movingSide)
}

func (state *GameState) IsInCheck(side Side) bool {
	// Find the king

	kingRow := -1
	kingCol := -1
	board := state.Board

	var otherside Side
	if side == White {
		otherside = Black
	} else {
		otherside = White
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board.PieceAt(i, j) == King && board.SideAt(i, j) == side {
				kingRow = i
				kingCol = j
			}
		}
	}

	// Look for vertical checks
	isVerticalHorizontalChecker := func(row int, col int) bool {
		if board.SideAt(row, col) == otherside {
			switch board.PieceAt(row, col) {
			case Rook:
				return true
			case Queen:
				return true
			default:
				return false
			}
		} else {
			return false
		}
	}

	// Look for diagonal checks
	isDiagonalChecker := func(row int, col int) bool {
		if board.SideAt(row, col) == otherside {
			switch board.PieceAt(row, col) {
			case Bishop:
				return true
			case Queen:
				return true
			case Pawn:
				var verticalRange int
				if side == White {
					verticalRange = 1
				} else {
					verticalRange = -1
				}
				if row-kingRow == verticalRange && math.Abs(float64(col-kingCol)) == 1 {
					return true
				} else {
					return false
				}
			default:
				return false
			}
		} else {
			return false
		}
	}

	isOpponentKnight := func(row int, col int) bool {
		return board.SideAt(row, col) == otherside && board.PieceAt(row, col) == Knight
	}

	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if isVerticalHorizontalChecker(row, col) && state.canSeeRook(row, col, kingRow, kingCol) {
				return true
			} else if isDiagonalChecker(row, col) && state.canSeeDiagonal(row, col, kingRow, kingCol) {
				return true
			} else if isOpponentKnight(row, col) && state.knightCanSee(row, col, kingRow, kingCol) {
				return true
			}
		}
	}

	return false
}

func (state *GameState) TakeTurn(move Move) bool {
	if !state.IsMoveLegal(move) {
		return false
	} else {
		movingPiece := state.Board.PieceAt(move.src_row, move.src_col)
		movingSide := state.Board.SideAt(move.src_row, move.src_col)
		if state.moveIsCastlesLong(move) {
			state.Board.RemovePiece(move.src_row, 7)
			state.Board.PlacePiece(move.src_row, 4, Rook, White)
		} else if state.moveIsCastlesShort(move) {
			state.Board.RemovePiece(move.src_row, 0)
			state.Board.PlacePiece(move.src_row, 2, Rook, White)
		} else if state.moveIsEnPassant(move) {
			state.Board.RemovePiece(move.src_row, move.dest_col)
		}
		state.Board.RemovePiece(move.dest_row, move.dest_col)
		state.Board.RemovePiece(move.src_row, move.src_col)
		if move.isPromotion {
			state.Board.PlacePiece(move.dest_row, move.dest_col, move.promotionPiece, movingSide)
		} else {
			state.Board.PlacePiece(move.dest_row, move.dest_col, movingPiece, movingSide)
		}

		if movingPiece == King {
			if movingSide == Black {
				state.BlackCanCastleKingside = false
				state.BlackCanCastleQueenside = false
			} else {
				state.WhiteCanCastleKingside = false
				state.WhiteCanCastleQueenside = false
			}
		} else if movingPiece == Rook {
			if movingSide == White {
				state.WhiteCanCastleKingside = state.WhiteCanCastleKingside && move.src_col != 0
				state.WhiteCanCastleQueenside = state.WhiteCanCastleQueenside && move.src_row != 7
			} else {
				state.BlackCanCastleKingside = state.BlackCanCastleKingside && move.src_col != 0
				state.BlackCanCastleQueenside = state.BlackCanCastleQueenside && move.src_row != 7
			}
		}

		prevState := *state
		state.PreviousState = &prevState
		state.PreviousMove = &move

		state.IsWhiteTurn = !state.IsWhiteTurn

		return true
	}
}

// The "move is special move" functions assume the move has already been determined to be legal. They are for identifying after the intial legality check.
func (state *GameState) moveIsCastlesShort(move Move) bool {
	return state.Board.PieceAt(move.src_row, move.src_col) == King && move.dest_col-move.src_col == -2
}

func (state *GameState) moveIsCastlesLong(move Move) bool {
	return state.Board.PieceAt(move.src_row, move.src_col) == King && move.dest_col-move.src_col == 2
}

func (state *GameState) moveIsEnPassant(move Move) bool {
	pa := state.Board.PieceAt(move.dest_row, move.dest_col)
	return state.Board.PieceAt(move.src_row, move.src_col) == Pawn && math.Abs(float64(move.dest_col)-float64(move.src_col)) == 1 && pa == Empty
}

func (state *GameState) squareTaken(side Side, row int, col int) bool {
	return state.Board.SideAt(row, col) == side
}

// Naming convention: canSee == can attack

func (state *GameState) canSeeRook(srcRow int, srcCol int, destRow int, destCol int) bool {
	if srcRow == destRow && srcCol != destCol {
		dir := int(math.Copysign(1, float64(destCol-srcCol)))
		for col := srcCol + dir; col != destCol; col += dir {
			if state.Board.PieceAt(srcRow, col) != Empty {
				return false
			}
		}
		return true
	} else if srcRow != destRow && srcCol == destCol {
		dir := int(math.Copysign(1, float64(destRow-srcRow)))
		for row := srcRow + dir; row != destRow; row += dir {
			if state.Board.PieceAt(row, srcCol) != Empty {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func (state *GameState) canSeeDiagonal(srcRow int, srcCol int, destRow int, destCol int) bool {
	if srcRow == destRow || srcCol == destCol || math.Abs(float64(srcRow-destRow)) != math.Abs(float64(srcCol-destCol)) {
		return false
	} else {
		rowDir := math.Copysign(1, float64(destRow-srcRow))
		colDir := math.Copysign(1, float64(destCol-srcCol))
		row := srcRow + int(rowDir)
		col := srcCol + int(colDir)
		for row > 0 && row < 8 && col > 0 && col < 8 && row != destRow && col != destCol {
			if state.Board.PieceAt(row, col) != Empty {
				return false
			}

			row += int(rowDir)
			col += int(colDir)
		}

		return true
	}
}

func (state *GameState) knightCanSee(srcRow int, srcCol int, destRow int, destCol int) bool {
	return (math.Abs(float64(srcRow-destRow)) == 2 && math.Abs(float64(srcCol-destCol)) == 1) || (math.Abs(float64(srcRow-destRow)) == 1 && math.Abs(float64(srcCol-destCol)) == 2)
}

func (state *GameState) pawnCanSee(side Side, srcRow int, srcCol int, destRow int, destCol int) bool {
	var otherside Side
	if side == Black {
		otherside = White
	} else {
		otherside = Black
	}

	var verticalRange int
	if side == White {
		verticalRange = 1
	} else {
		verticalRange = -1
	}

	if destRow-srcRow == verticalRange && math.Abs(float64(srcCol-destCol)) == 1 {
		if state.Board.SideAt(destRow, destCol) == otherside {
			return true
		} else {
			if state.Board.PieceAt(srcRow, destCol) == Pawn && state.Board.SideAt(srcRow, destCol) == otherside {
				if state.PreviousMove != nil && state.PreviousMove.dest_row == srcRow && state.PreviousMove.dest_col == destCol {
					return true
				}
			}
		}
	}

	return false
}

func (state *GameState) pawnCanMove(side Side, srcRow int, srcCol int, destRow int, destCol int) bool {
	var verticalRange int
	if side == White {
		verticalRange = 1
	} else {
		verticalRange = -1
	}

	if srcCol-destCol == 0 {
		var startingRow int
		if side == White {
			startingRow = 1
		} else {
			startingRow = 6
		}

		if srcRow == startingRow {
			verticalRange = int(math.Copysign(2, float64(verticalRange)))
		}

		dir := int(math.Copysign(1, float64(verticalRange)))

		if destRow-srcRow == verticalRange || destRow-srcRow == verticalRange-dir {
			for row := srcRow + dir; row <= destRow; row += dir {
				if state.Board.SideAt(row, srcCol) != Empty {
					return false
				}
			}
			return true
		}
	}

	return false
}

func (state *GameState) kingCanSee(side Side, srcRow int, srcCol int, destRow int, destCol int) bool {
	canSeeDiagonal := state.canSeeDiagonal(srcRow, srcCol, destRow, destCol)
	canSeeRook := state.canSeeRook(srcRow, srcCol, destRow, destCol)
	canMoveTo := canSeeDiagonal && canSeeRook && math.Abs(float64((srcRow-destRow)+(srcCol-destCol))) == 1

	if canMoveTo {
		return true
	} else {
		if srcRow == destRow && math.Abs(float64(destCol)-float64(srcCol)) == 2 && canSeeRook {
			testState := *state
			dir := int(math.Copysign(1, float64(destCol)-float64(srcCol)))
			for col := srcCol; col != destCol; col += dir {
				testState.Board.RemovePiece(destRow, col)
				testState.Board.PlacePiece(destRow, col+dir, King, side)
				if testState.IsInCheck(side) {
					return false
				}
			}

			if math.Signbit(float64(destCol) - float64(srcCol)) {
				if side == White {
					return state.WhiteCanCastleKingside
				} else {
					return state.BlackCanCastleKingside
				}
			} else {
				if side == White {
					return state.WhiteCanCastleQueenside
				} else {
					return state.BlackCanCastleQueenside
				}
			}
		} else {
			return false
		}
	}
}
