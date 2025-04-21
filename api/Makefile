proto_grpc_gen = --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative

proto_path_projects = ./proto/projects/projects.proto

gen-projects:
	protoc $(proto_grpc_gen) $(proto_path_projects)


gen-all: gen-projects