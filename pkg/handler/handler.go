package handler
import (
	"github.com/gin-gonic/gin"
	"github.com/inzarubin80/booking-services/pkg/service"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files" 
	_ "github.com/inzarubin80/booking-services/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/typeBusines")
		{
			lists.POST("/", h.createTypeBusines)
			lists.GET("/", h.getAllTypeBusines)
			lists.GET("/:id", h.getTypeBusinesById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteTypeBusines)
		}
		
		companies := api.Group("/companies")
		{
			companies.POST("/", h.createCompanies)
			companies.GET("/", h.getAllCompanies)
			companies.GET("/:id", h.getCompaniesById)
			companies.PUT("/:id", h.updateСompany)
			companies.DELETE("/:id", h.deleteСompany)
		}		

		serviceСenters := api.Group("/serviceСenters")
		{
			serviceСenters.POST("/", h.createServiceСenter)
			serviceСenters.GET("/", h.getAllServiceСenter)
			serviceСenters.GET("/:id", h.getServiceСenterById)
			serviceСenters.PUT("/:id", h.updateServiceСenter)
			serviceСenters.DELETE("/:id", h.deleteServiceСenter)
		}		

		serviceGroups := api.Group("/serviceGroups")
		{
			serviceGroups.POST("/", h.createServiceGroup)
			serviceGroups.GET("/", h.getAllServiceGroups)
			serviceGroups.GET("/:id", h.getServiceGroupById)
			serviceGroups.PUT("/:id", h.updateServiceGroup)
			serviceGroups.DELETE("/:id", h.deleteServiceGroup)
		}
		
		serviceProducers := api.Group("/serviceProducers")
		{
			serviceProducers.POST("/", h.createServiceProducer)
			serviceProducers.GET("/", h.getAllServiceProducers)
			serviceProducers.GET("/:id", h.getServiceProducerById)
			serviceProducers.PUT("/:id", h.updateServiceProducers)
			serviceProducers.DELETE("/:id", h.deleteServiceProducer)
		}


		serviceItems := api.Group("/serviceItems")
		{
			serviceItems.POST("/", h.createServiceItem)
			serviceItems.GET("/", h.getAllServiceItems)
			serviceItems.GET("/:id", h.getServiceItemById)
			serviceItems.PUT("/:id", h.updateServiceItems)
			serviceItems.DELETE("/:id", h.deleteServiceItem)
		}


		slots := api.Group("/slots")
		{
			slots.POST("/", h.createSlot)
			slots.GET("/", h.getAllSlots)
			slots.GET("/:id", h.getSlotById)
			slots.PUT("/:id", h.updateSlot)
			slots.DELETE("/:id", h.deleteSlot)
		}

		
		
	}
	return router
}