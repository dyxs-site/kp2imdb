以下是`kp2imdb v1.0.2`文档的整理：

---

# kp2imdb v1.0.2

[项目主页](https://github.com/oklookat/kp2imdb)  
[捐赠链接](https://donationalerts.com/r/oklookat)  
[Boosty 捐赠链接](https://boosty.to/oklookat/donate)

## 简介

该程序将来自Kинопоиск的电影（标题）转换为IMDB上的ID列表。

## 使用方法

```bash
./kp2imdb -m file.json
```

`-m` 标志允许在出现错误或找不到标题时手动输入IMDB ID。

程序将开始匹配标题，并生成一个名为 `links.json` 的文件，其中包含已匹配的标题。完成后，将创建一个格式为 `{当前时间_unix}.txt` 的文件。TXT文件中将包含按行分隔的IMDB ID列表。

TXT文件中的ID列表可以在IMDB List Importer脚本中使用：
[IMDB List Importer](https://greasyfork.org/en/scripts/23584-imdb-list-importer)

## 如何获取JSON文件？

### 1. 自行创建

示例JSON文件：

```json
[
    {"id":"462762","name":"铁人三部曲","alt_name":"Iron Man Three (2013) 125 min."},
    {"id":"471158","name":"拯救世界","alt_name":" (2010) 97 min."},
    {"id":"404900","name":"绝命毒师（剧集）","alt_name":"Breaking Bad (2008-2013) 47 min."},
    {"id":"655800","name":"黑镜（剧集）","alt_name":"Black Mirror (2011-...) 43 min."}
]
```

`id` 是Kинопоиск上的ID。`name` 和 `alt_name` 必须正确填写，示例中展示了电影和剧集的名称格式。

### 2. 使用Kинопоиск + TamperMonkey脚本

下载链接：
- [kpexport.js](https://github.com/oklookat/kp2imdb/blob/main/kpexport.js)
- [Kинопоиск Folder Exporter](https://greasyfork.org/en/scripts/487107-kinopoisk-folder-exporter)

简要指南：
1. 安装脚本（具体安装步骤不在此文档中）。
2. 进入Kинопоиск上的文件夹（例如收藏的电影）。
3. 将文件夹的“显示”设置为最大值（例如200）。
4. 在“全选”旁点击红色背景，下载JSON文件。
5. 如果有多个页面，转到下一页并重复第4步。

## 匹配过程

1. 在IMDB上搜索格式为 "name (year)" 的标题。例如："复仇者联盟 2012"。
2. 如果未找到电影或出现错误，使用 `alt_name` 重新进行搜索。
3. 如果仍然未找到电影或出现错误，可以使用 `-m` 标志手动输入ID。
4. 如果获得IMDB电影ID，则将其以 "id_кинопоиска:id_imdb" 的格式保存到 `links.json` 中。

需要注意的是，一些电影可能会被错误匹配或无法找到。这可能是由于电影的标题不常见或是用俄语表示的原因。

---
