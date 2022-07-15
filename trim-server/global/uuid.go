package global

import "github.com/aidarkhanov/nanoid/v2"

func UUID() (string, error) {
	return nanoid.GenerateString(nanoid.DefaultAlphabet, 16)
}
