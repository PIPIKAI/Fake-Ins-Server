package postservice

import (
	"github.com/PIPIKAI/Ins-gin-vue/server/common"
	"github.com/PIPIKAI/Ins-gin-vue/server/model"
)

// type RestSerivce interface {
// 	Create(ctx *gin.Context)
// 	Update(ctx *gin.Context)
// 	Show(ctx *gin.Context)
// 	Delete(ctx *gin.Context)
// }
// type CategoryService struct {
// 	repository repository.CategoryRepository
// }

// func NewCategoryController() CategoryService {
// 	repository := repository.NewCategoryRepository()
// 	return CategoryService{repository}
// }

// type ICategoryController interface {
// 	RestSerivce
// }

func (c CategoryService) CreateOrFind(names []string) []model.Category {
	categorys := make([]model.Category, len(names))
	DB := common.GetDB()
	for idx, name := range names {
		Mc := model.Category{}
		if err := DB.Where("name = ?", name).First(&Mc).Error; err != nil {
			categorys[idx].Name = name
		} else {
			categorys[idx] = Mc
		}

	}
	return categorys
}

// func (c CategoryService) Update(ctx *gin.Context) {
// 	// 绑定body中的参数
// 	var requestCategory vo.CreateCategoryRequest
// 	if ctx.ShouldBind(&requestCategory) != nil {
// 		util.Response.Error(ctx, nil, "数据验证错误，分类名必填")
// 		return
// 	}
// 	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
// 	selectedCategory, err := c.repository.SelectById(categoryId)
// 	if err != nil {
// 		util.Response.Error(ctx, nil, "分类不存在")
// 		panic(err)
// 	}
// 	// 更新
// 	updateCategory, err := c.repository.Update(*selectedCategory, requestCategory.Name)
// 	if err != nil {
// 		util.Response.Error(ctx, nil, "更新失败，已存在该类")
// 		return
// 	}
// 	util.Response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")

// }

// func (c CategoryService) Show(ctx *gin.Context) {
// 	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
// 	selectedCategory, err := c.repository.SelectById(categoryId)
// 	if err != nil {
// 		util.Response.Error(ctx, nil, "分类不存在")
// 		return
// 	}

// 	util.Response.Success(ctx, gin.H{"category": selectedCategory}, "修改成功")
// }

// func (c CategoryService) Delete(ctx *gin.Context) {
// 	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

// 	_, err := c.repository.SelectById(categoryId)
// 	if err != nil {
// 		panic(err)

// 	}
// 	if c.repository.DeleteById(categoryId) != nil {
// 		util.Response.Error(ctx, nil, "删除失败")
// 		return
// 	}
// 	util.Response.Success(ctx, nil, "删除成功")

// }
