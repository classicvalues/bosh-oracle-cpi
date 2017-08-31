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

// CreateInternetGatewayReader is a Reader for the CreateInternetGateway structure.
type CreateInternetGatewayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInternetGatewayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateInternetGatewayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateInternetGatewayBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateInternetGatewayUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateInternetGatewayNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateInternetGatewayConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateInternetGatewayInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateInternetGatewayDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateInternetGatewayOK creates a CreateInternetGatewayOK with default headers values
func NewCreateInternetGatewayOK() *CreateInternetGatewayOK {
	return &CreateInternetGatewayOK{}
}

/*CreateInternetGatewayOK handles this case with default header values.

The Internet Gateway was created.
*/
type CreateInternetGatewayOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.InternetGateway
}

func (o *CreateInternetGatewayOK) Error() string {
	return fmt.Sprintf("[POST /internetGateways][%d] createInternetGatewayOK  %+v", 200, o.Payload)
}

func (o *CreateInternetGatewayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.InternetGateway)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternetGatewayBadRequest creates a CreateInternetGatewayBadRequest with default headers values
func NewCreateInternetGatewayBadRequest() *CreateInternetGatewayBadRequest {
	return &CreateInternetGatewayBadRequest{}
}

/*CreateInternetGatewayBadRequest handles this case with default header values.

Bad Request
*/
type CreateInternetGatewayBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateInternetGatewayBadRequest) Error() string {
	return fmt.Sprintf("[POST /internetGateways][%d] createInternetGatewayBadRequest  %+v", 400, o.Payload)
}

func (o *CreateInternetGatewayBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternetGatewayUnauthorized creates a CreateInternetGatewayUnauthorized with default headers values
func NewCreateInternetGatewayUnauthorized() *CreateInternetGatewayUnauthorized {
	return &CreateInternetGatewayUnauthorized{}
}

/*CreateInternetGatewayUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateInternetGatewayUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateInternetGatewayUnauthorized) Error() string {
	return fmt.Sprintf("[POST /internetGateways][%d] createInternetGatewayUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateInternetGatewayUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternetGatewayNotFound creates a CreateInternetGatewayNotFound with default headers values
func NewCreateInternetGatewayNotFound() *CreateInternetGatewayNotFound {
	return &CreateInternetGatewayNotFound{}
}

/*CreateInternetGatewayNotFound handles this case with default header values.

Not Found
*/
type CreateInternetGatewayNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateInternetGatewayNotFound) Error() string {
	return fmt.Sprintf("[POST /internetGateways][%d] createInternetGatewayNotFound  %+v", 404, o.Payload)
}

func (o *CreateInternetGatewayNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternetGatewayConflict creates a CreateInternetGatewayConflict with default headers values
func NewCreateInternetGatewayConflict() *CreateInternetGatewayConflict {
	return &CreateInternetGatewayConflict{}
}

/*CreateInternetGatewayConflict handles this case with default header values.

Conflict
*/
type CreateInternetGatewayConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateInternetGatewayConflict) Error() string {
	return fmt.Sprintf("[POST /internetGateways][%d] createInternetGatewayConflict  %+v", 409, o.Payload)
}

func (o *CreateInternetGatewayConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternetGatewayInternalServerError creates a CreateInternetGatewayInternalServerError with default headers values
func NewCreateInternetGatewayInternalServerError() *CreateInternetGatewayInternalServerError {
	return &CreateInternetGatewayInternalServerError{}
}

/*CreateInternetGatewayInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateInternetGatewayInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateInternetGatewayInternalServerError) Error() string {
	return fmt.Sprintf("[POST /internetGateways][%d] createInternetGatewayInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateInternetGatewayInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternetGatewayDefault creates a CreateInternetGatewayDefault with default headers values
func NewCreateInternetGatewayDefault(code int) *CreateInternetGatewayDefault {
	return &CreateInternetGatewayDefault{
		_statusCode: code,
	}
}

/*CreateInternetGatewayDefault handles this case with default header values.

An error has occurred.
*/
type CreateInternetGatewayDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the create internet gateway default response
func (o *CreateInternetGatewayDefault) Code() int {
	return o._statusCode
}

func (o *CreateInternetGatewayDefault) Error() string {
	return fmt.Sprintf("[POST /internetGateways][%d] CreateInternetGateway default  %+v", o._statusCode, o.Payload)
}

func (o *CreateInternetGatewayDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}