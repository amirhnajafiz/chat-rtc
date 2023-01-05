package webrtc

import "sync"

type Room struct {
	Peers Peers
}

type Peers struct {
	ListLock    sync.RWMutex
	Connections []PeerConnectionState
	TrackLocals map[string]webrtc.TrackLocalStaticRTP
}

func (p *Peers) DispatchKeyFrame() {

}
