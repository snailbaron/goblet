package main

import "github.com/snailbaron/goblet/internal/sdl"

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func NewRGBColor(r, g, b uint8) Color {
	return Color{R: r, G: g, B: b, A: 255}
}

func NewRGBAColor(r, g, b, a uint8) Color {
	return Color{R: r, G: g, B: b, A: a}
}

func setDrawColor(rr *sdl.Renderer, color *Color) error {
	return rr.SetDrawColor(color.R, color.G, color.B, color.A)
}

func shrink(rect *sdl.FRect, value float32) sdl.FRect {
	if rect.W <= value || rect.H <= value {
		return *rect
	}

	return sdl.FRect{
		X: rect.X + value,
		Y: rect.Y + value,
		W: rect.W - value*2,
		H: rect.H - value*2,
	}
}

type Widget interface {
	Render(rr *sdl.Renderer) error
}

type Button struct {
	Position        sdl.FRect
	BorderColor     Color
	BorderWidth     float32
	BackgroundColor Color
}

func (b *Button) Render(rr *sdl.Renderer) error {
	if err := setDrawColor(rr, &b.BorderColor); err != nil {
		return err
	}
	if err := rr.FillRect(&b.Position); err != nil {
		return err
	}

	innerRect := shrink(&b.Position, b.BorderWidth)
	if err := setDrawColor(rr, &b.BackgroundColor); err != nil {
		return err
	}
	if err := rr.FillRect(&innerRect); err != nil {
		return err
	}

	return nil
}

type Screen struct {
	windowWidth     int
	windowHeight    int
	backgroundColor [4]uint8

	window   *sdl.Window
	renderer *sdl.Renderer

	widgets []Widget
}

func NewScreen(config *Config) *Screen {
	return &Screen{
		windowWidth:     config.WindowWidth,
		windowHeight:    config.WindowHeight,
		backgroundColor: config.BackgroundColor,
	}
}

func (s *Screen) Init() error {
	var err error

	s.window, err = sdl.CreateWindow(
		"Goblet", s.windowWidth, s.windowHeight, 0)
	if err != nil {
		return err
	}

	s.renderer, err = s.window.CreateRenderer()
	if err != nil {
		return err
	}

	return nil
}

func (s *Screen) Destroy() {
	if s.renderer != nil {
		s.renderer.Destroy()
		s.renderer = nil
	}
	if s.window != nil {
		s.window.Destroy()
		s.window = nil
	}
}

func (s *Screen) Update(delta float64) error {
	_ = delta
	return nil
}

func (s *Screen) Render() error {
	r, g, b, a := s.backgroundColor[0], s.backgroundColor[1],
		s.backgroundColor[2], s.backgroundColor[3]
	if err := s.renderer.SetDrawColor(r, g, b, a); err != nil {
		return err
	}
	if err := s.renderer.Clear(); err != nil {
		return err
	}

	for _, w := range s.widgets {
		if err := w.Render(s.renderer); err != nil {
			return err
		}
	}

	if err := s.renderer.Present(); err != nil {
		return err
	}

	return nil
}

func (s *Screen) AddWidget(w Widget) {
	s.widgets = append(s.widgets, w)
}
