package worker

import (
	"github.com/kubeflow/katib/pkg/api"
)

type Interface interface {
	SpawnWorker(wid string, workerConf *api.WorkerConfig) error
	StoreWorkerLog(wID string, namespace string) error
	IsWorkerComplete(wID string, namespace string) (bool, error)
	UpdateWorkerStatus(studyId string, namespace string) error
	StopWorkers(studyId string, wIDs []string, iscomplete bool, namespace string) error
	CleanWorkers(studyId string, namespace string) error
}
