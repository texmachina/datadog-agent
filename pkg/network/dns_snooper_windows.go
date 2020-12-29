// +build windows

package network

import (
	"time"

	"github.com/DataDog/datadog-agent/pkg/process/util"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

var _ ReverseDNS = &dnsSnooperWindows{}

type dnsSnooperWindows struct {
	cache           *reverseDNSCache
	di *DriverInterface
	parser *dnsParser
}

func NewDnsSnooperWindows(di *DriverInterface) *dnsSnooperWindows {
	snooper := &dnsSnooperWindows{
		cache: newReverseDNSCache(10000, 10*time.Minute, time.Minute),
		di:    di,
		parser: newDNSParser(true, false), // TODO: config
	}

	go snooper.Run()
	return snooper


}


func (d *dnsSnooperWindows) Resolve(cxs []ConnectionStats) map[util.Address][]string {
	return d.cache.Get(cxs, time.Now())
}


func (d *dnsSnooperWindows) GetDNSStats() map[dnsKey]map[string]dnsStats {
	// TODO: implement
	return map[dnsKey]map[string]dnsStats{}
}


func (d *dnsSnooperWindows) GetStats() map[string]int64 {
	return map[string]int64{}
}


func (d *dnsSnooperWindows) Close() {
	// OK
}


func (d *dnsSnooperWindows) Run() {
	for range time.Tick(time.Second) {
		dns := d.di.GetDNS()


		var translation translation
		d.parser.ParseInto(dns, &translation, &dnsPacketInfo{})
		log.Infof("parsed packet: %v", translation)
	}
}