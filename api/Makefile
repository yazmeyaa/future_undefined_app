proto_grpc_gen = --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative

proto_path_projects = ./proto/projects/projects.proto
proto_path_tasklist = ./proto/task_list/task_list.proto

gen-projects:
	protoc $(proto_grpc_gen) $(proto_path_projects)

gen-tasklist:
	protoc $(proto_grpc_gen) $(proto_path_tasklist)

gen-all: gen-projects gen-tasklist