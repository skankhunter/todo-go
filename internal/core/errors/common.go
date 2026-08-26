package core_errors

import "errors"

var (
	ErrNotFound         = errors.New("not found")
	ErrnInvalidArgument = errors.New("ivalid argument")
	ErrnConflict        = errors.New("conflict")
)
