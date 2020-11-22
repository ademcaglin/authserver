package models

type Client struct {
	ClientId     string
	Name         string
	RedirectUris []string
	ApiResources []ApiResource
	IsActive     bool
}
