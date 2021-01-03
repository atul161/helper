package helper

import (
	"errors"
	"fmt"
	"math"
	"os"
	"runtime"
	"strings"
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

// Trim float function trim the float number with certain digits
// in case the value of digit is zero or negative then we will return
// the actual number
// i.e := if number is 123.3454 and digit is 2 then result will be
// 123.34
func TrimFloat(no float64, digit int64) float64 {
	if digit <= 0 {
		return no
	}
	d := math.Pow(10, float64(digit))
	return float64(int64(no*d)) / d
}


type Directory struct {
	Override bool
	Name string
	Path string
	Permission int64
}

// Make Directory is a receiver function which
// will create a directory to the specified path.
// If we set override true then it will delete the specified folder and create new
// directory.
// Default Permission of the created directory will be 0755.
//
func (dir *Directory) MakeDirectory() error {
    filepath := dir.Path+"/"+dir.Name
	//Checking the operating system.
	if runtime.GOOS == "windows" {
		filepath = dir.Path+"\\"+dir.Name
	}

	if dir.Override {
		if err := os.RemoveAll(filepath); err != nil {
			return err
		}
	}

	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(filepath, os.FileMode(dir.Permission))
		if errDir != nil {
          return errDir
		}
	}
	return nil
}

func defaultDirectory() *Directory {
	return &Directory{
		Override:   false,
		Permission: 0755,
	}
}