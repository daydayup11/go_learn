package main

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"os"
)

func main() {
	// 创建登录阶段的上下文
	ctxLogin, cancelLogin := chromedp.NewContext(context.Background())
	defer cancelLogin()
	// 创建获取文件链接阶段的上下文
	// 定义两个变量分别用于存储两个阶段的截图
	var bufLogin []byte
	var nodes []*cdp.Node
	// 登录阶段
	if err := chromedp.Run(ctxLogin,
		chromedp.Navigate("https://shimo.im/login"),
		chromedp.WaitVisible(`[name="account"]`, chromedp.ByQuery),
		chromedp.SendKeys(`[name="account"]`, "xxx@shimo.im", chromedp.ByQuery),
		chromedp.SendKeys(`[name="password"]`, "xxx", chromedp.ByQuery),
		chromedp.Click(`div.StyledCheckBox-sc-RjILa-1`, chromedp.ByQuery),
		chromedp.Click(`div[type="black"] button[data-test="btn-submit"]`, chromedp.ByQuery),
		chromedp.WaitVisible(`.StyledUserCardContainer-sc-RevvT-9`, chromedp.ByQuery),
		chromedp.CaptureScreenshot(&bufLogin),
		// 获取所有文件链接的节点列表
		chromedp.Nodes(`ul.Wrapper-sc-34VWDK.bdKFMr div.Item-sc-RevvT-7 a.LinkName-sc-RevvT-4.hrBuEg`, &nodes, chromedp.ByQueryAll),
	); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("elementScreenshot.png", bufLogin, 0o644); err != nil {
		log.Fatal(err)
	}
	// 在这里处理两个阶段的截图数据 bufLogin 和 bufGetFiles
	// 打印找到的链接文本
	for _, link := range nodes {
		log.Println(link.AttributeValue("href"))
	}
}
