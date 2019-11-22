/*
 * @file
 * @copyright defined in aergo/LICENSE.txt
 */

package p2putil

import (
	"fmt"
	"github.com/aergoio/aergo-lib/log"
	"github.com/aergoio/aergo/p2p/p2pcommon"
	"github.com/aergoio/aergo/types"
	"github.com/rs/zerolog"
	"net"
)


// Deprecated
func DebugLogReceiveMsg(logger *log.Logger, protocol p2pcommon.SubProtocol, msgID string, peer p2pcommon.RemotePeer, additional interface{}) {
	if additional != nil {
		logger.Debug().Str(LogProtoID, protocol.String()).Str(LogMsgID, msgID).Str("from_peer", peer.Name()).Str("other", fmt.Sprint(additional)).
			Msg("Received a message")
	} else {
		logger.Debug().Str(LogProtoID, protocol.String()).Str(LogMsgID, msgID).Str("from_peer", peer.Name()).
			Msg("Received a message")
	}
}

// Deprecated
func DebugLogReceiveResponseMsg(logger *log.Logger, protocol p2pcommon.SubProtocol, msgID string, reqID string, peer p2pcommon.RemotePeer, additional interface{}) {
	if additional != nil {
		logger.Debug().Str(LogProtoID, protocol.String()).Str(LogMsgID, msgID).Str(LogOrgReqID, reqID).Str("from_peer", peer.Name()).Str("other", fmt.Sprint(additional)).
			Msg("Received a response message")
	} else {
		logger.Debug().Str(LogProtoID, protocol.String()).Str(LogMsgID, msgID).Str(LogOrgReqID, reqID).Str("from_peer", peer.Name()).
			Msg("Received a response message")
	}
}

func DebugLogReceive(logger *log.Logger, protocol p2pcommon.SubProtocol, msgID string, peer p2pcommon.RemotePeer, additional zerolog.LogObjectMarshaler) {
	if additional != nil {
		logger.Debug().Str(LogProtoID, protocol.String()).Str(LogMsgID, msgID).Str("from_peer", peer.Name()).Object("body", additional).Msg("Received a message")
	} else {
		logger.Debug().Str(LogProtoID, protocol.String()).Str(LogMsgID, msgID).Str("from_peer", peer.Name()).
			Msg("Received a message")
	}
}

func DebugLogReceiveResponse(logger *log.Logger, protocol p2pcommon.SubProtocol, msgID string, reqID string, peer p2pcommon.RemotePeer, additional zerolog.LogObjectMarshaler) {
	if additional != nil {
		logger.Debug().Str(LogProtoID, protocol.String()).Str(LogMsgID, msgID).Str(LogOrgReqID, reqID).Str("from_peer", peer.Name()).Object("body", additional).
			Msg("Received a response message")
	} else {
		logger.Debug().Str(LogProtoID, protocol.String()).Str(LogMsgID, msgID).Str(LogOrgReqID, reqID).Str("from_peer", peer.Name()).
			Msg("Received a response message")
	}
}

type LogStringersMarshaller struct {
	arr   []fmt.Stringer
	limit int
}

func NewLogStringersMarshaller(arr []fmt.Stringer, limit int) *LogStringersMarshaller {
	return &LogStringersMarshaller{arr: arr, limit: limit}
}

func (m *LogStringersMarshaller) MarshalZerologArray(a *zerolog.Array) {
	size := len(m.arr)
	if size > m.limit {
		for i := 0; i < m.limit-1; i++ {
			a.Str(m.arr[i].String())
		}
		a.Str(fmt.Sprintf("(and %d more)", size-m.limit+1))
	} else {
		for _, element := range m.arr {
			a.Str(element.String())
		}
	}
}

type LogPeerMetasMarshaller struct {
	metas []p2pcommon.PeerMeta
	limit int
}

func NewLogPeerMetasMarshaller(metas []p2pcommon.PeerMeta, limit int) *LogPeerMetasMarshaller {
	return &LogPeerMetasMarshaller{metas: metas, limit: limit}
}

func (m *LogPeerMetasMarshaller) MarshalZerologArray(a *zerolog.Array) {
	size := len(m.metas)
	if size > m.limit {
		for i := 0; i < m.limit-1; i++ {
			a.Str(ShortMetaForm(m.metas[i]))
		}
		a.Str(fmt.Sprintf("(and %d more)", size-m.limit+1))
	} else {
		for _, meta := range m.metas {
			a.Str(ShortMetaForm(meta))
		}
	}
}


type LogPeersMarshaller struct {
	metas []p2pcommon.RemotePeer
	limit int
}

func NewLogPeersMarshaller(metas []p2pcommon.RemotePeer, limit int) *LogPeersMarshaller {
	return &LogPeersMarshaller{metas: metas, limit: limit}
}

func (m *LogPeersMarshaller) MarshalZerologArray(a *zerolog.Array) {
	size := len(m.metas)
	if size > m.limit {
		for i := 0; i < m.limit-1; i++ {
			a.Str(m.metas[i].Name())
		}
		a.Str(fmt.Sprintf("(and %d more)", size-m.limit+1))
	} else {
		for _, meta := range m.metas {
			a.Str(meta.Name())
		}
	}
}


type LogStringsMarshaller struct {
	strs []string
	limit int
}

func NewLogStringsMarshaller(strs []string, limit int) *LogStringsMarshaller {
	return &LogStringsMarshaller{strs: strs, limit: limit}
}

func (m *LogStringsMarshaller) MarshalZerologArray(a *zerolog.Array) {
	size := len(m.strs)
	if size > m.limit {
		for i := 0; i < m.limit-1; i++ {
			a.Str(m.strs[i])
		}
		a.Str(fmt.Sprintf("(and %d more)", size-m.limit+1))
	} else {
		for _, meta := range m.strs {
			a.Str(meta)
		}
	}
}


type LogPeerIdsMarshaller struct {
	arr   []types.PeerID
	limit int
}

func NewLogPeerIdsMarshaller(arr []types.PeerID, limit int) *LogPeerIdsMarshaller {
	return &LogPeerIdsMarshaller{arr: arr, limit: limit}
}

func (m LogPeerIdsMarshaller) MarshalZerologArray(a *zerolog.Array) {
	size := len(m.arr)
	if size > m.limit {
		for i := 0; i < m.limit-1; i++ {
			a.Str(ShortForm(m.arr[i]))
		}
		a.Str(fmt.Sprintf("(and %d more)", size-m.limit+1))
	} else {
		for _, element := range m.arr {
			a.Str(ShortForm(element))
		}
	}
}


type LogIPNetMarshaller struct {
	arr   []*net.IPNet
	limit int
}

func NewLogIPNetMarshaller(arr []*net.IPNet, limit int) *LogIPNetMarshaller {
	return &LogIPNetMarshaller{arr: arr, limit: limit}
}

func (m LogIPNetMarshaller) MarshalZerologArray(a *zerolog.Array) {
	size := len(m.arr)
	if size > m.limit {
		for i := 0; i < m.limit-1; i++ {
			a.Str(m.arr[i].String())
		}
		a.Str(fmt.Sprintf("(and %d more)", size-m.limit+1))
	} else {
		for _, element := range m.arr {
			a.Str(element.String())
		}
	}
}