package manager

import (
	"apiLunchLite/internal/utils"
	"fmt"
	"os/exec"
)

func (m *ApiManager) InstallOneLib(name, lib, version string) error {
	result, err := m.Db.GetName(name)
	if err != nil {
		return err
	}

	if result.Pid != 0 {
		err := m.Stop(name, false)
		if err != nil {
			return nil
		}
	}

	pythonPath := utils.SearchPythonExe(result.PathFolder)
	libVersion := fmt.Sprintf("%s==%s", lib, version)
	cmd := exec.Command(pythonPath, "-m", "pip", "install", libVersion)
	cmd.Run()

	return nil
}

func (m *ApiManager) InstallVariousLib(name string, libList []string) error {
	result, err := m.Db.GetName(name)
	if err != nil {
		return err
	}

	if result.Pid != 0 {
		err := m.Stop(name, false)
		if err != nil {
			return nil
		}
	}

	pythonPath := utils.SearchPythonExe(result.PathFolder)

	for _, lib := range libList {
		cmdUpdate := exec.Command(pythonPath, "-m", "pip", "install", lib)
		cmdUpdate.Run()
	}
	return nil
}
