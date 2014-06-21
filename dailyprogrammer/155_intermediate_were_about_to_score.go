package main

type Bitboard uint64

type Chessboard struct {
	wp Bitboard // White Pawns
	wb Bitboard // White Bishops
	wr Bitboard // White Rooks
	wn Bitboard // White Knights
	wk Bitboard // White King
	wq Bitboard // White Queen

	bp Bitboard // Black Pawns
	bb Bitboard // Black Bishops
	br Bitboard // Black Rooks
	bn Bitboard // Black Knights
	bk Bitboard // Black King
	bq Bitboard // Black Queen

	allw Bitboard // All White Pieces
	allb Bitboard // All Black Pieces
	allp Bitboard // All Pieces
}

func NewChessboard() *Chessboard {
	return &Chessboard{}
}

var w = map[string]rune{"wp": '♙', "wb": '♗', "wr": '♖', "wn": '♘', "wk": '♔', "wq": '♕'}
var b = map[string]rune{"bp": '♟', "bb": '♝', "br": '♜', "bn": '♞', "bk": '♚', "bq": '♛'}

var wp, wb, wr, wn, wk, wq Bitboard
var bp, bb, br, bn, bk, bq Bitboard

var allwhite = wp | wb | wr | wn | wk | wq
var allblack = bp | bb | br | bn | bk | bq
var allpieces = allwhite | allblack

func main() {

}
