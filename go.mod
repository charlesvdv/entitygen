module github.com/charlesvdv/entitygen

go 1.13

require (
	github.com/VividCortex/ewma v1.1.1 // indirect
	github.com/aws/aws-sdk-go v1.34.32 // indirect
	github.com/biogo/store v0.0.0-20200525035639-8c94ae1e7c9c // indirect
	github.com/certifi/gocertifi v0.0.0-20200922220541-2c3bb06c6054 // indirect
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/cockroachdb/cmux v0.0.0-20170110192607-30d10be49292 // indirect
	github.com/cockroachdb/cockroach v20.1.6+incompatible
	github.com/cockroachdb/errors v1.7.5 // indirect
	github.com/cockroachdb/ttycolor v0.0.0-20180709150743-a1d5aaeb377d // indirect
	github.com/codahale/hdrhistogram v0.9.0 // indirect
	github.com/getsentry/raven-go v0.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.15.0 // indirect
	github.com/jaegertracing/jaeger v1.19.2 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/lib/pq v1.8.0 // indirect
	github.com/lightstep/lightstep-tracer-go v0.21.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/petermattis/goid v0.0.0-20180202154549-b0b1615b78e5 // indirect
	github.com/prometheus/common v0.14.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0 // indirect
	github.com/stretchr/testify v1.6.1
	go.etcd.io/etcd v3.3.25+incompatible // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	golang.org/x/net v0.0.0-20200925080053-05aa5d4ee321 // indirect
	golang.org/x/sys v0.0.0-20200923182605-d9f96fdee20d // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/tools v0.0.0-20200925191224-5d1fdd8fa346 // indirect
)

replace (
	github.com/abourget/teamcity => github.com/cockroachdb/teamcity v0.0.0-20180905144921-8ca25c33eb11
	github.com/cockroachdb/cockroach => github.com/cockroachdb/cockroach-gen v0.0.0-20200926032514-a63d9f6f2696
	github.com/openzipkin-contrib/zipkin-go-opentracing => github.com/openzipkin-contrib/zipkin-go-opentracing v0.3.5
	go.etcd.io/etcd => github.com/cockroachdb/etcd v0.4.7-0.20200615211340-a17df30d5955
	vitess.io/vitess => github.com/cockroachdb/vitess v2.2.0-rc.1.0.20180830030426-1740ce8b3188+incompatible
)
