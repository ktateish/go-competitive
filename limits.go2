package main

// vim:set ft=go:

import (
	"fmt"
	"math"
	"unsafe"
)

// snip ------------------------------------------------------------------------

var (
	// For 32bit, 64bit integers, useful values are used for competitive
	// programming that typically have 10^9 or 10^18 constraints for upper limits.
	// The `useful` means that the values are greater than the upper limits
	// while they don't overflow when they are multiplied by 2.
	limits_inf_int8  int8  = math.MaxInt8
	limits_inf_int16 int16 = math.MaxInt16
	limits_inf_int32 int32 = int32(1e9 + 1000) // useful for competitive programming
	limits_inf_int64 int64 = 1<<60 - 1         // useful for competitive programming
	limits_inf_int   int

	limits_inf_uint8  uint8  = math.MaxUint8
	limits_inf_uint16 uint16 = math.MaxUint16
	limits_inf_uint32 uint32 = uint32(limits_inf_int32)
	limits_inf_uint64 uint64 = uint64(limits_inf_int64)
	limits_inf_uint   uint

	limits_inf_float32 float32 = math.MaxFloat32
	limits_inf_float64 float64 = math.MaxFloat64

	limits_ninf_int8  int8  = math.MinInt8
	limits_ninf_int16 int16 = math.MinInt16
	limits_ninf_int32 int32 = -int32(1e9 + 1000) // useful for competitive programming
	limits_ninf_int64 int64 = -(1<<60 - 1)       // useful for competitive programming
	limits_ninf_int   int

	limits_ninf_uint8  uint8  = 0
	limits_ninf_uint16 uint16 = 0
	limits_ninf_uint32 uint32 = 0
	limits_ninf_uint64 uint64 = 0
	limits_ninf_uint   uint   = 0

	limits_ninf_float32 float32 = -math.MaxFloat32
	limits_ninf_float64 float64 = -math.MaxFloat64
)

func init() {
	var y int
	sz := unsafe.Sizeof(y)
	if sz == 4 {
		limits_inf_int = int(limits_inf_int32)
		limits_inf_uint = uint(limits_inf_uint32)
		limits_ninf_int = int(limits_ninf_int32)
	} else if sz == 8 {
		limits_inf_int = int(limits_inf_int64)
		limits_inf_uint = uint(limits_inf_uint64)
		limits_ninf_int = int(limits_ninf_int64)
	} else {
		panic(fmt.Sprintf("unknown int size: %d", sz))
	}
}

func GetInf[T OrderedNumeric]() T {
	var x T
	switch (interface{})(x).(type) {
	case int:
		return T(limits_inf_int)
	case int8:
		return T(limits_inf_int8)
	case int16:
		return T(limits_inf_int16)
	case int32:
		return T(limits_inf_int32)
	case int64:
		return T(limits_inf_int64)
	case uint:
		return T(limits_inf_uint)
	case uint8:
		return T(limits_inf_uint8)
	case uint16:
		return T(limits_inf_uint16)
	case uint32:
		return T(limits_inf_uint32)
	case uint64:
		return T(limits_inf_uint64)
	case float32:
		return T(limits_inf_float32)
	case float64:
		return T(limits_inf_float64)
	}
	panic("unknown type for GetInf")
}

func GetNInf[T OrderedNumeric]() T {
	var x T
	switch (interface{})(x).(type) {
	case int:
		return T(limits_ninf_int)
	case int8:
		return T(limits_ninf_int8)
	case int16:
		return T(limits_ninf_int16)
	case int32:
		return T(limits_ninf_int32)
	case int64:
		return T(limits_ninf_int64)
	case uint:
		return T(limits_ninf_uint)
	case uint8:
		return T(limits_ninf_uint8)
	case uint16:
		return T(limits_ninf_uint16)
	case uint32:
		return T(limits_ninf_uint32)
	case uint64:
		return T(limits_ninf_uint64)
	case float32:
		return T(limits_ninf_float32)
	case float64:
		return T(limits_ninf_float64)
	}
	panic("unknown type for GetInf")
}

func SetInf[T OrderedNumeric](x T) {
	switch (interface{})(x).(type) {
	case int:
		(limits_inf_int) = int(x)
	case int8:
		(limits_inf_int8) = int8(x)
	case int16:
		(limits_inf_int16) = int16(x)
	case int32:
		(limits_inf_int32) = int32(x)
	case int64:
		(limits_inf_int64) = int64(x)
	case uint:
		(limits_inf_uint) = uint(x)
	case uint8:
		(limits_inf_uint8) = uint8(x)
	case uint16:
		(limits_inf_uint16) = uint16(x)
	case uint32:
		(limits_inf_uint32) = uint32(x)
	case uint64:
		(limits_inf_uint64) = uint64(x)
	case float32:
		(limits_inf_float32) = float32(x)
	case float64:
		(limits_inf_float64) = float64(x)
	}
	panic("unknown type for GetInf")
}

func SetNInf[T OrderedNumeric](x T) {
	switch (interface{})(x).(type) {
	case int:
		(limits_ninf_int) = int(x)
	case int8:
		(limits_ninf_int8) = int8(x)
	case int16:
		(limits_ninf_int16) = int16(x)
	case int32:
		(limits_ninf_int32) = int32(x)
	case int64:
		(limits_ninf_int64) = int64(x)
	case uint:
		(limits_ninf_uint) = uint(x)
	case uint8:
		(limits_ninf_uint8) = uint8(x)
	case uint16:
		(limits_ninf_uint16) = uint16(x)
	case uint32:
		(limits_ninf_uint32) = uint32(x)
	case uint64:
		(limits_ninf_uint64) = uint64(x)
	case float32:
		(limits_ninf_float32) = float32(x)
	case float64:
		(limits_ninf_float64) = float64(x)
	}
	panic("unknown type for GetInf")
}

func numeric_limits_max[T OrderedNumeric]() T {
	var x T
	casti := func(x uint64) T { return T(x) }
	castf := func(x float64) T { return T(x) }
	switch (interface{})(x).(type) {
	case int:
		var y int
		sz := unsafe.Sizeof(y)
		if sz == 4 {
			return casti(math.MaxInt32)
		} else if sz == 8 {
			return casti(math.MaxInt64)
		}
	case int8:
		return casti(math.MaxInt8)
	case int16:
		return casti(math.MaxInt16)
	case int32:
		return casti(math.MaxInt32)
	case int64:
		return casti(math.MaxInt64)
	case uint:
		var y uint
		sz := unsafe.Sizeof(y)
		if sz == 4 {
			return casti(math.MaxUint32)
		} else if sz == 8 {
			return casti(math.MaxUint64)
		}
	case uint8:
		return casti(math.MaxUint8)
	case uint16:
		return casti(math.MaxUint16)
	case uint32:
		return casti(math.MaxUint32)
	case uint64:
		return casti(math.MaxUint64)
	case float32:
		return castf(math.MaxFloat32)
	case float64:
		return castf(math.MaxFloat64)
	}
	panic("unknown type for numeric_limits_max")
}

func GetMax[T OrderedNumeric]() T {
	return numeric_limits_max[T]()
}

func numeric_limits_min[T OrderedNumeric]() T {
	var x T
	casti := func(x int64) T { return T(x) }
	castf := func(x float64) T { return T(x) }
	switch (interface{})(x).(type) {
	case int:
		var y int
		sz := unsafe.Sizeof(y)
		if sz == 4 {
			return casti(math.MinInt32)
		} else if sz == 8 {
			return casti(math.MinInt64)
		}
	case int8:
		return casti(math.MinInt8)
	case int16:
		return casti(math.MinInt16)
	case int32:
		return casti(math.MinInt32)
	case int64:
		return casti(math.MinInt64)
	case uint:
		var y uint
		sz := unsafe.Sizeof(y)
		if sz == 4 {
			return casti(0)
		} else if sz == 8 {
			return casti(0)
		}
	case uint8:
		return casti(0)
	case uint16:
		return casti(0)
	case uint32:
		return casti(0)
	case uint64:
		return casti(0)
	case float32:
		return castf(-math.MaxFloat32)
	case float64:
		return castf(-math.MaxFloat64)
	}
	panic("unknown type for numeric_limits_max")
}

func GetMin[T OrderedNumeric]() T {
	return numeric_limits_min[T]()
}
