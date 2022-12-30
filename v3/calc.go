package v2

import (
	"errors"

	"github.com/sirupsen/logrus"
)

func Sum(num1, num2, num3 int) (int, error) {
	if num1 == 0 || num2 == 0 {
		return 0, errors.New("do not sum zero numbers")
	}
	logrus.Infof("The numbers are - num1=%d, num2=%d, num3=%d", num1, num2, num3)
	return (num1 + num2 + num3), nil
}
