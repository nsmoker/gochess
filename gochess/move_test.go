package gochess

import (
	"testing"
)

func TestReadPrettyMovePawn(t *testing.T) {
	state := MakeStartingState()
	moveStr := "e4"
	want := Move {
		SrcRow: 1,
		SrcCol: 3,
		DstRow: 3,
		DstCol: 3,
		IsPromotion: false,
		PromotionPiece: 0,
	}
	have := CreateMoveFromPrettyMove(&state, moveStr)
	if want != have {
		t.Fatalf("parsed %s as %v, want %v", moveStr, have, want)
	}
}

func TestReadPrettyMoveKnight(t *testing.T) {
	state := MakeStartingState()
	moveStr := "Nf3"
	want := Move {
		SrcRow: 0,
		SrcCol: 1,
		DstRow: 2,
		DstCol: 2,
		IsPromotion: false,
		PromotionPiece: 0,
	}
	have := CreateMoveFromPrettyMove(&state, moveStr)
	if want != have {
		t.Fatalf("parsed %s as %v, want %v", moveStr, have, want)
	}
}

func TestReadPrettyMoveAmbiguous(t *testing.T) {
	state, _ := ParseFEN("k7/8/6N1/8/2N5/8/8/7K w - - 1 2")
	moveStr := "Nce5"
	want := Move {
		SrcRow: 3,
		SrcCol: 5,
		DstRow: 4,
		DstCol: 3,
		IsPromotion: false,
		PromotionPiece: 0,
	}
	have := CreateMoveFromPrettyMove(&state, moveStr)
	if want != have {
		t.Fatalf("parsed %s as %v, want %v", moveStr, have, want)
	}
}

// Ahead, the term "AmbiguousHellCase" refers to the following position: "k7/8/2N3N1/8/2N5/8/8/7K w - - 1 2"
// which you may paste into your editor of choice to see why it earns that name.

func TestReadPrettyMoveAmbiguousHellCase1(t *testing.T) {
	state, _ := ParseFEN("k7/8/2N3N1/8/2N5/8/8/7K w - - 1 2")
	moveStr := "Nge5"
	want := Move {
		SrcRow: 5,
		SrcCol: 1,
		DstRow: 4,
		DstCol: 3,
		IsPromotion: false,
		PromotionPiece: 0,
	}
	have := CreateMoveFromPrettyMove(&state, moveStr)
	if want != have {
		t.Fatalf("parsed %s as %v, want %v", moveStr, have, want)
	}
}

func TestReadPrettyMoveAmbiguousHellCase2(t *testing.T) {
	state, _ := ParseFEN("k7/8/2N3N1/8/2N5/8/8/7K w - - 1 2")
	moveStr := "N4e5"
	want := Move {
		SrcRow: 3,
		SrcCol: 5,
		DstRow: 4,
		DstCol: 3,
		IsPromotion: false,
		PromotionPiece: 0,
	}
	have := CreateMoveFromPrettyMove(&state, moveStr)
	if want != have {
		t.Fatalf("parsed %s as %v, want %v", moveStr, have, want)
	}
}

func TestReadPrettyMoveAmbiguousHellCase3(t *testing.T) {
	state, _ := ParseFEN("k7/8/2N3N1/8/2N5/8/8/7K w - - 1 2")
	moveStr := "Nc6e5"
	want := Move {
		SrcRow: 5,
		SrcCol: 5,
		DstRow: 4,
		DstCol: 3,
		IsPromotion: false,
		PromotionPiece: 0,
	}
	have := CreateMoveFromPrettyMove(&state, moveStr)
	if want != have {
		t.Fatalf("parsed %s as %v, want %v", moveStr, have, want)
	}
}

func TestReadPrettyMovePawnCapture(t *testing.T) {
	state, _ := ParseFEN("k7/8/8/3p4/2P1P3/8/8/7K w - - 1 2")
	moveStr := "exd5"
	want := Move {
		SrcRow: 3,
		SrcCol: 3,
		DstRow: 4,
		DstCol: 4,
		IsPromotion: false,
		PromotionPiece: 0,
	}
	have := CreateMoveFromPrettyMove(&state, moveStr)
	if want != have {
		t.Fatalf("parsed %s as %v, want %v", moveStr, have, want)
	}
}