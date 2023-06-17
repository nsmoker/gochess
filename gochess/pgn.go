package gochess

import (
	"strings"
	"fmt"
)

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