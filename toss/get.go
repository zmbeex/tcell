package toss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func GetOssFile(makerStr string, prefixStr string) ([]oss.ObjectProperties, error) {
	var list []oss.ObjectProperties
	marker := oss.Marker(makerStr)
	prefix := oss.Prefix(prefixStr)
	for {
		lsRes, err := cache.Bucket.ListObjects(oss.MaxKeys(200), marker, prefix)
		if err != nil {
			return nil, err
		}
		marker = oss.Marker(lsRes.NextMarker)
		list = append(list, lsRes.Objects...)
		if !lsRes.IsTruncated {
			break
		}
	}
	return list, nil
}
