package mem

import (
	"bytes"
	"context"
	"encoding/gob"
	"sync"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/vgough/sequin/job"
)

type store struct {
	mu   sync.Mutex
	ids  map[string]*record
	tags map[string]map[string]*record // tag -> id -> rec
}

type record struct {
	mu sync.Mutex

	id   string
	tags []string

	snapshots []*snapshot
}

var _ job.Record = &record{}

type snapshot struct {
	r     *record
	index int

	desc  string
	state []byte
	err   error
	final bool
	when  time.Time
}

var _ job.Snapshot = &snapshot{}

// NewStore returns a new store instance.
func NewStore() job.Store {
	return &store{
		ids:  make(map[string]*record),
		tags: make(map[string]map[string]*record),
	}
}

func (s *store) Close() error {
	return nil
}

func (s *store) Create(ctx context.Context, job job.Runnable, tags ...string) (job.Record, error) {
	data, err := encodeRunnable(job)
	if err != nil {
		return nil, err
	}

	snap := &snapshot{
		state: data,
		when:  time.Now(),
	}

	r := &record{
		id:   uuid.NewV4().String(),
		tags: tags,
		snapshots: []*snapshot{
			snap,
		},
	}
	snap.r = r

	s.mu.Lock()
	s.ids[r.id] = r
	for _, tag := range tags {
		m, ok := s.tags[tag]
		if !ok {
			m = make(map[string]*record)
			s.tags[tag] = m
		}
		m[r.id] = r
	}
	s.mu.Unlock()

	return r, nil
}

func (s *store) Load(ctx context.Context, id string) (job.Record, error) {
	s.mu.Lock()
	r, ok := s.ids[id]
	s.mu.Unlock()

	if !ok {
		return nil, errors.Errorf("no such job: %q", id)
	}

	return r, nil
}

func (s *store) SearchByTag(ctx context.Context, tag string) (ids []string, err error) {
	s.mu.Lock()
	tagged := s.tags[tag]
	for id := range tagged {
		ids = append(ids, id)
	}
	s.mu.Unlock()
	return ids, nil
}

func (s *store) Delete(ctx context.Context, id string) error {
	s.mu.Lock()
	r, ok := s.ids[id]
	if !ok {
		s.mu.Unlock()
		return errors.Errorf("no such job: %q", id)
	}
	delete(s.ids, id)

	for _, tag := range r.tags {
		tags := s.tags[tag]
		delete(tags, id)
		if len(tags) == 0 {
			delete(s.tags, tag)
		}
	}
	s.mu.Unlock()
	return nil
}

func (r *record) JobID() string {
	return r.id
}

func (r *record) Tags() []string {
	return r.tags
}

func (r *record) AddSnapshot(ctx context.Context, run job.Runnable, opts ...job.SnapshotOpt) error {
	data, err := encodeRunnable(run)
	if err != nil {
		return err
	}

	r.mu.Lock()
	snap := &snapshot{
		r:     r,
		state: data,
		index: len(r.snapshots),
		when:  time.Now(),
	}
	for _, opt := range opts {
		opt(snap)
	}
	r.snapshots = append(r.snapshots, snap)

	r.mu.Unlock()
	return nil
}

func (r *record) LastSnapshot(ctx context.Context) (job.Snapshot, error) {
	r.mu.Lock()
	last := r.snapshots[len(r.snapshots)-1]
	r.mu.Unlock()
	return last, nil
}

func (r *record) SnapshotsSince(ctx context.Context, after job.Snapshot) ([]job.Snapshot, error) {
	var snap *snapshot
	if after != nil {
		snap = after.(*snapshot)
	}
	r.mu.Lock()
	var start = 0
	if snap != nil {
		start = snap.index + 1
		if snap.r != r {
			r.mu.Unlock()
			return nil, errors.New("invalid snapshot")
		}
	}
	var num = len(r.snapshots) - start
	if num == 0 {
		r.mu.Unlock()
		return nil, nil
	}

	snaps := make([]job.Snapshot, num)
	for i := start; i < len(r.snapshots); i++ {
		snaps[i-start] = r.snapshots[i]
	}
	r.mu.Unlock()

	return snaps, nil
}

func (r *record) WaitForUpdate(ctx context.Context, snap job.Snapshot) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-time.After(250 * time.Millisecond):
			last, err := r.LastSnapshot(ctx)
			if err != nil {
				return err
			}
			if last.When().After(snap.When()) {
				return nil
			}
		}
	}
}

func (s *snapshot) JobID() string {
	return s.r.id
}

func (s *snapshot) Description() string {
	return s.desc
}

func (s *snapshot) Runnable() (job.Runnable, error) {
	return decodeRunnable(s.state)
}

func (s *snapshot) Tags() []string {
	return s.r.tags
}

func (s *snapshot) IsFinal() bool {
	return s.final
}

func (s *snapshot) Result() error {
	return s.err
}

func (s *snapshot) When() time.Time {
	return s.when
}

func (s *snapshot) SetError(err error) {
	s.err = err
}

func (s *snapshot) SetDescription(desc string) {
	s.desc = desc
}

func (s *snapshot) SetFinal(final bool) {
	s.final = final
}

func encodeRunnable(r job.Runnable) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(&r); err != nil {
		return nil, errors.Wrap(err, "unable to encode job")
	}

	return buf.Bytes(), nil
}

func decodeRunnable(data []byte) (job.Runnable, error) {
	var runnable job.Runnable
	dec := gob.NewDecoder(bytes.NewBuffer(data))
	if err := dec.Decode(&runnable); err != nil {
		return nil, err
	}
	return runnable, nil
}
