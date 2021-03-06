// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_cloudagent_identity.proto

package security

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type IdentityRequest struct {
	Identitys            []*Identity      `protobuf:"bytes,1,rep,name=Identitys,proto3" json:"Identitys,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *IdentityRequest) Reset()         { *m = IdentityRequest{} }
func (m *IdentityRequest) String() string { return proto.CompactTextString(m) }
func (*IdentityRequest) ProtoMessage()    {}
func (*IdentityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df50a18f5e732c64, []int{0}
}

func (m *IdentityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityRequest.Unmarshal(m, b)
}
func (m *IdentityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityRequest.Marshal(b, m, deterministic)
}
func (m *IdentityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityRequest.Merge(m, src)
}
func (m *IdentityRequest) XXX_Size() int {
	return xxx_messageInfo_IdentityRequest.Size(m)
}
func (m *IdentityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityRequest proto.InternalMessageInfo

func (m *IdentityRequest) GetIdentitys() []*Identity {
	if m != nil {
		return m.Identitys
	}
	return nil
}

func (m *IdentityRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type IdentityResponse struct {
	Identitys            []*Identity         `protobuf:"bytes,1,rep,name=Identitys,proto3" json:"Identitys,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IdentityResponse) Reset()         { *m = IdentityResponse{} }
func (m *IdentityResponse) String() string { return proto.CompactTextString(m) }
func (*IdentityResponse) ProtoMessage()    {}
func (*IdentityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_df50a18f5e732c64, []int{1}
}

func (m *IdentityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityResponse.Unmarshal(m, b)
}
func (m *IdentityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityResponse.Marshal(b, m, deterministic)
}
func (m *IdentityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityResponse.Merge(m, src)
}
func (m *IdentityResponse) XXX_Size() int {
	return xxx_messageInfo_IdentityResponse.Size(m)
}
func (m *IdentityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityResponse proto.InternalMessageInfo

func (m *IdentityResponse) GetIdentitys() []*Identity {
	if m != nil {
		return m.Identitys
	}
	return nil
}

func (m *IdentityResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *IdentityResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type IdentityOperationRequest struct {
	Identities           []*Identity              `protobuf:"bytes,1,rep,name=Identities,proto3" json:"Identities,omitempty"`
	OperationType        common.IdentityOperation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.IdentityOperation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *IdentityOperationRequest) Reset()         { *m = IdentityOperationRequest{} }
func (m *IdentityOperationRequest) String() string { return proto.CompactTextString(m) }
func (*IdentityOperationRequest) ProtoMessage()    {}
func (*IdentityOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df50a18f5e732c64, []int{2}
}

func (m *IdentityOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityOperationRequest.Unmarshal(m, b)
}
func (m *IdentityOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityOperationRequest.Marshal(b, m, deterministic)
}
func (m *IdentityOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityOperationRequest.Merge(m, src)
}
func (m *IdentityOperationRequest) XXX_Size() int {
	return xxx_messageInfo_IdentityOperationRequest.Size(m)
}
func (m *IdentityOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityOperationRequest proto.InternalMessageInfo

func (m *IdentityOperationRequest) GetIdentities() []*Identity {
	if m != nil {
		return m.Identities
	}
	return nil
}

func (m *IdentityOperationRequest) GetOperationType() common.IdentityOperation {
	if m != nil {
		return m.OperationType
	}
	return common.IdentityOperation_REVOKE
}

type IdentityCertificateRequest struct {
	IdentityName         string                              `protobuf:"bytes,1,opt,name=IdentityName,proto3" json:"IdentityName,omitempty"`
	CSR                  []*CertificateSigningRequest        `protobuf:"bytes,2,rep,name=CSR,proto3" json:"CSR,omitempty"`
	OperationType        common.IdentityCertificateOperation `protobuf:"varint,3,opt,name=OperationType,proto3,enum=moc.IdentityCertificateOperation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *IdentityCertificateRequest) Reset()         { *m = IdentityCertificateRequest{} }
func (m *IdentityCertificateRequest) String() string { return proto.CompactTextString(m) }
func (*IdentityCertificateRequest) ProtoMessage()    {}
func (*IdentityCertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df50a18f5e732c64, []int{3}
}

func (m *IdentityCertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityCertificateRequest.Unmarshal(m, b)
}
func (m *IdentityCertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityCertificateRequest.Marshal(b, m, deterministic)
}
func (m *IdentityCertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityCertificateRequest.Merge(m, src)
}
func (m *IdentityCertificateRequest) XXX_Size() int {
	return xxx_messageInfo_IdentityCertificateRequest.Size(m)
}
func (m *IdentityCertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityCertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityCertificateRequest proto.InternalMessageInfo

func (m *IdentityCertificateRequest) GetIdentityName() string {
	if m != nil {
		return m.IdentityName
	}
	return ""
}

func (m *IdentityCertificateRequest) GetCSR() []*CertificateSigningRequest {
	if m != nil {
		return m.CSR
	}
	return nil
}

func (m *IdentityCertificateRequest) GetOperationType() common.IdentityCertificateOperation {
	if m != nil {
		return m.OperationType
	}
	return common.IdentityCertificateOperation_CREATE_CERTIFICATE
}

type IdentityCertificateResponse struct {
	Certificates         []*Certificate      `protobuf:"bytes,1,rep,name=Certificates,proto3" json:"Certificates,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IdentityCertificateResponse) Reset()         { *m = IdentityCertificateResponse{} }
func (m *IdentityCertificateResponse) String() string { return proto.CompactTextString(m) }
func (*IdentityCertificateResponse) ProtoMessage()    {}
func (*IdentityCertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_df50a18f5e732c64, []int{4}
}

func (m *IdentityCertificateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityCertificateResponse.Unmarshal(m, b)
}
func (m *IdentityCertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityCertificateResponse.Marshal(b, m, deterministic)
}
func (m *IdentityCertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityCertificateResponse.Merge(m, src)
}
func (m *IdentityCertificateResponse) XXX_Size() int {
	return xxx_messageInfo_IdentityCertificateResponse.Size(m)
}
func (m *IdentityCertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityCertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityCertificateResponse proto.InternalMessageInfo

func (m *IdentityCertificateResponse) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

func (m *IdentityCertificateResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *IdentityCertificateResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Identity struct {
	Name                 string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string                    `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ResourceGroup        string                    `protobuf:"bytes,3,opt,name=resourceGroup,proto3" json:"resourceGroup,omitempty"`
	Password             string                    `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Token                string                    `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	Certificate          []byte                    `protobuf:"bytes,6,opt,name=certificate,proto3" json:"certificate,omitempty"` // Deprecated: Do not use.
	Status               *common.Status            `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	LocationName         string                    `protobuf:"bytes,10,opt,name=locationName,proto3" json:"locationName,omitempty"`
	Tags                 *common.Tags              `protobuf:"bytes,11,opt,name=tags,proto3" json:"tags,omitempty"`
	Certificates         map[string]string         `protobuf:"bytes,12,rep,name=certificates,proto3" json:"certificates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TokenExpiry          int64                     `protobuf:"varint,13,opt,name=tokenExpiry,proto3" json:"tokenExpiry,omitempty"`
	ClientType           common.ClientType         `protobuf:"varint,14,opt,name=clientType,proto3,enum=moc.ClientType" json:"clientType,omitempty"`
	CloudFqdn            string                    `protobuf:"bytes,15,opt,name=cloudFqdn,proto3" json:"cloudFqdn,omitempty"`
	CloudPort            int32                     `protobuf:"varint,16,opt,name=cloudPort,proto3" json:"cloudPort,omitempty"`
	CloudAuthPort        int32                     `protobuf:"varint,17,opt,name=cloudAuthPort,proto3" json:"cloudAuthPort,omitempty"`
	AuthType             common.AuthenticationType `protobuf:"varint,18,opt,name=authType,proto3,enum=moc.AuthenticationType" json:"authType,omitempty"`
	Revoked              bool                      `protobuf:"varint,19,opt,name=revoked,proto3" json:"revoked,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Identity) Reset()         { *m = Identity{} }
func (m *Identity) String() string { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()    {}
func (*Identity) Descriptor() ([]byte, []int) {
	return fileDescriptor_df50a18f5e732c64, []int{5}
}

func (m *Identity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity.Unmarshal(m, b)
}
func (m *Identity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity.Marshal(b, m, deterministic)
}
func (m *Identity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity.Merge(m, src)
}
func (m *Identity) XXX_Size() int {
	return xxx_messageInfo_Identity.Size(m)
}
func (m *Identity) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity.DiscardUnknown(m)
}

var xxx_messageInfo_Identity proto.InternalMessageInfo

func (m *Identity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Identity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Identity) GetResourceGroup() string {
	if m != nil {
		return m.ResourceGroup
	}
	return ""
}

func (m *Identity) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Identity) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Deprecated: Do not use.
func (m *Identity) GetCertificate() []byte {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func (m *Identity) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Identity) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *Identity) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Identity) GetCertificates() map[string]string {
	if m != nil {
		return m.Certificates
	}
	return nil
}

func (m *Identity) GetTokenExpiry() int64 {
	if m != nil {
		return m.TokenExpiry
	}
	return 0
}

func (m *Identity) GetClientType() common.ClientType {
	if m != nil {
		return m.ClientType
	}
	return common.ClientType_CONTROLPLANE
}

func (m *Identity) GetCloudFqdn() string {
	if m != nil {
		return m.CloudFqdn
	}
	return ""
}

func (m *Identity) GetCloudPort() int32 {
	if m != nil {
		return m.CloudPort
	}
	return 0
}

func (m *Identity) GetCloudAuthPort() int32 {
	if m != nil {
		return m.CloudAuthPort
	}
	return 0
}

func (m *Identity) GetAuthType() common.AuthenticationType {
	if m != nil {
		return m.AuthType
	}
	return common.AuthenticationType_SELFSIGNED
}

func (m *Identity) GetRevoked() bool {
	if m != nil {
		return m.Revoked
	}
	return false
}

func init() {
	proto.RegisterType((*IdentityRequest)(nil), "moc.cloudagent.security.IdentityRequest")
	proto.RegisterType((*IdentityResponse)(nil), "moc.cloudagent.security.IdentityResponse")
	proto.RegisterType((*IdentityOperationRequest)(nil), "moc.cloudagent.security.IdentityOperationRequest")
	proto.RegisterType((*IdentityCertificateRequest)(nil), "moc.cloudagent.security.IdentityCertificateRequest")
	proto.RegisterType((*IdentityCertificateResponse)(nil), "moc.cloudagent.security.IdentityCertificateResponse")
	proto.RegisterType((*Identity)(nil), "moc.cloudagent.security.Identity")
	proto.RegisterMapType((map[string]string)(nil), "moc.cloudagent.security.Identity.CertificatesEntry")
}

func init() { proto.RegisterFile("moc_cloudagent_identity.proto", fileDescriptor_df50a18f5e732c64) }

var fileDescriptor_df50a18f5e732c64 = []byte{
	// 804 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0x66, 0x76, 0x93, 0x4d, 0xf6, 0xec, 0xe6, 0xa7, 0x53, 0x44, 0x06, 0x43, 0x91, 0xbb, 0xe4,
	0xc2, 0xdc, 0xd8, 0xb0, 0xe9, 0x05, 0x42, 0x48, 0x55, 0x12, 0x42, 0xe9, 0x0d, 0xa0, 0x49, 0x05,
	0x12, 0x37, 0x91, 0xe3, 0x9d, 0x38, 0xa3, 0xac, 0x3d, 0xee, 0xcc, 0xb8, 0x65, 0x6f, 0xb9, 0xe2,
	0x1d, 0x10, 0x0f, 0x01, 0x2f, 0xc2, 0x2b, 0xf0, 0x28, 0xc8, 0xc7, 0xf6, 0xda, 0xde, 0x34, 0xda,
	0x46, 0xf4, 0xca, 0x9e, 0x73, 0xbe, 0xf3, 0xcd, 0x77, 0xce, 0x7c, 0x33, 0xf0, 0x28, 0x51, 0xd1,
	0x45, 0x34, 0x57, 0xf9, 0x2c, 0x8c, 0x45, 0x6a, 0x2f, 0xe4, 0x4c, 0xa4, 0x56, 0xda, 0x85, 0x9f,
	0x69, 0x65, 0x15, 0x3d, 0x48, 0x54, 0xe4, 0x37, 0x69, 0xdf, 0x88, 0x28, 0xd7, 0xd2, 0x2e, 0x9c,
	0x4f, 0x62, 0xa5, 0xe2, 0xb9, 0x08, 0x10, 0x76, 0x99, 0x5f, 0x05, 0xaf, 0x75, 0x98, 0x65, 0x42,
	0x9b, 0xb2, 0xd0, 0x71, 0x57, 0x78, 0x23, 0xa1, 0xad, 0xbc, 0x92, 0x51, 0x68, 0x45, 0x85, 0x38,
	0x40, 0x84, 0x4a, 0x12, 0x95, 0x56, 0x9f, 0x2a, 0xf1, 0x61, 0x2b, 0x51, 0xef, 0x57, 0xa6, 0x26,
	0xbf, 0x13, 0xd8, 0x7b, 0x5e, 0x29, 0xe4, 0xe2, 0x65, 0x2e, 0x8c, 0xa5, 0x4f, 0x61, 0x58, 0x87,
	0x0c, 0x23, 0x6e, 0xdf, 0x1b, 0x4d, 0x1f, 0xfb, 0x77, 0xc8, 0xf6, 0x97, 0xc5, 0x4d, 0x0d, 0x7d,
	0x02, 0x3b, 0x3f, 0x64, 0x42, 0x87, 0x56, 0xaa, 0xf4, 0xc5, 0x22, 0x13, 0xac, 0xe7, 0x12, 0x6f,
	0x77, 0xba, 0x8b, 0x24, 0xcb, 0x0c, 0xef, 0x82, 0x26, 0x7f, 0x12, 0xd8, 0x6f, 0xa4, 0x98, 0x4c,
	0xa5, 0x46, 0xfc, 0x7f, 0x2d, 0x53, 0x18, 0x70, 0x61, 0xf2, 0xb9, 0x45, 0x11, 0xa3, 0xa9, 0xe3,
	0x97, 0x73, 0xf6, 0xeb, 0x39, 0xfb, 0x27, 0x4a, 0xcd, 0x7f, 0x0a, 0xe7, 0xb9, 0xe0, 0x15, 0x92,
	0xbe, 0x0f, 0x9b, 0x67, 0x5a, 0x2b, 0xcd, 0xfa, 0x2e, 0xf1, 0x86, 0xbc, 0x5c, 0x4c, 0xfe, 0x20,
	0xc0, 0x6a, 0xde, 0xa6, 0x89, 0x6a, 0x66, 0xc7, 0x00, 0x55, 0x4e, 0x8a, 0x7b, 0x08, 0x6d, 0x15,
	0xd1, 0xaf, 0xdf, 0x3c, 0xb5, 0x0f, 0x90, 0xe5, 0xf6, 0xc6, 0x2b, 0xd3, 0xfb, 0x87, 0x80, 0x53,
	0x83, 0x4e, 0x1b, 0x6b, 0xd4, 0xfa, 0x26, 0x30, 0xae, 0xb3, 0xdf, 0x87, 0x89, 0x60, 0x04, 0x3b,
	0xeb, 0xc4, 0xe8, 0x37, 0xd0, 0x3f, 0x3d, 0xe7, 0xac, 0x87, 0xe2, 0xa7, 0x77, 0x8a, 0x6f, 0xb1,
	0x9f, 0xcb, 0x38, 0x95, 0x69, 0x5c, 0x6d, 0xc2, 0x8b, 0x72, 0xfa, 0x6c, 0xb5, 0x8d, 0x3e, 0xb6,
	0xf1, 0xb8, 0xd3, 0x46, 0x8b, 0xe3, 0xce, 0x8e, 0xfe, 0x22, 0xf0, 0xd1, 0x1b, 0x3b, 0xaa, 0xac,
	0xf1, 0x1d, 0x8c, 0x5b, 0xe1, 0x7a, 0xe8, 0x87, 0x6f, 0xa3, 0x9b, 0x77, 0x2a, 0xdf, 0xa1, 0x47,
	0xfe, 0xde, 0x84, 0xed, 0x5a, 0x33, 0xa5, 0xb0, 0x91, 0x36, 0xb3, 0xc6, 0x7f, 0xba, 0x0b, 0x3d,
	0x39, 0xc3, 0x6d, 0x86, 0xbc, 0x27, 0x67, 0xf4, 0x10, 0x76, 0xb4, 0x30, 0x2a, 0xd7, 0x91, 0x78,
	0xa6, 0x55, 0x9e, 0x55, 0x74, 0xdd, 0x20, 0x75, 0x60, 0x3b, 0x0b, 0x8d, 0x79, 0xad, 0xf4, 0x8c,
	0x6d, 0x20, 0x60, 0xb9, 0x2e, 0x84, 0x58, 0x75, 0x23, 0x52, 0xb6, 0x59, 0x0a, 0xc1, 0x05, 0x3d,
	0x84, 0x51, 0xeb, 0x81, 0x60, 0x03, 0x97, 0x78, 0xe3, 0x93, 0x1e, 0x23, 0xbc, 0x1d, 0xa6, 0x9f,
	0xc2, 0xc0, 0xd8, 0xd0, 0xe6, 0x86, 0x6d, 0x61, 0xe3, 0x23, 0x1c, 0xde, 0x39, 0x86, 0x78, 0x95,
	0x2a, 0xac, 0x33, 0x57, 0x11, 0x9e, 0x0b, 0x5a, 0x07, 0x4a, 0xeb, 0xb4, 0x63, 0xf4, 0x11, 0x6c,
	0xd8, 0x30, 0x36, 0x6c, 0x84, 0x34, 0x43, 0xa4, 0x79, 0x11, 0xc6, 0x86, 0x63, 0x98, 0xfe, 0x0c,
	0xe3, 0xa8, 0x7d, 0x54, 0x63, 0x3c, 0xaa, 0xa3, 0xb5, 0xf7, 0xa3, 0x7d, 0x66, 0xe6, 0x2c, 0xb5,
	0x7a, 0xc1, 0x3b, 0x44, 0xd4, 0x85, 0x11, 0xf6, 0x7b, 0xf6, 0x6b, 0x26, 0xf5, 0x82, 0xed, 0xb8,
	0xc4, 0xeb, 0xf3, 0x76, 0x88, 0x06, 0x00, 0xd1, 0x5c, 0x8a, 0xd4, 0xa2, 0x17, 0x77, 0xd1, 0x8b,
	0x7b, 0xb8, 0xf1, 0xe9, 0x32, 0xcc, 0x5b, 0x10, 0xfa, 0x31, 0x0c, 0x51, 0xd2, 0xb7, 0x2f, 0x67,
	0x29, 0xdb, 0xc3, 0x5e, 0x9b, 0xc0, 0x32, 0xfb, 0xa3, 0xd2, 0x96, 0xed, 0xbb, 0xc4, 0xdb, 0xe4,
	0x4d, 0xa0, 0x38, 0x4d, 0x5c, 0x1c, 0xe7, 0xf6, 0x1a, 0x11, 0x0f, 0x10, 0xd1, 0x0d, 0xd2, 0x23,
	0xd8, 0x0e, 0x73, 0x7b, 0x8d, 0x82, 0x28, 0x0a, 0x3a, 0x40, 0x41, 0x05, 0xa0, 0xe8, 0x3b, 0x5a,
	0xde, 0x01, 0xbe, 0x04, 0x52, 0x06, 0x5b, 0x5a, 0xbc, 0x52, 0x37, 0x62, 0xc6, 0x1e, 0xba, 0xc4,
	0xdb, 0xe6, 0xf5, 0xd2, 0x79, 0x0a, 0x0f, 0x6e, 0x8d, 0x89, 0xee, 0x43, 0xff, 0x46, 0x2c, 0x2a,
	0xeb, 0x15, 0xbf, 0x85, 0x4f, 0x5e, 0x15, 0x0e, 0xae, 0xcc, 0x57, 0x2e, 0xbe, 0xea, 0x7d, 0x49,
	0xa6, 0xff, 0xf6, 0x60, 0xa7, 0x9e, 0xf8, 0x71, 0x71, 0x10, 0xf4, 0x02, 0x06, 0xcf, 0xd3, 0x82,
	0x9d, 0x7a, 0xeb, 0xdf, 0xb0, 0xf2, 0xf2, 0x3b, 0x9f, 0xbd, 0x05, 0xb2, 0xbc, 0xb9, 0x93, 0xf7,
	0xa8, 0x84, 0xad, 0xf2, 0xb2, 0x0b, 0xfa, 0xc5, 0xda, 0xba, 0xd5, 0xc7, 0xf6, 0x7e, 0x5b, 0xfd,
	0x46, 0xe0, 0x61, 0xb5, 0x57, 0xe7, 0xd2, 0xaf, 0x77, 0xdf, 0xed, 0x67, 0xd4, 0x79, 0x72, 0xbf,
	0xa2, 0x5a, 0xc4, 0xc9, 0xf4, 0x97, 0xcf, 0x63, 0x69, 0xaf, 0xf3, 0x4b, 0x3f, 0x52, 0x49, 0x90,
	0xc8, 0x48, 0x2b, 0xa3, 0xae, 0x6c, 0x90, 0xa8, 0x28, 0xd0, 0x59, 0x14, 0x34, 0x8c, 0x41, 0xcd,
	0x78, 0x39, 0xc0, 0xd7, 0xe7, 0xe8, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x8e, 0x47, 0x16,
	0x51, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdentityAgentClient is the client API for IdentityAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentityAgentClient interface {
	Invoke(ctx context.Context, in *IdentityRequest, opts ...grpc.CallOption) (*IdentityResponse, error)
	Operate(ctx context.Context, in *IdentityOperationRequest, opts ...grpc.CallOption) (*IdentityResponse, error)
	OperateCertificates(ctx context.Context, in *IdentityCertificateRequest, opts ...grpc.CallOption) (*IdentityCertificateResponse, error)
}

type identityAgentClient struct {
	cc *grpc.ClientConn
}

func NewIdentityAgentClient(cc *grpc.ClientConn) IdentityAgentClient {
	return &identityAgentClient{cc}
}

func (c *identityAgentClient) Invoke(ctx context.Context, in *IdentityRequest, opts ...grpc.CallOption) (*IdentityResponse, error) {
	out := new(IdentityResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.IdentityAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityAgentClient) Operate(ctx context.Context, in *IdentityOperationRequest, opts ...grpc.CallOption) (*IdentityResponse, error) {
	out := new(IdentityResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.IdentityAgent/Operate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityAgentClient) OperateCertificates(ctx context.Context, in *IdentityCertificateRequest, opts ...grpc.CallOption) (*IdentityCertificateResponse, error) {
	out := new(IdentityCertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.IdentityAgent/OperateCertificates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityAgentServer is the server API for IdentityAgent service.
type IdentityAgentServer interface {
	Invoke(context.Context, *IdentityRequest) (*IdentityResponse, error)
	Operate(context.Context, *IdentityOperationRequest) (*IdentityResponse, error)
	OperateCertificates(context.Context, *IdentityCertificateRequest) (*IdentityCertificateResponse, error)
}

// UnimplementedIdentityAgentServer can be embedded to have forward compatible implementations.
type UnimplementedIdentityAgentServer struct {
}

func (*UnimplementedIdentityAgentServer) Invoke(ctx context.Context, req *IdentityRequest) (*IdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (*UnimplementedIdentityAgentServer) Operate(ctx context.Context, req *IdentityOperationRequest) (*IdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Operate not implemented")
}
func (*UnimplementedIdentityAgentServer) OperateCertificates(ctx context.Context, req *IdentityCertificateRequest) (*IdentityCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OperateCertificates not implemented")
}

func RegisterIdentityAgentServer(s *grpc.Server, srv IdentityAgentServer) {
	s.RegisterService(&_IdentityAgent_serviceDesc, srv)
}

func _IdentityAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.IdentityAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityAgentServer).Invoke(ctx, req.(*IdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityAgent_Operate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityAgentServer).Operate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.IdentityAgent/Operate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityAgentServer).Operate(ctx, req.(*IdentityOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityAgent_OperateCertificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityAgentServer).OperateCertificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.IdentityAgent/OperateCertificates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityAgentServer).OperateCertificates(ctx, req.(*IdentityCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IdentityAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.security.IdentityAgent",
	HandlerType: (*IdentityAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _IdentityAgent_Invoke_Handler,
		},
		{
			MethodName: "Operate",
			Handler:    _IdentityAgent_Operate_Handler,
		},
		{
			MethodName: "OperateCertificates",
			Handler:    _IdentityAgent_OperateCertificates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_cloudagent_identity.proto",
}
