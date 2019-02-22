package transport

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"

	"github.com/vikashvverma/eventers/pkg/api/event"
	"github.com/vikashvverma/eventers/pkg/utl/model"
)

// HTTP represents event http service
type HTTP struct {
	svc event.Service
}

// NewHTTP creates new event http service
func NewHTTP(svc event.Service, er *echo.Group) {
	h := HTTP{svc}
	ur := er.Group("/events")
	// swagger:route POST /v1/events events eventCreate
	// Creates new event.
	// responses:
	//  200: eventResp
	//  400: errMsg
	//  500: err
	ur.POST("", h.create)

	// swagger:operation GET /v1/events/{id} events getEvent
	// ---
	// summary: Returns a single event.
	// description: Returns a single event by its ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of event
	//   type: int
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/eventResp"
	//   "400":
	//     "$ref": "#/responses/err"
	//   "404":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.GET("/:id", h.view)
}

// Event create request
// swagger:model eventCreate
type createReq struct {
	Name     string    `json:"name" validate:"required,min=3"`
	Date     time.Time `json:"date" validate:"required"`
	Location string    `json:"location" validate:"required,min=3"`
}

func (h *HTTP) create(c echo.Context) error {
	r := new(createReq)

	if err := c.Bind(r); err != nil {
		return err
	}

	usr, err := h.svc.Create(c, eventers.Event{
		Name:     r.Name,
		Date:     r.Date,
		Location: r.Location,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, usr)
}

func (h *HTTP) view(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return eventers.ErrBadRequest
	}

	result, err := h.svc.View(c, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
