package router

type Response struct {
	StatusCode int
	Body       []byte
}

func (r *Response) Write(p []byte) (n int, err error) {
	r.Body = append(r.Body, p...)
	return len(r.Body), nil
}
