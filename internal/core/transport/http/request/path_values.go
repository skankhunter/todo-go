package core_http_request

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/skankhunter/todo-go/internal/core/errors"
)

// GET /user/123

func GetIntPathValues(r *http.Request, key string) (int, error) {
	pathValue := r.PathValue(key)

	if pathValue == "" {
		return 0, fmt.Errorf(
			"No key='%s' in path values: %w",
			key,
			core_errors.ErrnInvalidArgument,
		)
	}
	val, err := strconv.Atoi(pathValue)

	if err != nil {
		return 0, fmt.Errorf(
			"Path value='%s'by key='%s' not a valid integer: %v: %w",
			pathValue,
			key,
			err,
			core_errors.ErrnInvalidArgument,
		)

	}

	return val, nil
}
