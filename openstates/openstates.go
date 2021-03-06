package openstates

import (
	"github.com/mylez/go-sunlight/internal"
)

// Open States API root URL.
var openstatesRoot string = "http://openstates.org/api/v1"

/**
 * Open States data Source object. Present on nearly all objects in the
 * Open States database.
 */
type Source struct {
	Url string
}

/**
 */
type Sources struct {
	Sources []Source `json:"sources"`
}

/**
 */
type Timestamps struct {
	UpdatedAt internal.Time `json:"updated_at"`
	CreatedAt internal.Time `json:"created_at"`
}
