module github.com/ccjshop/go-mall/application/portal

go 1.22.4

replace (
	github.com/ccjshop/go-mall/common => ../../common
	github.com/ccjshop/go-mall/proto => ../../proto
)

require (
	github.com/ccjshop/go-mall/common v0.0.0-00010101000000-000000000000
	github.com/ccjshop/go-mall/proto v0.0.0-00010101000000-000000000000
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.21.0
	github.com/ilyakaznacheev/cleanenv v1.5.0
	github.com/rs/cors v1.11.0
	github.com/shopspring/decimal v1.4.0
	golang.org/x/net v0.27.0
	google.golang.org/grpc v1.65.0
	gorm.io/gorm v1.25.11
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.4 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/rs/zerolog v1.33.0 // indirect
	golang.org/x/crypto v0.25.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240730163845-b1a4ccb954bf // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240725223205-93522f1f2a9f // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/mysql v1.5.7 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
