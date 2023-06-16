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

每天自动刷新，最近更新时间：`2023-06-16 16:04:51`

📝 帖子 &nbsp; 💬 评论 &nbsp; 🗣 回帖 &nbsp; 🌙 清月 &nbsp; 👨‍💻 用户 &nbsp; 🏷️ 标签 &nbsp; ⭐️ 关注 &nbsp; 👍 赞同 &nbsp; 💗 感谢 &nbsp; 💰 打赏 &nbsp; 🗃 收藏

* 💗💬 [数据库表的 created 错误 | 统计今天创建文件](https://ld246.com/article/1686673103683/comment/1686887189690#comments)

  > 我的那个代码是为 query 定制的，这样写 query 能有更好的显示，如果你是嵌入块使用，需要把前面的一串替换为 select * from blocks where。 另外 created 的数据没有问题，一样是因为你的模板代码中的.created 并不能调用底层的 created 数据，模板没有支持这个变量，具 ..
* 💗💬 [数据库表的 created 错误 | 统计今天创建文件](https://ld246.com/article/1686673103683/comment/1686889893479#comments)

  > 你的模板本身就是写错的，思源并不支持以你这样的方式获取创建时间，思源的模板变量并不包括.created 和 .updated。 正确的做法是使用 sql 查询语句获取创建时间。如何获取 @Reader 已有说明。如果你打算在模板中使用，请仔细阅读模板片段内的 queryBlocks 和 querySpans 如何使用。 ..
* 💬 [网页剪藏后文字排版能不能自定义](https://ld246.com/article/1686875286156/comment/1686879673897#comments)

  > 看下默认主题的 CSS。
* 💬 [2.9.1 版本出现内核因未知原因退出](https://ld246.com/article/1686725891806/comment/1686879627949#comments)

  > 试试删掉 C:\Users\用户名\.config\siyuan\workspace.json 后再启动。
* 💬 [请问思源笔记启动后，为什么会有一个长时间的“获取最新文件 ******”的阶段？](https://ld246.com/article/1686873911210/comment/1686877905740#comments)

  > 这个阶段是从本地数据仓库中获取文件列表，然后通过对比本地文件列表和云端文件列表的差异来决定需要同步的数据。 下个版本会对获取文件列表进行一些优化，另外，把工作空间放在固态硬盘上性能会好很多。
* 💗📝 [【思源插件】微信读书插件](https://ld246.com/article/1686672350774)

  > 参考 Obsidian 的微信读书插件，做了一个比较简陋的思源插件。 基本功能和 Ob 的插件一致： 点击插件图标导入所有微信读书的笔记和标注 支持微信扫码登录 支持自动刷新 Cookie（延长登录有效期） 自定义模板 另加了导入部分笔记的特性。 这是导入的样例： [图片] 这是设置界面： [图片] 导入部分笔记和标注 ..
* 💗💬 [2.9.1 history database index failed 错误提示如何处理](https://ld246.com/article/1686843046398/comment/1686844675258#comments)

  > 数据历史-文件历史-右边重建索引
* 💬 [什么原因 / 条件 / 情况会引起资源文件丢失？](https://ld246.com/article/1686785704804/comment/1686829968357#comments)

  > 复制以后在全局搜索中搜索一下，前提是要打开 设置 - 搜索 - 索引 - 资源文件路径


<!--events end -->
