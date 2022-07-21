package management

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/trustasia-com/go-sdk/wekey/management/types"
)

func TestClient_GetIdpGithubConfig(t *testing.T) {
	initTestClient()

	data, err := testClient.GetIdpGithubConfig()
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_AddIdpGithub(t *testing.T) {

	initTestClient()
	err := testClient.AddIdpGithub(types.AddIdpGithubRequest{
		Name:         "test-github",
		ClientID:     "sadascxvxcvxcsedfdsf",
		ClientSecret: "asdasdasdasdqwqerrgfhytew22er",
	})

	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_UpdateIdpGithub(t *testing.T) {
	initTestClient()
	err := testClient.UpdateIdpGithub(types.UpdateIdpGithub{
		Name:         "test-github",
		ClientID:     "sadascxvxcvxcsedfdsf",
		ClientSecret: "asdasdasdasdqwqerrgfhytew22er",
	})

	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_DelIdpGithub(t *testing.T) {
	initTestClient()
	err := testClient.DelIdpGithub()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("del success")
}

func TestClient_GetIdpGitlabConfig(t *testing.T) {
	initTestClient()

	data, err := testClient.GetIdpGitlabConfig()
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_AddIdpGitlab(t *testing.T) {
	initTestClient()
	err := testClient.AddIdpGitlab(types.AddIdpGitlabRequest{
		Name:         "test-gitlab",
		ClientID:     "sadascxvxcvxcsedfdsf",
		ClientSecret: "asdasdasdasdqwqerrgfhytew22er",
	})

	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_UpdateIdpGitlab(t *testing.T) {
	initTestClient()
	err := testClient.UpdateIdpGitlab(types.UpdateIdpGitlab{
		Name:         "test-gitlab",
		ClientID:     "sadascxvxcvxcsedfdsf",
		ClientSecret: "asdasdasdasdqwqerrgfhytew22er",
	})

	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_DelIdpGitlab(t *testing.T) {
	initTestClient()
	err := testClient.DelIdpGitlab()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("del success")
}

func TestClient_GetIdpWeiboConfig(t *testing.T) {
	initTestClient()

	data, err := testClient.GetIdpWeiboConfig()
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_AddIdpWeibo(t *testing.T) {
	initTestClient()
	err := testClient.AddIdpWeibo(types.AddIdpWeiboRequest{
		Name:         "test-weibo",
		ClientID:     "sadascxvxcvxcsedfdsf",
		ClientSecret: "asdasdasdasdqwqerrgfhytew22er",
	})

	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_UpdateIdpWeibo(t *testing.T) {
	initTestClient()
	err := testClient.UpdateIdpGitlab(types.UpdateIdpGitlab{
		Name:         "test-weibo1",
		ClientID:     "sadascxvxcvxcsedzxfdsf",
		ClientSecret: "asdasdasdasdqwqezrrgfhytew22er",
	})

	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_DelIdpWeibo(t *testing.T) {
	initTestClient()
	err := testClient.DelIdpWeibo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("del success")
}

func TestClient_GetIdpGitEEConfig(t *testing.T) {
	initTestClient()

	data, err := testClient.GetIdpGitEEConfig()
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_AddIdpGitEE(t *testing.T) {
	initTestClient()
	err := testClient.AddIdpGitEE(types.AddIdpGitEERequest{
		Name:         "test-gitee",
		ClientID:     "sadascxvxcvxcsedfdsf",
		ClientSecret: "asdasdasdasdqwqerrgfhytew22er",
	})

	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_UpdateIdpGiteEE(t *testing.T) {
	initTestClient()
	err := testClient.UpdateIdpGitEE(types.UpdateIdpGitEE{
		Name:         "test-gitee23",
		ClientID:     "sadascxvxcvxcsedzxfdsf",
		ClientSecret: "asdasdasdasdqwqezrrgfhytew22er",
	})

	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_DelIdpGitEE(t *testing.T) {
	initTestClient()
	err := testClient.DelIdpGitEE()
	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_GetIdpBaiduConfig(t *testing.T) {
	initTestClient()

	data, err := testClient.GetIdpBaiduConfig()
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_AddIdpBaidu(t *testing.T) {
	initTestClient()
	err := testClient.AddIdpBaidu(types.AddIdpBaiduRequest{
		Name:         "test-baidu",
		ClientID:     "sadascxvxcvxcsedfdsf",
		ClientSecret: "asdasdasdasdqwqerrgfhytew22er",
	})

	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_UpdateIdpBaidu(t *testing.T) {
	initTestClient()
	err := testClient.UpdateIdpBaidu(types.UpdateIdpBaidu{
		Name:         "test-baidu123",
		ClientID:     "sadascxvxcvxcsedzxfdsf",
		ClientSecret: "asdasdasdasdqwqezrrgfhytew22er",
	})

	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_DelIdpBaidu(t *testing.T) {
	initTestClient()
	err := testClient.DelIdpBaidu()
	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_GetIdpQQConfig(t *testing.T) {
	initTestClient()

	data, err := testClient.GetIdpQQConfig()
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_AddIdpQQ(t *testing.T) {
	initTestClient()
	err := testClient.AddIdpQQ(types.AddIdpQQRequest{
		Name:         "test-qq",
		ClientID:     "sadascxvxcvxcsedfdsf",
		ClientSecret: "asdasdasdasdqwqerrgfhytew22er",
	})

	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_UpdateIdpQQ(t *testing.T) {
	initTestClient()
	err := testClient.UpdateIdpQQ(types.UpdateIdpQQ{
		Name:         "test-qq",
		ClientID:     "sadascxvxcvxcsedzxfdsf",
		ClientSecret: "asdasdasdasdqwqezrrgfhytew22er",
	})

	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_DelIdpQQ(t *testing.T) {
	initTestClient()
	err := testClient.DelIdpQQ()
	if err != nil {
		log.Fatal(err)
	}
}
