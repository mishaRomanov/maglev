package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mishaRomanov/maglev/internal/shared"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	storage *redis.Client
}

func New(storage *redis.Client) *Handler {
	return &Handler{
		storage: storage,
	}
}

// Handles /register endpoint
// TODO переписать под http формы (?) каким образом буду получать дату (?)
func (h *Handler) Register(ctx *gin.Context) {
	var userCreds *shared.UserCreds

	bindError := ctx.BindJSON(&userCreds)
	if bindError != nil {
		log.Errorf("Error while binding json: %v\n", bindError)
		ctx.JSON(
			http.StatusBadRequest,
			shared.CustomError{
				GivenError:  bindError.Error(),
				Description: "can't unmarshall the data you sent",
			})
		return
	}
	// Generate userId and return it
	userID, uidError := shared.GenetareNewUUID(*userCreds)
	if uidError != nil {
		log.Error("Error while generating userid")
		ctx.JSON(http.StatusInternalServerError,
			shared.CustomError{
				GivenError:  uidError.Error(),
				Description: "failed to generate userID",
			})
		return
	}

	if redisErr := h.storage.Set(shared.OneMinuteContextTimeOut(), userID, userCreds, time.Hour); redisErr != nil {
		log.Error("Failed to write data to redis")
		ctx.JSON(
			http.StatusInternalServerError,
			shared.CustomError{
				GivenError:  redisErr.Err().Error(),
				Description: "failed to write data to database",
			})
		return
	}
	ctx.JSON(
		http.StatusOK,
		shared.UserID{
			Userid: userID,
		})
}

func (h *Handler) MainPage(ctx *gin.Context) {
	//TODO заменить на пашин html
	ctx.String(http.StatusOK, "Hello from a main page!")
}
