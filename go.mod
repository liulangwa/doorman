module github.com/youtube/doorman

go 1.16

require (
	github.com/chzyer/logex v1.1.10 // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e
	github.com/chzyer/test v0.0.0-20180213035817-a1ea475d72b1 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529
	github.com/golang/protobuf v1.5.2
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510
	github.com/pborman/uuid v1.2.1
	github.com/prometheus/client_golang v1.11.0
	go.etcd.io/etcd/client/v2 v2.305.0
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/examples v0.0.0-20210708170655-30dfb4b933a5
	google.golang.org/protobuf v1.26.0
)

//replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.6
