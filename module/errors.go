package module

import "errors"

var (
	INTERNAL_SERVER_ERROR = errors.New("Internal Server Error")
	NOT_FOUND_ERROR       = errors.New("Your requested Item is not found")
	CONFLIT_ERROR         = errors.New("Your Item already exist")
	EMPTY_ERROR           = errors.New("Empty data")
	FAILED_SAVE_ERROR     = errors.New("Failed save data")
	FAILED_UPDATE_ERROR   = errors.New("Failed update data")
	FAILED_GENERATE_ID    = errors.New("Failed generate id")
)
