package server

type ErrorBadRegistrationRequest struct {}
type ErrorBadAuthenticationRequest struct {}

func (*ErrorBadRegistrationRequest) Error() string {
	return "error: no credentials and/or email provided"
}

func (*ErrorBadAuthenticationRequest) Error() string {
	return "error: bad authentication request"
}

// func (*ErrorBadRequest) Error() string {
// 	return "error: no credentials and/or email provided"
// }