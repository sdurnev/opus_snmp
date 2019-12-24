// Copyright 2012-2014 The GoSNMP Authors. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in the
// LICENSE file.

package main

import (
	"fmt"
	g "github.com/soniah/gosnmp"
	"log"
	"math/big"
)

/*
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!! VERSION !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
*/
const version = "0.00.1"

/*
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!! VERSION !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
*/

func main() {

	paramBCMmib := []string{"1.3.6.1.4.1.26896.1.3.1",
		"1.3.6.1.4.1.26896.1.3.2",
		"1.3.6.1.4.1.26896.1.3.3",
		"1.3.6.1.4.1.26896.1.3.4",
		"1.3.6.1.4.1.26896.1.3.5",
		"1.3.6.1.4.1.26896.1.3.6",
		"1.3.6.1.4.1.26896.1.3.7"}

	paramBCM := [...]string{"bcmSystemVoltage",
		"bcmLoadCurrent",
		"bcmBatteryCurrent",
		"bcmTotalRectifierCurrent",
		"bcmTotalInverterCurrent",
		"bcmMaxBatteryTemperature",
		"bcmMaxSystemTemperature",
	}

	// Default is a pointer to a GoSNMP struct that contains sensible defaults
	// eg port 161, community public, etc
	g.Default.Target = "10.10.12.21"
	g.Default.Community = "Public"
	err := g.Default.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer g.Default.Conn.Close()

	oids := paramBCMmib
	//oids := []string{"1.3.6.1.2.1.1.4", "1.3.6.1.2.1.1.7"}
	//oids := []string{".1.3.6.1.4.1.26896.1.1.106.0", ".1.3.6.1.4.1.26896.1.1.112.0"}
	//oids := []string{".1.3.6.1.4.1.26896.1.1.111"}

	result, err2 := g.Default.GetNext(oids) // Get() accepts up to g.MAX_OIDS
	if err2 != nil {
		log.Fatalf("Get() err: %v", err2)
	}

	for i, variable := range result.Variables {
		//fmt.Printf("%d: oid: %s ", i, variable.Name)
		if i == 0 {
			fmt.Printf("{")
		}
		fmt.Printf("\"%s\":", paramBCM[i])

		a, _ := new(big.Float).SetInt(g.ToBigInt(variable.Value)).Float64()
		c := a / float64(1000)
		if i == len(result.Variables)-1 {
			fmt.Printf("%.2f,", c)
		} else {
			fmt.Printf("%.2f,", c) //!!!!!!!!!!! походу это лишнее
		}
	}
	fmt.Printf("\"version\": \"%s\"}", version)
}
