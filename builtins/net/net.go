// +build !appengine

// Package net implements net interface for anko script.
package net

import (
	pkg "net"

	"github.com/covrom/gonec/vm"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("net")
	m.DefineS("CIDRMask", pkg.CIDRMask)
	m.DefineS("Dial", pkg.Dial)
	m.DefineS("DialIP", pkg.DialIP)
	m.DefineS("DialTCP", pkg.DialTCP)
	m.DefineS("DialTimeout", pkg.DialTimeout)
	m.DefineS("DialUDP", pkg.DialUDP)
	m.DefineS("DialUnix", pkg.DialUnix)
	m.DefineS("ErrWriteToConnected", pkg.ErrWriteToConnected)
	m.DefineS("FileConn", pkg.FileConn)
	m.DefineS("FileListener", pkg.FileListener)
	m.DefineS("FilePacketConn", pkg.FilePacketConn)
	m.DefineS("FlagBroadcast", pkg.FlagBroadcast)
	m.DefineS("FlagLoopback", pkg.FlagLoopback)
	m.DefineS("FlagMulticast", pkg.FlagMulticast)
	m.DefineS("FlagPointToPoint", pkg.FlagPointToPoint)
	m.DefineS("FlagUp", pkg.FlagUp)
	m.DefineS("IPv4", pkg.IPv4)
	m.DefineS("IPv4Mask", pkg.IPv4Mask)
	m.DefineS("IPv4allrouter", pkg.IPv4allrouter)
	m.DefineS("IPv4allsys", pkg.IPv4allsys)
	m.DefineS("IPv4bcast", pkg.IPv4bcast)
	m.DefineS("IPv4len", pkg.IPv4len)
	m.DefineS("IPv4zero", pkg.IPv4zero)
	m.DefineS("IPv6interfacelocalallnodes", pkg.IPv6interfacelocalallnodes)
	m.DefineS("IPv6len", pkg.IPv6len)
	m.DefineS("IPv6linklocalallnodes", pkg.IPv6linklocalallnodes)
	m.DefineS("IPv6linklocalallrouters", pkg.IPv6linklocalallrouters)
	m.DefineS("IPv6loopback", pkg.IPv6loopback)
	m.DefineS("IPv6unspecified", pkg.IPv6unspecified)
	m.DefineS("IPv6zero", pkg.IPv6zero)
	m.DefineS("InterfaceAddrs", pkg.InterfaceAddrs)
	m.DefineS("InterfaceByIndex", pkg.InterfaceByIndex)
	m.DefineS("InterfaceByName", pkg.InterfaceByName)
	m.DefineS("Interfaces", pkg.Interfaces)
	m.DefineS("JoinHostPort", pkg.JoinHostPort)
	m.DefineS("Listen", pkg.Listen)
	m.DefineS("ListenIP", pkg.ListenIP)
	m.DefineS("ListenMulticastUDP", pkg.ListenMulticastUDP)
	m.DefineS("ListenPacket", pkg.ListenPacket)
	m.DefineS("ListenTCP", pkg.ListenTCP)
	m.DefineS("ListenUDP", pkg.ListenUDP)
	m.DefineS("ListenUnix", pkg.ListenUnix)
	m.DefineS("ListenUnixgram", pkg.ListenUnixgram)
	m.DefineS("LookupAddr", pkg.LookupAddr)
	m.DefineS("LookupCNAME", pkg.LookupCNAME)
	m.DefineS("LookupHost", pkg.LookupHost)
	m.DefineS("LookupIP", pkg.LookupIP)
	m.DefineS("LookupMX", pkg.LookupMX)
	m.DefineS("LookupNS", pkg.LookupNS)
	m.DefineS("LookupPort", pkg.LookupPort)
	m.DefineS("LookupSRV", pkg.LookupSRV)
	m.DefineS("LookupTXT", pkg.LookupTXT)
	m.DefineS("ParseCIDR", pkg.ParseCIDR)
	m.DefineS("ParseIP", pkg.ParseIP)
	m.DefineS("ParseMAC", pkg.ParseMAC)
	m.DefineS("Pipe", pkg.Pipe)
	m.DefineS("ResolveIPAddr", pkg.ResolveIPAddr)
	m.DefineS("ResolveTCPAddr", pkg.ResolveTCPAddr)
	m.DefineS("ResolveUDPAddr", pkg.ResolveUDPAddr)
	m.DefineS("ResolveUnixAddr", pkg.ResolveUnixAddr)
	m.DefineS("SplitHostPort", pkg.SplitHostPort)
	return m
}
