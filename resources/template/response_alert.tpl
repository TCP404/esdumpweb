<!-- 成功响应时的弹窗，展示下载地址 -->
<div id="customAlert" class="custom-alert">
    <div class="custom-alert-content">
        <span class="close-btn" onclick="closeAlert()">&times;</span>
        <p id="alertMessage"></p>
    </div>
</div>
<script>
    function showAlert(message) {
        document.getElementById('alertMessage').innerText = message;
        document.getElementById('customAlert').style.display = 'block';
    }
    function closeAlert() {
        document.getElementById('customAlert').style.display = 'none';
    }
</script>