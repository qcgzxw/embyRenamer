package renamer

type Client interface {
	// GetOriginName 获取原始文件名
	GetOriginName() string
	// GetEmbyDirName 获取重命名文件夹名称
	GetEmbyDirName() string
	// GetEmbyTitleName 获取重命名文件名称
	GetEmbyTitleName() string
	// Rename 重命名为规范格式
	Rename()
}
