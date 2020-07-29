package fetcher

type probe interface {
}

type Probe struct {
}

func NewProbe() *Probe {
	return &Probe{}
}
