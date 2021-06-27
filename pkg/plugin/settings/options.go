package settings

import (
	"github.com/evalsocket/policyreport-octant-plugin/pkg/plugin/controller"

	"github.com/vmware-tanzu/octant/pkg/plugin/service"
)

func GetOptions() []service.PluginOption {
	return []service.PluginOption{
		service.WithTabPrinter(controller.ResourceTabPrinter),
		//service.WithNavigation(
		//	func(_ *service.NavigationRequest) (nav navigation.Navigation, err error) {
		//		nav = navigation.Navigation{
		//			Title:    strings.Title(name),
		//			Path:     name,
		//			IconName: rootNavIcon,
		//		}
		//		return
		//	},
		//	controller.InitRoutes,
		//),
	}
}