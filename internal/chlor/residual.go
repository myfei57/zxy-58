package chlor

type Residual struct {
	Free  float64
	Total float64
}

func ComputeResidual(dose float64, flowRate float64) Residual {
	free := dose / (flowRate + 1)
	return Residual{Free: free, Total: free * 1.2}
}

func (r Residual) InRange(low, high float64) bool {
	return r.Free >= low && r.Free <= high
}
