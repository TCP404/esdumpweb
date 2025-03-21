# esdump

## 简介
`esdump` 是一个用于从 Elasticsearch 导出数据的 **工具**。

## 使用方法

1. 前往 [releases 页面 https://github.com/TCP404/esdumpweb/releases](https://github.com/TCP404/esdumpweb/releases)，下载适合您操作系统的最新版本 `esdump`。
2. 解压文件并将其放置在您选择的目录中。
3. 点击 `esdump` 可执行文件运行它，或者通过命令行运行以下命令：
   ```bash
   cd /path/to/esdump
   chmod +x esdump
   ./esdump
   ```
   
   然后，`esdump` 可执行文件将在端口 80 上启动一个 Web 服务器。您可以通过打开浏览器并访问 `http://localhost:80` 来访问 Web 界面。

> 您将看到一个包含以下字段的表单：
> 
> - **保存类型**
>   - `xlsx` - 以 xlsx 格式导出数据。
>   - `csv` - 以 csv 格式导出数据。
> - **主机**: Elasticsearch 服务器的主机。
>   - `online` - online 服务器。
>   - `agg` - agg 服务器。
> - **索引**: 要导出数据的索引。
> - **时间字段**: 用于过滤数据的时间字段。此字段为必填项。由于数据量过大，无法在没有时间字段的情况下查询数据。
>   - `insert_time` - 所谓的`捕获时间`。
>   - `event_time` - 所谓的`发生时间`。
> - **开始时间**: 用于过滤数据的开始时间。此字段为必填项。
> - **结束时间**: 用于过滤数据的结束时间。此字段为必填项。
> - **产品**: 用于过滤数据的产品。此字段为可选项。
> - **条件**: 用于过滤数据的查询条件。此字段为可选项。必须使用 `JSON` 格式，键必须是 `filter`、`should`、`must` 或 `must_not` 之一。例如：
>   ```json
>   {
>     "filter": [{your filter1}],
>     "should": [{your should1}, {your should2}],
>     "must": [],
>     "must_not": []
>   }
>   ```
>   然后 `esdump` 会将它们附加到 `query.bool` 中。`filter`、`should`、`must` 和 `must_not` 键均为可选项。最终，`esdump` 将生成如下查询：
>   ```json
>   {
>     "query": {
>       "bool": {
>         "filter": [
>             {
>                 "range": {
>                     "time_field": {
>                         "gte": "start_time",
>                         "lte": "end_time"
>                     }
>                 }
>             },
>             {
>                 "term": {
>                     "product": "product" // 仅当选择 `online` 作为主机时
>                 }
>             },
>             {your filter1}
>         ],
>         "should": [{your should1}, {your should2}],
>         "must": [],
>         "must_not": []
>       }
>     }
>   }
>   ```

4. 点击 `提交` 按钮提交表单。`esdump` 将开始从 Elasticsearch 导出数据，并将其保存到 `YOUR_DOWNLOAD_DIRECTORY` 中的文件中。文件名格式为 `indexname_product_timefield_YYYYMMDD_HHMMSS.extension`，其中 `index_name` 是索引名称，`YYYYMMDD_HHMMSS` 是当前日期和时间，`extension` 是您选择的文件扩展名。
5. 如果您未登录 Elasticsearch 服务器，系统会提示您输入用户名和密码。输入用户名和密码后，`esdump` 将开始从 Elasticsearch 导出数据，并将其保存到 `YOUR_CONFIG_DIRECTORY/esdumpweb/config.json` 文件中。
6. 您可以点击 `登出` 按钮从配置文件中删除您的用户名和密码，点击 `错误反馈` 按钮将运行时日志文件发送到飞书群。
7. 您可以点击 `结束` 按钮退出 `esdump` Web 界面。`esdump` 将停止 Web 服务器并退出。