package network

var _ PacketSource = &packetSource{}

type packetSource struct {
}

func (p *packetSource) VisitPackets(exit <-chan struct{}, visit func([]byte, time.Time) error) error {
}

func (p *packetSource) Stats() map[string]int64 {
	// TODO: explain why this is a no-op
}

func (p *packetSource) Close() {
	// TODO: explain why this is a no-op
}
