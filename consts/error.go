package consts

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrInternalServerError Error = "internal server error"
	ErrInvalidRequestBody  Error = "invalid request body"
	ErrInvalidRequest      Error = "invalid request"
)

const (
	ErrUserNotFound          Error = "user not found"
	ErrUserExists            Error = "user already exists"
	ErrInvalidCredentials    Error = "email or password is incorrect"
	ErrTokenGenerationFailed Error = "failed to generate token"
)

const (
	ErrDatabaseConnectionError Error = "failed to connect to database"
)
