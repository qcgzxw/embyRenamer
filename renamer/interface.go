package renamer

type Client interface {
	// GetOriginName 获取原始文件名
	GetOriginName() string
	// GetEmbyDirName 获取文件夹名称
	GetEmbyDirName() string
	// GetEmbyTitleName 获取文件夹名称
	GetEmbyTitleName() string
	// Rename 重命名为规范格式
	Rename()
}
