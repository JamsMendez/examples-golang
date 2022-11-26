package structs

import "testing"

// [01]
/* func TestPerimeter(t *testing.T) {
  got := Perimeter(10.0, 10.0)
  want := 40.0

  if got != want {
    t.Errorf("expected %.2f want %.2f", got, want)
  }
}

func TestArea(t *testing.T) {
  got := Area(12.0, 6.0)
  want := 72.0

  if got != want {
    t.Errorf("expected %.2f want %.2f", got, want)
  }
} */

// [02]
/* func TestPerimeter(t *testing.T) {
  rectangle := Rectangle{Width: 10.0, Height: 10.0}

  got := Perimeter(rectangle)
  want := 40.0

  if got != want {
    t.Errorf("expected %.2f want %.2f", got, want)
  }
}

func TestArea(t *testing.T) {
  rectangle := Rectangle{Width: 12.0, Height: 6.0}

  got := Area(rectangle)
  want := 72.0

  if got != want {
    t.Errorf("expected %.2f want %.2f", got, want)
  }
} */

// [03]
/* func TestArea(t *testing.T) {
  t.Run("rectangles", func(t *testing.T) {
    rectangle := Rectangle{Width: 12, Height: 6}

    got := rectangle.Area()
    want := 72.0

    if got != want {
    t.Errorf("expected %g want %g", got, want)
    }
  })

  t.Run("circles", func(t *testing.T) {
    circle := Circle{10}

    got := circle.Area()
    want := 314.1592653589793

    if got != want {
    t.Errorf("expected %g want %g", got, want)
    }
  })
} */

// [04]
/* func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()
		if got != want {
			t.Errorf("expected %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{Width: 12, Height: 6}
		want := 72.0

		checkArea(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793

		checkArea(t, circle, want)
	})
} */

// [05]
/* func TestArea(t *testing.T) {
	// areaTests := []struct {
	//   		shape Shape
	//   		want  float64
	//   	}{
	//   		{Rectangle{12, 6}, 72.0},
	//   		{Circle{10}, 314.1592653589793},
	//   		{Triangle{12, 6}, 36.0},
	//   	}

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("expected %g want %g", got, tt.want)
		}
	}
} */

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

  for _, tt := range areaTests {
    t.Run(tt.name, func(t *testing.T) {
      got := tt.shape.Area()
      
      if got != tt.hasArea {
        t.Errorf("%#v expected %g want %g", tt.shape, got, tt.hasArea)
      }
    })
  }
}
