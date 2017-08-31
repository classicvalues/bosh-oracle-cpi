package virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/baremetal/core/models"
)

// ListInternetGatewaysReader is a Reader for the ListInternetGateways structure.
type ListInternetGatewaysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListInternetGatewaysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListInternetGatewaysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListInternetGatewaysUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListInternetGatewaysNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListInternetGatewaysInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListInternetGatewaysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListInternetGatewaysOK creates a ListInternetGatewaysOK with default headers values
func NewListInternetGatewaysOK() *ListInternetGatewaysOK {
	return &ListInternetGatewaysOK{}
}

/*ListInternetGatewaysOK handles this case with default header values.

The list is being retrieved.
*/
type ListInternetGatewaysOK struct {
	/*For pagination of a list of items. When paging through a list, if this header appears in the response,
	then a partial list might have been returned. Include this value as the `page` parameter for the
	subsequent GET request to get the next batch of items.

	*/
	OpcNextPage string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload []*models.InternetGateway
}

func (o *ListInternetGatewaysOK) Error() string {
	return fmt.Sprintf("[GET /internetGateways][%d] listInternetGatewaysOK  %+v", 200, o.Payload)
}

func (o *ListInternetGatewaysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-next-page
	o.OpcNextPage = response.GetHeader("opc-next-page")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListInternetGatewaysUnauthorized creates a ListInternetGatewaysUnauthorized with default headers values
func NewListInternetGatewaysUnauthorized() *ListInternetGatewaysUnauthorized {
	return &ListInternetGatewaysUnauthorized{}
}

/*ListInternetGatewaysUnauthorized handles this case with default header values.

Unauthorized
*/
type ListInternetGatewaysUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListInternetGatewaysUnauthorized) Error() string {
	return fmt.Sprintf("[GET /internetGateways][%d] listInternetGatewaysUnauthorized  %+v", 401, o.Payload)
}

func (o *ListInternetGatewaysUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListInternetGatewaysNotFound creates a ListInternetGatewaysNotFound with default headers values
func NewListInternetGatewaysNotFound() *ListInternetGatewaysNotFound {
	return &ListInternetGatewaysNotFound{}
}

/*ListInternetGatewaysNotFound handles this case with default header values.

Not Found
*/
type ListInternetGatewaysNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListInternetGatewaysNotFound) Error() string {
	return fmt.Sprintf("[GET /internetGateways][%d] listInternetGatewaysNotFound  %+v", 404, o.Payload)
}

func (o *ListInternetGatewaysNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListInternetGatewaysInternalServerError creates a ListInternetGatewaysInternalServerError with default headers values
func NewListInternetGatewaysInternalServerError() *ListInternetGatewaysInternalServerError {
	return &ListInternetGatewaysInternalServerError{}
}

/*ListInternetGatewaysInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListInternetGatewaysInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListInternetGatewaysInternalServerError) Error() string {
	return fmt.Sprintf("[GET /internetGateways][%d] listInternetGatewaysInternalServerError  %+v", 500, o.Payload)
}

func (o *ListInternetGatewaysInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListInternetGatewaysDefault creates a ListInternetGatewaysDefault with default headers values
func NewListInternetGatewaysDefault(code int) *ListInternetGatewaysDefault {
	return &ListInternetGatewaysDefault{
		_statusCode: code,
	}
}

/*ListInternetGatewaysDefault handles this case with default header values.

An error has occurred.
*/
type ListInternetGatewaysDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the list internet gateways default response
func (o *ListInternetGatewaysDefault) Code() int {
	return o._statusCode
}

func (o *ListInternetGatewaysDefault) Error() string {
	return fmt.Sprintf("[GET /internetGateways][%d] ListInternetGateways default  %+v", o._statusCode, o.Payload)
}

func (o *ListInternetGatewaysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}