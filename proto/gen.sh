echo "生成 rpc 代码"

OUT=./

# admin

protoc \
-I ${OUT} \
-I model/*.proto \
--go_out=":./mall/"  \
--go-grpc_out=":./mall/"  \
--grpc-gateway_out=":./mall/" \
--validate_out="lang=go:./mall/" \
admin/*.proto model/*.proto


protoc \
-I ${OUT} \
-I model/*.proto \
--go_out=":./mall/"  \
--go-grpc_out=":./mall/"  \
--grpc-gateway_out=":./mall/" \
--validate_out="lang=go:./mall/" \
portal/*.proto


