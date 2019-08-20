package email

import (
    "fmt"
    "net/smtp"
)

type EmailHandlerContract interface {
    SendEmail(identifier string, message Message) error
}

type EmailHandler struct {
    sender string
    address string
    port string
    authUser AuthUserContract
}

func NewEmailHandler(sender string, address string, port string, authUser AuthUserContract) EmailHandlerContract {
    return EmailHandler{sender, address, port, authUser}
}

func (emailHandler EmailHandler) SendEmail(identifier string, message Message) error {
    address := fmt.Sprintf("%s:%s", emailHandler.address, emailHandler.port)
    err := smtp.SendMail(address, emailHandler.authUser.GetSMTPAuth(), "", []string{message.GetRecipient()}, []byte(message.GetContent()))
    if err != nil {
	return err
    }

    return nil
}
