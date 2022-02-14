## 说明
该程序是根据emby刮削生成的nfo文件来将文件名称整理成统一格式，让你的影视目录更整洁。

如果你未使用emby或者媒体库未开启 **将媒体图像保存到媒体文件夹中** 功能，请勿使用。

该程序不具备撤销能力，如需使用，请自行备份以免造成文件丢失。请务必要先按照以下方式生成测试文件，查看效果。

## 调试
### 文件准备
```shell
# 复制文件结构到新目录便于测试
cp -r --attributes-only /目标目录/ /测试目录/

# 复制目标目录下的所有nfo文件覆盖到测试目录
cd /目标目录/
find -iname "*.nfo" -exec cp {} /测试目录/{} \;
```
### 配置文件说明
复制 **config.json.example** 文件到 **config.json** ，并按需更改配置
```json
{
  "movieRename": true,# 电影重命名开关
  "movieDirPath": "example/originData/电影/",# 需要改名的目标文件夹
  "movieRootPath": "example/newData/电影/",# 用于保存新数据的文件夹
  "tvRename": true,# 电视剧重命名开关
  "tvDirPath": "example/originData/电视剧/",# 需要改名的目标文件夹
  "tvRootPath": "example/newData/电视剧/",# 用于保存新数据的文件夹
  "movieDirFormat": "{originaltitle} ({year}) [imdbid={imdbid}]",# 电影目录格式(花括号里的为变量)
  "movieTitleFormat": "{originaltitle}",# 电影文件名格式
  "tvDirFormat": "{originaltitle} ({year})",# 电视剧目录格式
  "episodeDirFormat": "Season {season}",# 电视剧季目录格式
  "episodeTitleFormat": "{originaltitle} S{SEASON}E{EPISODE}"# 电视剧文件名格式
}
```
### 重命名变量说明
| 变量 | 备注 | 例如 |
| --- | --- | --- | 
| {originaltitle} | 原始名称 | 살인자의 기억법 |
| {title} | 显示名称 | 杀人者的记忆法 |
| {year} | 发行年份 | 2017 |
| {season} | 第几季 | 3 |
| {SEASON} | 第几季 | 03 |
| {episode} | 第几集 | 5 |
| {episodetitle} | 集标题 |第一集标题 |
| {EPISODE} | 第几集 | 05 |
| {imdbid} | imdb id | tt5729348 |
| {tmdbid} | tmdb id | 432836  |
| {tvdbid} | tvdb id | 414615  |
| {releasedate} | 发行日期   | 2017-09-07 |
| {country} | 国家     | South Korea |
| {id} | id     | tt5729348 |

其中 **{season}** **{SEASON}** **{episode}** **{tvdbid}** **{episodetitle}** **{EPISODE}** 为电视独有

备注为空的说明字段有可能为空

### 运行
```
# 下载程序
git clone git@github.com:qcgzxw/embyRenamer.git
# 复制生成配置文件
copy config.json.example config.json
# 编辑配置文件 按照以上规则更改文件目录和相关配置
vim config.json
# 运行测试代码
go run main.go
```

## 预览
### 电影目录结构
默认采用emby官方推荐的格式 **Name (Year) [tmdbid=xxxx]**
```
└── Movie
    └── Name (Year) [tmdbid=xxxx]
        ├── Name.mkv
        ├── Name.nfo
        ├── poster.jpg
        └── fanart.jpg
```

[BEFORE](example/originData/电影目录结构.txt)

[AFTER](example/newData/电影目录结构.txt)

### 电视剧目录结构
```
└── TV
    └── Name (2021)
        ├── fanart.jpg
        ├── poster.jpg
        ├── Season 1
        │   ├── Name S01E01.mkv
        │   ├── Name S01E01.nfo
        │   ├── Name S01E01-thumb.jpg
        │   ├── Name S01E02.mkv
        │   ├── Name S01E02.nfo
        │   ├── Name S01E02-thumb.jpg
        │   ├── Name S01E03.mkv
        │   ├── Name S01E03.nfo
        │   └── Name S01E03-thumb.jpg
        └── tvshow.nfo
```

[BEFORE](example/originData/电视剧目录结构.txt)

[AFTER](example/newData/电视剧目录结构.txt)

## 版本变更

### v1.1
- 同一文件，多版本支持(https://github.com/qcgzxw/embyRenamer/issues/5)
### v1.0
- 电影文件重命名
- 电视剧文件重命名
