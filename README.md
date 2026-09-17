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

每天自动刷新，最近更新时间：`2026-09-17 21:07:13`

📝 帖子 &nbsp; 💬 评论 &nbsp; 🗣 回帖 &nbsp; 🌙 清月 &nbsp; 👨‍💻 用户 &nbsp; 🏷️ 标签 &nbsp; ⭐️ 关注 &nbsp; 👍 赞同 &nbsp; 💗 感谢 &nbsp; 💰 打赏 &nbsp; 🗃 收藏

* 💬 [安卓端很多问题都没解决就出了正式版](https://ld246.com/article/1789646203860/comment/1789647436018#comments)

  > 我这里打开 PDF 没问题，不需要安装插件；全选下个版本支持 [链接]
* 💬 [思源 3.8.3 版本麒麟系统无法输入中文](https://ld246.com/article/1789612119565/comment/1789639748430#comments)

  > 已找到思源启动命令 /usr/bin/siyuan。桌面快捷方式目前指向自定义脚本 /home/huawei/siyuan-fcitx.sh。 请先彻底退出思源，包括托盘中的思源，再在终端执行： /usr/bin/siyuan --ozone-platform=x11 启动后试一下搜索框和正文能否输入中文。如果仍然不行 ..
* 💬 [思源 3.8.3 版本麒麟系统无法输入中文](https://ld246.com/article/1789612119565/comment/1789637429851#comments)

  > 从输出看，你当前使用的是 Wayland 和 Fcitx，终端中的输入法环境变量已经设置。 最后一条命令中的 /实际路径/SiYuan 需要替换为电脑上思源程序的真实路径。这次并没有真正启动思源，所以还不能判断参数是否有效。 请在终端执行下面两条命令，并将输出贴过来，我们根据真实路径给你一条可以直接复制运行的命令： c ..
* 💬 [思源 3.8.3 版本麒麟系统无法输入中文](https://ld246.com/article/1789612119565/comment/1789634907495#comments)

  > 请先在终端执行以下命令，查看桌面会话和输入法： echo '$XDG_SESSION_TYPE' printenv GTK_IM_MODULE QT_IM_MODULE XMODIFIERS pgrep -a -f 'fcitx|ibus' 如果第一条输出 wayland，请彻底退出思源，包括托盘中的思源，再尝试： ' ..
* 💬 [uos 下 3.8.3 及 3.8.4 版都无法输入中文](https://ld246.com/article/1789627693596/comment/1789634879449#comments)

  > 请先彻底退出思源，包括托盘中的思源，然后在终端尝试： '/实际路径/SiYuan' --ozone-platform=x11 请将 '/实际路径/SiYuan' 替换为思源可执行文件的实际路径；如果使用 AppImage，就填写 AppImage 文件路径。 这个参数只让思源使用 X11 兼容后端，不需要切换整个桌面到 ..
* 💗💬 [移动端 3.8.4 数据库单选多选字段直接拖动就会排序](https://ld246.com/article/1789622279627/comment/1789629574999#comments)

  > @88250 有六个点图标的情况下只允许拖拽图标区域排序
* 💗📝 [移动端 3.8.4 数据库单选多选字段直接拖动就会排序](https://ld246.com/article/1789622279627)

  > 问题描述 在 3.8.4 和 3.8.3 均出现，数据库的字段是单选或者多选时，点击添加标签时，在移动端上只需要稍微移动就会拖动标签移动。正常应该是长按后才可以拖动排序，但是目前是只要点击拖动即=即可移动位置。 另外提一下，移动端数据库在点击单选和多选字段时，用户应该是希望首先自己去选择，找不到再点击搜索框搜索，而不是 ..
* 💬 [移动端 3.8.4 数据库单选多选字段直接拖动就会排序](https://ld246.com/article/1789622279627/comment/1789634360810#comments)

  > 感谢反馈，关联 [链接]


<!--events end -->
