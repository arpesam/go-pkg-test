package calc

import "github.com/sirupsen/logrus"

func Sum(num1, num2 int) int {
	logrus.Infof("The numbers are - num1=%d, num2=%d", num1, num2)
	return num1 + num2
}
