package gochess

import (
	"math"
)

type BoardCoord struct {
	Row int
	Col int
}

type GameState struct {
	Board MailboxBoard

	BlackCanCastleKingside  bool
	BlackCanCastleQueenside bool
	WhiteCanCastleKingside  bool
	WhiteCanCastleQueenside bool

	EpTarget *BoardCoord
	PlyClock int

	IsWhiteTurn   bool
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
	movingPiece := state.Board.PieceAt(move.SrcRow, move.SrcCol)
	movingSide := state.Board.SideAt(move.SrcRow, move.SrcCol)

	if movingSide == Black && state.IsWhiteTurn || movingSide == White && !state.IsWhiteTurn {
		return false
	}

	var canMoveTo bool

	switch movingPiece {
	case Empty:
		return false
	case Pawn:
		canMove := state.pawnCanMove(movingSide, move.SrcRow, move.SrcCol, move.DstRow, move.DstCol)
		canMoveTo = state.pawnCanSee(movingSide, move.SrcRow, move.SrcCol, move.DstRow, move.DstCol) || canMove
	case Bishop:
		canMoveTo = state.canSeeDiagonal(move.SrcRow, move.SrcCol, move.DstRow, move.DstCol)
	case Knight:
		canMoveTo = state.knightCanSee(move.SrcRow, move.SrcCol, move.DstRow, move.DstCol)
	case King:
		canMoveTo = state.kingCanSee(movingSide, move.SrcRow, move.SrcCol, move.DstRow, move.DstCol)
	case Queen:
		canSeeDiagonal := state.canSeeDiagonal(move.SrcRow, move.SrcCol, move.DstRow, move.DstCol)
		canMoveTo = state.canSeeRook(move.SrcRow, move.SrcCol, move.DstRow, move.DstCol) || canSeeDiagonal
	case Rook:
		canMoveTo = state.canSeeRook(move.SrcRow, move.SrcCol, move.DstRow, move.DstCol)
	}

	legalWithoutCheck := canMoveTo && !state.squareTaken(movingSide, move.DstRow, move.DstCol)

	var testState GameState
	testState.Board = state.Board
	testState.IsWhiteTurn = !state.IsWhiteTurn

	testState.Board.PlacePiece(move.DstRow, move.DstCol, movingPiece, movingSide)
	testState.Board.RemovePiece(move.SrcRow, move.SrcCol)

	return legalWithoutCheck && !testState.IsInCheck(movingSide)
}

func (state *GameState) locateKing(side Side) BoardCoord {
	loc := BoardCoord { Row: -1, Col: -1 }

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if state.Board.PieceAt(i, j) == King && state.Board.SideAt(i, j) == side {
				loc.Row = i
				loc.Col = j
			}
		}
	}

	return loc
}

func (state *GameState) IsInCheck(side Side) bool {
	kingLoc := state.locateKing(side)

	var otherside Side
	if side == White {
		otherside = Black
	} else {
		otherside = White
	}

	// Look for vertical checks
	isVerticalHorizontalChecker := func(row int, col int) bool {
		if state.Board.SideAt(row, col) == otherside {
			switch state.Board.PieceAt(row, col) {
			case Rook:
				return true
			case Queen:
				return true
			case King:
				return math.Abs(float64(row - kingLoc.Row)) <= 1 && math.Abs(float64(col - kingLoc.Col)) <= 1
			default:
				return false
			}
		} else {
			return false
		}
	}

	// Look for diagonal checks
	isDiagonalChecker := func(row int, col int) bool {
		if state.Board.SideAt(row, col) == otherside {
			switch state.Board.PieceAt(row, col) {
			case Bishop:
				return true
			case Queen:
				return true
			case King:
				return math.Abs(float64(row - kingLoc.Row)) <= 1 && math.Abs(float64(col - kingLoc.Col)) <= 1
			case Pawn:
				var verticalRange int
				if side == White {
					verticalRange = 1
				} else {
					verticalRange = -1
				}
				if row-kingLoc.Row == verticalRange && math.Abs(float64(col-kingLoc.Col)) == 1 {
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
		return state.Board.SideAt(row, col) == otherside && state.Board.PieceAt(row, col) == Knight
	}

	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if isVerticalHorizontalChecker(row, col) && state.canSeeRook(row, col, kingLoc.Row, kingLoc.Col) {
				return true
			} else if isDiagonalChecker(row, col) && state.canSeeDiagonal(row, col, kingLoc.Row, kingLoc.Col) {
				return true
			} else if isOpponentKnight(row, col) && state.knightCanSee(row, col, kingLoc.Row, kingLoc.Col) {
				return true
			}
		}
	}

	return false
}

func (state *GameState) IsInCheckmate(side Side) bool {
	if !state.IsInCheck(side) {
		return false
	} else {
		kingLoc := state.locateKing(side)
		for i := int(math.Max(float64(kingLoc.Row - 1), 0)); i <= int(math.Min(float64(kingLoc.Row) + 1, 7)); i++ {
			for j := int(math.Max(float64(kingLoc.Col - 1), 0)); j <= int(math.Min(float64(kingLoc.Col) + 1, 7)); j++ {
				move := Move {
					SrcRow: kingLoc.Row,
					SrcCol: kingLoc.Col,
					DstRow: i,
					DstCol: j,
				}
				if state.IsMoveLegal(move) {
					return false
				}
			}
		}

		return true
	}
}

func (oldState *GameState) TakeTurn(move Move) (GameState, bool) {
	state := *oldState
	if !state.IsMoveLegal(move) {
		return *oldState, false
	} else {
		movingPiece := state.Board.PieceAt(move.SrcRow, move.SrcCol)
		movingSide := state.Board.SideAt(move.SrcRow, move.SrcCol)
		if state.moveIsCastlesLong(move) {
			state.Board.RemovePiece(move.SrcRow, 7)
			state.Board.PlacePiece(move.SrcRow, 4, Rook, movingSide)
		} else if state.moveIsCastlesShort(move) {
			state.Board.RemovePiece(move.SrcRow, 0)
			state.Board.PlacePiece(move.SrcRow, 2, Rook, movingSide)
		} else if state.moveIsEnPassant(move) {
			state.Board.RemovePiece(move.SrcRow, move.DstCol)
		}
		state.Board.RemovePiece(move.DstRow, move.DstCol)
		state.Board.RemovePiece(move.SrcRow, move.SrcCol)
		if move.IsPromotion {
			state.Board.PlacePiece(move.DstRow, move.DstCol, move.PromotionPiece, movingSide)
		} else {
			state.Board.PlacePiece(move.DstRow, move.DstCol, movingPiece, movingSide)
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
				state.WhiteCanCastleKingside = state.WhiteCanCastleKingside && move.SrcCol != 0
				state.WhiteCanCastleQueenside = state.WhiteCanCastleQueenside && move.SrcRow != 7
			} else {
				state.BlackCanCastleKingside = state.BlackCanCastleKingside && move.SrcCol != 0
				state.BlackCanCastleQueenside = state.BlackCanCastleQueenside && move.SrcRow != 7
			}
		}

		if movingPiece == Pawn && math.Abs(float64(move.SrcRow) - float64(move.DstRow)) == 2 {
			oldState.EpTarget = &BoardCoord{}
			oldState.EpTarget.Col = move.SrcCol
			var sign int
			if math.Signbit(float64(move.DstRow) - float64(move.SrcRow)) {
				sign = -1
			} else {
				sign = 1
			}
			oldState.EpTarget.Row = move.SrcRow + sign
		} else {
			state.EpTarget = nil
		}

		state.IsWhiteTurn = !state.IsWhiteTurn

		return state, true
	}
}

// The "move is special move" functions assume the move has already been determined to be legal. They are for identifying after the intial legality check.
func (state *GameState) moveIsCastlesShort(move Move) bool {
	return state.Board.PieceAt(move.SrcRow, move.SrcCol) == King && move.DstCol-move.SrcCol == -2
}

func (state *GameState) moveIsCastlesLong(move Move) bool {
	return state.Board.PieceAt(move.SrcRow, move.SrcCol) == King && move.DstCol-move.SrcCol == 2
}

func (state *GameState) moveIsEnPassant(move Move) bool {
	pa := state.Board.PieceAt(move.DstRow, move.DstCol)
	return state.Board.PieceAt(move.SrcRow, move.SrcCol) == Pawn && math.Abs(float64(move.DstCol)-float64(move.SrcCol)) == 1 && pa == Empty
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
		for row >= 0 && row < 8 && col >= 0 && col < 8 && row != destRow && col != destCol {
			if state.Board.SideAt(row, col) != Empty {
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
			if state.EpTarget != nil && state.EpTarget.Row == destRow && state.EpTarget.Col == destCol {
				return true
			}
		}
	}

	return false
}

func (state *GameState) pawnCanMove(side Side, srcRow int, srcCol int, destRow int, destCol int) bool {
	var dir int
	if side == White {
		dir = 1
	} else {
		dir = -1
	}

	if state.Board.SideAt(destRow, destCol) != Empty {
		return false
	}

	if srcCol-destCol == 0 {
		var startingRow int
		if side == White {
			startingRow = 1
		} else {
			startingRow = 6
		}

		verticalRange := dir
		if srcRow == startingRow {
			verticalRange = int(math.Copysign(2, float64(dir)))
		}

		if destRow-srcRow == verticalRange || destRow-srcRow == verticalRange-dir {
			for row := srcRow + dir; row != destRow; row += dir {
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
	canMoveTo := (canSeeDiagonal || canSeeRook) && (math.Abs(float64((srcRow-destRow))) <= 1 && math.Abs(float64((srcCol-destCol))) <= 1)

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

func (state *GameState) piecesInRange(piece Piece, side Side, row int, col int) []BoardCoord {
	var ret []BoardCoord
	if (state.IsWhiteTurn && side == Black) || (!state.IsWhiteTurn && side == White) {
		return ret
	}

	for i := 0; i < 8; i += 1 {
		for j := 0; j < 8; j += 1 {
			if state.Board.SideAt(i, j) == side && state.Board.PieceAt(i, j) == piece {
				fakeMove := Move{
					SrcRow:         i,
					SrcCol:         j,
					DstRow:         row,
					DstCol:         col,
					IsPromotion:    false,
					PromotionPiece: 0,
				}
				if state.IsMoveLegal(fakeMove) {
					coord := BoardCoord{
						Row: i,
						Col: j,
					}
					ret = append(ret, coord)
				}
			}
		}
	}

	return ret
}
