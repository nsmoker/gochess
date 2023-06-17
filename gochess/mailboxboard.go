package gochess

import "strings"

const (
    King   = 0
    Queen  = 1
    Rook   = 2
    Bishop = 3
    Knight = 4
    Pawn   = 5
    White  = 0
    Black  = 1
    Empty  = 10
)

type Piece = uint8
type Side = uint8

type MailboxBoard struct {
    Pieces [64]Piece
    Sides  [64]Side
}

func (board *MailboxBoard) PlacePiece(row int, col int, piece Piece, side Side) {
    board.Pieces[row*8+col] = piece
    board.Sides[row*8+col] = side
}

func (board *MailboxBoard) RemovePiece(row int, col int) {
    board.Pieces[row*8+col] = Empty
    board.Sides[row*8+col] = Empty
}

func (board *MailboxBoard) PieceAt(row int, col int) Piece {
    return board.Pieces[row*8+col]
}

func (board *MailboxBoard) SideAt(row int, col int) Side {
    return board.Sides[row*8+col]
}

func MakeStartingBoard() MailboxBoard {
    pieces := [64]Piece{Rook, Knight, Bishop, King, Queen, Bishop, Knight, Rook,
        Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn,
        Rook, Knight, Bishop, King, Queen, Bishop, Knight, Rook}
    sides := [64]Side{White, White, White, White, White, White, White, White,
        White, White, White, White, White, White, White, White,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty,
        Black, Black, Black, Black, Black, Black, Black, Black,
        Black, Black, Black, Black, Black, Black, Black, Black}

    return MailboxBoard{pieces, sides}
}

func (board *MailboxBoard) PrettyPrint() string {
    var builder strings.Builder

    blackChars := []string{"♚", "♛", "♜", "♝", "♞", "♟︎"}
    whiteChars := []string{"♔", "♕", "♖", "♗", "♘", "♙"}

    builder.WriteString("\n")
    for i := 0; i < 8; i++ {
        for j := 0; j < 8; j++ {
            var charList []string
            if board.Sides[i*8+j] == White {
                charList = whiteChars
            } else {
                charList = blackChars
            }
            if board.Pieces[i*8+j] == Empty {
                builder.WriteString(" ")
            } else {
                builder.WriteString(charList[board.Pieces[i*8+j]])
            }
        }
        builder.WriteString("\n")
    }

    return builder.String()
}

// func (board *MailboxBoard) numberInColumn(column int, piece Piece, side Side) int {
//     numberInCol := 0
//     for row := 0; row < 8; row++ {
//         if board.PieceAt(row, column) == piece && board.SideAt(row, column) == side {
//             numberInCol += 1
//         }
//     }

//     return numberInCol
// }

// func (board *MailboxBoard) numberInRow(row int, piece Piece, side Side) int {
//     numberInRow := 0
//     for column := 0; column < 8; column++ {
//         if board.PieceAt(row, column) == piece && board.SideAt(row, column) == side {
//             numberInRow += 1
//         }
//     }

//     return numberInRow
// }