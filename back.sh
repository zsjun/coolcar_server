 protoc -I=. --go_out=paths=source_relative:.  --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:. ./OldPackageTest/grpc_token_auth_test/proto/helloworld.proto 


 protoc -I=. --go_out=paths=source_relative:.  --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:. ./user_srv/proto/user.proto