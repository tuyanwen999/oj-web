package login

import (
	"encoding/json"

	"errors"

	"strings"
)

type GithubErrorMess struct {
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
}

type GithubMess struct {
	UserName string `json:"login"`
	Avatar   string `json:"avatar_url"`
	NickName string `json:"name"`
	OpenId   int    `json:"id"`
}

func GithubCallback(code string) (string, error) {
	baseUrl := "https://github.com/login/oauth/access_token"
	param := map[string]string{
		"client_id":     "080191e49e855122ea33",
		"client_secret": "34b9a36397b171f01e83fc3c5b676177b29df79e",
		"code":          code,
	}
	body, err := (&Url{}).post(baseUrl, param)
	if err != nil {
		return "", err
	}
	if strings.Contains(body, "error_description") {
		return "", errors.New(body)
	}
	start := strings.Index(body, "=")
	end := strings.Index(body, "&")
	return body[start+1 : end], nil
}

func GetGithubMess(accessToken string) (*GithubMess, error) {
	baseUrl := "https://api.github.com/user"
	param := map[string]string{
		"access_token": accessToken,
	}
	body, err := (&Url{}).get(baseUrl, param)
	if err != nil {
		return nil, errors.New("get response fail")
	}
	if strings.Contains(body, "message") {
		mess := &GithubErrorMess{}
		err = json.Unmarshal([]byte(body), mess)
		if err != nil {
			return nil, errors.New("decode response fail")
		}
		return nil, errors.New(mess.Message)
	} else {
		mess := &GithubMess{}
		err = json.Unmarshal([]byte(body), mess)
		if err != nil {
			return nil, errors.New("decode response fail")
		}
		return mess, nil
	}
}
