package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/baremetal/core/models"
)

// GetVolumeAttachmentReader is a Reader for the GetVolumeAttachment structure.
type GetVolumeAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVolumeAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVolumeAttachmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetVolumeAttachmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetVolumeAttachmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetVolumeAttachmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetVolumeAttachmentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVolumeAttachmentOK creates a GetVolumeAttachmentOK with default headers values
func NewGetVolumeAttachmentOK() *GetVolumeAttachmentOK {
	return &GetVolumeAttachmentOK{}
}

/*GetVolumeAttachmentOK handles this case with default header values.

The volume attachment was retrieved.
*/
type GetVolumeAttachmentOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload models.VolumeAttachment
}

func (o *GetVolumeAttachmentOK) Error() string {
	return fmt.Sprintf("[GET /volumeAttachments/{volumeAttachmentId}][%d] getVolumeAttachmentOK  %+v", 200, o.Payload)
}

func (o *GetVolumeAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	// response payload as interface type
	payload, err := models.UnmarshalVolumeAttachment(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetVolumeAttachmentUnauthorized creates a GetVolumeAttachmentUnauthorized with default headers values
func NewGetVolumeAttachmentUnauthorized() *GetVolumeAttachmentUnauthorized {
	return &GetVolumeAttachmentUnauthorized{}
}

/*GetVolumeAttachmentUnauthorized handles this case with default header values.

Unauthorized
*/
type GetVolumeAttachmentUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetVolumeAttachmentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /volumeAttachments/{volumeAttachmentId}][%d] getVolumeAttachmentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVolumeAttachmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumeAttachmentNotFound creates a GetVolumeAttachmentNotFound with default headers values
func NewGetVolumeAttachmentNotFound() *GetVolumeAttachmentNotFound {
	return &GetVolumeAttachmentNotFound{}
}

/*GetVolumeAttachmentNotFound handles this case with default header values.

Not Found
*/
type GetVolumeAttachmentNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetVolumeAttachmentNotFound) Error() string {
	return fmt.Sprintf("[GET /volumeAttachments/{volumeAttachmentId}][%d] getVolumeAttachmentNotFound  %+v", 404, o.Payload)
}

func (o *GetVolumeAttachmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumeAttachmentInternalServerError creates a GetVolumeAttachmentInternalServerError with default headers values
func NewGetVolumeAttachmentInternalServerError() *GetVolumeAttachmentInternalServerError {
	return &GetVolumeAttachmentInternalServerError{}
}

/*GetVolumeAttachmentInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetVolumeAttachmentInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetVolumeAttachmentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /volumeAttachments/{volumeAttachmentId}][%d] getVolumeAttachmentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVolumeAttachmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumeAttachmentDefault creates a GetVolumeAttachmentDefault with default headers values
func NewGetVolumeAttachmentDefault(code int) *GetVolumeAttachmentDefault {
	return &GetVolumeAttachmentDefault{
		_statusCode: code,
	}
}

/*GetVolumeAttachmentDefault handles this case with default header values.

An error has occurred.
*/
type GetVolumeAttachmentDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the get volume attachment default response
func (o *GetVolumeAttachmentDefault) Code() int {
	return o._statusCode
}

func (o *GetVolumeAttachmentDefault) Error() string {
	return fmt.Sprintf("[GET /volumeAttachments/{volumeAttachmentId}][%d] GetVolumeAttachment default  %+v", o._statusCode, o.Payload)
}

func (o *GetVolumeAttachmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}