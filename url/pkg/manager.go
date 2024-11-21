package pkg

type RedirectManager interface {
	Get(item string) (*RedirectModel, error)
	Add(item *RedirectModel) error
}
