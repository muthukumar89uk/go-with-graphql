package graph

import (
	"context"
	"jobPostGraphQl/graph/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Resolver struct {
	DB *gorm.DB
}

func (r *Resolver) CreateUserHandler(g *gin.Context) {
	var user model.NewUser

	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.(validator.ValidationErrors),
		})
		return

	}

	res, err := r.Mutation().CreateUser(context.Background(), user)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusCreated, res)
}

func (r *Resolver) GetAllUsersHandler(g *gin.Context) {
	res, err := r.Query().GetAllCareer(g)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	g.JSON(http.StatusCreated, res)
}

func (r *Resolver) GetOneUserHandler(g *gin.Context) {
	id, check := g.Params.Get("id")
	if !check {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": "error at getting the param",
		})
		return
	}
	userId, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": "error at string to integer conversion",
		})
		return
	}
	res, err := r.Query().GetOneUserByID(g, userId)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	g.JSON(http.StatusCreated, res)
}

func (r *Resolver) UpdateUserHandler(g *gin.Context) {
	var user model.NewUser

	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.(validator.ValidationErrors),
		})
		return

	}

	id, check := g.Params.Get("id")
	if !check {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": "error at getting the param",
		})
		return
	}
	userId, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": "error at string to integer conversion",
		})
		return
	}
	res, err := r.Mutation().UpdateUser(context.Background(), userId, user)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusCreated, res)
}

func (r *Resolver) DeleteUserHandler(g *gin.Context) {
	id, check := g.Params.Get("id")
	if !check {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": "error at getting the param",
		})
		return
	}
	userId, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": "error at string to integer conversion",
		})
		return
	}
	res, err := r.Mutation().DeleteUser(context.Background(), userId)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusCreated, res)
}

func (r *Resolver) CreateProfileHandler(g *gin.Context) {
	var userProfile model.NewProfile

	if err := g.ShouldBindJSON(&userProfile); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(userProfile); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.(validator.ValidationErrors),
		})
		return

	}

	res, err := r.Mutation().CreateProfile(context.Background(), userProfile)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusCreated, res)
}

func (r *Resolver) GetAllProfileHandler(g *gin.Context) {
	res, err := r.Query().GetAllProfile(g)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	g.JSON(http.StatusCreated, res)
}

func (r *Resolver) GetOneProfileHandler(g *gin.Context) {
	id, check := g.Params.Get("id")
	if !check {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": "error at getting the param",
		})
		return
	}
	profileId, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": "error at string to integer conversion",
		})
		return
	}
	res, err := r.Query().GetOneProfile(g, profileId)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	g.JSON(http.StatusCreated, res)
}

func (r *Resolver) UpdateProfileHandler(g *gin.Context) {
	var userProfile model.NewProfile

	if err := g.ShouldBindJSON(&userProfile); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(userProfile); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.(validator.ValidationErrors),
		})
		return

	}

	id, check := g.Params.Get("id")
	if !check {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": "error at getting the param",
		})
		return
	}
	profileId, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": "error at string to integer conversion",
		})
		return
	}
	res, err := r.Mutation().UpdateProfile(context.Background(), profileId, userProfile)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusCreated, res)
}

func (r *Resolver) DeleteProfileHandler(g *gin.Context) {
	id, check := g.Params.Get("id")
	if !check {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": "error at getting the param",
		})
		return
	}
	profileId, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": "error at string to integer conversion",
		})
		return
	}
	res, err := r.Mutation().DeleteProfile(context.Background(), profileId)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusCreated, res)
}

func (r *Resolver) CreateCareerHandler(g *gin.Context) {
	var userCareer model.NewCareer

	if err := g.ShouldBindJSON(&userCareer); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(userCareer); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.(validator.ValidationErrors),
		})
		return

	}

	res, err := r.Mutation().CreateCareer(context.Background(), userCareer)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusCreated, res)
}

func (r *Resolver) GetAllCareerHandler(g *gin.Context) {
	res, err := r.Query().GetAllCareer(g)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	g.JSON(http.StatusCreated, res)
}

func (r *Resolver) GetOneCareerHandler(g *gin.Context) {
	id, check := g.Params.Get("id")
	if !check {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": "error at getting the param",
		})
		return
	}
	careerId, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": "error at string to integer conversion",
		})
		return
	}
	res, err := r.Query().GetOneCareer(g, careerId)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	g.JSON(http.StatusCreated, res)
}

func (r *Resolver) UpdateCareerHandler(g *gin.Context) {
	var userCareer model.NewCareer

	if err := g.ShouldBindJSON(&userCareer); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(userCareer); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.(validator.ValidationErrors),
		})
		return

	}

	id, check := g.Params.Get("id")
	if !check {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": "error at getting the param",
		})
		return
	}
	careerId, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": "error at string to integer conversion",
		})
		return
	}
	res, err := r.Mutation().UpdateCareer(context.Background(), careerId, userCareer)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusCreated, res)
}

func (r *Resolver) DeleteCareerHandler(g *gin.Context) {
	id, check := g.Params.Get("id")
	if !check {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": "error at getting the param",
		})
		return
	}
	careerId, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": "error at string to integer conversion",
		})
		return
	}
	res, err := r.Mutation().DeleteProfile(context.Background(), careerId)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusCreated, res)
}
