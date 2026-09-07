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

每天自动刷新，最近更新时间：`2026-09-07 22:02:42`

📝 帖子 &nbsp; 💬 评论 &nbsp; 🗣 回帖 &nbsp; 🌙 清月 &nbsp; 👨‍💻 用户 &nbsp; 🏷️ 标签 &nbsp; ⭐️ 关注 &nbsp; 👍 赞同 &nbsp; 💗 感谢 &nbsp; 💰 打赏 &nbsp; 🗃 收藏

* 💬 [版本更新测试问题](https://ld246.com/article/1788774301022/comment/1788785294324#comments)

  > 我本人赞同并正在实践 AI 开发，也认为 AI 能显著提高开发效率。我们更关注的是，如何在保持开发效率的同时，让测试验证能力也跟上开发速度。 对于近期新开发的部分特性，我们已经开始同步编写 docs，明确功能预期，并逐步补充端到端测试，相关项目地址是：[链接]。 另外，我们长期积累的大量 issues 对 AI 开发也 ..
* 💬 [手机和电脑互相导出导入吗？](https://ld246.com/article/1788773957345/comment/1788780483627#comments)

  > 没有问题的
* 💬 [为什么工作空间会出现 corrupted 这个文件夹？](https://ld246.com/article/1788771695404/comment/1788774398958#comments)

  > corrupted 是内核自动创建的损坏数据隔离目录。第一层目录是发生时间，第二层目录是笔记本 ID，所以与 data 下某个笔记本目录名称一致是正常的。 图中的目录创建于 2026-07-05，可能是加密笔记本早期版本误判密文后留下的，相关问题当天已经修复。现在目录为空，可以直接忽略或删除。 如果以后再次产生新的目录 ..
* 💗📝 [自定义图标和挂件无法显示（Flatpak 安装）](https://ld246.com/article/1788742378008)

  > 似乎被当作敏感文件拦截了 实际路径在~/.var/app/org.b3log.siyuan/SiYuan/data/emojis/ E 2026/09/07 08:02:02 serve.go:647: refuse to serve sensitive static file [/widgets/listChildD ..
* 💬 [自定义图标和挂件无法显示（Flatpak 安装）](https://ld246.com/article/1788742378008/comment/1788744019828#comments)

  > 感谢补充，已定位并修复工作空间路径解析导致的误拦截问题，修复会包含在下一版本中，无需移动工作空间或重装。 跟踪：[链接]
* 💬 [字体为红色的内容无法在安卓手机端 app 正确显示](https://ld246.com/article/1788743448772/comment/1788743626301#comments)

  > 使用默认字体试试
* 💬 [自定义图标和挂件无法显示（Flatpak 安装）](https://ld246.com/article/1788742378008/comment/1788742829806#comments)

  > 日志显示是内核的敏感路径检查拦截了资源，可能与工作空间路径中的符号链接有关，.var 目录本身不会触发拦截。 麻烦提供一下思源版本、Linux 发行版，以及下面两条命令的输出，帮助确认是否存在路径误判： readlink -f ~/.var/app/org.b3log.siyuan/SiYuan readlink -f ..
* 💬 [过几天就弹出：账号鉴权失败](https://ld246.com/article/1788694077046/comment/1788742539353#comments)

  > @JimAmadeus @GinkgoLight 等下个版本再试试


<!--events end -->
