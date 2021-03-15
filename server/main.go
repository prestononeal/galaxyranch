package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	// Flow support
	"context"
	"fmt"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
	"google.golang.org/grpc"
)

func handleErr(err error) {
	if err != nil {
		 panic(err)
	}
}

func getMoments(flowClient *client.Client, ctx context.Context, addr string) {
	getMomentScript := `
			import TopShot from 0x0b2a3299cc857e29
			pub fun main(owner:Address): {String: String} {
				let acct = getAccount(owner)
				let collectionRef = acct.getCapability(/public/MomentCollection)!.borrow<&{TopShot.MomentCollectionPublic}>() ?? panic("Could not borrow capability from public collection")
				let oneID = collectionRef.getIDs()[0]!
				let moment = collectionRef.borrowMoment(id: oneID)!
				return TopShot.getPlayMetaData(playID: moment.data.playID)!
			}
`

	res, err := flowClient.ExecuteScriptAtLatestBlock(context.Background(), []byte(getMomentScript), []cadence.Value{
		cadence.BytesToAddress(flow.HexToAddress(addr).Bytes()),
	})
	if err != nil {
		fmt.Println("error fetching sale moment from flow: %w", err)
	} else {
		fmt.Println("res: %w", res)
	}
}

func main() {

	// Set up connection to flow chain
	flowClient, err := client.New("access.mainnet.nodes.onflow.org:9000", grpc.WithInsecure())
  handleErr(err)
	ctx := context.Background()
  err = flowClient.Ping(ctx)
  handleErr(err)

	// Test
	getMoments(flowClient, ctx, "0xee95377cce1c3f2b")

	r := gin.Default()
	// Dont worry about this line just yet, it will make sense in the Dockerise bit!
	r.Use(static.Serve("/", static.LocalFile("./web", true)))
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}