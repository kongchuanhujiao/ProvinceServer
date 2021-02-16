package wenda

func CalculateQuestion(w *Detail) (calc *CalculationsTab) {

	var (
		rightStus []uint64
		wrongStus [][]uint64
	)

	rightStus = []uint64{}
	wrongStus = [][]uint64{}

	correctAnswer := w.Questions.Key

	for _, ans := range w.Answers {
		if ans.Answer == correctAnswer {
			rightStus = append(rightStus, ans.QQ)
		} else {
			for i, option := range w.Questions.Options {
				if option != ans.Answer {
					continue
				}

				op := wrongStus[i]
				if op == nil {
					op = []uint64{}
				}

				op = append(op, ans.QQ)
				wrongStus[i] = op
			}
		}
	}

	calc = &CalculationsTab{
		Question:    w.Questions.ID,
		AnswerCount: uint8(len(w.Answers)),
		Right:       rightStus,
		Wrong:       wrongStus,
	}

	return
}