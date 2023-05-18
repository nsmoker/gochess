package gochess

import "testing"

func TestStartingBoard(t *testing.T) {
    output := MakeStartingBoard()
    want := MailboxBoard{Pieces: [64]Piece{Rook, Knight, Bishop, King, Queen, Bishop, Knight, Rook,
        Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn,
        Rook, Knight, Bishop, King, Queen, Bishop, Knight, Rook},
        Sides: [64]Side{White, White, White, White, White, White, White, White,
            White, White, White, White, White, White, White, White,
            Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
            Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
            Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
            Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
            Black, Black, Black, Black, Black, Black, Black, Black,
            Black, Black, Black, Black, Black, Black, Black, Black}}

    if output.Pieces != want.Pieces || output.Sides != want.Sides {
        t.Errorf("Wanted %s, got %s", want.PrettyPrint(), output.PrettyPrint())
    }
}

func TestPlacePiece(t *testing.T) {
    board := MakeStartingBoard()
    board.PlacePiece(0, 0, Knight, White)

    want := MailboxBoard{Pieces: [64]Piece{Knight, Knight, Bishop, King, Queen, Bishop, Knight, Rook,
        Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn,
        Rook, Knight, Bishop, King, Queen, Bishop, Knight, Rook},
        Sides: [64]Side{White, White, White, White, White, White, White, White,
            White, White, White, White, White, White, White, White,
            Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
            Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
            Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
            Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
            Black, Black, Black, Black, Black, Black, Black, Black,
            Black, Black, Black, Black, Black, Black, Black, Black}}

    if board.Pieces[0] != Knight || board.Sides[0] != White {
        t.Errorf("Expected %s, got %s", board.PrettyPrint(), want.PrettyPrint())
    }
}

func TestRemovePiece(t *testing.T) {
    board := MakeStartingBoard()
    board.RemovePiece(0, 0)

    want := MailboxBoard{Pieces: [64]Piece{Empty, Knight, Bishop, King, Queen, Bishop, Knight, Rook,
        Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn,
        Rook, Knight, Bishop, King, Queen, Bishop, Knight, Rook},
        Sides: [64]Side{Empty, White, White, White, White, White, White, White,
            White, White, White, White, White, White, White, White,
            Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
            Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
            Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
            Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
            Black, Black, Black, Black, Black, Black, Black, Black,
            Black, Black, Black, Black, Black, Black, Black, Black}}

    if board.Pieces[0] != Empty || board.Sides[0] != Empty {
        t.Errorf("Expected %s, got %s", board.PrettyPrint(), want.PrettyPrint())
    }
}
