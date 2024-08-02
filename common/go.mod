module github.com/ccjshop/go-mall/common

go 1.22.4

replace github.com/ccjshop/go-mall/proto => ../proto

require (
	github.com/ccjshop/go-mall/proto v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.21.0
	github.com/rs/zerolog v1.33.0
	github.com/shopspring/decimal v1.4.0
	golang.org/x/crypto v0.25.0
	google.golang.org/grpc v1.65.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.11
)

require (
	github.com/envoyproxy/protoc-gen-validate v1.0.4 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240730163845-b1a4ccb954bf // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240725223205-93522f1f2a9f // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
