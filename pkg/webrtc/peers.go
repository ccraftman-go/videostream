package webrtc

import "sync"

type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}
type Peers struct {
	ListLock    sync.RWMutex
	Connections []PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStraticRTP
}

func (p *Peers) DispatchKeyFrame() {

}
