package kurs

import "context"

// IKursService interface
type IKursService interface {
	IndexingKurs(ctx context.Context) (ResultIndexing []DataKurs, err error)
}

// IKursRepository interface
type IKursRepository interface {
}
