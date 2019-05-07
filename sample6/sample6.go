

package service

import (
	"fmt"
)

type Trilean int
const Yes Trilean = 1
const No Trilean = 0
const IHopeNotButProbablyYes Trilean = -1

type ServiceError struct {
	Message string
	DidWePayMoneyForThis Trilean
}

func (serviceErr *ServiceError) Error() string {
	if serviceErr == nil {
		return "<nil serviceError>"
	}
	return serviceErr.Message
}

func GetResultFromService(id int) (result string, err error) {
	result, err = getResult(id)
	if err != nil {
		return result, fmt.Errorf("getResult(%d) failed with error: %v", id, err)
	}
	return result, nil
}

func getResult(id int) (string, *ServiceError) {
	return fmt.Sprintf("<Result for object with ID %d>", id), nil
}
