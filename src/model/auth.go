package model

import (
	"github.com/dgrijalva/jwt-go"
)

// BearerClaimsWithMobileNumber contains mobile number
type BearerClaimsWithMobileNumber struct {
	MobileNumber string `json:"mobileNumber"`
	jwt.StandardClaims
}

// BearerClaimsWithUsername contains username
type BearerClaimsWithUsername struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
