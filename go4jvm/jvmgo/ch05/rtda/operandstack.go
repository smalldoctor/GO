package rtda

import "math"

type OperandStack struct {
	/*
	通过数组模拟栈，因为数组的容量已经确定，所以可以通过下标索引进行模拟;
	*/
	size  uint
	slots []Slot
}

func newOperandStack(maxStacks uint) *OperandStack {
	if maxStacks > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStacks),
		}
	}
	return nil
}

func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++
}

func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}

func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}

func (self *OperandStack) PopFloat() float32 {
	self.size--
	val := uint32(self.slots[self.size].num)
	return math.Float32frombits(val)
}

func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size].num = int32(val)
	self.size++
	self.slots[self.size].num = int32(val >> 32)
	self.size++
}

func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint64(self.slots[self.size].num)
	high := uint64(self.slots[self.size+1].num)
	return int64(high<<32) | int64(low)
}

func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}

func (self *OperandStack) PopDouble() float64 {
	long := uint64(self.PopLong())
	return math.Float64frombits(long)
}

func (self *OperandStack) PushRef(ref *Object) {
	self.slots[self.size].ref = ref
	self.size++
}
func (self *OperandStack) PopRef() *Object {
	self.size--
	return self.slots[self.size].ref
}

func (self *OperandStack) PushSlot(slot Slot) {
	self.slots[self.size] = slot
	self.size++
}

func (self *OperandStack) PopSlot() Slot {
	self.size--
	return self.slots[self.size]
}
