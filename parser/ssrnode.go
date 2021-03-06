package parser

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

// SSRNode ssr节点信息
type SSRNode struct {
	// 节点名字
	NodeName string `json:"node_name"`
	// 节点类型
	Type string `json:"-"`

	// 节点IP地址
	IP     string `json:"server"`
	Port   int64  `json:"server_port"`
	Passwd string `json:"password"`

	// 加密算法
	Crypto string `json:"method"`
	// 连接协议
	Proto string `json:"protocol"`
	// 混淆算法
	Minx string `json:"obfs"`
}

// Store 将配置信息存入json文件
func (s *SSRNode) Store(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		return err
	}

	if _, err := f.Write(data); err != nil {
		return err
	}

	return nil
}

// Load 从配置文件读取node信息
func (s *SSRNode) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, s)
	if err != nil {
		return err
	}

	return nil
}

// NameNumber 获取节点的编号
func (s *SSRNode) NameNumber() string {
	nameSplit := strings.Split(s.NodeName, "_")
	return "节点" + nameSplit[len(nameSplit)-1]
}
