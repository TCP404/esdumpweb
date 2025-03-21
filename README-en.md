# esdump

## Description
`esdump` is a **tool** for exporting data from Elasticsearch.

## Usage

1. Go to the [releases page](https://github.com/TCP404/esdumpweb/releases) and download the latest version of `esdump` for your OS.
2. Unzip the file and place it in a directory of your choice.
3. Click on the `esdump` executable to run it. Or run it from the command line with the following command:
   ```bash
   cd /path/to/esdump
   chmod +x esdump
   ./esdump
   ```

> Then the `esdump` executable will run up a web server on port 80. You can access the web interface by opening your web browser and navigating to `http://localhost:80`.
>
> Okay, now you can see a form with the following fields:
>
> - **保存类型**
>   - `xlsx` - Export data in xlsx format.
>   - `csv` - Export data in csv format.
> - **主机**: The host of the Elasticsearch server.
>   - `online` - The online Elasticsearch server.
>   - `agg` - The aggregation Elasticsearch server.
> - **索引**: The index to export data from.
> - **时间字段**: The time field to filter data by. It's a required field. You can not query data without the time field because the total data is too large.
>   - `insert_time` - The so-called 捕获时间
>   - `event_time` - The so-called 发生时间
> - **开始时间**: The start time to filter data by. It's a required field.
> - **结束时间**: The end time to filter data by. It's a required field.
> - **产品**: The product to filter data by. It's an optional field.
> - **条件**: The query to filter data by. It's an optional field. You must use the `JSON` format and the key must one of `filter`, `should`, `must`, `must_not`. For example:
>   ```json
>   {
>     "filter": [{your filter1}],
>     "should": [{your should1}, {your should2}],
>     "must": [],
>     "must_not": []
>   }
>   ```
>   and `esdump` will append them to the `query.bool`. The `filter`, `should`, `must`, and `must_not` keys are all optional. Finally, the `esdump` will generate a query like this:
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
>                     "product": "product" // Only if you select `online` as the host
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

4. Click the button `submit` to submit the form. The `esdump` will start exporting data from Elasticsearch and save it to a file in the `YOUR_DOWNLOAD_DIRECTORY`. The file name will be in the format of `indexname_product_timefield_YYYYMMDD_HHMMSS.extension`, where `index_name` is the name of the index, `YYYYMMDD_HHMMSS` is the current date and time, and `extension` is the file extension you selected.
5. If you are not logging in to the Elasticsearch server, you will be prompted to enter the username and password. After you enter the username and password, the `esdump` will start exporting data from Elasticsearch and save it to a file in the `YOUR_CONFIG_DIRECTORY/esdumpweb/config.json`.
6. You can click the button `Login out` to drop your username and password from config file, the button `Error feedback` to send a runtime log file to feishu group.
7. You can click the button `Exit` to exit the `esdump` web interface. The `esdump` will stop the web server and exit.