package policy

import (
	"scheduler/entity"
	"strings"
)

type Factory struct {

}

func (this *Factory)FindPolicy( jobinfo *entity.JobInfo)(Policy)  {


	switch jobinfo.Type {
	case "PRIORITY":
		return &PriorityPolicy{urls:strings.Split(jobinfo.Urls,","),retryMaxTimes:3}
	default:
		return &RandomPolicy{urls:strings.Split(jobinfo.Urls,","),retryMaxTimes:3}

	}
}


