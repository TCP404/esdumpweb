<!-- 登录模态框 -->
<div id="loginModal" class="modal">
    <div class="modal-content">
        <span class="close">&times;</span>
        <h2>登录</h2>
        <label for="username">账号：</label>
        <input type="text" id="username" name="username" required>
        <label for="password">密码：</label>
        <input type="password" id="password" name="password" required>
        <button id="loginButton">登录</button>
    </div>
</div>

<script>
    function showLoginModal(host) {
        return new Promise((resolve) => {
            const modal = document.getElementById('loginModal');
            const closeBtn = modal.querySelector('.close');
            const loginButton = modal.querySelector('#loginButton');
            const usernameInput = modal.querySelector('#username');
            const passwordInput = modal.querySelector('#password');

            modal.style.display = 'block';

            closeBtn.onclick = function () {
                modal.style.display = 'none';
                resolve(false);
            };

            loginButton.onclick = function () {
                submitLoginForm(host, resolve);
            };

            // 监听回车键提交
            usernameInput.onkeypress = function (event) {
                if (event.key === 'Enter') {
                    event.preventDefault();
                    submitLoginForm(host, resolve);
                }
            };

            passwordInput.onkeypress = function (event) {
                if (event.key === 'Enter') {
                    event.preventDefault();
                    submitLoginForm(host, resolve);
                }
            };
        });
    }

    function submitLoginForm(host, resolve) {
        const modal = document.getElementById('loginModal');
        const username = modal.querySelector('#username').value;
        const password = modal.querySelector('#password').value;
        if (!username || !password) {
            alert('请输入账号和密码');
            return;
        }

        const method = "POST";
        const url = "/login";
        const header = { "Content-Type": "application/json" };
        const jsonData = JSON.stringify({ "username": username, "password": password, "host": host });
        fetch(url, { method: method, headers: header, body: jsonData })
            .then(response => {
                modal.style.display = 'none';
                resolve(response.status === 200);
            })
            .catch(error => {
                alert("登录失败", error);
                resolve(false);
            });
    }

    window.onclick = function (event) {
        const modal = document.getElementById('loginModal');
        if (event.target == modal) {
            modal.style.display = 'none';
        }
    };
</script>
<style>
    /* 登录模态框样式 */
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
        width: 300px; /* 调整模态框宽度 */
        border-radius: 5px; /* 添加圆角 */
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
</style>
