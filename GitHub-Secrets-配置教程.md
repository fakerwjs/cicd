# GitHub Secrets 配置教程

本教程详细说明如何配置 CI/CD 流程中需要的 GitHub Secrets。

## 📋 目录

- [阿里云容器镜像服务配置](#阿里云容器镜像服务配置)
  - [ALIYUN_USERNAME](#aliyun_username)
  - [ALIYUN_PASSWORD](#aliyun_password)
  - [ALIYUN_NAMESPACE](#aliyun_namespace)
- [服务器部署配置](#服务器部署配置)
  - [SERVER_HOST](#server_host)
  - [SERVER_USER](#server_user)
  - [SERVER_SSH_KEY](#server_ssh_key)
- [如何在 GitHub 添加 Secrets](#如何在-github-添加-secrets)

---

## 阿里云容器镜像服务配置

### ALIYUN_USERNAME

**作用**：阿里云容器镜像服务的登录用户名

**获取步骤**：

1. 登录 [阿里云容器镜像服务控制台](https://cr.console.aliyun.com/)
2. 选择 **个人实例** 或 **企业实例**（本项目使用的是上海区域的个人实例）
3. 点击左侧菜单 **访问凭证**
4. 在页面中找到 **用户名**，通常格式为：`your-aliyun-account@example.com` 或阿里云账号ID

**示例值**：
```
your-aliyun-account@aliyun.com
```

---

### ALIYUN_PASSWORD

**作用**：阿里云容器镜像服务的登录密码（固定密码）

**获取步骤**：

1. 在 [阿里云容器镜像服务控制台](https://cr.console.aliyun.com/) 的 **访问凭证** 页面
2. 找到 **固定密码** 部分
3. 如果未设置，点击 **重置Docker登录密码** 或 **设置固定密码**
4. 设置一个新密码（建议使用强密码，包含大小写字母、数字和特殊字符）
5. ⚠️ **重要**：密码设置后无法查看，请务必保存好

**注意事项**：
- 这个密码与阿里云账号密码不同，是专门用于 Docker 登录的
- 如果忘记，只能重置，无法找回

---

### ALIYUN_NAMESPACE

**作用**：阿里云容器镜像仓库的命名空间

**获取步骤**：

1. 在 [阿里云容器镜像服务控制台](https://cr.console.aliyun.com/)
2. 点击左侧菜单 **命名空间**
3. 查看已有的命名空间列表，或者创建新的命名空间
4. 复制命名空间名称

**创建新命名空间**：
1. 点击 **创建命名空间**
2. 输入命名空间名称（例如：`myapps`、`production`、`your-project`）
3. 选择 **公开** 或 **私有**（建议选择私有）
4. 点击确定

**示例值**：
```
myapps
```

**当前项目使用的命名空间**：根据你的配置应该是 `myapps` 或类似的名称

---

## 服务器部署配置

### SERVER_HOST

**作用**：部署目标服务器的 IP 地址或域名

**获取步骤**：

1. 登录你的云服务器控制台（阿里云ECS、腾讯云、AWS等）
2. 找到服务器实例，查看 **公网IP** 或 **弹性IP**
3. 复制 IP 地址

**示例值**：
```
123.45.67.89
```

或者使用域名：
```
myserver.example.com
```

---

### SERVER_USER

**作用**：SSH 登录服务器的用户名

**获取步骤**：

通常是服务器的默认用户或你创建的用户：

- **Ubuntu/Debian 系统**：默认用户通常是 `ubuntu` 或 `root`
- **CentOS/RHEL 系统**：默认用户通常是 `root` 或 `centos`
- **自定义用户**：如果你创建了其他用户，使用该用户名

**示例值**：
```
root
```

或
```
ubuntu
```

**注意事项**：
- 建议不要使用 `root` 用户，而是创建一个具有 sudo 权限的普通用户
- 确保该用户有权限执行 Docker 命令（加入 docker 用户组）

---

### SERVER_SSH_KEY

**作用**：SSH 私钥，用于免密登录服务器

**获取步骤**：

#### 方法一：使用已有的 SSH 密钥

1. 在本地电脑打开终端
2. 查看是否已有 SSH 密钥：
   ```bash
   ls -la ~/.ssh/
   ```
3. 如果存在 `id_rsa` 或 `id_ed25519` 等文件，查看私钥内容：
   ```bash
   cat ~/.ssh/id_rsa
   ```
4. 复制完整的私钥内容（包括 `-----BEGIN RSA PRIVATE KEY-----` 和 `-----END RSA PRIVATE KEY-----`）

#### 方法二：生成新的 SSH 密钥对

1. 在本地终端执行：
   ```bash
   ssh-keygen -t ed25519 -C "your-email@example.com" -f ~/.ssh/deploy_key
   ```
   
   或使用 RSA：
   ```bash
   ssh-keygen -t rsa -b 4096 -C "your-email@example.com" -f ~/.ssh/deploy_key
   ```

2. 按提示操作（可以不设置密码，直接回车）

3. 查看私钥：
   ```bash
   cat ~/.ssh/deploy_key
   ```

4. 查看公钥：
   ```bash
   cat ~/.ssh/deploy_key.pub
   ```

#### 将公钥添加到服务器

1. 登录到你的服务器
2. 编辑 `authorized_keys` 文件：
   ```bash
   vim ~/.ssh/authorized_keys
   ```
3. 将公钥内容（`deploy_key.pub` 的内容）粘贴到文件末尾
4. 保存退出

5. 设置正确的权限：
   ```bash
   chmod 700 ~/.ssh
   chmod 600 ~/.ssh/authorized_keys
   ```

#### 测试连接

在本地测试 SSH 连接是否成功：
```bash
ssh -i ~/.ssh/deploy_key user@server-ip
```

**示例值**（私钥格式）：
```
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACBK9...（很长的密钥内容）...
-----END OPENSSH PRIVATE KEY-----
```

**⚠️ 重要提示**：
- 私钥是敏感信息，绝对不要泄露或提交到代码仓库
- 只将私钥添加到 GitHub Secrets，不要公开
- 公钥添加到服务器，私钥保存在本地和 GitHub Secrets

---

## 如何在 GitHub 添加 Secrets

### 步骤说明

1. **打开 GitHub 仓库**
   - 进入你的项目仓库页面

2. **进入 Settings**
   - 点击仓库顶部的 **Settings** 标签

3. **找到 Secrets and variables**
   - 在左侧菜单找到 **Security** → **Secrets and variables** → **Actions**

4. **添加新 Secret**
   - 点击右上角的 **New repository secret** 按钮
   - 输入 **Name**（Secret 名称，必须与 CI/CD 配置中的名称完全一致）
   - 输入 **Value**（Secret 的值）
   - 点击 **Add secret**

5. **重复以上步骤**，添加所有需要的 Secrets

### 需要添加的所有 Secrets 清单

| Secret 名称 | 说明 | 示例 |
|------------|------|------|
| `ALIYUN_USERNAME` | 阿里云容器镜像服务用户名 | `your-account@aliyun.com` |
| `ALIYUN_PASSWORD` | 阿里云容器镜像服务密码 | `YourStrongPassword123!` |
| `ALIYUN_NAMESPACE` | 阿里云镜像仓库命名空间 | `myapps` |
| `SERVER_HOST` | 服务器 IP 或域名 | `123.45.67.89` |
| `SERVER_USER` | SSH 登录用户名 | `ubuntu` |
| `SERVER_SSH_KEY` | SSH 私钥（完整内容） | `-----BEGIN OPENSSH PRIVATE KEY-----...` |

### 验证 Secrets 配置

1. 所有 Secrets 添加完成后，在 **Actions** 页面可以看到列表
2. Secret 的值添加后无法查看，只能删除或更新
3. 推送代码到 `main` 分支，触发 CI/CD 流程
4. 在 **Actions** 标签页查看工作流运行结果
5. 如果出现认证失败，检查对应的 Secret 值是否正确

---

## 📝 注意事项

1. **安全性**
   - 所有 Secrets 都是敏感信息，切勿泄露
   - 定期更新密码和密钥
   - 使用最小权限原则

2. **命名规范**
   - Secret 名称必须与 CI/CD 配置文件中的完全一致
   - 区分大小写

3. **阿里云镜像仓库地址**
   - 当前项目使用的是：`crpi-kc1xjy517ehrbwcm.cn-shanghai.personal.cr.aliyuncs.com`
   - 这是你的个人实例地址，不同账号地址不同

4. **服务器要求**
   - 确保服务器已安装 Docker
   - 确保 SSH 用户有 Docker 权限
   - 防火墙开放必要端口（如 8080）

---

## 🔍 常见问题

### 1. Docker 登录失败？
- 检查 `ALIYUN_USERNAME` 和 `ALIYUN_PASSWORD` 是否正确
- 确认镜像仓库地址是否与你的实例匹配
- 在阿里云控制台重置 Docker 登录密码

### 2. SSH 连接失败？
- 检查 `SERVER_HOST` 是否可访问
- 确认 SSH 密钥格式正确（包含完整的 BEGIN 和 END 行）
- 检查服务器的 `~/.ssh/authorized_keys` 是否包含对应的公钥
- 验证文件权限是否正确（700 和 600）

### 3. 镜像推送失败？
- 确认命名空间 `ALIYUN_NAMESPACE` 已在阿里云创建
- 检查镜像仓库是否存在，或设置为自动创建

### 4. 容器启动失败？
- SSH 登录服务器，查看容器日志：`docker logs go-app`
- 检查端口 8080 是否被占用：`netstat -tlnp | grep 8080`
- 确认镜像是否拉取成功：`docker images`

---

## 🎉 配置完成

完成以上所有配置后，当你推送代码到 `main` 分支时，GitHub Actions 将自动：

1. ✅ 运行 Go 测试
2. ✅ 构建 Docker 镜像
3. ✅ 推送镜像到阿里云
4. ✅ SSH 登录服务器
5. ✅ 拉取最新镜像
6. ✅ 重启应用容器

享受自动化部署的便利吧！🚀
