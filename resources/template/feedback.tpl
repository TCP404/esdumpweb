<button class="feedback-button" onclick="sendFeedback()">错误反馈</button>
<script>
    function sendFeedback() {
        fetch('/feedback', { method: 'GET' })
            .then(response => response.json())
            .then(data => {
                console.log('获取反馈信息成功，返回数据：', data);
                alert("反馈成功");
            })
            .catch(error => {
                alert("获取反馈信息失败", error);
            });
    }
</script>
<style>
    /* 错误反馈按钮样式 */
    .feedback-button {
        position: fixed;
        top: 50px;
        left: 10px;
        background-color: #FF5733;
        color: white;
        padding: 10px 20px;
        border: none;
        border-radius: 5px;
        cursor: pointer;
        z-index: 1000;
    }

    .feedback-button:hover {
        background-color: #C70039;
    }
</style>