package entity

type Request struct {
	URL         string
	Requests    int
	Concurrency int
}

func NewRequest(url string, requests int, concurrency int) *Request {
	return &Request{
		URL:         url,
		Requests:    requests,
		Concurrency: concurrency,
	}
}
