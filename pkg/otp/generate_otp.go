package otp

import (
	"fmt"
	"math/rand"
)

func GenerateOTP() string {

	return fmt.Sprintf("%06d", rand.Intn(1000000))

}
