package toss

//
//func DeleteFile(keys ...string) ([]string, error) {
//	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
//	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
//	res, err := cache.Bucket.DeleteObjects(keys)
//	if err != nil {
//		return []string{}, err
//	}
//
//	return res.DeletedObjects, err
//}
//
//// 删除文件，通过前缀
//func DeleteFileFromPrefix(prefix string, fileKeys []string) ([]string, error) {
//	resList, err := GetOssFile("", prefix)
//	if err != nil {
//		return nil, err
//	}
//	var keys []string
//	for _, v := range resList {
//		keys = append(keys, v.Key)
//	}
//	return DeleteFile(keys...)
//}
