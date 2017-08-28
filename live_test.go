package fx

import (
	"fmt"
	"testing"
)

func TestLFO(t *testing.T) {
	lfo := &LFO{}
	lfo.ModType = ModSync
	lfo.Rate = 1 / 4.0
	lfo.Hold = true
	lfo.Shape = Sawtooth

	fmt.Printf("%s = %#v\n", lfo, lfo)
}
