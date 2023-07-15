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
	Next []StateLst
}

// It is best for us to handle most of PGN parsing, it requires too much rule awareness for the GUI to be reasonably expected to do it.
func ParsePgn(pgn string) {
	scanner := bufio.NewScanner(strings.NewReader(pgn))
	inMove := false
	move := ""
	//state := MakeStartingState()
	for scanner.Scan() {
		line := scanner.Text()
		// We don't care about tag pairs, those are trivial for the GUI to parse and display. We want the move text.
		// We parse move by move. At least for now, we don't care about anything but getting all variations and comments. 
		if !strings.HasPrefix(strings.TrimSpace(line), "[") {
			for i := 0; i < len(line); i += 1 {
				c := line[i]
				if c == ' ' {
					if inMove {
						inMove = false
						fmt.Println(move)
						move = ""
					}
				} else if !inMove && unicode.IsLetter(rune(c)) {
					move += string(c)
					inMove = true
				} else if inMove {
					move += string(c)
				}
			}
		}
	}
}

func ToPGN(state *GameState) string {
	var builder strings.Builder
	states := []*GameState{}
	for currentState := state; currentState.PreviousState != nil; {
		states = append(states, currentState)
	}

	for i := len(states) - 1; i >= 0; i++ {
		state := states[i]
		moveNumber := len(states) - i
		builder.WriteString(fmt.Sprintf("%d. %s ", moveNumber, state.PreviousMove.PrettyPrint(state)))
	}

	return builder.String()
}