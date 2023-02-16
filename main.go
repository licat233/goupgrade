package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

const (
	// goInstallURL = "https://golang.org/dl/"
	goInstallURL = "http://mirrors.ustc.edu.cn/golang/"
	goVersionURL = "https://golang.org/VERSION?m=text"
)

func main() {
	// 获取当前安装的Go版本
	installedVersion := getInstalledVersion()

	// 获取最新的Go版本
	latestVersion := getLatestVersion()

	// 比较当前版本和最新版本
	if installedVersion == latestVersion {
		fmt.Printf("已安装最新版本 %s\n", installedVersion)
		return
	}

	// 下载并安装最新版本的Go
	downloadAndInstall(latestVersion)

	fmt.Printf("已安装最新版本 %s\n", latestVersion)
}

// getInstalledVersion 获取当前安装的Go版本
func getInstalledVersion() string {
	out, err := exec.Command("go", "version").Output()
	if err != nil {
		return ""
	}
	return string(out)
}

// getLatestVersion 获取最新的Go版本
func getLatestVersion() string {
	resp, err := http.Get(goVersionURL)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	latestVersion, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(latestVersion)
}

// downloadAndInstall 下载并安装最新版本的Go
func downloadAndInstall(version string) error {
	url := fmt.Sprintf("%sgo%s.%s-%s.tar.gz", goInstallURL, version, runtime.GOOS, runtime.GOARCH)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 创建临时文件
	tmpFile, err := os.CreateTemp("", "go.*.tar.gz")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	// 将下载的文件写入临时文件中
	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		return err
	}

	// 解压文件
	untarCmd := exec.Command("tar", "-C", "/usr/local", "-xzf", tmpFile.Name())
	if err := untarCmd.Run(); err != nil {
		return err
	}

	// 更新环境变量
	os.Setenv("PATH", fmt.Sprintf("/usr/local/go/bin:%s", os.Getenv("PATH")))

	return nil
}
