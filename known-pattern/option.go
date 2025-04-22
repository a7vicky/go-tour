package knownpattern

type Server struct {
	Addr string
	TLS  bool
}

type Option func(*Server)

func WithTLS(tls bool) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServer(addr string, opts ...Option) *Server {
	s := &Server{Addr: addr}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
