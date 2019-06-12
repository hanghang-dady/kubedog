package pod

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/flant/kubedog/pkg/tracker/indicators"
)

type PodReadyIndicator struct {
	PhaseIndicator *indicators.StringEqualConditionIndicator
	IsReady        bool
}

type PodStatus struct {
	corev1.PodStatus
	ReadyIndicator PodReadyIndicator

	IsFailed     bool
	FailedReason string
}

func NewPodStatus(readyIndicator PodReadyIndicator, isFailed bool, failedReason string, kubeStatus corev1.PodStatus) PodStatus {
	return PodStatus{
		PodStatus:      kubeStatus,
		ReadyIndicator: readyIndicator,
		IsFailed:       isFailed,
		FailedReason:   failedReason,
	}
}

func NewPodReadyIndicator(pod *corev1.Pod, newStatus *corev1.PodStatus) PodReadyIndicator {
	res := PodReadyIndicator{
		IsReady:        false,
		PhaseIndicator: &indicators.StringEqualConditionIndicator{},
	}

	res.PhaseIndicator.Value = string(newStatus.Phase)
	res.PhaseIndicator.TargetValue = string(corev1.PodRunning)

	for _, cond := range newStatus.Conditions {
		if cond.Type == corev1.PodReady && cond.Status == corev1.ConditionTrue {
			res.IsReady = true
			break
		}
	}

	return res
}