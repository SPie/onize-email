package email

import (
    "net/smtp"
)

type AuthUserContract interface {
    GetSMTPAuth() smtp.Auth
}

type AuthUser struct {
    username string
    password string
    host string
}

func NewAuthUser(username string, password string, host string) AuthUserContract {
    return AuthUser{username, password, host}
}

func (authUser AuthUser) GetSMTPAuth() smtp.Auth {
    return smtp.PlainAuth("", authUser.username, authUser.password, authUser.host)
}
