package main

import "testing"

// [01]
// func TestHello(t *testing.T)  {
//   got := Hello("Jams")
//   want := "Hello, Jams"
//
//   if got != want {
//     t.Errorf("got %q want %q", got, want)
//   }
// }

// [02]
// func TestHello(t *testing.T) {
// 	t.Run("saying hello to people", func(t *testing.T) {
// 		got := Hello("Jams")
// 		want := "Hello, Jams"
//
// 		assertCorrectMessage(t, got, want)
// 	})
//
// 	t.Run("empty string default to 'world'", func(t *testing.T) {
// 		got := Hello("")
// 		want := "Hello, world"
//
// 		assertCorrectMessage(t, got, want)
// 	})
// }
//
// func assertCorrectMessage(t testing.TB, got, want string) {
//   // Show line failed in the test, if omit, the line will be t.Errorf
// 	t.Helper()
//
// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

func TestHello(t *testing.T) {
  t.Run("in English", func(t *testing.T) {
    got := Hello("Jams", "English")
    want := "Hello, Jams"

    assertCorrectMessage(t, got, want)
  })

  t.Run("in French", func(t *testing.T) {
    got := Hello("Jams", "French")
    want := "Bonjour, Jams"

    assertCorrectMessage(t, got, want)
  })


  t.Run("in Spanish", func(t *testing.T) {
    got := Hello("Jams", "Spanish")
    want := "Hola, Jams"

    assertCorrectMessage(t, got, want)
  })
}

func assertCorrectMessage(t testing.TB, got, want string) {
  // Show line failed in the test, if omit, the line will be t.Errorf
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
