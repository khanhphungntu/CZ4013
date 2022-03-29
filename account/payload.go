package account

type request interface {
	unmarshal(data []byte) error
}

type response interface {
	marshal() []byte
}
