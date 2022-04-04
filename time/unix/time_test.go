/* For license and copyright information please see LEGAL file in repository */

package unix

import (
	"fmt"
	"testing"
	"time"
)

func Test_ShowNow(t *testing.T) {
	var ti Time = Now()
	fmt.Println(ti)
	fmt.Println(ti.NanoElapsed())
	fmt.Println(ti.MicroElapsed())
	fmt.Println(ti.MilliElapsed())
	fmt.Println(ti.SecElapsed())

	var timeNow = time.Now()
	fmt.Println(timeNow.Unix())
	fmt.Println(timeNow.UnixMilli())
	fmt.Println(timeNow.UnixMicro())
	fmt.Println(timeNow.UnixNano())
}
