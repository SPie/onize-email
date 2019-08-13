package email

import (
    "net/smtp"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

type AuthUserMock struct {
    mock.Mock
}

func (authUser AuthUserMock) GetSMTPAuth() smtp.Auth {
    return authUser.Called().Get(0).(smtp.Auth)
}

type ParserMock struct {
    mock.Mock
}

func (parser ParserMock) Parse(templateName string, data map[string]interface{}) ([]byte, error) {
    args := parser.Called(templateName, data)

    return args.Get(0).([]byte), args.Error(1)
}

func TestCreateNewEmailHandler(t *testing.T) {
    emailHandler := NewEmailHandler("example@domain.dev", "address", "1234", AuthUserMock{}, ParserMock{})

    assert.Implements(t, new(EmailHandlerContract), emailHandler)
}
