package constant

import "errors"

var ErrInsertDatabase error = errors.New("invalid Add Data in Database")
var ErrInvalidRequest error = errors.New("invalid Request")
var ErrEmptyInput error = errors.New("title cannot be empty")
var ErrDuplicatedData error = errors.New("duplicated data")
