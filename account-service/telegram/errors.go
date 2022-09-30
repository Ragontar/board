package telegram

type ErrorSessionIsAlreadyRegistered struct{}

func (*ErrorSessionIsAlreadyRegistered) Error() string {
	return "error: session is already registered"
}