package email

type EmailHandlerContract interface {
    SendEmail(identifier string, message Message) error
}

type EmailHandler struct {
    authUser AuthUserContract
}

func NewEmailHandler(authUser AuthUserContract) EmailHandlerContract {
    return EmailHandler{authUser}
}

func (emailHandler EmailHandler) SendEmail(identifier string, message Message) error {
    // TODO
    return nil
}
