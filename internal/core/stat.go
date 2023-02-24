package core

import (
	"context"
	"fmt"
	tool "github.com/wuranxu/mouse-tool"
	"log"
	"time"
)

type stepStat struct {
	requestNum int64
	failedNum  int64
	successNum int64
	cost       int64
}

type sceneStat struct {
	stepStat
	name string
}

type RequestStat struct {
	taskId       int64
	stepStat     map[time.Time]map[string]*stepStat
	sceneStat    map[time.Time]*sceneStat
	success      chan *TestResult
	failure      chan *TestResult
	sceneSuccess chan *TestResult
	sceneFailure chan *TestResult

	sceneNum        int64
	sceneSuccessNum int64
	sceneFailedNum  int64
	cost            int64
	startTime       time.Time

	addr   string
	influx *tool.InfluxdbClient
}

type TestResult struct {
	Name       string
	StatusCode int
	Result     bool
	Response   string
	Elapsed    int64
	Exception  string
	EndTime    time.Time
}

func (s *RequestStat) stat(ctx context.Context) {
	// save startTime
	go s.statStep(ctx)
	go s.statScene(ctx)
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.Upload()
		}
	}
}

// Upload collect statistic info and report to console
func (s *RequestStat) Upload() {
	// qps average
	ms := time.Since(s.startTime).Milliseconds()
	qpsAvg := s.sceneNum / (ms / 1000.0)
	// RT
	rt := s.cost / s.sceneNum
	log.Println("scene qps: ", qpsAvg)
	log.Println("total scene: ", s.sceneNum)
	log.Println("total cost: ", s.cost)
	log.Println("rt(ms): ", rt)

}

func (s *RequestStat) statStep(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-s.success:
			if err := s.write(data); err != nil {
				log.Fatal("write record to influxdb error: ", err)
			}
			//if _, ok := s.stepStat[data.EndTime]; !ok {
			//	s.stepStat[data.EndTime] = make(map[string]*stepStat, 0)
			//}
			//if _, ok := s.stepStat[data.EndTime][data.Name]; !ok {
			//	s.stepStat[data.EndTime][data.Name] = new(stepStat)
			//}
			//s.stepStat[data.EndTime][data.Name].requestNum++
			//s.stepStat[data.EndTime][data.Name].successNum++
			//s.stepStat[data.EndTime][data.Name].cost += data.Elapsed
		case data := <-s.failure:
			if err := s.write(data); err != nil {
				log.Fatal("write record to influxdb error: ", err)
			}
			//if _, ok := s.stepStat[data.EndTime]; !ok {
			//	s.stepStat[data.EndTime] = make(map[string]*stepStat, 0)
			//}
			//if _, ok := s.stepStat[data.EndTime][data.Name]; !ok {
			//	s.stepStat[data.EndTime][data.Name] = new(stepStat)
			//}
			//s.stepStat[data.EndTime][data.Name].requestNum++
			//s.stepStat[data.EndTime][data.Name].failedNum++
			//s.stepStat[data.EndTime][data.Name].cost += data.Elapsed
		}
	}
}

func (s *RequestStat) write(result *TestResult) error {
	return s.influx.Write("report", map[string]string{
		"taskId":   fmt.Sprintf("%v", s.taskId),
		"hostname": s.addr,
	}, map[string]interface{}{
		"elapsed":     result.Elapsed,
		"response":    result.Response,
		"name":        result.Name,
		"status_code": result.StatusCode,
		"exception":   result.Exception,
		"result":      result.Result,
	}, result.EndTime)
}

func (s *RequestStat) statScene(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-s.sceneSuccess:
			if _, ok := s.sceneStat[data.EndTime]; !ok {
				s.sceneStat[data.EndTime] = new(sceneStat)
			}
			s.sceneStat[data.EndTime].requestNum++
			s.sceneStat[data.EndTime].successNum++
			s.sceneStat[data.EndTime].cost += data.Elapsed
			s.sceneNum++
			s.sceneSuccessNum++
			s.cost += data.Elapsed
		case data := <-s.sceneFailure:
			if _, ok := s.sceneStat[data.EndTime]; !ok {
				s.sceneStat[data.EndTime] = new(sceneStat)
			}
			s.sceneStat[data.EndTime].requestNum++
			s.sceneStat[data.EndTime].failedNum++
			s.sceneStat[data.EndTime].cost += data.Elapsed
			s.sceneNum++
			s.sceneFailedNum++
			s.cost += data.Elapsed
		}
	}
}

func Success(name string, elapsed int64, status int, response string, now time.Time) *TestResult {
	return &TestResult{
		Result:     true,
		Name:       name,
		StatusCode: status,
		Response:   response,
		EndTime:    now,
		Elapsed:    elapsed,
	}
}

func Failed(name string, elapsed int64, status int, response string, now time.Time, err error) *TestResult {
	return &TestResult{
		Exception:  err.Error(),
		Result:     false,
		Name:       name,
		StatusCode: status,
		Response:   response,
		EndTime:    now,
		Elapsed:    elapsed,
	}
}
