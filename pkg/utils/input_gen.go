package utils

import (
	"math/rand"
	"reflect"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomValue generates a random value based on the given type.
func RandomValue(t reflect.Type) reflect.Value {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(randomInt()).Convert(t)
	case reflect.Float32, reflect.Float64:
		return reflect.ValueOf(randomFloat()).Convert(t)
	case reflect.String:
		return reflect.ValueOf(randomString())
	case reflect.Bool:
		return reflect.ValueOf(randomBool())
	default:
		return reflect.Zero(t)
	}
}

func randomInt() int {
	return rand.Intn(2001) - 1000 // Range from -1000 to 1000
}

func randomFloat() float64 {
	return rand.Float64()*2000 - 1000 // Range from -1000.0 to 1000.0
}

func randomString() string {
	length := rand.Intn(20) + 1
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func randomBool() bool {
	return rand.Intn(2) == 0
}
