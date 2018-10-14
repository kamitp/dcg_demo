package interfaces

type TokenAuthInterface interface {
	GetTokenStatusFromServer(token string) error
}
