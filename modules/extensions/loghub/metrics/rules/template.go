package rules

import (
	"github.com/erda-project/erda-infra/providers/i18n"
	"github.com/erda-project/erda/modules/monitor/core/metrics"
)

// ListConfigTemplate .
func (p *provider) ListConfigTemplate(scope string, lang i18n.LanguageCodes) []map[string]interface{} {
	return []map[string]interface{}{
		map[string]interface{}{
			"name": "nginx",
			"desc": p.t.Text(lang, "在 nginx 的访问日志中提取指标"),
		},
		map[string]interface{}{
			"name": "tomcat",
			"desc": p.t.Text(lang, "在 tomcat 的访问日志中提取指标"),
		},
	}
}

// GetConfigTemplate .
func (p *provider) GetConfigTemplate(scope, name string, lang i18n.LanguageCodes) map[string]interface{} {
	switch name {
	case "nginx":
		return map[string]interface{}{
			"name": "nginx",
			"filters": []*Tag{
				&Tag{
					Key:   "dice_project_name",
					Value: "<input your project name>",
				},
				&Tag{
					Key:   "dice_application_name",
					Value: "<input your application name>",
				},
				&Tag{
					Key:   "dice_service",
					Value: "<input your service name>",
				},
			},
			"content": "[21/Jul/2020:15:39:24 +0800]\t127.0.0.1, 42.120.75.141\t10.118.183.0\thttp\tterminus-org.dev.terminus.io\tGET /api/deployments/actions/list-pending-approval?pageSize=15&type=BUILD&pageNo=1 HTTP/1.1\t200\t4.608\t1307\t2309\thttp://terminus-org.dev.terminus.io/workBench/approval/my-approve/pending?operator[]=1000019&pageNo=1&type=BUILD\tMozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36\t-\t-\t-\t10.96.239.163:9529\t200\t4.608\t20571\t1",
			"processors": []*ProcessorConfig{
				&ProcessorConfig{
					Type: "regexp",
					Config: map[string]interface{}{
						"pattern": `^[^\t]*\t([^\t,]*)[^\t]*\t[^\t]*\t([^\t]*)\t([^\t]*)\t([^\t]*)\t([^\t]*)\t([^\t]*)\t([^\t]*)\t([^\t]*)\t([^\t]*)\t([^\t]*)\t[^\t]*\t[^\t]*\t[^\t]*\t([^\t]*)\t([^\t,]*)[^\t]*\t([^\t,]*)[^\t]*\t.*`,
						"keys": []*metrics.FieldDefine{
							&metrics.FieldDefine{Key: "userIp", Type: "string", Name: "userIp"},
							&metrics.FieldDefine{Key: "scheme", Type: "string", Name: "scheme"},
							&metrics.FieldDefine{Key: "host", Type: "string", Name: "host"},
							&metrics.FieldDefine{Key: "requestLine", Type: "string", Name: "requestLine"},
							&metrics.FieldDefine{Key: "status", Type: "string", Name: "status"},
							&metrics.FieldDefine{Key: "latency", Type: "number", Name: "latency"},
							&metrics.FieldDefine{Key: "requestSize", Type: "number", Name: "requestSize"},
							&metrics.FieldDefine{Key: "responseSize", Type: "number", Name: "responseSize"},
							&metrics.FieldDefine{Key: "referer", Type: "string", Name: "referer"},
							&metrics.FieldDefine{Key: "ua", Type: "string", Name: "ua"},
							&metrics.FieldDefine{Key: "upstreamAddr", Type: "string", Name: "upstreamAddr"},
							&metrics.FieldDefine{Key: "upstreamStatus", Type: "number", Name: "upstreamStatus"},
							&metrics.FieldDefine{Key: "upstreamLatency", Type: "number", Name: "upstreamLatency"},
						},
					},
				},
			},
		}
	case "tomcat":
		return map[string]interface{}{
			"name": "tomcat",
			"filters": []*Tag{
				&Tag{
					Key:   "dice_project_name",
					Value: "<input your project name>",
				},
				&Tag{
					Key:   "dice_application_name",
					Value: "<input your project name>",
				},
				&Tag{
					Key:   "dice_service",
					Value: "<input your project name>",
				},
			},
			"content": `11.11.11.11 - - [25/Jan/2000:14:00:01 +0100] "GET /1986.js HTTP/1.1" 200 932 "-" "Mozilla/5.0 (Windows; U; Windows NT 5.1; de; rv:1.9.1.7) Gecko/20091221 Firefox/3.5.7 GTB6"`,
			"processors": []*ProcessorConfig{
				&ProcessorConfig{
					Type: "regexp",
					Config: map[string]interface{}{
						"pattern": `^(?P<remoteIP>\S+) \S+ \S+ \[(?P<time>[\w:\/]+\s[+\-]\d{4})\] "(?P<method>\S+)\s?(?P<path>\S+)?\s?(?P<version>\S+)?" (?P<httpCode>\d{3}|-) (\d+|-)\s?"?([^"]*)"?\s?"?(?P<client>[^"]*)?"?$`,
						"keys": []*metrics.FieldDefine{
							&metrics.FieldDefine{Key: "remoteIP", Type: "string", Name: "remoteIP"},
							&metrics.FieldDefine{Key: "time", Type: "string", Name: "time"},
							&metrics.FieldDefine{Key: "method", Type: "string", Name: "method"},
							&metrics.FieldDefine{Key: "path", Type: "string", Name: "path"},
							&metrics.FieldDefine{Key: "version", Type: "string", Name: "version"},
							&metrics.FieldDefine{Key: "httpCode", Type: "string", Name: "httpCode"},
							&metrics.FieldDefine{Key: "group9", Type: "string", Name: "group9"},
							&metrics.FieldDefine{Key: "group10", Type: "string", Name: "group10"},
							&metrics.FieldDefine{Key: "ua", Type: "string", Name: "ua"},
						},
					},
				},
			},
		}
	}
	return nil
}
