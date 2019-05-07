//+build ignore

package service

import (
	"fmt"
)

type ServiceError struct {
	Message              string
	DidWePayMoneyForThis bool
}

func (serviceErr *ServiceError) Error() string {
	if serviceErr == nil {
		return "<nil serviceError>"
	}
	return serviceErr.Message
}

func GetResultFromService(id int) (result string, err error) {
	// Do not assign the *ServiceError to err, or it will be a typed nil interface-value, which is != nil.
	// An even better solution would be if getResult() returned an error instead of a *ServiceError,
	// making it impossible to accidentally run into the problem.
	result, serviceErr := getResult(id)
	if serviceErr != nil {
		return result, fmt.Errorf("getResult(%d) failed with error: %v", id, serviceErr)
	}
	return result, nil
}

func getResult(id int) (string, *ServiceError) {
	return fmt.Sprintf("Glorious result for object with id %d.", id), nil
}
