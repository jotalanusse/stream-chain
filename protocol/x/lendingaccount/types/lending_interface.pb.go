// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dydxprotocol/lendingaccount/lending_interface.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type LendingInterface struct {
	// The name of the lending manager.
	ManagerName string `protobuf:"bytes,1,opt,name=manager_name,json=managerName,proto3" json:"manager_name,omitempty"`
	// The base denomination for the lending interface.
	BaseDenom string `protobuf:"bytes,2,opt,name=base_denom,json=baseDenom,proto3" json:"base_denom,omitempty"`
	// The maximum amount that can be borrowed per block.
	MaxBorrowPerBlock string `protobuf:"bytes,3,opt,name=max_borrow_per_block,json=maxBorrowPerBlock,proto3" json:"max_borrow_per_block,omitempty"`
	// The minimum amount that can be borrowed per account.
	MinBorrowPerAccount string `protobuf:"bytes,4,opt,name=min_borrow_per_account,json=minBorrowPerAccount,proto3" json:"min_borrow_per_account,omitempty"`
	// The maximum amount that can be borrowed per account.
	MaxBorrowPerAccount string `protobuf:"bytes,5,opt,name=max_borrow_per_account,json=maxBorrowPerAccount,proto3" json:"max_borrow_per_account,omitempty"`
	// The cumulative loss for the lending interface.
	CumulativeLoss string `protobuf:"bytes,6,opt,name=cumulative_loss,json=cumulativeLoss,proto3" json:"cumulative_loss,omitempty"`
	// The limit for the cumulative loss.
	CumulativeLossLimit string `protobuf:"bytes,7,opt,name=cumulative_loss_limit,json=cumulativeLossLimit,proto3" json:"cumulative_loss_limit,omitempty"`
	// The total debt for the lending interface.
	TotalDebt string `protobuf:"bytes,8,opt,name=total_debt,json=totalDebt,proto3" json:"total_debt,omitempty"`
	// The limit for the total debt.
	TotalDebtLimit string `protobuf:"bytes,9,opt,name=total_debt_limit,json=totalDebtLimit,proto3" json:"total_debt_limit,omitempty"`
}

func (m *LendingInterface) Reset()         { *m = LendingInterface{} }
func (m *LendingInterface) String() string { return proto.CompactTextString(m) }
func (*LendingInterface) ProtoMessage()    {}
func (*LendingInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2533b4832dc1102, []int{0}
}
func (m *LendingInterface) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LendingInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LendingInterface.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LendingInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LendingInterface.Merge(m, src)
}
func (m *LendingInterface) XXX_Size() int {
	return m.Size()
}
func (m *LendingInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_LendingInterface.DiscardUnknown(m)
}

var xxx_messageInfo_LendingInterface proto.InternalMessageInfo

func (m *LendingInterface) GetManagerName() string {
	if m != nil {
		return m.ManagerName
	}
	return ""
}

func (m *LendingInterface) GetBaseDenom() string {
	if m != nil {
		return m.BaseDenom
	}
	return ""
}

func (m *LendingInterface) GetMaxBorrowPerBlock() string {
	if m != nil {
		return m.MaxBorrowPerBlock
	}
	return ""
}

func (m *LendingInterface) GetMinBorrowPerAccount() string {
	if m != nil {
		return m.MinBorrowPerAccount
	}
	return ""
}

func (m *LendingInterface) GetMaxBorrowPerAccount() string {
	if m != nil {
		return m.MaxBorrowPerAccount
	}
	return ""
}

func (m *LendingInterface) GetCumulativeLoss() string {
	if m != nil {
		return m.CumulativeLoss
	}
	return ""
}

func (m *LendingInterface) GetCumulativeLossLimit() string {
	if m != nil {
		return m.CumulativeLossLimit
	}
	return ""
}

func (m *LendingInterface) GetTotalDebt() string {
	if m != nil {
		return m.TotalDebt
	}
	return ""
}

func (m *LendingInterface) GetTotalDebtLimit() string {
	if m != nil {
		return m.TotalDebtLimit
	}
	return ""
}

func init() {
	proto.RegisterType((*LendingInterface)(nil), "dydxprotocol.lendingaccount.LendingInterface")
}

func init() {
	proto.RegisterFile("dydxprotocol/lendingaccount/lending_interface.proto", fileDescriptor_b2533b4832dc1102)
}

var fileDescriptor_b2533b4832dc1102 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcd, 0x4e, 0xf2, 0x40,
	0x14, 0x86, 0xe9, 0xc7, 0x27, 0xca, 0x68, 0x14, 0x2b, 0x9a, 0x46, 0x63, 0xa3, 0x6e, 0x64, 0x03,
	0x4d, 0xe4, 0x0a, 0x24, 0xc4, 0xc4, 0x84, 0x10, 0xa2, 0x3b, 0x37, 0x93, 0xe9, 0xf4, 0x58, 0x26,
	0x76, 0x66, 0xc8, 0x74, 0xaa, 0xe5, 0x2e, 0xdc, 0x78, 0x4f, 0x2e, 0x59, 0xba, 0x34, 0x70, 0x23,
	0xa6, 0xd3, 0x96, 0xbf, 0x1d, 0xbc, 0xef, 0xf3, 0x9c, 0xd3, 0x4c, 0x0e, 0xea, 0x06, 0xd3, 0x20,
	0x9d, 0x28, 0xa9, 0x25, 0x95, 0x91, 0x17, 0x81, 0x08, 0x98, 0x08, 0x09, 0xa5, 0x32, 0x11, 0xba,
	0xfc, 0x8b, 0x99, 0xd0, 0xa0, 0x5e, 0x09, 0x85, 0x8e, 0x21, 0xed, 0x8b, 0x75, 0xa9, 0xb3, 0x29,
	0x9d, 0x37, 0x43, 0x19, 0x4a, 0x53, 0x7a, 0xd9, 0xaf, 0x5c, 0xb9, 0xf9, 0xaa, 0xa2, 0xc6, 0x20,
	0x07, 0x1f, 0xcb, 0x69, 0xf6, 0x35, 0x3a, 0xe0, 0x44, 0x90, 0x10, 0x14, 0x16, 0x84, 0x83, 0x63,
	0x5d, 0x59, 0xad, 0xfa, 0xd3, 0x7e, 0x91, 0x0d, 0x09, 0x07, 0xfb, 0x12, 0x21, 0x9f, 0xc4, 0x80,
	0x03, 0x10, 0x92, 0x3b, 0xff, 0x0c, 0x50, 0xcf, 0x92, 0x7e, 0x16, 0xd8, 0x1e, 0x6a, 0x72, 0x92,
	0x62, 0x5f, 0x2a, 0x25, 0x3f, 0xf0, 0x04, 0x14, 0xf6, 0x23, 0x49, 0xdf, 0x9c, 0xaa, 0x01, 0x8f,
	0x39, 0x49, 0x7b, 0xa6, 0x1a, 0x81, 0xea, 0x65, 0x85, 0xdd, 0x45, 0x67, 0x9c, 0x89, 0x75, 0xa1,
	0xf8, 0x6e, 0xe7, 0xbf, 0x51, 0x4e, 0x38, 0x13, 0x4b, 0xe5, 0x3e, 0xaf, 0x8c, 0xb4, 0xb9, 0xa5,
	0x94, 0x76, 0x0a, 0x69, 0x6d, 0x4f, 0x29, 0xdd, 0xa2, 0x23, 0x9a, 0xf0, 0x24, 0x22, 0x9a, 0xbd,
	0x03, 0x8e, 0x64, 0x1c, 0x3b, 0x35, 0x43, 0x1f, 0xae, 0xe2, 0x81, 0x8c, 0x63, 0xfb, 0x0e, 0x9d,
	0x6e, 0x81, 0x38, 0x62, 0x9c, 0x69, 0x67, 0x37, 0x1f, 0xbe, 0x89, 0x0f, 0xb2, 0x2a, 0x7b, 0x16,
	0x2d, 0x35, 0x89, 0x70, 0x00, 0xbe, 0x76, 0xf6, 0xf2, 0x67, 0x31, 0x49, 0x1f, 0x7c, 0x6d, 0xb7,
	0x50, 0x63, 0x55, 0x17, 0xd3, 0xea, 0xf9, 0xf2, 0x25, 0x64, 0x06, 0xf5, 0xc6, 0xdf, 0x73, 0xd7,
	0x9a, 0xcd, 0x5d, 0xeb, 0x77, 0xee, 0x5a, 0x9f, 0x0b, 0xb7, 0x32, 0x5b, 0xb8, 0x95, 0x9f, 0x85,
	0x5b, 0x79, 0x19, 0x86, 0x4c, 0x8f, 0x13, 0xbf, 0x43, 0x25, 0xf7, 0x9e, 0xb5, 0x02, 0xc2, 0x1f,
	0x98, 0x20, 0x82, 0x42, 0x7b, 0x54, 0x9e, 0x4b, 0x6c, 0xe2, 0x36, 0x1d, 0x13, 0x26, 0xbc, 0xe5,
	0x11, 0xa5, 0xdb, 0x67, 0xa4, 0xa7, 0x13, 0x88, 0xfd, 0x9a, 0x01, 0xba, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x16, 0xfb, 0x44, 0x2e, 0x72, 0x02, 0x00, 0x00,
}

func (m *LendingInterface) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LendingInterface) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LendingInterface) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TotalDebtLimit) > 0 {
		i -= len(m.TotalDebtLimit)
		copy(dAtA[i:], m.TotalDebtLimit)
		i = encodeVarintLendingInterface(dAtA, i, uint64(len(m.TotalDebtLimit)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.TotalDebt) > 0 {
		i -= len(m.TotalDebt)
		copy(dAtA[i:], m.TotalDebt)
		i = encodeVarintLendingInterface(dAtA, i, uint64(len(m.TotalDebt)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.CumulativeLossLimit) > 0 {
		i -= len(m.CumulativeLossLimit)
		copy(dAtA[i:], m.CumulativeLossLimit)
		i = encodeVarintLendingInterface(dAtA, i, uint64(len(m.CumulativeLossLimit)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.CumulativeLoss) > 0 {
		i -= len(m.CumulativeLoss)
		copy(dAtA[i:], m.CumulativeLoss)
		i = encodeVarintLendingInterface(dAtA, i, uint64(len(m.CumulativeLoss)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.MaxBorrowPerAccount) > 0 {
		i -= len(m.MaxBorrowPerAccount)
		copy(dAtA[i:], m.MaxBorrowPerAccount)
		i = encodeVarintLendingInterface(dAtA, i, uint64(len(m.MaxBorrowPerAccount)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.MinBorrowPerAccount) > 0 {
		i -= len(m.MinBorrowPerAccount)
		copy(dAtA[i:], m.MinBorrowPerAccount)
		i = encodeVarintLendingInterface(dAtA, i, uint64(len(m.MinBorrowPerAccount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.MaxBorrowPerBlock) > 0 {
		i -= len(m.MaxBorrowPerBlock)
		copy(dAtA[i:], m.MaxBorrowPerBlock)
		i = encodeVarintLendingInterface(dAtA, i, uint64(len(m.MaxBorrowPerBlock)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BaseDenom) > 0 {
		i -= len(m.BaseDenom)
		copy(dAtA[i:], m.BaseDenom)
		i = encodeVarintLendingInterface(dAtA, i, uint64(len(m.BaseDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ManagerName) > 0 {
		i -= len(m.ManagerName)
		copy(dAtA[i:], m.ManagerName)
		i = encodeVarintLendingInterface(dAtA, i, uint64(len(m.ManagerName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLendingInterface(dAtA []byte, offset int, v uint64) int {
	offset -= sovLendingInterface(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LendingInterface) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ManagerName)
	if l > 0 {
		n += 1 + l + sovLendingInterface(uint64(l))
	}
	l = len(m.BaseDenom)
	if l > 0 {
		n += 1 + l + sovLendingInterface(uint64(l))
	}
	l = len(m.MaxBorrowPerBlock)
	if l > 0 {
		n += 1 + l + sovLendingInterface(uint64(l))
	}
	l = len(m.MinBorrowPerAccount)
	if l > 0 {
		n += 1 + l + sovLendingInterface(uint64(l))
	}
	l = len(m.MaxBorrowPerAccount)
	if l > 0 {
		n += 1 + l + sovLendingInterface(uint64(l))
	}
	l = len(m.CumulativeLoss)
	if l > 0 {
		n += 1 + l + sovLendingInterface(uint64(l))
	}
	l = len(m.CumulativeLossLimit)
	if l > 0 {
		n += 1 + l + sovLendingInterface(uint64(l))
	}
	l = len(m.TotalDebt)
	if l > 0 {
		n += 1 + l + sovLendingInterface(uint64(l))
	}
	l = len(m.TotalDebtLimit)
	if l > 0 {
		n += 1 + l + sovLendingInterface(uint64(l))
	}
	return n
}

func sovLendingInterface(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLendingInterface(x uint64) (n int) {
	return sovLendingInterface(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LendingInterface) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLendingInterface
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LendingInterface: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LendingInterface: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ManagerName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLendingInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLendingInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ManagerName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLendingInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLendingInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBorrowPerBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLendingInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLendingInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxBorrowPerBlock = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBorrowPerAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLendingInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLendingInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinBorrowPerAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBorrowPerAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLendingInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLendingInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxBorrowPerAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CumulativeLoss", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLendingInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLendingInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CumulativeLoss = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CumulativeLossLimit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLendingInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLendingInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CumulativeLossLimit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalDebt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLendingInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLendingInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalDebt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalDebtLimit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLendingInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLendingInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLendingInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalDebtLimit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLendingInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLendingInterface
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLendingInterface(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLendingInterface
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLendingInterface
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLendingInterface
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLendingInterface
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLendingInterface
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLendingInterface
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLendingInterface        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLendingInterface          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLendingInterface = fmt.Errorf("proto: unexpected end of group")
)