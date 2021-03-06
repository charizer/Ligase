// Copyright (C) 2020 Finogeeks Co., Ltd
//
// This program is free software: you can redistribute it and/or  modify
// it under the terms of the GNU Affero General Public License, version 3,
// as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by capnpc-go. DO NOT EDIT.

package external

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type GetUserProfileAvatarURLRequestCapn struct{ capnp.Struct }

// GetUserProfileAvatarURLRequestCapn_TypeID is the unique identifier for the type GetUserProfileAvatarURLRequestCapn.
const GetUserProfileAvatarURLRequestCapn_TypeID = 0xc88cd0641eca2e2d

func NewGetUserProfileAvatarURLRequestCapn(s *capnp.Segment) (GetUserProfileAvatarURLRequestCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return GetUserProfileAvatarURLRequestCapn{st}, err
}

func NewRootGetUserProfileAvatarURLRequestCapn(s *capnp.Segment) (GetUserProfileAvatarURLRequestCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return GetUserProfileAvatarURLRequestCapn{st}, err
}

func ReadRootGetUserProfileAvatarURLRequestCapn(msg *capnp.Message) (GetUserProfileAvatarURLRequestCapn, error) {
	root, err := msg.RootPtr()
	return GetUserProfileAvatarURLRequestCapn{root.Struct()}, err
}

func (s GetUserProfileAvatarURLRequestCapn) String() string {
	str, _ := text.Marshal(0xc88cd0641eca2e2d, s.Struct)
	return str
}

func (s GetUserProfileAvatarURLRequestCapn) UserID() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s GetUserProfileAvatarURLRequestCapn) HasUserID() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s GetUserProfileAvatarURLRequestCapn) UserIDBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s GetUserProfileAvatarURLRequestCapn) SetUserID(v string) error {
	return s.Struct.SetText(0, v)
}

// GetUserProfileAvatarURLRequestCapn_List is a list of GetUserProfileAvatarURLRequestCapn.
type GetUserProfileAvatarURLRequestCapn_List struct{ capnp.List }

// NewGetUserProfileAvatarURLRequestCapn creates a new list of GetUserProfileAvatarURLRequestCapn.
func NewGetUserProfileAvatarURLRequestCapn_List(s *capnp.Segment, sz int32) (GetUserProfileAvatarURLRequestCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return GetUserProfileAvatarURLRequestCapn_List{l}, err
}

func (s GetUserProfileAvatarURLRequestCapn_List) At(i int) GetUserProfileAvatarURLRequestCapn {
	return GetUserProfileAvatarURLRequestCapn{s.List.Struct(i)}
}

func (s GetUserProfileAvatarURLRequestCapn_List) Set(i int, v GetUserProfileAvatarURLRequestCapn) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s GetUserProfileAvatarURLRequestCapn_List) String() string {
	str, _ := text.MarshalList(0xc88cd0641eca2e2d, s.List)
	return str
}

// GetUserProfileAvatarURLRequestCapn_Promise is a wrapper for a GetUserProfileAvatarURLRequestCapn promised by a client call.
type GetUserProfileAvatarURLRequestCapn_Promise struct{ *capnp.Pipeline }

func (p GetUserProfileAvatarURLRequestCapn_Promise) Struct() (GetUserProfileAvatarURLRequestCapn, error) {
	s, err := p.Pipeline.Struct()
	return GetUserProfileAvatarURLRequestCapn{s}, err
}

type GetUserProfileDisplayNameRequestCapn struct{ capnp.Struct }

// GetUserProfileDisplayNameRequestCapn_TypeID is the unique identifier for the type GetUserProfileDisplayNameRequestCapn.
const GetUserProfileDisplayNameRequestCapn_TypeID = 0xfd6e19082b99c691

func NewGetUserProfileDisplayNameRequestCapn(s *capnp.Segment) (GetUserProfileDisplayNameRequestCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return GetUserProfileDisplayNameRequestCapn{st}, err
}

func NewRootGetUserProfileDisplayNameRequestCapn(s *capnp.Segment) (GetUserProfileDisplayNameRequestCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return GetUserProfileDisplayNameRequestCapn{st}, err
}

func ReadRootGetUserProfileDisplayNameRequestCapn(msg *capnp.Message) (GetUserProfileDisplayNameRequestCapn, error) {
	root, err := msg.RootPtr()
	return GetUserProfileDisplayNameRequestCapn{root.Struct()}, err
}

func (s GetUserProfileDisplayNameRequestCapn) String() string {
	str, _ := text.Marshal(0xfd6e19082b99c691, s.Struct)
	return str
}

func (s GetUserProfileDisplayNameRequestCapn) UserID() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s GetUserProfileDisplayNameRequestCapn) HasUserID() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s GetUserProfileDisplayNameRequestCapn) UserIDBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s GetUserProfileDisplayNameRequestCapn) SetUserID(v string) error {
	return s.Struct.SetText(0, v)
}

// GetUserProfileDisplayNameRequestCapn_List is a list of GetUserProfileDisplayNameRequestCapn.
type GetUserProfileDisplayNameRequestCapn_List struct{ capnp.List }

// NewGetUserProfileDisplayNameRequestCapn creates a new list of GetUserProfileDisplayNameRequestCapn.
func NewGetUserProfileDisplayNameRequestCapn_List(s *capnp.Segment, sz int32) (GetUserProfileDisplayNameRequestCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return GetUserProfileDisplayNameRequestCapn_List{l}, err
}

func (s GetUserProfileDisplayNameRequestCapn_List) At(i int) GetUserProfileDisplayNameRequestCapn {
	return GetUserProfileDisplayNameRequestCapn{s.List.Struct(i)}
}

func (s GetUserProfileDisplayNameRequestCapn_List) Set(i int, v GetUserProfileDisplayNameRequestCapn) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s GetUserProfileDisplayNameRequestCapn_List) String() string {
	str, _ := text.MarshalList(0xfd6e19082b99c691, s.List)
	return str
}

// GetUserProfileDisplayNameRequestCapn_Promise is a wrapper for a GetUserProfileDisplayNameRequestCapn promised by a client call.
type GetUserProfileDisplayNameRequestCapn_Promise struct{ *capnp.Pipeline }

func (p GetUserProfileDisplayNameRequestCapn_Promise) Struct() (GetUserProfileDisplayNameRequestCapn, error) {
	s, err := p.Pipeline.Struct()
	return GetUserProfileDisplayNameRequestCapn{s}, err
}

type GetUserProfileRequestCapn struct{ capnp.Struct }

// GetUserProfileRequestCapn_TypeID is the unique identifier for the type GetUserProfileRequestCapn.
const GetUserProfileRequestCapn_TypeID = 0xd0cf8d14918db5d7

func NewGetUserProfileRequestCapn(s *capnp.Segment) (GetUserProfileRequestCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return GetUserProfileRequestCapn{st}, err
}

func NewRootGetUserProfileRequestCapn(s *capnp.Segment) (GetUserProfileRequestCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return GetUserProfileRequestCapn{st}, err
}

func ReadRootGetUserProfileRequestCapn(msg *capnp.Message) (GetUserProfileRequestCapn, error) {
	root, err := msg.RootPtr()
	return GetUserProfileRequestCapn{root.Struct()}, err
}

func (s GetUserProfileRequestCapn) String() string {
	str, _ := text.Marshal(0xd0cf8d14918db5d7, s.Struct)
	return str
}

func (s GetUserProfileRequestCapn) UserID() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s GetUserProfileRequestCapn) HasUserID() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s GetUserProfileRequestCapn) UserIDBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s GetUserProfileRequestCapn) SetUserID(v string) error {
	return s.Struct.SetText(0, v)
}

// GetUserProfileRequestCapn_List is a list of GetUserProfileRequestCapn.
type GetUserProfileRequestCapn_List struct{ capnp.List }

// NewGetUserProfileRequestCapn creates a new list of GetUserProfileRequestCapn.
func NewGetUserProfileRequestCapn_List(s *capnp.Segment, sz int32) (GetUserProfileRequestCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return GetUserProfileRequestCapn_List{l}, err
}

func (s GetUserProfileRequestCapn_List) At(i int) GetUserProfileRequestCapn {
	return GetUserProfileRequestCapn{s.List.Struct(i)}
}

func (s GetUserProfileRequestCapn_List) Set(i int, v GetUserProfileRequestCapn) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s GetUserProfileRequestCapn_List) String() string {
	str, _ := text.MarshalList(0xd0cf8d14918db5d7, s.List)
	return str
}

// GetUserProfileRequestCapn_Promise is a wrapper for a GetUserProfileRequestCapn promised by a client call.
type GetUserProfileRequestCapn_Promise struct{ *capnp.Pipeline }

func (p GetUserProfileRequestCapn_Promise) Struct() (GetUserProfileRequestCapn, error) {
	s, err := p.Pipeline.Struct()
	return GetUserProfileRequestCapn{s}, err
}

type PostUserSearchRequestCapn struct{ capnp.Struct }

// PostUserSearchRequestCapn_TypeID is the unique identifier for the type PostUserSearchRequestCapn.
const PostUserSearchRequestCapn_TypeID = 0x96745109c202f5a6

func NewPostUserSearchRequestCapn(s *capnp.Segment) (PostUserSearchRequestCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PostUserSearchRequestCapn{st}, err
}

func NewRootPostUserSearchRequestCapn(s *capnp.Segment) (PostUserSearchRequestCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PostUserSearchRequestCapn{st}, err
}

func ReadRootPostUserSearchRequestCapn(msg *capnp.Message) (PostUserSearchRequestCapn, error) {
	root, err := msg.RootPtr()
	return PostUserSearchRequestCapn{root.Struct()}, err
}

func (s PostUserSearchRequestCapn) String() string {
	str, _ := text.Marshal(0x96745109c202f5a6, s.Struct)
	return str
}

func (s PostUserSearchRequestCapn) SearchTerm() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s PostUserSearchRequestCapn) HasSearchTerm() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PostUserSearchRequestCapn) SearchTermBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s PostUserSearchRequestCapn) SetSearchTerm(v string) error {
	return s.Struct.SetText(0, v)
}

func (s PostUserSearchRequestCapn) Limit() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s PostUserSearchRequestCapn) SetLimit(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

// PostUserSearchRequestCapn_List is a list of PostUserSearchRequestCapn.
type PostUserSearchRequestCapn_List struct{ capnp.List }

// NewPostUserSearchRequestCapn creates a new list of PostUserSearchRequestCapn.
func NewPostUserSearchRequestCapn_List(s *capnp.Segment, sz int32) (PostUserSearchRequestCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return PostUserSearchRequestCapn_List{l}, err
}

func (s PostUserSearchRequestCapn_List) At(i int) PostUserSearchRequestCapn {
	return PostUserSearchRequestCapn{s.List.Struct(i)}
}

func (s PostUserSearchRequestCapn_List) Set(i int, v PostUserSearchRequestCapn) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s PostUserSearchRequestCapn_List) String() string {
	str, _ := text.MarshalList(0x96745109c202f5a6, s.List)
	return str
}

// PostUserSearchRequestCapn_Promise is a wrapper for a PostUserSearchRequestCapn promised by a client call.
type PostUserSearchRequestCapn_Promise struct{ *capnp.Pipeline }

func (p PostUserSearchRequestCapn_Promise) Struct() (PostUserSearchRequestCapn, error) {
	s, err := p.Pipeline.Struct()
	return PostUserSearchRequestCapn{s}, err
}

type PostUserSearchResponseCapn struct{ capnp.Struct }

// PostUserSearchResponseCapn_TypeID is the unique identifier for the type PostUserSearchResponseCapn.
const PostUserSearchResponseCapn_TypeID = 0x9e37487e38669298

func NewPostUserSearchResponseCapn(s *capnp.Segment) (PostUserSearchResponseCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PostUserSearchResponseCapn{st}, err
}

func NewRootPostUserSearchResponseCapn(s *capnp.Segment) (PostUserSearchResponseCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PostUserSearchResponseCapn{st}, err
}

func ReadRootPostUserSearchResponseCapn(msg *capnp.Message) (PostUserSearchResponseCapn, error) {
	root, err := msg.RootPtr()
	return PostUserSearchResponseCapn{root.Struct()}, err
}

func (s PostUserSearchResponseCapn) String() string {
	str, _ := text.Marshal(0x9e37487e38669298, s.Struct)
	return str
}

func (s PostUserSearchResponseCapn) Results() (UserCapn_List, error) {
	p, err := s.Struct.Ptr(0)
	return UserCapn_List{List: p.List()}, err
}

func (s PostUserSearchResponseCapn) HasResults() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PostUserSearchResponseCapn) SetResults(v UserCapn_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewResults sets the results field to a newly
// allocated UserCapn_List, preferring placement in s's segment.
func (s PostUserSearchResponseCapn) NewResults(n int32) (UserCapn_List, error) {
	l, err := NewUserCapn_List(s.Struct.Segment(), n)
	if err != nil {
		return UserCapn_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s PostUserSearchResponseCapn) Limited() bool {
	return s.Struct.Bit(0)
}

func (s PostUserSearchResponseCapn) SetLimited(v bool) {
	s.Struct.SetBit(0, v)
}

// PostUserSearchResponseCapn_List is a list of PostUserSearchResponseCapn.
type PostUserSearchResponseCapn_List struct{ capnp.List }

// NewPostUserSearchResponseCapn creates a new list of PostUserSearchResponseCapn.
func NewPostUserSearchResponseCapn_List(s *capnp.Segment, sz int32) (PostUserSearchResponseCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return PostUserSearchResponseCapn_List{l}, err
}

func (s PostUserSearchResponseCapn_List) At(i int) PostUserSearchResponseCapn {
	return PostUserSearchResponseCapn{s.List.Struct(i)}
}

func (s PostUserSearchResponseCapn_List) Set(i int, v PostUserSearchResponseCapn) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s PostUserSearchResponseCapn_List) String() string {
	str, _ := text.MarshalList(0x9e37487e38669298, s.List)
	return str
}

// PostUserSearchResponseCapn_Promise is a wrapper for a PostUserSearchResponseCapn promised by a client call.
type PostUserSearchResponseCapn_Promise struct{ *capnp.Pipeline }

func (p PostUserSearchResponseCapn_Promise) Struct() (PostUserSearchResponseCapn, error) {
	s, err := p.Pipeline.Struct()
	return PostUserSearchResponseCapn{s}, err
}

type PutUserProfileAvatarURLRequestCapn struct{ capnp.Struct }

// PutUserProfileAvatarURLRequestCapn_TypeID is the unique identifier for the type PutUserProfileAvatarURLRequestCapn.
const PutUserProfileAvatarURLRequestCapn_TypeID = 0xb43eb716e2a47a96

func NewPutUserProfileAvatarURLRequestCapn(s *capnp.Segment) (PutUserProfileAvatarURLRequestCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return PutUserProfileAvatarURLRequestCapn{st}, err
}

func NewRootPutUserProfileAvatarURLRequestCapn(s *capnp.Segment) (PutUserProfileAvatarURLRequestCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return PutUserProfileAvatarURLRequestCapn{st}, err
}

func ReadRootPutUserProfileAvatarURLRequestCapn(msg *capnp.Message) (PutUserProfileAvatarURLRequestCapn, error) {
	root, err := msg.RootPtr()
	return PutUserProfileAvatarURLRequestCapn{root.Struct()}, err
}

func (s PutUserProfileAvatarURLRequestCapn) String() string {
	str, _ := text.Marshal(0xb43eb716e2a47a96, s.Struct)
	return str
}

func (s PutUserProfileAvatarURLRequestCapn) UserID() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s PutUserProfileAvatarURLRequestCapn) HasUserID() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PutUserProfileAvatarURLRequestCapn) UserIDBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s PutUserProfileAvatarURLRequestCapn) SetUserID(v string) error {
	return s.Struct.SetText(0, v)
}

func (s PutUserProfileAvatarURLRequestCapn) AvatarURL() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s PutUserProfileAvatarURLRequestCapn) HasAvatarURL() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s PutUserProfileAvatarURLRequestCapn) AvatarURLBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s PutUserProfileAvatarURLRequestCapn) SetAvatarURL(v string) error {
	return s.Struct.SetText(1, v)
}

// PutUserProfileAvatarURLRequestCapn_List is a list of PutUserProfileAvatarURLRequestCapn.
type PutUserProfileAvatarURLRequestCapn_List struct{ capnp.List }

// NewPutUserProfileAvatarURLRequestCapn creates a new list of PutUserProfileAvatarURLRequestCapn.
func NewPutUserProfileAvatarURLRequestCapn_List(s *capnp.Segment, sz int32) (PutUserProfileAvatarURLRequestCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return PutUserProfileAvatarURLRequestCapn_List{l}, err
}

func (s PutUserProfileAvatarURLRequestCapn_List) At(i int) PutUserProfileAvatarURLRequestCapn {
	return PutUserProfileAvatarURLRequestCapn{s.List.Struct(i)}
}

func (s PutUserProfileAvatarURLRequestCapn_List) Set(i int, v PutUserProfileAvatarURLRequestCapn) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s PutUserProfileAvatarURLRequestCapn_List) String() string {
	str, _ := text.MarshalList(0xb43eb716e2a47a96, s.List)
	return str
}

// PutUserProfileAvatarURLRequestCapn_Promise is a wrapper for a PutUserProfileAvatarURLRequestCapn promised by a client call.
type PutUserProfileAvatarURLRequestCapn_Promise struct{ *capnp.Pipeline }

func (p PutUserProfileAvatarURLRequestCapn_Promise) Struct() (PutUserProfileAvatarURLRequestCapn, error) {
	s, err := p.Pipeline.Struct()
	return PutUserProfileAvatarURLRequestCapn{s}, err
}

type PutUserProfileDisplayNameRequestCapn struct{ capnp.Struct }

// PutUserProfileDisplayNameRequestCapn_TypeID is the unique identifier for the type PutUserProfileDisplayNameRequestCapn.
const PutUserProfileDisplayNameRequestCapn_TypeID = 0xc216de4a31403538

func NewPutUserProfileDisplayNameRequestCapn(s *capnp.Segment) (PutUserProfileDisplayNameRequestCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return PutUserProfileDisplayNameRequestCapn{st}, err
}

func NewRootPutUserProfileDisplayNameRequestCapn(s *capnp.Segment) (PutUserProfileDisplayNameRequestCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return PutUserProfileDisplayNameRequestCapn{st}, err
}

func ReadRootPutUserProfileDisplayNameRequestCapn(msg *capnp.Message) (PutUserProfileDisplayNameRequestCapn, error) {
	root, err := msg.RootPtr()
	return PutUserProfileDisplayNameRequestCapn{root.Struct()}, err
}

func (s PutUserProfileDisplayNameRequestCapn) String() string {
	str, _ := text.Marshal(0xc216de4a31403538, s.Struct)
	return str
}

func (s PutUserProfileDisplayNameRequestCapn) UserID() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s PutUserProfileDisplayNameRequestCapn) HasUserID() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PutUserProfileDisplayNameRequestCapn) UserIDBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s PutUserProfileDisplayNameRequestCapn) SetUserID(v string) error {
	return s.Struct.SetText(0, v)
}

func (s PutUserProfileDisplayNameRequestCapn) DisplayName() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s PutUserProfileDisplayNameRequestCapn) HasDisplayName() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s PutUserProfileDisplayNameRequestCapn) DisplayNameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s PutUserProfileDisplayNameRequestCapn) SetDisplayName(v string) error {
	return s.Struct.SetText(1, v)
}

// PutUserProfileDisplayNameRequestCapn_List is a list of PutUserProfileDisplayNameRequestCapn.
type PutUserProfileDisplayNameRequestCapn_List struct{ capnp.List }

// NewPutUserProfileDisplayNameRequestCapn creates a new list of PutUserProfileDisplayNameRequestCapn.
func NewPutUserProfileDisplayNameRequestCapn_List(s *capnp.Segment, sz int32) (PutUserProfileDisplayNameRequestCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return PutUserProfileDisplayNameRequestCapn_List{l}, err
}

func (s PutUserProfileDisplayNameRequestCapn_List) At(i int) PutUserProfileDisplayNameRequestCapn {
	return PutUserProfileDisplayNameRequestCapn{s.List.Struct(i)}
}

func (s PutUserProfileDisplayNameRequestCapn_List) Set(i int, v PutUserProfileDisplayNameRequestCapn) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s PutUserProfileDisplayNameRequestCapn_List) String() string {
	str, _ := text.MarshalList(0xc216de4a31403538, s.List)
	return str
}

// PutUserProfileDisplayNameRequestCapn_Promise is a wrapper for a PutUserProfileDisplayNameRequestCapn promised by a client call.
type PutUserProfileDisplayNameRequestCapn_Promise struct{ *capnp.Pipeline }

func (p PutUserProfileDisplayNameRequestCapn_Promise) Struct() (PutUserProfileDisplayNameRequestCapn, error) {
	s, err := p.Pipeline.Struct()
	return PutUserProfileDisplayNameRequestCapn{s}, err
}

type UserCapn struct{ capnp.Struct }

// UserCapn_TypeID is the unique identifier for the type UserCapn.
const UserCapn_TypeID = 0xde55d01bc299d3b9

func NewUserCapn(s *capnp.Segment) (UserCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return UserCapn{st}, err
}

func NewRootUserCapn(s *capnp.Segment) (UserCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return UserCapn{st}, err
}

func ReadRootUserCapn(msg *capnp.Message) (UserCapn, error) {
	root, err := msg.RootPtr()
	return UserCapn{root.Struct()}, err
}

func (s UserCapn) String() string {
	str, _ := text.Marshal(0xde55d01bc299d3b9, s.Struct)
	return str
}

func (s UserCapn) UserID() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s UserCapn) HasUserID() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UserCapn) UserIDBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s UserCapn) SetUserID(v string) error {
	return s.Struct.SetText(0, v)
}

func (s UserCapn) DisplayName() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s UserCapn) HasDisplayName() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s UserCapn) DisplayNameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s UserCapn) SetDisplayName(v string) error {
	return s.Struct.SetText(1, v)
}

func (s UserCapn) AvatarURL() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s UserCapn) HasAvatarURL() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s UserCapn) AvatarURLBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s UserCapn) SetAvatarURL(v string) error {
	return s.Struct.SetText(2, v)
}

// UserCapn_List is a list of UserCapn.
type UserCapn_List struct{ capnp.List }

// NewUserCapn creates a new list of UserCapn.
func NewUserCapn_List(s *capnp.Segment, sz int32) (UserCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return UserCapn_List{l}, err
}

func (s UserCapn_List) At(i int) UserCapn { return UserCapn{s.List.Struct(i)} }

func (s UserCapn_List) Set(i int, v UserCapn) error { return s.List.SetStruct(i, v.Struct) }

func (s UserCapn_List) String() string {
	str, _ := text.MarshalList(0xde55d01bc299d3b9, s.List)
	return str
}

// UserCapn_Promise is a wrapper for a UserCapn promised by a client call.
type UserCapn_Promise struct{ *capnp.Pipeline }

func (p UserCapn_Promise) Struct() (UserCapn, error) {
	s, err := p.Pipeline.Struct()
	return UserCapn{s}, err
}

const schema_87ebf98ac187faed = "x\xda\x94\x92OHT]\x18\xc6\xdf\xe7\x9c\x99O\x05" +
	"\x95o\xb8\x83\x10Q\xcd\x1d\x0b\x8aR\xd4\x92F\xc1f" +
	"\xb2\x91\x1c\xb1\x98\xa3\xce&\xda\xdc\xf4H\x033\xe3x" +
	"\xef\x9d\xa2\x16\xb9\x14$\x17\x06I\xbah\xd5\x1fZD" +
	"\x9b$\x0cZ\xe4\"\xa2MB\x8bv.\xdaF\xd0\xa6" +
	"EA\xdc8\xa7q\x9c\x99F\x9b\xb8\xbb\x97\x1f\xcf\xfb" +
	";\xcf{\xbbz\x9ab\xac\xdb\xff8F4\xfe\x95\xf9" +
	"\xff\xf3\x1e}c\x1bM\xc2]&\xd1\x0ax_~\xcc" +
	"\xbf^\xf8\xfey\x9e\xfch :\xd9\x17\x1a\x84\x91\x08" +
	"5\x10\x19C\xa1\xeb\x04\xef\xde\x9d\xe9\xc8\xad\xe1\xd3\xf7" +
	"k\xd2\x0fC#0\xd65\xbd\xa6\xe9\xe5\x9b\x0f>\xb5" +
	"\xbd8\xf3\x9c\x02\xad\xe50S\xb0i.\xc0\xe83\x15" +
	"\xdck*8\xd2\x1b\xeb\x1e\xd9j\xdb\xa8\x05\xdf5W" +
	"a<\xd5\xf0\x13\x0dwt\xbe;8\xb5y\xfbm\x15" +
	"\xac5\x9a\xc2\x0b0\xcc\xb0\x82\x0f\x84\xa3\x04\xef\xe3\xda" +
	"\xe2Rp\xf1\xfdf-x <\x08\xe3\x82\x86\x13\x1a" +
	"^\xff\xb0\xb2\xb1\x7f3\xb5U\x05sE\xa4\xc3/\x8d" +
	"Y\xcdf\xc3\xcf\x08\xde\xd2\x9b\x95\xe3\x8d\xfbr?k" +
	"\x05\xb7\xb4\xaf\xc28\xd2\xae`\xb3=J^\xc5\xf7\xca" +
	"+8\xd2\x8e[\xae\xe5\xeb\x9c\xb4\xf2\xb9\xfe\xe4\x8c\xe3" +
	"\xa6\x1ci\x8fK\xcb\x9e\xbc:&g\x0b\xd2q\xcfY" +
	"\xf9\x1c%\x01\xd1\xc8}D>\x10\x05\x8e]\"\x12G" +
	"9\xc4)\x06 \x085\xeb\xee!\x12'8D\x84\xc1" +
	"st\xc0\x84$ng\xd1L\x0c\xcd\x84C\x99t6" +
	"\xed\xc2O\x0c~\xc2\xdfV;\xf9\x99\x9c#\xf5\xee\xea" +
	"\xe5\x83D\xe20\x87\x88\xed,\x1fP\xb3\x08\x87\x883" +
	"\xcc\xd9\xd2)d\\\x07\xad\x84$\x07\xfe\xdf)\x93\xa0" +
	"\x86sZDN\x01\xc4\x802\x15\x7fQ\xa5\xa0M\x92" +
	"\xf6\xcct:#\xcf^\xb3\\\xcbN\x8d\x8d\x96\xd7Q" +
	"\xad\xd4_T\xeab\x08l;u\x8c\xed\x14\x12U+" +
	"\x12\xf1\xed*<\xab\x18J\x18-\xcd\xf6\xb4\x88\xa7\x9d" +
	"|\xc6\xbaq\xd1\xca\xca\x92\x07\xcf\xe7\xea\xd1\xb8\xb2\xbb" +
	"\xc6T1\x95\x1a\xac\xac\xdcM\xe4\xbc\xac\xaf\x0e_\xc9" +
	"\xa3Ey4r\x88\xe0\x9f+\xab\xce^\x19^\xfd\xc7" +
	"\xfd[$~G\xa6\x9c\xa8\xb4U\x84Jh.%\x0c" +
	"\xa9\x84\x18\x87\x18-+'\xa1\xca\x19\xe6\x10\x13\x0c\x01" +
	"\xc6\x82`D\x01\xa1\x0e\x97\xe4\x10\x97\xebm\xac\x8es" +
	"V>t\x8fs\xd6\xf5\xe6_\x01\x00\x00\xff\xffPB" +
	"c\xb1"

func init() {
	schemas.Register(schema_87ebf98ac187faed,
		0x96745109c202f5a6,
		0x9e37487e38669298,
		0xb43eb716e2a47a96,
		0xc216de4a31403538,
		0xc88cd0641eca2e2d,
		0xd0cf8d14918db5d7,
		0xde55d01bc299d3b9,
		0xfd6e19082b99c691)
}
