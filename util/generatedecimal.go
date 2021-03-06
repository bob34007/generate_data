/**
 * @Author: guobob
 * @Description:
 * @File:  generateint.go
 * @Version: 1.0.0
 * @Date: 2022/3/24 23:05
 */

package util

import (
	"fmt"
	"math/rand"

	"github.com/fufuok/random"
	deci "github.com/shopspring/decimal"
)

func Randdecimal_1(a *Property) (deci.Decimal, error) {
	if (len(a.DefaultVal)) != 0 {
		strnum := len(a.DefaultVal)
		num, err := deci.NewFromString(a.DefaultVal[rand.Intn(strnum)])
		return num, err
	}
	var num deci.Decimal
	bytesInteger := make([]byte, a.Length-a.SuffixLen)
	bytesDec := make([]byte, a.SuffixLen)
	if len(a.StartKey) > 0 {
		err := fmt.Errorf("decimal can not contain startkey ")
		return num, err
	}

	if len(a.EndKey) > 0 {
		err := fmt.Errorf("decimal can not contain endkey ")
		return num, err
	}

	if a.CharFormat == nil {
		for i := 0; i < (a.Length - a.SuffixLen); i++ {
			b := rand.Intn(10) + 48
			bytesInteger[i] = byte(b)
		}
		for i := 0; i < a.SuffixLen; i++ {
			b := rand.Intn(10) + 48
			bytesDec[i] = byte(b)
		}
	} else {
		for i := 0; i < (a.Length - a.SuffixLen); i++ {
			num := len(a.CharFormat)
			b := a.CharFormat[rand.Intn(num)]
			bytesInteger[i] = byte(b)
		}
		for i := 0; i < a.SuffixLen; i++ {
			num := len(a.CharFormat)
			b := a.CharFormat[rand.Intn(num)]
			bytesDec[i] = byte(b)
		}
	}

	randomStr := string(bytesInteger) + "." + string(bytesDec)
	num, err := deci.NewFromString(randomStr)
	return num, err

}

func Randdecimal(a *Property) (deci.Decimal, error) {
	if (len(a.DefaultVal)) != 0 {
		strnum := len(a.DefaultVal)
		num, err := deci.NewFromString(a.DefaultVal[random.FastIntn(strnum)])
		return num, err
	}
	var num deci.Decimal
	bytesInteger := make([]byte, a.Length-a.SuffixLen)
	bytesDec := make([]byte, a.SuffixLen)
	if len(a.StartKey) > 0 {
		err := fmt.Errorf("decimal can not contain startkey ")
		return num, err
	}

	if len(a.EndKey) > 0 {
		err := fmt.Errorf("decimal can not contain endkey ")
		return num, err
	}

	if a.CharFormat == nil {
		for i := 0; i < (a.Length - a.SuffixLen); i++ {
			b := random.FastIntn(10) + 48
			bytesInteger[i] = byte(b)
		}
		for i := 0; i < a.SuffixLen; i++ {
			b := random.FastIntn(10) + 48
			bytesDec[i] = byte(b)
		}
	} else {
		for i := 0; i < (a.Length - a.SuffixLen); i++ {
			num := len(a.CharFormat)
			b := a.CharFormat[random.FastIntn(num)]
			bytesInteger[i] = byte(b)
		}
		for i := 0; i < a.SuffixLen; i++ {
			num := len(a.CharFormat)
			b := a.CharFormat[random.FastIntn(num)]
			bytesDec[i] = byte(b)
		}
	}

	randomStr := string(bytesInteger) + "." + string(bytesDec)
	num, err := deci.NewFromString(randomStr)
	return num, err

}
