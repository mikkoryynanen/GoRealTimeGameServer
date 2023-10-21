run-client:
	@go run cmd/client/main.go --port=3000 --is_server=false --cert_path=certs/client.pem --key_path=certs/client.key debug=true

run-server:
	@go run cmd/server/main.go --port=3000 --is_server=true --cert_path=certs/server.pem --key_path=certs/server.key debug=true