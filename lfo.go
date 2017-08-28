package fx

import (
	"fmt"
)

type LFOModulationFreqOrSync int
type LFOWaveShape int

const (
	ModFreq LFOModulationFreqOrSync = iota
	ModSync
)

const (
	Sine LFOWaveShape = iota
	Square
	Sawtooth
)

type Mapper interface {
	Map(to *FX)
}

type FX struct {
	status bool

	Mapping Mapper
}

type FXToggler interface {
	Enable() *FX
	Disable() *FX
}

type Dial struct {
	input *FX
}

func (e *FX) Enable() *FX {
	e.status = true
	return e
}

func (e *FX) Disable() *FX {
	e.status = false
	return e
}

func (e *FX) String() string { return fmt.Sprintf("FX: %t", e.status) }

func (e *FX) Map(fx *FX) {
	e.Mapping = fx
}

type AudioFX interface {
	FXToggler
	fmt.Stringer
}

type LFO struct {
	FX

	Offset  float64
	ModType LFOModulationFreqOrSync
	Rate    float64
	Depth   float64
	Phase   float64
	Smooth  float64
	Hold    bool
	Shape   LFOWaveShape

	Mapping AudioFX
}

func (osc *LFO) Map(fx *FX) {
	osc.Mapping = fx
}

func (l LFO) String() string {
	return "LFO"
}
