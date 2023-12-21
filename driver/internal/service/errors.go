package service

import "errors"

var (
	ErrEntityNotFound     = errors.New("no such entity")
	ErrEntityAlreadyExist = errors.New("entity already exist")
)
