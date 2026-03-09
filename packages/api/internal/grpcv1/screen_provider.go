package grpcv1

import (
	"context"
	"fmt"

	"github.com/smysnk/sikuligo/pkg/sikuli"
)

var listScreensFn = listScreens

func WithScreenLister(fn func(context.Context) ([]sikuli.Screen, error)) ServerOption {
	return func(s *Server) {
		s.listScreens = fn
	}
}

func (s *Server) screens(ctx context.Context) ([]sikuli.Screen, error) {
	if s != nil && s.listScreens != nil {
		return s.listScreens(ctx)
	}
	return listScreensFn(ctx)
}

func (s *Server) primaryScreen(ctx context.Context) (sikuli.Screen, error) {
	screens, err := s.screens(ctx)
	if err != nil {
		return sikuli.Screen{}, err
	}
	for _, screen := range screens {
		if screen.Primary {
			return screen, nil
		}
	}
	if len(screens) == 0 {
		return sikuli.Screen{}, fmt.Errorf("%w: no screens available", sikuli.ErrBackendUnsupported)
	}
	return screens[0], nil
}

func (s *Server) screenByID(ctx context.Context, id int) (sikuli.Screen, error) {
	screens, err := s.screens(ctx)
	if err != nil {
		return sikuli.Screen{}, err
	}
	for _, screen := range screens {
		if screen.ID == id {
			return screen, nil
		}
	}
	return sikuli.Screen{}, fmt.Errorf("%w: screen id %d not found", sikuli.ErrInvalidTarget, id)
}
