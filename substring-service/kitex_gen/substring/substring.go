// Code generated by thriftgo (0.2.12). DO NOT EDIT.

package substring

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type SubstringRequest struct {
	MainString string `thrift:"mainString,1,required" frugal:"1,required,string" json:"mainString"`
	SubString  string `thrift:"subString,2,required" frugal:"2,required,string" json:"subString"`
}

func NewSubstringRequest() *SubstringRequest {
	return &SubstringRequest{}
}

func (p *SubstringRequest) InitDefault() {
	*p = SubstringRequest{}
}

func (p *SubstringRequest) GetMainString() (v string) {
	return p.MainString
}

func (p *SubstringRequest) GetSubString() (v string) {
	return p.SubString
}
func (p *SubstringRequest) SetMainString(val string) {
	p.MainString = val
}
func (p *SubstringRequest) SetSubString(val string) {
	p.SubString = val
}

var fieldIDToName_SubstringRequest = map[int16]string{
	1: "mainString",
	2: "subString",
}

func (p *SubstringRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetMainString bool = false
	var issetSubString bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetMainString = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetSubString = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetMainString {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetSubString {
		fieldId = 2
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SubstringRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_SubstringRequest[fieldId]))
}

func (p *SubstringRequest) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.MainString = v
	}
	return nil
}

func (p *SubstringRequest) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.SubString = v
	}
	return nil
}

func (p *SubstringRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("SubstringRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *SubstringRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("mainString", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.MainString); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *SubstringRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("subString", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.SubString); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *SubstringRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SubstringRequest(%+v)", *p)
}

func (p *SubstringRequest) DeepEqual(ano *SubstringRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.MainString) {
		return false
	}
	if !p.Field2DeepEqual(ano.SubString) {
		return false
	}
	return true
}

func (p *SubstringRequest) Field1DeepEqual(src string) bool {

	if strings.Compare(p.MainString, src) != 0 {
		return false
	}
	return true
}
func (p *SubstringRequest) Field2DeepEqual(src string) bool {

	if strings.Compare(p.SubString, src) != 0 {
		return false
	}
	return true
}

type SubstringResponse struct {
	Positions []int32 `thrift:"positions,1,required" frugal:"1,required,list<i32>" json:"positions"`
}

func NewSubstringResponse() *SubstringResponse {
	return &SubstringResponse{}
}

func (p *SubstringResponse) InitDefault() {
	*p = SubstringResponse{}
}

func (p *SubstringResponse) GetPositions() (v []int32) {
	return p.Positions
}
func (p *SubstringResponse) SetPositions(val []int32) {
	p.Positions = val
}

var fieldIDToName_SubstringResponse = map[int16]string{
	1: "positions",
}

func (p *SubstringResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetPositions bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetPositions = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetPositions {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SubstringResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_SubstringResponse[fieldId]))
}

func (p *SubstringResponse) ReadField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.Positions = make([]int32, 0, size)
	for i := 0; i < size; i++ {
		var _elem int32
		if v, err := iprot.ReadI32(); err != nil {
			return err
		} else {
			_elem = v
		}

		p.Positions = append(p.Positions, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}

func (p *SubstringResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("SubstringResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *SubstringResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("positions", thrift.LIST, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteListBegin(thrift.I32, len(p.Positions)); err != nil {
		return err
	}
	for _, v := range p.Positions {
		if err := oprot.WriteI32(v); err != nil {
			return err
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *SubstringResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SubstringResponse(%+v)", *p)
}

func (p *SubstringResponse) DeepEqual(ano *SubstringResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Positions) {
		return false
	}
	return true
}

func (p *SubstringResponse) Field1DeepEqual(src []int32) bool {

	if len(p.Positions) != len(src) {
		return false
	}
	for i, v := range p.Positions {
		_src := src[i]
		if v != _src {
			return false
		}
	}
	return true
}

type SubstringService interface {
	FindSubstring(ctx context.Context, req *SubstringRequest) (r *SubstringResponse, err error)
}

type SubstringServiceClient struct {
	c thrift.TClient
}

func NewSubstringServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *SubstringServiceClient {
	return &SubstringServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewSubstringServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *SubstringServiceClient {
	return &SubstringServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewSubstringServiceClient(c thrift.TClient) *SubstringServiceClient {
	return &SubstringServiceClient{
		c: c,
	}
}

func (p *SubstringServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *SubstringServiceClient) FindSubstring(ctx context.Context, req *SubstringRequest) (r *SubstringResponse, err error) {
	var _args SubstringServiceFindSubstringArgs
	_args.Req = req
	var _result SubstringServiceFindSubstringResult
	if err = p.Client_().Call(ctx, "findSubstring", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type SubstringServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      SubstringService
}

func (p *SubstringServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *SubstringServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *SubstringServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewSubstringServiceProcessor(handler SubstringService) *SubstringServiceProcessor {
	self := &SubstringServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("findSubstring", &substringServiceProcessorFindSubstring{handler: handler})
	return self
}
func (p *SubstringServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type substringServiceProcessorFindSubstring struct {
	handler SubstringService
}

func (p *substringServiceProcessorFindSubstring) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := SubstringServiceFindSubstringArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("findSubstring", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := SubstringServiceFindSubstringResult{}
	var retval *SubstringResponse
	if retval, err2 = p.handler.FindSubstring(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing findSubstring: "+err2.Error())
		oprot.WriteMessageBegin("findSubstring", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("findSubstring", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type SubstringServiceFindSubstringArgs struct {
	Req *SubstringRequest `thrift:"req,1" frugal:"1,default,SubstringRequest" json:"req"`
}

func NewSubstringServiceFindSubstringArgs() *SubstringServiceFindSubstringArgs {
	return &SubstringServiceFindSubstringArgs{}
}

func (p *SubstringServiceFindSubstringArgs) InitDefault() {
	*p = SubstringServiceFindSubstringArgs{}
}

var SubstringServiceFindSubstringArgs_Req_DEFAULT *SubstringRequest

func (p *SubstringServiceFindSubstringArgs) GetReq() (v *SubstringRequest) {
	if !p.IsSetReq() {
		return SubstringServiceFindSubstringArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *SubstringServiceFindSubstringArgs) SetReq(val *SubstringRequest) {
	p.Req = val
}

var fieldIDToName_SubstringServiceFindSubstringArgs = map[int16]string{
	1: "req",
}

func (p *SubstringServiceFindSubstringArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SubstringServiceFindSubstringArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SubstringServiceFindSubstringArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *SubstringServiceFindSubstringArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Req = NewSubstringRequest()
	if err := p.Req.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *SubstringServiceFindSubstringArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("findSubstring_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *SubstringServiceFindSubstringArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *SubstringServiceFindSubstringArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SubstringServiceFindSubstringArgs(%+v)", *p)
}

func (p *SubstringServiceFindSubstringArgs) DeepEqual(ano *SubstringServiceFindSubstringArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *SubstringServiceFindSubstringArgs) Field1DeepEqual(src *SubstringRequest) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type SubstringServiceFindSubstringResult struct {
	Success *SubstringResponse `thrift:"success,0,optional" frugal:"0,optional,SubstringResponse" json:"success,omitempty"`
}

func NewSubstringServiceFindSubstringResult() *SubstringServiceFindSubstringResult {
	return &SubstringServiceFindSubstringResult{}
}

func (p *SubstringServiceFindSubstringResult) InitDefault() {
	*p = SubstringServiceFindSubstringResult{}
}

var SubstringServiceFindSubstringResult_Success_DEFAULT *SubstringResponse

func (p *SubstringServiceFindSubstringResult) GetSuccess() (v *SubstringResponse) {
	if !p.IsSetSuccess() {
		return SubstringServiceFindSubstringResult_Success_DEFAULT
	}
	return p.Success
}
func (p *SubstringServiceFindSubstringResult) SetSuccess(x interface{}) {
	p.Success = x.(*SubstringResponse)
}

var fieldIDToName_SubstringServiceFindSubstringResult = map[int16]string{
	0: "success",
}

func (p *SubstringServiceFindSubstringResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SubstringServiceFindSubstringResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SubstringServiceFindSubstringResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *SubstringServiceFindSubstringResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewSubstringResponse()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *SubstringServiceFindSubstringResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("findSubstring_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *SubstringServiceFindSubstringResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *SubstringServiceFindSubstringResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SubstringServiceFindSubstringResult(%+v)", *p)
}

func (p *SubstringServiceFindSubstringResult) DeepEqual(ano *SubstringServiceFindSubstringResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *SubstringServiceFindSubstringResult) Field0DeepEqual(src *SubstringResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}
