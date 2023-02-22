package core

import (
	"context"
	json "github.com/json-iterator/go"
	"github.com/wuranxu/mouse-client/internal/core/scene"
	"github.com/wuranxu/mouse-client/internal/entity"
	"github.com/wuranxu/mouse-client/internal/protocol/http"
	"log"
	"regexp"
	"strings"
)

var (
	ParamRegex = `\$\{(.+?)\}`
	compiler   = regexp.MustCompile(ParamRegex)
)

type Runner struct {
	sceneId int64
	scene   *scene.Scene
	client  *http.Client
	stat    *RequestStat
}

// Run scene case
func (r *Runner) Run(ctx context.Context) {
	go func() {
		for {
			select {
			case <-r.stat.success:
				log.Println("success")
			case data := <-r.stat.failure:
				log.Println("error: ", data.Exception)
			}
		}
	}()
	select {
	case <-ctx.Done():
		return
	default:
		r.run()
	}
}

func (r *Runner) run() {
	params := make(map[string][]byte)
	for _, step := range r.scene.Steps {
		// build http
		req := r.buildHTTP(step, params)
		// request
		resp := r.client.Do(req)
		if resp.Error != nil {
			r.stat.failure <- Failed(resp.Error)
			return
		}
		// assert
		if err := scene.Assert(step.Check); err != nil {
			r.stat.failure <- Failed(resp.Error)
			return
		}
		// extract parameters
		if err := r.extractParameters(resp, step.Out, &params); err != nil {
			r.stat.failure <- Failed(err)
			return
		}
	}
	r.stat.success <- Success()
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

func NewRunner(sceneId int64, sceneData []byte) (*Runner, error) {
	var sc scene.Scene
	// load scene data
	if err := scene.Load(sceneData, &sc); err != nil {
		return nil, err
	}
	return &Runner{
		sceneId: sceneId,
		scene:   &sc,
		client:  http.NewHTTPClient(),
		stat: &RequestStat{
			success: make(chan *TestResult, 500),
			failure: make(chan *TestResult, 500),
		},
	}, nil
}
