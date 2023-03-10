package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/server"
	"flag"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)
var addr = flag.String("addr", ":8080", "address listen")
var authAddr = flag.String("auth_addr", "localhost:8081", "address for auth service")
var tripAddr = flag.String("trip_addr", "localhost:8082", "address for trip service")
var profileAddr = flag.String("profile_addr", "localhost:8082", "address for profile service")
var carAddr = flag.String("car_addr", "localhost:8084", "address for car service")

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile)
	lg, err := server.NewZapLogger();
	if err != nil {
		log.Fatalf("cannot create zap:logger: %v", err)
	}
	// 建立空的上下文
	c := context.Background();
	c, cancel := context.WithCancel(c);
	defer cancel()
	// 分发器，当有服务的时候，分发到哪个服务器上面
	// customMar := jsonpb.Marshaler{
	// 	EnumsAsInts: true,
	// }
	// mar := runtime.JSONPb(customMar)
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseEnumNumbers:true,
				UseProtoNames: true,
			},
		},
	))
	serverConfig := []struct{
		name string;
		addr string;
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{
			name:"auth",
			addr:"localhost:8081",
			registerFunc:authpb.RegisterAuthServiceHandlerFromEndpoint,
		},
		{
			name:"trip",
			addr:"localhost:8082",
			registerFunc:rentalpb.RegisterTripServiceHandlerFromEndpoint,
		},
		{
			name:"profile",
			addr:"localhost:8082",
			registerFunc:rentalpb.RegisterProfileServiceHandlerFromEndpoint,
		},

	}
	for _, s := range serverConfig {
		err := s.registerFunc(c, mux, ":8081", []grpc.DialOption{grpc.WithInsecure()})
		if err != nil {
			log.Fatalf("cannot register service %s: %v",s.name,err)
		}
	}
	lg.Sugar().Infof("grpc server started ad: %s", addr);
	err = http.ListenAndServe(*addr, mux);
	if err != nil {
		log.Fatalf("cannot listen to 8080 %v", err)
	}
	

}

package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	carpb "coolcar/car/api/gen/v1"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/auth"
	"coolcar/shared/server"
	"log"
	"net/http"
	"net/textproto"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/namsral/flag"
	"google.golang.org/grpc"
)

var addr = flag.String("addr", ":8080", "address to listen")
var authAddr = flag.String("auth_addr", "localhost:8081", "address for auth service")
var tripAddr = flag.String("trip_addr", "localhost:8082", "address for trip service")
var profileAddr = flag.String("profile_addr", "localhost:8082", "address for profile service")
var carAddr = flag.String("car_addr", "localhost:8084", "address for car service")

func main() {
	flag.Parse()

	lg, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create zap logger: %v", err)
	}
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			EnumsAsInts: true,
			OrigName:    true,
		},
	), runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
		if key == textproto.CanonicalMIMEHeaderKey(runtime.MetadataHeaderPrefix+auth.ImpersonateAccountHeader) {
			return "", false
		}
		return runtime.DefaultHeaderMatcher(key)
	}))

	serverConfig := []struct {
		name         string
		addr         string
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{
			name:         "auth",
			addr:         *authAddr,
			registerFunc: authpb.RegisterAuthServiceHandlerFromEndpoint,
		},
		{
			name:         "trip",
			addr:         *tripAddr,
			registerFunc: rentalpb.RegisterTripServiceHandlerFromEndpoint,
		},
		{
			name:         "profile",
			addr:         *profileAddr,
			registerFunc: rentalpb.RegisterProfileServiceHandlerFromEndpoint,
		},
		{
			name:         "car",
			addr:         *carAddr,
			registerFunc: carpb.RegisterCarServiceHandlerFromEndpoint,
		},
	}

	for _, s := range serverConfig {
		err := s.registerFunc(
			c, mux, s.addr,
			[]grpc.DialOption{grpc.WithInsecure()},
		)
		if err != nil {
			lg.Sugar().Fatalf("cannot register service %s: %v", s.name, err)
		}
	}
	http.HandleFunc("/healthz", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("ok"))
	})
	http.Handle("/", mux)
	lg.Sugar().Infof("grpc gateway started at %s", *addr)
	lg.Sugar().Fatal(http.ListenAndServe(*addr, nil))
}


// func startGRPCGateway() {
	
// }