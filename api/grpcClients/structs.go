package grpcclients


type ErrorBadInsertData struct{}
type ErrorAlreadyRegistered struct{}
type ErrorSecretExpiredOrAbsent struct{}

func (e *ErrorBadInsertData) Error() string {
	return "error: corrupted or absent data for insertion"
}

func (e *ErrorAlreadyRegistered) Error() string {
	return "error: such uid/email/credentials already exists"
}

func (e *ErrorSecretExpiredOrAbsent) Error() string {
	return "error: session expired or absent"
}
