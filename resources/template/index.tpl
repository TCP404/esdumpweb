<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>表单示例</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.65.13/codemirror.min.css">
    <style>
        /* 整体表单样式 */
        form {
            background-color: #f9f9f9;
            border: 1px solid #ddd;
            border-radius: 5px;
            padding: 20px;
            width: 400px;
            margin: 0 auto;
        }

        /* 表单标签样式 */
        label {
            display: block;
            margin-bottom: 5px;
            font-weight: bold;
        }

        /* 表单输入框、下拉框、日期时间选择器样式 */
        input[type="text"],
        select,
        input[type="datetime-local"] {
            box-sizing: border-box;
            width: 100%;
            padding: 10px;
            margin-bottom: 15px;
            border: 1px solid #ccc;
            border-radius: 3px;
        }

        /* 对日期时间选择器添加额外样式 */
        input[type="datetime-local"] {
            appearance: None;
            -webkit-appearance: None;
            -moz-appearance: None;
            background-color: white;
            font-size: 14px;
        }

        /* 文件选择器相关的隐藏输入框样式 */
        input[type="file"] {
            opacity: 0;
            position: absolute;
            z-index: -1;
        }

        /* 保存地址按钮样式 */
        #saveLocationButton {
            display: inline-block;
            background-color: #007BFF;
            color: white;
            padding: 10px 20px;
            border-radius: 3px;
            cursor: pointer;
            text-align: center;
            line-height: 1;
            vertical-align: middle;
            margin-bottom: 15px;
        }

        /* 提交按钮样式 */
        input[type="submit"] {
            background-color: #007BFF;
            color: white;
            padding: 10px 20px;
            border: none;
            border-radius: 3px;
            cursor: pointer;
        }

        input[type="submit"]:hover {
            background-color: #0056b3;
        }
    </style>
    <style>
        /* 条件 */
        .collapsible {
            background-color: #f1f1f1;
            color: black;
            cursor: pointer;
            padding: 10px;
            width: 100%;
            border: none;
            text-align: left;
            outline: none;
            font-size: 15px;
        }

        .content {
            padding: 0 18px;
            display: none;
            overflow: hidden;
            background-color: #f9f9f9;
            resize: both;
            /* 允许调整大小 */
            overflow: auto;
        }

        textarea {
            width: 100%;
            height: 100%;
            resize: none;
            /* 禁止调整大小 */
        }
    </style>
    <style>
        /* 弹窗样式 */
        .custom-alert {
            display: none;
            position: fixed;
            z-index: 1;
            left: 0;
            top: 0;
            width: 100%;
            height: 100%;
            overflow: auto;
            background-color: rgb(0, 0, 0);
            background-color: rgba(0, 0, 0, 0.4);
        }

        .custom-alert-content {
            background-color: #fefefe;
            margin: 15% auto;
            padding: 20px;
            border: 1px solid #888;
            width: 80%;
        }

        .close-btn {
            color: #aaa;
            float: right;
            font-size: 28px;
            font-weight: bold;
        }

        .close-btn:hover,
        .close-btn:focus {
            color: black;
            text-decoration: none;
            cursor: pointer;
        }
    </style>
    <style>
        /* 结束按钮样式 */
        .shutdown-button {
            position: absolute;
            top: 10px;
            right: 10px;
            color: white;
            background-color: red;
            font-weight: bold;
            border: none;
            padding: 10px 20px;
            font-size: 16px;
            border: none;
            border-radius: 5px;
        }
    </style>
</head>

<body>
    <!-- <button class="settings-button" onclick="showSettingsModal()">设置</button> -->
    {{ template "response_alert.tpl" }}

    {{ template "form.tpl" }}

    {{ template "feedback.tpl" }}

    {{ template "logout.tpl" }}

    {{ template "shutdown.tpl" }}

    {{ template "login.tpl" }}

    <script>
        // 获取昨天的日期并格式化为指定格式
        function getYesterdayFormatted() {
            const yesterday = new Date();
            yesterday.setDate(yesterday.getDate() - 1);
            return yesterday.toISOString().slice(0, 16);
        }

        // 获取今天的日期并格式化为指定格式
        function getTodayFormatted() {
            const today = new Date();
            return today.toISOString().slice(0, 16);
        }

        window.onload = function () {
            // 设置起始时间和结束时间的默认值
            const startTimeInput = document.getElementById('startTime');
            const endTimeInput = document.getElementById('endTime');
            startTimeInput.value = getYesterdayFormatted();
            endTimeInput.value = getTodayFormatted();
        }
    </script>
</body>

</html>