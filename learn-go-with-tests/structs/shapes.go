package structs

import "math"

// [01]
/* func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

func Area(width, height float64) float64 {
	return width * height
} */

type Shape interface {
  Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
  Radius float64
}

type Triangle struct {
  Base float64
  Height float64
}

func (r Rectangle) Area() float64 {
  return r.Width * r.Height
}

func (c Circle) Area() float64 {
  return (c.Radius * c.Radius) * math.Pi
}

func (t Triangle) Area() float64 {
  return (t.Base * t.Height) * 0.5
}

// [02]
/* func Perimeter(rectangle Rectangle) float64 {
  return 2 * (rectangle.Width + rectangle.Height)
}

func (rectangle Rectangle) float64 {
  return rectangle.Width * rectangle.Height
} */

