<form id="myForm">
    <label for="host">主机：</label>
    <select id="host" name="host" required>
        {{ if eq .host "agg" }}
        <option value="agg" selected>agg</option>
        <option value="online">online</option>
        {{ else }}
        <option value="agg">agg</option>
        <option value="online" selected>online</option>
        {{ end }}
    </select>
    <br>

    <label for="index">索引：</label>
    <input type="text" id="index" name="index" required value="{{.index}}">

    <label for="timeField">时间字段：</label>
    <select id="timeField" name="timeField" required>
        {{ if eq .timeField "insert_time" }}
        <option value="insert_time" selected>insert_time</opti on>
        <option value="event_time">event_time</option>
        {{ else }}
        <option value="insert_time">insert_time</option>
        <option value="event_time" selected>event_time</option>
        {{ end }}
    </select>
    <br>

    <label for="startTime">起始时间：</label>
    <input type="datetime-local" id="startTime" name="startTime" required>
    <br>

    <label for="endTime">结束时间：</label>
    <input type="datetime-local" id="endTime" name="endTime" required>
    <br>

    <label for="product">产品：</label>
    <input type="text" id="product" name="product" value="{{.product}}">
    <br>

    <label for="condition">条件：</label>
    <button class="collapsible" onclick="toggleTextarea(event)">编辑条件</button>
    <div class="content" id="textareaContent">
        <textarea id="condition" name="condition" rows="10" cols="50"></textarea>
    </div>
    <br>

    <input type="submit" value="提交">
</form>

<script src="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.65.13/codemirror.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.65.13/mode/javascript/javascript.min.js"></script>
<script>
    document.getElementById('myForm').addEventListener('submit', function (e) {
        e.preventDefault(); // 阻止表单默认提交行为
        const formDataObj = getFormData(this);
        if (!formDataObj) return; // 如果表单数据无效，则退出
        sendFormRequest(formDataObj, false);
    });

    function getFormData(form) {
        const formData = new FormData(form);
        const formDataObj = {};
        formData.forEach((value, key) => {
            if (key === 'startTime' || key === 'endTime') {
                value = value.replace('Z', '') + ':00+08:00';
            }
            formDataObj[key] = value;
        });

        // 检查JSON编辑框中的内容是否语法正确
        const conditionValue = editor.getValue();
        if (conditionValue.length !== 0) {
            try {
                JSON.parse(conditionValue);
                formDataObj['condition'] = conditionValue;
            } catch (error) {
                alert('“条件”输入的JSON格式不正确，请检查后重新提交。');
                return null;
            }
        }
        return formDataObj;
    }

    function sendFormRequest(formDataObj, retried) {
        const url = '/dump';
        const method = 'POST';
        const header = { 'Content-Type': 'application/json' };
        const jsonData = JSON.stringify(formDataObj);  // 将表单数据转换为JSON字符串

        fetch(url, { method: method, headers: header, body: jsonData })
            .then(response => {
                if (response.status === 403 && !retried) {
                    showLoginModal(formDataObj["host"]).then(success => {
                        if (success) {
                            sendFormRequest(formDataObj, true); // 重新发送请求
                        }
                    });
                    return;
                }
                return response.json();
            })
            .then(data => {
                if (data) {
                    console.log('请求成功，返回数据：', data);
                    // 弹窗显示请求结果
                    if (data.total > 0) {
                        showAlert("任务创建成功，数据量：" + data.total + "，正在下载. \n\n请到" + data.output + "查收");
                    } else {
                        showAlert("任务创建成功，但数据量为0，无需下载");
                    }
                }
            })
            .catch(error => {
                alert("好像还没启动 esdump ?");
                console.error('请求出错：', error);
            });
    }

    function toggleTextarea(event) {
        event.preventDefault(); // 阻止默认行为
        var content = document.getElementById("textareaContent");
        if (content.style.display === "block") {
            content.style.display = "none";
        } else {
            content.style.display = "block";
        }
    }

    const editor = CodeMirror.fromTextArea(document.getElementById('condition'), {
        mode: "application/json", // 设置为JSON模式，以获得更好的语法支持
        lineNumbers: true,
        tabSize: 2,
        indentWithTabs: false,
        gutters: ["CodeMirror-lint-markers"], // 添加用于显示语法错误标记的 gutter
        lint: true // 开启语法检查功能
    });
</script>