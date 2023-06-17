package gochess

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)

func TestIsInCheckHorizontalVertical(t *testing.T) {
	boardVerticalRook := MailboxBoard{Pieces: [64]Piece{
		Rook, King, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		King, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		Sides: [64]Side{
			White, White, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Black, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty}}
	gamestate := MakeStartingState()
	gamestate.Board = boardVerticalRook
	if gamestate.IsInCheck(White) {
		t.Fatalf("White should not be in check for %s", boardVerticalRook.PrettyPrint())
	} else if !gamestate.IsInCheck(Black) {
		t.Fatalf("Black should be in check for %s", boardVerticalRook.PrettyPrint())
	}

	boardHorizontalRook := MailboxBoard{Pieces: [64]Piece{
		Rook, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, King, Empty, Empty, Empty, Rook,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, King},
		Sides: [64]Side{
			White, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, White, Empty, Empty, Empty, Black,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Black}}
	gamestate.Board = boardHorizontalRook
	if gamestate.IsInCheck(Black) {
		t.Fatalf("Black should not be in check for %s", boardVerticalRook.PrettyPrint())
	} else if !gamestate.IsInCheck(White) {
		t.Fatalf("White should be in check for %s", boardVerticalRook.PrettyPrint())
	}

	boardHorizontalRookBlocked := MailboxBoard{Pieces: [64]Piece{
		Rook, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, King, Empty, Pawn, Empty, Rook,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, King},
		Sides: [64]Side{
			White, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, White, Empty, White, Empty, Black,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Black}}
	gamestate.Board = boardHorizontalRookBlocked
	if gamestate.IsInCheck(Black) {
		t.Fatalf("Black should not be in check for %s", boardHorizontalRookBlocked.PrettyPrint())
	} else if gamestate.IsInCheck(White) {
		t.Fatalf("White should not be in check for %s", boardHorizontalRookBlocked.PrettyPrint())
	}

	boardHorizontalRookBlockedSameSide := MailboxBoard{Pieces: [64]Piece{
		Rook, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, King, Empty, Pawn, Empty, Rook,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, King},
		Sides: [64]Side{
			White, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, White, Empty, Black, Empty, Black,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Black}}

	gamestate.Board = boardHorizontalRookBlockedSameSide
	if gamestate.IsInCheck(Black) {
		t.Fatalf("Black should not be in check for %s", boardHorizontalRookBlocked.PrettyPrint())
	} else if gamestate.IsInCheck(White) {
		t.Fatalf("White should not be in check for %s", boardHorizontalRookBlocked.PrettyPrint())
	}
}

func TestIsInCheckDiagonal(t *testing.T) {
	boardUpperRightBishop := MailboxBoard{Pieces: [64]Piece{
		Empty, King, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Bishop, Empty, Empty, Empty, Empty, Empty, Empty,
		King, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		Sides: [64]Side{
			Empty, White, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, White, Empty, Empty, Empty, Empty, Empty, Empty,
			Black, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty}}

	gamestate := MakeStartingState()
	gamestate.Board = boardUpperRightBishop
	if gamestate.IsInCheck(White) {
		t.Fatalf("White should not be in check for %s", boardUpperRightBishop.PrettyPrint())
	} else if !gamestate.IsInCheck(Black) {
		t.Fatalf("Black should be in check for %s", boardUpperRightBishop.PrettyPrint())
	}

	boardUpperRightPawn := MailboxBoard{Pieces: [64]Piece{
		Empty, King, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Pawn, Empty, Empty, Empty, Empty, Empty, Empty,
		King, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		Sides: [64]Side{
			Empty, White, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, White, Empty, Empty, Empty, Empty, Empty, Empty,
			Black, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty}}

	gamestate.Board = boardUpperRightPawn
	if gamestate.IsInCheck(White) {
		t.Fatalf("White should not be in check for %s", boardUpperRightPawn.PrettyPrint())
	} else if !gamestate.IsInCheck(Black) {
		t.Fatalf("Black should be in check for %s", boardUpperRightPawn.PrettyPrint())
	}

	boardUpperLeftBishop := MailboxBoard{Pieces: [64]Piece{
		Empty, King, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Bishop, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, King, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		Sides: [64]Side{
			Empty, White, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			White, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Black, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty}}

	gamestate.Board = boardUpperLeftBishop
	if gamestate.IsInCheck(White) {
		t.Fatalf("White should not be in check for %s", boardUpperLeftBishop.PrettyPrint())
	} else if !gamestate.IsInCheck(Black) {
		t.Fatalf("Black should be in check for %s", boardUpperLeftBishop.PrettyPrint())
	}

	boardUpperLeftPawn := MailboxBoard{Pieces: [64]Piece{
		Empty, King, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Pawn, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, King, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		Sides: [64]Side{
			Empty, White, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			White, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Black, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty}}

	gamestate.Board = boardUpperLeftPawn
	if gamestate.IsInCheck(White) {
		t.Fatalf("White should not be in check for %s", boardUpperLeftPawn.PrettyPrint())
	} else if !gamestate.IsInCheck(Black) {
		t.Fatalf("Black should be in check for %s", boardUpperLeftPawn.PrettyPrint())
	}

	boardLowerRightBishop := MailboxBoard{Pieces: [64]Piece{
		Empty, King, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Bishop, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		King, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		Sides: [64]Side{
			Empty, White, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Black, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Black, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty}}

	gamestate.Board = boardLowerRightBishop
	if gamestate.IsInCheck(Black) {
		t.Fatalf("Black should not be in check for %s", boardLowerRightBishop.PrettyPrint())
	} else if !gamestate.IsInCheck(White) {
		t.Fatalf("White should be in check for %s", boardLowerRightBishop.PrettyPrint())
	}

	boardLowerRightPawn := MailboxBoard{Pieces: [64]Piece{
		Empty, King, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Pawn, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		King, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		Sides: [64]Side{
			Empty, White, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Black, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Black, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty}}

	gamestate.Board = boardLowerRightPawn
	if gamestate.IsInCheck(Black) {
		t.Fatalf("Black should not be in check for %s", boardLowerRightPawn.PrettyPrint())
	} else if !gamestate.IsInCheck(White) {
		t.Fatalf("White should be in check for %s", boardLowerRightPawn.PrettyPrint())
	}

	boardLowerLeftBishop := MailboxBoard{Pieces: [64]Piece{
		Empty, King, Empty, Empty, Empty, Empty, Empty, Empty,
		Bishop, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		King, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		Sides: [64]Side{
			Empty, White, Empty, Empty, Empty, Empty, Empty, Empty,
			Black, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Black, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty}}

	gamestate.Board = boardLowerLeftBishop
	if gamestate.IsInCheck(Black) {
		t.Fatalf("Black should not be in check for %s", boardLowerLeftBishop.PrettyPrint())
	} else if !gamestate.IsInCheck(White) {
		t.Fatalf("White should be in check for %s", boardLowerLeftBishop.PrettyPrint())
	}

	boardLowerLeftPawn := MailboxBoard{Pieces: [64]Piece{
		Empty, King, Empty, Empty, Empty, Empty, Empty, Empty,
		Pawn, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		King, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		Sides: [64]Side{
			Empty, White, Empty, Empty, Empty, Empty, Empty, Empty,
			Black, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Black, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty}}

	gamestate.Board = boardLowerLeftPawn
	if gamestate.IsInCheck(Black) {
		t.Fatalf("Black should not be in check for %s", boardLowerLeftPawn.PrettyPrint())
	} else if !gamestate.IsInCheck(White) {
		t.Fatalf("White should be in check for %s", boardLowerLeftPawn.PrettyPrint())
	}

	boardLowerRightBishopBlocked := MailboxBoard{Pieces: [64]Piece{
		Empty, King, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Pawn, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Bishop, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		King, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		Sides: [64]Side{
			Empty, White, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, White, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Black, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Black, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty}}

	gamestate.Board = boardLowerRightBishopBlocked
	if gamestate.IsInCheck(Black) {
		t.Fatalf("Black should not be in check for %s", boardLowerRightBishopBlocked.PrettyPrint())
	} else if gamestate.IsInCheck(White) {
		t.Fatalf("White should not be in check for %s", boardLowerRightBishopBlocked.PrettyPrint())
	}

	boardLowerRightPawnTooFar := MailboxBoard{Pieces: [64]Piece{
		Empty, King, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Pawn, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		King, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		Sides: [64]Side{
			Empty, White, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Black, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Black, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty}}

	gamestate.Board = boardLowerRightPawnTooFar
	if gamestate.IsInCheck(Black) {
		t.Fatalf("Black should not be in check for %s", boardLowerRightPawnTooFar.PrettyPrint())
	} else if gamestate.IsInCheck(White) {
		t.Fatalf("White should not be in check for %s", boardLowerRightPawnTooFar.PrettyPrint())
	}
}

func TestKnightCheck(t *testing.T) {
	gamestate := MakeStartingState()

	boardKnightCheck := MailboxBoard{Pieces: [64]Piece{
		Knight, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, King, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, King},
		Sides: [64]Side{
			White, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Black, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, White}}
	gamestate.Board = boardKnightCheck
	if gamestate.IsInCheck(White) {
		t.Fatalf("White should not be in check for %s", boardKnightCheck.PrettyPrint())
	} else if !gamestate.IsInCheck(Black) {
		t.Fatalf("Black should be in check for %s", boardKnightCheck.PrettyPrint())
	}
}

func TestIsMoveLegal(t *testing.T) {
	state := MakeStartingState()

	e4 := Move{Src_row: 1, Src_col: 3, Dest_row: 3, Dest_col: 3}
	if !state.IsMoveLegal(e4) {
		t.Fatalf("%s should be legal for %s", e4.PrettyPrint(&state), state.Board.PrettyPrint())
	}

	ke2 := Move{Src_row: 0, Src_col: 3, Dest_row: 1, Dest_col: 3}

	if state.IsMoveLegal(ke2) {
		t.Fatalf("%s should not be legal for %s", ke2.PrettyPrint(&state), state.Board.PrettyPrint())
	}

	nf3 := Move{Src_row: 0, Src_col: 1, Dest_row: 2, Dest_col: 2}

	if !state.IsMoveLegal(nf3) {
		t.Fatalf("%s should be legal for %s", nf3.PrettyPrint(&state), state.Board.PrettyPrint())
	}

	ne2 := Move{Src_row: 0, Src_col: 1, Dest_row: 1, Dest_col: 3}

	if state.IsMoveLegal(ne2) {
		t.Fatalf("%s should not be legal for %s", ne2.PrettyPrint(&state), state.Board.PrettyPrint())
	}

	state.Board.RemovePiece(1, 4)
	state.Board.RemovePiece(6, 4)

	kd2 := Move{Src_row: 0, Src_col: 3, Dest_row: 1, Dest_col: 4}
	if state.IsMoveLegal(kd2) {
		t.Fatalf("%s should not be legal for %s", kd2.PrettyPrint(&state), state.Board.PrettyPrint())
	}
}

func TestEnPassantLegal(t *testing.T) {
	state := MakeStartingState()

	state.Board.PlacePiece(3, 3, Pawn, White)
	state.Board.PlacePiece(3, 4, Pawn, Black)

	pm := Move{Src_row: 0, Src_col: 0, Dest_row: 3, Dest_col: 4}

	state.PreviousMove = &pm

	ep := Move{Src_row: 3, Src_col: 3, Dest_row: 4, Dest_col: 4}

	if !state.IsMoveLegal(ep) {
		t.Fatalf("%s should be legal for %s", ep.PrettyPrint(&state), state.Board.PrettyPrint())
	}

	state.PreviousMove = nil

	if state.IsMoveLegal(ep) {
		t.Fatalf("%s should not be legal for %s with no previous move", ep.PrettyPrint(&state), state.Board.PrettyPrint())
	}
}

func TestPawnMoveOccupied(t *testing.T) {
	state := MakeStartingState()

	state.Board.PlacePiece(3, 3, Pawn, Black)

	e4 := Move{Src_row: 1, Src_col: 3, Dest_row: 3, Dest_col: 3}

	if state.IsMoveLegal(e4) {
		t.Fatalf("%s should not be legal for %s", e4.PrettyPrint(&state), state.Board.PrettyPrint())
	}
}

func TestCanCastle(t *testing.T) {
	state := MakeStartingState()

	state.Board.RemovePiece(0, 1)
	state.Board.RemovePiece(0, 2)

	state.WhiteCanCastleKingside = true

	castleShort := Move{Src_row: 0, Src_col: 3, Dest_row: 0, Dest_col: 1}

	if !state.IsMoveLegal(castleShort) {
		t.Fatalf("%s should be legal for %s", castleShort.PrettyPrint(&state), state.Board.PrettyPrint())
	}
}

func TestCantCastle(t *testing.T) {
	state := MakeStartingState()

	state.Board.RemovePiece(0, 1)
	state.Board.RemovePiece(0, 2)
	state.WhiteCanCastleKingside = false

	castleShort := Move{Src_row: 0, Src_col: 3, Dest_row: 0, Dest_col: 1}

	if state.IsMoveLegal(castleShort) {
		t.Fatalf("%s should not be legal for %s when castling rights lost", castleShort.PrettyPrint(&state), state.Board.PrettyPrint())
	}
}

func TestCantCastleThroughPieces(t *testing.T) {
	state := MakeStartingState()

	castleShort := Move{Src_row: 0, Src_col: 3, Dest_row: 0, Dest_col: 1}

	if state.IsMoveLegal(castleShort) {
		t.Fatalf("%s should not be legal for %s", castleShort.PrettyPrint(&state), state.Board.PrettyPrint())
	}
}

func TestCanCastleQueenside(t *testing.T) {
	state := MakeStartingState()

	state.Board.RemovePiece(0, 4)
	state.Board.RemovePiece(0, 5)
	state.Board.RemovePiece(0, 6)

	castleLong := Move{Src_row: 0, Src_col: 3, Dest_row: 0, Dest_col: 5}

	if !state.IsMoveLegal(castleLong) {
		t.Fatalf("%s should be legal for %s", castleLong.PrettyPrint(&state), state.Board.PrettyPrint())
	}
}

func TestCantCastleThroughCheck(t *testing.T) {
	state := MakeStartingState()

	state.Board.RemovePiece(0, 4)
	state.Board.RemovePiece(0, 5)
	state.Board.RemovePiece(0, 6)
	state.Board.RemovePiece(1, 4)
	state.Board.RemovePiece(6, 4)

	castleLong := Move{Src_row: 0, Src_col: 3, Dest_row: 0, Dest_col: 5}

	if state.IsMoveLegal(castleLong) {
		t.Fatalf("%s should not be legal for %s", castleLong.PrettyPrint(&state), state.Board.PrettyPrint())
	}
}

func TestTakeTurnSimple(t *testing.T) {
	state := MakeStartingState()

	e4 := Move{Src_row: 1, Src_col: 3, Dest_row: 3, Dest_col: 3}

	state.TakeTurn(e4)

	want := MailboxBoard{Pieces: [64]Piece{Rook, Knight, Bishop, King, Queen, Bishop, Knight, Rook,
		Pawn, Pawn, Pawn, Empty, Pawn, Pawn, Pawn, Pawn,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Pawn, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn,
		Rook, Knight, Bishop, King, Queen, Bishop, Knight, Rook},
		Sides: [64]Side{White, White, White, White, White, White, White, White,
			White, White, White, Empty, White, White, White, White,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, White, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Black, Black, Black, Black, Black, Black, Black, Black,
			Black, Black, Black, Black, Black, Black, Black, Black}}

	if state.Board != want {
		t.Fatalf("Should get %s after %s. Got %s", want.PrettyPrint(), e4.PrettyPrint(&state), state.Board.PrettyPrint())
	}

	if *(state.PreviousMove) != e4 {
		t.Fatalf("Previous move should be %s after %s", e4.PrettyPrint(&state), e4.PrettyPrint(&state))
	}
}

func TestTakeTurnPromotion(t *testing.T) {
	state := MakeStartingState()

	state.Board.RemovePiece(7, 6)
	state.Board.PlacePiece(6, 6, Pawn, White)

	move := Move{Src_row: 6, Src_col: 6, Dest_row: 7, Dest_col: 6, IsPromotion: true, PromotionPiece: Queen}

	state.TakeTurn(move)

	want := MailboxBoard{Pieces: [64]Piece{Rook, Knight, Bishop, King, Queen, Bishop, Knight, Rook,
		Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Empty, Pawn,
		Rook, Knight, Bishop, King, Queen, Bishop, Queen, Rook},
		Sides: [64]Side{White, White, White, White, White, White, White, White,
			White, White, White, White, White, White, White, White,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Black, Black, Black, Black, Black, Black, Empty, Black,
			Black, Black, Black, Black, Black, Black, White, Black}}

	if !cmp.Equal(state.Board, want) {
		t.Log(cmp.Diff(state.Board, want))
		t.Fatalf("Got %s after %s, want %s", state.Board.PrettyPrint(), move.PrettyPrint(&state), want.PrettyPrint())
	}
}

func TestTakeTurnCastles(t *testing.T) {
	state := MakeStartingState()

	state.Board.RemovePiece(0, 1)
	state.Board.RemovePiece(0, 2)

	move := Move{Src_row: 0, Src_col: 3, Dest_row: 0, Dest_col: 1}

	state.TakeTurn(move)

	want := MailboxBoard{Pieces: [64]Piece{Empty, King, Rook, Empty, Queen, Bishop, Knight, Rook,
		Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn,
		Rook, Knight, Bishop, King, Queen, Bishop, Knight, Rook},
		Sides: [64]Side{Empty, White, White, Empty, White, White, White, White,
			White, White, White, White, White, White, White, White,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Black, Black, Black, Black, Black, Black, Black, Black,
			Black, Black, Black, Black, Black, Black, Black, Black}}

	if !cmp.Equal(state.Board, want) {
		t.Log(cmp.Diff(state.Board, want))
		t.Fatalf("Got %s after %s, want %s", state.Board.PrettyPrint(), move.PrettyPrint(&state), want.PrettyPrint())
	}

	if state.WhiteCanCastleKingside || state.WhiteCanCastleQueenside {
		t.Fatalf("White should not be able to castle in %s", state.Board.PrettyPrint())
	}
}

func TestTakeTurnEnPassant(t *testing.T) {
	state := MakeStartingState()

	state.Board.PlacePiece(3, 3, Pawn, White)
	state.Board.PlacePiece(3, 4, Pawn, Black)

	pm := Move{Src_row: 0, Src_col: 0, Dest_row: 3, Dest_col: 4}

	state.PreviousMove = &pm

	ep := Move{Src_row: 3, Src_col: 3, Dest_row: 4, Dest_col: 4}

	state.TakeTurn(ep)

	want := MailboxBoard{Pieces: [64]Piece{Rook, Knight, Bishop, King, Queen, Bishop, Knight, Rook,
		Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Pawn, Empty, Empty, Empty,
		Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
		Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn,
		Rook, Knight, Bishop, King, Queen, Bishop, Knight, Rook},
		Sides: [64]Side{White, White, White, White, White, White, White, White,
			White, White, White, White, White, White, White, White,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, White, Empty, Empty, Empty,
			Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
			Black, Black, Black, Black, Black, Black, Black, Black,
			Black, Black, Black, Black, Black, Black, Black, Black}}

	if !cmp.Equal(state.Board, want) {
		t.Log(cmp.Diff(state.Board, want))
		t.Fatalf("Got %s after %s, want %s", state.Board.PrettyPrint(), ep.PrettyPrint(&state), want.PrettyPrint())
	}
}