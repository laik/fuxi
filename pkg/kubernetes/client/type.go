package client

import "k8s.io/apimachinery/pkg/runtime/schema"

type ResourceName string

// nuwa resource
const (
	Water       ResourceName = "water"
	Deployment  ResourceName = "deployment"
	Stone       ResourceName = "stone"
	StatefulSet ResourceName = "statefulset"
	DaemonSet   ResourceName = "daemonsets"
	Injector    ResourceName = "injector"
	Pod         ResourceName = "pod"
	CronJobs    ResourceName = "cronjobs"
)

// GroupVersionResources describe resource collection
var GroupVersionResources = map[ResourceName]schema.GroupVersionResource{
	Water:       {Group: "nuwa.nip.io", Version: "v1", Resource: "waters"},
	Deployment:  {Group: "apps", Version: "v1", Resource: "deployments"},
	Stone:       {Group: "nuwa.nip.io", Version: "v1", Resource: "stones"},
	StatefulSet: {Group: "apps", Version: "v1", Resource: "statefulsets"},
	DaemonSet:   {Group: "apps", Version: "v1", Resource: "daemonsets"},
	Injector:    {Group: "nuwa.nip.io", Version: "v1", Resource: "injectors"},
	Pod:         {Group: "", Version: "v1", Resource: "pods"},
	CronJobs:    {Group: "batch", Version: "v1beta1", Resource: "cronjobs"},
}

var (
	ResourceWater       = GetGVR(Water)
	ResourceDeployment  = GetGVR(Deployment)
	ResourceStone       = GetGVR(Stone)
	ResourceStatefulSet = GetGVR(StatefulSet)
	ResourceDaemonSet   = GetGVR(DaemonSet)
	ResourceInjector    = GetGVR(Injector)
	ResourcePod         = GetGVR(Pod)
	ResourceCronJobs    = GetGVR(CronJobs)
)

func GetGVR(rs ResourceName) schema.GroupVersionResource {
	gvr, exist := GroupVersionResources[rs]
	if !exist {
		panic("try to get an undefined resource")
	}
	return gvr
}
