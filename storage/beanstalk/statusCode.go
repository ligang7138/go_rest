package beanstalk

import "errors"

var (
	ErrBadFormat  = errors.New("bad command format")
	ErrBuried     = errors.New("buried")
	ErrDeadline   = errors.New("deadline soon")
	ErrDraining   = errors.New("draining")
	ErrInternal   = errors.New("internal error")
	ErrJobTooBig  = errors.New("job too big")
	ErrNoCRLF     = errors.New("expected CR LF")
	ErrNotFound   = errors.New("not found")
	ErrNotIgnored = errors.New("not ignored")
	ErrOOM        = errors.New("server is out of memory")
	ErrTimeout    = errors.New("timeout")
	ErrUnknown    = errors.New("unknown command")
	ErrEmpty      = errors.New("name is empty")
	ErrBadChar    = errors.New("name has bad char") // contains a character not in NameChars
	ErrTooLong    = errors.New("name is too long")
)
