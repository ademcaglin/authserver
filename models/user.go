package models

import (
	"context"
	"time"
)

type CredentialType int

const (
	TPM CredentialType = iota
	Packed
)

type UserCredential struct {
	CredentialId   string
	Type           CredentialType
	PublicKey      string
	SignatureCount uint
	IsActive       bool
}

type UserKey struct {
	Code           string
	CodeTime       time.Time
	ClientId       string
	Username       string
	PublicKey      string
	LastAccessTime time.Time
	IsActive       bool
	IsRevoked      bool
}

type User struct {
	Username    string
	Name        string
	CreatedAt   time.Time
	Credentials []UserCredential
	IsActive    bool
}

type UserStore interface {
	GetOne(ctx context.Context, username string) (User, error)
	//GetOneByKey(ctx context.Context, key string) (UserPublicKey, error)
	//GetOneByCode(ctx context.Context, code string) (UserPublicKey, error)
	Save(ctx context.Context, username string, name string) error
}
