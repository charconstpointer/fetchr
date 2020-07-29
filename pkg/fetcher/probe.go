package fetcher

type Probe struct {
	url      string
	interval int
}

func NewProbe(url string, interval int) *Probe {
	return &Probe{
		url:      url,
		interval: interval,
	}
}
