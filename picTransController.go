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

    fileserver := iris.FileServer("./static")
    h := iris.StripPrefix("/static", fileserver)
    app.Get("/static/{f:path}", h)

    app.RegisterView(templates)
    //app.StaticWeb("/static","./static")

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

        cmd := buildCmd(style, filename)
        
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

func buildCmd(style string, filename string) *exec.Cmd {
    switch style {
    case "0":
        return exec.Command("python3", "./pictrans.py", "--input", filename)
    default:
        model := modelMap[style]
        return exec.Command("python3", "./evaluate.py", "--checkpoint", model, 
        "--in-path", "./static/pics/"+filename+".jpg", 
        "--out-pat", "./static/pics/res"+filename+".jpg")
    }
}

var modelMap = map[string]string{
    "1" : "model/la_muse.ckpt",
    "2" : "model/rain_princess.ckpt",
    "3" : "model/scream.ckpt",
    "4" : "model/udnie.ckpt",
    "5" : "model/wave.ckpt",
    "6" : "model/sunset",
    "7" : "model/star",
    "8" : "model/city",
    "9" : "model/monet",
    "10" : "model/vastland",
}