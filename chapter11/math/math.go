package math

// find the average of a series of numbers
func Average(xs ... float32)float32  {
	total := float32(0)
	for _, v := range xs{
		total+=v
	}
	return total / float32(len(xs))
}

// find the max from a slice
func Max(xs ... float32)float32  {
	mx := float32(0)
	for _, v := range xs{
		if v > mx {
			mx = v
		}
	}
	return mx
}

// find the minimum from a slice.
// it is also a variadic function
func Min(xs ... float32)float32  {
	mn := float32(9999999)
	for _, v := range xs{
		if v < mn{
			mn = v
		}
	}
	return  mn
}