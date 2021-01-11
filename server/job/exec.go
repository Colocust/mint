package job

import (
	"fmt"
)

func Exec(node *Node) {
	fmt.Println("node", node)
	//builder := client.NewBuilder()
	//builder.SetUrl(node.Url).SetContent(node.Content)
	//
	//sender := client.NewSender(builder)
	//result, ret := sender.Send(node.Method)
	//fmt.Println(result, ret)
}
