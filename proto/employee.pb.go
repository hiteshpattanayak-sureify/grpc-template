// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: employee.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
	// This is a compile-time assertion to ensure that this generated file
	// is compatible with the grpc package it is being compiled against.
	// Requires gRPC-Go v1.32.0 or later.
	 _ = grpc.SupportPackageIsVersion7
)

type Employee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Dept   string   `protobuf:"bytes,3,opt,name=dept,proto3" json:"dept,omitempty"`
	Skills []string `protobuf:"bytes,4,rep,name=skills,proto3" json:"skills,omitempty"`
}

func (x *Employee) Reset() {
	*x = Employee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
	mi := &file_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_employee_proto_rawDescGZIP(), []int{0}
}

func (x *Employee) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Employee) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Employee) GetDept() string {
	if x != nil {
		return x.Dept
	}
	return ""
}

func (x *Employee) GetSkills() []string {
	if x != nil {
		return x.Skills
	}
	return nil
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_employee_proto_rawDescGZIP(), []int{1}
}

func (x *ID) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_employee_proto protoreflect.FileDescriptor

var file_employee_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x70, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x70, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x73, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0xc8, 0x01, 0x0a, 0x0f, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x1a, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x44, 0x22, 0x00, 0x12, 0x2b, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x09, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x44, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x1a, 0x09, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x44, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x09, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x44, 0x1a, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x44, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_employee_proto_rawDescOnce sync.Once
	file_employee_proto_rawDescData = file_employee_proto_rawDesc
)

func file_employee_proto_rawDescGZIP() []byte {
	file_employee_proto_rawDescOnce.Do(func() {
		file_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_employee_proto_rawDescData)
	})
	return file_employee_proto_rawDescData
}

var file_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_employee_proto_goTypes = []interface{}{
	(*Employee)(nil), // 0: proto.Employee
	(*ID)(nil),       // 1: proto.ID
}
var file_employee_proto_depIdxs = []int32{
	0, // 0: proto.EmployeeService.CreateEmployee:input_type -> proto.Employee
	1, // 1: proto.EmployeeService.GetEmployee:input_type -> proto.ID
	0, // 2: proto.EmployeeService.UpdateEmployee:input_type -> proto.Employee
	1, // 3: proto.EmployeeService.DeleteEmployee:input_type -> proto.ID
	1, // 4: proto.EmployeeService.CreateEmployee:output_type -> proto.ID
	0, // 5: proto.EmployeeService.GetEmployee:output_type -> proto.Employee
	1, // 6: proto.EmployeeService.UpdateEmployee:output_type -> proto.ID
	1, // 7: proto.EmployeeService.DeleteEmployee:output_type -> proto.ID
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_employee_proto_init() }
func file_employee_proto_init() {
	if File_employee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_employee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Employee); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_employee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_employee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_employee_proto_goTypes,
		DependencyIndexes: file_employee_proto_depIdxs,
		MessageInfos:      file_employee_proto_msgTypes,
	}.Build()
	File_employee_proto = out.File
	file_employee_proto_rawDesc = nil
	file_employee_proto_goTypes = nil
	file_employee_proto_depIdxs = nil
}

// EmployeeServiceClient is the client API for EmployeeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployeeServiceClient interface {
	CreateEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*ID, error)
	GetEmployee(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Employee, error)
	UpdateEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*ID, error)
	DeleteEmployee(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ID, error)
}

type employeeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeServiceClient(cc grpc.ClientConnInterface) EmployeeServiceClient {
	return &employeeServiceClient{cc}
}

func (c *employeeServiceClient) CreateEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/proto.EmployeeService/CreateEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) GetEmployee(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Employee, error) {
	out := new(Employee)
	err := c.cc.Invoke(ctx, "/proto.EmployeeService/GetEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) UpdateEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/proto.EmployeeService/UpdateEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) DeleteEmployee(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/proto.EmployeeService/DeleteEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeServiceServer is the server API for EmployeeService service.
// All implementations must embed UnimplementedEmployeeServiceServer
// for forward compatibility
type EmployeeServiceServer interface {
	CreateEmployee(context.Context, *Employee) (*ID, error)
	GetEmployee(context.Context, *ID) (*Employee, error)
	UpdateEmployee(context.Context, *Employee) (*ID, error)
	DeleteEmployee(context.Context, *ID) (*ID, error)
	mustEmbedUnimplementedEmployeeServiceServer()
}

// UnimplementedEmployeeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmployeeServiceServer struct {
}

func (UnimplementedEmployeeServiceServer) CreateEmployee(context.Context, *Employee) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) GetEmployee(context.Context, *ID) (*Employee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) UpdateEmployee(context.Context, *Employee) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) DeleteEmployee(context.Context, *ID) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) mustEmbedUnimplementedEmployeeServiceServer() {}

// UnsafeEmployeeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployeeServiceServer will
// result in compilation errors.
type UnsafeEmployeeServiceServer interface {
	mustEmbedUnimplementedEmployeeServiceServer()
}

func RegisterEmployeeServiceServer(s grpc.ServiceRegistrar, srv EmployeeServiceServer) {
	s.RegisterService(&EmployeeService_ServiceDesc, srv)
}

func _EmployeeService_CreateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).CreateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmployeeService/CreateEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).CreateEmployee(ctx, req.(*Employee))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_GetEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).GetEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmployeeService/GetEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).GetEmployee(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_UpdateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).UpdateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmployeeService/UpdateEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).UpdateEmployee(ctx, req.(*Employee))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_DeleteEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).DeleteEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmployeeService/DeleteEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).DeleteEmployee(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// EmployeeService_ServiceDesc is the grpc.ServiceDesc for EmployeeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmployeeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EmployeeService",
	HandlerType: (*EmployeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEmployee",
			Handler:    _EmployeeService_CreateEmployee_Handler,
		},
		{
			MethodName: "GetEmployee",
			Handler:    _EmployeeService_GetEmployee_Handler,
		},
		{
			MethodName: "UpdateEmployee",
			Handler:    _EmployeeService_UpdateEmployee_Handler,
		},
		{
			MethodName: "DeleteEmployee",
			Handler:    _EmployeeService_DeleteEmployee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "employee.proto",
}