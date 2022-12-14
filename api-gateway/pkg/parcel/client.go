package parcel
import (
	"fmt"
	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/parcel/pb"
	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ParcelServiceClient
}

func InitServiceClient(c *config.Config) pb.ParcelServiceClient {
	cc, err := grpc.Dial(c.ParcelSvcUrl,grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect ", err)
	}

	return pb.NewParcelServiceClient(cc)
}