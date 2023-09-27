package model

import "github.com/golang-jwt/jwt"

type UserRole string

const (
	ADMIN  UserRole = "admin"
	PUBLIC UserRole = "public"
)

type Auth struct {
	Id       string `json:"id"`
	Username string `json:"Username"`
	jwt.StandardClaims
}

type AuthResp struct {
	Token string `json:"token"`
}
