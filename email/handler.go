package email

type EmailHandlerContract interface {
    SendEmail(identifier string, message Message) error
}

type EmailHandler struct {}

func NewEmailHandler() EmailHandlerContract {
    return EmailHandler{}
}

func (emailHandler EmailHandler) SendEmail(identifier string, message Message) error {
    // TODO
    return nil
}
