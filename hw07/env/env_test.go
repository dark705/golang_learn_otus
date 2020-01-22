package env

import (
	"io/ioutil"
	"os"
	"testing"
)

const (
	testEnvDir = "./test_env_dir/"
	envFile1   = "ENV1"
	envFile2   = "ENV2"
	val1       = "val1"
	val2       = "123"
)

func TestReadDirEnv(t *testing.T) {
	var err error
	defer func() {
		_ = os.RemoveAll(testEnvDir)
	}()

	err = os.Mkdir(testEnvDir, 0755)
	if err != nil {
		t.Error("Can't create dir for test")
	}

	err = ioutil.WriteFile(testEnvDir+envFile1, []byte(val1), 0644)
	if err != nil {
		t.Error("Can't write file for test")
	}

	err = ioutil.WriteFile(testEnvDir+envFile2, []byte(val2), 0644)
	if err != nil {
		t.Error("Can't write file for test")
	}

	environment, err := ReadDir(testEnvDir)
	if err != nil {
		t.Error("Can't read dir", err)
	}

	if environment[envFile1] != val1 || environment[envFile2] != val2 {
		t.Error("Wrong environment for dir")
	}
}

func TestRunCmdWithArgs(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	exitCode := RunCmd([]string{"echo", "123"}, map[string]string{})
	if exitCode != 0 {
		t.Error("RunCmd not return success exit code")
	}

	_ = w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if string(out) != "123\n" {
		t.Error(`Run test command: "echo 123" not output 123`)
	}
}
