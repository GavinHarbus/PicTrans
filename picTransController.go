package main

import (
    "io"
    "os"
    "os/exec"
    "time"

    "github.com/kataras/iris"
    "github.com/kataras/iris/middleware/logger"
    "github.com/kataras/iris/middleware/recover"
)

func main() {
    app := iris.New()
    app.Use(recover.New())
    app.Use(logger.New())

    templates := iris.HTML("./views",".html")
    templates.Reload(true)
    /*templates.AddFunc("displayMessage", func (message string) string {
        return message
    })*/

    /*fileserver := iris.FileServer("./static")
    h := iris.StripPrefix("/static", fileserver)
    app.Get("/static/{f:path}", h)*/

    app.RegisterView(templates)
    app.StaticWeb("/static","./static")

    app.Get("/", func (ctx iris.Context) {
        ctx.ViewData("rawpath","/static/pics/20190710155351.jpg")
        ctx.ViewData("respath","/static/pics/res20190710155351.jpg")
        ctx.ViewData("message","Please upload your picture!")
        if err := ctx.View("index.html"); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString(err.Error())
        }
    })

    app.Post("upload", func (ctx iris.Context) {
        file, _, err := ctx.FormFile("upload")
        style := ctx.FormValue("style")

        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.ViewData("message","Upload failed!")
            ctx.View("index.html")
            return
        }

        defer file.Close()

        //filename := info.Filename
        filename := time.Now().Format("20060102150405")

        out, err := os.OpenFile("./static/pics/"+string(filename)+".jpg", os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.ViewData("message","Save failed!")
            ctx.View("index.html")
            return
        }
        defer out.Close()

        io.Copy(out,file)

        var cmd *exec.Cmd
        if style == "0" {
            cmd = exec.Command("./pictrans.py", "--input", filename)
        } else if style == "1" {
            cmd = exec.Command("./evaluate.py", "--checkpoint", "model/la_muse.ckpt", 
                "--in-path", "./static/pics/"+filename+".jpg", 
                "--out-pat", "./static/pics/res"+filename+".jpg")
        } else if style == "2" {
            cmd = exec.Command("./evaluate.py", "--checkpoint", "model/rain_princess.ckpt", 
                "--in-path", "./static/pics/"+filename+".jpg", 
                "--out-pat", "./static/pics/res"+filename+".jpg")
        } else if style == "3" {
            cmd = exec.Command("./evaluate.py", "--checkpoint", "model/scream.ckpt", 
                "--in-path", "./static/pics/"+filename+".jpg", 
                "--out-pat", "./static/pics/res"+filename+".jpg")
        } else if style == "4" {
            cmd = exec.Command("./evaluate.py", "--checkpoint", "model/udnie.ckpt", 
                "--in-path", "./static/pics/"+filename+".jpg", 
                "--out-pat", "./static/pics/res"+filename+".jpg")
        } else if style == "5" {
            cmd = exec.Command("./evaluate.py", "--checkpoint", "model/wave.ckpt", 
                "--in-path", "./static/pics/"+filename+".jpg", 
                "--out-pat", "./static/pics/res"+filename+".jpg")
        } else if style == "6" {
            cmd = exec.Command("./evaluate.py", "--checkpoint", "model/wreck.ckpt", 
                "--in-path", "./static/pics/"+filename+".jpg", 
                "--out-pat", "./static/pics/res"+filename+".jpg")
        } else if style == "7" {
            cmd = exec.Command("./evaluate.py", "--checkpoint", "model/sunset", 
                "--in-path", "./static/pics/"+filename+".jpg", 
                "--out-pat", "./static/pics/res"+filename+".jpg")
        } else if style == "8" {
            cmd = exec.Command("./evaluate.py", "--checkpoint", "model/star", 
                "--in-path", "./static/pics/"+filename+".jpg", 
                "--out-pat", "./static/pics/res"+filename+".jpg")
        } else if style == "9" {
            cmd = exec.Command("./evaluate.py", "--checkpoint", "model/city", 
                "--in-path", "./static/pics/"+filename+".jpg", 
                "--out-pat", "./static/pics/res"+filename+".jpg")
        }
        
        _, err = cmd.Output()

        if err != nil {
            ctx.ViewData("message","Transform failed!")
            ctx.View("index.html")
            return
        }
        ctx.ViewData("filename",filename+".jpg")
        ctx.ViewData("rawpath","/static/pics/"+filename+".jpg")
        ctx.ViewData("respath","/static/pics/res"+filename+".jpg")
        ctx.ViewData("message","Upload and transform success!")
        ctx.View("index.html")
    })

    app.Run(iris.Addr(":8080"))
}