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

type SharedService interface {
  // Parameters:
  //  - Key
  GetStruct(key int32) (_r *SharedStruct, err error)
}

type SharedServiceClientInterface interface {
  thrift.ClientInterface
  // Parameters:
  //  - Key
  GetStruct(key int32) (_r *SharedStruct, err error)
}

type SharedServiceClient struct {
  SharedServiceClientInterface
  CC thrift.ClientConn
}

func(client *SharedServiceClient) Open() error {
  return client.CC.Open()
}

func(client *SharedServiceClient) Close() error {
  return client.CC.Close()
}

func(client *SharedServiceClient) IsOpen() bool {
  return client.CC.IsOpen()
}

func NewSharedServiceClientFactory(t thrift.Transport, f thrift.ProtocolFactory) *SharedServiceClient {
  return &SharedServiceClient{ CC: thrift.NewClientConn(t, f) }
}

func NewSharedServiceClient(t thrift.Transport, iprot thrift.Protocol, oprot thrift.Protocol) *SharedServiceClient {
  return &SharedServiceClient{ CC: thrift.NewClientConnWithProtocols(t, iprot, oprot) }
}

func NewSharedServiceClientProtocol(prot thrift.Protocol) *SharedServiceClient {
  return NewSharedServiceClient(prot.Transport(), prot, prot)
}

// Parameters:
//  - Key
func (p *SharedServiceClient) GetStruct(key int32) (_r *SharedStruct, err error) {
  args := SharedServiceGetStructArgs{
    Key : key,
  }
  err = p.CC.SendMsg("getStruct", &args, thrift.CALL)
  if err != nil { return }
  return p.recvGetStruct()
}


func (p *SharedServiceClient) recvGetStruct() (value *SharedStruct, err error) {
  var result SharedServiceGetStructResult
  err = p.CC.RecvMsg("getStruct", &result)
  if err != nil { return }

  return result.GetSuccess(), nil
}


type SharedServiceThreadsafeClient struct {
  SharedServiceClientInterface
  CC thrift.ClientConn
  Mu sync.Mutex
}

func(client *SharedServiceThreadsafeClient) Open() error {
  client.Mu.Lock()
  defer client.Mu.Unlock()
  return client.CC.Open()
}

func(client *SharedServiceThreadsafeClient) Close() error {
  client.Mu.Lock()
  defer client.Mu.Unlock()
  return client.CC.Close()
}

func(client *SharedServiceThreadsafeClient) IsOpen() bool {
  client.Mu.Lock()
  defer client.Mu.Unlock()
  return client.CC.IsOpen()
}

func NewSharedServiceThreadsafeClientFactory(t thrift.Transport, f thrift.ProtocolFactory) *SharedServiceThreadsafeClient {
  return &SharedServiceThreadsafeClient{ CC: thrift.NewClientConn(t, f) }
}

func NewSharedServiceThreadsafeClient(t thrift.Transport, iprot thrift.Protocol, oprot thrift.Protocol) *SharedServiceThreadsafeClient {
  return &SharedServiceThreadsafeClient{ CC: thrift.NewClientConnWithProtocols(t, iprot, oprot) }
}

func NewSharedServiceThreadsafeClientProtocol(prot thrift.Protocol) *SharedServiceThreadsafeClient {
  return NewSharedServiceThreadsafeClient(prot.Transport(), prot, prot)
}

// Parameters:
//  - Key
func (p *SharedServiceThreadsafeClient) GetStruct(key int32) (_r *SharedStruct, err error) {
  p.Mu.Lock()
  defer p.Mu.Unlock()
  args := SharedServiceGetStructArgs{
    Key : key,
  }
  err = p.CC.SendMsg("getStruct", &args, thrift.CALL)
  if err != nil { return }
  return p.recvGetStruct()
}


func (p *SharedServiceThreadsafeClient) recvGetStruct() (value *SharedStruct, err error) {
  var result SharedServiceGetStructResult
  err = p.CC.RecvMsg("getStruct", &result)
  if err != nil { return }

  return result.GetSuccess(), nil
}


type SharedServiceChannelClient struct {
  RequestChannel thrift.RequestChannel
}

func (c *SharedServiceChannelClient) Close() error {
  return c.RequestChannel.Close()
}

func (c *SharedServiceChannelClient) IsOpen() bool {
  return c.RequestChannel.IsOpen()
}

func (c *SharedServiceChannelClient) Open() error {
  return c.RequestChannel.Open()
}

func NewSharedServiceChannelClient(channel thrift.RequestChannel) *SharedServiceChannelClient {
  return &SharedServiceChannelClient{RequestChannel: channel}
}

// Parameters:
//  - Key
func (p *SharedServiceChannelClient) GetStruct(ctx context.Context, key int32) (_r *SharedStruct, err error) {
  args := SharedServiceGetStructArgs{
    Key : key,
  }
  var result SharedServiceGetStructResult
  err = p.RequestChannel.Call(ctx, "getStruct", &args, &result)
  if err != nil { return }

  return result.GetSuccess(), nil
}


type SharedServiceProcessor struct {
  processorMap map[string]thrift.ProcessorFunction
  functionServiceMap map[string]string
  handler SharedService
}

func (p *SharedServiceProcessor) AddToProcessorMap(key string, processor thrift.ProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *SharedServiceProcessor) AddToFunctionServiceMap(key, service string) {
  p.functionServiceMap[key] = service
}

func (p *SharedServiceProcessor) GetProcessorFunction(key string) (processor thrift.ProcessorFunction, err error) {
  if processor, ok := p.processorMap[key]; ok {
    return processor, nil
  }
  return nil, nil // generic error message will be sent
}

func (p *SharedServiceProcessor) ProcessorMap() map[string]thrift.ProcessorFunction {
  return p.processorMap
}

func (p *SharedServiceProcessor) FunctionServiceMap() map[string]string {
  return p.functionServiceMap
}

func NewSharedServiceProcessor(handler SharedService) *SharedServiceProcessor {
  self0 := &SharedServiceProcessor{handler:handler, processorMap:make(map[string]thrift.ProcessorFunction), functionServiceMap:make(map[string]string)}
  self0.processorMap["getStruct"] = &sharedServiceProcessorGetStruct{handler:handler}
  self0.functionServiceMap["getStruct"] = "SharedService"
  return self0
}

type sharedServiceProcessorGetStruct struct {
  handler SharedService
}

func (p *SharedServiceGetStructResult) Exception() thrift.WritableException {
  if p == nil { return nil }
  return nil
}

func (p *sharedServiceProcessorGetStruct) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
  args := SharedServiceGetStructArgs{}
  if err := args.Read(iprot); err != nil {
    return nil, err
  }
  iprot.ReadMessageEnd()
  return &args, nil
}

func (p *sharedServiceProcessorGetStruct) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
  var err2 error
  messageType := thrift.REPLY
  switch result.(type) {
  case thrift.ApplicationException:
    messageType = thrift.EXCEPTION
  }
  if err2 = oprot.WriteMessageBegin("getStruct", messageType, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  return err
}

func (p *sharedServiceProcessorGetStruct) Run(argStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
  args := argStruct.(*SharedServiceGetStructArgs)
  var result SharedServiceGetStructResult
  if retval, err := p.handler.GetStruct(args.Key); err != nil {
    switch err.(type) {
    default:
      x := thrift.NewApplicationExceptionCause(thrift.INTERNAL_ERROR, "Internal error processing getStruct: " + err.Error(), err)
      return x, x
    }
  } else {
    result.Success = retval
  }
  return &result, nil
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Key
type SharedServiceGetStructArgs struct {
  thrift.IRequest
  Key int32 `thrift:"key,1" db:"key" json:"key"`
}

func NewSharedServiceGetStructArgs() *SharedServiceGetStructArgs {
  return &SharedServiceGetStructArgs{}
}


func (p *SharedServiceGetStructArgs) GetKey() int32 {
  return p.Key
}
type SharedServiceGetStructArgsBuilder struct {
  obj *SharedServiceGetStructArgs
}

func NewSharedServiceGetStructArgsBuilder() *SharedServiceGetStructArgsBuilder{
  return &SharedServiceGetStructArgsBuilder{
    obj: NewSharedServiceGetStructArgs(),
  }
}

func (p SharedServiceGetStructArgsBuilder) Emit() *SharedServiceGetStructArgs{
  return &SharedServiceGetStructArgs{
    Key: p.obj.Key,
  }
}

func (s *SharedServiceGetStructArgsBuilder) Key(key int32) *SharedServiceGetStructArgsBuilder {
  s.obj.Key = key
  return s
}

func (s *SharedServiceGetStructArgs) SetKey(key int32) *SharedServiceGetStructArgs {
  s.Key = key
  return s
}

func (p *SharedServiceGetStructArgs) Read(iprot thrift.Protocol) error {
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

func (p *SharedServiceGetStructArgs)  ReadField1(iprot thrift.Protocol) error {
  if v, err := iprot.ReadI32(); err != nil {
    return thrift.PrependError("error reading field 1: ", err)
  } else {
    p.Key = v
  }
  return nil
}

func (p *SharedServiceGetStructArgs) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("getStruct_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *SharedServiceGetStructArgs) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("key", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err) }
  if err := oprot.WriteI32(int32(p.Key)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err) }
  return err
}

func (p *SharedServiceGetStructArgs) String() string {
  if p == nil {
    return "<nil>"
  }

  keyVal := fmt.Sprintf("%v", p.Key)
  return fmt.Sprintf("SharedServiceGetStructArgs({Key:%s})", keyVal)
}

// Attributes:
//  - Success
type SharedServiceGetStructResult struct {
  thrift.IResponse
  Success *SharedStruct `thrift:"success,0,optional" db:"success" json:"success,omitempty"`
}

func NewSharedServiceGetStructResult() *SharedServiceGetStructResult {
  return &SharedServiceGetStructResult{}
}

var SharedServiceGetStructResult_Success_DEFAULT *SharedStruct
func (p *SharedServiceGetStructResult) GetSuccess() *SharedStruct {
  if !p.IsSetSuccess() {
    return SharedServiceGetStructResult_Success_DEFAULT
  }
return p.Success
}
func (p *SharedServiceGetStructResult) IsSetSuccess() bool {
  return p != nil && p.Success != nil
}

type SharedServiceGetStructResultBuilder struct {
  obj *SharedServiceGetStructResult
}

func NewSharedServiceGetStructResultBuilder() *SharedServiceGetStructResultBuilder{
  return &SharedServiceGetStructResultBuilder{
    obj: NewSharedServiceGetStructResult(),
  }
}

func (p SharedServiceGetStructResultBuilder) Emit() *SharedServiceGetStructResult{
  return &SharedServiceGetStructResult{
    Success: p.obj.Success,
  }
}

func (s *SharedServiceGetStructResultBuilder) Success(success *SharedStruct) *SharedServiceGetStructResultBuilder {
  s.obj.Success = success
  return s
}

func (s *SharedServiceGetStructResult) SetSuccess(success *SharedStruct) *SharedServiceGetStructResult {
  s.Success = success
  return s
}

func (p *SharedServiceGetStructResult) Read(iprot thrift.Protocol) error {
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
    case 0:
      if err := p.ReadField0(iprot); err != nil {
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

func (p *SharedServiceGetStructResult)  ReadField0(iprot thrift.Protocol) error {
  p.Success = NewSharedStruct()
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *SharedServiceGetStructResult) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("getStruct_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField0(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *SharedServiceGetStructResult) writeField0(oprot thrift.Protocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *SharedServiceGetStructResult) String() string {
  if p == nil {
    return "<nil>"
  }

  var successVal string
  if p.Success == nil {
    successVal = "<nil>"
  } else {
    successVal = fmt.Sprintf("%v", p.Success)
  }
  return fmt.Sprintf("SharedServiceGetStructResult({Success:%s})", successVal)
}

