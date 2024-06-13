package validator

import (
	"bytes"
	"e-commerce-order-service/model"
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetRequestUnableToBindZwError() model.Errors {
	zwErr := model.Errors{
		Message: errors.New("request body is invalid"),
		Type:    "422, invalid_request_error",
	}
	return zwErr
}
func GetValidationZwError(e url.Values) model.Errors {
	zwErr := model.Errors{}
	for param, message := range e {
		zwErr = model.Errors{
			Param:   param,
			Message: errors.New(message[0]),
			Type:    "invalid_request_error",
			Code:    "parameter_missing",
		}
		break
	}
	return zwErr
}
func ReturnJsonStruct(c *gin.Context, genericStruct interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Set("responsestring", genericStruct)
	json.NewEncoder(c.Writer).Encode(genericStruct)
}
func ValidateUnknownParams(reqBody interface{}, ctx *gin.Context) model.Errors {
	ginErr := model.Errors{
		Error: "the request body is invalid",
		Type:  "invalid_request_error",
	}
	decoder := json.NewDecoder(ctx.Request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqBody)
	if err != nil {
		return ginErr
	}
	payloadBS, err := json.Marshal(&reqBody)
	if err != nil {
		return ginErr
	}
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(payloadBS))
	return model.Errors{}
}
func ReturnJsonUnescapedStruct(c *gin.Context, data interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(data)
	if err != nil {

		_ = json.NewEncoder(c.Writer).Encode(data)
		return
	}
	str := string(b)
	str = strings.ReplaceAll(str, "\\t", " ")
	str = strings.ReplaceAll(str, "\\r", " ")
	str = strings.ReplaceAll(str, "\\n", " ")
	str = strings.ReplaceAll(str, "\\\"", "\"")
	str = strings.ReplaceAll(str, "\\\\\"", "\"")
	str = strings.ReplaceAll(str, "\"{", "{")
	str = strings.ReplaceAll(str, "}\"", "}")
	str = strings.ReplaceAll(str, "\"[", "[")
	str = strings.ReplaceAll(str, "]\"", "]")
	str = strings.ReplaceAll(str, "]\"", "]")
	str = strings.ReplaceAll(str, "\\n\"", " ")
	str = strings.ReplaceAll(str, "\\", " ")
	str = strings.ReplaceAll(str, "  u0026", "&")

	c.Set("responsestring", str)
	_, _ = c.Writer.WriteString(str)
}

func IsValidPincode(pincode string) bool {
	regex := `^[0-9]{6}$`
	matched, _ := regexp.MatchString(regex, pincode)
	return matched
}

func IsValidDOB(dob string) bool {
	regex := `^\d{4}-\d{2}-\d{2}$`
	matched, _ := regexp.MatchString(regex, dob)
	return matched
}

func IsValidEmail(email string) bool {
	regex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	matched, _ := regexp.MatchString(regex, email)
	return matched
}
