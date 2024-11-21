package pkg

type RedirectService interface {
	Get(item string) (*RedirectModel, error)
	Add(item *RedirectModel) error
}
