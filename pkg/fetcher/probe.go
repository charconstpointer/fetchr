package fetcher

type Probe struct {
	id       int
	url      string
	interval int
}

func NewProbe(id int, url string, interval int) *Probe {
	return &Probe{
		id:       id,
		url:      url,
		interval: interval,
	}
}
