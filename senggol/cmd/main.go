package main

import (
	"senggol/controller/httpserver"
	"senggol/pkg"
	"senggol/repository/postgresql"
	"senggol/service"
)

func main() {
	repositories, err := postgresql.GetRepositories("db", "5432", "admin", "senggol", "admin123", "disable")
	if err != nil {
		panic(err)
	}

	services := service.Services{
		Register:             service.NewRegisterService(repositories.CreateUser, pkg.NewSequenceNumberGenerator()),
		GetAuthenticatedUser: service.NewGetAuthenticatedUserService(repositories.GetUserByUsername),
		GetPeers:             service.NewGetPeersService(repositories.GetPeers, repositories.GetPeersCount),
		PostDirectMessage:    service.NewPostDirectMessageService(repositories.CreateDirectMessage, pkg.NewSequenceNumberGenerator(), pkg.NewSequenceNumberGenerator()),
		GetDirectMessages:    service.NewGetDirectMessagesService(repositories.GetPeerDirectMessages),
		SetMessageSeen:       service.NewSetMessageSeenService(repositories.UpdateDirectMessageSeenAt),
	}

	httpserver.StartServer("8080", services)
}
