package explorer

import (
	"context"
	"github.com/cloudreve/Cloudreve/v3/pkg/util"
	"strings"

	model "github.com/cloudreve/Cloudreve/v3/models"
	"github.com/cloudreve/Cloudreve/v3/pkg/filesystem"
	"github.com/cloudreve/Cloudreve/v3/pkg/hashid"
	"github.com/cloudreve/Cloudreve/v3/pkg/serializer"
	"github.com/gin-gonic/gin"
)

// ItemSearchService 文件搜索服务
type ItemSearchService struct {
	Type     string `uri:"type" binding:"required"`
	Keywords string `uri:"keywords" binding:"required"`
}

type OtherSearchFileReq struct {
	Type string   `json:"type"`
	MD5  []string `json:"md_5"`
}

func (f *OtherSearchFileReq) Validate() (bool, string) {
	if len(f.Type) <= 0 {
		return false, "type is empty"
	}
	if len(f.MD5) <= 0 {
		return false, "MD5 is empty"
	}

	return true, ""
}

type OtherSearchFileResult struct {
	Name       string `json:"name"`
	SourcePath string `json:"source_path"`
	Size       uint64 `json:"size"`
	PolicyId   uint   `json:"policy_id"`
	MD5        string `json:"md5"`
}

// Search 执行搜索
func (s *OtherSearchFileReq) Search(c *gin.Context, req *OtherSearchFileReq) serializer.Response {

	if ok, msg := req.Validate(); !ok {
		util.Log().Warning("OtherSearchFileReq validate fail: %v", msg)
		return serializer.ParamErr(msg, nil)
	}
	// 创建文件系统
	fs, err := filesystem.NewFileSystemFromContext(c)
	if err != nil {
		return serializer.Err(serializer.CodePolicyNotAllowed, err.Error(), err)
	}
	defer fs.Recycle()

	switch s.Type {
	case "md5":
		f := model.File{}
		getFiles, err := f.GetFilesByMD5(fs.User.ID, req.MD5)
		if err != nil {
			util.Log().Error("GetFilesByMD5 failed: ", err)
			return serializer.Err(serializer.CodeDBError, err.Error(), err)
		}
		var res []*OtherSearchFileResult
		for _, v := range getFiles {
			res = append(res, &OtherSearchFileResult{
				Name:       v.Name,
				SourcePath: v.SourceName,
				Size:       v.Size,
				PolicyId:   v.PolicyID,
				MD5:        v.MD5,
			})
		}
		return serializer.Response{
			Data: res,
		}
	default:
		return serializer.ParamErr("未知搜索类型", nil)
	}

}

// Search 执行搜索
func (service *ItemSearchService) Search(c *gin.Context) serializer.Response {
	// 创建文件系统
	fs, err := filesystem.NewFileSystemFromContext(c)
	if err != nil {
		return serializer.Err(serializer.CodePolicyNotAllowed, err.Error(), err)
	}
	defer fs.Recycle()

	switch service.Type {
	case "keywords":
		return service.SearchKeywords(c, fs, "%"+service.Keywords+"%")
	case "image":
		return service.SearchKeywords(c, fs, "%.bmp", "%.iff", "%.png", "%.gif", "%.jpg", "%.jpeg", "%.psd", "%.svg", "%.webp")
	case "video":
		return service.SearchKeywords(c, fs, "%.mp4", "%.flv", "%.avi", "%.wmv", "%.mkv", "%.rm", "%.rmvb", "%.mov", "%.ogv")
	case "audio":
		return service.SearchKeywords(c, fs, "%.mp3", "%.flac", "%.ape", "%.wav", "%.acc", "%.ogg", "%.midi", "%.mid")
	case "doc":
		return service.SearchKeywords(c, fs, "%.txt", "%.md", "%.pdf", "%.doc", "%.docx", "%.ppt", "%.pptx", "%.xls", "%.xlsx", "%.pub")
	case "tag":
		if tid, err := hashid.DecodeHashID(service.Keywords, hashid.TagID); err == nil {
			if tag, err := model.GetTagsByID(tid, fs.User.ID); err == nil {
				if tag.Type == model.FileTagType {
					exp := strings.Split(tag.Expression, "\n")
					expInput := make([]interface{}, len(exp))
					for i := 0; i < len(exp); i++ {
						expInput[i] = exp[i]
					}
					return service.SearchKeywords(c, fs, expInput...)
				}
			}
		}
		return serializer.Err(serializer.CodeNotFound, "标签不存在", nil)
	default:
		return serializer.ParamErr("未知搜索类型", nil)
	}
}

// SearchKeywords 根据关键字搜索文件
func (service *ItemSearchService) SearchKeywords(c *gin.Context, fs *filesystem.FileSystem, keywords ...interface{}) serializer.Response {
	// 上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 获取子项目
	objects, err := fs.Search(ctx, keywords...)
	if err != nil {
		return serializer.Err(serializer.CodeNotSet, err.Error(), err)
	}

	return serializer.Response{
		Code: 0,
		Data: map[string]interface{}{
			"parent":  0,
			"objects": objects,
		},
	}
}
