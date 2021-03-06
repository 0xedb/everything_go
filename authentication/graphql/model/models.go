// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Auth struct {
	User *User `json:"user"`
}

type Mine struct {
	Users *Auth    `json:"users"`
	Link  []string `json:"link"`
}

type Nothing struct {
	More string `json:"more"`
}

type User struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Etag string `json:"etag"`
}
