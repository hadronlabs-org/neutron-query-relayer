// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NeutronInterchainadapterInterchainqueriesBlock neutron interchainadapter interchainqueries block
//
// swagger:model neutron.interchainadapter.interchainqueries.Block
type NeutronInterchainadapterInterchainqueriesBlock struct {

	// header
	Header *NeutronInterchainadapterInterchainqueriesBlockHeader `json:"header,omitempty"`

	// next block header
	NextBlockHeader *NeutronInterchainadapterInterchainqueriesBlockNextBlockHeader `json:"next_block_header,omitempty"`

	// tx
	Tx *NeutronInterchainadapterInterchainqueriesBlockTx `json:"tx,omitempty"`
}

// Validate validates this neutron interchainadapter interchainqueries block
func (m *NeutronInterchainadapterInterchainqueriesBlock) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeader(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextBlockHeader(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTx(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NeutronInterchainadapterInterchainqueriesBlock) validateHeader(formats strfmt.Registry) error {
	if swag.IsZero(m.Header) { // not required
		return nil
	}

	if m.Header != nil {
		if err := m.Header.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header")
			}
			return err
		}
	}

	return nil
}

func (m *NeutronInterchainadapterInterchainqueriesBlock) validateNextBlockHeader(formats strfmt.Registry) error {
	if swag.IsZero(m.NextBlockHeader) { // not required
		return nil
	}

	if m.NextBlockHeader != nil {
		if err := m.NextBlockHeader.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next_block_header")
			}
			return err
		}
	}

	return nil
}

func (m *NeutronInterchainadapterInterchainqueriesBlock) validateTx(formats strfmt.Registry) error {
	if swag.IsZero(m.Tx) { // not required
		return nil
	}

	if m.Tx != nil {
		if err := m.Tx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tx")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this neutron interchainadapter interchainqueries block based on the context it is used
func (m *NeutronInterchainadapterInterchainqueriesBlock) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHeader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNextBlockHeader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTx(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NeutronInterchainadapterInterchainqueriesBlock) contextValidateHeader(ctx context.Context, formats strfmt.Registry) error {

	if m.Header != nil {
		if err := m.Header.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header")
			}
			return err
		}
	}

	return nil
}

func (m *NeutronInterchainadapterInterchainqueriesBlock) contextValidateNextBlockHeader(ctx context.Context, formats strfmt.Registry) error {

	if m.NextBlockHeader != nil {
		if err := m.NextBlockHeader.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next_block_header")
			}
			return err
		}
	}

	return nil
}

func (m *NeutronInterchainadapterInterchainqueriesBlock) contextValidateTx(ctx context.Context, formats strfmt.Registry) error {

	if m.Tx != nil {
		if err := m.Tx.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tx")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchainqueriesBlock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchainqueriesBlock) UnmarshalBinary(b []byte) error {
	var res NeutronInterchainadapterInterchainqueriesBlock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NeutronInterchainadapterInterchainqueriesBlockHeader We need to know block X to verify inclusion of transaction for block X
//
// `Any` contains an arbitrary serialized protocol buffer message along with a
// URL that describes the type of the serialized message.
//
// Protobuf library provides support to pack/unpack Any values in the form
// of utility functions or additional generated methods of the Any type.
//
// Example 1: Pack and unpack a message in C++.
//
//     Foo foo = ...;
//     Any any;
//     any.PackFrom(foo);
//     ...
//     if (any.UnpackTo(&foo)) {
//       ...
//     }
//
// Example 2: Pack and unpack a message in Java.
//
//     Foo foo = ...;
//     Any any = Any.pack(foo);
//     ...
//     if (any.is(Foo.class)) {
//       foo = any.unpack(Foo.class);
//     }
//
//  Example 3: Pack and unpack a message in Python.
//
//     foo = Foo(...)
//     any = Any()
//     any.Pack(foo)
//     ...
//     if any.Is(Foo.DESCRIPTOR):
//       any.Unpack(foo)
//       ...
//
//  Example 4: Pack and unpack a message in Go
//
//      foo := &pb.Foo{...}
//      any, err := anypb.New(foo)
//      if err != nil {
//        ...
//      }
//      ...
//      foo := &pb.Foo{}
//      if err := any.UnmarshalTo(foo); err != nil {
//        ...
//      }
//
// The pack methods provided by protobuf library will by default use
// 'type.googleapis.com/full.type.name' as the type URL and the unpack
// methods only use the fully qualified type name after the last '/'
// in the type URL, for example "foo.bar.com/x/y.z" will yield type
// name "y.z".
//
//
// JSON
// ====
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//     package google.profile;
//     message Person {
//       string first_name = 1;
//       string last_name = 2;
//     }
//
//     {
//       "@type": "type.googleapis.com/google.profile.Person",
//       "firstName": <string>,
//       "lastName": <string>
//     }
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the `@type`
// field. Example (for message [google.protobuf.Duration][]):
//
//     {
//       "@type": "type.googleapis.com/google.protobuf.Duration",
//       "value": "1.212s"
//     }
//
// swagger:model NeutronInterchainadapterInterchainqueriesBlockHeader
type NeutronInterchainadapterInterchainqueriesBlockHeader struct {

	// A URL/resource name that uniquely identifies the type of the serialized
	// protocol buffer message. This string must contain at least
	// one "/" character. The last segment of the URL's path must represent
	// the fully qualified name of the type (as in
	// `path/google.protobuf.Duration`). The name should be in a canonical form
	// (e.g., leading "." is not accepted).
	//
	// In practice, teams usually precompile into the binary all types that they
	// expect it to use in the context of Any. However, for URLs which use the
	// scheme `http`, `https`, or no scheme, one can optionally set up a type
	// server that maps type URLs to message definitions as follows:
	//
	// * If no scheme is provided, `https` is assumed.
	// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//   value in binary format, or produce an error.
	// * Applications are allowed to cache lookup results based on the
	//   URL, or have them precompiled into a binary to avoid any
	//   lookup. Therefore, binary compatibility needs to be preserved
	//   on changes to types. (Use versioned type names to manage
	//   breaking changes.)
	//
	// Note: this functionality is not currently available in the official
	// protobuf release, and it is not used for type URLs beginning with
	// type.googleapis.com.
	//
	// Schemes other than `http`, `https` (or the empty scheme) might be
	// used with implementation specific semantics.
	AtType string `json:"@type,omitempty"`
}

// Validate validates this neutron interchainadapter interchainqueries block header
func (m *NeutronInterchainadapterInterchainqueriesBlockHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this neutron interchainadapter interchainqueries block header based on context it is used
func (m *NeutronInterchainadapterInterchainqueriesBlockHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchainqueriesBlockHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchainqueriesBlockHeader) UnmarshalBinary(b []byte) error {
	var res NeutronInterchainadapterInterchainqueriesBlockHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NeutronInterchainadapterInterchainqueriesBlockNextBlockHeader We need to know block X+1 to verify response of transaction for block X
// since LastResultsHash is root hash of all results from the txs from the previous block
//
// `Any` contains an arbitrary serialized protocol buffer message along with a
// URL that describes the type of the serialized message.
//
// Protobuf library provides support to pack/unpack Any values in the form
// of utility functions or additional generated methods of the Any type.
//
// Example 1: Pack and unpack a message in C++.
//
//     Foo foo = ...;
//     Any any;
//     any.PackFrom(foo);
//     ...
//     if (any.UnpackTo(&foo)) {
//       ...
//     }
//
// Example 2: Pack and unpack a message in Java.
//
//     Foo foo = ...;
//     Any any = Any.pack(foo);
//     ...
//     if (any.is(Foo.class)) {
//       foo = any.unpack(Foo.class);
//     }
//
//  Example 3: Pack and unpack a message in Python.
//
//     foo = Foo(...)
//     any = Any()
//     any.Pack(foo)
//     ...
//     if any.Is(Foo.DESCRIPTOR):
//       any.Unpack(foo)
//       ...
//
//  Example 4: Pack and unpack a message in Go
//
//      foo := &pb.Foo{...}
//      any, err := anypb.New(foo)
//      if err != nil {
//        ...
//      }
//      ...
//      foo := &pb.Foo{}
//      if err := any.UnmarshalTo(foo); err != nil {
//        ...
//      }
//
// The pack methods provided by protobuf library will by default use
// 'type.googleapis.com/full.type.name' as the type URL and the unpack
// methods only use the fully qualified type name after the last '/'
// in the type URL, for example "foo.bar.com/x/y.z" will yield type
// name "y.z".
//
//
// JSON
// ====
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//     package google.profile;
//     message Person {
//       string first_name = 1;
//       string last_name = 2;
//     }
//
//     {
//       "@type": "type.googleapis.com/google.profile.Person",
//       "firstName": <string>,
//       "lastName": <string>
//     }
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the `@type`
// field. Example (for message [google.protobuf.Duration][]):
//
//     {
//       "@type": "type.googleapis.com/google.protobuf.Duration",
//       "value": "1.212s"
//     }
//
// swagger:model NeutronInterchainadapterInterchainqueriesBlockNextBlockHeader
type NeutronInterchainadapterInterchainqueriesBlockNextBlockHeader struct {

	// A URL/resource name that uniquely identifies the type of the serialized
	// protocol buffer message. This string must contain at least
	// one "/" character. The last segment of the URL's path must represent
	// the fully qualified name of the type (as in
	// `path/google.protobuf.Duration`). The name should be in a canonical form
	// (e.g., leading "." is not accepted).
	//
	// In practice, teams usually precompile into the binary all types that they
	// expect it to use in the context of Any. However, for URLs which use the
	// scheme `http`, `https`, or no scheme, one can optionally set up a type
	// server that maps type URLs to message definitions as follows:
	//
	// * If no scheme is provided, `https` is assumed.
	// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//   value in binary format, or produce an error.
	// * Applications are allowed to cache lookup results based on the
	//   URL, or have them precompiled into a binary to avoid any
	//   lookup. Therefore, binary compatibility needs to be preserved
	//   on changes to types. (Use versioned type names to manage
	//   breaking changes.)
	//
	// Note: this functionality is not currently available in the official
	// protobuf release, and it is not used for type URLs beginning with
	// type.googleapis.com.
	//
	// Schemes other than `http`, `https` (or the empty scheme) might be
	// used with implementation specific semantics.
	AtType string `json:"@type,omitempty"`
}

// Validate validates this neutron interchainadapter interchainqueries block next block header
func (m *NeutronInterchainadapterInterchainqueriesBlockNextBlockHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this neutron interchainadapter interchainqueries block next block header based on context it is used
func (m *NeutronInterchainadapterInterchainqueriesBlockNextBlockHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchainqueriesBlockNextBlockHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchainqueriesBlockNextBlockHeader) UnmarshalBinary(b []byte) error {
	var res NeutronInterchainadapterInterchainqueriesBlockNextBlockHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NeutronInterchainadapterInterchainqueriesBlockTx neutron interchainadapter interchainqueries block tx
//
// swagger:model NeutronInterchainadapterInterchainqueriesBlockTx
type NeutronInterchainadapterInterchainqueriesBlockTx struct {

	// is body of the transaction
	// Format: byte
	Data strfmt.Base64 `json:"data,omitempty"`

	// delivery proof
	DeliveryProof *NeutronInterchainadapterInterchainqueriesBlockTxDeliveryProof `json:"delivery_proof,omitempty"`

	// inclusion proof
	InclusionProof *NeutronInterchainadapterInterchainqueriesBlockTxInclusionProof `json:"inclusion_proof,omitempty"`

	// response
	Response *NeutronInterchainadapterInterchainqueriesBlockTxResponse `json:"response,omitempty"`
}

// Validate validates this neutron interchainadapter interchainqueries block tx
func (m *NeutronInterchainadapterInterchainqueriesBlockTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeliveryProof(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInclusionProof(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NeutronInterchainadapterInterchainqueriesBlockTx) validateDeliveryProof(formats strfmt.Registry) error {
	if swag.IsZero(m.DeliveryProof) { // not required
		return nil
	}

	if m.DeliveryProof != nil {
		if err := m.DeliveryProof.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tx" + "." + "delivery_proof")
			}
			return err
		}
	}

	return nil
}

func (m *NeutronInterchainadapterInterchainqueriesBlockTx) validateInclusionProof(formats strfmt.Registry) error {
	if swag.IsZero(m.InclusionProof) { // not required
		return nil
	}

	if m.InclusionProof != nil {
		if err := m.InclusionProof.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tx" + "." + "inclusion_proof")
			}
			return err
		}
	}

	return nil
}

func (m *NeutronInterchainadapterInterchainqueriesBlockTx) validateResponse(formats strfmt.Registry) error {
	if swag.IsZero(m.Response) { // not required
		return nil
	}

	if m.Response != nil {
		if err := m.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tx" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this neutron interchainadapter interchainqueries block tx based on the context it is used
func (m *NeutronInterchainadapterInterchainqueriesBlockTx) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeliveryProof(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInclusionProof(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NeutronInterchainadapterInterchainqueriesBlockTx) contextValidateDeliveryProof(ctx context.Context, formats strfmt.Registry) error {

	if m.DeliveryProof != nil {
		if err := m.DeliveryProof.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tx" + "." + "delivery_proof")
			}
			return err
		}
	}

	return nil
}

func (m *NeutronInterchainadapterInterchainqueriesBlockTx) contextValidateInclusionProof(ctx context.Context, formats strfmt.Registry) error {

	if m.InclusionProof != nil {
		if err := m.InclusionProof.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tx" + "." + "inclusion_proof")
			}
			return err
		}
	}

	return nil
}

func (m *NeutronInterchainadapterInterchainqueriesBlockTx) contextValidateResponse(ctx context.Context, formats strfmt.Registry) error {

	if m.Response != nil {
		if err := m.Response.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tx" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchainqueriesBlockTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchainqueriesBlockTx) UnmarshalBinary(b []byte) error {
	var res NeutronInterchainadapterInterchainqueriesBlockTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NeutronInterchainadapterInterchainqueriesBlockTxDeliveryProof is the Merkle Proof which proves existence of response in block with height next_block_header.Height
//
// swagger:model NeutronInterchainadapterInterchainqueriesBlockTxDeliveryProof
type NeutronInterchainadapterInterchainqueriesBlockTxDeliveryProof struct {

	// aunts
	Aunts []strfmt.Base64 `json:"aunts"`

	// index
	Index string `json:"index,omitempty"`

	// leaf hash
	// Format: byte
	LeafHash strfmt.Base64 `json:"leaf_hash,omitempty"`

	// total
	Total string `json:"total,omitempty"`
}

// Validate validates this neutron interchainadapter interchainqueries block tx delivery proof
func (m *NeutronInterchainadapterInterchainqueriesBlockTxDeliveryProof) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this neutron interchainadapter interchainqueries block tx delivery proof based on context it is used
func (m *NeutronInterchainadapterInterchainqueriesBlockTxDeliveryProof) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchainqueriesBlockTxDeliveryProof) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchainqueriesBlockTxDeliveryProof) UnmarshalBinary(b []byte) error {
	var res NeutronInterchainadapterInterchainqueriesBlockTxDeliveryProof
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NeutronInterchainadapterInterchainqueriesBlockTxInclusionProof is the Merkle Proof which proves existence of data in block with height header.Height
//
// swagger:model NeutronInterchainadapterInterchainqueriesBlockTxInclusionProof
type NeutronInterchainadapterInterchainqueriesBlockTxInclusionProof struct {

	// aunts
	Aunts []strfmt.Base64 `json:"aunts"`

	// index
	Index string `json:"index,omitempty"`

	// leaf hash
	// Format: byte
	LeafHash strfmt.Base64 `json:"leaf_hash,omitempty"`

	// total
	Total string `json:"total,omitempty"`
}

// Validate validates this neutron interchainadapter interchainqueries block tx inclusion proof
func (m *NeutronInterchainadapterInterchainqueriesBlockTxInclusionProof) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this neutron interchainadapter interchainqueries block tx inclusion proof based on context it is used
func (m *NeutronInterchainadapterInterchainqueriesBlockTxInclusionProof) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchainqueriesBlockTxInclusionProof) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchainqueriesBlockTxInclusionProof) UnmarshalBinary(b []byte) error {
	var res NeutronInterchainadapterInterchainqueriesBlockTxInclusionProof
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NeutronInterchainadapterInterchainqueriesBlockTxResponse neutron interchainadapter interchainqueries block tx response
//
// swagger:model NeutronInterchainadapterInterchainqueriesBlockTxResponse
type NeutronInterchainadapterInterchainqueriesBlockTxResponse struct {

	// code
	Code int64 `json:"code,omitempty"`

	// codespace
	Codespace string `json:"codespace,omitempty"`

	// data
	// Format: byte
	Data strfmt.Base64 `json:"data,omitempty"`

	// events
	Events []*NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0 `json:"events"`

	// gas used
	GasUsed string `json:"gas_used,omitempty"`

	// gas wanted
	GasWanted string `json:"gas_wanted,omitempty"`

	// info
	Info string `json:"info,omitempty"`

	// log
	Log string `json:"log,omitempty"`
}

// Validate validates this neutron interchainadapter interchainqueries block tx response
func (m *NeutronInterchainadapterInterchainqueriesBlockTxResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NeutronInterchainadapterInterchainqueriesBlockTxResponse) validateEvents(formats strfmt.Registry) error {
	if swag.IsZero(m.Events) { // not required
		return nil
	}

	for i := 0; i < len(m.Events); i++ {
		if swag.IsZero(m.Events[i]) { // not required
			continue
		}

		if m.Events[i] != nil {
			if err := m.Events[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tx" + "." + "response" + "." + "events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this neutron interchainadapter interchainqueries block tx response based on the context it is used
func (m *NeutronInterchainadapterInterchainqueriesBlockTxResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NeutronInterchainadapterInterchainqueriesBlockTxResponse) contextValidateEvents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Events); i++ {

		if m.Events[i] != nil {
			if err := m.Events[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tx" + "." + "response" + "." + "events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchainqueriesBlockTxResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchainqueriesBlockTxResponse) UnmarshalBinary(b []byte) error {
	var res NeutronInterchainadapterInterchainqueriesBlockTxResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0 Event allows application developers to attach additional information to
// ResponseBeginBlock, ResponseEndBlock, ResponseCheckTx and ResponseDeliverTx.
// Later, transactions may be queried using these events.
//
// swagger:model NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0
type NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0 struct {

	// attributes
	Attributes []*NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0AttributesItems0 `json:"attributes"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this neutron interchainadapter interchainqueries block tx response events items0
func (m *NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0) validateAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	for i := 0; i < len(m.Attributes); i++ {
		if swag.IsZero(m.Attributes[i]) { // not required
			continue
		}

		if m.Attributes[i] != nil {
			if err := m.Attributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this neutron interchainadapter interchainqueries block tx response events items0 based on the context it is used
func (m *NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attributes); i++ {

		if m.Attributes[i] != nil {
			if err := m.Attributes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0) UnmarshalBinary(b []byte) error {
	var res NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0AttributesItems0 EventAttribute is a single key-value pair, associated with an event.
//
// swagger:model NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0AttributesItems0
type NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0AttributesItems0 struct {

	// index
	Index bool `json:"index,omitempty"`

	// key
	// Format: byte
	Key strfmt.Base64 `json:"key,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this neutron interchainadapter interchainqueries block tx response events items0 attributes items0
func (m *NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0AttributesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this neutron interchainadapter interchainqueries block tx response events items0 attributes items0 based on context it is used
func (m *NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0AttributesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0AttributesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0AttributesItems0) UnmarshalBinary(b []byte) error {
	var res NeutronInterchainadapterInterchainqueriesBlockTxResponseEventsItems0AttributesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
