package a

func f[_ any]() {} 	  // want "f has a type parameter"
type t[_ any] struct{}    // want "t has a type parameter"

// using generic types and functions is OK
type _ = t[int] // OK
var _ = f[int]  // OK
