package controllers

import (
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"simpeg/app/models"
	"simpeg/app/routes"
)

type App struct {
	*revel.Controller
	Database
}

func (c App) connected() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	if username, ok := c.Session["user"]; ok {
		return c.GetUser(username)
	}
	return nil
}

func (c App) GetUser(username string) *models.User {
	var user models.User
	err := c.Database.initDb().Debug().Where("username = ? OR email = ?", username, username).Find(&user)
	CheckError(err.Error)
	return &user
}
func (c App) GetUserWitId(idUser int64) *models.User {
	var users models.User
	err := c.initDb().First(&users, idUser)
	CheckError(err.Error)
	return &users
}

func (p App) GetPegawaiWithUser(userid int64) *models.Pegawai {
	var pegawai models.Pegawai
	db := p.initDb().Where("id_user = ?", userid).Find(&pegawai)
	if db.RecordNotFound() {
		return nil
	}
	return &pegawai
}

func (p App) GetJabatanPegawai(pegawaiId int64) *models.Jabatan {
	var jabatan models.Jabatan
	db := p.initDb().Where("id_pegawai = ?", pegawaiId).Find(&jabatan)
	if db.RecordNotFound() {
		return nil
	}
	return &jabatan
}

func (c App) Index() revel.Result {
	user := c.connected()
	if user != nil {
		c.RenderArgs["user"] = user
		return c.Render(user)
	}
	c.Flash.Error("Anda tidak memiliki izin untuk mengakses sistem")
	return c.Redirect(routes.App.Login())
}

func (c App) Pegawai() revel.Result {
	user := c.connected()
	if user != nil {
		c.RenderArgs["user"] = user
		return c.Render(user)
	}
	c.Flash.Error("Anda tidak memiliki izin untuk mengakses sistem")
	return c.Redirect(routes.App.Login())
}

func (c App) Pendidikan() revel.Result {
	user := c.connected()
	if user != nil {
		c.RenderArgs["user"] = user
		return c.Render(user)
	}
	c.Flash.Error("Anda tidak memiliki izin untuk mengakses sistem")
	return c.Redirect(routes.App.Login())
}

func (c App) Prestasi() revel.Result {
	user := c.connected()
	if user != nil {
		c.RenderArgs["user"] = user
		return c.Render(user)
	}
	c.Flash.Error("Anda tidak memiliki izin untuk mengakses sistem")
	return c.Redirect(routes.App.Login())
}
func (c App) Pengumuman() revel.Result {
	user := c.connected()
	if user != nil {
		c.RenderArgs["user"] = user
		return c.Render(user)
	}
	c.Flash.Error("Anda tidak memiliki izin untuk mengakses sistem")
	return c.Redirect(routes.App.Login())
}
func (c App) Profile() revel.Result {
	user := c.connected()
	if user != nil {
		c.RenderArgs["user"] = user
		return c.Render(user)
	}
	c.Flash.Error("Anda tidak memiliki izin untuk mengakses sistem")
	return c.Redirect(routes.App.Login())
}

func (c App) Login() revel.Result {
	user := c.connected()
	if user != nil {
		return c.Redirect(routes.App.Index())
	}
	return c.Render()
}

func (c App) Auth(username, password string) revel.Result {
	user := c.GetUser(username)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
		if err == nil {
			c.Session["user"] = user.Username
			c.Flash.Success("Selamat Datang " + user.Username)
			c.RenderArgs["user"] = username
			return c.Redirect(routes.App.Index())
		}
	}
	c.Flash.Out["username"] = username
	c.Flash.Error("Login gagal")
	return c.Redirect(routes.App.Login())
}

func (c App) Register() revel.Result {
	user := c.connected()
	if user != nil {
		return c.Redirect(routes.App.Index())
	}
	return c.Render()
}

func (c App) AddUser(user models.User, password, verifyPass string) revel.Result {
	c.Validation.Required(password)
	c.Validation.Required(verifyPass)
	c.Validation.Required(password == verifyPass).Message("Password tidak sama")
	user.Password, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Validasi(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Register())
	}
	db := c.initDb().Create(&user)
	if db.Error != nil {
		panic(db.Error)
	}
	c.Flash.Success("User berhasil di buat")
	return c.Redirect(routes.App.Login())
}

func (a App) Logout() revel.Result {
	if user := a.connected(); &user == nil {
		a.Flash.Error("Login terlebih dahulu")
		return a.Redirect(routes.App.Login())
	}

	for k := range a.Session {
		delete(a.Session, k)
	}
	return a.Redirect(routes.App.Login())
}
