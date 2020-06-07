package test

import (
	"context"
	"errors"
)

type Server struct{}

var CustomErr = errors.New("custom error")

func (s *Server) Handle(context.Context, *Req) (*Res, error) {
	return nil, CustomErr
}
