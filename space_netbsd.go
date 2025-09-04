//go:build netbsd
// +build netbsd

package rotatelogs

import "golang.org/x/sys/unix"

func GetDiskSize(dir string) (total uint64, avail uint64, err error) {
	fs := unix.Statvfs_t{}
	err = unix.Statvfs(dir, &fs)
	if err != nil {
		return 0, 0, err
	}
	avail = uint64(fs.Bavail) * uint64(fs.Frsize)
	total = uint64(fs.Blocks) * uint64(fs.Frsize)
	return total, avail, nil
}
