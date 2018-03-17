package autodiscover

import (
	// include all metricbeat specific builders
	_ "github.com/elastic/beats/metricbeat/autodiscover/builder/metrics"

	// include all metricbeat specific appenders
	_ "github.com/elastic/beats/metricbeat/autodiscover/appender/kubernetes/token"
)
