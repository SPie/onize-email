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
    parser ParserContract
}

func NewEmailHandler(sender string, address string, port string, authUser AuthUserContract, parser ParserContract) EmailHandlerContract {
    return EmailHandler{sender, address, port, authUser, parser}
}

func (emailHandler EmailHandler) SendEmail(identifier string, message Message) error {
    text, err := emailHandler.parser.Parse(identifier, message.GetData())
    if err != nil {
	return err
    }

    address := fmt.Sprintf("%s:%s", emailHandler.address, emailHandler.port)
    err = smtp.SendMail(address, emailHandler.authUser.GetSMTPAuth(), "", []string{message.GetRecipient()}, text)
    if err != nil {
	return err
    }

    return nil
}
