package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
    "distributed-systems-take-home/lib/client"
)


func main() {
    servers_env, exists := os.LookupEnv("SERVERS")  
    if !exists {
        log.Fatal("No servers provided, specify the servers using the SERVERS environment variable")
    }
    servers := strings.Split(servers_env, ",")
    c := client.AddServers(servers) 
    log.Println("Servers: ", servers)
    if len(servers) == 0 {
        log.Fatal("No comma separated servers provided, specify the servers using the SERVERS environment variable")
    }
    ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)

	// Register the channel to receive SIGTERM and SIGINT
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

    userIds := &client.UserIds{}
MAIN:
    for {
        select {
        case <-sigs:
            log.Println("Signal received, cancelling the context")
            cancel()
            break 
        case <-ctx.Done():
            log.Println("Context cancelled")
            break MAIN
        default:
            user := userIds.CreateUserOrUseExisting()
            if client.AnomolyProbability(user.BadActor) {
                user = client.UserInfo{
                    ID: user.ID, 
                    UserAgent: client.RandomUserAgent(), 
                    IPAddr: client.RandomIP(), 
                    BadActor: user.BadActor,
                }
            }
            err := c.SendData(ctx, user)
            if err != nil {
                log.Println("Error sending data: ", err)
            }
        }
    }
    log.Println("Exiting")
}

