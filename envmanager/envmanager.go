package envmanager

import (
	"fmt"
	"os/exec"
)

type EnvManager interface {
	Create() error
}

type VEnv struct {
	path string
}

func (v VEnv) Create() error {
	cmd := exec.Command("python3", "-m", "venv", v.path)

	return cmd.Run()
}

func NewEnvManager(name, path string) (EnvManager, error) {
	if name == "venv" {
		return VEnv{path: path}, nil
	}

	return nil, fmt.Errorf("invalid environment manager")
}
