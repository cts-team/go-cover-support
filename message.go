package go_cover_support

import "net/http"

type ResponseMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseContext interface {
	JSON(code int, obj interface{})
	AbortWithStatusJSON(code int, jsonObj interface{})
}

func Msg(data ...interface{}) ResponseMessage {
	result := ResponseMessage{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    nil,
	}
	count := len(data)
	if count > 0 {
		result.Data = data[0]
	}
	if count > 1 {
		if message, ok := data[1].(string); ok && message != "" {
			result.Message = message
		}
	}
	if count > 2 {
		if code, ok := data[2].(int); ok && code > 0 {
			result.Code = code
		}
	}
	return result
}

func Succeed(ctx ResponseContext, data ...interface{}) {
	var newData interface{}
	newMessage := ""
	count := len(data)
	if count > 0 {
		newData = data[0]
	}
	if count > 1 {
		if message, ok := data[1].(string); ok {
			newMessage = message
		}
	}
	ctx.JSON(http.StatusOK, Msg(newData, newMessage, http.StatusOK))
}

func Error(ctx ResponseContext, code int, err error, data ...interface{}) {
	var newData interface{}
	if len(data) > 0 {
		newData = data[0]
	}
	ctx.AbortWithStatusJSON(http.StatusOK, Msg(newData, err.Error(), code))
}

func ErrorBadRequest(ctx ResponseContext, err error, data ...interface{}) {
	Error(ctx, http.StatusBadRequest, err, data)
}

func ErrorInternalServerError(ctx ResponseContext, err error, data ...interface{}) {
	Error(ctx, http.StatusInternalServerError, err, data)
}

func ErrorUnauthorized(ctx ResponseContext, err error, data ...interface{}) {
	Error(ctx, http.StatusUnauthorized, err, data)
}

func ErrorNotFound(ctx ResponseContext, err error, data ...interface{}) {
	Error(ctx, http.StatusNotFound, err, data)
}
