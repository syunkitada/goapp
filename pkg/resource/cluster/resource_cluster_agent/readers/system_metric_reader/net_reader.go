package system_metric_reader

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/config"
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
	Tw                        int64
	TwRecycled                int64
	TwKilled                  int64
	PawsActive                int64
	PawsEstab                 int64
	DelayedAcks               int64
	DelayedAckLocked          int64
	DelayedAckLost            int64
	ListenOverflows           int64
	ListenDrops               int64
	TcpHpHits                 int64
	TcpPureAcks               int64
	TcpHpAcks                 int64
	TcpRenoRecovery           int64
	TcpSackRecovery           int64
	TcpSackReneging           int64
	TcpSackReorder            int64
	TcpRenoReorder            int64
	TcpTsReorder              int64
	TcpFullUndo               int64
	TcpPartialUndo            int64
	TcpDsackUndo              int64
	TcpLossUndo               int64
	TcpLostRetransmit         int64
	TcpRenoFailures           int64
	TcpSackFailures           int64
	TcpLossFailures           int64
	TcpFastRetrans            int64
	TcpSlowStartRetrans       int64
	TcpTimeouts               int64
	TcpLossProbes             int64
	TcpLossProbeRecovery      int64
	TcpRenoRecoveryFail       int64
	TcpSackRecoveryFail       int64
	TcpRcvCollapsed           int64
	TcpBacklogCoalesce        int64
	TcpDsackOldSent           int64
	TcpDsackOfoSent           int64
	TcpDsackRecv              int64
	TcpDsackOfoRecv           int64
	TcpAbortOnData            int64
	TcpAbortOnClose           int64
	TcpAbortOnMemory          int64
	TcpAbortOnTimeout         int64
	TcpAbortOnLinger          int64
	TcpAbortFailed            int64
	TcpMemoryPressures        int64
	TcpMemoryPressuresChrono  int64
	TcpSackDiscard            int64
	TcpDsackIgnoredOld        int64
	TcpDsackIgnoredNoUndo     int64
	TcpSpuriousRTOs           int64
	TcpMd5NotFound            int64
	TcpMd5Unexpected          int64
	TcpMd5Failure             int64
	TcpSackShifted            int64
	TcpSackMerged             int64
	TcpSackShiftFallback      int64
	TcpBacklogDrop            int64
	PfMemallocDrop            int64
	TcpMinTtlDrop             int64
	TcpDeferAcceptDrop        int64
	IpReversePathFilter       int64
	TcpTimeWaitOverflow       int64
	TcpReqQFullDoCookies      int64
	TcpReqQFullDrop           int64
	TcpRetransFail            int64
	TcpRcvCoalesce            int64
	TcpOfoQueue               int64
	TcpOfoDrop                int64
	TcpOfoMerge               int64
	TcpChallengeACK           int64
	TcpSynChallenge           int64
	TcpFastOpenActive         int64
	TcpFastOpenActiveFail     int64
	TcpFastOpenPassive        int64
	TcpFastOpenPassiveFail    int64
	TcpFastOpenListenOverflow int64
	TcpFastOpenCookieReqd     int64
	TcpFastOpenBlackhole      int64
	TcpSpuriousRtxHostQueues  int64
	BusyPollRxPackets         int64
	TcpAutoCorking            int64
	TcpFromZeroWindowAdv      int64
	TcpToZeroWindowAdv        int64
	TcpWantZeroWindowAdv      int64
	TcpSynRetrans             int64
	TcpOrigDataSent           int64
	TcpHystartTrainDetect     int64
	TcpHystartTrainCwnd       int64
	TcpHystartDelayDetect     int64
	TcpHystartDelayCwnd       int64
	TcpAckSkippedSynRecv      int64
	TcpAckSkippedPAWS         int64
	TcpAckSkippedSeq          int64
	TcpAckSkippedFinWait2     int64
	TcpAckSkippedTimeWait     int64
	TcpAckSkippedChallenge    int64
	TcpWinProbe               int64
	TcpKeepAlive              int64
	TcpMtupFail               int64
	TcpMtupSuccess            int64
	TcpDelivered              int64
	TcpDeliveredCE            int64
	TcpAckCompressed          int64
	TcpZeroWindowDrop         int64
	TcpRcvQDrop               int64
	TcpWqueueTooBig           int64
	TcpFastOpenPassiveAltKey  int64
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
	Timestamp                 time.Time
	ReportStatus              int // 0, 1(GetReport), 2(Reported)
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
	Tw                        int64
	TwRecycled                int64
	TwKilled                  int64
	PawsActive                int64
	PawsEstab                 int64
	DelayedAcks               int64
	DelayedAckLocked          int64
	DelayedAckLost            int64
	ListenOverflows           int64
	ListenDrops               int64
	TcpHpHits                 int64
	TcpPureAcks               int64
	TcpHpAcks                 int64
	TcpRenoRecovery           int64
	TcpSackRecovery           int64
	TcpSackReneging           int64
	TcpSackReorder            int64
	TcpRenoReorder            int64
	TcpTsReorder              int64
	TcpFullUndo               int64
	TcpPartialUndo            int64
	TcpDsackUndo              int64
	TcpLossUndo               int64
	TcpLostRetransmit         int64
	TcpRenoFailures           int64
	TcpSackFailures           int64
	TcpLossFailures           int64
	TcpFastRetrans            int64
	TcpSlowStartRetrans       int64
	TcpTimeouts               int64
	TcpLossProbes             int64
	TcpLossProbeRecovery      int64
	TcpRenoRecoveryFail       int64
	TcpSackRecoveryFail       int64
	TcpRcvCollapsed           int64
	TcpBacklogCoalesce        int64
	TcpDsackOldSent           int64
	TcpDsackOfoSent           int64
	TcpDsackRecv              int64
	TcpDsackOfoRecv           int64
	TcpAbortOnData            int64
	TcpAbortOnClose           int64
	TcpAbortOnMemory          int64
	TcpAbortOnTimeout         int64
	TcpAbortOnLinger          int64
	TcpAbortFailed            int64
	TcpMemoryPressures        int64
	TcpMemoryPressuresChrono  int64
	TcpSackDiscard            int64
	TcpDsackIgnoredOld        int64
	TcpDsackIgnoredNoUndo     int64
	TcpSpuriousRTOs           int64
	TcpMd5NotFound            int64
	TcpMd5Unexpected          int64
	TcpMd5Failure             int64
	TcpSackShifted            int64
	TcpSackMerged             int64
	TcpSackShiftFallback      int64
	TcpBacklogDrop            int64
	PfMemallocDrop            int64
	TcpMinTtlDrop             int64
	TcpDeferAcceptDrop        int64
	IpReversePathFilter       int64
	TcpTimeWaitOverflow       int64
	TcpReqQFullDoCookies      int64
	TcpReqQFullDrop           int64
	TcpRetransFail            int64
	TcpRcvCoalesce            int64
	TcpOfoQueue               int64
	TcpOfoDrop                int64
	TcpOfoMerge               int64
	TcpChallengeACK           int64
	TcpSynChallenge           int64
	TcpFastOpenActive         int64
	TcpFastOpenActiveFail     int64
	TcpFastOpenPassive        int64
	TcpFastOpenPassiveFail    int64
	TcpFastOpenListenOverflow int64
	TcpFastOpenCookieReqd     int64
	TcpFastOpenBlackhole      int64
	TcpSpuriousRtxHostQueues  int64
	BusyPollRxPackets         int64
	TcpAutoCorking            int64
	TcpFromZeroWindowAdv      int64
	TcpToZeroWindowAdv        int64
	TcpWantZeroWindowAdv      int64
	TcpSynRetrans             int64
	TcpOrigDataSent           int64
	TcpHystartTrainDetect     int64
	TcpHystartTrainCwnd       int64
	TcpHystartDelayDetect     int64
	TcpHystartDelayCwnd       int64
	TcpAckSkippedSynRecv      int64
	TcpAckSkippedPAWS         int64
	TcpAckSkippedSeq          int64
	TcpAckSkippedFinWait2     int64
	TcpAckSkippedTimeWait     int64
	TcpAckSkippedChallenge    int64
	TcpWinProbe               int64
	TcpKeepAlive              int64
	TcpMtupFail               int64
	TcpMtupSuccess            int64
	TcpDelivered              int64
	TcpDeliveredCE            int64
	TcpAckCompressed          int64
	TcpZeroWindowDrop         int64
	TcpRcvQDrop               int64
	TcpWqueueTooBig           int64
	TcpFastOpenPassiveAltKey  int64
}

type IpExtStat struct {
	Timestamp       time.Time
	ReportStatus    int // 0, 1(GetReport), 2(Reported)
	InNoRoutes      int64
	InTruncatedPkts int64
	InCsumErrors    int64
}

type NetReader struct {
	conf          *config.ResourceMetricSystemConfig
	cacheLength   int
	tmpTcpExtStat *TmpTcpExtStat
	tmpIpExtStat  *TmpIpExtStat
	tcpExtStats   []TcpExtStat
	ipExtStats    []IpExtStat
}

func NewNetReader(conf *config.ResourceMetricSystemConfig) SubMetricReader {
	return &NetReader{
		conf:        conf,
		cacheLength: conf.CacheLength,
		tcpExtStats: make([]TcpExtStat, 0, conf.CacheLength),
		ipExtStats:  make([]IpExtStat, 0, conf.CacheLength),
	}
}

func (reader *NetReader) Read(tctx *logger.TraceContext) {
	timestamp := time.Now()

	if reader.tmpTcpExtStat == nil {
		reader.tmpTcpExtStat, reader.tmpIpExtStat = reader.readTmpNetStat(tctx)
	} else {
		tmpTcpExtStat, tmpIpExtStat := reader.readTmpNetStat(tctx)

		if len(reader.tcpExtStats) > reader.cacheLength {
			reader.tcpExtStats = reader.tcpExtStats[1:]
		}
		reader.tcpExtStats = append(reader.tcpExtStats, TcpExtStat{
			ReportStatus:   0,
			Timestamp:      timestamp,
			SyncookiesSent: tmpTcpExtStat.SyncookiesSent - reader.tmpTcpExtStat.SyncookiesSent,

			SyncookiesRecv:            tmpTcpExtStat.SyncookiesRecv - reader.tmpTcpExtStat.SyncookiesRecv,
			SyncookiesFailed:          tmpTcpExtStat.SyncookiesFailed - reader.tmpTcpExtStat.SyncookiesFailed,
			EmbryonicRsts:             tmpTcpExtStat.EmbryonicRsts - reader.tmpTcpExtStat.EmbryonicRsts,
			PruneCalled:               tmpTcpExtStat.PruneCalled - reader.tmpTcpExtStat.PruneCalled,
			RcvPruned:                 tmpTcpExtStat.RcvPruned - reader.tmpTcpExtStat.RcvPruned,
			OfoPruned:                 tmpTcpExtStat.OfoPruned - reader.tmpTcpExtStat.OfoPruned,
			OutOfWindowIcmps:          tmpTcpExtStat.OutOfWindowIcmps - reader.tmpTcpExtStat.OutOfWindowIcmps,
			LockDroppedIcmps:          tmpTcpExtStat.LockDroppedIcmps - reader.tmpTcpExtStat.LockDroppedIcmps,
			ArpFilter:                 tmpTcpExtStat.ArpFilter - reader.tmpTcpExtStat.ArpFilter,
			Tw:                        tmpTcpExtStat.Tw - reader.tmpTcpExtStat.Tw,
			TwRecycled:                tmpTcpExtStat.TwRecycled - reader.tmpTcpExtStat.TwRecycled,
			TwKilled:                  tmpTcpExtStat.TwKilled - reader.tmpTcpExtStat.TwKilled,
			PawsActive:                tmpTcpExtStat.PawsActive - reader.tmpTcpExtStat.PawsActive,
			PawsEstab:                 tmpTcpExtStat.PawsEstab - reader.tmpTcpExtStat.PawsEstab,
			DelayedAcks:               tmpTcpExtStat.DelayedAcks - reader.tmpTcpExtStat.DelayedAcks,
			DelayedAckLocked:          tmpTcpExtStat.DelayedAckLocked - reader.tmpTcpExtStat.DelayedAckLocked,
			DelayedAckLost:            tmpTcpExtStat.DelayedAckLost - reader.tmpTcpExtStat.DelayedAckLost,
			ListenOverflows:           tmpTcpExtStat.ListenOverflows - reader.tmpTcpExtStat.ListenOverflows,
			ListenDrops:               tmpTcpExtStat.ListenDrops - reader.tmpTcpExtStat.ListenDrops,
			TcpHpHits:                 tmpTcpExtStat.TcpHpHits - reader.tmpTcpExtStat.TcpHpHits,
			TcpPureAcks:               tmpTcpExtStat.TcpPureAcks - reader.tmpTcpExtStat.TcpPureAcks,
			TcpHpAcks:                 tmpTcpExtStat.TcpHpAcks - reader.tmpTcpExtStat.TcpHpAcks,
			TcpRenoRecovery:           tmpTcpExtStat.TcpRenoRecovery - reader.tmpTcpExtStat.TcpRenoRecovery,
			TcpSackRecovery:           tmpTcpExtStat.TcpSackRecovery - reader.tmpTcpExtStat.TcpSackRecovery,
			TcpSackReneging:           tmpTcpExtStat.TcpSackReneging - reader.tmpTcpExtStat.TcpSackReneging,
			TcpSackReorder:            tmpTcpExtStat.TcpSackReorder - reader.tmpTcpExtStat.TcpSackReorder,
			TcpRenoReorder:            tmpTcpExtStat.TcpRenoReorder - reader.tmpTcpExtStat.TcpRenoReorder,
			TcpTsReorder:              tmpTcpExtStat.TcpTsReorder - reader.tmpTcpExtStat.TcpTsReorder,
			TcpFullUndo:               tmpTcpExtStat.TcpFullUndo - reader.tmpTcpExtStat.TcpFullUndo,
			TcpPartialUndo:            tmpTcpExtStat.TcpPartialUndo - reader.tmpTcpExtStat.TcpPartialUndo,
			TcpDsackUndo:              tmpTcpExtStat.TcpDsackUndo - reader.tmpTcpExtStat.TcpDsackUndo,
			TcpLossUndo:               tmpTcpExtStat.TcpLossUndo - reader.tmpTcpExtStat.TcpLossUndo,
			TcpLostRetransmit:         tmpTcpExtStat.TcpLostRetransmit - reader.tmpTcpExtStat.TcpLostRetransmit,
			TcpRenoFailures:           tmpTcpExtStat.TcpRenoFailures - reader.tmpTcpExtStat.TcpRenoFailures,
			TcpSackFailures:           tmpTcpExtStat.TcpSackFailures - reader.tmpTcpExtStat.TcpSackFailures,
			TcpLossFailures:           tmpTcpExtStat.TcpLossFailures - reader.tmpTcpExtStat.TcpLossFailures,
			TcpFastRetrans:            tmpTcpExtStat.TcpFastRetrans - reader.tmpTcpExtStat.TcpFastRetrans,
			TcpSlowStartRetrans:       tmpTcpExtStat.TcpSlowStartRetrans - reader.tmpTcpExtStat.TcpSlowStartRetrans,
			TcpTimeouts:               tmpTcpExtStat.TcpTimeouts - reader.tmpTcpExtStat.TcpTimeouts,
			TcpLossProbes:             tmpTcpExtStat.TcpLossProbes - reader.tmpTcpExtStat.TcpLossProbes,
			TcpLossProbeRecovery:      tmpTcpExtStat.TcpLossProbeRecovery - reader.tmpTcpExtStat.TcpLossProbeRecovery,
			TcpRenoRecoveryFail:       tmpTcpExtStat.TcpRenoRecoveryFail - reader.tmpTcpExtStat.TcpRenoRecoveryFail,
			TcpSackRecoveryFail:       tmpTcpExtStat.TcpSackRecoveryFail - reader.tmpTcpExtStat.TcpSackRecoveryFail,
			TcpRcvCollapsed:           tmpTcpExtStat.TcpRcvCollapsed - reader.tmpTcpExtStat.TcpRcvCollapsed,
			TcpBacklogCoalesce:        tmpTcpExtStat.TcpBacklogCoalesce - reader.tmpTcpExtStat.TcpBacklogCoalesce,
			TcpDsackOldSent:           tmpTcpExtStat.TcpDsackOldSent - reader.tmpTcpExtStat.TcpDsackOldSent,
			TcpDsackOfoSent:           tmpTcpExtStat.TcpDsackOfoSent - reader.tmpTcpExtStat.TcpDsackOfoSent,
			TcpDsackRecv:              tmpTcpExtStat.TcpDsackRecv - reader.tmpTcpExtStat.TcpDsackRecv,
			TcpDsackOfoRecv:           tmpTcpExtStat.TcpDsackOfoRecv - reader.tmpTcpExtStat.TcpDsackOfoRecv,
			TcpAbortOnData:            tmpTcpExtStat.TcpAbortOnData - reader.tmpTcpExtStat.TcpAbortOnData,
			TcpAbortOnClose:           tmpTcpExtStat.TcpAbortOnClose - reader.tmpTcpExtStat.TcpAbortOnClose,
			TcpAbortOnMemory:          tmpTcpExtStat.TcpAbortOnMemory - reader.tmpTcpExtStat.TcpAbortOnMemory,
			TcpAbortOnTimeout:         tmpTcpExtStat.TcpAbortOnTimeout - reader.tmpTcpExtStat.TcpAbortOnTimeout,
			TcpAbortOnLinger:          tmpTcpExtStat.TcpAbortOnLinger - reader.tmpTcpExtStat.TcpAbortOnLinger,
			TcpAbortFailed:            tmpTcpExtStat.TcpAbortFailed - reader.tmpTcpExtStat.TcpAbortFailed,
			TcpMemoryPressures:        tmpTcpExtStat.TcpMemoryPressures - reader.tmpTcpExtStat.TcpMemoryPressures,
			TcpMemoryPressuresChrono:  tmpTcpExtStat.TcpMemoryPressuresChrono - reader.tmpTcpExtStat.TcpMemoryPressuresChrono,
			TcpSackDiscard:            tmpTcpExtStat.TcpSackDiscard - reader.tmpTcpExtStat.TcpSackDiscard,
			TcpDsackIgnoredOld:        tmpTcpExtStat.TcpDsackIgnoredOld - reader.tmpTcpExtStat.TcpDsackIgnoredOld,
			TcpDsackIgnoredNoUndo:     tmpTcpExtStat.TcpDsackIgnoredNoUndo - reader.tmpTcpExtStat.TcpDsackIgnoredNoUndo,
			TcpSpuriousRTOs:           tmpTcpExtStat.TcpSpuriousRTOs - reader.tmpTcpExtStat.TcpSpuriousRTOs,
			TcpMd5NotFound:            tmpTcpExtStat.TcpMd5NotFound - reader.tmpTcpExtStat.TcpMd5NotFound,
			TcpMd5Unexpected:          tmpTcpExtStat.TcpMd5Unexpected - reader.tmpTcpExtStat.TcpMd5Unexpected,
			TcpMd5Failure:             tmpTcpExtStat.TcpMd5Failure - reader.tmpTcpExtStat.TcpMd5Failure,
			TcpSackShifted:            tmpTcpExtStat.TcpSackShifted - reader.tmpTcpExtStat.TcpSackShifted,
			TcpSackMerged:             tmpTcpExtStat.TcpSackMerged - reader.tmpTcpExtStat.TcpSackMerged,
			TcpSackShiftFallback:      tmpTcpExtStat.TcpSackShiftFallback - reader.tmpTcpExtStat.TcpSackShiftFallback,
			TcpBacklogDrop:            tmpTcpExtStat.TcpBacklogDrop - reader.tmpTcpExtStat.TcpBacklogDrop,
			PfMemallocDrop:            tmpTcpExtStat.PfMemallocDrop - reader.tmpTcpExtStat.PfMemallocDrop,
			TcpMinTtlDrop:             tmpTcpExtStat.TcpMinTtlDrop - reader.tmpTcpExtStat.TcpMinTtlDrop,
			TcpDeferAcceptDrop:        tmpTcpExtStat.TcpDeferAcceptDrop - reader.tmpTcpExtStat.TcpDeferAcceptDrop,
			IpReversePathFilter:       tmpTcpExtStat.IpReversePathFilter - reader.tmpTcpExtStat.IpReversePathFilter,
			TcpTimeWaitOverflow:       tmpTcpExtStat.TcpTimeWaitOverflow - reader.tmpTcpExtStat.TcpTimeWaitOverflow,
			TcpReqQFullDoCookies:      tmpTcpExtStat.TcpReqQFullDoCookies - reader.tmpTcpExtStat.TcpReqQFullDoCookies,
			TcpReqQFullDrop:           tmpTcpExtStat.TcpReqQFullDrop - reader.tmpTcpExtStat.TcpReqQFullDrop,
			TcpRetransFail:            tmpTcpExtStat.TcpRetransFail - reader.tmpTcpExtStat.TcpRetransFail,
			TcpRcvCoalesce:            tmpTcpExtStat.TcpRcvCoalesce - reader.tmpTcpExtStat.TcpRcvCoalesce,
			TcpOfoQueue:               tmpTcpExtStat.TcpOfoQueue - reader.tmpTcpExtStat.TcpOfoQueue,
			TcpOfoDrop:                tmpTcpExtStat.TcpOfoDrop - reader.tmpTcpExtStat.TcpOfoDrop,
			TcpOfoMerge:               tmpTcpExtStat.TcpOfoMerge - reader.tmpTcpExtStat.TcpOfoMerge,
			TcpChallengeACK:           tmpTcpExtStat.TcpChallengeACK - reader.tmpTcpExtStat.TcpChallengeACK,
			TcpSynChallenge:           tmpTcpExtStat.TcpSynChallenge - reader.tmpTcpExtStat.TcpSynChallenge,
			TcpFastOpenActive:         tmpTcpExtStat.TcpFastOpenActive - reader.tmpTcpExtStat.TcpFastOpenActive,
			TcpFastOpenActiveFail:     tmpTcpExtStat.TcpFastOpenActiveFail - reader.tmpTcpExtStat.TcpFastOpenActiveFail,
			TcpFastOpenPassive:        tmpTcpExtStat.TcpFastOpenPassive - reader.tmpTcpExtStat.TcpFastOpenPassive,
			TcpFastOpenPassiveFail:    tmpTcpExtStat.TcpFastOpenPassiveFail - reader.tmpTcpExtStat.TcpFastOpenPassiveFail,
			TcpFastOpenListenOverflow: tmpTcpExtStat.TcpFastOpenListenOverflow - reader.tmpTcpExtStat.TcpFastOpenListenOverflow,
			TcpFastOpenCookieReqd:     tmpTcpExtStat.TcpFastOpenCookieReqd - reader.tmpTcpExtStat.TcpFastOpenCookieReqd,
			TcpFastOpenBlackhole:      tmpTcpExtStat.TcpFastOpenBlackhole - reader.tmpTcpExtStat.TcpFastOpenBlackhole,
			TcpSpuriousRtxHostQueues:  tmpTcpExtStat.TcpSpuriousRtxHostQueues - reader.tmpTcpExtStat.TcpSpuriousRtxHostQueues,
			BusyPollRxPackets:         tmpTcpExtStat.BusyPollRxPackets - reader.tmpTcpExtStat.BusyPollRxPackets,
			TcpAutoCorking:            tmpTcpExtStat.TcpAutoCorking - reader.tmpTcpExtStat.TcpAutoCorking,
			TcpFromZeroWindowAdv:      tmpTcpExtStat.TcpFromZeroWindowAdv - reader.tmpTcpExtStat.TcpFromZeroWindowAdv,
			TcpToZeroWindowAdv:        tmpTcpExtStat.TcpToZeroWindowAdv - reader.tmpTcpExtStat.TcpToZeroWindowAdv,
			TcpWantZeroWindowAdv:      tmpTcpExtStat.TcpWantZeroWindowAdv - reader.tmpTcpExtStat.TcpWantZeroWindowAdv,
			TcpSynRetrans:             tmpTcpExtStat.TcpSynRetrans - reader.tmpTcpExtStat.TcpSynRetrans,
			TcpOrigDataSent:           tmpTcpExtStat.TcpOrigDataSent - reader.tmpTcpExtStat.TcpOrigDataSent,
			TcpHystartTrainDetect:     tmpTcpExtStat.TcpHystartTrainDetect - reader.tmpTcpExtStat.TcpHystartTrainDetect,
			TcpHystartTrainCwnd:       tmpTcpExtStat.TcpHystartTrainCwnd - reader.tmpTcpExtStat.TcpHystartTrainCwnd,
			TcpHystartDelayDetect:     tmpTcpExtStat.TcpHystartDelayDetect - reader.tmpTcpExtStat.TcpHystartDelayDetect,
			TcpHystartDelayCwnd:       tmpTcpExtStat.TcpHystartDelayCwnd - reader.tmpTcpExtStat.TcpHystartDelayCwnd,
			TcpAckSkippedSynRecv:      tmpTcpExtStat.TcpAckSkippedSynRecv - reader.tmpTcpExtStat.TcpAckSkippedSynRecv,
			TcpAckSkippedPAWS:         tmpTcpExtStat.TcpAckSkippedPAWS - reader.tmpTcpExtStat.TcpAckSkippedPAWS,
			TcpAckSkippedSeq:          tmpTcpExtStat.TcpAckSkippedSeq - reader.tmpTcpExtStat.TcpAckSkippedSeq,
			TcpAckSkippedFinWait2:     tmpTcpExtStat.TcpAckSkippedFinWait2 - reader.tmpTcpExtStat.TcpAckSkippedFinWait2,
			TcpAckSkippedTimeWait:     tmpTcpExtStat.TcpAckSkippedTimeWait - reader.tmpTcpExtStat.TcpAckSkippedTimeWait,
			TcpAckSkippedChallenge:    tmpTcpExtStat.TcpAckSkippedChallenge - reader.tmpTcpExtStat.TcpAckSkippedChallenge,
			TcpWinProbe:               tmpTcpExtStat.TcpWinProbe - reader.tmpTcpExtStat.TcpWinProbe,
			TcpKeepAlive:              tmpTcpExtStat.TcpKeepAlive - reader.tmpTcpExtStat.TcpKeepAlive,
			TcpMtupFail:               tmpTcpExtStat.TcpMtupFail - reader.tmpTcpExtStat.TcpMtupFail,
			TcpMtupSuccess:            tmpTcpExtStat.TcpMtupSuccess - reader.tmpTcpExtStat.TcpMtupSuccess,
			TcpDelivered:              tmpTcpExtStat.TcpDelivered - reader.tmpTcpExtStat.TcpDelivered,
			TcpDeliveredCE:            tmpTcpExtStat.TcpDeliveredCE - reader.tmpTcpExtStat.TcpDeliveredCE,
			TcpAckCompressed:          tmpTcpExtStat.TcpAckCompressed - reader.tmpTcpExtStat.TcpAckCompressed,
			TcpZeroWindowDrop:         tmpTcpExtStat.TcpZeroWindowDrop - reader.tmpTcpExtStat.TcpZeroWindowDrop,
			TcpRcvQDrop:               tmpTcpExtStat.TcpRcvQDrop - reader.tmpTcpExtStat.TcpRcvQDrop,
			TcpWqueueTooBig:           tmpTcpExtStat.TcpWqueueTooBig - reader.tmpTcpExtStat.TcpWqueueTooBig,
			TcpFastOpenPassiveAltKey:  tmpTcpExtStat.TcpFastOpenPassiveAltKey - reader.tmpTcpExtStat.TcpFastOpenPassiveAltKey,
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

func (reader *NetReader) readTmpNetStat(tctx *logger.TraceContext) (tmpTcpExtStat *TmpTcpExtStat, tmpIpExtStat *TmpIpExtStat) {
	netstatFile, _ := os.Open("/proc/net/netstat")
	defer netstatFile.Close()
	tmpReader := bufio.NewReader(netstatFile)

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
		SyncookiesSent:            tcpExtMap["SyncookiesSent"],
		SyncookiesRecv:            tcpExtMap["SyncookiesRecv"],
		SyncookiesFailed:          tcpExtMap["SyncookiesFailed"],
		EmbryonicRsts:             tcpExtMap["EmbryonicRsts"],
		PruneCalled:               tcpExtMap["PruneCalled"],
		RcvPruned:                 tcpExtMap["RcvPruned"],
		OfoPruned:                 tcpExtMap["OfoPruned"],
		OutOfWindowIcmps:          tcpExtMap["OutOfWindowIcmps"],
		LockDroppedIcmps:          tcpExtMap["LockDroppedIcmps"],
		ArpFilter:                 tcpExtMap["ArpFilter"],
		Tw:                        tcpExtMap["TW"],
		TwRecycled:                tcpExtMap["TWRecycled"],
		TwKilled:                  tcpExtMap["TWKilled"],
		PawsActive:                tcpExtMap["PAWSActive"],
		PawsEstab:                 tcpExtMap["PAWSEstab"],
		DelayedAcks:               tcpExtMap["DelayedACKs"],
		DelayedAckLocked:          tcpExtMap["DelayedACKLocked"],
		DelayedAckLost:            tcpExtMap["DelayedACKLost"],
		ListenOverflows:           tcpExtMap["ListenOverflows"],
		ListenDrops:               tcpExtMap["ListenDrops"],
		TcpHpHits:                 tcpExtMap["TCPHPHits"],
		TcpPureAcks:               tcpExtMap["TCPPureAcks"],
		TcpHpAcks:                 tcpExtMap["TCPHPAcks"],
		TcpRenoRecovery:           tcpExtMap["TCPRenoRecovery"],
		TcpSackRecovery:           tcpExtMap["TCPSackRecovery"],
		TcpSackReneging:           tcpExtMap["TCPSACKReneging"],
		TcpSackReorder:            tcpExtMap["TCPSACKReorder"],
		TcpRenoReorder:            tcpExtMap["TCPRenoReorder"],
		TcpTsReorder:              tcpExtMap["TCPTSReorder"],
		TcpFullUndo:               tcpExtMap["TCPFullUndo"],
		TcpPartialUndo:            tcpExtMap["TCPPartialUndo"],
		TcpDsackUndo:              tcpExtMap["TCPDSACKUndo"],
		TcpLossUndo:               tcpExtMap["TCPLossUndo"],
		TcpLostRetransmit:         tcpExtMap["TCPLostRetransmit"],
		TcpRenoFailures:           tcpExtMap["TCPRenoFailures"],
		TcpSackFailures:           tcpExtMap["TCPSackFailures"],
		TcpLossFailures:           tcpExtMap["TCPLossFailures"],
		TcpFastRetrans:            tcpExtMap["TCPFastRetrans"],
		TcpSlowStartRetrans:       tcpExtMap["TCPSlowStartRetrans"],
		TcpTimeouts:               tcpExtMap["TCPTimeouts"],
		TcpLossProbes:             tcpExtMap["TCPLossProbes"],
		TcpLossProbeRecovery:      tcpExtMap["TCPLossProbeRecovery"],
		TcpRenoRecoveryFail:       tcpExtMap["TCPRenoRecoveryFail"],
		TcpSackRecoveryFail:       tcpExtMap["TCPSackRecoveryFail"],
		TcpRcvCollapsed:           tcpExtMap["TCPRcvCollapsed"],
		TcpBacklogCoalesce:        tcpExtMap["TCPBacklogCoalesce"],
		TcpDsackOldSent:           tcpExtMap["TCPDSACKOldSent"],
		TcpDsackOfoSent:           tcpExtMap["TCPDSACKOfoSent"],
		TcpDsackRecv:              tcpExtMap["TCPDSACKRecv"],
		TcpDsackOfoRecv:           tcpExtMap["TCPDSACKOfoRecv"],
		TcpAbortOnData:            tcpExtMap["TCPAbortOnData"],
		TcpAbortOnClose:           tcpExtMap["TCPAbortOnClose"],
		TcpAbortOnMemory:          tcpExtMap["TCPAbortOnMemory"],
		TcpAbortOnTimeout:         tcpExtMap["TCPAbortOnTimeout"],
		TcpAbortOnLinger:          tcpExtMap["TCPAbortOnLinger"],
		TcpAbortFailed:            tcpExtMap["TCPAbortFailed"],
		TcpMemoryPressures:        tcpExtMap["TCPMemoryPressures"],
		TcpMemoryPressuresChrono:  tcpExtMap["TCPMemoryPressuresChrono"],
		TcpSackDiscard:            tcpExtMap["TCPSACKDiscard"],
		TcpDsackIgnoredOld:        tcpExtMap["TCPDSACKIgnoredOld"],
		TcpDsackIgnoredNoUndo:     tcpExtMap["TCPDSACKIgnoredNoUndo"],
		TcpSpuriousRTOs:           tcpExtMap["TCPSpuriousRTOs"],
		TcpMd5NotFound:            tcpExtMap["TCPMD5NotFound"],
		TcpMd5Unexpected:          tcpExtMap["TCPMD5Unexpected"],
		TcpMd5Failure:             tcpExtMap["TCPMD5Failure"],
		TcpSackShifted:            tcpExtMap["TCPSackShifted"],
		TcpSackMerged:             tcpExtMap["TCPSackMerged"],
		TcpSackShiftFallback:      tcpExtMap["TCPSackShiftFallback"],
		TcpBacklogDrop:            tcpExtMap["TCPBacklogDrop"],
		PfMemallocDrop:            tcpExtMap["PFMemallocDrop"],
		TcpMinTtlDrop:             tcpExtMap["TCPMinTTLDrop"],
		TcpDeferAcceptDrop:        tcpExtMap["TCPDeferAcceptDrop"],
		IpReversePathFilter:       tcpExtMap["IPReversePathFilter"],
		TcpTimeWaitOverflow:       tcpExtMap["TCPTimeWaitOverflow"],
		TcpReqQFullDoCookies:      tcpExtMap["TCPReqQFullDoCookies"],
		TcpReqQFullDrop:           tcpExtMap["TCPReqQFullDrop"],
		TcpRetransFail:            tcpExtMap["TCPRetransFail"],
		TcpRcvCoalesce:            tcpExtMap["TCPRcvCoalesce"],
		TcpOfoQueue:               tcpExtMap["TCPOFOQueue"],
		TcpOfoDrop:                tcpExtMap["TCPOFODrop"],
		TcpOfoMerge:               tcpExtMap["TCPOFOMerge"],
		TcpChallengeACK:           tcpExtMap["TCPChallengeACK"],
		TcpSynChallenge:           tcpExtMap["TCPSYNChallenge"],
		TcpFastOpenActive:         tcpExtMap["TCPFastOpenActive"],
		TcpFastOpenActiveFail:     tcpExtMap["TCPFastOpenActiveFail"],
		TcpFastOpenPassive:        tcpExtMap["TCPFastOpenPassive"],
		TcpFastOpenPassiveFail:    tcpExtMap["TCPFastOpenPassiveFail"],
		TcpFastOpenListenOverflow: tcpExtMap["TCPFastOpenListenOverflow"],
		TcpFastOpenCookieReqd:     tcpExtMap["TCPFastOpenCookieReqd"],
		TcpFastOpenBlackhole:      tcpExtMap["TCPFastOpenBlackhole"],
		TcpSpuriousRtxHostQueues:  tcpExtMap["TCPSpuriousRtxHostQueues"],
		BusyPollRxPackets:         tcpExtMap["BusyPollRxPackets"],
		TcpAutoCorking:            tcpExtMap["TCPAutoCorking"],
		TcpFromZeroWindowAdv:      tcpExtMap["TCPFromZeroWindowAdv"],
		TcpToZeroWindowAdv:        tcpExtMap["TCPToZeroWindowAdv"],
		TcpWantZeroWindowAdv:      tcpExtMap["TCPWantZeroWindowAdv"],
		TcpSynRetrans:             tcpExtMap["TCPSynRetrans"],
		TcpOrigDataSent:           tcpExtMap["TCPOrigDataSent"],
		TcpHystartTrainDetect:     tcpExtMap["TCPHystartTrainDetect"],
		TcpHystartTrainCwnd:       tcpExtMap["TCPHystartTrainCwnd"],
		TcpHystartDelayDetect:     tcpExtMap["TCPHystartDelayDetect"],
		TcpHystartDelayCwnd:       tcpExtMap["TCPHystartDelayCwnd"],
		TcpAckSkippedSynRecv:      tcpExtMap["TCPACKSkippedSynRecv"],
		TcpAckSkippedPAWS:         tcpExtMap["TCPACKSkippedPAWS"],
		TcpAckSkippedSeq:          tcpExtMap["TCPACKSkippedSeq"],
		TcpAckSkippedFinWait2:     tcpExtMap["TCPACKSkippedFinWait2"],
		TcpAckSkippedTimeWait:     tcpExtMap["TCPACKSkippedTimeWait"],
		TcpAckSkippedChallenge:    tcpExtMap["TCPACKSkippedChallenge"],
		TcpWinProbe:               tcpExtMap["TCPWinProbe"],
		TcpKeepAlive:              tcpExtMap["TCPKeepAlive"],
		TcpMtupFail:               tcpExtMap["TCPMTUPFail"],
		TcpMtupSuccess:            tcpExtMap["TCPMTUPSuccess"],
		TcpDelivered:              tcpExtMap["TCPDelivered"],
		TcpDeliveredCE:            tcpExtMap["TCPDeliveredCE"],
		TcpAckCompressed:          tcpExtMap["TCPAckCompressed"],
		TcpZeroWindowDrop:         tcpExtMap["TCPZeroWindowDrop"],
		TcpRcvQDrop:               tcpExtMap["TCPRcvQDrop"],
		TcpWqueueTooBig:           tcpExtMap["TCPWqueueTooBig"],
		TcpFastOpenPassiveAltKey:  tcpExtMap["TCPFastOpenPassiveAltKey"],
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
		InCsumErrors:    ipExtMap["InCsumErrors"],
	}

	return
}

func (reader *NetReader) ReportMetrics() (metrics []spec.ResourceMetric) {
	metrics = make([]spec.ResourceMetric, 0, len(reader.tcpExtStats)+len(reader.ipExtStats))
	for _, stat := range reader.tcpExtStats {
		if stat.ReportStatus == ReportStatusReported {
			continue
		}

		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_tcp_netstat",
			Time: stat.Timestamp,
			Metric: map[string]interface{}{
				"tw":               stat.Tw,
				"tw_recycled":      stat.TwRecycled,
				"abort":            stat.TcpAbortOnData + stat.TcpAbortOnClose + stat.TcpAbortOnMemory + stat.TcpAbortOnTimeout + stat.TcpAbortOnLinger,
				"abort_failed":     stat.TcpAbortFailed,
				"retrans":          stat.TcpSynRetrans + stat.TcpLostRetransmit + stat.TcpFastRetrans + stat.TcpSlowStartRetrans,
				"retrans_failed":   stat.TcpRetransFail,
				"drops":            stat.TcpBacklogDrop + stat.PfMemallocDrop + stat.TcpMinTtlDrop + stat.TcpReqQFullDrop + stat.TcpRcvQDrop,
				"listen_drops":     stat.ListenDrops,
				"listen_overflows": stat.ListenOverflows,
				"delayed_acks":     stat.DelayedAcks,
			},
		})
	}

	for _, stat := range reader.ipExtStats {
		if stat.ReportStatus == ReportStatusReported {
			continue
		}
		metrics = append(metrics, spec.ResourceMetric{
			Name: "system_ip_netstat",
			Time: stat.Timestamp,
			Metric: map[string]interface{}{
				"in_no_routes":      stat.InNoRoutes,
				"in_truncated_pkts": stat.InTruncatedPkts,
				"in_csum_errors":    stat.InCsumErrors,
			},
		})
	}

	return
}

func (reader *NetReader) ReportEvents() (events []spec.ResourceEvent) {
	return
}

func (reader *NetReader) Reported() {
	for i := range reader.tcpExtStats {
		reader.tcpExtStats[i].ReportStatus = ReportStatusReported
	}
	for i := range reader.ipExtStats {
		reader.ipExtStats[i].ReportStatus = ReportStatusReported
	}
	return
}
