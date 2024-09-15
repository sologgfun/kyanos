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

type AgentOldConnEvtT struct {
	ConnInfo AgentOldConnInfoT
	ConnType AgentOldConnTypeT
	_        [4]byte
	Ts       uint64
}

type AgentOldConnIdS_t struct {
	TgidFd  uint64
	NoTrace bool
	_       [7]byte
}

type AgentOldConnInfoT struct {
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
	ReadBytes  uint64
	WriteBytes uint64
	Laddr      struct {
		In4 struct {
			SinFamily uint16
			SinPort   uint16
			SinAddr   struct{ S_addr uint32 }
			Pad       [8]uint8
		}
		_ [12]byte
	}
	Raddr struct {
		In4 struct {
			SinFamily uint16
			SinPort   uint16
			SinAddr   struct{ S_addr uint32 }
			Pad       [8]uint8
		}
		_ [12]byte
	}
	Protocol            AgentOldTrafficProtocolT
	Role                AgentOldEndpointRoleT
	PrevCount           uint64
	PrevBuf             [4]int8
	PrependLengthHeader bool
	NoTrace             bool
	_                   [2]byte
}

type AgentOldConnTypeT uint32

const (
	AgentOldConnTypeTKConnect       AgentOldConnTypeT = 0
	AgentOldConnTypeTKClose         AgentOldConnTypeT = 1
	AgentOldConnTypeTKProtocolInfer AgentOldConnTypeT = 2
)

type AgentOldControlValueIndexT uint32

const (
	AgentOldControlValueIndexTKTargetTGIDIndex   AgentOldControlValueIndexT = 0
	AgentOldControlValueIndexTKStirlingTGIDIndex AgentOldControlValueIndexT = 1
	AgentOldControlValueIndexTKEnabledXdpIndex   AgentOldControlValueIndexT = 2
	AgentOldControlValueIndexTKNumControlValues  AgentOldControlValueIndexT = 3
)

type AgentOldEndpointRoleT uint32

const (
	AgentOldEndpointRoleTKRoleClient  AgentOldEndpointRoleT = 1
	AgentOldEndpointRoleTKRoleServer  AgentOldEndpointRoleT = 2
	AgentOldEndpointRoleTKRoleUnknown AgentOldEndpointRoleT = 4
)

type AgentOldKernEvt struct {
	FuncName [16]int8
	Ts       uint64
	Seq      uint64
	Len      uint32
	Flags    uint8
	_        [3]byte
	ConnIdS  AgentOldConnIdS_t
	IsSample int32
	Step     AgentOldStepT
}

type AgentOldKernEvtData struct {
	Ke      AgentOldKernEvt
	BufSize uint32
	Msg     [30720]int8
	_       [4]byte
}

type AgentOldSockKey struct {
	Sip    uint32
	Dip    uint32
	Sport  uint32
	Dport  uint32
	Family uint32
}

type AgentOldStepT uint32

const (
	AgentOldStepTStart       AgentOldStepT = 0
	AgentOldStepTSYSCALL_OUT AgentOldStepT = 1
	AgentOldStepTTCP_OUT     AgentOldStepT = 2
	AgentOldStepTIP_OUT      AgentOldStepT = 3
	AgentOldStepTQDISC_OUT   AgentOldStepT = 4
	AgentOldStepTDEV_OUT     AgentOldStepT = 5
	AgentOldStepTNIC_OUT     AgentOldStepT = 6
	AgentOldStepTNIC_IN      AgentOldStepT = 7
	AgentOldStepTDEV_IN      AgentOldStepT = 8
	AgentOldStepTIP_IN       AgentOldStepT = 9
	AgentOldStepTTCP_IN      AgentOldStepT = 10
	AgentOldStepTUSER_COPY   AgentOldStepT = 11
	AgentOldStepTSYSCALL_IN  AgentOldStepT = 12
	AgentOldStepTEnd         AgentOldStepT = 13
)

type AgentOldTrafficDirectionT uint32

const (
	AgentOldTrafficDirectionTKEgress  AgentOldTrafficDirectionT = 0
	AgentOldTrafficDirectionTKIngress AgentOldTrafficDirectionT = 1
)

type AgentOldTrafficProtocolT uint32

const (
	AgentOldTrafficProtocolTKProtocolUnset   AgentOldTrafficProtocolT = 0
	AgentOldTrafficProtocolTKProtocolUnknown AgentOldTrafficProtocolT = 1
	AgentOldTrafficProtocolTKProtocolHTTP    AgentOldTrafficProtocolT = 2
	AgentOldTrafficProtocolTKProtocolHTTP2   AgentOldTrafficProtocolT = 3
	AgentOldTrafficProtocolTKProtocolMySQL   AgentOldTrafficProtocolT = 4
	AgentOldTrafficProtocolTKProtocolCQL     AgentOldTrafficProtocolT = 5
	AgentOldTrafficProtocolTKProtocolPGSQL   AgentOldTrafficProtocolT = 6
	AgentOldTrafficProtocolTKProtocolDNS     AgentOldTrafficProtocolT = 7
	AgentOldTrafficProtocolTKProtocolRedis   AgentOldTrafficProtocolT = 8
	AgentOldTrafficProtocolTKProtocolNATS    AgentOldTrafficProtocolT = 9
	AgentOldTrafficProtocolTKProtocolMongo   AgentOldTrafficProtocolT = 10
	AgentOldTrafficProtocolTKProtocolKafka   AgentOldTrafficProtocolT = 11
	AgentOldTrafficProtocolTKProtocolMux     AgentOldTrafficProtocolT = 12
	AgentOldTrafficProtocolTKProtocolAMQP    AgentOldTrafficProtocolT = 13
	AgentOldTrafficProtocolTKNumProtocols    AgentOldTrafficProtocolT = 14
)

// LoadAgentOld returns the embedded CollectionSpec for AgentOld.
func LoadAgentOld() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_AgentOldBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load AgentOld: %w", err)
	}

	return spec, err
}

// LoadAgentOldObjects loads AgentOld and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*AgentOldObjects
//	*AgentOldPrograms
//	*AgentOldMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadAgentOldObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadAgentOld()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// AgentOldSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type AgentOldSpecs struct {
	AgentOldProgramSpecs
	AgentOldMapSpecs
}

// AgentOldSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type AgentOldProgramSpecs struct {
	DevHardStartXmit                   *ebpf.ProgramSpec `ebpf:"dev_hard_start_xmit"`
	DevQueueXmit                       *ebpf.ProgramSpec `ebpf:"dev_queue_xmit"`
	IpQueueXmit                        *ebpf.ProgramSpec `ebpf:"ip_queue_xmit"`
	IpQueueXmit2                       *ebpf.ProgramSpec `ebpf:"ip_queue_xmit2"`
	IpRcvCore                          *ebpf.ProgramSpec `ebpf:"ip_rcv_core"`
	SecuritySocketRecvmsgEnter         *ebpf.ProgramSpec `ebpf:"security_socket_recvmsg_enter"`
	SecuritySocketSendmsgEnter         *ebpf.ProgramSpec `ebpf:"security_socket_sendmsg_enter"`
	SkbCopyDatagramIovec               *ebpf.ProgramSpec `ebpf:"skb_copy_datagram_iovec"`
	SkbCopyDatagramIter                *ebpf.ProgramSpec `ebpf:"skb_copy_datagram_iter"`
	SockAllocRet                       *ebpf.ProgramSpec `ebpf:"sock_alloc_ret"`
	TcpDestroySock                     *ebpf.ProgramSpec `ebpf:"tcp_destroy_sock"`
	TcpQueueRcv                        *ebpf.ProgramSpec `ebpf:"tcp_queue_rcv"`
	TcpRcvEstablished                  *ebpf.ProgramSpec `ebpf:"tcp_rcv_established"`
	TcpV4DoRcv                         *ebpf.ProgramSpec `ebpf:"tcp_v4_do_rcv"`
	TcpV4Rcv                           *ebpf.ProgramSpec `ebpf:"tcp_v4_rcv"`
	TracepointNetifReceiveSkb          *ebpf.ProgramSpec `ebpf:"tracepoint__netif_receive_skb"`
	TracepointSyscallsSysEnterAccept4  *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_accept4"`
	TracepointSyscallsSysEnterClose    *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_close"`
	TracepointSyscallsSysEnterConnect  *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_connect"`
	TracepointSyscallsSysEnterRead     *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_read"`
	TracepointSyscallsSysEnterReadv    *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_readv"`
	TracepointSyscallsSysEnterRecvfrom *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_recvfrom"`
	TracepointSyscallsSysEnterRecvmsg  *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_recvmsg"`
	TracepointSyscallsSysEnterSendmsg  *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_sendmsg"`
	TracepointSyscallsSysEnterSendto   *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_sendto"`
	TracepointSyscallsSysEnterWrite    *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_write"`
	TracepointSyscallsSysEnterWritev   *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_writev"`
	TracepointSyscallsSysExitAccept4   *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_accept4"`
	TracepointSyscallsSysExitClose     *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_close"`
	TracepointSyscallsSysExitConnect   *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_connect"`
	TracepointSyscallsSysExitRead      *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_read"`
	TracepointSyscallsSysExitReadv     *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_readv"`
	TracepointSyscallsSysExitRecvfrom  *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_recvfrom"`
	TracepointSyscallsSysExitRecvmsg   *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_recvmsg"`
	TracepointSyscallsSysExitSendmsg   *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_sendmsg"`
	TracepointSyscallsSysExitSendto    *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_sendto"`
	TracepointSyscallsSysExitWrite     *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_write"`
	TracepointSyscallsSysExitWritev    *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_writev"`
	XdpProxy                           *ebpf.ProgramSpec `ebpf:"xdp_proxy"`
}

// AgentOldMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type AgentOldMapSpecs struct {
	AcceptArgsMap        *ebpf.MapSpec `ebpf:"accept_args_map"`
	CloseArgsMap         *ebpf.MapSpec `ebpf:"close_args_map"`
	ConnEvtRb            *ebpf.MapSpec `ebpf:"conn_evt_rb"`
	ConnInfoMap          *ebpf.MapSpec `ebpf:"conn_info_map"`
	ConnInfoT_map        *ebpf.MapSpec `ebpf:"conn_info_t_map"`
	ConnectArgsMap       *ebpf.MapSpec `ebpf:"connect_args_map"`
	ControlValues        *ebpf.MapSpec `ebpf:"control_values"`
	EnabledLocalIpv4Map  *ebpf.MapSpec `ebpf:"enabled_local_ipv4_map"`
	EnabledLocalPortMap  *ebpf.MapSpec `ebpf:"enabled_local_port_map"`
	EnabledRemoteIpv4Map *ebpf.MapSpec `ebpf:"enabled_remote_ipv4_map"`
	EnabledRemotePortMap *ebpf.MapSpec `ebpf:"enabled_remote_port_map"`
	Rb                   *ebpf.MapSpec `ebpf:"rb"`
	ReadArgsMap          *ebpf.MapSpec `ebpf:"read_args_map"`
	SockKeyConnIdMap     *ebpf.MapSpec `ebpf:"sock_key_conn_id_map"`
	SockXmitMap          *ebpf.MapSpec `ebpf:"sock_xmit_map"`
	SyscallDataMap       *ebpf.MapSpec `ebpf:"syscall_data_map"`
	SyscallRb            *ebpf.MapSpec `ebpf:"syscall_rb"`
	WriteArgsMap         *ebpf.MapSpec `ebpf:"write_args_map"`
}

// AgentOldObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadAgentOldObjects or ebpf.CollectionSpec.LoadAndAssign.
type AgentOldObjects struct {
	AgentOldPrograms
	AgentOldMaps
}

func (o *AgentOldObjects) Close() error {
	return _AgentOldClose(
		&o.AgentOldPrograms,
		&o.AgentOldMaps,
	)
}

// AgentOldMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadAgentOldObjects or ebpf.CollectionSpec.LoadAndAssign.
type AgentOldMaps struct {
	AcceptArgsMap        *ebpf.Map `ebpf:"accept_args_map"`
	CloseArgsMap         *ebpf.Map `ebpf:"close_args_map"`
	ConnEvtRb            *ebpf.Map `ebpf:"conn_evt_rb"`
	ConnInfoMap          *ebpf.Map `ebpf:"conn_info_map"`
	ConnInfoT_map        *ebpf.Map `ebpf:"conn_info_t_map"`
	ConnectArgsMap       *ebpf.Map `ebpf:"connect_args_map"`
	ControlValues        *ebpf.Map `ebpf:"control_values"`
	EnabledLocalIpv4Map  *ebpf.Map `ebpf:"enabled_local_ipv4_map"`
	EnabledLocalPortMap  *ebpf.Map `ebpf:"enabled_local_port_map"`
	EnabledRemoteIpv4Map *ebpf.Map `ebpf:"enabled_remote_ipv4_map"`
	EnabledRemotePortMap *ebpf.Map `ebpf:"enabled_remote_port_map"`
	Rb                   *ebpf.Map `ebpf:"rb"`
	ReadArgsMap          *ebpf.Map `ebpf:"read_args_map"`
	SockKeyConnIdMap     *ebpf.Map `ebpf:"sock_key_conn_id_map"`
	SockXmitMap          *ebpf.Map `ebpf:"sock_xmit_map"`
	SyscallDataMap       *ebpf.Map `ebpf:"syscall_data_map"`
	SyscallRb            *ebpf.Map `ebpf:"syscall_rb"`
	WriteArgsMap         *ebpf.Map `ebpf:"write_args_map"`
}

func (m *AgentOldMaps) Close() error {
	return _AgentOldClose(
		m.AcceptArgsMap,
		m.CloseArgsMap,
		m.ConnEvtRb,
		m.ConnInfoMap,
		m.ConnInfoT_map,
		m.ConnectArgsMap,
		m.ControlValues,
		m.EnabledLocalIpv4Map,
		m.EnabledLocalPortMap,
		m.EnabledRemoteIpv4Map,
		m.EnabledRemotePortMap,
		m.Rb,
		m.ReadArgsMap,
		m.SockKeyConnIdMap,
		m.SockXmitMap,
		m.SyscallDataMap,
		m.SyscallRb,
		m.WriteArgsMap,
	)
}

// AgentOldPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadAgentOldObjects or ebpf.CollectionSpec.LoadAndAssign.
type AgentOldPrograms struct {
	DevHardStartXmit                   *ebpf.Program `ebpf:"dev_hard_start_xmit"`
	DevQueueXmit                       *ebpf.Program `ebpf:"dev_queue_xmit"`
	IpQueueXmit                        *ebpf.Program `ebpf:"ip_queue_xmit"`
	IpQueueXmit2                       *ebpf.Program `ebpf:"ip_queue_xmit2"`
	IpRcvCore                          *ebpf.Program `ebpf:"ip_rcv_core"`
	SecuritySocketRecvmsgEnter         *ebpf.Program `ebpf:"security_socket_recvmsg_enter"`
	SecuritySocketSendmsgEnter         *ebpf.Program `ebpf:"security_socket_sendmsg_enter"`
	SkbCopyDatagramIovec               *ebpf.Program `ebpf:"skb_copy_datagram_iovec"`
	SkbCopyDatagramIter                *ebpf.Program `ebpf:"skb_copy_datagram_iter"`
	SockAllocRet                       *ebpf.Program `ebpf:"sock_alloc_ret"`
	TcpDestroySock                     *ebpf.Program `ebpf:"tcp_destroy_sock"`
	TcpQueueRcv                        *ebpf.Program `ebpf:"tcp_queue_rcv"`
	TcpRcvEstablished                  *ebpf.Program `ebpf:"tcp_rcv_established"`
	TcpV4DoRcv                         *ebpf.Program `ebpf:"tcp_v4_do_rcv"`
	TcpV4Rcv                           *ebpf.Program `ebpf:"tcp_v4_rcv"`
	TracepointNetifReceiveSkb          *ebpf.Program `ebpf:"tracepoint__netif_receive_skb"`
	TracepointSyscallsSysEnterAccept4  *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_accept4"`
	TracepointSyscallsSysEnterClose    *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_close"`
	TracepointSyscallsSysEnterConnect  *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_connect"`
	TracepointSyscallsSysEnterRead     *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_read"`
	TracepointSyscallsSysEnterReadv    *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_readv"`
	TracepointSyscallsSysEnterRecvfrom *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_recvfrom"`
	TracepointSyscallsSysEnterRecvmsg  *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_recvmsg"`
	TracepointSyscallsSysEnterSendmsg  *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_sendmsg"`
	TracepointSyscallsSysEnterSendto   *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_sendto"`
	TracepointSyscallsSysEnterWrite    *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_write"`
	TracepointSyscallsSysEnterWritev   *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_writev"`
	TracepointSyscallsSysExitAccept4   *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_accept4"`
	TracepointSyscallsSysExitClose     *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_close"`
	TracepointSyscallsSysExitConnect   *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_connect"`
	TracepointSyscallsSysExitRead      *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_read"`
	TracepointSyscallsSysExitReadv     *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_readv"`
	TracepointSyscallsSysExitRecvfrom  *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_recvfrom"`
	TracepointSyscallsSysExitRecvmsg   *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_recvmsg"`
	TracepointSyscallsSysExitSendmsg   *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_sendmsg"`
	TracepointSyscallsSysExitSendto    *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_sendto"`
	TracepointSyscallsSysExitWrite     *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_write"`
	TracepointSyscallsSysExitWritev    *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_writev"`
	XdpProxy                           *ebpf.Program `ebpf:"xdp_proxy"`
}

func (p *AgentOldPrograms) Close() error {
	return _AgentOldClose(
		p.DevHardStartXmit,
		p.DevQueueXmit,
		p.IpQueueXmit,
		p.IpQueueXmit2,
		p.IpRcvCore,
		p.SecuritySocketRecvmsgEnter,
		p.SecuritySocketSendmsgEnter,
		p.SkbCopyDatagramIovec,
		p.SkbCopyDatagramIter,
		p.SockAllocRet,
		p.TcpDestroySock,
		p.TcpQueueRcv,
		p.TcpRcvEstablished,
		p.TcpV4DoRcv,
		p.TcpV4Rcv,
		p.TracepointNetifReceiveSkb,
		p.TracepointSyscallsSysEnterAccept4,
		p.TracepointSyscallsSysEnterClose,
		p.TracepointSyscallsSysEnterConnect,
		p.TracepointSyscallsSysEnterRead,
		p.TracepointSyscallsSysEnterReadv,
		p.TracepointSyscallsSysEnterRecvfrom,
		p.TracepointSyscallsSysEnterRecvmsg,
		p.TracepointSyscallsSysEnterSendmsg,
		p.TracepointSyscallsSysEnterSendto,
		p.TracepointSyscallsSysEnterWrite,
		p.TracepointSyscallsSysEnterWritev,
		p.TracepointSyscallsSysExitAccept4,
		p.TracepointSyscallsSysExitClose,
		p.TracepointSyscallsSysExitConnect,
		p.TracepointSyscallsSysExitRead,
		p.TracepointSyscallsSysExitReadv,
		p.TracepointSyscallsSysExitRecvfrom,
		p.TracepointSyscallsSysExitRecvmsg,
		p.TracepointSyscallsSysExitSendmsg,
		p.TracepointSyscallsSysExitSendto,
		p.TracepointSyscallsSysExitWrite,
		p.TracepointSyscallsSysExitWritev,
		p.XdpProxy,
	)
}

func _AgentOldClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed agentold_x86_bpfel.o
var _AgentOldBytes []byte
