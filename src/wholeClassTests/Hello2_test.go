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
	"strings"
	"testing"
)

/*
 * Tests for Hello2.class, which is one of the first classes Jacobin executed. Source code:
 *
 *		public static void main( String[] args) {
 *			int x;
 *			for( int i = 0; i < 10; i++) {
 *				x = addTwo(i, i-1);
 *				System.out.println( x );
 *          }
 *      }
 *
 *	    static int addTwo(int j, int k) {
 *		    return j + k;
 *	    }
 *
 * These tests check the output with various options for verbosity and features set on the command line.
 */

func initVarsHello2() {
	_JACOBIN = "d:\\GoogleDrive\\Dev\\jacobin\\src\\jacobin.exe"
	_JVM_ARGS = ""
	_TESTCLASS = "d:\\GoogleDrive\\Dev\\jacobin\\testdata\\Hello2.class" // the class to test
	_APP_ARGS = ""
}

func TestRunHello2(t *testing.T) {
	initVarsHello2()
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
	if len(slurp) != 0 {
		t.Errorf("Got unexpected output to stderr: %s", string(slurp))
	}

	slurp, _ = io.ReadAll(stdout)
	if !strings.HasPrefix(string(slurp), "Jacobin VM") {
		t.Errorf("Stdout did not begin with Jacobin copyright, instead: %s", string(slurp))
	}

	if !strings.Contains(string(slurp), "-1") && !strings.Contains(string(slurp), "17") {
		t.Errorf("Did not get expected output to stdout. Got: %s", string(slurp))
	}
}

func TestRunHello2VerboseClass(t *testing.T) {
	initVarsHello2()
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
	if !strings.Contains(string(slurp), "Class: Hello2, loader: bootstrap") {
		t.Errorf("Got unexpected output to stderr: %s", string(slurp))
	}

	slurp, _ = io.ReadAll(stdout)
	if !strings.HasPrefix(string(slurp), "Jacobin VM") {
		t.Errorf("Stdout did not begin with Jacobin copyright, instead: %s", string(slurp))
	}

	if !strings.Contains(string(slurp), "-1") && !strings.Contains(string(slurp), "17") {
		t.Errorf("Did not get expected output to stdout. Got: %s", string(slurp))
	}
}

func TestRunHello2VerboseFinest(t *testing.T) {
	initVarsHello2()
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
	if !strings.Contains(string(slurp), "Class Hello2 has been format-checked.") {
		t.Errorf("Got unexpected output to stderr: %s", string(slurp))
	}

	slurp, _ = io.ReadAll(stdout)
	if !strings.HasPrefix(string(slurp), "Jacobin VM") {
		t.Errorf("Stdout did not begin with Jacobin copyright, instead: %s", string(slurp))
	}

	if !strings.Contains(string(slurp), "13") {
		t.Errorf("Did not get expected output to stdout. Got: %s", string(slurp))
	}
}

func TestRunHello2TraceInst(t *testing.T) {
	initVarsHello2()
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
	if !strings.Contains(string(slurp), "class: Hello2, meth: main, pc: 29, inst: RETURN, tos: -1") {
		t.Errorf("Got unexpected output to stderr: %s", string(slurp))
	}

	slurp, _ = io.ReadAll(stdout)
	if !strings.HasPrefix(string(slurp), "Jacobin VM") {
		t.Errorf("Stdout did not begin with Jacobin copyright, instead: %s", string(slurp))
	}

	if !strings.Contains(string(slurp), "15") {
		t.Errorf("Did not get expected output to stdout. Got: %s", string(slurp))
	}
}
