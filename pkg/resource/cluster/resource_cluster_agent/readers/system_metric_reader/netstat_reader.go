package system_metric_reader

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type TmpTcpExtStat struct {
	SyncookiesSent            int64
	SyncookiesRecv            int64
	SyncookiesFailed          int64
	EmbryonicRsts             int64
	PruneCalled               int64
	RcvPruned                 int64
	OfoPruned                 int64
	OutOfWindowIcmps          int64
	LockDroppedIcmps          int64
	ArpFilter                 int64
	TW                        int64
	TWRecycled                int64
	TWKilled                  int64
	PAWSActive                int64
	PAWSEstab                 int64
	DelayedACKs               int64
	DelayedACKLocked          int64
	DelayedACKLost            int64
	ListenOverflows           int64
	ListenDrops               int64
	TCPHPHits                 int64
	TCPPureAcks               int64
	TCPHPAcks                 int64
	TCPRenoRecovery           int64
	TCPSackRecovery           int64
	TCPSACKReneging           int64
	TCPSACKReorder            int64
	TCPRenoReorder            int64
	TCPTSReorder              int64
	TCPFullUndo               int64
	TCPPartialUndo            int64
	TCPDSACKUndo              int64
	TCPLossUndo               int64
	TCPLostRetransmit         int64
	TCPRenoFailures           int64
	TCPSackFailures           int64
	TCPLossFailures           int64
	TCPFastRetrans            int64
	TCPSlowStartRetrans       int64
	TCPTimeouts               int64
	TCPLossProbes             int64
	TCPLossProbeRecovery      int64
	TCPRenoRecoveryFail       int64
	TCPSackRecoveryFail       int64
	TCPRcvCollapsed           int64
	TCPBacklogCoalesce        int64
	TCPDSACKOldSent           int64
	TCPDSACKOfoSent           int64
	TCPDSACKRecv              int64
	TCPDSACKOfoRecv           int64
	TCPAbortOnData            int64
	TCPAbortOnClose           int64
	TCPAbortOnMemory          int64
	TCPAbortOnTimeout         int64
	TCPAbortOnLinger          int64
	TCPAbortFailed            int64
	TCPMemoryPressures        int64
	TCPMemoryPressuresChrono  int64
	TCPSACKDiscard            int64
	TCPDSACKIgnoredOld        int64
	TCPDSACKIgnoredNoUndo     int64
	TCPSpuriousRTOs           int64
	TCPMD5NotFound            int64
	TCPMD5Unexpected          int64
	TCPMD5Failure             int64
	TCPSackShifted            int64
	TCPSackMerged             int64
	TCPSackShiftFallback      int64
	TCPBacklogDrop            int64
	PFMemallocDrop            int64
	TCPMinTTLDrop             int64
	TCPDeferAcceptDrop        int64
	IPReversePathFilter       int64
	TCPTimeWaitOverflow       int64
	TCPReqQFullDoCookies      int64
	TCPReqQFullDrop           int64
	TCPRetransFail            int64
	TCPRcvCoalesce            int64
	TCPOFOQueue               int64
	TCPOFODrop                int64
	TCPOFOMerge               int64
	TCPChallengeACK           int64
	TCPSYNChallenge           int64
	TCPFastOpenActive         int64
	TCPFastOpenActiveFail     int64
	TCPFastOpenPassive        int64
	TCPFastOpenPassiveFail    int64
	TCPFastOpenListenOverflow int64
	TCPFastOpenCookieReqd     int64
	TCPFastOpenBlackhole      int64
	TCPSpuriousRtxHostQueues  int64
	BusyPollRxPackets         int64
	TCPAutoCorking            int64
	TCPFromZeroWindowAdv      int64
	TCPToZeroWindowAdv        int64
	TCPWantZeroWindowAdv      int64
	TCPSynRetrans             int64
	TCPOrigDataSent           int64
	TCPHystartTrainDetect     int64
	TCPHystartTrainCwnd       int64
	TCPHystartDelayDetect     int64
	TCPHystartDelayCwnd       int64
	TCPACKSkippedSynRecv      int64
	TCPACKSkippedPAWS         int64
	TCPACKSkippedSeq          int64
	TCPACKSkippedFinWait2     int64
	TCPACKSkippedTimeWait     int64
	TCPACKSkippedChallenge    int64
	TCPWinProbe               int64
	TCPKeepAlive              int64
	TCPMTUPFail               int64
	TCPMTUPSuccess            int64
	TCPDelivered              int64
	TCPDeliveredCE            int64
	TCPAckCompressed          int64
	TCPZeroWindowDrop         int64
	TCPRcvQDrop               int64
	TCPWqueueTooBig           int64
	TCPFastOpenPassiveAltKey  int64
}

type TmpIpExtStat struct {
	InNoRoutes      int64
	InTruncatedPkts int64
	InMcastPkts     int64
	OutMcastPkts    int64
	InBcastPkts     int64
	OutBcastPkts    int64
	InOctets        int64
	OutOctets       int64
	InMcastOctets   int64
	OutMcastOctets  int64
	InBcastOctets   int64
	OutBcastOctets  int64
	InCsumErrors    int64
	InNoECTPkts     int64
	InECT1Pkts      int64
	InECT0Pkts      int64
	InCEPkts        int64
	ReasmOverlaps   int64
}

type TcpExtStat struct {
	Timestamp        time.Time
	ReportStatus     int // 0, 1(GetReport), 2(Reported)
	SyncookiesSent   int64
	SyncookiesRecv   int64
	SyncookiesFailed int64
}

type IpExtStat struct {
	Timestamp       time.Time
	ReportStatus    int // 0, 1(GetReport), 2(Reported)
	InNoRoutes      int64
	InTruncatedPkts int64
}

func (reader *SystemMetricReader) ReadNetStat(tctx *logger.TraceContext) {
	timestamp := time.Now()

	if reader.tmpTcpExtStat == nil {
		reader.tmpTcpExtStat, reader.tmpIpExtStat = reader.ReadTmpNetStat(tctx)
	} else {
		tmpTcpExtStat, tmpIpExtStat := reader.ReadTmpNetStat(tctx)

		if len(reader.tcpExtStats) > reader.cacheLength {
			reader.tcpExtStats = reader.tcpExtStats[1:]
		}
		reader.tcpExtStats = append(reader.tcpExtStats, TcpExtStat{
			ReportStatus:     0,
			Timestamp:        timestamp,
			SyncookiesSent:   tmpTcpExtStat.SyncookiesSent - reader.tmpTcpExtStat.SyncookiesSent,
			SyncookiesRecv:   tmpTcpExtStat.SyncookiesRecv - reader.tmpTcpExtStat.SyncookiesRecv,
			SyncookiesFailed: tmpTcpExtStat.SyncookiesFailed - reader.tmpTcpExtStat.SyncookiesFailed,
		})

		if len(reader.ipExtStats) > reader.cacheLength {
			reader.ipExtStats = reader.ipExtStats[1:]
		}
		reader.ipExtStats = append(reader.ipExtStats, IpExtStat{
			ReportStatus:    0,
			Timestamp:       timestamp,
			InNoRoutes:      tmpIpExtStat.InNoRoutes - reader.tmpIpExtStat.InNoRoutes,
			InTruncatedPkts: tmpIpExtStat.InTruncatedPkts - reader.tmpIpExtStat.InTruncatedPkts,
		})

		reader.tmpTcpExtStat, reader.tmpIpExtStat = tmpTcpExtStat, tmpIpExtStat
	}
	return
}

func (reader *SystemMetricReader) ReadTmpNetStat(tctx *logger.TraceContext) (tmpTcpExtStat *TmpTcpExtStat, tmpIpExtStat *TmpIpExtStat) {
	netstatFile, _ := os.Open("/proc/net/netstat")
	defer netstatFile.Close()
	tmpReader := bufio.NewReader(netstatFile)

	// tcpExt
	tmpBytes, _, _ := tmpReader.ReadLine()
	tcpExtKeys := strings.Split(string(tmpBytes), " ")
	lenKeys := len(tcpExtKeys)

	tmpBytes, _, _ = tmpReader.ReadLine()
	tcpExtValues := strings.Split(string(tmpBytes), " ")

	tcpExtMap := map[string]int64{}
	for i := 1; i < lenKeys; i++ {
		tcpExtMap[tcpExtKeys[i]], _ = strconv.ParseInt(tcpExtValues[i], 10, 64)
	}

	tmpTcpExtStat = &TmpTcpExtStat{
		SyncookiesSent:   tcpExtMap["SyncookiesSent"],
		SyncookiesRecv:   tcpExtMap["SyncookiesRecv"],
		SyncookiesFailed: tcpExtMap["SyncookiesFailed"],
	}

	// ipExt
	tmpBytes, _, _ = tmpReader.ReadLine()
	ipExtKeys := strings.Split(string(tmpBytes), " ")
	lenKeys = len(ipExtKeys)

	tmpBytes, _, _ = tmpReader.ReadLine()
	ipExtValues := strings.Split(string(tmpBytes), " ")

	ipExtMap := map[string]int64{}
	for i := 1; i < lenKeys; i++ {
		ipExtMap[ipExtKeys[i]], _ = strconv.ParseInt(ipExtValues[i], 10, 64)
	}

	tmpIpExtStat = &TmpIpExtStat{
		InNoRoutes:      ipExtMap["InNoRoutes"],
		InTruncatedPkts: ipExtMap["InTruncatedPkts"],
	}

	return
}

func (reader *SystemMetricReader) GetNetStatMetrics() (metrics []spec.ResourceMetric) {
	metrics = make([]spec.ResourceMetric, len(reader.tcpExtStats)+len(reader.ipExtStats))
	for _, stat := range reader.tcpExtStats {
		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_tcp_netstat",
			Time: stat.Timestamp,
			Metric: map[string]interface{}{
				"syncookies_sent":   stat.SyncookiesSent,
				"syncookies_recv":   stat.SyncookiesRecv,
				"syncookies_failed": stat.SyncookiesFailed,
			},
		})
	}

	for _, stat := range reader.ipExtStats {
		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_tcp_netstat",
			Time: stat.Timestamp,
			Metric: map[string]interface{}{
				"in_no_routes":      stat.InNoRoutes,
				"in_truncated_pkts": stat.InTruncatedPkts,
			},
		})
	}
	return
}
