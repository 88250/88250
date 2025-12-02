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

每天自动刷新，最近更新时间：`2025-12-02 16:08:29`

📝 帖子 &nbsp; 💬 评论 &nbsp; 🗣 回帖 &nbsp; 🌙 清月 &nbsp; 👨‍💻 用户 &nbsp; 🏷️ 标签 &nbsp; ⭐️ 关注 &nbsp; 👍 赞同 &nbsp; 💗 感谢 &nbsp; 💰 打赏 &nbsp; 🗃 收藏

* 💬 [鸿蒙 6 平板上偶尔闪退](https://ld246.com/article/1764554163301/comment/1764647458434#comments)

  > 是的，移动端只支持同时运行一个内核，连不上可能是因为端口也受到了系统的隔离限制。
* 💬 [鸿蒙 6 平板上偶尔闪退](https://ld246.com/article/1764554163301/comment/1764645829927#comments)

  > 思源笔记内核是通过后台长时任务（数据传输）运行的，所以会显示这个提示。 切换隐私空间可能会使内核端口无法访问所以会出现这个问题，目前看可能没有解决方案，只能凑合用了。
* 💬 [数据丢失，求如何排查恢复?](https://ld246.com/article/1764553165374/comment/1764583206414#comments)

  > 得上传一下日志文件以便分析
* 💬 [数据丢失，求如何排查恢复?](https://ld246.com/article/1764553165374/comment/1764582113054#comments)

  > 云端损坏的话报错是“云端数据已经损坏”
* 💬 [数据丢失，求如何排查恢复?](https://ld246.com/article/1764553165374/comment/1764581449664#comments)

  > 看报错是本地的数据仓库损坏了，可以试试本地重置数据仓库后再同步。
* 💬 [紧急求助 PC 和手机端升级 v3.4.1 后七牛云同步失败](https://ld246.com/article/1764293269018/comment/1764548811846#comments)

  > 请发一下重置后的日志
* 💬 [请问输入【【引用文档时，弹出的搜索框上显示的推荐文档怎么排序？](https://ld246.com/article/1764415490843/comment/1764473073722#comments)

  > 排查后发现之前的实现一直都有问题，根据 refs.id 排序并不能保证固定，因为 refs 会重新生成，导致 id 被刷新（不再是引用时的时间）。 目前看没有啥解决方案，因为更新数据索引时生成 refs 是 delete 后再 insert，无法实现 update。 综上，只能做到将最近更新文档中的 refs 排序在前 ..
* 💗📝 [超级块复制粘贴文本 *，聚焦引用只显示原文第一行](https://ld246.com/article/1764395026961)

  > [图片]


<!--events end -->
