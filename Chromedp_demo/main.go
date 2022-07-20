package main

import (
	"context"
	_ "github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"log"
	_ "math"
	_ "testing"
	"time"
)

//  chromedp.NewContext() 初始化chromedp的上下文，后续这个页面都使用这个上下文进行操作
//  chromedp.Run() 运行一个chrome的一系列操作
//  chromedp.Navigate() 将浏览器导航到某个页面
//  chromedp.WaitVisible() 等候某个元素可见，再继续执行。
//  chromedp.Click() 模拟鼠标点击某个元素
//  chromedp.Value() 获取某个元素的value值
//  chromedp.ActionFunc() 再当前页面执行某些自定义函数
//  chromedp.Text() 读取某个元素的text值
//  chromedp.Evaluate() 执行某个js，相当于控制台输入js
//  network.SetExtraHTTPHeaders() 截取请求，额外增加header头
//  chromedp.SendKeys() 模拟键盘操作，输入字符
//  chromedp.Nodes() 根据xpath获取某些元素，并存储进入数组
//  chromedp.NewRemoteAllocator
//  chromedp.OuterHTML() 获取元素的outer html
//  chromedp.Screenshot() 根据某个元素截图
//  page.CaptureScreenshot() 截取整个页面的元素
//  chromedp.Submit() 提交某个表单
//  chromedp.WaitNotPresent() 等候某个元素不存在，比如“正在搜索。。。”
//  chromedp.Tasks{} 一系列Action组成的任务
// https://www.jianshu.com/p/d282b4a57596
//
//chromedp.BySearch     //如果不写，默认会使用这个选择器，类似devtools ctrl+f 搜索
//chromedp.ByID         //只id来选择元素
//chromedp.ByQuery      //根据document.querySelector的规则选择元素，返回单个节点
//chromedp.ByQueryAll   // 根据document.querySelectorAll返回所有匹配的节点
//chromedp.ByNodeIP    //检索特定节点(必须先有分配的节点IP)

//chromedp.Navigate("https://xxxx") // 跳转的url
//chromedp.WaitVisible(`#username`, chromedp.ByID),       //使用chromedp.ByID选择器。所以就是等待id=username的标签元素加载完。
//chromedp.SendKeys(`#username`, "username", chromedp.ByID),    //使用chromedp.ByID选择器。向id=username的标签输入username。
//chromedp.SetValue     设置值，和sendkey不同的是，如果原有的输入框中有内容，sendkey不会清空原来的值，setva
//chromedp.Value(`#input1`, val1, chromedp.ByID),    // 获取id=input1的值，并把值传给val1
//chromedp.Click("btn-submit",chromedp.Bysearch),    // 触发点击事件，
//chromedp.Click(`//*[@value="Click me"]`, chromedp.BySearch)
func main() {
	// 禁用chrome headless
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("User-Agent", "Mozilla/5.0 (hp-tablet; Linux; hpwOS/3.0.0; U; en-US) AppleWebKit/534.6 (KHTML, like Gecko) wOSBrowser/233.70 Safari/534.6 TouchPad/1.0"),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 10030*time.Second)
	defer cancel()

	loginUrl := `https://passport.csdn.net/login?code=account`
	username := "yy要上进"
	password := "Xuyiyun110"
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(loginUrl),
		chromedp.Sleep(2 * time.Second),
		chromedp.SendKeys(`.login-form-item .base-input-text `, username, chromedp.NodeVisible),
		chromedp.SendKeys(`.base-input .base-input-text type="password"`, password, chromedp.ByID),
		chromedp.Sleep(1 * time.Second),
		chromedp.Click(`button[data-type="account"]`, chromedp.ByQuery),
	})

	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 100000)
}
