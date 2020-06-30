/**
 * @Author: liyalei
 * @Description:
 * @Version:
 * @Date: 2020/6/22 3:05 下午
 */
package types

//node
type NodeItem struct {
	OrgName  string `json:"org_name,omitempty"`
	NodeName string `json:"node_name,omitempty"`
	NodeAddr string `json:"node_addr,omitempty"`
}
type OrgNodesMap map[string]NodeItem
type QueryOrgResult struct {
	Records []NodeItem `json:"node_list"`
}

//task
type Task struct {
	Owner   string `json:"t_owner"`
	Name    string `json:"t_name"`
	Version string `json:"t_version"`
	Data    string `json:"t_data"`
}
