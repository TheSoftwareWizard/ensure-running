//# Copyright (C) 2026 @TheSoftwareWizard
//
//# This program is free software: you can redistribute it and/or modify
//# it under the terms of the GNU General Public License as published by
//# the Free Software Foundation, either version 3 of the License, or
//# (at your option) any later version.
//
//# This program is distributed in the hope that it will be useful,
//# but WITHOUT ANY WARRANTY; without even the implied warranty of
//# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//# GNU General Public License for more details.
//
//# You should have received a copy of the GNU General Public License
//# along with this program.  If not, see <https://gnu.org>.
package main
import (
	"os/exec"
	"fmt"
	"flag"
	"time"
)

func main(){
	process := flag.String("p", "", "Process binary name.")
	processNames := flag.String("n", "", "Processes that if running do not execute process.")
	processArgs := flag.String("a", "", "Extra arguments for process.")
	processSleepBefore := flag.Int("sb", 0, "Delay in milliseconds before running process.")
	processSleepAfter := flag.Int("sa", 0, "Delay in milliseconds after running process.")

	flag.Parse()

	if *processNames == "" {
		*processNames = *process
	}
	
	if *processSleepBefore > 0 { time.Sleep(time.Duration(*processSleepBefore) * time.Millisecond)}
	ensureRunning := `/usr/bin/pgrep -x "%v" >/dev/null 2>&1 || "/usr/bin/%v" %v >/dev/null 2>&1 &`

	script := fmt.Sprintf(ensureRunning, *processNames, *process, *processArgs)
	cmd := exec.Command("/bin/sh", "-c", script)
	_ = cmd.Run()
	if *processSleepAfter > 0 { time.Sleep(time.Duration(*processSleepAfter) * time.Millisecond)}
}
