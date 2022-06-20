package entity

import "errors"

//ErrNotFound not found
var ErrNotFound = errors.New("not found")

//ErrEmailAlreadyRegistred email already in DB
var ErrEmailAlreadyRegistred = errors.New("email already registred")

//ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("invalid entity")

//ErrCannotBeCreated cannot be created
var ErrCannotBeCreated = errors.New("cannot be created")

//ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("cannot be Deleted")

//ErrCannotBeUpdated cannot be deleted
var ErrCannotBeUpdated = errors.New("cannot be updated")

//ErrCannotBeDeleted cannot be deleted
var ErrPasswordGenaration = errors.New("pssword cannot be generated")
