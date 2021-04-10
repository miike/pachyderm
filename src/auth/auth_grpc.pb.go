// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package auth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIClient interface {
	// Activate/Deactivate the auth API. 'Activate' sets an initial set of admins
	// for the Pachyderm cluster, and 'Deactivate' removes all ACLs, tokens, and
	// admins from the Pachyderm cluster, making all data publicly accessable
	Activate(ctx context.Context, in *ActivateRequest, opts ...grpc.CallOption) (*ActivateResponse, error)
	Deactivate(ctx context.Context, in *DeactivateRequest, opts ...grpc.CallOption) (*DeactivateResponse, error)
	GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*GetConfigurationResponse, error)
	SetConfiguration(ctx context.Context, in *SetConfigurationRequest, opts ...grpc.CallOption) (*SetConfigurationResponse, error)
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error)
	WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIResponse, error)
	ModifyRoleBinding(ctx context.Context, in *ModifyRoleBindingRequest, opts ...grpc.CallOption) (*ModifyRoleBindingResponse, error)
	GetRoleBinding(ctx context.Context, in *GetRoleBindingRequest, opts ...grpc.CallOption) (*GetRoleBindingResponse, error)
	GetOIDCLogin(ctx context.Context, in *GetOIDCLoginRequest, opts ...grpc.CallOption) (*GetOIDCLoginResponse, error)
	GetAuthToken(ctx context.Context, in *GetAuthTokenRequest, opts ...grpc.CallOption) (*GetAuthTokenResponse, error)
	GetRobotToken(ctx context.Context, in *GetRobotTokenRequest, opts ...grpc.CallOption) (*GetRobotTokenResponse, error)
	ExtendAuthToken(ctx context.Context, in *ExtendAuthTokenRequest, opts ...grpc.CallOption) (*ExtendAuthTokenResponse, error)
	RevokeAuthToken(ctx context.Context, in *RevokeAuthTokenRequest, opts ...grpc.CallOption) (*RevokeAuthTokenResponse, error)
	SetGroupsForUser(ctx context.Context, in *SetGroupsForUserRequest, opts ...grpc.CallOption) (*SetGroupsForUserResponse, error)
	ModifyMembers(ctx context.Context, in *ModifyMembersRequest, opts ...grpc.CallOption) (*ModifyMembersResponse, error)
	GetGroups(ctx context.Context, in *GetGroupsRequest, opts ...grpc.CallOption) (*GetGroupsResponse, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	ExtractAuthTokens(ctx context.Context, in *ExtractAuthTokensRequest, opts ...grpc.CallOption) (*ExtractAuthTokensResponse, error)
	RestoreAuthToken(ctx context.Context, in *RestoreAuthTokenRequest, opts ...grpc.CallOption) (*RestoreAuthTokenResponse, error)
}

type aPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIClient(cc grpc.ClientConnInterface) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Activate(ctx context.Context, in *ActivateRequest, opts ...grpc.CallOption) (*ActivateResponse, error) {
	out := new(ActivateResponse)
	err := c.cc.Invoke(ctx, "/auth.API/Activate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Deactivate(ctx context.Context, in *DeactivateRequest, opts ...grpc.CallOption) (*DeactivateResponse, error) {
	out := new(DeactivateResponse)
	err := c.cc.Invoke(ctx, "/auth.API/Deactivate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*GetConfigurationResponse, error) {
	out := new(GetConfigurationResponse)
	err := c.cc.Invoke(ctx, "/auth.API/GetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) SetConfiguration(ctx context.Context, in *SetConfigurationRequest, opts ...grpc.CallOption) (*SetConfigurationResponse, error) {
	out := new(SetConfigurationResponse)
	err := c.cc.Invoke(ctx, "/auth.API/SetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/auth.API/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error) {
	out := new(AuthorizeResponse)
	err := c.cc.Invoke(ctx, "/auth.API/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIResponse, error) {
	out := new(WhoAmIResponse)
	err := c.cc.Invoke(ctx, "/auth.API/WhoAmI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ModifyRoleBinding(ctx context.Context, in *ModifyRoleBindingRequest, opts ...grpc.CallOption) (*ModifyRoleBindingResponse, error) {
	out := new(ModifyRoleBindingResponse)
	err := c.cc.Invoke(ctx, "/auth.API/ModifyRoleBinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetRoleBinding(ctx context.Context, in *GetRoleBindingRequest, opts ...grpc.CallOption) (*GetRoleBindingResponse, error) {
	out := new(GetRoleBindingResponse)
	err := c.cc.Invoke(ctx, "/auth.API/GetRoleBinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetOIDCLogin(ctx context.Context, in *GetOIDCLoginRequest, opts ...grpc.CallOption) (*GetOIDCLoginResponse, error) {
	out := new(GetOIDCLoginResponse)
	err := c.cc.Invoke(ctx, "/auth.API/GetOIDCLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetAuthToken(ctx context.Context, in *GetAuthTokenRequest, opts ...grpc.CallOption) (*GetAuthTokenResponse, error) {
	out := new(GetAuthTokenResponse)
	err := c.cc.Invoke(ctx, "/auth.API/GetAuthToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetRobotToken(ctx context.Context, in *GetRobotTokenRequest, opts ...grpc.CallOption) (*GetRobotTokenResponse, error) {
	out := new(GetRobotTokenResponse)
	err := c.cc.Invoke(ctx, "/auth.API/GetRobotToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ExtendAuthToken(ctx context.Context, in *ExtendAuthTokenRequest, opts ...grpc.CallOption) (*ExtendAuthTokenResponse, error) {
	out := new(ExtendAuthTokenResponse)
	err := c.cc.Invoke(ctx, "/auth.API/ExtendAuthToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) RevokeAuthToken(ctx context.Context, in *RevokeAuthTokenRequest, opts ...grpc.CallOption) (*RevokeAuthTokenResponse, error) {
	out := new(RevokeAuthTokenResponse)
	err := c.cc.Invoke(ctx, "/auth.API/RevokeAuthToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) SetGroupsForUser(ctx context.Context, in *SetGroupsForUserRequest, opts ...grpc.CallOption) (*SetGroupsForUserResponse, error) {
	out := new(SetGroupsForUserResponse)
	err := c.cc.Invoke(ctx, "/auth.API/SetGroupsForUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ModifyMembers(ctx context.Context, in *ModifyMembersRequest, opts ...grpc.CallOption) (*ModifyMembersResponse, error) {
	out := new(ModifyMembersResponse)
	err := c.cc.Invoke(ctx, "/auth.API/ModifyMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetGroups(ctx context.Context, in *GetGroupsRequest, opts ...grpc.CallOption) (*GetGroupsResponse, error) {
	out := new(GetGroupsResponse)
	err := c.cc.Invoke(ctx, "/auth.API/GetGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, "/auth.API/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ExtractAuthTokens(ctx context.Context, in *ExtractAuthTokensRequest, opts ...grpc.CallOption) (*ExtractAuthTokensResponse, error) {
	out := new(ExtractAuthTokensResponse)
	err := c.cc.Invoke(ctx, "/auth.API/ExtractAuthTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) RestoreAuthToken(ctx context.Context, in *RestoreAuthTokenRequest, opts ...grpc.CallOption) (*RestoreAuthTokenResponse, error) {
	out := new(RestoreAuthTokenResponse)
	err := c.cc.Invoke(ctx, "/auth.API/RestoreAuthToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
// All implementations must embed UnimplementedAPIServer
// for forward compatibility
type APIServer interface {
	// Activate/Deactivate the auth API. 'Activate' sets an initial set of admins
	// for the Pachyderm cluster, and 'Deactivate' removes all ACLs, tokens, and
	// admins from the Pachyderm cluster, making all data publicly accessable
	Activate(context.Context, *ActivateRequest) (*ActivateResponse, error)
	Deactivate(context.Context, *DeactivateRequest) (*DeactivateResponse, error)
	GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error)
	SetConfiguration(context.Context, *SetConfigurationRequest) (*SetConfigurationResponse, error)
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error)
	WhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIResponse, error)
	ModifyRoleBinding(context.Context, *ModifyRoleBindingRequest) (*ModifyRoleBindingResponse, error)
	GetRoleBinding(context.Context, *GetRoleBindingRequest) (*GetRoleBindingResponse, error)
	GetOIDCLogin(context.Context, *GetOIDCLoginRequest) (*GetOIDCLoginResponse, error)
	GetAuthToken(context.Context, *GetAuthTokenRequest) (*GetAuthTokenResponse, error)
	GetRobotToken(context.Context, *GetRobotTokenRequest) (*GetRobotTokenResponse, error)
	ExtendAuthToken(context.Context, *ExtendAuthTokenRequest) (*ExtendAuthTokenResponse, error)
	RevokeAuthToken(context.Context, *RevokeAuthTokenRequest) (*RevokeAuthTokenResponse, error)
	SetGroupsForUser(context.Context, *SetGroupsForUserRequest) (*SetGroupsForUserResponse, error)
	ModifyMembers(context.Context, *ModifyMembersRequest) (*ModifyMembersResponse, error)
	GetGroups(context.Context, *GetGroupsRequest) (*GetGroupsResponse, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	ExtractAuthTokens(context.Context, *ExtractAuthTokensRequest) (*ExtractAuthTokensResponse, error)
	RestoreAuthToken(context.Context, *RestoreAuthTokenRequest) (*RestoreAuthTokenResponse, error)
	mustEmbedUnimplementedAPIServer()
}

// UnimplementedAPIServer must be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (UnimplementedAPIServer) Activate(context.Context, *ActivateRequest) (*ActivateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}
func (UnimplementedAPIServer) Deactivate(context.Context, *DeactivateRequest) (*DeactivateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deactivate not implemented")
}
func (UnimplementedAPIServer) GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfiguration not implemented")
}
func (UnimplementedAPIServer) SetConfiguration(context.Context, *SetConfigurationRequest) (*SetConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfiguration not implemented")
}
func (UnimplementedAPIServer) Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedAPIServer) Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedAPIServer) WhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhoAmI not implemented")
}
func (UnimplementedAPIServer) ModifyRoleBinding(context.Context, *ModifyRoleBindingRequest) (*ModifyRoleBindingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyRoleBinding not implemented")
}
func (UnimplementedAPIServer) GetRoleBinding(context.Context, *GetRoleBindingRequest) (*GetRoleBindingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleBinding not implemented")
}
func (UnimplementedAPIServer) GetOIDCLogin(context.Context, *GetOIDCLoginRequest) (*GetOIDCLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOIDCLogin not implemented")
}
func (UnimplementedAPIServer) GetAuthToken(context.Context, *GetAuthTokenRequest) (*GetAuthTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthToken not implemented")
}
func (UnimplementedAPIServer) GetRobotToken(context.Context, *GetRobotTokenRequest) (*GetRobotTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRobotToken not implemented")
}
func (UnimplementedAPIServer) ExtendAuthToken(context.Context, *ExtendAuthTokenRequest) (*ExtendAuthTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtendAuthToken not implemented")
}
func (UnimplementedAPIServer) RevokeAuthToken(context.Context, *RevokeAuthTokenRequest) (*RevokeAuthTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeAuthToken not implemented")
}
func (UnimplementedAPIServer) SetGroupsForUser(context.Context, *SetGroupsForUserRequest) (*SetGroupsForUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGroupsForUser not implemented")
}
func (UnimplementedAPIServer) ModifyMembers(context.Context, *ModifyMembersRequest) (*ModifyMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyMembers not implemented")
}
func (UnimplementedAPIServer) GetGroups(context.Context, *GetGroupsRequest) (*GetGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroups not implemented")
}
func (UnimplementedAPIServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedAPIServer) ExtractAuthTokens(context.Context, *ExtractAuthTokensRequest) (*ExtractAuthTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtractAuthTokens not implemented")
}
func (UnimplementedAPIServer) RestoreAuthToken(context.Context, *RestoreAuthTokenRequest) (*RestoreAuthTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestoreAuthToken not implemented")
}
func (UnimplementedAPIServer) mustEmbedUnimplementedAPIServer() {}

// UnsafeAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServer will
// result in compilation errors.
type UnsafeAPIServer interface {
	mustEmbedUnimplementedAPIServer()
}

func RegisterAPIServer(s grpc.ServiceRegistrar, srv APIServer) {
	s.RegisterService(&API_ServiceDesc, srv)
}

func _API_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Activate(ctx, req.(*ActivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Deactivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Deactivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/Deactivate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Deactivate(ctx, req.(*DeactivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/GetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetConfiguration(ctx, req.(*GetConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_SetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).SetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/SetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).SetConfiguration(ctx, req.(*SetConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Authorize(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_WhoAmI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhoAmIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).WhoAmI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/WhoAmI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).WhoAmI(ctx, req.(*WhoAmIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ModifyRoleBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyRoleBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ModifyRoleBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/ModifyRoleBinding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ModifyRoleBinding(ctx, req.(*ModifyRoleBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetRoleBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetRoleBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/GetRoleBinding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetRoleBinding(ctx, req.(*GetRoleBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetOIDCLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOIDCLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetOIDCLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/GetOIDCLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetOIDCLogin(ctx, req.(*GetOIDCLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetAuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetAuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/GetAuthToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetAuthToken(ctx, req.(*GetAuthTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetRobotToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRobotTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetRobotToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/GetRobotToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetRobotToken(ctx, req.(*GetRobotTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ExtendAuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtendAuthTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ExtendAuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/ExtendAuthToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ExtendAuthToken(ctx, req.(*ExtendAuthTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_RevokeAuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeAuthTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).RevokeAuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/RevokeAuthToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).RevokeAuthToken(ctx, req.(*RevokeAuthTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_SetGroupsForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGroupsForUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).SetGroupsForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/SetGroupsForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).SetGroupsForUser(ctx, req.(*SetGroupsForUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ModifyMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ModifyMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/ModifyMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ModifyMembers(ctx, req.(*ModifyMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/GetGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetGroups(ctx, req.(*GetGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ExtractAuthTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtractAuthTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ExtractAuthTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/ExtractAuthTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ExtractAuthTokens(ctx, req.(*ExtractAuthTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_RestoreAuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreAuthTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).RestoreAuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.API/RestoreAuthToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).RestoreAuthToken(ctx, req.(*RestoreAuthTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// API_ServiceDesc is the grpc.ServiceDesc for API service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var API_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Activate",
			Handler:    _API_Activate_Handler,
		},
		{
			MethodName: "Deactivate",
			Handler:    _API_Deactivate_Handler,
		},
		{
			MethodName: "GetConfiguration",
			Handler:    _API_GetConfiguration_Handler,
		},
		{
			MethodName: "SetConfiguration",
			Handler:    _API_SetConfiguration_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _API_Authenticate_Handler,
		},
		{
			MethodName: "Authorize",
			Handler:    _API_Authorize_Handler,
		},
		{
			MethodName: "WhoAmI",
			Handler:    _API_WhoAmI_Handler,
		},
		{
			MethodName: "ModifyRoleBinding",
			Handler:    _API_ModifyRoleBinding_Handler,
		},
		{
			MethodName: "GetRoleBinding",
			Handler:    _API_GetRoleBinding_Handler,
		},
		{
			MethodName: "GetOIDCLogin",
			Handler:    _API_GetOIDCLogin_Handler,
		},
		{
			MethodName: "GetAuthToken",
			Handler:    _API_GetAuthToken_Handler,
		},
		{
			MethodName: "GetRobotToken",
			Handler:    _API_GetRobotToken_Handler,
		},
		{
			MethodName: "ExtendAuthToken",
			Handler:    _API_ExtendAuthToken_Handler,
		},
		{
			MethodName: "RevokeAuthToken",
			Handler:    _API_RevokeAuthToken_Handler,
		},
		{
			MethodName: "SetGroupsForUser",
			Handler:    _API_SetGroupsForUser_Handler,
		},
		{
			MethodName: "ModifyMembers",
			Handler:    _API_ModifyMembers_Handler,
		},
		{
			MethodName: "GetGroups",
			Handler:    _API_GetGroups_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _API_GetUsers_Handler,
		},
		{
			MethodName: "ExtractAuthTokens",
			Handler:    _API_ExtractAuthTokens_Handler,
		},
		{
			MethodName: "RestoreAuthToken",
			Handler:    _API_RestoreAuthToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}
