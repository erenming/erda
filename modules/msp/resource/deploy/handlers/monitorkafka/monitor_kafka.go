// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package monitorkafka

import (
	"strings"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/msp/instance/db"
	"github.com/erda-project/erda/modules/msp/resource/deploy/handlers"
	"github.com/erda-project/erda/modules/msp/resource/utils"
	"github.com/erda-project/erda/modules/orchestrator/services/addon"
	"github.com/erda-project/erda/pkg/parser/diceyml"
)

func (p *provider) IsMatch(tmc *db.Tmc) bool {
	return tmc.Engine == handlers.ResourceMonitorKafka
}

func (p *provider) BuildServiceGroupRequest(resourceInfo *handlers.ResourceInfo, tmcInstance *db.Instance, clusterConfig map[string]string) interface{} {
	req := p.DefaultDeployHandler.BuildServiceGroupRequest(resourceInfo, tmcInstance, clusterConfig).(*apistructs.ServiceGroupCreateV2Request)

	req.GroupLabels["ADDON_GROUPS"] = "3"

	instanceOptions := map[string]string{}
	utils.JsonConvertObjToType(tmcInstance.Options, &instanceOptions)

	for _, service := range resourceInfo.Dice.Services {
		nodeId := utils.GetRandomId()

		// envs
		env := map[string]string{
			"ADDON_ID":                tmcInstance.ID,
			"ADDON_NODE_ID":           nodeId,
			"KAFKA_LOG_DIRS":          "/kafka/data",
			"KAFKA_ZOOKEEPER_CONNECT": instanceOptions["ZK_HOSTS"],
		}
		utils.AppendMap(service.Envs, env)

		// labels
		if service.Labels == nil {
			service.Labels = make(map[string]string)
		}
		options := map[string]string{}
		utils.JsonConvertObjToType(tmcInstance.Options, &options)
		utils.SetlabelsFromOptions(options, service.Labels)

		// volumes
		if p.IsNotDCOSCluster(clusterConfig["DICE_CLUSTER_TYPE"]) {
			/*
				service.Binds = diceyml.Binds{
					nodeId + "_data:/kafka/data:rw",
				}
			*/
			//  /kafka/data volume
			vol := addon.SetAddonVolumes(options, "/kafka/data", false)
			service.Volumes = diceyml.Volumes{vol}
		} else {
			service.Binds = diceyml.Binds{
				clusterConfig["DICE_STORAGE_MOUNTPOINT"] + "/addon/kafka/" + nodeId + "/kafka/data:/kafka/data:rw",
			}
		}
	}

	return req
}

func (p *provider) BuildTmcInstanceConfig(tmcInstance *db.Instance, serviceGroupDeployResult interface{},
	clusterConfig map[string]string, additionalConfig map[string]string) map[string]string {
	instanceOptions := map[string]string{}
	utils.JsonConvertObjToType(tmcInstance.Options, &instanceOptions)

	var hosts []string
	serviceGroup := serviceGroupDeployResult.(*apistructs.ServiceGroup)
	for _, service := range serviceGroup.Services {
		hosts = append(hosts, service.Vip+":9092")
	}

	config := map[string]string{
		"KAFKA_HOSTS": strings.Join(hosts, ","),
		"ZK_HOSTS":    instanceOptions["ZK_HOSTS"],
	}

	return config
}
