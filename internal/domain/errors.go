package domain

import "errors"

var ErrRefreshTokenExpired = errors.New("refresh token expired")
var ErrRefreshTokenParse = errors.New("parse refresh token error")
var ErrRefreshToken = errors.New("refresh tokens error")

var ErrUserCredNotFound = errors.New("user with such credentials not found")
