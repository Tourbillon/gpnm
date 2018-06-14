// Copyright (c) 2018-present Anbillon Team (anbillonteam@gmail.com).

package util

import (
	"strconv"
)

func ParseInt(origin string) int64 {
	i, err := strconv.ParseInt(origin, 0, 0)
	if err != nil {
		return -1
	}
	return i
}
