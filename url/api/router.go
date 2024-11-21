// This package defines the API for handling redirect operations, including retrieving and adding redirects.

package api

import (
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"

	"url/pkg"
	js "url/serializer/json"
	ms "url/serializer/msgpack"
)

// RedirectHandler interface defines the methods for handling HTTP requests related to redirects.
type RedirectHandler interface {
	Get(http.ResponseWriter, *http.Request)  // Method to handle GET requests for retrieving a redirect.
	Post(http.ResponseWriter, *http.Request) // Method to handle POST requests for adding a new redirect.
}

// handler struct implements the RedirectHandler interface and holds a reference to the redirect service.
type handler struct {
	redirectService pkg.RedirectService // Service for managing redirects.
}

// NewHandler initializes a new handler with the provided redirect service.
func NewHandler(redirectService pkg.RedirectService) RedirectHandler {
	return &handler{redirectService: redirectService} // Return a new handler instance.
}

// setupResponse prepares the HTTP response with the specified content type, body, and status code.
func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType) // Set the Content-Type header.
	w.WriteHeader(statusCode)                   // Set the HTTP status code.
	if _, err := w.Write(body); err != nil {    // Write the response body and log any errors.
		log.Println(err)
	}
}

// serializer returns the appropriate serializer based on the content type of the request.
func (h *handler) serializer(contentType string) pkg.RedirectSerializer {
	if contentType == "application/x-msgpack" { // Check if the content type is MsgPack.
		return &ms.Redirect{} // Return MsgPack serializer.
	}
	return &js.Redirect{} // Default to JSON serializer.
}

// Get handles GET requests to retrieve a redirect based on the provided code.
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")              // Extract the redirect code from the URL parameters.
	redirect, err := h.redirectService.Get(code) // Retrieve the redirect using the service.
	if err != nil {                              // If an error occurs, handle it.
		h.handleError(w, err)
		return
	}
	http.Redirect(w, r, redirect.URL, http.StatusMovedPermanently) // Redirect the client to the stored URL.
}

// Post handles POST requests to add a new redirect.
func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type") // Get the Content-Type from the request header.
	requestBody, err := io.ReadAll(r.Body)      // Read the request body.
	if err != nil {                             // Handle any errors that occur while reading the body.
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	redirect, err := h.serializer(contentType).Decode(requestBody) // Decode the request body into a Redirect object.
	if err != nil {                                                // Handle any errors during decoding.
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err = h.redirectService.Add(redirect); err != nil { // Attempt to add the redirect using the service.
		h.handleError(w, err) // Handle any errors that occur during the addition.
		return
	}
	responseBody, err := h.serializer(contentType).Encode(redirect) // Encode the redirect object for the response.
	if err != nil {                                                 // Handle any errors during encoding.
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated) // Send the response with a 201 Created status.
}

// handleError processes errors and sends the appropriate HTTP response based on the error type.
func (h *handler) handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Cause(err) == pkg.ErrRedirectNotFound: // Check if the error is a redirect not found error.
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound) // Respond with 404 Not Found.
	case errors.Cause(err) == pkg.ErrRedirectInvalid: // Check if the error is a redirect invalid error.
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // Respond with 400 Bad Request.
	default: // For any other errors, respond with 500 Internal Server Error.
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
