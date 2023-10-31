run-client:
	@go run cmd/client/main.go \
	--port=3000 \
	--is_server=false \
	--cert_path=certs/client.pem \
	--key_path=certs/client.key \
	debug=true

run-server:
	@go run cmd/server/main.go \
	--port=3000 \
	--is_server=true \
	--cert_path=certs/server.pem \
	--key_path=certs/server.key \
	debug=true

build-proto:
	@protoc \
	--go_out=generated \
	--go_opt=paths=source_relative \
	proto/messages.proto

generate-certs:
	@openssl genrsa -out server.key 2042 && \
	openssl req -new -nodes -x509 -out certs/client.pem -keyout certs/client.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=www.random.com/emailAddress=$1"