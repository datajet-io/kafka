package consumergroup

import (
	"errors"
)

var ErrEmptyConsumerList = errors.New("get empty consumer list from zookeeper")
var ErrGetConsumerList = errors.New("failed to get list of registered consumer instancess")
var ErrRegisterConsumer = errors.New("failed to register consumer instance")
