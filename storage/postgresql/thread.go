package postgresql

import (
	"github.com/GabrielBG99/vxchan/domain"
	"gorm.io/gorm"
)

type (
	Reply struct {
		gorm.Model
		ThreadID   int64
		Thread     Thread
		Content    string
		Attachment string
	}

	Thread struct {
		gorm.Model
		BoardID     string
		Board       Board `gorm:"references:tag"`
		Description string
		Attachment  string
		Replies     []Reply
	}
)

func (r Reply) ToDomain() domain.Reply {
	return domain.Reply{
		ID:         int64(r.ID),
		CreatedAt:  r.CreatedAt,
		ThreadID:   int64(r.ThreadID),
		Content:    r.Content,
		Attachment: domain.Attachment(r.Attachment),
	}
}

func (t Thread) ToDomain() domain.Thread {
	replies := make([]domain.Reply, 0)
	for _, r := range t.Replies {
		replies = append(replies, r.ToDomain())
	}

	return domain.Thread{
		ID:          int64(t.ID),
		CreatedAt:   t.CreatedAt,
		BoardID:     t.BoardID,
		Description: t.Description,
		Attachment:  domain.Attachment(t.Attachment),
		Replies:     replies,
	}
}

func (c connector) initThread() error {
	if err := c.db.AutoMigrate(&Thread{}); err != nil {
		return err
	}

	return c.db.AutoMigrate(&Reply{})
}

func (c connector) CreateThread(boardID, content string, attachment domain.Attachment) (domain.Thread, error) {
	t := Thread{
		BoardID:     boardID,
		Description: content,
		Attachment:  string(attachment),
	}

	r := c.db.Create(&t)
	if err := r.Error; err != nil {
		return domain.Thread{}, err
	}

	thread := t.ToDomain()
	return thread, nil
}

func (c connector) ListThreads(boardID string, page, nItems int) ([]domain.Thread, error) {
	var data []Thread
	r := c.db.Model(&Thread{}).Offset(page * nItems).Limit(nItems).Where(&Thread{BoardID: boardID}).Find(&data)
	if err := r.Error; err != nil {
		return nil, err
	}

	threads := make([]domain.Thread, 0)
	for _, d := range data {
		threads = append(threads, d.ToDomain())
	}

	return threads, nil
}

func (c connector) ReplyThread(threadID int64, comment string, attachment domain.Attachment) (domain.Reply, error) {
	rp := Reply{
		ThreadID:   threadID,
		Content:    comment,
		Attachment: string(attachment),
	}

	r := c.db.Create(rp)
	if err := r.Error; err != nil {
		return domain.Reply{}, err
	}

	reply := rp.ToDomain()

	return reply, nil
}
