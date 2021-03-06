package service

import (
	"category/domain/model"
	"category/domain/repository"
)

type ICategoryDataService interface {
	AddCategory(Category *model.Category) (categoryId int64, err error)
	DeleteCategory(categoryId int64) error
	UpdateCategory(category *model.Category) error
	FindCategoryById(categoryId int64) (category *model.Category, err error)
	FindAllCategory() (category []model.Category, err error)
	FindCategoryByName(categoryName string) (*model.Category, error)
	FindCategoryByLevel(level uint32) ([]model.Category, error)
	FindCategoryByParent(parent int64) ([]model.Category, error)
}

//创建实例
func NewCategoryDataService(CategoryRepository repository.ICategoryRepository) ICategoryDataService {
	return &CategoryDataService{CategoryRepository: CategoryRepository}
}

type CategoryDataService struct {
	CategoryRepository repository.ICategoryRepository
}

//插入用户
func (u *CategoryDataService) AddCategory(category *model.Category) (categoryId int64, err error) {
	return u.CategoryRepository.CreateCategory(category)
}

//删除用户
func (u *CategoryDataService) DeleteCategory(categoryId int64) error {
	return u.CategoryRepository.DeleteCategoryByID(categoryId)
}

func (u *CategoryDataService) UpdateCategory(category *model.Category) error {
	return u.CategoryRepository.UpdateCategory(category)
}

func (u *CategoryDataService) FindCategoryById(categoryId int64) (category *model.Category, err error) {
	return u.CategoryRepository.FindCategoryByID(categoryId)
}

func (u *CategoryDataService) FindAllCategory() (category []model.Category, err error) {
	return u.CategoryRepository.FindAll()
}

func (u *CategoryDataService) FindCategoryByName(categoryName string) (category *model.Category, err error) {
	return u.CategoryRepository.FindCategoryByName(categoryName)
}

func (u *CategoryDataService) FindCategoryByLevel(level uint32) (category []model.Category, err error) {
	return u.CategoryRepository.FindCategoryByLevel(level)
}

func (u *CategoryDataService) FindCategoryByParent(parent int64) (category []model.Category, err error) {
	return u.CategoryRepository.FindCategoryByParent(parent)
}
