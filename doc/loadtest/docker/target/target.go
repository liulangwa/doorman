package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"net"
	"net/http"

	log "github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

var (
	port      = flag.Int("port", 0, "port to bind to")
	debugPort = flag.Int("debug_port", 0, "port to bind to")
)

var (
	requests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "requests",
		Help: "How many requests have been served.",
	},
		[]string{"resource"})
)

func init() {
	prometheus.MustRegister(requests)
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	requests.WithLabelValues(in.Name).Inc()
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		log.Exitf("failed to listen: %v", err)
	}

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(fmt.Sprintf(":%v", *debugPort), nil)
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
