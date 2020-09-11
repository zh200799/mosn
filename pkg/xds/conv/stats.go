package conv

import (
	"mosn.io/mosn/pkg/metrics"
	types "mosn.io/mosn/pkg/types"
)

var Stats types.XdsStats

// LDS， RDS， CDS， EDS以及SDS
// Listener,Router,Cluster, Endpoint, Secret
func InitStats() {
	m := metrics.NewXdsStats()
	Stats = types.XdsStats{
		CdsUpdateSuccess: m.Counter(metrics.CdsUpdateSuccessTotal),
		CdsUpdateReject:  m.Counter(metrics.CdsUpdateRejectTotal),
		LdsUpdateSuccess: m.Counter(metrics.LdsUpdateSuccessTotal),
		LdsUpdateReject:  m.Counter(metrics.LdsUpdateRejectTotal),
	}
}
