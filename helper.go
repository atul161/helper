package helper

import (
	"math"
)

type Container struct {
	list []float64
}

// Power function is a receiver method which accept the list
// of numbers and power every number accordingly
// ie : let a number be a , b , c then its power will be
// ((a)^b)^c
//In case list is empty or nil we will return -1
func (c *Container) Power() float64 {
	if len(c.list) == 0 || c.list == nil {
		return -1
	}
	var result float64
	for i := 0; i < len(c.list); i++ {
		if i == 0 {
			result = c.list[i]
			continue
		}
		result = math.Pow(result, c.list[i])
	}
	return result
}

// trim float function trim the float no with certain digits
// in case the value of digit is zero or negative then we will return
//the actual number
// i.e := if no is 123.3454 and digit is 2 then result will be
// 123.34
func TrimFloat(no float64, digit int64) float64 {
	if digit <= 0 {
		return no
	}
	d := math.Pow(10, float64(digit))
	return float64(int64(no*d)) / d
}
