package refreshtoken

import (
	"net/http"

	"github.com/golang-microservices/template/config"
	"github.com/golang-microservices/template/model"
	"github.com/labstack/echo/v4"
)

// Handler manage all request and dependency
type Handler struct {
	*config.Environment
}

// New return a new Handler
func New(env *config.Environment) *Handler {
	return &Handler{env}
}

// Handler register all path to echo.Echo
func (h *Handler) Handler(e *echo.Group) {
	e.GET("/refresh", h.refreshtoken(), h.JWT.RefreshTokentMiddleware(h.Config.Token.Refreshtoken.PublicKey))
}

func (h *Handler) refreshtoken() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse query string for query user by username
		request := model.GetUserByQueryString{}
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// Query user from DB by username
		user, err := h.Database.GetUser(request.Username)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}

		if user.Username != request.Username {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}

		accessToken, refreshToken, err := h.JWT.CreateNewTokens(h.Config.Token.Accesstoken.PrivateKey, h.Config.Token.Refreshtoken.PrivateKey, user.Email, "normal", h.Config.Token.Accesstoken.JWTTimeout, true)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		})
	}
}