package consumergroup

import (
	"errors"
)

var ErrEmptyConsumerList = errors.New("get empty consumer list from zookeeper")
