package emailutils

import (
	"time"

	"github.com/Polo4444/emailutils/otp"
	"github.com/Polo4444/emailutils/otpstore"
)

func InitOtp() {
	otp.NewOTP(otpstore.NewInMemory(), time.Duration(time.Second*10), 6)
}
