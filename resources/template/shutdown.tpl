<div id="shutdown-wrapper">
    <button class="shutdown-button" onclick="shutdown()">结束</button>
</div>
<script>
    function shutdown() {
        fetch('/shutdown', { method: 'GET' })
            .then(resp => {
                alert(resp.ok ? '拜拜了您嘞~👋🏻' : '拜了个拜~👋🏻');
                if (resp.ok) {
                    const submitButton = document.querySelector('input[type="submit"]');
                    submitButton.disabled = true;
                    submitButton.value = '服务已关闭';
                    submitButton.style.backgroundColor = 'gray';
                }
            });
    }
</script>