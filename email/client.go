package email

type ClientContract interface {
    Auth(authUser AuthUserContract) error
}

type Client struct {}
