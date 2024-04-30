test:
	go test -v -short -cover ./...

server:
	go run main.go	

local:
	ENVIRONMENT=local go run .	

docker-run:
	docker run -p 8080:8080 -e ENVIRONMENT=local -e SIRIUS_HOST=192.168.143.202 -e SIRIUS_PORT=12580 \
	-d --name authhub authhub:latest

docker:
	docker build -f Dockerfile -t authhub:latest .

proto: 
	rm -rf pb/*
	rm -rf doc/swagger/*.swagger.json
	protoc --proto_path=proto \
	--go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger  --openapiv2_opt=allow_merge=true \
	--openapiv2_opt=merge_file_name=$(DB_NAME) --openapiv2_opt=json_names_for_fields=false \
    proto/*.proto

outerpb: 
	protoc --proto_path=outerpb \
	--go_out=outerpb --go_opt=paths=source_relative \
    --go-grpc_out=outerpb --go-grpc_opt=paths=source_relative \
    outerpb/openocr/*.proto

.PHONY: test local server mock docker proto outerpb