package pkg

type RedirectSerializer interface {
	Decode(item []byte) (*RedirectModel, error)
	Encode(item *RedirectModel) ([]byte, error)
}
