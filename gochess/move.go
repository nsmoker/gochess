package gochess

import (
	"strconv"
	"strings"
	"unicode"
)

// This is very (frustratingly) minimal, but it's intended to minimize the work that the GUI has to do. Ideally, the GUI should need to
// know absolutely nothing about the rules of chess beyond the fact that there's an 8x8 grid and promotion.

type Move struct {
	SrcRow         int
	SrcCol         int
	DstRow         int
	DstCol         int
	IsPromotion    bool
	PromotionPiece Piece
}

func getColumnFromColumnName(columnName string) int {
	return 104 - int(columnName[0])
}

func getColumnName(column int) string {
	return string(rune(104 - column))
}

func getRowFromRowName(rowName string) int {
	i, _ := strconv.ParseInt(rowName, 10, 32)
	return int(i) - 1
}

func getRowName(row int) string {
	return strconv.FormatInt(int64(1+row), 10)
}

func getPieceFromPieceName(pieceName string) Piece {
	switch pieceName {
	case "K":
		return King
	case "N":
		return Knight
	case "B":
		return Bishop
	case "Q":
		return Queen
	case "R":
		return Rook
	default:
		return 0
	}
}

func replaceAllInSet(str string, killSet string) string {
	ret := ""
	for i := 0; i < len(str); i += 1 {
		if !strings.ContainsRune(killSet, rune(str[i])) {
			ret += string(rune(str[i]))
		}
	}
	return ret
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

func CreateMoveFromPrettyMove(state *GameState, prettyMove string) Move {
	if prettyMove == "O-O" {
		srcCol := 3
		srcRow := 0
		if !state.IsWhiteTurn {
			srcRow = 7
		}

		return Move {
			SrcRow: srcRow,
			SrcCol: srcCol,
			DstRow: srcRow,
			DstCol: 1,
			IsPromotion: false,
			PromotionPiece: 0,
		}
	} else if prettyMove == "O-O-O" {
		srcCol := 3
		srcRow := 0
		if !state.IsWhiteTurn {
			srcRow = 7
		}

		return Move {
			SrcRow: srcRow,
			SrcCol: srcCol,
			DstRow: srcRow,
			DstCol: 5,
			IsPromotion: false,
			PromotionPiece: 0,
		}
	}

	var piece Piece
	var side Side
	if state.IsWhiteTurn {
		side = White
	} else {
		side = Black
	}
	var srcRow int
	var srcCol int
	var dstRow int
	var dstCol int
	promotionPiece := uint8(0)
	// Trim check, mate, and capture symbols (we can figure that out ourselves)
	prettyMove = replaceAllInSet(prettyMove, "+#x")
	// Retrieve and then discard promotion information
	if strings.Contains(prettyMove, "=") {
		split := strings.Split(prettyMove, "=")
		promotionPiece = getPieceFromPieceName(string(rune(split[1][1])))
		prettyMove = split[0]
	}

	if unicode.IsLower(rune(prettyMove[0])) {
		piece = Pawn
		col := getColumnFromColumnName(string(rune(prettyMove[0])))
		srcCol = col;
		if len(prettyMove) == 2 {
			dstCol = col
			dstRow = getRowFromRowName(string(rune(prettyMove[1])))

			// Should only ever be 1.
			possibleMovers := state.piecesInRange(piece, side, dstRow, dstCol)
			srcRow = possibleMovers[0].Row
		} else {
			dstCol = getColumnFromColumnName(string(rune(prettyMove[1])))
			dstRow = getRowFromRowName(string(rune(prettyMove[2])))

			possibleMovers := state.piecesInRange(piece, side, dstRow, dstCol)
			var mover BoardCoord
			for i := 0; i < len(possibleMovers); i += 1 {
				if possibleMovers[i].Col == srcCol {
					mover = possibleMovers[i]
				}
			}
			srcRow = mover.Row
		}
	} else {
		piece = getPieceFromPieceName(string(rune(prettyMove[0])))
		dstCol = getColumnFromColumnName(string(rune(prettyMove[len(prettyMove) - 2])))
		dstRow = getRowFromRowName(string(rune(prettyMove[len(prettyMove) - 1])))
		possibleMovers := state.piecesInRange(piece, side, dstRow, dstCol)
		if len(possibleMovers) == 1 {
			srcCol = possibleMovers[0].Col
			srcRow = possibleMovers[0].Row
		} else {
			// Check if all have the same rows/cols
			var rowOccurences [8]int
			var colOccurences [8]int
			for i := 0; i < len(possibleMovers); i += 1 {
				mover := possibleMovers[i]
				rowOccurences[mover.Row] += 1
				colOccurences[mover.Col] += 1 
			}

			if unicode.IsLetter(rune(prettyMove[1])) {
				srcCol = getColumnFromColumnName(string(rune(prettyMove[1])))
				if colOccurences[srcCol] == 1 {
					for i := 0; i < len(possibleMovers); i += 1 {
						if possibleMovers[i].Col == srcCol {
							srcRow = possibleMovers[i].Row
						}
					}
				} else {
					srcRow = getRowFromRowName(string(rune(prettyMove[2])))
				}
			} else {
				srcRow = getRowFromRowName(string(rune(prettyMove[1])))
				for i := 0; i < len(possibleMovers); i += 1 {
					if possibleMovers[i].Row == srcRow {
						srcCol = possibleMovers[i].Col
					}
				}
			}
		}
	}

	return Move {
		SrcRow: srcRow,
		SrcCol: srcCol,
		DstRow: dstRow,
		DstCol: dstCol,
		PromotionPiece: promotionPiece,
	}
}

func (move *Move) PrettyPrint(state *GameState) string {
	var builder strings.Builder
	piece := state.Board.PieceAt(move.SrcRow, move.SrcCol)
	side := state.Board.SideAt(move.SrcRow, move.SrcCol)
	possibleMovers := state.piecesInRange(piece, side, move.SrcRow, move.SrcCol)

	if state.moveIsCastlesShort(*move) {
		return "O-O"
	} else if state.moveIsCastlesLong(*move) {
		return "O-O-O"
	}

	var sameRowCount int
	var sameColCount int
	for i := 0; i < len(possibleMovers); i += 1 {
		if possibleMovers[i].Row == move.SrcRow {
			sameRowCount += 1
		}
		if possibleMovers[i].Col == move.SrcCol {
			sameColCount += 1
		}
	}

	dstColumnName := getColumnName(move.DstCol)
	dstRowName := getRowName(move.DstRow)

	pieceName := getPieceName(piece)

	if piece != Pawn {
		builder.WriteString(pieceName)
	}

	if sameColCount > 1 || piece == Pawn {
		builder.WriteString(getColumnName(move.SrcCol))
	}
	if sameRowCount > 1 {
		builder.WriteString(getRowName(move.SrcRow))
	}

	if state.Board.PieceAt(move.DstRow, move.DstCol) != Empty {
		builder.WriteString("x")
	}

	if piece != Pawn || state.Board.PieceAt(move.DstRow, move.SrcCol) != Empty {
		builder.WriteString(dstColumnName)
	}
	builder.WriteString(dstRowName)

	if move.IsPromotion {
		promotionPieceName := getPieceName(move.PromotionPiece)
		builder.WriteString("=")
		builder.WriteString(promotionPieceName)
	}

	return builder.String()
}
