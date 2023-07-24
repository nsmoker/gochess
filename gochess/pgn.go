package gochess

import (
	"bufio"
	"strings"
	"unicode"
)

type StateLst struct {
	State GameState
	Comments []string
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
	ret.Comments = stateList.Comments
	ret.WhitesTurn = stateList.State.IsWhiteTurn
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
	inCommentTagPair := false
	var move strings.Builder
	var comment strings.Builder
	rootState := StateLst {
		State: MakeStartingState(),
		Comments: []string{},
		Next: []*StateLst{},
	}
	currentState := &rootState
	var prev *StateLst
	prev = nil
	stack := []*StateLst{}
	for scanner.Scan() {
		line := scanner.Text()
		// We don't care about tag pairs, those are trivial for the GUI to parse and display. We want the move text.
		// At least for now, we don't care about anything but getting all variations and comments through. 
		if !strings.HasPrefix(strings.TrimSpace(line), "[") {
			for i := 0; i < len(line); i += 1 {
				c := line[i]
				if c == '[' {
					inCommentTagPair = true
				} else if c == ']' {
					inCommentTagPair = false
					continue
				}
				if c == '{' {
					inCurlyBrace = true
					continue
				} else if c == '}' {
					inCurlyBrace = false
					currentState.Comments = append(currentState.Comments, comment.String())
					comment.Reset()
				}

				if c == '(' {
					stack = append(stack, currentState)
					currentState = prev
				}

				if !inCurlyBrace {
					if c == ' ' || !isMoveText(c) {
						if inMove {
							inMove = false
							mv := CreateMoveFromPrettyMove(&currentState.State, move.String())
							newState, _ := currentState.State.TakeTurn(mv)
							sl := StateLst {
								State: newState,
								Comments: []string{},
								Algebraic: move.String(),
								Next: []*StateLst{},
							}
							currentState.Next = append(currentState.Next, &sl)
							prev = currentState
							currentState = &sl
							move.Reset()
						} 
					} else if !inMove && unicode.IsLetter(rune(c)) {
						move.WriteByte(c)
						inMove = true
					} else if inMove && isMoveText(c) {
						move.WriteByte(c)
					}
				} else if !inCommentTagPair {
					comment.WriteByte(c)
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