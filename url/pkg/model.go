package pkg

type RedirectModel struct {
	Code      string `json:"code" msgpack:"code"`
	URL       string `json:"url" msgpack:"url" validate:"empty=false & format=url"` // validate the url format
	CreatedAt int64  `json:"created_at" msgpack:"created_at"`
}
