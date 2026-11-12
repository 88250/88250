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

每天自动刷新，最近更新时间：`2026-11-13 00:38:02`

📝 帖子 &nbsp; 💬 评论 &nbsp; 🗣 回帖 &nbsp; 🌙 清月 &nbsp; 👨‍💻 用户 &nbsp; 🏷️ 标签 &nbsp; ⭐️ 关注 &nbsp; 👍 赞同 &nbsp; 💗 感谢 &nbsp; 💰 打赏 &nbsp; 🗃 收藏

* 💬 [v3.8.10 更新后手机端展示问题求助](https://ld246.com/article/1794477303285/comment/1794482513933#comments)

  > 弹窗移到底部是新版的界面调整，高度会随内容变化。截图中底部露出背景的问题已修复，相关弹窗的按钮样式也已统一，会在后续版本更新。感谢反馈！
* 💗💬 [v3.8.10 更新后手机端展示问题求助](https://ld246.com/article/1794477303285/comment/1794477846324#comments)

  > @88250 看起来抽屉的高度不对（统一显示到屏幕中间），而且底部还镂空（要留空而不是镂空）
* 💗📝 [v3.8.10 更新后手机端展示问题求助](https://ld246.com/article/1794477303285)

  > 手机端 iOS 版本的思源笔记更新完 v3.8.10 后，之前中间显示的都整体下移了，这个是更新的内容么？还是 bug？ [图片]
* 💬 [产品对环境变量的重写逻辑探讨](https://ld246.com/article/1794472539862/comment/1794477176335#comments)

  > 感谢反馈，已确认存在这个问题。读取 PATH 时会展开变量引用，写回还会将 REG_EXPAND_SZ 改成 REG_SZ，导致后续修改 JAVA_HOME 等变量时，对应路径无法随之变化。 已在开发分支修复，覆盖安装和卸载的系统级、用户级 PATH 处理：保留原有引用、注册表类型及其他条目，无变化时不写入。相关隔离回 ..
* 💗📝 [产品对环境变量的重写逻辑探讨](https://ld246.com/article/1794472539862)

  > $k = '$INSTDIR\resources\kernel\' $p = [Environment]::GetEnvironmentVariable('Path', 'Machine') if ((-not $p) -or -not ($p.Split(';') -contains $k)) { $p = '$k; ..
* 💗🌙 [感谢了清风明月](https://ld246.com/member/IMTOP/breezemoons/1794455805681)

  > 软件搞得不错，把 joplin 慢慢就转过来了。期待越做越好。
* 💬 [关于多个数据库之间及其绑定的子文档的搭建求助](https://ld246.com/article/1794374982328/comment/1794469360999#comments)

  > 我的意思是另外建一个「对象清单」数据库，用关联字段指向统一的条文数据库，备注则写在对象清单中。 例如，同一条条文适用于 A 和 B，就在对象清单中建两条记录：一条选 A，一条选 B，两条都关联到同一条条文，但可以分别填写备注。需要查看 A 或 B 的清单时，再按对象筛选视图即可。 这样条文内容只维护一份，各对象的备注可 ..
* 💬 [知乎的视频不能剪藏](https://ld246.com/article/1794455350459/comment/1794460221019#comments)

  > 知乎视频剪藏后只能保留封面图片，不支持剪藏为可播放的视频。可以通过原文链接回到知乎观看。


<!--events end -->
