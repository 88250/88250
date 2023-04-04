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

每天自动刷新，最近更新时间：`2023-04-04 16:05:06`

📝 帖子 &nbsp; 💬 评论 &nbsp; 🗣 回帖 &nbsp; 🌙 清月 &nbsp; 👨‍💻 用户 &nbsp; 🏷️ 标签 &nbsp; ⭐️ 关注 &nbsp; 👍 赞同 &nbsp; 💗 感谢 &nbsp; 💰 打赏 &nbsp; 🗃 收藏

* 💬 [docker 版本无法导出 pdf 吗？](https://ld246.com/article/1680594404377/comment/1680594637384#comments)

  > 是的，Docker 部署不支持导出 PDF 等格式。
* 💬 [求助：2.0.18 离线升级到 2.8.2 后，之前的文档全都没有了](https://ld246.com/article/1680593250881/comment/1680593654587#comments)

  > 安装新版本会卸载旧版本，如果工作空间放在安装路径下的话会在卸载时被一并删除，只能通过其他备份数据恢复或者试下磁盘恢复工具了。
* 💗💬 [删除含有闪卡的文档后，闪卡管理包里会存在很多不存在符合条件的内容块](https://ld246.com/article/1680575149128/comment/1680592975158#comments)

  > 谢谢 D 大的回复。 前面的可以在打开闪卡管理界面的最上面显示就好了（这样就可以很方便的手动删除了，如果不可以也没事，我改变一下习惯，先一键删除闪卡再删除文档） 后面的建议是想有闪卡时会显示，没有闪卡时不会显示（就像现在没有子文档就不会显示包含了几个子文档一样，没有也没事，主要是挺想要在总复习和管理界面可以选择 all ..
* 💬 [频繁遇到思源笔记卡死的情况](https://ld246.com/article/1680438992079/comment/1680591034785#comments)

  > 日志中没有发现异常，估计只能再观察看看了……
* 💬 [删除含有闪卡的文档后，闪卡管理包里会存在很多不存在符合条件的内容块](https://ld246.com/article/1680575149128/comment/1680590958615#comments)

  > 文档删除后可以通过快照或者文件历史回滚，但是闪卡不行，所以这里不做关联删除。 后面的提议如果不用闪卡的话都是 0，显示会比较冗余，所以也暂时不考虑增加了。
* 💬 [siyuan-2.8.2-mac-arm64 重装后选择工作目录后没有反应](https://ld246.com/article/1680518813121/comment/1680590866356#comments)

  > 云端损坏的话只能重建新的云端目录再同步，可能是之前版本的缺陷所致。
* 💬 [siyuan-2.8.2-mac-arm64 重装后选择工作目录后没有反应](https://ld246.com/article/1680518813121/comment/1680581776274#comments)

  > 帮忙试试这个地方的代码： 在安装目录找下 init.html，大概路径应该是安装包下/Contents/Resources/app/electron/init.html 打开后看下能否编辑保存，将其中 isICloudPath 这个函数实现直接返回 false [图片]
* 💬 [siyuan-2.8.2-mac-arm64 重装后选择工作目录后没有反应](https://ld246.com/article/1680518813121/comment/1680579510715#comments)

  > 初始化界面不应该有这么多的磁盘读流量才对，不过我也不太确定是不是因为那个 iCloud 路径校验引起的。 配置文件大概是这几个地方：~/.config/siyuan、userData/SiYuan-Electron，后面这个是 Electron 生成的，macOS 上 userData 的具体路径你得自己找一下看看。


<!--events end -->
