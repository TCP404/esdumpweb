<button class="logout-button" onclick="sendLogout()">登出</button>
<script>
    function sendLogout() {
        fetch("/logout", { method: "GET" })
            .then(response => response.json())
            .then(data => {
                console.log('登出成功，返回数据：', data);
                alert("登出成功");
            })
            .catch(error => {
                alert("登出失败", error);
            });
    }
</script>
<style>
    /* 错误反馈按钮样式 */
    .logout-button {
        position: fixed;
        top: 10px;
        left: 10px;
        background-color: #FF5733;
        color: white;
        padding: 10px 20px;
        border: none;
        border-radius: 5px;
        cursor: pointer;
        z-index: 1000;
    }

    .logout-button:hover {
        background-color: #C70039;
    }
</style>