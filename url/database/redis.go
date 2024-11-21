// This package provides a Redis-based implementation of the RedirectManager interface.
// It allows for storing and retrieving redirect information in a Redis database.

package redis

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"url/pkg"
)

// redisRepository is a struct that holds a Redis client for interacting with the database.
type redisRepository struct {
	client *redis.Client
}

// newRedisClient initializes a new Redis client using the provided Redis URL.
// It parses the URL and checks the connection by sending a Ping command.
func newRedisClient(redisURL string) (*redis.Client, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err // Return an error if the URL parsing fails.
	}
	client := redis.NewClient(opts) // Create a new Redis client with the parsed options.
	_, err = client.Ping().Result() // Ping the Redis server to check connectivity.
	if err != nil {
		return nil, err // Return an error if the Ping command fails.
	}
	return client, nil // Return the initialized Redis client.
}

// NewRedisRepository creates a new instance of redisRepository.
// It initializes the Redis client and returns the repository instance.
func NewRedisRepository(redisURL string) (pkg.RedirectManager, error) {
	repo := &redisRepository{}              // Create a new redisRepository instance.
	client, err := newRedisClient(redisURL) // Initialize the Redis client.
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewRedisRepository") // Wrap and return the error if client creation fails.
	}
	repo.client = client // Assign the Redis client to the repository.
	return repo, nil     // Return the repository instance.
}

// generateKey constructs a Redis key for storing redirect information based on the provided code.
func (r *redisRepository) generateKey(code string) string {
	return fmt.Sprintf("redirect:%s", code) // Format the key as "redirect:<code>".
}

// Get retrieves a redirect from the Redis database using the provided code.
// It returns the redirect information or an error if not found.
func (r *redisRepository) Get(code string) (*pkg.RedirectModel, error) {
	redirect := &pkg.RedirectModel{}            // Create a new Redirect instance.
	key := r.generateKey(code)                  // Generate the Redis key for the redirect.
	data, err := r.client.HGetAll(key).Result() // Fetch all fields of the hash stored at the key.
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.Get") // Wrap and return the error if fetching fails.
	}
	if len(data) == 0 {
		return nil, errors.Wrap(pkg.ErrRedirectNotFound, "repository.Redirect.Get") // Return not found error if no data is returned.
	}
	// Parse the creation timestamp from the retrieved data.
	createdAt, err := strconv.ParseInt(data["created_at"], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.Get") // Wrap and return the error if parsing fails.
	}
	// Populate the Redirect instance with the retrieved data.
	redirect.Code = data["code"]
	redirect.URL = data["url"]
	redirect.CreatedAt = createdAt
	return redirect, nil // Return the populated Redirect instance.
}

// Add stores a new redirect in the Redis database.
// It takes a Redirect instance and saves its information in a Redis hash.
func (r *redisRepository) Add(redirect *pkg.RedirectModel) error {
	key := r.generateKey(redirect.Code) // Generate the Redis key for the redirect.
	// Prepare the data to be stored in the Redis hash.
	data := map[string]interface{}{
		"code":       redirect.Code,
		"url":        redirect.URL,
		"created_at": redirect.CreatedAt,
	}
	_, err := r.client.HMSet(key, data).Result() // Store the data in the Redis hash.
	if err != nil {
		return errors.Wrap(err, "repository.Redirect.Add") // Wrap and return the error if storing fails.
	}
	return nil // Return nil if the operation is successful.
}
