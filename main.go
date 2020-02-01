package main

import ("fmt"
	"golang.org/x/net/context"
    "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	
	ba "github.com/prokopovich/grpc-server/api/basic"
)

const (
    port = ":50051"
)

type apiBasic struct {
}

func (a *apiBasic) Test(ctx context.Context,in *ba.TestRequest) (*ba.TestResponse, error){
	return &ba.TestResponse{
		"ololo"+in.Message
	}, nil
}

func main() {
	fmt.Println(sayHi("Marco"))
}

func sayHi(person string) string {
	return fmt.Sprintf("Hi %s", person)
}
