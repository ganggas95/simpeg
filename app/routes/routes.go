// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) Pegawai(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Pegawai", args).Url
}

func (_ tApp) Pendidikan(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Pendidikan", args).Url
}

func (_ tApp) Prestasi(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Prestasi", args).Url
}

func (_ tApp) Pengumuman(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Pengumuman", args).Url
}

func (_ tApp) Profile(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Profile", args).Url
}

func (_ tApp) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Login", args).Url
}

func (_ tApp) Auth(
		username string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("App.Auth", args).Url
}

func (_ tApp) Register(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Register", args).Url
}

func (_ tApp) AddUser(
		user interface{},
		password string,
		verifyPass string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "verifyPass", verifyPass)
	return revel.MainRouter.Reverse("App.AddUser", args).Url
}

func (_ tApp) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Logout", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


