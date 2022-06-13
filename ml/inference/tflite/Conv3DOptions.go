// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Conv3DOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsConv3DOptions(buf []byte, offset flatbuffers.UOffsetT) *Conv3DOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Conv3DOptions{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsConv3DOptions(buf []byte, offset flatbuffers.UOffsetT) *Conv3DOptions {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Conv3DOptions{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Conv3DOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Conv3DOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Conv3DOptions) Padding() Padding {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Padding(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Conv3DOptions) MutatePadding(n Padding) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func (rcv *Conv3DOptions) StrideD() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Conv3DOptions) MutateStrideD(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *Conv3DOptions) StrideW() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Conv3DOptions) MutateStrideW(n int32) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func (rcv *Conv3DOptions) StrideH() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Conv3DOptions) MutateStrideH(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

func (rcv *Conv3DOptions) FusedActivationFunction() ActivationFunctionType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return ActivationFunctionType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Conv3DOptions) MutateFusedActivationFunction(n ActivationFunctionType) bool {
	return rcv._tab.MutateInt8Slot(12, int8(n))
}

func (rcv *Conv3DOptions) DilationDFactor() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 1
}

func (rcv *Conv3DOptions) MutateDilationDFactor(n int32) bool {
	return rcv._tab.MutateInt32Slot(14, n)
}

func (rcv *Conv3DOptions) DilationWFactor() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 1
}

func (rcv *Conv3DOptions) MutateDilationWFactor(n int32) bool {
	return rcv._tab.MutateInt32Slot(16, n)
}

func (rcv *Conv3DOptions) DilationHFactor() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 1
}

func (rcv *Conv3DOptions) MutateDilationHFactor(n int32) bool {
	return rcv._tab.MutateInt32Slot(18, n)
}

func Conv3DOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(8)
}
func Conv3DOptionsAddPadding(builder *flatbuffers.Builder, padding Padding) {
	builder.PrependInt8Slot(0, int8(padding), 0)
}
func Conv3DOptionsAddStrideD(builder *flatbuffers.Builder, strideD int32) {
	builder.PrependInt32Slot(1, strideD, 0)
}
func Conv3DOptionsAddStrideW(builder *flatbuffers.Builder, strideW int32) {
	builder.PrependInt32Slot(2, strideW, 0)
}
func Conv3DOptionsAddStrideH(builder *flatbuffers.Builder, strideH int32) {
	builder.PrependInt32Slot(3, strideH, 0)
}
func Conv3DOptionsAddFusedActivationFunction(builder *flatbuffers.Builder, fusedActivationFunction ActivationFunctionType) {
	builder.PrependInt8Slot(4, int8(fusedActivationFunction), 0)
}
func Conv3DOptionsAddDilationDFactor(builder *flatbuffers.Builder, dilationDFactor int32) {
	builder.PrependInt32Slot(5, dilationDFactor, 1)
}
func Conv3DOptionsAddDilationWFactor(builder *flatbuffers.Builder, dilationWFactor int32) {
	builder.PrependInt32Slot(6, dilationWFactor, 1)
}
func Conv3DOptionsAddDilationHFactor(builder *flatbuffers.Builder, dilationHFactor int32) {
	builder.PrependInt32Slot(7, dilationHFactor, 1)
}
func Conv3DOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}