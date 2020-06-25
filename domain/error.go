package domain

import "errors"

//ErrNotFound not found
var ErrNotFound = errors.New("Not found")

//ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("Cannot Be Deleted")

//ErrNotEnoughBooks cannot borrow
var ErrNotEnoughBooks = errors.New("Not enough books")

//ErrBookAlreadyBorrowed cannot borrow
var ErrBookAlreadyBorrowed = errors.New("Book already borrowed")

//ErrBookNotBorrowed cannot return
var ErrBookNotBorrowed = errors.New("Book not borrowed")
