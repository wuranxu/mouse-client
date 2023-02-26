package core

import (
	"context"
	json "github.com/json-iterator/go"
	"github.com/wuranxu/mouse-client/internal/core/scene"
	"github.com/wuranxu/mouse-client/internal/entity"
	"github.com/wuranxu/mouse-client/internal/protocol/http"
	tool "github.com/wuranxu/mouse-tool"
	"regexp"
	"strings"
	"time"
)

var (
	ParamRegex = `\$\{(.+?)\}`
	compiler   = regexp.MustCompile(ParamRegex)
)

type Runner struct {
	taskId int64
	scene  *scene.Scene
	client *http.Client
	stat   *RequestStat
	addr   string
}

func NewRunner(taskId int64, sceneData []byte) (*Runner, error) {
	var sc scene.Scene
	// load scene data
	if err := scene.Load(sceneData, &sc); err != nil {
		return nil, err
	}
	return &Runner{
		taskId: taskId,
		scene:  &sc,
		client: http.NewHTTPClient(),
		stat: &RequestStat{
			taskId:       taskId,
			result:       make(chan *TestResult, 2000),
			success:      make(chan *TestResult, 2000),
			failure:      make(chan *TestResult, 2000),
			sceneSuccess: make(chan *TestResult, 2000),
			sceneFailure: make(chan *TestResult, 2000),
			stepStat:     make(map[time.Time]map[string]*stepStat),
			sceneStat:    make(map[time.Time]*sceneStat),
		},
	}, nil
}

func (r *Runner) SetInflux(client *tool.InfluxdbClient) {
	r.stat.influx = client
}

func (r *Runner) SetAddr(addr string) {
	r.addr = addr
	r.stat.addr = addr
}

// Run scene case
func (r *Runner) Run(ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	default:
		r.run()
	}
}

func (r *Runner) storeStepResult(result *TestResult, ok bool) {
	if ok {
		r.stat.success <- result
		return
	}
	r.stat.failure <- result
	r.stat.sceneFailure <- result
}

func (r *Runner) StatInfo() *RequestStat {
	return r.stat
}

func (r *Runner) Stat(ctx context.Context) {
	r.stat.startTime = time.Now()
	r.stat.stat(ctx)
}

func (r *Runner) run() {
	params := make(map[string][]byte)
	var (
		timestamp time.Time
		elapsed   int64
	)
	for _, step := range r.scene.Steps {
		// build http
		req := r.buildHTTP(step, params)

		// request
		resp := r.client.Do(req)
		timestamp = time.Now()
		elapsed += resp.Elapsed
		if resp.Error != nil {
			r.stat.result <- Failed(step.Name, elapsed, resp.StatusCode, resp.Data, timestamp, resp.Error)
			break
		}

		// extract parameters
		if err := r.extractParameters(resp, step.Out, &params); err != nil {
			r.stat.result <- Failed(step.Name, elapsed, resp.StatusCode, resp.Data, timestamp, err)
			break
		}
		// assert
		if err := scene.Assert(step.Check, params); err != nil {
			r.stat.result <- Failed(step.Name, elapsed, resp.StatusCode, resp.Data, timestamp, err)
			break
		}

		r.stat.result <- Success(step.Name, elapsed, resp.StatusCode, resp.Data, timestamp)
		//r.storeStepResult(Success(step.Name, elapsed, resp.StatusCode, resp.Data, timestamp), true)
	}

	// store scene success
	//r.stat.sceneSuccess <- Success(r.scene.Name, elapsed, 200, "", timestamp)
}

// buildHTTP build http request
// include request headers and body
func (r *Runner) buildHTTP(step *scene.Step, params map[string][]byte) *entity.HTTPRequest {
	req := http.NewRequest(step.Url, entity.HTTPMethod(step.Method),
		entity.WithQuery(step.Query), entity.WithBody(step.Body),
		entity.WithHeaders(step.Headers), entity.WithTimeout(step.Timeout),
	)
	// replace params
	if len(params) > 0 {
		req.Url = r.replaceParams(req.Url, params).(string)
		req.Headers = r.replaceParams(req.Headers, params).(map[string]string)
		req.Body = r.replaceParams(req.Body, params).(string)
	}
	return req
}

// replaceParams
func (r *Runner) replaceParams(value any, params map[string][]byte) any {
	restore := false
	var result string
	switch val := value.(type) {
	case map[string]string:
		restore = true
		bt, err := json.Marshal(value)
		if err != nil {
			return value
		}
		result = scene.ToString(bt)
	case string:
		restore = false
		result = val
	}
	ans := compiler.FindAllString(result, -1)
	for _, key := range ans {
		s, ok := params[key]
		if !ok {
			continue
		}
		var str string
		if err := json.Unmarshal(s, &str); err != nil {
			continue
		}
		result = strings.ReplaceAll(result, key, str)
	}
	if restore {
		var hd map[string]string
		if err := json.Unmarshal(scene.ToBytes(result), &hd); err != nil {
			return value
		}
		return hd
	}
	return result
}

// extractParameters extract params for usage
func (r *Runner) extractParameters(resp *entity.HTTPResponse, out []*scene.Out, par *map[string][]byte) error {
	for _, ot := range out {
		extract, err := scene.NewExtractor(ot).Extract(resp)
		if err != nil {
			return err
		}
		(*par)["${"+ot.Variable+"}"] = extract
	}
	return nil
}
