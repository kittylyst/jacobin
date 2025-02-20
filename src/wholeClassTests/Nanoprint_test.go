/*
 * Jacobin VM - A Java virtual machine
 * Copyright (c) 2022 by Andrew Binstock. All rights reserved.
 * Licensed under Mozilla Public License 2.0 (MPL 2.0)
 */

package wholeClassTests

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

// Test for Nanoprint class, which calls java.lang.System.nanoTime() twice and prints the result.
// This is a test of returning a value (here, a long with the nano count) from a go-style function,
// which is standing in for the Java call. In addition, this class is one of the first to use the
// lstore and lload instructions.
//
// Source code:
//
// import static java.lang.System.nanoTime;
//
// public class NanoPrint {
//
//     public static void main( String[] args) {
//          long nano1 = nanoTime();
//          long nano2 = nanoTime();
//          System.out.println( nano1 );
//          System.out.println( nano2 );
//     }
// }
//
// These tests check the output with various options for verbosity and features set on the command line.

func initVarsNanoPrint() {
	_JACOBIN = "d:\\GoogleDrive\\Dev\\jacobin\\src\\jacobin.exe"
	_JVM_ARGS = ""
	_TESTCLASS = "d:\\GoogleDrive\\Dev\\jacobin\\testdata\\NanoPrint.class" // the class to test
	_APP_ARGS = ""
}

func TestRunNanoprint(t *testing.T) {
	initVarsNanoPrint()
	var cmd *exec.Cmd

	if testing.Short() { // don't run if running quick tests only. (Used primarily so GitHub doesn't run and bork)
		t.Skip()
	}

	// test that executable exists
	if _, err := os.Stat(_JACOBIN); err != nil {
		t.Errorf("Missing Jacobin executable, which was specified as %s", _JACOBIN)
	}

	// run the various combinations of args. This is necessary b/c the empty string is viewed as
	// an actual specified option on the command line.
	if len(_JVM_ARGS) > 0 {
		if len(_APP_ARGS) > 0 {
			cmd = exec.Command(_JACOBIN, _JVM_ARGS, _TESTCLASS, _APP_ARGS)
		} else {
			cmd = exec.Command(_JACOBIN, _JVM_ARGS, _TESTCLASS)
		}
	} else {
		if len(_APP_ARGS) > 0 {
			cmd = exec.Command(_JACOBIN, _TESTCLASS, _APP_ARGS)
		} else {
			cmd = exec.Command(_JACOBIN, _TESTCLASS)
		}
	}

	// get the stdout and stderr contents from the file execution
	stderr, err := cmd.StderrPipe()
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// run the command
	if err = cmd.Start(); err != nil {
		t.Errorf("Got error running Jacobin: %s", err.Error())
	}

	// Here begin the actual tests on the output to stderr and stdout
	slurp, _ := io.ReadAll(stderr)
	slurpErr := string(slurp)
	if len(slurp) != 0 {
		t.Errorf("Got unexpected output to stderr: %s", slurpErr)
	}

	slurp, _ = io.ReadAll(stdout)
	slurpOut := string(slurp)
	if !strings.HasPrefix(slurpOut, "Jacobin VM") {
		t.Errorf("Stdout did not begin with Jacobin copyright, instead: %s", slurpOut)
	}

	outStrings := strings.Split(strings.ReplaceAll(slurpOut, "\r\n", "\n"), "\n")
	time1, err1 := strconv.Atoi(outStrings[1])
	time2, err2 := strconv.Atoi(outStrings[2])

	if err1 != nil || err2 != nil {
		t.Errorf("Error converting nanoTimes to integers in lines[1] and [2]: %s", outStrings)
	}

	if time2 < time1 {
		t.Errorf("expected time2 to be >= to time1, but got: time1 = %d, time2 = %d", time1, time2)
	}
}

func TestRunNanoPrintVerboseClass(t *testing.T) {
	initVarsNanoPrint()
	var cmd *exec.Cmd

	if testing.Short() { // don't run if running quick tests only. (Used primarily so GitHub doesn't run and bork)
		t.Skip()
	}

	// test that executable exists
	if _, err := os.Stat(_JACOBIN); err != nil {
		t.Errorf("Missing Jacobin executable, which was specified as %s", _JACOBIN)
	}

	_JVM_ARGS = "-verbose:class"
	// run the various combinations of args. This is necessary b/c the empty string is viewed as
	// an actual specified option on the command line.
	if len(_JVM_ARGS) > 0 {
		if len(_APP_ARGS) > 0 {
			cmd = exec.Command(_JACOBIN, _JVM_ARGS, _TESTCLASS, _APP_ARGS)
		} else {
			cmd = exec.Command(_JACOBIN, _JVM_ARGS, _TESTCLASS)
		}
	} else {
		if len(_APP_ARGS) > 0 {
			cmd = exec.Command(_JACOBIN, _TESTCLASS, _APP_ARGS)
		} else {
			cmd = exec.Command(_JACOBIN, _TESTCLASS)
		}
	}

	// get the stdout and stderr contents from the file execution
	stderr, err := cmd.StderrPipe()
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// run the command
	if err = cmd.Start(); err != nil {
		t.Errorf("Got error running Jacobin: %s", err.Error())
	}

	// Here begin the actual tests on the output to stderr and stdout
	slurp, _ := io.ReadAll(stderr)
	slurpErr := string(slurp)
	if !strings.Contains(slurpErr, "Class: NanoPrint, loader: bootstrap") {
		t.Errorf("Got unexpected output to stderr: %s", slurpErr)
	}

	slurp, _ = io.ReadAll(stdout)
	slurpOut := string(slurp)
	if !strings.HasPrefix(string(slurp), "Jacobin VM") {
		t.Errorf("Stdout did not begin with Jacobin copyright, instead: %s", slurpOut)
	}
}

func TestRunNanoPrintVerboseFinest(t *testing.T) {
	initVarsNanoPrint()
	var cmd *exec.Cmd

	if testing.Short() { // don't run if running quick tests only. (Used primarily so GitHub doesn't run and bork)
		t.Skip()
	}

	// test that executable exists
	if _, err := os.Stat(_JACOBIN); err != nil {
		t.Errorf("Missing Jacobin executable, which was specified as %s", _JACOBIN)
	}

	_JVM_ARGS = "-verbose:finest"
	// run the various combinations of args. This is necessary b/c the empty string is viewed as
	// an actual specified option on the command line.
	if len(_JVM_ARGS) > 0 {
		if len(_APP_ARGS) > 0 {
			cmd = exec.Command(_JACOBIN, _JVM_ARGS, _TESTCLASS, _APP_ARGS)
		} else {
			cmd = exec.Command(_JACOBIN, _JVM_ARGS, _TESTCLASS)
		}
	} else {
		if len(_APP_ARGS) > 0 {
			cmd = exec.Command(_JACOBIN, _TESTCLASS, _APP_ARGS)
		} else {
			cmd = exec.Command(_JACOBIN, _TESTCLASS)
		}
	}

	// get the stdout and stderr contents from the file execution
	stderr, err := cmd.StderrPipe()
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// run the command
	if err = cmd.Start(); err != nil {
		t.Errorf("Got error running Jacobin: %s", err.Error())
	}

	// Here begin the actual tests on the output to stderr and stdout
	slurp, _ := io.ReadAll(stderr)
	slurpErr := string(slurp)
	if !strings.Contains(slurpErr, "Class NanoPrint has been format-checked.") {
		t.Errorf("Got unexpected output to stderr: %s", slurpErr)
	}

	slurp, _ = io.ReadAll(stdout)
	slurpOut := string(slurp)
	if !strings.HasPrefix(slurpOut, "Jacobin VM") {
		t.Errorf("Stdout did not begin with Jacobin copyright, instead: %s", slurpOut)
	}
}

func TestRunNanoPrintTraceInst(t *testing.T) {
	var cmd *exec.Cmd

	if testing.Short() { // don't run if running quick tests only. (Used primarily so GitHub doesn't run and bork)
		t.Skip()
	}

	// test that executable exists
	if _, err := os.Stat(_JACOBIN); err != nil {
		t.Errorf("Missing Jacobin executable, which was specified as %s", _JACOBIN)
	}

	_JVM_ARGS = "-trace:inst"
	// run the various combinations of args. This is necessary b/c the empty string is viewed as
	// an actual specified option on the command line.
	if len(_JVM_ARGS) > 0 {
		if len(_APP_ARGS) > 0 {
			cmd = exec.Command(_JACOBIN, _JVM_ARGS, _TESTCLASS, _APP_ARGS)
		} else {
			cmd = exec.Command(_JACOBIN, _JVM_ARGS, _TESTCLASS)
		}
	} else {
		if len(_APP_ARGS) > 0 {
			cmd = exec.Command(_JACOBIN, _TESTCLASS, _APP_ARGS)
		} else {
			cmd = exec.Command(_JACOBIN, _TESTCLASS)
		}
	}

	// get the stdout and stderr contents from the file execution
	stderr, err := cmd.StderrPipe()
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// run the command
	if err = cmd.Start(); err != nil {
		t.Errorf("Got error running Jacobin: %s", err.Error())
	}

	// Here begin the actual tests on the output to stderr and stdout
	slurp, _ := io.ReadAll(stderr)
	slurpErr := string(slurp)
	if !strings.Contains(slurpErr, "class: NanoPrint, meth: main, pc: 22, inst: RETURN, tos: -1") {
		t.Errorf("Got unexpected output to stderr: %s", slurpErr)
	}

	slurp, _ = io.ReadAll(stdout)
	slurpOut := string(slurp)
	if !strings.HasPrefix(slurpOut, "Jacobin VM") {
		t.Errorf("Stdout did not begin with Jacobin copyright, instead: %s", slurpOut)
	}
}
