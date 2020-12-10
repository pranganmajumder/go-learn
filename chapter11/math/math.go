package math

func Average(xs ... float32)float32  {
	total := float32(0)
	for _, v := range xs{
		total+=v
	}
	return total / float32(len(xs))
}

func Max(xs ... float32)float32  {
	mx := float32(0)
	for _, v := range xs{
		if v > mx {
			mx = v
		}
	}
	return mx
}

func Min(xs ... float32)float32  {
	mn := float32(9999999)
	for _, v := range xs{
		if v < mn{
			mn = v
		}
	}
	return  mn
}