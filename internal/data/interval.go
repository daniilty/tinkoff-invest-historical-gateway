package data

type Interval struct {
	Step uint32
	From uint32
	To   uint32
}

func GetMissingIntervals(timestamps []uint32, on *Interval) []*Interval {
	missing := getMissingIntervals(timestamps, on)
	joined := joinIntervals(missing)

	return joined
}

func joinIntervals(intervals []*Interval) []*Interval {
	const minLen = 2

	if len(intervals) < minLen {
		return intervals
	}

	joined := []*Interval{
		intervals[0],
	}

	for i := 1; i < len(intervals); i++ {
		lastJoined := joined[len(joined)-1]

		if isNextTimestamp(lastJoined.To, intervals[i].To, intervals[i].Step) {
			joined[len(joined)-1].To = intervals[i].To

			continue
		}

		joined = append(joined, intervals[i])
	}

	return joined
}

func getMissingIntervals(timestamps []uint32, on *Interval) []*Interval {
	const minLen = 1

	if len(timestamps) < minLen {
		return []*Interval{on}
	}

	if on.Step == 0 {
		panic("sent 0 interval step")
	}

	from, to := getFromTo(timestamps)
	missing := getLeadingAndTrailingIntervals(from, to, on)

	timestampsMap := convertTimestampsToMap(timestamps)

	for from <= to {
		_, ok := timestampsMap[from]
		if ok {
			from += on.Step

			continue
		}

		if to-from < on.Step {
			break
		}

		// if we need to download only one piece of data
		// we need to set the last interval at that value
		missing = append(missing, &Interval{
			Step: on.Step,
			From: from - on.Step,
			To:   from,
		})

		from += on.Step
	}

	return missing
}

func isNextTimestamp(currentTs uint32, nextTs uint32, interval uint32) bool {
	return nextTs-currentTs == interval
}

func getLeadingAndTrailingIntervals(from uint32, to uint32, on *Interval) []*Interval {
	intervals := []*Interval{}

	// if <from> inside interval - append it
	if from-on.From >= on.Step {
		intervals = append(intervals, &Interval{
			From: on.From,
			To:   from,
			Step: on.Step,
		})
	}

	// same logic
	if on.To-to >= on.Step {
		intervals = append(intervals, &Interval{
			From: to,
			To:   on.To,
			Step: on.Step,
		})
	}

	return intervals
}

func getFromTo(timestamps []uint32) (uint32, uint32) {
	const minLen = 2

	if len(timestamps) < minLen {
		return 0, 0
	}

	return timestamps[0], timestamps[len(timestamps)-1]
}
