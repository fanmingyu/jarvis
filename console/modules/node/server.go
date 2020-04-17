package node

import(
	"log"

	"smsgate/utils/registry"
)

var Registry *registry.Registry

func InitRegistry(conf registry.EtcdConf) error {
	var err error
	Registry, err = registry.NewRegistry(conf)
	if err != nil {
		log.Printf("registry init failed. err:%v", err)
	}

	Registry.WatchNodes()
	return err
}
