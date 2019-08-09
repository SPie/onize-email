package email

import (
    "net/smtp"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestCreateNewAuthUser(t *testing.T) {
    authUser := NewAuthUser("Username", "Password", "Host")

    assert.Implements(t, new(AuthUserContract), authUser)
}

func TestGetSMTPAuth(t *testing.T) {
    authUser := AuthUser{"Username", "Password", "Host"}

    smtpAuth := authUser.GetSMTPAuth()

    assert.Implements(t, new(smtp.Auth), smtpAuth)
}
