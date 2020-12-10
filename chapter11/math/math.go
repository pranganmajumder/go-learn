package math

func Average(xs ... float32)float32  {
	total := float32(0)
	for _, v := range xs{
		total+=v
	}
	return total / float32(len(xs))
}