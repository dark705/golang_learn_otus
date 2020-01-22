package env

import (
	"io/ioutil"
	"os"
	"testing"
)

const (
	testEnvDir = "./test_env_dir/"
	env1       = "ENV1"
	env2       = "ENV2"
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

	err = ioutil.WriteFile(testEnvDir+env1, []byte(val1), 0644)
	if err != nil {
		t.Error("Can't write file for test")
	}

	err = ioutil.WriteFile(testEnvDir+env2, []byte(val2), 0644)
	if err != nil {
		t.Error("Can't write file for test")
	}

	environment, err := ReadDir(testEnvDir)
	if err != nil {
		t.Error("Can't read dir", err)
	}

	if environment[env1] != val1 || environment[env2] != val2 {
		t.Error("Wrong environment for dir")
	}
}
