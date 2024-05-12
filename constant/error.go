package constant

import "errors"

var ErrInsertDatabase error = errors.New("invalid Add Data in Database")
var ErrInvalidRequest error = errors.New("invalid Request")
var ErrEmptyInput error = errors.New("input cannot be empty")
var ErrDuplicatedData error = errors.New("duplicated data")
var ErrNotFound error = errors.New("not found")
var ErrInvalidEmailOrPassword error = errors.New("invalid email or password")
var ErrNotAuthorized error = errors.New("not authorized")
var ErrJobAlreadyFull error = errors.New("job already full")
var ErrJobAlreadyDone error = errors.New("job already marked done")
var ErrJobAlreadyOnProgress error = errors.New("job already on progress")
var ErrJobStillOpened error = errors.New("job still opened, please mark on progress first")
var ErrJobAlreadyClosed error = errors.New("job already closed")
var ErrFailedUpdate error = errors.New("failed to update the data")
var ErrHelperAlreadyTakeTheJob error = errors.New("helper already take the job")
var ErrPaymentGateway error = errors.New("payment gateway error")
var ErrTokenNotFound error = errors.New("token not found")
var ErrTokenNotValid error = errors.New("token not valid")
