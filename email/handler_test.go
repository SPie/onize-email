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

func TestCreateNewEmailHandler(t *testing.T) {
    emailHandler := NewEmailHandler(AuthUserMock{})

    assert.Implements(t, new(EmailHandlerContract), emailHandler)
}
