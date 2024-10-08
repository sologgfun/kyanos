// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package bpf

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type Openssl310ConnEvtT struct {
	ConnInfo Openssl310ConnInfoT
	ConnType Openssl310ConnTypeT
	_        [4]byte
	Ts       uint64
}

type Openssl310ConnIdS_t struct {
	TgidFd  uint64
	NoTrace bool
	_       [7]byte
}

type Openssl310ConnInfoT struct {
	ConnId struct {
		Upid struct {
			Pid            uint32
			_              [4]byte
			StartTimeTicks uint64
		}
		Fd   int32
		_    [4]byte
		Tsid uint64
	}
	ReadBytes     uint64
	WriteBytes    uint64
	SslReadBytes  uint64
	SslWriteBytes uint64
	Laddr         struct {
		In6 struct {
			Sin6Family   uint16
			Sin6Port     uint16
			Sin6Flowinfo uint32
			Sin6Addr     struct{ In6U struct{ U6Addr8 [16]uint8 } }
			Sin6ScopeId  uint32
		}
	}
	Raddr struct {
		In6 struct {
			Sin6Family   uint16
			Sin6Port     uint16
			Sin6Flowinfo uint32
			Sin6Addr     struct{ In6U struct{ U6Addr8 [16]uint8 } }
			Sin6ScopeId  uint32
		}
	}
	Protocol            Openssl310TrafficProtocolT
	Role                Openssl310EndpointRoleT
	PrevCount           uint64
	PrevBuf             [4]int8
	PrependLengthHeader bool
	NoTrace             bool
	Ssl                 bool
	_                   [1]byte
}

type Openssl310ConnTypeT uint32

const (
	Openssl310ConnTypeTKConnect       Openssl310ConnTypeT = 0
	Openssl310ConnTypeTKClose         Openssl310ConnTypeT = 1
	Openssl310ConnTypeTKProtocolInfer Openssl310ConnTypeT = 2
)

type Openssl310ControlValueIndexT uint32

const (
	Openssl310ControlValueIndexTKTargetTGIDIndex          Openssl310ControlValueIndexT = 0
	Openssl310ControlValueIndexTKStirlingTGIDIndex        Openssl310ControlValueIndexT = 1
	Openssl310ControlValueIndexTKEnabledXdpIndex          Openssl310ControlValueIndexT = 2
	Openssl310ControlValueIndexTKEnableFilterByPid        Openssl310ControlValueIndexT = 3
	Openssl310ControlValueIndexTKEnableFilterByLocalPort  Openssl310ControlValueIndexT = 4
	Openssl310ControlValueIndexTKEnableFilterByRemotePort Openssl310ControlValueIndexT = 5
	Openssl310ControlValueIndexTKEnableFilterByRemoteHost Openssl310ControlValueIndexT = 6
	Openssl310ControlValueIndexTKNumControlValues         Openssl310ControlValueIndexT = 7
)

type Openssl310EndpointRoleT uint32

const (
	Openssl310EndpointRoleTKRoleClient  Openssl310EndpointRoleT = 1
	Openssl310EndpointRoleTKRoleServer  Openssl310EndpointRoleT = 2
	Openssl310EndpointRoleTKRoleUnknown Openssl310EndpointRoleT = 4
)

type Openssl310KernEvt struct {
	FuncName [16]int8
	Ts       uint64
	Seq      uint64
	Len      uint32
	Flags    uint8
	_        [3]byte
	Ifindex  uint32
	_        [4]byte
	ConnIdS  Openssl310ConnIdS_t
	Step     Openssl310StepT
	_        [4]byte
}

type Openssl310KernEvtData struct {
	Ke      Openssl310KernEvt
	BufSize uint32
	Msg     [30720]int8
	_       [4]byte
}

type Openssl310SockKey struct {
	Sip   [2]uint64
	Dip   [2]uint64
	Sport uint16
	Dport uint16
	_     [4]byte
}

type Openssl310StepT uint32

const (
	Openssl310StepTStart       Openssl310StepT = 0
	Openssl310StepTSSL_OUT     Openssl310StepT = 1
	Openssl310StepTSYSCALL_OUT Openssl310StepT = 2
	Openssl310StepTTCP_OUT     Openssl310StepT = 3
	Openssl310StepTIP_OUT      Openssl310StepT = 4
	Openssl310StepTQDISC_OUT   Openssl310StepT = 5
	Openssl310StepTDEV_OUT     Openssl310StepT = 6
	Openssl310StepTNIC_OUT     Openssl310StepT = 7
	Openssl310StepTNIC_IN      Openssl310StepT = 8
	Openssl310StepTDEV_IN      Openssl310StepT = 9
	Openssl310StepTIP_IN       Openssl310StepT = 10
	Openssl310StepTTCP_IN      Openssl310StepT = 11
	Openssl310StepTUSER_COPY   Openssl310StepT = 12
	Openssl310StepTSYSCALL_IN  Openssl310StepT = 13
	Openssl310StepTSSL_IN      Openssl310StepT = 14
	Openssl310StepTEnd         Openssl310StepT = 15
)

type Openssl310TrafficDirectionT uint32

const (
	Openssl310TrafficDirectionTKEgress  Openssl310TrafficDirectionT = 0
	Openssl310TrafficDirectionTKIngress Openssl310TrafficDirectionT = 1
)

type Openssl310TrafficProtocolT uint32

const (
	Openssl310TrafficProtocolTKProtocolUnset   Openssl310TrafficProtocolT = 0
	Openssl310TrafficProtocolTKProtocolUnknown Openssl310TrafficProtocolT = 1
	Openssl310TrafficProtocolTKProtocolHTTP    Openssl310TrafficProtocolT = 2
	Openssl310TrafficProtocolTKProtocolHTTP2   Openssl310TrafficProtocolT = 3
	Openssl310TrafficProtocolTKProtocolMySQL   Openssl310TrafficProtocolT = 4
	Openssl310TrafficProtocolTKProtocolCQL     Openssl310TrafficProtocolT = 5
	Openssl310TrafficProtocolTKProtocolPGSQL   Openssl310TrafficProtocolT = 6
	Openssl310TrafficProtocolTKProtocolDNS     Openssl310TrafficProtocolT = 7
	Openssl310TrafficProtocolTKProtocolRedis   Openssl310TrafficProtocolT = 8
	Openssl310TrafficProtocolTKProtocolNATS    Openssl310TrafficProtocolT = 9
	Openssl310TrafficProtocolTKProtocolMongo   Openssl310TrafficProtocolT = 10
	Openssl310TrafficProtocolTKProtocolKafka   Openssl310TrafficProtocolT = 11
	Openssl310TrafficProtocolTKProtocolMux     Openssl310TrafficProtocolT = 12
	Openssl310TrafficProtocolTKProtocolAMQP    Openssl310TrafficProtocolT = 13
	Openssl310TrafficProtocolTKNumProtocols    Openssl310TrafficProtocolT = 14
)

// LoadOpenssl310 returns the embedded CollectionSpec for Openssl310.
func LoadOpenssl310() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Openssl310Bytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load Openssl310: %w", err)
	}

	return spec, err
}

// LoadOpenssl310Objects loads Openssl310 and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*Openssl310Objects
//	*Openssl310Programs
//	*Openssl310Maps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadOpenssl310Objects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadOpenssl310()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// Openssl310Specs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type Openssl310Specs struct {
	Openssl310ProgramSpecs
	Openssl310MapSpecs
}

// Openssl310Specs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type Openssl310ProgramSpecs struct {
	SSL_readEntryNestedSyscall    *ebpf.ProgramSpec `ebpf:"SSL_read_entry_nested_syscall"`
	SSL_readEntryOffset           *ebpf.ProgramSpec `ebpf:"SSL_read_entry_offset"`
	SSL_readExEntryNestedSyscall  *ebpf.ProgramSpec `ebpf:"SSL_read_ex_entry_nested_syscall"`
	SSL_readExRetNestedSyscall    *ebpf.ProgramSpec `ebpf:"SSL_read_ex_ret_nested_syscall"`
	SSL_readRetNestedSyscall      *ebpf.ProgramSpec `ebpf:"SSL_read_ret_nested_syscall"`
	SSL_readRetOffset             *ebpf.ProgramSpec `ebpf:"SSL_read_ret_offset"`
	SSL_writeEntryNestedSyscall   *ebpf.ProgramSpec `ebpf:"SSL_write_entry_nested_syscall"`
	SSL_writeEntryOffset          *ebpf.ProgramSpec `ebpf:"SSL_write_entry_offset"`
	SSL_writeExEntryNestedSyscall *ebpf.ProgramSpec `ebpf:"SSL_write_ex_entry_nested_syscall"`
	SSL_writeExRetNestedSyscall   *ebpf.ProgramSpec `ebpf:"SSL_write_ex_ret_nested_syscall"`
	SSL_writeRetNestedSyscall     *ebpf.ProgramSpec `ebpf:"SSL_write_ret_nested_syscall"`
	SSL_writeRetOffset            *ebpf.ProgramSpec `ebpf:"SSL_write_ret_offset"`
}

// Openssl310MapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type Openssl310MapSpecs struct {
	ActiveSslReadArgsMap  *ebpf.MapSpec `ebpf:"active_ssl_read_args_map"`
	ActiveSslWriteArgsMap *ebpf.MapSpec `ebpf:"active_ssl_write_args_map"`
	ConnEvtRb             *ebpf.MapSpec `ebpf:"conn_evt_rb"`
	ConnInfoMap           *ebpf.MapSpec `ebpf:"conn_info_map"`
	FilterMntnsMap        *ebpf.MapSpec `ebpf:"filter_mntns_map"`
	FilterNetnsMap        *ebpf.MapSpec `ebpf:"filter_netns_map"`
	FilterPidMap          *ebpf.MapSpec `ebpf:"filter_pid_map"`
	FilterPidnsMap        *ebpf.MapSpec `ebpf:"filter_pidns_map"`
	Rb                    *ebpf.MapSpec `ebpf:"rb"`
	SslDataMap            *ebpf.MapSpec `ebpf:"ssl_data_map"`
	SslRb                 *ebpf.MapSpec `ebpf:"ssl_rb"`
	SslUserSpaceCallMap   *ebpf.MapSpec `ebpf:"ssl_user_space_call_map"`
	SyscallDataMap        *ebpf.MapSpec `ebpf:"syscall_data_map"`
	SyscallRb             *ebpf.MapSpec `ebpf:"syscall_rb"`
}

// Openssl310Objects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadOpenssl310Objects or ebpf.CollectionSpec.LoadAndAssign.
type Openssl310Objects struct {
	Openssl310Programs
	Openssl310Maps
}

func (o *Openssl310Objects) Close() error {
	return _Openssl310Close(
		&o.Openssl310Programs,
		&o.Openssl310Maps,
	)
}

// Openssl310Maps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadOpenssl310Objects or ebpf.CollectionSpec.LoadAndAssign.
type Openssl310Maps struct {
	ActiveSslReadArgsMap  *ebpf.Map `ebpf:"active_ssl_read_args_map"`
	ActiveSslWriteArgsMap *ebpf.Map `ebpf:"active_ssl_write_args_map"`
	ConnEvtRb             *ebpf.Map `ebpf:"conn_evt_rb"`
	ConnInfoMap           *ebpf.Map `ebpf:"conn_info_map"`
	FilterMntnsMap        *ebpf.Map `ebpf:"filter_mntns_map"`
	FilterNetnsMap        *ebpf.Map `ebpf:"filter_netns_map"`
	FilterPidMap          *ebpf.Map `ebpf:"filter_pid_map"`
	FilterPidnsMap        *ebpf.Map `ebpf:"filter_pidns_map"`
	Rb                    *ebpf.Map `ebpf:"rb"`
	SslDataMap            *ebpf.Map `ebpf:"ssl_data_map"`
	SslRb                 *ebpf.Map `ebpf:"ssl_rb"`
	SslUserSpaceCallMap   *ebpf.Map `ebpf:"ssl_user_space_call_map"`
	SyscallDataMap        *ebpf.Map `ebpf:"syscall_data_map"`
	SyscallRb             *ebpf.Map `ebpf:"syscall_rb"`
}

func (m *Openssl310Maps) Close() error {
	return _Openssl310Close(
		m.ActiveSslReadArgsMap,
		m.ActiveSslWriteArgsMap,
		m.ConnEvtRb,
		m.ConnInfoMap,
		m.FilterMntnsMap,
		m.FilterNetnsMap,
		m.FilterPidMap,
		m.FilterPidnsMap,
		m.Rb,
		m.SslDataMap,
		m.SslRb,
		m.SslUserSpaceCallMap,
		m.SyscallDataMap,
		m.SyscallRb,
	)
}

// Openssl310Programs contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadOpenssl310Objects or ebpf.CollectionSpec.LoadAndAssign.
type Openssl310Programs struct {
	SSL_readEntryNestedSyscall    *ebpf.Program `ebpf:"SSL_read_entry_nested_syscall"`
	SSL_readEntryOffset           *ebpf.Program `ebpf:"SSL_read_entry_offset"`
	SSL_readExEntryNestedSyscall  *ebpf.Program `ebpf:"SSL_read_ex_entry_nested_syscall"`
	SSL_readExRetNestedSyscall    *ebpf.Program `ebpf:"SSL_read_ex_ret_nested_syscall"`
	SSL_readRetNestedSyscall      *ebpf.Program `ebpf:"SSL_read_ret_nested_syscall"`
	SSL_readRetOffset             *ebpf.Program `ebpf:"SSL_read_ret_offset"`
	SSL_writeEntryNestedSyscall   *ebpf.Program `ebpf:"SSL_write_entry_nested_syscall"`
	SSL_writeEntryOffset          *ebpf.Program `ebpf:"SSL_write_entry_offset"`
	SSL_writeExEntryNestedSyscall *ebpf.Program `ebpf:"SSL_write_ex_entry_nested_syscall"`
	SSL_writeExRetNestedSyscall   *ebpf.Program `ebpf:"SSL_write_ex_ret_nested_syscall"`
	SSL_writeRetNestedSyscall     *ebpf.Program `ebpf:"SSL_write_ret_nested_syscall"`
	SSL_writeRetOffset            *ebpf.Program `ebpf:"SSL_write_ret_offset"`
}

func (p *Openssl310Programs) Close() error {
	return _Openssl310Close(
		p.SSL_readEntryNestedSyscall,
		p.SSL_readEntryOffset,
		p.SSL_readExEntryNestedSyscall,
		p.SSL_readExRetNestedSyscall,
		p.SSL_readRetNestedSyscall,
		p.SSL_readRetOffset,
		p.SSL_writeEntryNestedSyscall,
		p.SSL_writeEntryOffset,
		p.SSL_writeExEntryNestedSyscall,
		p.SSL_writeExRetNestedSyscall,
		p.SSL_writeRetNestedSyscall,
		p.SSL_writeRetOffset,
	)
}

func _Openssl310Close(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed openssl310_x86_bpfel.o
var _Openssl310Bytes []byte