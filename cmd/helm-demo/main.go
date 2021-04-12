package main

import (
	"fmt"
	"github.com/gzlj/helm-demo/pkg/localhelm"
	"github.com/gzlj/helm-demo/pkg/localhelm/chart/loader"
	"github.com/gzlj/helm-demo/pkg/localhelm/engine"
	"github.com/gzlj/helm-demo/pkg/localhelm/module"

)


func main() {
	c, err := loader.Load("/opt/helm/repo/nginx")

	if err != nil {

	}
	subsys := map[string]interface{}{}
	ic := ImageConfigMap{
		"repository":  "demo.io/cas",
	}
	cas := map[string]interface{}{}
	cas["image"] = ic
	subsys["cas"] = cas

	ic = ImageConfigMap{
		"repository":  "demo.io/identify",
	}
	identify := map[string]interface{}{}
	identify["image"] = ic

	subsys["identify"] = identify

	overrides := map[string]interface{}{
		 "subsys": subsys,
	}

	caps := &module.Capabilities{}
	options := module.ReleaseOptions{
		Name:      "cloud",
		Namespace: "cloud",
		Revision:  1,
	}

	finalValues, err := localhelm.ToRenderValues(c, overrides, options, caps)

	//finalValues, err := chartutil.CoalesceValues(c, overrides)
	if err != nil {
		return
	}

	out, err := engine.Render(c, finalValues)

	if err != nil {
		fmt.Println("err:", err)
		return
	}
	for fileName, content := range out {
		fmt.Println("file name:", fileName)
		fmt.Println("file content:", content)
		fmt.Println()
	}
}

type ImageConfigMap map[string]string




