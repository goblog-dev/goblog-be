package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/michaelwp/goblog/api/controller"
	"github.com/michaelwp/goblog/tool"
	"net/http"
	"strconv"
	"strings"
)

func AuthMiddleware(config *controller.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		response := &controller.Response{
			Status:    controller.ERROR,
			Message:   "unauthorized",
			Translate: "unauthorized",
			HttpCode:  http.StatusUnauthorized,
		}

		bearerToken := c.Request.Header.Get("Authorization")
		if bearerToken == "" {
			response.Message = tool.PrintLog("get header authorization", errors.New("token required")).Error()
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		bearerTokenSplit := strings.Split(bearerToken, " ")
		if len(bearerTokenSplit) < 2 {
			response.Message = tool.PrintLog("bearer token split", errors.New("token required")).Error()
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		token := bearerTokenSplit[1]

		claims, err := tool.VerifyJWT(token)
		if err != nil {
			response.Message = tool.PrintLog("verify JWT", err).Error()
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		userIdFloat := claims["id"].(float64)
		userIdStr := strconv.FormatUint(uint64(userIdFloat), 10)

		resultToken, err := config.RedisClient.Get(c, userIdStr).Result()
		if err != nil {
			response.Message = tool.PrintLog("get redis token", err).Error()
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		if resultToken != token {
			response.Message = tool.PrintLog("verify token with redis", errors.New("token invalid")).Error()
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		userIdInt, err := strconv.ParseInt(userIdStr, 10, 64)
		if resultToken != token {
			response.Message = tool.PrintLog("parseInt userId", err).Error()
			c.JSON(http.StatusInternalServerError, response)
			c.Abort()
			return
		}

		c.Set("user_id", userIdInt)
		c.Next()
	}
}
