package gochess

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

type StateLst struct {
	State GameState
	Comment string
	Algebraic string
	Next []*StateLst
}

func DisplayStateList(stateList *StateLst) *PgnDisplay {
	var nextList []*PgnDisplay
	for i := 0; i < len(stateList.Next); i += 1 {
		nextList = append(nextList, DisplayStateList(stateList.Next[i]))
	}
	var ret PgnDisplay
	ret.Fen = ToFEN(&stateList.State)
	ret.Algebraic = stateList.Algebraic
	ret.Comment = stateList.Comment
	ret.Next = nextList

	return &ret
}

func isMoveText(c byte) bool {
	return unicode.IsLetter(rune(c)) || unicode.IsNumber(rune(c)) || c == '=' || c == '-'
}

// It is best for us to handle most of PGN parsing, it requires too much rule awareness for the GUI to be reasonably expected to do it.
func ParsePgn(pgn string) StateLst {
	scanner := bufio.NewScanner(strings.NewReader(pgn))
	inMove := false
	inCurlyBrace := false
	move := ""
	rootState := StateLst {
		State: MakeStartingState(),
		Comment: "",
		Next: []*StateLst{},
	}
	currentState := &rootState
	var prev *StateLst
	prev = nil
	stack := []*StateLst{}
	for scanner.Scan() {
		line := scanner.Text()
		// We don't care about tag pairs, those are trivial for the GUI to parse and display. We want the move text.
		// We parse move by move. At least for now, we don't care about anything but getting all variations. 
		if !strings.HasPrefix(strings.TrimSpace(line), "[") {
			for i := 0; i < len(line); i += 1 {
				c := line[i]
				if c == '{' {
					inCurlyBrace = true
				} else if c == '}' {
					inCurlyBrace = false
				}

				if c == '(' {
					stack = append(stack, currentState)
					currentState = prev
				}

				if !inCurlyBrace {
					if c == ' ' || !isMoveText(c) {
						if inMove {
							inMove = false
							mv := CreateMoveFromPrettyMove(&currentState.State, move)
							newState, made := currentState.State.TakeTurn(mv)
							if !made {
								fmt.Println(currentState.State.Board.PrettyPrint())
								fmt.Println(move)
							}
							sl := StateLst {
								State: newState,
								Comment: "",
								Algebraic: move,
								Next: []*StateLst{},
							}
							currentState.Next = append(currentState.Next, &sl)
							prev = currentState
							currentState = &sl
							move = ""
						} 
					} else if !inMove && unicode.IsLetter(rune(c)) {
						move += string(c)
						inMove = true
					} else if inMove && isMoveText(c) {
						move += string(c)
					}
				}

				if c == ')' {
					// Pop stack
					currentState = stack[len(stack) - 1]
					stack = stack[:len(stack) - 1]
				}
			}
		}
	}

	return rootState
}