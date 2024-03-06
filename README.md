### 我和妻子的开源生活

我和我的妻子 [Vanessa](https://github.com/Vanessa219) 从 2009 年开始编写开源软件，其中我们一直在维护的 Solo 项目到现在已经有 10 年了。我们的创作领域主要围绕博客和社区系统开展，我们一直在努力实现 [B3log 构思 - 分布式社区网络](https://ld246.com/article/1546941897596)。

2018 年初，我和 Vanessa 从公司离职，正式开始了“全职做开源、自由职业者”的创业生涯。我们建立了一家公司，主要产品是 Sym 社区系统，它的[社区版](https://github.com/88250/symphony)是完全开源的，个人可基于其开源协议免费使用。另外，我们也运营着一个超过 5 万用户的社区[链滴](https://ld246.com)，作为 B3log 分布式社区网络的社区端节点。

最近我们开启了新项目[思源笔记](https://github.com/siyuan-note/siyuan)，这是一款本地离线优先的个人知识管理系统，支持细粒度块级引用和 Markdown 所见即所得，欢迎大家来试用和反馈。

对我和 Vanessa 来说，开源已经不仅仅只是一种爱好，更是一种生活方式，我们对这条“开源生活”之路的未来充满了信心。希望在这条路上我们能通过开源软件帮助到其他人，同时其他人也能帮助到我们。

开源连接你我，开源构建未来，让我们一起走进开源的世界！

### 我在 GitHub 上的统计

<a title="Hits" target="_blank" href="https://github.com/88250/88250"><img src="https://hits.b3log.org/88250/88250.svg"></a>

[![Github Stats](https://github-readme-stats.vercel.app/api?username=88250&theme=tokyonight&show_icons=true)](https://github.com/88250)

<!--events start -->

### 我在链滴的近期动态

每天自动刷新，最近更新时间：`2024-03-06 16:05:57`

📝 帖子 &nbsp; 💬 评论 &nbsp; 🗣 回帖 &nbsp; 🌙 清月 &nbsp; 👨‍💻 用户 &nbsp; 🏷️ 标签 &nbsp; ⭐️ 关注 &nbsp; 👍 赞同 &nbsp; 💗 感谢 &nbsp; 💰 打赏 &nbsp; 🗃 收藏

* 💬 [右键复制粘贴菜单消失不见；拼音输入时首字母异常](https://ld246.com/article/1709703804844/comment/1709710626572#comments)

  > 这个是触发了块级菜单，我这里右键不会触发，建议排查一下是不是主题或者插件影响 更新一下搜狗输入法看看，我这里也是搜狗，没有这个问题
* 💬 [建议有图标时钉住时也能显示笔记名称](https://ld246.com/article/1709710315016/comment/1709710390533#comments)

  > 这个大概是主题特性吧，默认主题是全显示的。
* 💬 [WebDAV 数据同步失败 HTTP 413](https://ld246.com/article/1709698004765/comment/1709709128814#comments)

  > 从内核日志上看是因为 WebDAV 服务端返回了 HTTP 状态码 413 导致的上传失败： E 2024/03/06 11:23:19 webdav.go:89: upload object [main/siyuan/repo/objects/9a/6c9beeddfcb6d11236c2d2d1d8cfd4f2d5 ..
* 💬 [WebDAV 数据同步失败 HTTP 413](https://ld246.com/article/1709698004765/comment/1709707942023#comments)

  > 还是传一下内核日志吧，同步数据是并发操作，其中某次上传失败后会退出同步，退出时肯定要删锁的，所以不是你认为的靠这几行日志就能判断出问题，真正的问题可能是其他原因导致。
* 💬 [WebDAV 数据同步失败 HTTP 413](https://ld246.com/article/1709698004765/comment/1709707535357#comments)

  > 在工作空间文件夹里。
* 💬 [WebDAV 数据同步失败 HTTP 413](https://ld246.com/article/1709698004765/comment/1709702721342#comments)

  > 内核日志位于工作空间/temp/siyuan.log
* 💬 [求助 Docker 搭建，输入正确的授权码后 302，继续输入授权码](https://ld246.com/article/1709697747489/comment/1709702675589#comments)

  > 配置过反代吗？配置过的话检测一下域名反代，另外不要使用 URL 重写。
* 💬 [WebDAV 数据同步失败 HTTP 413](https://ld246.com/article/1709698004765/comment/1709698464399#comments)

  > 获取不到锁文件的话会创建锁，删除不存在的锁返回 204 不会导致报错，这些都是正常逻辑。 多半是 WebDAV 服务器或者网络导致的其他问题，建议结合思源内核日志定位问题。


<!--events end -->
