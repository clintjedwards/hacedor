package master

import (
	"context"
	"net"
	"time"

	"github.com/clintjedwards/cursor/api"
	"github.com/clintjedwards/cursor/config"
	"github.com/clintjedwards/cursor/storage"
	"github.com/clintjedwards/cursor/utils"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// CursorMaster represents a cursor master server
type CursorMaster struct {
	storage storage.Engine
	config  *config.Config
}

func initCursorMaster() *CursorMaster {
	cursorMaster := CursorMaster{}

	config, err := config.FromEnv()
	if err != nil {
		utils.StructuredLog(utils.LogLevelFatal, "failed to get config", err)
	}

	storage, err := storage.InitStorage(storage.StorageEngineBoltDB)
	if err != nil {
		utils.StructuredLog(utils.LogLevelFatal, "failed init storage", err)
	}

	cursorMaster.config = config
	cursorMaster.storage = storage

	return &cursorMaster
}

// CreatePipeline registers a new pipeline
func (master *CursorMaster) CreatePipeline(context context.Context, request *api.CreatePipelineRequest) (*api.CreatePipelineResponse, error) {

	newPipeline := api.Pipeline{
		Id:          string(utils.GenerateRandString(master.config.Master.IDLength)),
		Name:        request.Name,
		Description: request.Description,
		Created:     time.Now().Unix(),
		Modified:    time.Now().Unix(),
		GitRepo: &api.GitRepo{
			Url:    request.GitRepo.Url,
			Branch: request.GitRepo.Branch,
		},
	}

	protoNewPipeline, err := proto.Marshal(&newPipeline)
	if err != nil {
		return nil, err
	}

	err = master.storage.Add("pipelines", []byte(newPipeline.Id), protoNewPipeline)
	if err != nil {
		return nil, err
	}

	return &api.CreatePipelineResponse{Id: newPipeline.Id}, nil
}

func (master *CursorMaster) ListPipelines(context context.Context, request *api.ListPipelinesRequest) (*api.ListPipelinesResponse, error) {
	return nil, nil
}

func (master *CursorMaster) GetPipeline(context context.Context, request *api.GetPipelineRequest) (*api.GetPipelineResponse, error) {
	return nil, nil
}

func (master *CursorMaster) DeletePipeline(context context.Context, request *api.DeletePipelineRequest) (*api.DeletePipelineResponse, error) {
	return nil, nil
}

func (master *CursorMaster) RegisterWorker(context context.Context, request *api.RegisterWorkerRequest) (*api.RegisterWorkerResponse, error) {
	return nil, nil
}

// StartServer initializes a GRPC server
func StartServer() {
	config, err := config.FromEnv()
	if err != nil {
		utils.StructuredLog(utils.LogLevelFatal, "failed to get config", err)
	}

	listen, err := net.Listen("tcp", config.ServerURL)
	if err != nil {
		utils.StructuredLog(utils.LogLevelFatal, "could not initialize tcp listener", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	cursorMaster := initCursorMaster()

	api.RegisterCursorMasterServer(grpcServer, cursorMaster)

	utils.StructuredLog(utils.LogLevelInfo,
		"started cursor master",
		map[string]string{"server_url": config.ServerURL})
	grpcServer.Serve(listen)
}
