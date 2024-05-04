package constant

import "errors"

var ErrInsertDatabase error = errors.New("invalid Add Data in Database")
var ErrInvalidRequest error = errors.New("invalid Request")
var ErrEmptyInput error = errors.New("input cannot be empty")
var ErrDuplicatedData error = errors.New("duplicated data")
var ErrNotFound error = errors.New("not found")
var ErrInvalidEmailOrPassword error = errors.New("invalid email or password")
var ErrNotAuthorized error = errors.New("not authorized")
var ErrInternalServer error = errors.New("internal server error")
