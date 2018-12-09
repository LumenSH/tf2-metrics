package tasks

import (
	"git.lumen.sh/xNevo/tf2-metrics/vectors"
	"github.com/prometheus/common/log"
	"sync"
)

var Tasks []func()

func init() {
	Tasks = []func(){
		SteamStatus,
		SteamTF2,
	}
}

func Run() {
	wg := &sync.WaitGroup{}
	wg.Add(len(Tasks))

	vectors.Reset()
	for _, f := range Tasks {
		go func(f func()) {
			f()
			defer func() {
				wg.Done()
				if err := recover(); err != nil {
					log.Errorf("Recovered from error: %s", string(err.(string)))
				}
			}()
		}(f)
	}

	wg.Wait()
}
