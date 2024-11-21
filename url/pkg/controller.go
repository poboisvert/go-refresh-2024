package pkg

import (
	"errors"
	"time"

	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	validate "gopkg.in/dealancer/validate.v2"
)

// Define custom error messages for redirect operations.
var (
	ErrRedirectNotFound = errors.New("redirect not found") // Error when a redirect is not found.
	ErrRedirectInvalid  = errors.New("redirect invalid")   // Error when a redirect is invalid.
)

// redirectService struct implements the RedirectService interface.
type redirectService struct {
	item RedirectService // Holds a reference to the RedirectService for data operations.
}

// NewRedirectService initializes a new redirectService with the provided RedirectService.
func NewRedirectService(redirectRepo RedirectService) RedirectService {
	return &redirectService{
		item: redirectRepo, // Assign the provided repository to the service.
	}
}

// Get retrieves a Redirect by its unique code from the repository.
func (r *redirectService) Get(code string) (*RedirectModel, error) {
	return r.item.Get(code) // Call the Get method on the repository to fetch the redirect.
}

// Add validates the Redirect object, generates a unique code and timestamp, and stores it in the repository.
func (r *redirectService) Add(redirect *RedirectModel) error {
	// Validate the redirect object to ensure it meets the required criteria.
	if err := validate.Validate(redirect); err != nil {
		return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Add") // Wrap and return an error if validation fails.
	}
	redirect.Code = shortid.MustGenerate()       // Generate a unique short ID for the redirect using the shortid library.
	redirect.CreatedAt = time.Now().UTC().Unix() // Set the creation timestamp to the current time in UTC.
	return r.item.Add(redirect)                  // Store the validated redirect in the repository.
}
