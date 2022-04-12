// Autogenerated by Thrift Compiler (facebook)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// @generated

package shared

import (
	"bytes"
	"context"
	"sync"
	"fmt"
	thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = sync.Mutex{}
var _ = bytes.Equal
var _ = context.Background

var GoUnusedProtection__ int;

// Attributes:
//  - Key
//  - Value
type SharedStruct struct {
  Key int32 `thrift:"key,1" db:"key" json:"key"`
  Value string `thrift:"value,2" db:"value" json:"value"`
}

func NewSharedStruct() *SharedStruct {
  return &SharedStruct{}
}


func (p *SharedStruct) GetKey() int32 {
  return p.Key
}

func (p *SharedStruct) GetValue() string {
  return p.Value
}
type SharedStructBuilder struct {
  obj *SharedStruct
}

func NewSharedStructBuilder() *SharedStructBuilder{
  return &SharedStructBuilder{
    obj: NewSharedStruct(),
  }
}

func (p SharedStructBuilder) Emit() *SharedStruct{
  return &SharedStruct{
    Key: p.obj.Key,
    Value: p.obj.Value,
  }
}

func (s *SharedStructBuilder) Key(key int32) *SharedStructBuilder {
  s.obj.Key = key
  return s
}

func (s *SharedStructBuilder) Value(value string) *SharedStructBuilder {
  s.obj.Value = value
  return s
}

func (s *SharedStruct) SetKey(key int32) *SharedStruct {
  s.Key = key
  return s
}

func (s *SharedStruct) SetValue(value string) *SharedStruct {
  s.Value = value
  return s
}

func (p *SharedStruct) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *SharedStruct)  ReadField1(iprot thrift.Protocol) error {
  if v, err := iprot.ReadI32(); err != nil {
    return thrift.PrependError("error reading field 1: ", err)
  } else {
    p.Key = v
  }
  return nil
}

func (p *SharedStruct)  ReadField2(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 2: ", err)
  } else {
    p.Value = v
  }
  return nil
}

func (p *SharedStruct) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("SharedStruct"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *SharedStruct) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("key", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err) }
  if err := oprot.WriteI32(int32(p.Key)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err) }
  return err
}

func (p *SharedStruct) writeField2(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("value", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:value: ", p), err) }
  if err := oprot.WriteString(string(p.Value)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.value (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:value: ", p), err) }
  return err
}

func (p *SharedStruct) String() string {
  if p == nil {
    return "<nil>"
  }

  keyVal := fmt.Sprintf("%v", p.Key)
  valueVal := fmt.Sprintf("%v", p.Value)
  return fmt.Sprintf("SharedStruct({Key:%s Value:%s})", keyVal, valueVal)
}

