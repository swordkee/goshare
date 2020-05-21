package hub

import (
	"errors"
	"time"

	proto "github.com/golang/protobuf/proto"
	ctp "github.com/mineralres/protos/src/go/ctp"
)

// SyncAdapter  sync
type SyncAdapter struct {
	adapter *Adapter
	chIn    chan *Packet
	Data    interface{}
}

// NewSyncAdapter create new sync adapter
func NewSyncAdapter(host string, timeout time.Duration, fronts []string, h func(*Packet)) (*SyncAdapter, error) {
	var err error
	ret := &SyncAdapter{}
	ret.adapter, err = NewAdapter(host, timeout, func(pkt *Packet) {
		switch ctp.CtpMessageType(pkt.MsgType) {
		case ctp.CtpMessageType_HEARTBEAT:
			return
		case ctp.CtpMessageType_TD_OnRtnOrder:
			// send order 可能要求同步返回
			var rtn ctp.CThostFtdcOrderField
			err := pkt.Get1(&rtn)
			if err == nil {
				pkt.RequestID = rtn.RequestID
			}
			h(pkt)
		case ctp.CtpMessageType_TD_OnRtnInstrumentStatus, ctp.CtpMessageType_TD_OnRtnTrade:
			h(pkt)
			return
		}
		ret.chIn <- pkt
	})
	if err != nil {
		return nil, err
	}
	ret.chIn = make(chan *Packet, 1000)
	var req ctp.CThostFtdcReqRegisterFrontField
	req.Fronts = fronts
	ret.adapter.Post(int32(ctp.CtpMessageType_TD_RegisterFront), &req, 0)
	ret.adapter.Post(int32(ctp.CtpMessageType_TD_Init), &req, 0)
	for {
		select {
		case <-time.After(timeout):
			ret.adapter.Close()
			return ret, errors.New("CtpTimeout")
		case pkt := <-ret.chIn:
			if pkt.MsgType == int32(ctp.CtpMessageType_TD_OnFrontConnected) {
				return ret, nil
			}
		}
	}
	return ret, nil
}

// Send send msg
func (sa *SyncAdapter) Send(msgType int32, req proto.Message, requestID int32, timeout time.Duration) ([]*Packet, error) {
	sa.adapter.Post(msgType, req, requestID)
	var ret []*Packet
	for {
		select {
		case <-time.After(timeout):
			sa.adapter.Close()
			return ret, errors.New("wait timeout")
		case pkt := <-sa.chIn:
			if pkt.RequestID == requestID {
				ret = append(ret, pkt)
				if pkt.IsLast > 0 {
					return ret, nil
				}
			}
		}
	}
}
