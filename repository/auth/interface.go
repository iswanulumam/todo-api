package auth

type Auth interface {
	Login(username string, password string) (string, error)
}
