package go_testing

import "errors"

func Divide(num, den int)(int, error){
	if den == 0{
		err := errors.New("o denominador não pode ser zero")
		return 0, err
	}
	
	return num/den, nil
}