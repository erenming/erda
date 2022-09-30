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

package clusters

import (
	"context"
	"encoding/base64"
	"net/http"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"

	clusterpb "github.com/erda-project/erda-proto-go/core/clustermanager/cluster/pb"
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/bundle"
	"github.com/erda-project/erda/internal/apps/cmp/dbclient"
	"github.com/erda-project/erda/pkg/mock"
)

const (
	fakeKubeconfigWithCert  = "YXBpVmVyc2lvbjogdjEKY2x1c3RlcnM6Ci0gY2x1c3RlcjoKICAgIHNlcnZlcjogaHR0cHM6Ly8xMjcuMC4wLjE6NjQzNTEKICBuYW1lOiBraW5kLW1vb24KY29udGV4dHM6Ci0gY29udGV4dDoKICAgIGNsdXN0ZXI6IGtpbmQtbW9vbgogICAgdXNlcjoga2luZC1tb29uCiAgbmFtZToga2luZC1tb29uCmtpbmQ6IENvbmZpZwp1c2VyczoKLSBuYW1lOiBraW5kLW1vb24KICB1c2VyOgogICAgY2xpZW50LWNlcnRpZmljYXRlLWRhdGE6IExTMHRMUzFDUlVkSlRpQkRSVkpVU1VaSlEwRlVSUzB0TFMwdENrMUpTVU00YWtORFFXUnhaMEYzU1VKQlowbEpVR0Z3TkdGSmIyZDBPREIzUkZGWlNrdHZXa2xvZG1OT1FWRkZURUpSUVhkR1ZFVlVUVUpGUjBFeFZVVUtRWGhOUzJFelZtbGFXRXAxV2xoU2JHTjZRV1ZHZHpCNVRWUkZlRTFFUlhkTmFsVXlUVVJHWVVaM01IbE5ha1Y0VFVSRmQwMXFWVEpOUkZKaFRVUlJlQXBHZWtGV1FtZE9Wa0pCYjFSRWJrNDFZek5TYkdKVWNIUlpXRTR3V2xoS2VrMVNhM2RHZDFsRVZsRlJSRVY0UW5Ka1YwcHNZMjAxYkdSSFZucE1WMFpyQ21KWGJIVk5TVWxDU1dwQlRrSm5hM0ZvYTJsSE9YY3dRa0ZSUlVaQlFVOURRVkU0UVUxSlNVSkRaMHREUVZGRlFXMTZSa1l3TkdWUmRqbFlWRGRGY21ZS2MxZFhlalV2Y3paYWQySkJXbTVrU25GNFVVdDVVbmsyWjJKd1kyUk9RVU5ZVTFScVdVdEJTM2xaVEd4bGVWbzFUVkV4TmxJNVUyNU9aVmR0V0hCUmNRcFFja0ZyWlRCMk9WcENkV3htUTFCb2NteHViRGRIYWpJMlZqZHhlakV4U0doc1pGRlRTM0pHTWpoTVNUSkpZamQ2VG5OTVlUTlljRFl3VkN0RGNXRTVDbXhYVERGSWFqaDFPRmxFTjB4WGVuWlBORU1yWmpkTlZETTFjRzgxVTNReFFuQnNWVWRTTjNFNE4wdGhUV055VVVWbWIxbDJPVmd5ZEZoMU0xZExUa1VLWXk5SVlsVXdkVUZTVDFSV1JUSXlOakU1TldRM2NUZFJZMDFVTkhCUk1qazRWamhNVDBKek1WZGpSVlpyWnpsYWFWQkZOMUpFVTBJMFJ6VkZhbEl6YWdvdk0yTkRWRGxPYUV0emQwNVlhVVEzZUZCc1MwUnZaM2hEYVdOc2MweHdjMUZzTTNSeFlWRndUVGRhUm1aM2RXSlBPQ3N2ZURWd2IxQklkVGRZUW5kR0NrTjVNV2RIVVVsRVFWRkJRbTk1WTNkS1ZFRlBRbWRPVmtoUk9FSkJaamhGUWtGTlEwSmhRWGRGZDFsRVZsSXdiRUpCZDNkRFoxbEpTM2RaUWtKUlZVZ0tRWGRKZDBSUldVcExiMXBKYUhaalRrRlJSVXhDVVVGRVoyZEZRa0ZHYWpWSWNtdHROa3BXY0VWSmJuRXJLMVY2VW5OcGNra3pNU3REY25ZclNXMUxlQXBFY2tGRU9WcFFkbHBJU20xb09FczJRVk5YZFcxbWR6UjJVV05wWTFocVdUWlJMM2Q0U2taeU5qVTVMMjlITHpCclRtRjFXbWh1UzBoSFFVdG1kRTVQQ25SemRHOU1aMWxwVjNKRk0zbHZRa3BaYjBoSFkzSkxURVZEWjJ4RGJVaFBOMVpqZURBM2QzRjZUM2g0TUhSeFpVaEhkWFEyTWxKMGFqWlRMM012UWpBS2FXUXpUMUYwUVhVNVpUbEtWV2xpUldSSk5scE9VM1pTZWtFMVpuTlNWVUl3TjNwYU1ITmhTRmRsWjFGdlUzWlVLMDFoUVhRd1ZVMW5jWGhZTUc5TFN3cGxWbXhJVWt0VmNpOWhNMnBhYzJoNk0yMDFabFJYUmxSUWJsVm1ka3cyYjI1aWJuVkVWMWhEU1c1ak5HUkZUR3RVU0RSaUsxZE5OV2RaU0ZGS0t6SldDbWxKVFhOcGEyWm5XVEZrY1cxVmEzaDZURzFKYWs0MFlVWlJNM056Tld4TlUzbFhWbEppYW5CMlFtSnNUVWR5ZEdVcmF6MEtMUzB0TFMxRlRrUWdRMFZTVkVsR1NVTkJWRVV0TFMwdExRbz0KICAgIGNsaWVudC1rZXktZGF0YTogTFMwdExTMUNSVWRKVGlCU1UwRWdVRkpKVmtGVVJTQkxSVmt0TFMwdExRcE5TVWxGYjJkSlFrRkJTME5CVVVWQmJYcEdSakEwWlZGMk9WaFVOMFZ5Wm5OWFYzbzFMM00yV25kaVFWcHVaRXB4ZUZGTGVWSjVObWRpY0dOa1RrRkRDbGhUVkdwWlMwRkxlVmxNYkdWNVdqVk5VVEUyVWpsVGJrNWxWMjFZY0ZGeFVISkJhMlV3ZGpsYVFuVnNaa05RYUhKc2JtdzNSMm95TmxZM2NYb3hNVWdLYUd4a1VWTkxja1l5T0V4Sk1rbGlOM3BPYzB4aE0xaHdOakJVSzBOeFlUbHNWMHd4U0dvNGRUaFpSRGRNVjNwMlR6UkRLMlkzVFZRek5YQnZOVk4wTVFwQ2NHeFZSMUkzY1RnM1MyRk5ZM0pSUldadldYWTVXREowV0hVelYwdE9SV012U0dKVk1IVkJVazlVVmtVeU1qWXhPVFZrTjNFM1VXTk5WRFJ3VVRJNUNqaFdPRXhQUW5NeFYyTkZWbXRuT1ZwcFVFVTNVa1JUUWpSSE5VVnFVak5xTHpOalExUTVUbWhMYzNkT1dHbEVOM2hRYkV0RWIyZDRRMmxqYkhOTWNITUtVV3d6ZEhGaFVYQk5OMXBHWm5kMVlrODRLeTk0TlhCdlVFaDFOMWhDZDBaRGVURm5SMUZKUkVGUlFVSkJiMGxDUVVkamVVWlpaRFpRZFU5Q1ZWbEdkZ3B5UmpKMlJHcHNNVVZxYms5QlNWSnBWVUZrUm5SWlNsUTJlVlV5WjJsS1pUWjVUall3ZDA4cmFVVXpSa1E0YldKS1RERmpaeTlLWkRCeE5rSlFaV0p2Q2xOVmMxVklOelJ3U1hBMk9YbzBXRVEyYkRkbloyOWFOSEZoT0ROSldtRjNTbFUyVFc5MVRXWlJOa0pHV0hZMk1HazFNa2RFYjNKMmFIQlBaV1JtYzNBS1RYbHRlQ3QxUzFwRGVVMXRURVl2YUdKNFl6TklabGRYV1N0SGMwb3pkMHQyTlhVMFJIaE5hVUZDWld0S2JTODNNVlpRTURsMVZrdGphRFJWU0RoNVpncFJVMjVEU1dKd2RtTXlLMmxsWlhvclozVTJRa1J6UTJoRFltOUpZVUZPUTNRMVFrRklXVmwxY0hVMk1sRkZkVmMwWVU5QlpreEtZV1E0VDBwTk1WUndDaTh3T0dvMGRtVkZPWHBPUzNKNmN6VjRaVmhDTUhwQ1lqRk5SVlpHVFVrek9YbHFOMDFMYzBReVNrNW9iaXRFYm1ocFFqSlFkemhqUVhVNE9GVXJieklLVEZST1kxTklNRU5uV1VWQmVVVnZiRVJQV2pGTmJqRTFWa2REYWtSSGQxQTFlVGt2U2twemFtZDNjMFUzZW14bFkzZEdOMUZFTlRjMVlWVTFiR1pTUndwb1lXdGlSamh2ZEdoWFRIVkZPVzhyTUZac1JTOUVTbXRJYUVsd0t6QllTVzRyTUVaa1IydG1OMDFNUVZGaVlYazNhMWw0Tm10d2VWUnFRVEZxTms5SENqTm1RbTVEVFRGM1VUbGlSblpyYUM5VFltRm9NVTFCVEdFeU5VTnFPVzByYWxoalVUTldkR3hKU1c5d1ZWTkhOa3hDYWtkclF6aERaMWxGUVhoc2RuVUtlSEpDVURsb2FHcHRRVFpTWjBNeWFXWlFRWEZsUkZFMGJVNU1PV0l5YVdGMVNHMWhNV2MyVFZCTU1HTndTRzVFVm01TlFtTlBOelZHVDFoTlpuVjFiZ3BWWmtneU1XeERjVmhqWjNOa1dVdDVSVkZCUjJ4b2JISklOVUpEYVV4aVdGUnJaVUpZVlhWWmJtb3pVMG96ZEU5TlptZ3dTbWxOVkZGWWJXMXZVRUp6Q25GUFJ6RXdNVmhFZG05WFlVbGxWbFl5VUcwMVpsSjRhVGh2U3pJemVuTm5jM0pQVTJWcVkwTm5XVUkwZGxKSFYzUkVjVVZLUmxkSGQzWmhibGd5UlRZS01sVjNTVGhwYVROTU1scGlSemRrVUVaME5rbDJUMGxPWW5OMlNFWk5iVXRsTkZkcmEzSkdZMVJ0V0RoMU4zcFJhMDFNWjFWVVVTOVhSek5OVm5GWU5RcHRibGRIVm01VWVsVm1aRFUxWW5ZMWJsbEZjbTV3VGtob1VrcFZjelY0ZVc5RFFrMXJlVTlTUldGRFJUZzRhMll3Tms4NGVqaFNVMm81V1RaU01saDVDbVUwZEVWR1dIUm9ORlZtUlRZNVRVTkZSRVp4Tm5kTFFtZEhjRzl3VDFnellsWXpVMXBhVFV4MFlXZDVWWFZzT1VkUWRGVTBSWFY2Vnk5MGJ5OVNPR0lLVGxOSVFuVm9PVmhRT0dSSUwzRTRaMnRPVWxwcWEzUlZiakpuZVRKS1lWRkhPU3RMWkVkd1RUSlFWRTVCVDAwemVGWkRLMGxCV2xaeFEzVmFSV3hXVHdvM1MyNHdlRTVIVjNwdFkxcHpZa0ZETTJGUllubDBaV292SzNnd1RGbGFTRloyWjJoSFVtNWhjQzh5Ynk5dE1tZFJXVFZLVEZKa01WbGpOMDA1UmxOS0NrZHViblpCYjBkQlYzbFVWbGhFVlhGUVdXaFNVRlZ1VTBNdmRsUTRNV05MVW5GVmJHdEllRVYyTkVJeWVWcEZLM05oVVdJNFVEQlZRemRDTWt4TFUzUUtabmt6ZG5oVWRVMHhlRVpsV2t3emVHUllXRVZzWlVWTFVXSlBVM1JXY2tJMlVXVjRkRTVyY0dsM1MwTlVTVlJOUTJFM0wwazVlRzhyVVhBMFVqZHJiUXBJYmxwdlRFNVdkbEZtUkZoTFF6VnNNakpKYjNjclVIcHRaMmRtV1ZNNE1qUkhaVmRJUldKdWFtRkRiRE5pT0ZSb2FuYzlDaTB0TFMwdFJVNUVJRkpUUVNCUVVrbFdRVlJGSUV0RldTMHRMUzB0Q2c9PQo="
	fakeKubeconfigWithToken = "YXBpVmVyc2lvbjogdjEKa2luZDogQ29uZmlnCmNsdXN0ZXJzOgotIG5hbWU6ICJraW5kLWtpbmQiCiAgY2x1c3RlcjoKICAgIHNlcnZlcjogImh0dHBzOi8vbG9jYWxob3N0OjY0NDMvazhzL2NsdXN0ZXJzL2MtbS1mNGxzaGtoOSIKdXNlcnM6Ci0gbmFtZTogImtpbmQta2luZCIKICB1c2VyOgogICAgdG9rZW46ICJrdWJlY29uZmlnLXVzZXItc2ZoaHN4anNmYTp6cTk3OTJ3cXY3Y3B4cHR2Z3ZucGd6bHFzZnZjcTh4MnJjNzZucHI0ZDVnZnZmNHJraHBmNzYiCmNvbnRleHRzOgotIG5hbWU6ICJraW5kLWtpbmQiCiAgY29udGV4dDoKICAgIHVzZXI6ICJraW5kLWtpbmQiCiAgICBjbHVzdGVyOiAia2luZC1raW5kIgpjdXJyZW50LWNvbnRleHQ6ICJraW5kLWtpbmQi"
)

func Test_ImportCluster(t *testing.T) {
	var bdl *bundle.Bundle
	var db *dbclient.DBClient
	c := New(db, bdl, nil, &fakeClusterServiceServer{}, nil, &mock.MockPipelineServiceServer{})

	monkey.PatchInstanceMethod(reflect.TypeOf(bdl), "CreateClusterWithOrg", func(bundle *bundle.Bundle, userID string, orgID uint64, req *apistructs.ClusterCreateRequest, header ...http.Header) error {
		return nil
	})

	defer monkey.UnpatchAll()

	err := c.importCluster(context.Background(), "1", &apistructs.ImportCluster{
		ClusterName:    "",
		Credential:     apistructs.ICCredential{},
		CredentialType: ProxyType,
		OrgID:          1,
		ClusterType:    "k8s",
		WildcardDomain: "fake-domain",
		DisplayName:    "unit-test",
	})

	assert.NoError(t, err)
}

func Test_ParseManageConfigFromCredential(t *testing.T) {
	var datas = []struct {
		credentialType string
		credential     apistructs.ICCredential
		manageConfig   apistructs.ManageConfig
	}{
		{
			credentialType: KubeconfigType,
			credential: apistructs.ICCredential{
				// fake kubeconfig
				Content: fakeKubeconfigWithCert,
			},
			manageConfig: apistructs.ManageConfig{
				Type:             apistructs.ManageCert,
				CredentialSource: KubeconfigType,
			},
		},
		{
			credentialType: SAType,
			credential: apistructs.ICCredential{
				Address: "1.1.1.1",
				// fake secret with default auth
				Content: "YXBpVmVyc2lvbjogdjEKZGF0YToKICBjYS5jcnQ6IExTMHRMUzFDUlVkSlRpQkRSVkpVU1VaSlEw\nRlVSUzB0TFMwdENrMUpTVVJQZWtORFFXbFBaMEYzU1VKQlowbENRVVJCVGtKbmEzRm9hMmxIT1hj\nd1FrRlJjMFpCUkVFclRWTmpkMFIzV1VSV1VWRkxSWGRvYjFsWE5XNEtaVzFvZG1SVVFWVkNaMDVX\nUWtGdlZFUlhSbk5oVjBwb1dXMUZaMWt5ZUhaa1YxRjRSWHBCVWtKblRsWkNRVTFVUTIxME1WbHRW\nbmxpYlZZd1dsaE5kd3BKUW1OT1RXcEZkMDU2U1RWTlJHZDRUbnBCTUZkb1oxQk5ha0V4VFZSQk0w\nMXFTWGRQUkVVelRVUlNZVTFFTkhoS2VrRlFRbWRPVmtKQmIxUkRSMmhvQ21KdFpEWmhSemt4VFVK\nUlIwRXhWVVZEYUUxT1dWZDRjRmx0Um1sWlUwSnFZa2M1TVZwRVJWUk5Ra1ZIUVRGVlJVRjRUVXRo\nTTFacFdsaEtkVnBZVW13S1kzcERRMEZUU1hkRVVWbEtTMjlhU1doMlkwNUJVVVZDUWxGQlJHZG5S\nVkJCUkVORFFWRnZRMmRuUlVKQlRWTkplakpOZEdZMWF5OVZhR2xEZFc5UWJRcFdVa1ZJZGtWWVMz\naHhiV0YwWkN0c1dHbHhSMk5NUnpSRWVrVnhXVWN2Wm5WUFFUVXhVM0U1YVRGSVNWWnBOM0IzV1U1\nTlpFUlZjRXBLYWxWTk5HWmpDa2hzWkd4dlVYSXZNMnN4WVUxMk9YWllWbXB3Vm5sNU9HUnRlREpa\nZVdKT1ZYVTFMMGhzU0RZeVpVOXBSbG8yV2xCa1RTODFVa2hpYW5kTFJWazVjRElLUmt0aVMyMWhV\nSEZpZVRGQ2QyOXliWFkxVTBSdGNEazNXRTlIWTFwaVRrc3dlRWN4ZUhsblVUQlNSQ3RJU1haaVMx\nSnFORnBFUVd0UmNHeGpSVVZNVndwU1pFRXJhbmxzU0dRcmJ6a3JUM3BzWlU1NmEzUlJlWGhSYWs5\nSk4wMXZabUZOYUhGblNVNHdiVlJMTUdKSWMxUm9kWEoxTmpRd1VERlllVEp6U2pCTENtaFJZM0Z0\nYkVaWVIyeHVOemsxUnpGTlFuZERWemRaUmtOR2JXUnZWVmcyWm1keGFqaDJOMjAwZHpGeFZWQkRl\nV1kxYURodVQyTnVla3Q0Y2poMWVXTUtNVTlGUTBGM1JVRkJZVTVEVFVWQmQwUm5XVVJXVWpCUVFW\nRklMMEpCVVVSQlowdHJUVUU0UjBFeFZXUkZkMFZDTDNkUlJrMUJUVUpCWmpoM1NGRlpSQXBXVWpC\nUFFrSlpSVVpITnpsU00xSnlkMHRxVVZsMmRsVjBjVmhUUWl0dFV6UmhhbWxOUVRCSFExTnhSMU5K\nWWpORVVVVkNRM2RWUVVFMFNVSkJVVU5uQ2pGNk1sb3lhRkU1VVZGbUwxSkZhRFEyY1RCcE9HRldj\nM0EwUzIxc1IxbFdNMVpIVkdveWExaEhTU3R1Unl0RGExRndZMlZSZERGNmVrbHJOWEJGTjBNS1JX\nZzRRbkprYTA1aVZFMUdVVVo1YmxaVmVrcExTR2xKSzI1VVVFUnBRMEpZZDNSck9XZHhXRU5OV1Vk\nVWRGVkplVXhGWlN0cE5FRkpRbEpyUmxSTVFRcG9OMFpyVlhCcFJETlNaR050TDFWcFFYWkpRMDFO\nYm5WaWRXUkdWRUZETWs0ck9UaDNWRW92WTBOSmIyRXdVVlpOZVdOb1pVUmFLMGhqUms4elkyTmhD\nbEl2VlVoTlNtMTNjbU5XYVhGaE5tZHRTbEpRYUU1YVpWQmlSR3RPYXpsTFRWazVWMFJ5UVU5M0sw\nbHRhMmxEWm5WUGNVMVVPRmhLYzJvMlpYTlNXRGNLYmxSRlducEpWV0ZVVm5NcmJEVnBabEZSV0M5\neWVWRlBSSFEzZGtKR04wbHpTWEI2ZWpGc01rZFphRXBHY0dFeVRrZG5ia2RDUWtwWU1FZG9NSFZN\nYVFvclNGVkxTMWMxVVVwT1ZXRlFXa1ZUV1dzeE5nb3RMUzB0TFVWT1JDQkRSVkpVU1VaSlEwRlVS\nUzB0TFMwdENnPT0KICBuYW1lc3BhY2U6IFpHVm1ZWFZzZEE9PQogIHRva2VuOiBaWGxLYUdKSFky\nbFBhVXBUVlhwSk1VNXBTWE5KYlhSd1drTkpOa2xyV25aT01qRm9UMVZaZUdKSGVEQmlSMlIzVFRC\nc1JFOVlhRnBOYXpsdVUwVmFkVm95V2sxTlZGWjBUMWRrZDA5SGNGaFhSbFpHVFRKemFXWlJMbVY1\nU25Cak0wMXBUMmxLY21SWFNteGpiVFZzWkVkV2Vrd3pUbXhqYmxwd1dUSldhRmt5VG5aa1Z6VXdT\nV2wzYVdFelZtbGFXRXAxV2xoU2JHTjVOWEJpZVRsNldsaEtNbUZYVG14WlYwNXFZak5XZFdSRE9Y\nVlpWekZzWXpOQ2FGa3lWV2xQYVVwcldsZGFhR1JYZURCSmFYZHBZVE5XYVZwWVNuVmFXRkpzWTNr\nMWNHSjVPWHBhV0VveVlWZE9iRmxYVG1waU0xWjFaRU01ZWxwWFRubGFXRkYxWW0xR2RGcFRTVFpK\nYlZKc1dtMUdNV0pJVVhSa1J6bHlXbGMwZEZscVZucGhlbU5wVEVOS2NtUlhTbXhqYlRWc1pFZFdl\na3h0YkhaTU0wNXNZMjVhY0ZreVZtaFpNazUyWkZjMU1Fd3pUbXhqYmxwd1dUSlZkRmxYVG1waU0x\nWjFaRU0xZFZsWE1XeEphbTlwV2tkV2JWbFlWbk5rUTBselNXMTBNVmx0Vm5saWJWWXdXbGhOZFdG\nWE9IWmpNbFo1Wkcxc2FscFhSbXBaTWpreFltNVJkbU15Vm5sa2JXeHFXbE14YUZreVRuWmtWelV3\nVEc1V2NGcERTVFpKYlZWNFdWUkJlazlVVm1oTVZFNXJUVVJWZEU1RVdUQlBRekZvV2xkWk0weFVW\nVFZQVkVrelRXMUZlVmw2V1RKUFEwbHpTVzVPTVZscFNUWkpiazQxWXpOU2JHSlVjSHBhV0VveVlW\nZE9iRmxYVG1waU0xWjFaRVJ3YTFwWFdtaGtWM2d3VDIxU2JGcHRSakZpU0ZGcFpsRXVTVEZwWDA4\nMlQySmlOa0ZtY3pWVmFYWTNTWEJWVW5KbWNUVXhSMDl0ZEVaVGFrRk5VMnBLVFdkRmVsazJjMlkx\nWnpodFFrbFVVMXBwTjI5UE5tSmpVRTR6ZEVWbllYbDFWV2x2YjA0emVtbHFXR1ZyTjBad04xOUJl\nV1F3WVhWS1VERkpaSGxzVGtSeGRteGxUbXBCZUhvNFRsWlZZMlp2TkZWZlFpMUpkazVKT0dSNmIw\nMW5XWGxGY1dOeGNWQnBabVZ6VkZaWVl6TkVjVjkxVVdWZldubExkVko2UVVwT2RGUkJVSGxqU2tW\nMWRscHBhME5mZVZoSWNWSjFaM1JNTUdkbVdUVkhSVE4zYUROYU5tOWZWRjlwVjFoV05raDJWVEF3\nTlVkamNqUnpSMmRtWTJ0ZlJVdzJPVmhCVmtoV05WZFhUbmxqV0RORE9WaHhVWFppTkVkTGRGbFdS\nV1JPWlZSS1YzZENNSFJmWlhrdFJuSTJMVkJ1UjJneFluUnJVRk5HTUdGcWFsVnVVamRZZFVObFJI\naDRZa2MyT1hWcmRFUkZPR05hTkRBNE5tRmhkR2sxVmxCQlkwWlVjVXcwVkhCbmRqQm4Ka2luZDog\nU2VjcmV0Cm1ldGFkYXRhOgogIGFubm90YXRpb25zOgogICAga3ViZXJuZXRlcy5pby9zZXJ2aWNl\nLWFjY291bnQubmFtZTogZGVmYXVsdAogICAga3ViZXJuZXRlcy5pby9zZXJ2aWNlLWFjY291bnQu\ndWlkOiBlMWEwMzk1YS0zZDA1LTQ2NDgtYWVmNy01OTkyNzJhMmM2NjgKICBjcmVhdGlvblRpbWVz\ndGFtcDogIjIwMjEtMDctMjlUMDg6MjM6MzhaIgogIG1hbmFnZWRGaWVsZHM6CiAgLSBhcGlWZXJz\naW9uOiB2MQogICAgZmllbGRzVHlwZTogRmllbGRzVjEKICAgIGZpZWxkc1YxOgogICAgICBmOmRh\ndGE6CiAgICAgICAgLjoge30KICAgICAgICBmOmNhLmNydDoge30KICAgICAgICBmOm5hbWVzcGFj\nZToge30KICAgICAgICBmOnRva2VuOiB7fQogICAgICBmOm1ldGFkYXRhOgogICAgICAgIGY6YW5u\nb3RhdGlvbnM6CiAgICAgICAgICAuOiB7fQogICAgICAgICAgZjprdWJlcm5ldGVzLmlvL3NlcnZp\nY2UtYWNjb3VudC5uYW1lOiB7fQogICAgICAgICAgZjprdWJlcm5ldGVzLmlvL3NlcnZpY2UtYWNj\nb3VudC51aWQ6IHt9CiAgICAgIGY6dHlwZToge30KICAgIG1hbmFnZXI6IGt1YmUtY29udHJvbGxl\nci1tYW5hZ2VyCiAgICBvcGVyYXRpb246IFVwZGF0ZQogICAgdGltZTogIjIwMjEtMDctMjlUMDg6\nMjM6MzhaIgogIG5hbWU6IGRlZmF1bHQtdG9rZW4tYjVzazcKICBuYW1lc3BhY2U6IGRlZmF1bHQK\nICByZXNvdXJjZVZlcnNpb246ICIzMjEiCiAgc2VsZkxpbms6IC9hcGkvdjEvbmFtZXNwYWNlcy9k\nZWZhdWx0L3NlY3JldHMvZGVmYXVsdC10b2tlbi1iNXNrNwogIHVpZDogNDc3MmQ5NmUtMWQyNS00\nNjYyLTgzNzAtODI1NTM3NjE5OTg0CnR5cGU6IGt1YmVybmV0ZXMuaW8vc2VydmljZS1hY2NvdW50\nLXRva2VuCg==",
			},
			manageConfig: apistructs.ManageConfig{
				Type:             apistructs.ManageToken,
				CredentialSource: SAType,
			},
		},
		{
			credentialType: ProxyType,
			credential:     apistructs.ICCredential{},
			manageConfig: apistructs.ManageConfig{
				Type:             apistructs.ManageProxy,
				CredentialSource: ProxyType,
			},
		},
		{
			credentialType: "fake-type",
		},
	}

	for _, data := range datas {
		mc, err := ParseManageConfigFromCredential(data.credentialType, data.credential)
		switch data.credentialType {
		case ProxyType:
			assert.Equal(t, mc.Type, apistructs.ManageProxy)
			assert.Equal(t, mc.CredentialSource, ProxyType)
			assert.NotEqual(t, mc.AccessKey, "")
		case SAType:
			assert.Equal(t, mc.Type, apistructs.ManageToken)
			assert.Equal(t, mc.Address, "1.1.1.1")
			assert.NotEqual(t, mc.Token, "")
			assert.NoError(t, err, nil)
		case KubeconfigType:
			assert.Equal(t, mc.Type, apistructs.ManageCert)
			assert.Equal(t, mc.Address, "https://127.0.0.1:64351")
			assert.Equal(t, mc.Token, "")
			assert.Equal(t, mc.CaData, "")
			assert.NotEqual(t, mc.CertData, "")
			assert.NotEqual(t, mc.KeyData, "")
			assert.NoError(t, err, nil)
		default:
			assert.Error(t, err)
		}
	}
}

func Test_Generate(t *testing.T) {
	_ = generateSetValues(&clusterpb.ClusterInfo{
		Type:           "k8s",
		WildcardDomain: "fake-domain",
		Name:           "fake-cluster",
	}, "fake-key")

	ak := generateAccessKey(10)
	assert.Equal(t, len(ak), 10)
	assert.Equal(t, renderReleaseImageAddr(), "/cluster-ops:v")
	assert.Equal(t, generateInitJobName(1, "fake-cluster"), "erda-cluster-init-1-fake-cluster")
	assert.Equal(t, getWorkerNamespace(), "")
}

func Test_ParseKubeConfig_Token(t *testing.T) {
	res, err := base64.StdEncoding.DecodeString(fakeKubeconfigWithToken)
	assert.NoError(t, err)

	mc, err := ParseKubeconfig(res)
	assert.NoError(t, err)
	assert.NotEqual(t, mc.Token, "")
}

func Test_convertSchedConfigToPbSchedConfig(t *testing.T) {
	type args struct {
		in *apistructs.ClusterSchedConfig
	}

	tests := []struct {
		name string
		args args
		want *clusterpb.ClusterSchedConfig
	}{
		{
			name: "case 1",
			args: args{
				in: &apistructs.ClusterSchedConfig{
					EnableTag:         true,
					EnableWorkspace:   true,
					CPUSubscribeRatio: "10",
				},
			},
			want: &clusterpb.ClusterSchedConfig{
				EnableTag:         true,
				EnableWorkspace:   true,
				CpuSubscribeRatio: "10",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertSchedConfigToPbSchedConfig(tt.args.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertSchedConfigToPbSchedConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
