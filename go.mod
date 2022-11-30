module github.com/NPTLI/volc-sdk-golang

go 1.14

require (
	github.com/Shopify/sarama v1.30.1
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/cenkalti/backoff/v4 v4.1.2
	github.com/go-kit/kit v0.12.0
	github.com/gogo/protobuf v1.3.2
	github.com/google/go-querystring v1.1.0
	github.com/google/uuid v1.3.0
	github.com/minio/minio-go/v7 v7.0.34
	github.com/pierrec/lz4 v2.6.1+incompatible
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b
	google.golang.org/protobuf v1.28.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/volcengine/volc-sdk-golang => ../volc-sdk-golang
