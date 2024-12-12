package verify

import (
	"net/http"
	"reflect"

	"esdumpweb/constant"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

func validater[BodyType any](c *gin.Context) (BodyType, error) {
	var (
		body BodyType
		err  error
	)
	if c.Request.Method == http.MethodGet {
		err = c.ShouldBindQuery(&body)
	} else {
		// Cannot use c.ShouldBind() or c.Bind() cause they only unmarshal and not set the body into the context.
		// Use c.ShouldBindBodyWith() instead of them.
		err = c.ShouldBindBodyWith(&body, BindingDefault(c.ContentType()))
	}
	if err != nil {
		handleErrMsg(c, err, body)
		return body, err
	}
	return body, nil
}

func Validate[BodyType any](c *gin.Context) {
	body, err := validater[BodyType](c)
	if err != nil {
		return
	}
	c.Set(constant.CtxKeyReqParams, body)
	c.Next()
}

func ValidateX[BodyType any](opts ...func(*BodyType) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := validater[BodyType](c)
		if err != nil {
			return
		}
		for _, f := range opts {
			if err = f(&body); err != nil {
				handleErrMsg(c, err, body)
				return
			}
		}
		c.Set(constant.CtxKeyReqParams, body)
		c.Next()
	}
}

func handleErrMsg[BodyType any](c *gin.Context, err error, body BodyType) {
	if errors.Is(err, &validator.InvalidValidationError{}) {
		c.Status(http.StatusUnprocessableEntity)
		_ = c.Error(errors.WithStack(err))
		c.Abort()
		return
	}

	validationErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.Status(http.StatusUnprocessableEntity)
		_ = c.Error(errors.WithStack(err))
		c.Abort()
		return
	}

	for _, validationErr := range validationErrs {
		v := reflect.ValueOf(body)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		field, ok := v.Type().FieldByName(validationErr.Field())
		if !ok {
			continue
		}
		prompt := field.Tag.Get("prompt") // get prompt info from tag
		if prompt == "" {
			continue
		}
		errMsg := prompt
		c.Status(http.StatusUnprocessableEntity)
		_ = c.Error(errors.WithMessage(errors.WithStack(err), errMsg))
		c.Abort()
		break
	}
}

func GetReqParams[BodyType any](c *gin.Context) BodyType {
	return c.MustGet(constant.CtxKeyReqParams).(BodyType)
}
