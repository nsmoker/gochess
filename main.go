package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"syscall"

	"github.com/nsmoker/gochess/gochess"
	"google.golang.org/protobuf/proto"
)

func listenOnUrl(waitGroup *sync.WaitGroup, url string, handler func(net.Conn)) {
	waitGroup.Add(1)
	err := syscall.Unlink(url)
	if err != nil {
		log.Printf("couldn't unlink %s: %s", url, err)
	}

	listener, err := net.Listen("unix", url)
	if err != nil {
		log.Fatal(err)
	} else {
		go func() {
			defer listener.Close()
			for {
				conn, err := listener.Accept()

				if err != nil {
					log.Fatal(err)
				} else {
					go handler(conn)
				}
			}
		}()
	}
}

func checkLegal(conn net.Conn) {
	buf := make([]byte, 1024)
	conn.Read(buf)
	moveAndPos := gochess.MoveInPosition{}
	proto.Unmarshal(buf, &moveAndPos)

	position, err := gochess.ParseFEN(moveAndPos.Position.Fen)
	// 07/14/2023 proto.Unmarshal throws invalid wire format errors even on valid wire formats that unmarshal successfully, so we need to do our own error checking.
	if err != nil {
		conn.Write([]byte(fmt.Sprintf("Couldn't unmarshal message: %v, error %v", string(buf), err)))
		conn.Close()
		return
	}
	log.Printf("From position: %s\n", position.Board.PrettyPrint())
	move := gochess.Move{
		SrcRow:         int(moveAndPos.Move.From.Row),
		SrcCol:         int(moveAndPos.Move.From.Col),
		DstRow:         int(moveAndPos.Move.To.Row),
		DstCol:         int(moveAndPos.Move.To.Col),
		IsPromotion:    false,
		PromotionPiece: 0,
	}
	log.Println(move)
	if err == nil {
		isLegal := position.IsMoveLegal(move)
		log.Println(isLegal)
		var legalMsg gochess.MoveLegal
		var posNew gochess.Position
		legalMsg.Legal = isLegal
		legalMsg.PrettyMove = move.PrettyPrint(&position)
		sNew, legal := position.TakeTurn(move)
		posNew.Fen = gochess.ToFEN(&sNew)
		legalMsg.Position = &posNew
		legalMsg.Legal = legal
		marshalled, err := proto.Marshal(&legalMsg)
		if err != nil {
			log.Println(err)
		}
		_, err = conn.Write(marshalled)

		if err != nil {
			log.Println(err)
		}
		conn.Close()
		log.Printf("Position after: %s\n", position.Board.PrettyPrint())
	} else {
		log.Println(err)
	}
}

func main() {
	var wg sync.WaitGroup
	listenOnUrl(&wg, "/tmp/checkLegal.sock", checkLegal)

	wg.Wait()
}
