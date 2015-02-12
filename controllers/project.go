package controllers

import (
	"fmt"
	"net/http"

	ctx "github.com/jdkanani/smalldocs/context"
	"github.com/jdkanani/smalldocs/models"
	"github.com/jdkanani/smalldocs/utils"

	"labix.org/v2/mgo/bson"
)

//
// Project page
//
func ProjectIndex(context *ctx.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	return 301, nil
}

//
// Project name check
//
func CheckProject(context *ctx.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	var db = context.Config.Get("db.database")

	// get mongodb session
	session := context.DBSession.Copy()
	defer session.Close()

	var data = make(map[string]string)
	context.ReadJson(r, &data)
	title, ok := data["title"]
	if !ok {
		return 412, fmt.Errorf("Title is required")
	}

	name := utils.FormatName(title)
	collection := session.DB(db).C("projects")
	count, err := collection.Find(bson.M{"name": name}).Count()
	if err != nil || count > 0 || name == "" {
		return 403, nil
	}

	// send name
	context.JSON(w, &map[string]string{
		"title": title,
		"name":  name,
	})
	return 200, nil
}

//
// Get all projects
//
func GetAllProjects(context *ctx.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	var db = context.Config.Get("db.database")

	// get mongodb session
	session := context.DBSession.Copy()
	defer session.Close()

	collection := session.DB(db).C("projects")

	var projects []models.Project = make([]models.Project, 0)
	if err := collection.Find(nil).All(&projects); err != nil {
		return 500, err
	}

	return 200, context.JSON(w, &projects)
}

//
// Get project by Id
//
func GetProject(context *ctx.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	var db = context.Config.Get("db.database")

	// get mongodb session
	session := context.DBSession.Copy()
	defer session.Close()

	collection := session.DB(db).C("projects")

	var project models.Project
	if err := collection.FindId("").One(&project); err != nil {
		return 500, err
	}

	return 200, context.JSON(w, project)
}

//
// Create new project
//
func PostProject(context *ctx.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	var db = context.Config.Get("db.database")

	// get mongodb session
	session := context.DBSession.Copy()
	defer session.Close()

	collection := session.DB(db).C("projects")

	var project = new(models.Project)
	if err := context.ReadJson(r, project); err != nil {
		return 500, err
	}

	project.Name = utils.FormatName(project.Title)
	if project.Name == "" {
		return 412, fmt.Errorf("Invalid title for project!")
	}

	if err := collection.Insert(project); err != nil {
		return 500, err
	}

	return 200, context.JSON(w, &project)
}

//
// Save project id
//
func SaveProject(context *ctx.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	var db = context.Config.Get("db.database")

	// get mongodb session
	session := context.DBSession.Copy()
	defer session.Close()

	collection := session.DB(db).C("projects")

	var project = new(models.Project)
	if err := context.ReadJson(r, project); err != nil {
		return 500, err
	}

	project.Name = utils.FormatName(project.Title)
	if project.Name == "" {
		return 412, fmt.Errorf("Invalid title for project!")
	}

	if err := collection.Insert(project); err != nil {
		return 500, err
	}

	return 200, context.JSON(w, &project)
}

//
// Delete project by id
//
func DeleteProject(context *ctx.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	var db = context.Config.Get("db.database")

	// get mongodb session
	session := context.DBSession.Copy()
	defer session.Close()

	collection := session.DB(db).C("projects")

	var projects []models.Project = make([]models.Project, 0)
	if err := collection.Find(nil).All(&projects); err != nil {
		return 500, err
	}

	return 200, context.JSON(w, &projects)
}
