package workhandler

import (
	"fmt"
	"log"
	"os"

	"github.com/MrDweller/service-registry-connection/models"
	serviceregistry "github.com/MrDweller/service-registry-connection/service-registry"
	httpserver "github.com/MrDweller/work-handler/http-server"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type WorkHandler struct {
	models.SystemDefinition
	ServiceRegistryConnection serviceregistry.ServiceRegistryConnection
	Services                  []models.ServiceDefinition
	Service                   Service
}

func New(address string, port int, systemName string, serviceRegistryAddress string, serviceRegistryPort int, services []models.ServiceDefinition) (*WorkHandler, error) {
	system := models.SystemDefinition{
		Address:            address,
		Port:               port,
		SystemName:         systemName,
		AuthenticationInfo: "",
	}

	serviceRegistryConnection, err := serviceregistry.NewConnection(
		serviceregistry.ServiceRegistry{
			Address: serviceRegistryAddress,
			Port:    serviceRegistryPort,
		},
		serviceregistry.ServiceRegistryImplementationType(os.Getenv("SERVICE_REGISTRY_IMPLEMENTATION")),
		models.CertificateInfo{
			CertFilePath: os.Getenv("CERT_FILE_PATH"),
			KeyFilePath:  os.Getenv("KEY_FILE_PATH"),
			Truststore:   os.Getenv("TRUSTSTORE_FILE_PATH"),
		},
	)
	if err != nil {
		return nil, err
	}

	return &WorkHandler{
		SystemDefinition:          system,
		ServiceRegistryConnection: serviceRegistryConnection,
		Services:                  services,
		Service:                   &ServiceImplementation{},
	}, nil
}

func (w *WorkHandler) Run() error {
	router := gin.Default()

	url := fmt.Sprintf("%s:%d", w.Address, w.Port)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	w.setupEnpoints(router)
	w.registerServices()

	log.Printf("Starting digital twin framework on: https://%s", url)
	log.Printf("Swagger documentation is available on: https://%s", url+"/docs/index.html")

	server, err := httpserver.NewServer(url, router)
	if err != nil {
		return err
	}
	err = server.StartServer()
	return err

}

func (w *WorkHandler) setupEnpoints(router *gin.Engine) error {
	controller := Controller{service: w.Service}

	router.POST("/create-work", controller.CreateWork)
	router.POST("/assign-worker", controller.CheckCurrentWorker, controller.AssignWorker)
	// router.POST("/worker", controller.RegisterWorker)

	return nil
}

func (w *WorkHandler) registerServices() {
	for _, service := range w.Services {
		w.ServiceRegistryConnection.RegisterService(
			service,
			[]string{
				"HTTP-SECURE-JSON",
			},
			map[string]string{},
			w.SystemDefinition,
		)
	}

}

func (w *WorkHandler) Stop() error {
	log.Printf("Unregistering the work hadnler services from the service registry!")
	for _, service := range w.Services {
		err := w.ServiceRegistryConnection.UnRegisterService(service, w.SystemDefinition)
		if err != nil {
			return err
		}
	}

	err := w.ServiceRegistryConnection.UnRegisterSystem(w.SystemDefinition)
	if err != nil {
		return err
	}

	return nil

}
