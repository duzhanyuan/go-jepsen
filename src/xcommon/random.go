/*
 * go-jepsen
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xcommon

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// RandInt64 return the random int64 values between min and max.
func RandInt64(min int64, max int64) int64 {
	return min + int64(rand.Int63n(int64(max-min)))
}
