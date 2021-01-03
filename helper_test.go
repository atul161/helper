package helper

import (
	"math"
	"testing"
)

func TestContainer_Power(t *testing.T) {
   //Power of integer number
   cont := Container{list:[]float64{
   	  2 , 3 , 4 ,
   }}
  if  cont.Power() != 4096 {
  	t.Error("power of list is incorrect" , cont.list)
  }

  //Power of float number
	cont = Container{list:[]float64{
		2.2 , 3.1 ,
	}}
	if  TrimFloat(cont.Power() , 2) != 11.52 {
		t.Error("power of list is incorrect" , cont.list)
	}

	//Power of negative number
	cont = Container{list:[]float64{
		-2,2,
	}}
	if  cont.Power() != 4 {
		t.Error("power of list is incorrect" , cont.list)
	}

	//Power of negative floating number
	cont = Container{list:[]float64{
		-2.2 , -3.1 ,
	}}
	if  cont.Power() == math.NaN() {
		t.Error("power of list is incorrect" , cont.list)
	}

	//Power 0 , 0
	cont = Container{list:[]float64{
		0, 0 ,
	}}
	if  cont.Power() == math.NaN() {
		t.Error("power of list is incorrect" , cont.list)
	}

	cont = Container{list:[]float64{
	}}
	if  cont.Power() != -1 {
		t.Error("power of list is incorrect" , cont.list)
	}

	cont = Container{list:nil}
	if  cont.Power() != -1 {
		t.Error("power of list is incorrect" , cont.list)
	}

}


func TestTrimFloat(t *testing.T) {
  f := TrimFloat(12.345566 , 2)
  if f != 12.34 {
  	t.Error("not able to trim")
  }

  f = TrimFloat(12.657 , -1)

  if f != 12.657 {
	  t.Error("not able to trim")
  }
}



func TestDirectory_MakeDirectory(t *testing.T) {
	dir := Directory{
		Override:   false,
		Name:       "test",
		Path:       "./",
		Permission: 0755,
	}
	if err := dir.MakeDirectory(); err != nil {
		t.Error(err)
	}

	dir.Override = true
	if err := dir.MakeDirectory(); err != nil {
		t.Error(err)
	}


}
