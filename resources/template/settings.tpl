<button class="settings-button" onclick="showSettingsModal()">设置</button>
<!-- 设置模态框 -->
<div id="settingsModal" class="modal">
    <div class="modal-content">
        <span class="close">&times;</span>
        <h2>设置</h2>
        <label for="saveLocation">保存地址：</label>
        <input type="text" id="saveLocation" name="saveLocation" required>
        <label for="username">账号：</label>
        <input type="text" id="username" name="username" required>
        <label for="password">密码：</label>
        <input type="password" id="password" name="password" required>
        <button id="saveSettingsButton">保存</button>
    </div>
</div>

<script>
    function showSettingsModal() {
        const modal = document.getElementById('settingsModal');
        const closeBtn = modal.querySelector('.close');
        const saveButton = modal.querySelector('#saveSettingsButton');

        modal.style.display = 'block';

        closeBtn.onclick = function () {
            modal.style.display = 'none';
        };

        saveButton.onclick = function () {
            saveSettings();
        };

        // Fetch settings from the server and populate the form
        fetch('/settings', { method: 'GET' })
            .then(response => response.json())
            .then(data => {
                document.getElementById('saveLocation').value = data.saveLocation || '';
                document.getElementById('username').value = data.username || '';
                document.getElementById('password').value = data.password || '';
            })
            .catch(error => {
                alert("获取设置失败", error);
            });
    }

    function saveSettings() {
        const saveLocation = document.getElementById('saveLocation').value;
        const username = document.getElementById('username').value;
        const password = document.getElementById('password').value;

        const settingsData = {
            saveLocation: saveLocation,
            username: username,
            password: password
        };

        fetch('/settings', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(settingsData)
        })
            .then(response => {
                if (response.ok) {
                    alert("设置保存成功");
                    document.getElementById('settingsModal').style.display = 'none';
                } else {
                    alert("设置保存失败");
                }
            })
            .catch(error => {
                alert("设置保存失败", error);
            });
    }

    window.onclick = function (event) {
        const modal = document.getElementById('settingsModal');
        if (event.target == modal) {
            modal.style.display = 'none';
        }
    };
</script>
<style>
    /* 设置模态框样式 */
    .modal {
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

    .modal-content {
        background-color: #fefefe;
        margin: 15% auto;
        padding: 20px;
        border: 1px solid #888;
        width: 300px;
        /* 调整模态框宽度 */
        border-radius: 5px;
        /* 添加圆角 */
    }

    .modal-content label {
        display: block;
        margin-bottom: 5px;
        font-weight: bold;
    }

    .modal-content input[type="text"],
    .modal-content input[type="password"] {
        width: 100%;
        padding: 10px;
        margin-bottom: 15px;
        border: 1px solid #ccc;
        border-radius: 3px;
        box-sizing: border-box;
    }

    .modal-content button {
        width: 100%;
        padding: 10px;
        background-color: #007BFF;
        color: white;
        border: none;
        border-radius: 3px;
        cursor: pointer;
    }

    .modal-content button:hover {
        background-color: #0056b3;
    }

    .close {
        color: #aaa;
        float: right;
        font-size: 28px;
        font-weight: bold;
    }

    .close:hover,
    .close:focus {
        color: black;
        text-decoration: none;
        cursor: pointer;
    }

    /* 设置按钮样式 */
    .settings-button {
        position: fixed;
        top: 10px;
        left: 10px;
        background-color: #007BFF;
        color: white;
        padding: 10px 20px;
        border: none;
        border-radius: 5px;
        cursor: pointer;
        z-index: 1000;
    }

    .settings-button:hover {
        background-color: #0056b3;
    }
</style>