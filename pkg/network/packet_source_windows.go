package network

import (
	"time"

	"github.com/DataDog/datadog-agent/pkg/util/log"
)

var _ PacketSource = &packetSource{}

type packetSource struct {
	di *DriverInterface
}

func NewPacketSource(di *DriverInterface) PacketSource {
	return &packetSource{di: di}
}

func (p *packetSource) VisitPackets(exit <-chan struct{}, visit func([]byte, time.Time) error) error {
	log.Infof("visited packet")
	bs := p.di.GetDNS()
	if len(bs) > 0 {
		err := visit(bs, time.Now())
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *packetSource) Stats() map[string]int64 {
	// TODO: explain why this is a no-op
	return map[string]int64{}
}

func (p *packetSource) Close() {
	// TODO: explain why this is a no-op
}
