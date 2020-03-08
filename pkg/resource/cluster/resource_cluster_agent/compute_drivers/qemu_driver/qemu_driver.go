package qemu_driver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"path/filepath"
	"sync"
	"time"

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
	Bytes [][]int
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
	defer func() {
		// This is not closed when websocket closed
		err = c.Close()
		fmt.Println("DEBUG Serial Socket CLOSED", err)
	}()

	chMutex := sync.Mutex{}
	isDone := false
	doneCh := make(chan bool, 2)
	readCh := make(chan []byte, 10)

	defer func() {
		chMutex.Lock()
		isDone = true
		close(doneCh)
		close(readCh)
		chMutex.Unlock()
	}()

	var message []byte
	go func() {
		for {
			fmt.Println("Waiting Messages on WebSocket: ", input.Name)
			_, message, err = conn.ReadMessage()
			if err != nil {
				chMutex.Lock()
				if !isDone {
					logger.Warningf(tctx, "Faild ReadMessage: %s", err.Error())
					doneCh <- true
				}
				chMutex.Unlock()
				return
			}
			var i ConsoleInput
			if err = json.Unmarshal(message, &i); err != nil {
				chMutex.Lock()
				if !isDone {
					logger.Warningf(tctx, "Faild Unmarshal: %s", err.Error())
					doneCh <- true
				}
				chMutex.Unlock()
				return
			}
			fmt.Println("DEBUG message", i, i.Value)

			_, err = c.Write([]byte(i.Value))
			if err != nil {
				chMutex.Lock()
				if !isDone {
					logger.Warningf(tctx, "Faild Write SerialSocket: %s", err.Error())
					doneCh <- true
				}
				chMutex.Unlock()
				return
			}
		}
	}()

	go func() {
		for {
			buf := make([]byte, 1024)
			n, tmpErr := c.Read(buf[:])
			if tmpErr != nil {
				chMutex.Lock()
				if !isDone {
					err = fmt.Errorf("Failed Read: err=%s", tmpErr.Error())
					doneCh <- true
				}
				chMutex.Unlock()
				return
			}
			fmt.Println("DEBUG READ", string(buf[0:n]))
			fmt.Println("DEBUG READ2", buf[0:n])
			// stringではなくbyteで管理して、js側で制御させたほうがよさそう
			// 8, 27, 91, 74
			// ここでとっても分割される[8, 27], [91, 74]
			// if string(buf[0:n]) == string([]byte{8, 27, 91, 74}) {
			// 	fmt.Println("DEBUG backspace")
			// 	readCh <- "\\d"
			// 	continue
			// }
			// fmt.Fprint(outputLogfile, string(buf[0:n]))
			readCh <- buf[0:n]
		}
	}()

	// 逐次出力せずに、バッファしてから出力する
	// 10msec 出力が途切れたら(timeoutしたら 、まとめて出力する
	var strs [][]byte
	timeout := time.Duration(60 * time.Second)
	for {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		select {
		case <-doneCh:
			cancel()
			log.Printf("\nExit by doneCh\n")
			return
		case str := <-readCh:
			cancel()
			strs = append(strs, str)
			timeout = time.Duration(10 * time.Millisecond)
		case <-ctx.Done():
			cancel()
			value := ""
			var outbytes [][]int
			for _, str := range strs {
				value += string(str)
				var ibyte []int
				for _, st := range str {
					ibyte = append(ibyte, int(st))
				}
				fmt.Println("DEBUG strs", str)
				fmt.Println("DEBUG byte", ibyte)
				outbytes = append(outbytes, ibyte)
			}
			output := ConsoleOutput{
				Value: value,
				Bytes: outbytes,
			}
			var bytes []byte
			bytes, err = json.Marshal(&output)
			if err = conn.WriteMessage(websocket.TextMessage, bytes); err != nil {
				logger.Warningf(tctx, "Faild WriteMessage: %s", err.Error())
				return
			}

			strs = [][]byte{}
			timeout = time.Duration(60 * time.Second)
		}
	}
}
