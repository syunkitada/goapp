package qemu_driver

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"path/filepath"

	"github.com/gorilla/websocket"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_models"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type QemuDriver struct {
	conf *config.ResourceComputeExConfig
	name string
}

func New(conf *config.ResourceComputeExConfig) *QemuDriver {
	driver := QemuDriver{
		conf: conf,
		name: "qemu",
	}
	return &driver
}

func (driver *QemuDriver) GetName() string {
	return ""
}

func (driver *QemuDriver) Deploy(tctx *logger.TraceContext) error {
	return nil
}

func (driver *QemuDriver) ConfirmDeploy(tctx *logger.TraceContext) (bool, error) {
	return false, nil
}

func (driver *QemuDriver) SyncActivatingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]spec.ComputeAssignmentEx,
	computeNetnsPortsMap map[uint][]compute_models.NetnsPort) error {
	return driver.syncActivatingAssignmentMap(tctx, assignmentMap, computeNetnsPortsMap)
}

func (driver *QemuDriver) ConfirmActivatingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]spec.ComputeAssignmentEx) (bool, error) {
	return true, nil
}

func (driver *QemuDriver) SyncDeletingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]spec.ComputeAssignmentEx) error {
	return driver.syncDeletingAssignmentMap(tctx, assignmentMap)
}

func (driver *QemuDriver) ConfirmDeletingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]spec.ComputeAssignmentEx) (bool, error) {
	return true, nil
}

type ConsoleInput struct {
	Code  int
	Shift bool
	Ctrl  bool
	Alt   bool
	Value string
}

type ConsoleOutput struct {
	Code  int
	Shift bool
	Ctrl  bool
	Alt   bool
	Value string
}

func (driver *QemuDriver) ProxyConsole(tctx *logger.TraceContext, input *spec.GetComputeConsole, conn *websocket.Conn) (err error) {
	defer func() {
		if tmpErr := conn.Close(); tmpErr != nil {
			logger.Warningf(tctx, "Failed websocket Close: %s", tmpErr.Error())
		}
	}()

	vmDir := filepath.Join(driver.conf.VmsDir, input.Name)
	vmSerialSocketPath := filepath.Join(vmDir, "serial.sock")
	fmt.Println("DEBUG socket", input.Name, vmSerialSocketPath)
	var c net.Conn
	c, err = net.Dial("unix", vmSerialSocketPath)
	if err != nil {
		return
	}
	defer c.Close()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := c.Read(buf[:])
			if err != nil {
				return
			}
			println("Client got:", string(buf[0:n]))

			output := ConsoleOutput{
				Value: string(buf[0:n]),
			}
			var bytes []byte
			bytes, err = json.Marshal(&output)
			if err = conn.WriteMessage(websocket.TextMessage, bytes); err != nil {
				logger.Warningf(tctx, "Faild WriteMessage: %s", err.Error())
				return
			}
		}
	}()

	var message []byte
	for {
		fmt.Println("Waiting Messages on WebSocket: ", input.Name)
		_, message, err = conn.ReadMessage()
		if err != nil {
			logger.Warningf(tctx, "Faild ReadMessage: %s", err.Error())
			return
		}
		var i ConsoleInput
		if err = json.Unmarshal(message, &i); err != nil {
			return
		}
		fmt.Println("DEBUG message", i, i.Value)

		_, err = c.Write([]byte(i.Value))
		if err != nil {
			log.Fatal("Write error:", err)
			break
		}

		// fmt.Println("DEBUG message", messageType, string(message))
		// if err = conn.WriteMessage(messageType, message); err != nil {
		// 	logger.Warningf(tctx, "Faild WriteMessage: %s", err.Error())
		// 	return
		// }
	}

	return
}
