# 写在前面的话
**这是一个基于自己兴趣编写的项目，虽然代码写的很差，功能简单不存在独立的知识产权，但是还是希望各位资本大大手下留情，不要直接将项目pull之后用做商用，本项目也不会用做商用。**
# 对于这个项目简单的描述
这是一个没有考虑用户体验的winform应用，首先说下简单的功能
 - .exe打开时，会先加载粤语对应的字典。字典是一个`txt`文件，文件中默认第一列为`简体中文`、第二列为`繁体中文`、第三列为`粤语音标`。该文件是相对路径`./resource/dict/dict.txt`，并且是固定文件名称和类型。可以被替换，只需要文件名称和类型相同，然后覆盖该文件即可。
 - 字典的使用，实际是在加载时生成了两组`key-value`的map类型，因此实际是通过中文`key`去取值音标`value`。
 - 字典来源于https://github.com/YuhangGe/yueyu，明确的说，我拿这个字典没有获取到这个作者的授权，因此如果确实涉及到侵权问题，请这个作者联系以作者的账户给我留言，我将及时删除该字典文件。
 - 该项目采用了fyne/v2的gui模版，在打开资源文件管理器的界面由于没有找到修改该界面样式的方法，因此样式默认为liunx的样式。
# 对于这个应用的功能描述
这个应用的功能非常非常简单，只有三个按钮
 - 打开文件，打开需要标记音标的`.txt`文件，会在按钮下方打印你打开的文本文件内容
 - 音标，将文件中的中文内容转换为音标，并以`音标\n中文`的方式打印出来
 - 保存，保存转换后的音标文件，并以原文件路径，以及`原文件_tmp.原文件类型`类型保存，例如：原文件`E:/handld.txt`，保存后的文件`E:/handld_tmp.txt`
# 关于这个应用中我认为的缺陷，即：可能在后续不定时的更新中会进行优化的地方
 1. 对于多音字的处理，目前只能在保存后使用文本编辑器进行处理，推荐使用超级文本编辑器(例如:vscode,subline txt等；个人主观不推荐n++)，原因是在转换的过程中多音字的音标前会加上`#数字$数字`，中文里面也会加上对应内容，方便识别
 2. 多音字的功能原来是准备采用界面点击选择的方式，但是由于fyne对于文字的处理，不fyne使用不熟练的情况下，只能退而求其次，在应用之外由使用人处理。
 3. 关于中文标记粤语音标的功能，目前还是推荐使用[百度api](https://fanyi-api.baidu.com/product/11)，需要注意的是免费的api只返回翻译的内容，没有读音。读音属于收费内容。
 # 关于使用
 获取后直接解压`bin.zip`，打开其中唯一的一个`.exe`文件即可