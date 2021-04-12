package localhelm

import (
	"github.com/gzlj/helm-demo/pkg/localhelm/chart"
	"github.com/gzlj/helm-demo/pkg/localhelm/module"
	"github.com/gzlj/helm-demo/pkg/localhelm/util"

)

func ToRenderValues(chrt *chart.Chart, chrtVals map[string]interface{}, options module.ReleaseOptions, caps *module.Capabilities) (util.Values, error) {

	top := map[string]interface{}{
		"Chart":        chrt.Metadata,
		"Capabilities": caps,
		"Release": map[string]interface{}{
			"Name":      options.Name,
			"Namespace": options.Namespace,
			"IsUpgrade": options.IsUpgrade,
			"IsInstall": options.IsInstall,
			"Revision":  options.Revision,
			"Service":   "Helm",
		},
	}

	vals, err := util.CoalesceValues(chrt, chrtVals)
	if err != nil {
		return top, err
	}

	top["Values"] = vals
	return top, nil
}
