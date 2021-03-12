// Package ppsdb contains the database schema that PPS uses.
package ppsdb

import (
	"path"

	etcd "github.com/coreos/etcd/clientv3"

	col "github.com/pachyderm/pachyderm/v2/src/internal/collection"
	"github.com/pachyderm/pachyderm/v2/src/pps"
)

const (
	pipelinesPrefix = "/pipelines"
	jobsPrefix      = "/jobs"
)

var (
	// JobsPipelineIndex maps pipeline to jobs started by the pipeline
	JobsPipelineIndex = &col.Index{Field: "Pipeline"}

	// JobsInputIndex maps job inputs (repos + pipeline version) to output
	// commit. This is how we know if we need to start a job.
	JobsInputIndex = &col.Index{Field: "Input"}

	// JobsOutputIndex maps job outputs to the job that create them.
	JobsOutputIndex = &col.Index{Field: "OutputCommit"}
)

// Pipelines returns an EtcdCollection of pipelines
func Pipelines(etcdClient *etcd.Client, etcdPrefix string) col.EtcdCollection {
	return col.NewEtcdCollection(
		etcdClient,
		path.Join(etcdPrefix, pipelinesPrefix),
		nil,
		&pps.EtcdPipelineInfo{},
		nil,
		nil,
	)
}

// Jobs returns an EtcdCollection of jobs
func Jobs(etcdClient *etcd.Client, etcdPrefix string) col.EtcdCollection {
	return col.NewEtcdCollection(
		etcdClient,
		path.Join(etcdPrefix, jobsPrefix),
		[]*col.Index{JobsPipelineIndex, JobsOutputIndex},
		&pps.EtcdJobInfo{},
		nil,
		nil,
	)
}
