module github.com/pachyderm/pachyderm/v2

go 1.16

require (
	cloud.google.com/go v0.84.0
	cloud.google.com/go/storage v1.10.0
	github.com/Azure/azure-sdk-for-go v54.0.0+incompatible
	github.com/Microsoft/hcsshim v0.8.7 // indirect
	github.com/aws/aws-lambda-go v1.17.0
	github.com/aws/aws-sdk-go v1.38.41
	github.com/blend/go-sdk v1.20210908.5 // indirect
	github.com/c-bata/go-prompt v0.2.3
	github.com/cevaris/ordered_map v0.0.0-20190319150403-3adeae072e73
	github.com/chmduquesne/rollinghash v4.0.0+incompatible
	github.com/containerd/continuity v0.0.0-20200107194136-26c1120b8d41 // indirect
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f
	github.com/dexidp/dex v0.0.0-00010101000000-000000000000
	github.com/dexidp/dex/api/v2 v2.0.0
	github.com/dlclark/regexp2 v1.2.0 // indirect
	github.com/dlmiddlecote/sqlstats v1.0.2
	github.com/dnaeon/go-vcr v1.2.0 // indirect
	github.com/docker/docker v20.10.6+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/docker/go-units v0.4.0
	github.com/elazarl/goproxy v0.0.0-20191011121108-aa519ddbe484 // indirect
	github.com/evanphx/json-patch v4.11.0+incompatible
	github.com/fatih/camelcase v1.0.0
	github.com/fatih/color v1.9.0
	github.com/fsouza/go-dockerclient v1.4.1
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grafana/loki v1.6.1
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.1-0.20191002090509-6af20e3a5340
	github.com/hanwen/go-fuse/v2 v2.0.3
	github.com/hashicorp/golang-lru v0.5.4
	github.com/itchyny/gojq v0.11.2
	github.com/jackc/pgconn v1.10.0
	github.com/jackc/pgerrcode v0.0.0-20201024163028-a0d42d470451
	github.com/jackc/pgx/v4 v4.13.0
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/juju/ansiterm v0.0.0-20180109212912-720a0952cc2a
	github.com/kr/pretty v0.2.1 // indirect
	github.com/lib/pq v1.10.2
	github.com/lunixbochs/vtclean v1.0.0 // indirect
	github.com/mattn/go-isatty v0.0.12
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/mattn/go-tty v0.0.3 // indirect
	github.com/minio/minio-go/v6 v6.0.56
	github.com/minio/minio-go/v7 v7.0.14
	github.com/moby/sys/mount v0.3.0 // indirect
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6 // indirect
	github.com/onsi/ginkgo v1.15.0 // indirect
	github.com/onsi/gomega v1.10.5 // indirect
	github.com/opentracing-contrib/go-grpc v0.0.0-20210225150812-73cb765af46e
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pachyderm/ohmyglob v0.0.0-20210308211843-d5b47775fc36
	github.com/pachyderm/s2 v0.0.0-20200609183354-d52f35094520
	github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4
	github.com/pkg/errors v0.9.1
	github.com/pkg/term v0.0.0-20190109203006-aa71e9d9e942 // indirect
	github.com/pquerna/cachecontrol v0.0.0-20180517163645-1555304b9b35 // indirect
	github.com/prometheus/client_golang v1.11.0
	github.com/prometheus/common v0.26.0
	github.com/remyoudompheng/bigfft v0.0.0-20170806203942-52369c62f446 // indirect
	github.com/robfig/cron v1.2.0
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
	github.com/segmentio/analytics-go v0.0.0-20160426181448-2d840d861c32
	github.com/segmentio/backo-go v0.0.0-20160424052352-204274ad699c // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	github.com/uber/jaeger-client-go v2.28.0+incompatible
	github.com/vbauerster/mpb/v6 v6.0.2
	github.com/wcharczuk/go-chart v2.0.1+incompatible
	github.com/x-cray/logrus-prefixed-formatter v0.5.2
	github.com/xtgo/uuid v0.0.0-20140804021211-a0b114877d4c // indirect
	go.etcd.io/etcd/api/v3 v3.5.1
	go.etcd.io/etcd/client/v3 v3.5.1
	go.etcd.io/etcd/server/v3 v3.5.1
	go.uber.org/automaxprocs v1.4.0
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5
	golang.org/x/oauth2 v0.0.0-20210615190721-d04028783cf1
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20211025201205-69cdffdb9359
	golang.org/x/term v0.0.0-20210220032956-6a3ed077a48d
	google.golang.org/api v0.49.0
	google.golang.org/grpc v1.38.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	gotest.tools/v3 v3.0.3 // indirect
	k8s.io/api v0.22.3
	k8s.io/apimachinery v0.22.3
	k8s.io/client-go v12.0.0+incompatible
	modernc.org/mathutil v1.0.0
)

// until the changes in github.com/pachyderm/dex are upstreamed to github.com/dexidp/dex, we swap in our repo
replace github.com/dexidp/dex => github.com/pachyderm/dex v0.0.0-20211020185745-ebfeda600c26

// loki requires k8s.io/client-go@v12.0.0+incompatible via dependency chain ending in:
// ... -> github.com/prometheus/alertmanager@v0.19.0
// -> github.com/prometheus/prometheus@v0.0.0-20190818123050-43acd0e2e93f
replace k8s.io/client-go => k8s.io/client-go v0.22.3

// Depending on both etcd v3.5.1 and loki v1.6.0 causes a conflict
// since etcd requires more up to date dependencies on prometheus. Specifically:
// 1.) etcd version v3.5.1 requires client_golang@v1.11.0 which requires github.com/prometheus/common@v0.26.0
// 2.) loki v1.6.0 requires github.com/prometheus/common@v0.10.0
// the two cannot live in harmony because the config.NewClientFromConfig functions signature changed several times,
// and was updated to the signature present in v0.26.0 during the v0.21.0 release
replace github.com/prometheus/common => github.com/prometheus/common v0.9.1
