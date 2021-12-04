package data

func convertTimestampsToMap(tss []uint32) map[uint32]struct{} {
	converted := make(map[uint32]struct{}, len(tss))

	for i := range tss {
		converted[tss[i]] = struct{}{}
	}

	return converted
}
