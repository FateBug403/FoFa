package result

type Result struct {
	InFos []InFo
}

type InFo struct {
	Id       int64     `json:"id"` //ID
	Ip       string    `json:"ip"`    //目标ip
	Port     string    `json:"port"` //端口
	Protocol string    `json:"protocol"` //协议名
	Host     string    `json:"host"` //主机名
	Domain   string    `json:"domain"` //域名
	Os       string    `json:"os"` //操作系统
	Server   string    `json:"server"` //网站server
	Icp      string    `json:"icp"` //icp备案号
	Title    string    `json:"title"` //网站标题
}

func (receiver *Result) GetHosts() []string {
	var hosts []string
	for _,value:=range receiver.InFos{
		hosts = append(hosts,value.Host)
		//log.Println(value.Host)
	}
	return hosts

}