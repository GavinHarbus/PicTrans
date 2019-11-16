# PicTrans
---

#### About

This is a ZJU AI homework project. This project aims to transfer our pictures to another artistic style such as Monet, Vincent van Gogh... I hope every one of us can enjoy the beauty of art while suffering from the boring life. This may does a little to that aim.  

![](http://45.32.68.50/large/006y8mN6ly1g8wf72u861j32020u04qu.jpg)

#### Requirements

1. Go v1.13.4 or above
2. Go Iris v12.0.1 or above
3. Python3.6
4. ffmpeg

#### Installation

```shell
git clone https://github.com/GavinHarbus/PicTrans.git
``` 

#### Usage

1. install Iris

	```shell
	go get -u github.com/kataras/iris
	```

2. install python packages
	
	```shell
	pip install -r requirements.txt
	```

3. download models

	```shell
	cd PicTrans
	wget http://45.32.68.50/file/model.zip
	unzip model.zip
	```
	
4. start the service

	```shell
	go build picTransController.go
	chmod +x picTransController
	nohup ./picTransController > log &
	```

#### Preview

1. Index
	![](http://45.32.68.50/large/pictransindex.jpg)
2. Styles  
	![](http://45.32.68.50/large/pictransstyle.jpg)
3. Transfer style  
	![](http://45.32.68.50/large/pictranstransfer.jpg)

#### Theory

![](http://45.32.68.50/large/theory.jpg)  

#### Style Transfer

![](http://45.32.68.50/large/scream.gif)

#### References

* [**Johnson J, Alahi A, Fei-Fei L. Perceptual losses for real-time style transfer and super-resolution[C]//European conference on computer vision. Springer, Cham, 2016: 694-711.**](https://link.springer.com/chapter/10.1007/978-3-319-46475-6_43)

#### License

[**MIT**](https://github.com/GavinHarbus/PicTrans/LICENSE)

#### Group Members

<a href="https://github.com/GavinHarbus">
    <img src="https://avatars1.githubusercontent.com/u/23392860?s=460&v=4" width="50px">
</a>
<a href="https://github.com/SuperVivian">
    <img src="https://avatars3.githubusercontent.com/u/29708092?s=400&v=4" width="50px">
</a>
<a href="https://github.com/jojofly">
    <img src="https://avatars3.githubusercontent.com/u/47408058?s=400&v=4" width="50px">
</a>
<a href="https://github.com/pppig1994">
    <img src="https://avatars1.githubusercontent.com/u/46932277?s=400&v=4" width="50px">
</a>
<a href="https://github.com/ls889">
    <img src="https://avatars3.githubusercontent.com/u/53811920?s=400&v=4" width="50px">
</a>
<a href="https://github.com/victor1999">
    <img src="https://avatars2.githubusercontent.com/u/630149?s=400&v=4" width="50px">
</a>

***Thanks for every one of our group!***