package job

import "mint/http/client"

func Exec(node *Node) {
	builder := client.NewBuilder()
	builder.SetUrl(node.Url).SetContent(node.Content)

	sender := client.NewSender(builder)
	sender.Send(node.Method)
}
