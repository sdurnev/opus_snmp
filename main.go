// Copyright 2012-2014 The GoSNMP Authors. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in the
// LICENSE file.

package main

import (
	"flag"
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

	paramBCMmib := []string{
		"1.3.6.1.4.1.26896.1.3.1",
		"1.3.6.1.4.1.26896.1.3.2",
		"1.3.6.1.4.1.26896.1.3.3",
		"1.3.6.1.4.1.26896.1.3.4",
		"1.3.6.1.4.1.26896.1.3.5",
		"1.3.6.1.4.1.26896.1.3.6",
		"1.3.6.1.4.1.26896.1.3.7"}

	paramBCOmib := []string{
		"1.3.6.1.4.1.26896.1.2.1",
		"1.3.6.1.4.1.26896.1.2.2",
		"1.3.6.1.4.1.26896.1.2.3",
		"1.3.6.1.4.1.26896.1.2.4",
	}

	paramVIDImib := []string{
		"1.3.6.1.4.1.26896.1.1.100",
		"1.3.6.1.4.1.26896.1.1.101",
		"1.3.6.1.4.1.26896.1.1.102",
		"1.3.6.1.4.1.26896.1.1.103",
		"1.3.6.1.4.1.26896.1.1.104",
		"1.3.6.1.4.1.26896.1.1.105",
		"1.3.6.1.4.1.26896.1.1.106",
		"1.3.6.1.4.1.26896.1.1.110",
		"1.3.6.1.4.1.26896.1.1.111",
		"1.3.6.1.4.1.26896.1.1.112",
	}

	paramBCM := [...]string{
		"bcmSystemVoltage",
		"bcmLoadCurrent",
		"bcmBatteryCurrent",
		"bcmTotalRectifierCurrent",
		"bcmTotalInverterCurrent",
		"bcmMaxBatteryTemperature",
		"bcmMaxSystemTemperature",
	}

	paramBCO := [...]string{
		"bcoChargeState",
		"bcoTemperatureCompensation",
		"bcoTestMode",
		"bcoBoostChargeMode",
	}

	paramVIDI := []string{
		"vidiAlarmId",
		"vidiAlarmMessage",
		"vidiAlarmCode",
		"vidiAlarmState",
		"vidiAlarmAckState",
		"vidiAlarmStartTime",
		"vidiAlarmDuration",
		"vidiNumActiveAlarms",
		"vidiNumNonAckAlarms",
		"vidiAlarmRelays",
	}

	addressIP := flag.String("ip", "localhost", "a string")
	tcpPort := flag.Uint("port", 161, "a Uint")
	community := flag.String("com", "Public", "an string")
	typeOfPar := flag.Int("type", 1, "an int")
	flag.Parse()

	g.Default.Target = fmt.Sprint(*addressIP)
	g.Default.Community = fmt.Sprint(*community)
	g.Default.Port = uint16(int16(*tcpPort))
	var param int = *typeOfPar

	if param == 1 {
		err := g.Default.Connect()
		if err != nil {
			log.Fatalf("Connect() err: %v", err)
		}
		defer g.Default.Conn.Close()

		oids := paramBCMmib
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
	}
	if param == 2 {
		err := g.Default.Connect()
		if err != nil {
			log.Fatalf("Connect() err: %v", err)
		}
		defer g.Default.Conn.Close()
		oids := paramBCOmib
		result, err2 := g.Default.GetNext(oids) // Get() accepts up to g.MAX_OIDS
		if err2 != nil {
			log.Fatalf("Get() err: %v", err2)
		}

		for i, variable := range result.Variables {
			if i == 0 {
				fmt.Printf("{")
			}
			fmt.Printf("\"%s\":", paramBCO[i])

			// the Value of each variable returned by Get() implements
			// interface{}. You could do a type switch...
			switch variable.Type {
			case g.OctetString:
				fmt.Printf("string: %s\n", string(variable.Value.([]byte)))
			default:
				// ... or often you're just interested in numeric values.
				// ToBigInt() will return the Value as a BigInt, for plugging
				// into your calculations.
				fmt.Printf("%d,", g.ToBigInt(variable.Value))
			}
		}
	}
	if param == 3 {
		err := g.Default.Connect()
		if err != nil {
			log.Fatalf("Connect() err: %v", err)
		}
		defer g.Default.Conn.Close()
		oids := paramVIDImib
		result, err2 := g.Default.GetNext(oids) // Get() accepts up to g.MAX_OIDS
		if err2 != nil {
			log.Fatalf("Get() err: %v", err2)
		}
		/*fmt.Print("\n")
		fmt.Print(result)
		fmt.Print("\n")*/
		for i, variable := range result.Variables {
			if i == 0 {
				fmt.Printf("{")
			}
			fmt.Printf("\"%s\":", paramVIDI[i])

			// the Value of each variable returned by Get() implements
			// interface{}. You could do a type switch...
			switch variable.Type {
			case g.OctetString:
				fmt.Printf("string: %s\n", string(variable.Value.([]byte)))
			default:
				// ... or often you're just interested in numeric values.
				// ToBigInt() will return the Value as a BigInt, for plugging
				// into your calculations.
				fmt.Printf("%d,", g.ToBigInt(variable.Value))
			}
		}
	}
	fmt.Printf("\"version\":\"%s\"}", version)
}
