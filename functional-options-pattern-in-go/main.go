package functional_options_pattern_in_go

// doc https://halls-of-valhalla.org/beta/articles/functional-options-pattern-in-go,54/

type Connection struct {

}


type stuffClient struct {
	conn    Connection
	timeout int
	retries int
}

type (
	StuffClientOption func(*StuffClientOptions)
	StuffClientOptions struct {
		Retries int
		Timeout int
	}
)

func WithRetries(r int) StuffClientOption {
	return func(o *StuffClientOptions) {
		o.Retries = r
	}
}

func WithTimeout(t int) StuffClientOption {
	return func(o *StuffClientOptions) {
		o.Timeout = t
	}
}

var defaultStuffClientOptions = StuffClientOptions{
	Retries: 3,
	Timeout: 2,
}

func NewStuffClient(conn Connection,opts ...StuffClientOption) stuffClient {
	options := defaultStuffClientOptions
	for _, o := range opts {
		o(&options)
	}
	return stuffClient{
		conn: conn,
		timeout: options.Timeout,
		retries: options.Retries,
	}
}