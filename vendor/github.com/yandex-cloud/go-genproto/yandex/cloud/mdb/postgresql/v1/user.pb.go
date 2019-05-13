// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/postgresql/v1/user.proto

package postgresql // import "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/postgresql/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserSettings_SynchronousCommit int32

const (
	UserSettings_SYNCHRONOUS_COMMIT_UNSPECIFIED  UserSettings_SynchronousCommit = 0
	UserSettings_SYNCHRONOUS_COMMIT_ON           UserSettings_SynchronousCommit = 1
	UserSettings_SYNCHRONOUS_COMMIT_OFF          UserSettings_SynchronousCommit = 2
	UserSettings_SYNCHRONOUS_COMMIT_LOCAL        UserSettings_SynchronousCommit = 3
	UserSettings_SYNCHRONOUS_COMMIT_REMOTE_WRITE UserSettings_SynchronousCommit = 4
	UserSettings_SYNCHRONOUS_COMMIT_REMOTE_APPLY UserSettings_SynchronousCommit = 5
)

var UserSettings_SynchronousCommit_name = map[int32]string{
	0: "SYNCHRONOUS_COMMIT_UNSPECIFIED",
	1: "SYNCHRONOUS_COMMIT_ON",
	2: "SYNCHRONOUS_COMMIT_OFF",
	3: "SYNCHRONOUS_COMMIT_LOCAL",
	4: "SYNCHRONOUS_COMMIT_REMOTE_WRITE",
	5: "SYNCHRONOUS_COMMIT_REMOTE_APPLY",
}
var UserSettings_SynchronousCommit_value = map[string]int32{
	"SYNCHRONOUS_COMMIT_UNSPECIFIED":  0,
	"SYNCHRONOUS_COMMIT_ON":           1,
	"SYNCHRONOUS_COMMIT_OFF":          2,
	"SYNCHRONOUS_COMMIT_LOCAL":        3,
	"SYNCHRONOUS_COMMIT_REMOTE_WRITE": 4,
	"SYNCHRONOUS_COMMIT_REMOTE_APPLY": 5,
}

func (x UserSettings_SynchronousCommit) String() string {
	return proto.EnumName(UserSettings_SynchronousCommit_name, int32(x))
}
func (UserSettings_SynchronousCommit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_86f81a23ffeb3ce2, []int{3, 0}
}

type UserSettings_LogStatement int32

const (
	UserSettings_LOG_STATEMENT_UNSPECIFIED UserSettings_LogStatement = 0
	UserSettings_LOG_STATEMENT_NONE        UserSettings_LogStatement = 1
	UserSettings_LOG_STATEMENT_DDL         UserSettings_LogStatement = 2
	UserSettings_LOG_STATEMENT_MOD         UserSettings_LogStatement = 3
	UserSettings_LOG_STATEMENT_ALL         UserSettings_LogStatement = 4
)

var UserSettings_LogStatement_name = map[int32]string{
	0: "LOG_STATEMENT_UNSPECIFIED",
	1: "LOG_STATEMENT_NONE",
	2: "LOG_STATEMENT_DDL",
	3: "LOG_STATEMENT_MOD",
	4: "LOG_STATEMENT_ALL",
}
var UserSettings_LogStatement_value = map[string]int32{
	"LOG_STATEMENT_UNSPECIFIED": 0,
	"LOG_STATEMENT_NONE":        1,
	"LOG_STATEMENT_DDL":         2,
	"LOG_STATEMENT_MOD":         3,
	"LOG_STATEMENT_ALL":         4,
}

func (x UserSettings_LogStatement) String() string {
	return proto.EnumName(UserSettings_LogStatement_name, int32(x))
}
func (UserSettings_LogStatement) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_86f81a23ffeb3ce2, []int{3, 1}
}

type UserSettings_TransactionIsolation int32

const (
	UserSettings_TRANSACTION_ISOLATION_UNSPECIFIED      UserSettings_TransactionIsolation = 0
	UserSettings_TRANSACTION_ISOLATION_READ_UNCOMMITTED UserSettings_TransactionIsolation = 1
	UserSettings_TRANSACTION_ISOLATION_READ_COMMITTED   UserSettings_TransactionIsolation = 2
	UserSettings_TRANSACTION_ISOLATION_REPEATABLE_READ  UserSettings_TransactionIsolation = 3
	UserSettings_TRANSACTION_ISOLATION_SERIALIZABLE     UserSettings_TransactionIsolation = 4
)

var UserSettings_TransactionIsolation_name = map[int32]string{
	0: "TRANSACTION_ISOLATION_UNSPECIFIED",
	1: "TRANSACTION_ISOLATION_READ_UNCOMMITTED",
	2: "TRANSACTION_ISOLATION_READ_COMMITTED",
	3: "TRANSACTION_ISOLATION_REPEATABLE_READ",
	4: "TRANSACTION_ISOLATION_SERIALIZABLE",
}
var UserSettings_TransactionIsolation_value = map[string]int32{
	"TRANSACTION_ISOLATION_UNSPECIFIED":      0,
	"TRANSACTION_ISOLATION_READ_UNCOMMITTED": 1,
	"TRANSACTION_ISOLATION_READ_COMMITTED":   2,
	"TRANSACTION_ISOLATION_REPEATABLE_READ":  3,
	"TRANSACTION_ISOLATION_SERIALIZABLE":     4,
}

func (x UserSettings_TransactionIsolation) String() string {
	return proto.EnumName(UserSettings_TransactionIsolation_name, int32(x))
}
func (UserSettings_TransactionIsolation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_86f81a23ffeb3ce2, []int{3, 2}
}

// A PostgreSQL User resource. For more information, see
// the [Developer's Guide](/docs/managed-postgresql/concepts).
type User struct {
	// Name of the PostgreSQL user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ID of the PostgreSQL cluster the user belongs to.
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Set of permissions granted to the user.
	Permissions []*Permission `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// Number of database connections available to the user.
	ConnLimit int64 `protobuf:"varint,4,opt,name=conn_limit,json=connLimit,proto3" json:"conn_limit,omitempty"`
	// Postgresql settings for this user
	Settings             *UserSettings `protobuf:"bytes,5,opt,name=settings,proto3" json:"settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_86f81a23ffeb3ce2, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *User) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *User) GetConnLimit() int64 {
	if m != nil {
		return m.ConnLimit
	}
	return 0
}

func (m *User) GetSettings() *UserSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

type Permission struct {
	// Name of the database that the permission grants access to.
	DatabaseName         string   `protobuf:"bytes,1,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_86f81a23ffeb3ce2, []int{1}
}
func (m *Permission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission.Unmarshal(m, b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
}
func (dst *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(dst, src)
}
func (m *Permission) XXX_Size() int {
	return xxx_messageInfo_Permission.Size(m)
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

func (m *Permission) GetDatabaseName() string {
	if m != nil {
		return m.DatabaseName
	}
	return ""
}

type UserSpec struct {
	// Name of the PostgreSQL user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Password of the PostgreSQL user.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// Set of permissions to grant to the user.
	Permissions []*Permission `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// Number of database connections that should be available to the user.
	ConnLimit *wrappers.Int64Value `protobuf:"bytes,4,opt,name=conn_limit,json=connLimit,proto3" json:"conn_limit,omitempty"`
	// Postgresql settings for this user
	Settings             *UserSettings `protobuf:"bytes,5,opt,name=settings,proto3" json:"settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UserSpec) Reset()         { *m = UserSpec{} }
func (m *UserSpec) String() string { return proto.CompactTextString(m) }
func (*UserSpec) ProtoMessage()    {}
func (*UserSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_86f81a23ffeb3ce2, []int{2}
}
func (m *UserSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSpec.Unmarshal(m, b)
}
func (m *UserSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSpec.Marshal(b, m, deterministic)
}
func (dst *UserSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSpec.Merge(dst, src)
}
func (m *UserSpec) XXX_Size() int {
	return xxx_messageInfo_UserSpec.Size(m)
}
func (m *UserSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UserSpec proto.InternalMessageInfo

func (m *UserSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserSpec) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserSpec) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *UserSpec) GetConnLimit() *wrappers.Int64Value {
	if m != nil {
		return m.ConnLimit
	}
	return nil
}

func (m *UserSpec) GetSettings() *UserSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

// Postgresql user settings config
type UserSettings struct {
	DefaultTransactionIsolation UserSettings_TransactionIsolation `protobuf:"varint,1,opt,name=default_transaction_isolation,json=defaultTransactionIsolation,proto3,enum=yandex.cloud.mdb.postgresql.v1.UserSettings_TransactionIsolation" json:"default_transaction_isolation,omitempty"`
	// in milliseconds.
	LockTimeout *wrappers.Int64Value `protobuf:"bytes,2,opt,name=lock_timeout,json=lockTimeout,proto3" json:"lock_timeout,omitempty"`
	// in milliseconds.
	LogMinDurationStatement *wrappers.Int64Value           `protobuf:"bytes,3,opt,name=log_min_duration_statement,json=logMinDurationStatement,proto3" json:"log_min_duration_statement,omitempty"`
	SynchronousCommit       UserSettings_SynchronousCommit `protobuf:"varint,4,opt,name=synchronous_commit,json=synchronousCommit,proto3,enum=yandex.cloud.mdb.postgresql.v1.UserSettings_SynchronousCommit" json:"synchronous_commit,omitempty"`
	// in bytes.
	TempFileLimit        *wrappers.Int64Value `protobuf:"bytes,5,opt,name=temp_file_limit,json=tempFileLimit,proto3" json:"temp_file_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UserSettings) Reset()         { *m = UserSettings{} }
func (m *UserSettings) String() string { return proto.CompactTextString(m) }
func (*UserSettings) ProtoMessage()    {}
func (*UserSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_86f81a23ffeb3ce2, []int{3}
}
func (m *UserSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSettings.Unmarshal(m, b)
}
func (m *UserSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSettings.Marshal(b, m, deterministic)
}
func (dst *UserSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSettings.Merge(dst, src)
}
func (m *UserSettings) XXX_Size() int {
	return xxx_messageInfo_UserSettings.Size(m)
}
func (m *UserSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSettings.DiscardUnknown(m)
}

var xxx_messageInfo_UserSettings proto.InternalMessageInfo

func (m *UserSettings) GetDefaultTransactionIsolation() UserSettings_TransactionIsolation {
	if m != nil {
		return m.DefaultTransactionIsolation
	}
	return UserSettings_TRANSACTION_ISOLATION_UNSPECIFIED
}

func (m *UserSettings) GetLockTimeout() *wrappers.Int64Value {
	if m != nil {
		return m.LockTimeout
	}
	return nil
}

func (m *UserSettings) GetLogMinDurationStatement() *wrappers.Int64Value {
	if m != nil {
		return m.LogMinDurationStatement
	}
	return nil
}

func (m *UserSettings) GetSynchronousCommit() UserSettings_SynchronousCommit {
	if m != nil {
		return m.SynchronousCommit
	}
	return UserSettings_SYNCHRONOUS_COMMIT_UNSPECIFIED
}

func (m *UserSettings) GetTempFileLimit() *wrappers.Int64Value {
	if m != nil {
		return m.TempFileLimit
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "yandex.cloud.mdb.postgresql.v1.User")
	proto.RegisterType((*Permission)(nil), "yandex.cloud.mdb.postgresql.v1.Permission")
	proto.RegisterType((*UserSpec)(nil), "yandex.cloud.mdb.postgresql.v1.UserSpec")
	proto.RegisterType((*UserSettings)(nil), "yandex.cloud.mdb.postgresql.v1.UserSettings")
	proto.RegisterEnum("yandex.cloud.mdb.postgresql.v1.UserSettings_SynchronousCommit", UserSettings_SynchronousCommit_name, UserSettings_SynchronousCommit_value)
	proto.RegisterEnum("yandex.cloud.mdb.postgresql.v1.UserSettings_LogStatement", UserSettings_LogStatement_name, UserSettings_LogStatement_value)
	proto.RegisterEnum("yandex.cloud.mdb.postgresql.v1.UserSettings_TransactionIsolation", UserSettings_TransactionIsolation_name, UserSettings_TransactionIsolation_value)
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/postgresql/v1/user.proto", fileDescriptor_user_86f81a23ffeb3ce2)
}

var fileDescriptor_user_86f81a23ffeb3ce2 = []byte{
	// 848 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdf, 0x6e, 0xdb, 0x54,
	0x18, 0xc7, 0x49, 0x8a, 0xd2, 0x2f, 0xed, 0x70, 0x8f, 0xd8, 0xc8, 0x3a, 0x52, 0x4a, 0xc6, 0xa6,
	0xb4, 0x22, 0xf6, 0x9c, 0xa1, 0x69, 0x08, 0x56, 0xc9, 0x49, 0x5c, 0x66, 0xc9, 0xb1, 0x23, 0xdb,
	0x05, 0x56, 0x84, 0x8e, 0x9c, 0xf8, 0xd4, 0xb3, 0xb0, 0x7d, 0x82, 0xcf, 0x71, 0xc7, 0xb8, 0xe7,
	0xa6, 0xaf, 0xc3, 0x3b, 0xd0, 0x3d, 0x01, 0x8f, 0x80, 0xb8, 0x44, 0x5c, 0x72, 0x85, 0xe2, 0x24,
	0x4d, 0xbb, 0x84, 0x54, 0x93, 0xd8, 0x9d, 0xfd, 0xfb, 0xf3, 0x7d, 0xc7, 0xbf, 0xf3, 0x93, 0x0c,
	0x7b, 0x2f, 0xbd, 0xc4, 0x27, 0x3f, 0xc9, 0xc3, 0x88, 0x66, 0xbe, 0x1c, 0xfb, 0x03, 0x79, 0x44,
	0x19, 0x0f, 0x52, 0xc2, 0x7e, 0x8c, 0xe4, 0x53, 0x45, 0xce, 0x18, 0x49, 0xa5, 0x51, 0x4a, 0x39,
	0x45, 0x3b, 0x13, 0xa9, 0x94, 0x4b, 0xa5, 0xd8, 0x1f, 0x48, 0x73, 0xa9, 0x74, 0xaa, 0x6c, 0xef,
	0x04, 0x94, 0x06, 0x11, 0x91, 0x73, 0xf5, 0x20, 0x3b, 0x91, 0x5f, 0xa4, 0xde, 0x68, 0x44, 0x52,
	0x36, 0xf1, 0x6f, 0xd7, 0xae, 0xac, 0x3a, 0xf5, 0xa2, 0xd0, 0xf7, 0x78, 0x48, 0x93, 0x09, 0x5d,
	0xff, 0x4b, 0x80, 0xd2, 0x11, 0x23, 0x29, 0x42, 0x50, 0x4a, 0xbc, 0x98, 0x54, 0x85, 0x5d, 0xa1,
	0xb1, 0x6e, 0xe7, 0xcf, 0xa8, 0x06, 0x30, 0x8c, 0x32, 0xc6, 0x49, 0x8a, 0x43, 0xbf, 0x5a, 0xc8,
	0x99, 0xf5, 0x29, 0xa2, 0xfb, 0xc8, 0x80, 0xca, 0x88, 0xa4, 0x71, 0xc8, 0x58, 0x48, 0x13, 0x56,
	0x2d, 0xee, 0x16, 0x1b, 0x95, 0xd6, 0xbe, 0xb4, 0xfa, 0xc0, 0x52, 0xff, 0xc2, 0x62, 0x5f, 0xb6,
	0xe7, 0xcb, 0x68, 0x92, 0xe0, 0x28, 0x8c, 0x43, 0x5e, 0x2d, 0xed, 0x0a, 0x8d, 0xa2, 0xbd, 0x3e,
	0x46, 0x8c, 0x31, 0x80, 0x9e, 0x42, 0x99, 0x11, 0xce, 0xc3, 0x24, 0x60, 0xd5, 0xb5, 0x5d, 0xa1,
	0x51, 0x69, 0x7d, 0x7a, 0xdd, 0xa6, 0xf1, 0x77, 0x39, 0x53, 0x8f, 0x7d, 0xe1, 0xae, 0x2b, 0x00,
	0xf3, 0x33, 0xa0, 0xbb, 0xb0, 0xe9, 0x7b, 0xdc, 0x1b, 0x78, 0x8c, 0xe0, 0x4b, 0x01, 0x6c, 0xcc,
	0x40, 0xd3, 0x8b, 0x49, 0xfd, 0xb7, 0x02, 0x94, 0xf3, 0x69, 0x23, 0x32, 0x44, 0xca, 0xe5, 0xa4,
	0xda, 0xb5, 0x3f, 0xcf, 0x15, 0xe1, 0xef, 0x73, 0x65, 0xf3, 0x3b, 0xaf, 0xf9, 0xb3, 0xda, 0x3c,
	0x7e, 0xd0, 0xfc, 0x1c, 0x7f, 0xbf, 0x7f, 0xf6, 0x4a, 0x29, 0x7d, 0xf9, 0xe4, 0xd1, 0xc3, 0x69,
	0x90, 0x7b, 0x50, 0x1e, 0x79, 0x8c, 0xbd, 0xa0, 0xe9, 0x34, 0xc6, 0xf6, 0xe6, 0xd8, 0x76, 0xf6,
	0x4a, 0x59, 0x7b, 0xdc, 0x54, 0x5a, 0x8f, 0xed, 0x0b, 0xfa, 0x7f, 0x0e, 0xb5, 0xbb, 0x10, 0x6a,
	0xa5, 0x75, 0x47, 0x9a, 0x54, 0x46, 0x9a, 0x55, 0x46, 0xd2, 0x13, 0xfe, 0xe8, 0xb3, 0xaf, 0xbd,
	0x28, 0x23, 0xed, 0xf2, 0x3f, 0xe7, 0x4a, 0xe9, 0xe0, 0x89, 0xf2, 0xe0, 0xed, 0x64, 0xff, 0x6b,
	0x19, 0x36, 0x2e, 0x53, 0xe8, 0x17, 0x01, 0x6a, 0x3e, 0x39, 0xf1, 0xb2, 0x88, 0x63, 0x9e, 0x7a,
	0x09, 0xf3, 0x86, 0xe3, 0x76, 0xe2, 0x90, 0xd1, 0x28, 0xef, 0x69, 0x1e, 0xf3, 0x8d, 0x96, 0xfa,
	0x26, 0x0b, 0x25, 0x77, 0x3e, 0x49, 0x9f, 0x0d, 0xb2, 0xef, 0x4c, 0xf7, 0x2c, 0x23, 0xd1, 0x01,
	0x6c, 0x44, 0x74, 0xf8, 0x03, 0xe6, 0x61, 0x4c, 0x68, 0xc6, 0xf3, 0x5b, 0x5a, 0x1d, 0x95, 0x5d,
	0x19, 0x1b, 0xdc, 0x89, 0x1e, 0x7d, 0x0b, 0xdb, 0x11, 0x0d, 0x70, 0x1c, 0x26, 0xd8, 0xcf, 0xd2,
	0x7c, 0x26, 0x66, 0xdc, 0xe3, 0x24, 0x26, 0x09, 0xaf, 0x16, 0xaf, 0x9f, 0xf6, 0x41, 0x44, 0x83,
	0x5e, 0x98, 0x74, 0xa7, 0x66, 0x67, 0xe6, 0x45, 0x31, 0x20, 0xf6, 0x32, 0x19, 0x3e, 0x4f, 0x69,
	0x42, 0x33, 0x86, 0x87, 0x34, 0x9e, 0x5d, 0xe5, 0x8d, 0xd6, 0xc1, 0x1b, 0xa5, 0xe2, 0xcc, 0xc7,
	0x74, 0xf2, 0x29, 0xf6, 0x16, 0x7b, 0x1d, 0x42, 0x1d, 0x78, 0x8f, 0x93, 0x78, 0x84, 0x4f, 0xc2,
	0x88, 0x4c, 0x6b, 0xb3, 0x76, 0xfd, 0xe9, 0x37, 0xc7, 0x9e, 0xc3, 0x30, 0x22, 0x79, 0x61, 0xea,
	0xbf, 0x0b, 0xb0, 0xb5, 0xb0, 0x0d, 0xd5, 0x61, 0xc7, 0x79, 0x66, 0x76, 0x9e, 0xda, 0x96, 0x69,
	0x1d, 0x39, 0xb8, 0x63, 0xf5, 0x7a, 0xba, 0x8b, 0x8f, 0x4c, 0xa7, 0xaf, 0x75, 0xf4, 0x43, 0x5d,
	0xeb, 0x8a, 0xef, 0xa0, 0xdb, 0x70, 0x73, 0x89, 0xc6, 0x32, 0x45, 0x01, 0x6d, 0xc3, 0xad, 0x65,
	0xd4, 0xe1, 0xa1, 0x58, 0x40, 0x1f, 0x42, 0x75, 0x09, 0x67, 0x58, 0x1d, 0xd5, 0x10, 0x8b, 0xe8,
	0x2e, 0x7c, 0xb4, 0x84, 0xb5, 0xb5, 0x9e, 0xe5, 0x6a, 0xf8, 0x1b, 0x5b, 0x77, 0x35, 0xb1, 0xb4,
	0x5a, 0xa4, 0xf6, 0xfb, 0xc6, 0x33, 0x71, 0xad, 0x7e, 0x26, 0xc0, 0x86, 0x41, 0x83, 0xf9, 0xed,
	0xd4, 0xe0, 0xb6, 0x61, 0x7d, 0x85, 0x1d, 0x57, 0x75, 0xb5, 0x9e, 0x66, 0xbe, 0xfe, 0x39, 0xb7,
	0x00, 0x5d, 0xa5, 0x4d, 0xcb, 0xd4, 0x44, 0x01, 0xdd, 0x84, 0xad, 0xab, 0x78, 0xb7, 0x6b, 0x88,
	0x85, 0x45, 0xb8, 0x67, 0x75, 0xc5, 0xe2, 0x22, 0xac, 0x1a, 0x86, 0x58, 0xaa, 0xff, 0x21, 0xc0,
	0xfb, 0x4b, 0xcb, 0x7c, 0x0f, 0x3e, 0x76, 0x6d, 0xd5, 0x74, 0xd4, 0x8e, 0xab, 0x5b, 0x26, 0xd6,
	0x1d, 0xcb, 0x50, 0xf3, 0xa7, 0xab, 0x87, 0xdb, 0x87, 0xfb, 0xcb, 0x65, 0xb6, 0xa6, 0x76, 0xf1,
	0x91, 0x39, 0x89, 0xc0, 0xd5, 0xba, 0xa2, 0x80, 0x1a, 0xf0, 0xc9, 0x0a, 0xed, 0x5c, 0x59, 0x40,
	0x7b, 0x70, 0xef, 0xbf, 0x94, 0x7d, 0x4d, 0x75, 0xd5, 0xb6, 0xa1, 0xe5, 0x26, 0xb1, 0x88, 0xee,
	0x43, 0x7d, 0xb9, 0xd4, 0xd1, 0x6c, 0x5d, 0x35, 0xf4, 0xe3, 0xb1, 0x58, 0x2c, 0xb5, 0xad, 0xe3,
	0x5e, 0x10, 0xf2, 0xe7, 0xd9, 0x40, 0x1a, 0xd2, 0x58, 0x9e, 0x54, 0xbe, 0x39, 0xf9, 0xa1, 0x05,
	0xb4, 0x19, 0x90, 0x24, 0xaf, 0xa4, 0xbc, 0xfa, 0xa7, 0xfa, 0xc5, 0xfc, 0x6d, 0xf0, 0x6e, 0x6e,
	0x78, 0xf8, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x72, 0x36, 0x79, 0x88, 0x07, 0x00, 0x00,
}