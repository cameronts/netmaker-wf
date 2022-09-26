package functions

import (
	"time"

	"github.com/netmakerio/netmaker/logger"
	"github.com/netmakerio/netmaker/netclient/daemon"
)

// Install - installs binary/daemon
func Install() error {
	daemon.Stop()
	if err := daemon.InstallDaemon(); err != nil {
		logger.Log(0, "error installing daemon", err.Error())
		return err
	}
	time.Sleep(time.Second * 5)
	return daemon.Restart()
}
