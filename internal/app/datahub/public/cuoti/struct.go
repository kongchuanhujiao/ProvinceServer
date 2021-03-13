package cuoti

type Tab struct {
	ID             uint32 `json:"id" db:"id"`                             // 唯一标识符
	UserID         uint32 `json:"user_id" db:"user_id"`                   // 用户 QQ 号
	Question       string `json:"question" db:"question"`                 // 错题题目
	ImageURL       string `json:"image_url" db:"image_url"`               // 错题图片 URL
	QuestionAnswer string `json:"question_answer" db:"question_answer"`   // 问题正确答案
	Duration       uint8  `json:"notice_duration" db:"notice_duration"`   // 提醒周期, 单位天
	LastNoticeTime uint64 `json:"last_notice_time" db:"last_notice_time"` // 上次提醒时间
}
