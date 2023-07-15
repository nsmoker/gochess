package gochess

import (
	"bufio"
	//"fmt"
	"strings"
	"unicode"
)

type StateLst struct {
	State GameState
	Comment string
	Next []*StateLst
}

type DisplayList struct {
	Fen string
	Comment string
	Next []*DisplayList
}

func DisplayStateList(stateList *StateLst) *PgnDisplay {
	var nextList []*PgnDisplay
	for i := 0; i < len(stateList.Next); i += 1 {
		nextList = append(nextList, DisplayStateList(stateList.Next[i]))
	}
	var ret PgnDisplay
	ret.Fen = ToFEN(&stateList.State)
	ret.Comment = stateList.Comment
	ret.Next = nextList

	return &ret
}

// It is best for us to handle most of PGN parsing, it requires too much rule awareness for the GUI to be reasonably expected to do it.
func ParsePgn(pgn string) StateLst {
	scanner := bufio.NewScanner(strings.NewReader(pgn))
	inMove := false
	inCurlyBrace := false
	encounteredNumber := false
	move := ""
	rootState := StateLst {
		State: MakeStartingState(),
		Comment: "",
		Next: []*StateLst{},
	}
	currentState := &rootState
	//stack := []StateLst{}
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

				if encounteredNumber && c == '.' {
					encounteredNumber = false
				}

				if !inCurlyBrace && !encounteredNumber {
					if c == ' ' {
						if inMove {
							inMove = false
							mv := CreateMoveFromPrettyMove(&currentState.State, move)
							newState, _ := currentState.State.TakeTurn(mv)
							sl := StateLst {
								State: newState,
								Comment: "",
								Next: []*StateLst{},
							}
							currentState.Next = append(currentState.Next, &sl)
							currentState = &sl
							//fmt.Println(newState.Board.PrettyPrint())
							move = ""
						} 
					} else if !inMove && unicode.IsLetter(rune(c)) {
						move += string(c)
						inMove = true
					} else if inMove && (unicode.IsLetter(rune(c)) || unicode.IsNumber(rune(c)) || c == '=' || c == '-') {
						move += string(c)
					} else if !inMove && unicode.IsNumber(rune(c)) {
						encounteredNumber = true
					}
				}
			}
		}
	}

	return rootState
}