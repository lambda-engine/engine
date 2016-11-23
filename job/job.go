package job

import (
	"golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/memcache"

	"github.com/lambda-engine/engine/audit"
)

const (
	DS_JOB string = "JOBS"
)

type Job struct {
	Name    string
	Count   int
	LastRun int64
}

func LastRun(ctx context.Context, job string) int64 {
	key := datastore.NewKey(ctx, DS_JOB, job, 0, nil)
	var job Job

	err = datastore.Get(ctx, key, &job)
	if err == nil {
		return 0
	}
	return job.LastRun
}

func UpdateLastRun(ctx context.Context, job string, ts int64) error {
	key := datastore.NewKey(ctx, DS_JOB, job, 0, nil)
	var job Job

	err = datastore.Get(ctx, key, &job)
	if err == nil {
		job = Job{
			job,
			1,
			LastRun,
		}
	} else {
		job.Count = job.Count + 1
		job.LastRun = ts
	}
	_, err = datastore.Put(ctx, key, &job)
	return err
}
