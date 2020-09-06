package mastodon

import (
	"encoding/json"
	"errors"

	"github.com/fatih/color"
	"github.com/go-playground/validator"
)

// Mastodon 总入口结构体
type Mastodon struct {
	Token  string
	Debug  bool
	Domain string
}

// GetHomeTimeLines 获取首页列表数据
func (mastodon *Mastodon) GetHomeTimeLines() ([]HomeResp, error) {
	body, err := Get(mastodon.Domain+HomeTimeLines, mastodon.Token)
	if err != nil {
		color.Red(err.Error())
		return nil, err
	}

	var result []HomeResp

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		color.Red("解析字符串出错！", err)
		return nil, err
	}

	return result, nil
}

// SendStatuses 发嘟
func (mastodon *Mastodon) SendStatuses(params *StatusParams) (*StatusRes, error) {
	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return nil, errors.New(err.Value().(string))
		}
	}

	body, err := Post(mastodon.Domain+sendStatusesURL, mastodon.Token, params)
	if err != nil {
		return nil, err
	}

	var result StatusRes
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (mastodon *Mastodon) errHandler(msg string, err error) {
	if mastodon.Debug {
		color.Red("%s - %s", msg, err.Error())
	}
}
