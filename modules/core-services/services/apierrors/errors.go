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

package apierrors

import (
	"github.com/erda-project/erda/pkg/http/httpserver/errorresp"
)

var (
	ErrCreateOrg                            = err("ErrCreateOrg", "创建企业失败")
	ErrUpdateOrg                            = err("ErrUpdateOrg", "更新企业失败")
	ErrUpdateOrgIngress                     = err("ErrUpdateOrgIngress", "更新企业入口失败")
	ErrSetReleaseCrossCluster               = err("ErrSetReleaseCrossCluster", "设置制品跨集群部署失败")
	ErrListOrg                              = err("ErrListOrg", "获取企业列表失败")
	ErrListPublicOrg                        = err("ErrListPublicOrg", "获取公开企业列表失败")
	ErrGetOrg                               = err("ErrGetOrg", "获取企业失败")
	ErrDeleteOrg                            = err("ErrDeleteOrg", "删除企业失败")
	ErrChangeOrg                            = err("ErrChangeOrg", "切换企业失败")
	ErrFetchOrgResources                    = err("ErrFetchOrgResources", "获取企业资源失败")
	ErrGetOrgClusterRelation                = err("ErrGetOrgClusterRelation", "获取企业与集群关联关系失败")
	ErrGetNotifyConfig                      = err("ErrGetNotifyConfig", "获取通知配置")
	ErrGetOrgVerifiCode                     = err("ErrGetOrgVerifiCode", "获取企业邀请验证码失败")
	ErrRelateCluster                        = err("ErrRelateCluster", "关联集群失败")
	ErrGetOrgClusterRelationsByOrg          = err("ErrGetOrgClusterRelationsByOrg", "获取企业集群关系失败")
	ErrCreateProject                        = err("ErrCreateProject", "创建项目失败")
	ErrUpdateProject                        = err("ErrUpdateProject", "更新项目失败")
	ErrDeleteProject                        = err("ErrDeleteProject", "删除项目失败")
	ErrGetProject                           = err("ErrGetProject", "获取项目失败")
	ErrGetProjectQuota                      = err("ErrGetProjectQuota", "获取项目配额失败")
	ErrGetNamespacesBelongsTo               = err("ErrGetNamespacesBelongsTo", "查询命令空间的项目归属失败")
	ErrListQuotaRecords                     = err("ErrListQuotaRecords", "查询项目资源配额列表失败")
	ErrListProject                          = err("ErrListProject", "获取项目列表失败")
	ErrListProjectByStates                  = err("ErrListProjectByStates", "通过state获取项目列表失败")
	ErrListProjectID                        = err("ErrListProjectID", "获取项目ID列表失败")
	ErrListProjectUsage                     = err("ErrListProjectUsage", "获取项目资源使用列表失败")
	ErrListAllProject                       = err("ErrListAllProject", "获取所有项目列表失败")
	ErrReferProject                         = err("ErrReferProject", "查看集群项目引用失败")
	ErrReferCluster                         = err("ErrReferCluster", "查看企业集群引用失败")
	ErrFillProjectQuota                     = err("ErrFillProjectQuota", "填充项目配额失败")
	ErrCreateApplication                    = err("ErrCreateApplication", "创建应用失败")
	ErrInitApplication                      = err("ErrInitApplication", "应用初始化失败")
	ErrListAppTemplates                     = err("ErrListAppTemplates", "获取应用模板列表失败")
	ErrGetAppIDByNames                      = err("ErrGetAppIDByNames", "批量获取应用ID失败")
	ErrUpdateApplication                    = err("ErrUpdateApplication", "更新应用失败")
	ErrPinApplication                       = err("ErrPinApplication", "pin应用失败")
	ErrUnPinApplication                     = err("ErrUnPinApplication", "unpin应用失败")
	ErrGetApplication                       = err("ErrGetApplication", "获取应用失败")
	ErrDeleteApplication                    = err("ErrDeleteApplication", "删除应用失败")
	ErrListApplication                      = err("ErrListApplication", "获取应用列表失败")
	ErrCountApplication                     = err("ErrCountApplication", "获取应用列表数量失败")
	ErrGetApplicationPublishItemRelation    = err("ErrGetApplicationPublishItemRelation", "获取应用发布关联失败")
	ErrUpdateApplicationPublishItemRelation = err("ErrUpdateApplicationPublishItemRelation", "设置应用发布关联失败")
	ErrRemoveApplicationPublishItemRelation = err("ErrRemoveApplicationPublishItemRelation", "删除应用发布关联失败")
	ErrListAppUsage                         = err("ErrListAppUsage", "获取应用资源使用列表失败")
	ErrCreateTicket                         = err("ErrCreateTicket", "创建工单失败")
	ErrUpdateTicket                         = err("ErrUpdateTicket", "更新工单失败")
	ErrDeleteTicket                         = err("ErrDeleteTicket", "删除工单失败")
	ErrCloseTicket                          = err("ErrCloseTicket", "关闭工单失败")
	ErrReopenTicket                         = err("ErrReopenTicket", "重新打开工单失败")
	ErrListTicket                           = err("ErrListTicket", "获取工单列表失败")
	ErrGetTicket                            = err("ErrGetTicket", "获取工单失败")
	ErrCreateComment                        = err("ErrCreateComment", "创建评论失败")
	ErrUpdateComment                        = err("ErrUpdateComment", "更新评论失败")
	ErrListComment                          = err("ErrListComment", "获取评论列表失败")
	ErrListUser                             = err("ErrListUser", "获取用户列表失败")
	ErrGetUser                              = err("ErrGetUser", "获取用户详情失败")
	ErrListActivity                         = err("ErrListActivity", "获取活动列表失败")
	ErrUploadImage                          = err("ErrUploadImage", "上传图片失败")
	ErrListPermission                       = err("ErrListPermission", "获取权限列表失败")
	ErrAccessPermission                     = err("ErrAccessPermission", "获取访问权限失败")
	ErrCheckPermission                      = err("ErrCheckPermission", "鉴权失败")
	ErrAddMember                            = err("ErrAddMember", "添加成员失败")
	ErrAddMemberOwner                       = err("ErrAddMemberOwner", "添加所有者失败，数量超过上限")
	ErrUpdateMemeberUserInfo                = err("ErrUpdateMemeberUserInfo", "更新成员信息失败")
	ErrDeleteMember                         = err("ErrDeleteMember", "删除成员失败")
	ErrListMember                           = err("ErrListMember", "获取成员列表失败")
	ErrListMemberRoles                      = err("ErrListMemberRoles", "获取成员角色列表失败")
	ListMembersWithoutExtraByScope          = err("ListMembersWithoutExtraByScope", "获取成员角色列表失败")
	GetMemberByUserAndScope                 = err("GetMemberByUserAndScope", "获取成员失败")
	ErrGetHost                              = err("ErrGetHost", "获取主机失败")
	ErrGetHostUsage                         = err("ErrGetHostUsage", "获取主机资源使用失败")
	ErrListHost                             = err("ErrListHost", "获取主机列表失败")
	ErrListHostResource                     = err("ErrListHostResource", "获取主机资源列表失败")
	ErrListHostUsage                        = err("ErrListHostUsage", "获取主机资源使用列表失败")
	ErrListContainer                        = err("ErrListContainer", "获取容器列表失败")
	ErrCreateCluster                        = err("ErrCreateCluster", "创建集群失败")
	ErrDereferenceCluster                   = err("ErrDereferenceCluster", "解除集群关联关系失败")
	ErrUpdateCluster                        = err("ErrUpdateCluster", "更新集群失败")
	ErrGetCluster                           = err("ErrGetCluster", "获取集群失败")
	ErrListCluster                          = err("ErrListCluster", "获取集群列表失败")
	ErrDeleteCluster                        = err("ErrDeleteCluster", "删除集群失败")
	ErrListClusterUsage                     = err("ErrListClusterUsage", "获取集群资源使用列表失败")
	ErrListAddonUsage                       = err("ErrListAddonUsage", "获取addon资源使用失败")
	ErrListComponentUsage                   = err("ErrListComponentUsage", "获取平台组件资源使用失败")
	ErrListRuntimeUsage                     = err("ErrListRuntimeUsage", "获取runtime资源使用列表失败")
	ErrListServiceUsage                     = err("ErrListServiceUsage", "获取service资源使用列表失败")
	ErrListInstanceUsage                    = err("ErrListInstanceUsage", "获取实例资源使用列表失败")
	ErrListInstance                         = err("ErrListInstance", "获取实例列表失败")
	ErrSyncInstance                         = err("ErrSyncInstance", "同步实例信息失败")
	ErrSearchCluster                        = err("ErrSearchCluster", "获取集群范围内搜索结果失败")
	ErrCreateNamespace                      = err("ErrCreateNamespace", "创建namespace失败")
	ErrDeleteNamespace                      = err("ErrDeleteNamespace", "删除namespace失败")
	ErrDeleteNamespaceRelation              = err("ErrDeleteNamespaceRelation", "删除namespace关联关系失败")
	ErrCreateNamespaceRelation              = err("ErrCreateNamespaceRelation", "删除namespace关联关系失败")
	ErrAddEnvConfig                         = err("ErrAddEnvConfig", "添加环境变量配置失败")
	ErrUpdateEnvConfig                      = err("ErrUpdateEnvConfig", "更新环境变量配置失败")
	ErrDeleteEnvConfig                      = err("ErrDeleteEnvConfig", "删除环境变量配置失败")
	ErrGetEnvConfig                         = err("ErrGetEnvConfig", "获取环境变量配置失败")
	ErrImportEnvConfig                      = err("ErrImportEnvConfig", "导入环境变量配置失败")
	ErrExportEnvConfig                      = err("ErrExportEnvConfig", "导出环境变量配置失败")
	ErrGetNamespaceEnvConfig                = err("ErrGetNamespaceEnvConfig", "获取指定namespace环境变量配置失败")
	ErrGetDeployEnvConfig                   = err("ErrGetDeployEnvConfig", "获取部署环境变量配置失败")
	ErrGetMultiNamespaceEnvConfigs          = err("ErrGetMultiNamespaceEnvConfigs", "获取多个namespace环境变量配置失败")
	ErrStatisticsClusterResource            = err("ErrStatisticsClusterResource", "统计项目/应用/主机和Runtime数量失败")
	ErrListClusterAbnormalHosts             = err("ErrListClusterAbnormalHosts", "获取集群异常主机数量失败")
	ErrStatisticsClusterLabelsPercentage    = err("ErrStatisticsClusterLabelsPercentage", "统计集群各标签占用主机数量失败")
	ErrListClusterServices                  = err("ErrListClusterServices", "获取集群服务和addon列表失败")
	ErrListOrgRunningTasks                  = err("ErrListOrgRunningTasks", "获取集群正在运行中的服务或者job列表失败")
	ErrDealTaskEvents                       = err("ErrDealTaskEvents", "处理接收到的任务事件失败")
	ErrCreateNotifyGroup                    = err("ErrCreateNotifyGroup", "创建通知组失败")
	ErrUpdateNotifyGroup                    = err("ErrUpdateNotifyGroup", "更新通知组失败")
	ErrDeleteNotifyGroup                    = err("ErrDeleteNotifyGroup", "删除通知组失败")
	ErrGetNotifyGroup                       = err("ErrGetNotifyGroup", "获取通知组失败")
	ErrQueryNotifyGroup                     = err("ErrQueryNotifyGroup", "查询通知组列表失败")
	ErrGetNotifyGroupDetail                 = err("ErrGetNotifyGroupDetail", "获取通知组详情失败")
	ErrQueryNotifyItem                      = err("ErrQueryNotifyItem", "查询通知项列表失败")
	ErrUpdateNotifyItem                     = err("ErrUpdateNotifyItem", "更新通知项失败")
	ErrCreateNotifyHistory                  = err("ErrCreateNotifyHistory", "创建通知历史失败")
	ErrQueryNotifyHistory                   = err("ErrQueryNotifyHistory", "查询通知发送历史失败")
	ErrCreateNotify                         = err("ErrCreateNotify", "创建通知失败")
	ErrQueryNotify                          = err("ErrQueryNotify", "查询通知列表失败")
	ErrDeleteNotify                         = err("ErrDeleteNotify", "删除通知失败")
	ErrGetNotify                            = err("ErrGetNotify", "获取通知失败")
	ErrUpdateNotify                         = err("ErrUpdateNotify", "更新通知失败")
	ErrNotifyEnable                         = err("ErrNotifyEnable", "启用通知失败")
	ErrNotifyDisable                        = err("ErrNotifyDisable", "禁用通知失败")
	ErrDeleteNotifySource                   = err("ErrDeleteNotifySource", "删除通知源失败")
	ErrGetLicense                           = err("ErrGetLicense", "获取license失败")
	ErrListHostStatus                       = err("ErrListHostStatus", "根据IP列表获取机器状态失败")
	ErrCreateLabel                          = err("ErrCreateLabel", "创建标签失败")
	ErrDeleteLabel                          = err("ErrDeleteLabel", "删除标签失败")
	ErrUpdateLabel                          = err("ErrUpdateLabel", "更新标签失败")
	ErrGetLabels                            = err("ErrGetLabels", "获取标签列表失败")
	ErrListByNamesAndProjectID              = err("ErrListByNamesAndProjectID", "通过name和projectID获取标签列表失败")
	ErrListLabelByIDs                       = err("ErrListLabelByIDs", "通过IDs获取标签列表失败")
	ErrListProjectResourceUsage             = err("ErrListProjectResourceUsage", "获取项目资源使用率历史数据失败")

	ErrGetIteration                  = err("ErrGetIteration", "查询迭代失败")
	ErrCreateIteration               = err("ErrCreateIteration", "创建迭代失败")
	ErrUpdateIteration               = err("ErrUpdateIteration", "更新迭代失败")
	ErrDeleteIteration               = err("ErrDeleteIteration", "删除迭代失败")
	ErrPagingIterations              = err("ErrPagingIterations", "分页查询迭代失败")
	ErrCreateIssue                   = err("ErrCreateIssue", "创建 issue 失败")
	ErrPagingIssues                  = err("ErrPagingIssues", "分页查询 issue 失败")
	ErrUpdateIssue                   = err("ErrUpdateIssue", "更新 issue 失败")
	ErrDeleteIssue                   = err("ErrDeleteIssue", "删除 issue 失败")
	ErrBatchUpdateIssue              = err("ErrBatchUpdateIssue", "批量更新 issue 失败")
	ErrUpdateIssueState              = err("ErrUpdateIssueState", "更新 issue 状态失败")
	ErrGetIssue                      = err("ErrGetIssue", "查询 issue 失败")
	ErrExportExcelIssue              = err("ErrExportExcelIssue", "导出 issue 失败")
	ErrImportExcelIssue              = err("ErrImportExcelIssue", "导入 issue 失败")
	ErrGetIssueManHourSum            = err("ErrGetIssueManHourSum", "查询任务总和失败")
	ErrGetIssueBugPercentage         = err("ErrGetIssueBugPercentage", "查询缺陷率失败")
	ErrGetIssueBugStatusPercentage   = err("ErrGetIssueBugStatusPercentage", "查询缺陷状态失败")
	ErrGetIssueBugSeverityPercentage = err("ErrGetIssueBugSeverityPercentage", "查询缺陷等级失败")
	ErrCreateIssueRelation           = err("ErrCreateIssueRelation", "添加关联事件失败")
	ErrGetIssueRelations             = err("ErrGetIssueRelations", "查看关联事件失败")
	ErrDeleteIssueRelation           = err("ErrDeleteIssueRelation", "删除关联事件失败")
	ErrBatchCreateIssueTestCaseRel   = err("ErrBatchCreateIssueTestCaseRel", "事件批量关联测试计划用例失败")
	ErrDeleteIssueTestCaseRel        = err("ErrDeleteIssueTestCaseRel", "事件取消关联测试计划用例失败")
	ErrListIssueTestCaseRels         = err("ErrListIssueTestCaseRels", "查询事件用例关联列表失败")
	ErrGetTestPlanCaseRel            = err("ErrGetTestPlanCaseRel", "查询事件关联的测试用例失败")

	ErrCreateIssueProperty      = err("ErrCreateIssueProperty", "创建事项字段失败")
	ErrUpdateIssueProperty      = err("ErrUpdateIssueProperty", "更新事项字段失败")
	ErrDeleteIssueProperty      = err("ErrDeleteIssueProperty", "删除事项字段失败")
	ErrGetIssueProperty         = err("ErrGetIssueProperty", "查询事项字段失败")
	ErrCreateIssuePropertyValue = err("ErrCreateIssuePropertyValue ", "创建事项字段枚举值失败")
	ErrDeleteIssuePropertyValue = err("ErrDeleteIssuePropertyValue ", "删除事项字段枚举值失败")

	ErrGetIssueStateRelation    = err("ErrGetIssueStateRelation", "事件获取状态关联失败")
	ErrUpdateIssueStateRelation = err("ErrUpdateIssueStateRelation", "事件修改状态关联失败")
	ErrCreateIssueState         = err("ErrCreateIssueState", "创建工作流状态失败")
	ErrDeleteIssueState         = err("ErrDeleteIssueState", "删除工作流状态失败")
	ErrGetIssueState            = err("ErrGetIssueState", "获取工作流状态失败")

	ErrCreateIssuePanel = err("ErrCreateIssuePanel", "创建自定义面板失败")
	ErrUpdateIssuePanel = err("ErrUpdateIssuePanel", "更新自定义面板失败")
	ErrDeleteIssuePanel = err("ErrDeleteIssuePanel", "删除自定义面板失败")
	ErrGetIssuePanel    = err("ErrGetIssuePanel", "查询自定义面板失败")

	ErrCreateMBox        = err("ErrCreateMBox", "创建站内信失败")
	ErrQueryMBox         = err("ErrQueryMBox", "查询站内信失败")
	ErrGetMBoxStats      = err("ErrGetMBoxStats", "获取站内信统计信息失败")
	ErrSetMBoxReadStatus = err("ErrSetMBoxReadStatus", "设置站内信已读失败")
	ErrCleanUnreadMboxs  = err("ErrCleanUnreadMboxs", "一键已读站内信失败")

	ErrCreateIssueStream = err("ErrCreateIssueStream", "创建活动记录列表失败")
	ErrPagingIssueStream = err("ErrPagingIssueStream", "分页查询活动记录失败")
	ErrListIssueStream   = err("ErrListIssueStream", "获取活动记录列表失败")

	ErrUploadFile          = err("ErrUploadFile", "上传文件失败")
	ErrUploadFileEncrypt   = err("ErrUploadFileEncrypt", "加密上传文件失败")
	ErrUploadTooLargeFile  = err("ErrUploadTooLargeFile", "上传的文件超过限制大小")
	ErrDownloadFile        = err("ErrDownloadFile", "下载文件失败")
	ErrDownloadFileDecrypt = err("ErrDownloadFileDecrypt", "解密下载文件失败")
	ErrCleanExpiredFile    = err("ErrCleanExpiredFile", "清理过期文件失败")
	ErrDeleteFile          = err("ErrDeleteFile", "删除文件失败")
	ErrBackup              = err("ErrBackup", "备份失败")
	ErrInvalidRef          = err("ErrInvalidRef", "invalid ref")

	ErrCreatePublisher = err("ErrCreatePublisher", "创建Publisher失败")
	ErrUpdatePublisher = err("ErrUpdatePublisher", "更新Publisher失败")
	ErrDeletePublisher = err("ErrDeletePublisher", "删除Publisher失败")
	ErrGetPublisher    = err("ErrGetPublisher", "获取Publisher失败")
	ErrListPublisher   = err("ErrListPublisher", "获取Publisher列表失败")

	ErrCreateNotice    = err("ErrCreateNotice", "创建公告失败")
	ErrUpdateNotice    = err("ErrUpdateNotice", "编辑公告失败")
	ErrPublishNotice   = err("ErrPublishNotice", "发布公告失败")
	ErrUnpublishNotice = err("ErrUnpublishNotice", "停用公告失败")
	ErrDeleteNotice    = err("ErrDeleteNotice", "删除公告失败")
	ErrListNotice      = err("ErrListNotice", "获取公告列表失败")

	ErrEnsurePhysicsNexusBlobStore     = err("ErrEnsurePhysicsNexusBlobStore", "保存 nexus blob store 失败")
	ErrEnsurePhysicsNexusRepo          = err("ErrEnsurePhysicsNexusRepo", "保存 nexus 物理 repo 失败")
	ErrGetPhysicsNexusRepo             = err("ErrGetPhysicsNexusRepo", "查询 nexus 物理 repo 失败")
	ErrEnsurePhysicsNexusUser          = err("ErrEnsurePhysicsNexusUser", "保存 nexus 物理 user 失败")
	ErrGetPhysicsNexusUser             = err("ErrGetPhysicsNexusUser", "查询 nexus 物理 user 失败")
	ErrGetNexusUserRecord              = err("ErrGetNexusUserRecord", "查询 nexus user 失败")
	ErrEnsureNexusRepoRecord           = err("ErrEnsureNexusRepoRecord", "保存 nexus repo 记录失败")
	ErrEnsureNexusUserRecord           = err("ErrEnsureNexusUserRecord", "保存 nexus user 记录失败")
	ErrListNexusRepos                  = err("ErrListNexusRepos", "查询 nexus repo 列表失败")
	ErrListNexusUsers                  = err("ErrListNexusUsers", "查询 nexus user 列表失败")
	ErrGetNexusRepoRecord              = err("ErrGetNexusRepoRecord", "查询 nexus repo 失败")
	ErrEncryptPassword                 = err("ErrEncryptPassword", "密码加密失败")
	ErrSyncConfigToPipelineCM          = err("ErrSyncConfigToPipelineCM", "同步配置至 pipeline 配置管理失败")
	ErrGetOrgNexus                     = err("ErrGetOrgNexus", "获取企业级别 nexus 配置失败")
	ErrShowOrgNexusPassword            = err("ErrShowOrgNexusPassword", "查询企业 nexus 密码失败")
	ErrHandleNexusDockerRepo           = err("ErrHandleNexusDockerRepo", "处理 nexus docker repo 失败")
	ErrGetNexusDockerCredentialByImage = err("ErrGetNexusDockerCredentialByImage", "根据镜像获取 docker 认证信息失败")

	ErrCreateCertificate = err("ErrCreateCertificate", "创建证书失败")
	ErrUpdateCertificate = err("ErrUpdateCertificate", "更新证书失败")
	ErrDeleteCertificate = err("ErrDeleteCertificate", "删除证书失败")
	ErrGetCertificate    = err("ErrGetCertificate", "获取证书失败")
	ErrListCertificate   = err("ErrListCertificate", "获取证书列表失败")

	ErrQuoteCertificate       = err("ErrQuoteCertificate", "应用引用证书失败")
	ErrCancelQuoteCertificate = err("ErrCancelQuoteCertificate", "取消引用证书失败")
	ErrListQuoteCertificate   = err("ErrListQuoteCertificate", "获取应用的证书列表失败")

	ErrPushCertificateConfigs = err("ErrPushCertificateConfigs", "推送证书配置到配置管理失败")

	ErrCreateApprove = err("ErrCreateApprove", "创建审批失败")
	ErrUpdateApprove = err("ErrUpdateApprove", "更新审批失败")
	ErrDeleteApprove = err("ErrDeleteApprove", "删除审批失败")
	ErrGetApprove    = err("ErrGetApprove", "获取审批失败")
	ErrListApprove   = err("ErrListApprove", "获取审批列表失败")

	ErrCreateLibReference      = err("ErrCreateLibReference", "创建库引用失败")
	ErrDeleteLibReference      = err("ErrDeleteLibReference", "删除库引用失败")
	ErrListLibReference        = err("ErrListLibReference", "获取库引用列表失败")
	ErrListLibReferenceVersion = err("ErrListLibReferenceVersion", "获取库引用版本列表失败")

	ErrApprovalStatusChanged = err("ErrApprovalStatusChanged", "审批流状态变更通知失败")

	ErrQueryBranchRule       = err("ErrQueryBranchRule", "查询分支规则失败")
	ErrCreateBranchRule      = err("ErrCreateBranchRule", "创建分支规则失败")
	ErrUpdateBranchRule      = err("ErrUpdateBranchRule", "更新分支规则失败")
	ErrDeleteBranchRule      = err("ErrDeleteBranchRule", "删除分支规则失败")
	ErrFillProjectBranchRule = err("ErrFillProjectBranchRule", "填充项目分支规则失败")

	ErrCreateAudit         = err("ErrCreateAudit", "创建审计事件失败")
	ErrListAudit           = err("ErrListAudit", "查看审计事件失败")
	ErrUpdateAuditSettings = err("ErrUpdateAuditSettings", "修改审计事件设置失败")
	ErrListAuditSettings   = err("ErrListAuditSettings", "查看审计事件设置失败")
	ErrExportExcelAudit    = err("ErrExportExcelAudit", "导出审计失败")

	ErrAddErrorLog  = err("ErrAddErrorLog", "记录错误日志失败")
	ErrListErrorLog = err("ErrListErrorLog", "查看错误日志失败")

	ErrUpdateApproval        = err("ErrupdateApproval", "更改审批记录失败")
	ErrGetReviewsByUserId    = err("ErrGetReviewsByUserId", "查询用户需要审批的列表失败")
	ErrCreateReview          = err("ErrCreateReview", "创建审批记录失败")
	ErrCreateReviewUser      = err("ErrCreateReviewUser", "创建用户审批记录失败")
	ErrGetReviewByTaskId     = err("ErrGetReviewByTaskId", "创建用户审批记录失败")
	ErrGetReviewsBySponsorId = err("ErrGetReviewsBySponsorId", "查询sponsorId(发起人)发起部署审批列表失败")
	ErrGetAuthorityByUserId  = err("ErrGetAuthorityByUserId", "查询用户权限失败")

	ErrCreateFileTreeNode        = err("ErrCreateFileTreeNode", "创建目录树节点失败")
	ErrDeleteFileTreeNode        = err("ErrDeleteFileTreeNode", "删除目录树节点失败")
	ErrUpdateSetBasicInfo        = err("ErrUpdateSetBasicInfo", "更新目录树节点基础信息失败")
	ErrMoveFileTreeNode          = err("ErrMoveFileTreeNode", "移动目录树节点失败")
	ErrCopyFileTreeNode          = err("ErrCopyFileTreeNode", "复制目录树节点失败")
	ErrGetFileTreeNode           = err("ErrGetFileTreeNode", "查询目录树节点详情失败")
	ErrListFileTreeNodes         = err("ErrListFileTreeNodes", "查询目录树节点列表失败")
	ErrListFileTreeNodeHistory   = err("ErrListFileTreeNodeHistory", "查询目录树节点历史列表失败")
	ErrFuzzySearchFileTreeNodes  = err("ErrFuzzySearchFileTreeNodes", "模糊搜索目录树节点失败")
	ErrSaveFileTreeNodePipeline  = err("ErrSaveFileTreeNodePipeline", "保存流水线失败")
	ErrFindFileTreeNodeAncestors = err("ErrFindFileTreeNodeAncestors", "目录树节点寻祖失败")

	ErrGetWorkBenchData = err("ErrGetWorkBenchData", "failed to query workbench data")

	ErrCreateAccessKey = err("ErrCreateSecretKey", "创建密钥对失败")
	ErrUpdateAccessKey = err("ErrCreateSecretKey", "更新密钥对失败")
	ErrGetAccessKey    = err("ErrGetAccessKey", "查询密钥对失败")
	ErrDeleteAccessKey = err("ErrDeleteAccessKey", "删除密钥对失败")

	ErrCreateSubscribe = err("ErrCreateSubscribe", "创建关注失败")
	ErrDeleteSubscribe = err("ErrDeleteSubscribe", "删除关注失败")
	ErrGetSubscribe    = err("ErrGetSubscribe", "获取关注信心失败")
)

func err(template, defaultValue string) *errorresp.APIError {
	return errorresp.New(errorresp.WithTemplateMessage(template, defaultValue))
}
